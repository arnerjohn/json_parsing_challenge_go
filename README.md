JsonParsingChallengeGo
=====================

This was a fun exercise that was listed by Jorin Vogel in a gist: ["https://gist.github.com/jorin-vogel/2e43ffa981a97bc17259"](JSON Challenge). I think that parsing data is a great programming exercise. 

I first did the challenge in `Elixir`. This is the `Go` version. I have spent more time in `Elixir` and `Erlang` than I have in `Go`.
Thus, I'm sure there are better ways to write the `Go` code. 

One of the interesting differences between the two languages, beside their vastly different programming paradigms, was the dependencies.
In `Elixir`, I had to use a number of libraries that are not part of the standard library. Unfortunately, these are libraries that I consider
necessary for nearly anything done on the web these days (e.g. JSON parsing, HTTP requests, CSV read/write). In `Go`, this was not a problem.
The standard library in `Go` is quite deep. The standard library has all of the major dependencies built-in. In my book, this is a huge positive.

## Running the Program

```
~$ git clone git@github.com:arnerjohn/json_parsing_challenge_go.git
~$ cd json_parsing_challenge_go/

~$ go run json_challenge.go 
```

This will output a CSV file with the days date as the filename.
