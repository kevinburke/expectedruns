# Expected Runs

The goal of this is to get something you can check at the game to show the
team's chances of scoring based on the current situation. Also, to show people
that sacrifice bunts early in the game are dumb, but that's a separate story.

## Usage

Go to [expectedruns.appspot.com](http://expectedruns.appspot.com). Educate your
fellow fans

## Install

The backend of the application was built using [Go][golang], a relatively new
programming language from Google. [Here are instructions on installing
Go][install].

If you'd like to develop this locally, check out the [Google App Engine Go
Runtime][appengine]. Clone this repository on your computer, then run:

    dev_appserver.py /path/to/this/repo

You should be able to browse to [localhost:8080](http://localhost:8080) and
view your application.

## Contributing

Contributions are welcome! If you can't get Go running but want to help out
with CSS/JS/HTML, just send a diff and I'll try to get it to work. Here are
instructions on [how to submit pull requests in Github][pulls]

Here are some feature ideas:

- Show expected runs in real-time, based on run situations for each team in
  each game, maybe at a new URL like /score, or for each team in the MLB, like
  /oak
- Show the cost of various actions, if they make sense and we can get the data 
  (Steal, sacrifice fly, suicide squeeze, sacrifice bunt)
- Update/customize the data based on teams, pitcher/batter matchup or batting order
- A better/more compressed visual layout

[appengine]: https://developers.google.com/appengine/docs/go/gettingstarted/
[install]: http://golang.org/doc/install
[golang]: http://golang.org/
[pulls]: https://help.github.com/articles/using-pull-requests
