## December 01

This is the first entry of this log. A colleague of mine sent me a link to the [adventofcode](https://adventofcode.com/)
webpage. This seems to be a fun and entertaining way of writing code.

Today's exercise doesn't look difficult. First approach would be to think of recursion, but the dataset isn't that big:
1000 lines - a brutal approach would cost O(n²), which is tolerable here.

I've written a `for`/`for`/`if` func I'm not proud of. I've got one big main.go file that performs several tasks -
reading from the disk, computing stuff, printing the solution...

Currently, `main.go` dwells inside `2020/01`, but that's not really what we'll want in the end, is it? My implementation
choices so far have been limited to using `go mod`, which isn't even a choice, as it's recommended since Go 1.13.

So, I've finished Part 1. Let's have a look at Part 2.

Yeah. Well. So much for recursion. Still, I don't think it's worth it. I'd rather improve my forforif horror. I found a
way of reducing my O(n²) down to O(n), which is something. I'm not going to generate a 1-million-lines file just for the
fun of watching one implementation run faster than the other. I'm sure these exercises will provide such opportunities.

At the end of December 01, I'm missing lots of stuff. I'm missing lots of unit tests. I'm missing some genericity. I'm
missing a nice Makefile.
