jQuery.fn.outerHtml = function(s) {
  return s
    ? this.before(s).remove()
    : jQuery("<p>").append(this.eq(0).clone()).html();
};

function first(obj) {
    for (var a in obj) return obj[a];
}

$(function(){
  var commitsDiv = $('#commits');

  function linkToCommit(commit) {
    return $('<a>').prop('href', commit['html_url']).text(commit['sha'].substring(0,8)).outerHtml();
  }

  var commitsByDay = {},
      limit = moment().subtract(7, 'days').format('YYYY-MM-DD');
      bucketCommits = function(page, cb) {
        var limitReached = false;

        if (page > 4) {
          arst; // safety precaution
        }

        $.getJSON('/commits' , {page: page}, function(data) {
          console.log(page, data);
          data.forEach(function(commit,i) {
            var date = moment(commit['commit']['committer']['date']), // github shows author date, we use committer date
                dateKey = date.format('YYYY-MM-DD');

            if (dateKey >= limit) {
              if (!(dateKey in commitsByDay)) {
                commitsByDay[dateKey] = {};
              }
              commitsByDay[dateKey][date.format('x')] = commit;
            } else {
              limitReached = true;
            }
          });

          if (limitReached) {
            cb();
          } else {
            bucketCommits(page+1);
          }
        });
      };

  bucketCommits(1, function() {
    Object.keys(commitsByDay).sort().reverse().forEach(function(dateKey) {
      $('<h2>')
      .text(moment(first(commitsByDay[dateKey])['commit']['committer']['date']).format('dddd'))
      .appendTo(commitsDiv);

      Object.keys(commitsByDay[dateKey]).sort().reverse().forEach(function(commitKey) {
        var commit = commitsByDay[dateKey][commitKey],
            newlinePos = commit['commit']['message'].indexOf("\n"),
            shortMsg = commit['commit']['message'].substr(0, newlinePos > 0 ? newlinePos : 9999);

        // console.log(commit);

        $('<div>')
        .addClass('commit')
        .append($('<div>').addClass('link').html(linkToCommit(commit)))
        .append($('<div>').addClass('date').text(moment(commit['commit']['committer']['date']).format('HH:mm')))
        .append($('<div>').addClass('author').html($('<img>').prop('src', commit['author']['avatar_url']).prop('title', commit['commit']['author']['name'])))
        .append($('<div>').addClass('message').text(shortMsg))
        .appendTo(commitsDiv);
      });
    });
  });
});