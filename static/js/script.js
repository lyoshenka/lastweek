jQuery.fn.outerHtml = function(s) {
  return s
    ? this.before(s).remove()
    : jQuery("<p>").append(this.eq(0).clone()).html();
};

function first(obj) {
    for (var a in obj) return obj[a];
}

$(function(){
  var commitsDiv = $('#commits'),
      commitTmpl = Handlebars.compile($('#commitTmpl').html()),
      commitsByDay = {},
      limit = moment().subtract(7, 'days').format('YYYY-MM-DD');

  // Mustache.parse(commitTmpl);

  var bucketCommits = function(page, cb) {
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
            shortMsg = newlinePos > 0 ? commit['commit']['message'].substr(0, newlinePos) : commit['commit']['message'],
            restOfMsg = newlinePos > 0 ? commit['commit']['message'].substr(newlinePos+1) : "";

        // console.log(commit);

        commitsDiv.append(commitTmpl({
          commit_url: commit['html_url'],
          commit_sha: commit['sha'].substring(0,8),
          date: moment(commit['commit']['committer']['date']).format('HH:mm'),
          committer_img: commit['author']['avatar_url'],
          committer_name: commit['commit']['author']['name'],
          short_msg: shortMsg.trim(),
          long_msg: restOfMsg.trim()
        }));

      });
    });
  });
});