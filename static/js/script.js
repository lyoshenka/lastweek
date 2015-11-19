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
      diffstatTmpl = Handlebars.compile($('#diffstatTmpl').html()),
      commitsByDay = {},
      limit = moment().subtract(7, 'days').format('YYYY-MM-DD');

  var diffstat = function (sha) {
    var storedDiff = store.get('diffstat_'+sha),
        updateDiffstat = function(sha, diffstat) {
          $('#' + sha + ' .diffstat').html(diffstatTmpl({
            additions: diffstat.stats.additions,
            deletions: diffstat.stats.deletions,
            files_changed: diffstat.files_changed,
            files_plural: diffstat.files_changed != 1
          }));

          if (diffstat.stats.total > 50) {
            $('#' + sha).addClass('major');
          }
        };

    if (storedDiff && storedDiff.stats) {
      updateDiffstat(sha, storedDiff);
    } else {
      $.getJSON('/diffstat' , {sha: sha}, function(data) {
        store.set('diffstat_'+sha, data);
        updateDiffstat(sha, data);
      });
    }
  };

  var bucketCommits = function(page, cb) {
        var limitReached = false;

        if (page > 5) {
          return; // safety precaution
        }

        $.getJSON('/commits' , {page: page}, function(data) {
          console.log(page, data);
          data.forEach(function(commit,i) {
            var date = moment(commit['commit']['committer']['date']), // github shows author date, we use committer date
                dateKey = date.format('YYYY-MM-DD');

            if (dateKey >= limit) {
              if (!(dateKey in commitsByDay)) {
                commitsByDay[dateKey] = [];
              }
              commitsByDay[dateKey].push(commit);
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

      commitsByDay[dateKey].forEach(function(commit) {
        var newlinePos = commit['commit']['message'].indexOf("\n"),
            shortMsg = newlinePos > 0 ? commit['commit']['message'].substr(0, newlinePos) : commit['commit']['message'],
            restOfMsg = newlinePos > 0 ? commit['commit']['message'].substr(newlinePos+1) : "",
            commitSha = commit['sha'].substring(0,8);

        // console.log(commit);

        commitsDiv.append(commitTmpl({
          commit_url: commit['html_url'],
          commit_sha: commitSha,
          date: moment(commit['commit']['committer']['date']).format('HH:mm'),
          committer_img: commit['author']['avatar_url'],
          committer_name: commit['commit']['author']['name'],
          short_msg: shortMsg.trim(),
          long_msg: restOfMsg.trim()
        }));

        diffstat(commitSha);
      });
    });
  });
});