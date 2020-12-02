_This file serves as a changelog and a blog. Don't expect anything fancy from it. I've decided to write it downwards, rather than upwards, so you don't get spoilt the results of whatever happens on Decembre 24th._

## December 01
This is the first entry of this log. A colleague of mine sent me a link to the [adventofcode](https://adventofcode.com/) webpage.
This seems to be a fun and entertaining way of writing code.

Today's exercise doesn't look difficult. First approach would be to think of recursion, but the dataset isn't that big: 1000 lines - a brutal approach would cost O(n²), which is tolerable here.
 
I've written a `for`/`for`/`if` func I'm not proud of. I've got one big main.go file that performs several tasks - reading
from the disk, computing stuff, printing the solution...

Currently, `main.go` dwells inside `2020/01`, but that's not really what we'll want in the end, is it?
My implementation choices so far have been limited to using `go mod`, which isn't even a choice, as it's recommended since Go 1.13.

So, I've finished Part 1. Let's have a look at Part 2.

Yeah. Well. So much for recursion. Still, I don't think it's worth it. I'd rather improve my forforif horror. I found a way of reducing my O(n²) down to O(n), which is something. I'm not going to generate a 1-million-lines file just for the fun of watching one implementation run faster than the other. I'm sure these exercises will provide such opportunities. 

At the end of December 01, I'm missing lots of stuff. I'm missing lots of unit tests. I'm missing some genericity. I'm missing a nice Makefile.

## December 02
I've reworked the Makefile, moved the main.go, added support for a `-door #` flag thanks to the help of a precious friend. I now have a `door` package with nice interfaces in it that I'll re-use on a day-by-day basis. I had to rewrite the flesh (but not the bones) of the December 01, but that's fine, as it provided the opportunity to write some unit tests over there.

December 2nd's door was about having fun with regexps. Part 1 required the same implementation as I needed for `dec01` - parse a file, line by line, and extract an object from each line. Clearly, I want to write a func that goes in the lines of 
```go
// ReadLines reads a file and returns a slice of objects built from its lines.
//  Or an error, if something went awry.
func ReadLines(inputFilePath string, objectBuilder func(line string)(interface{}, error)) ([]interface{}, error)
```
but that's not (yet) available. Go Generics would solve this for me.

Anyways. Part 1 wasn't very difficult, but I had to make sure I was calling the right `regexp` func. I wonder how much effort I should put in testing against invalid input files. Obviously, the website is expecting an integer anser, therefore it wouldn't generate bad stuff, right?

I read the rest of Part2 too fast, and I submitted a wrong answer (the website tells you it's wrong, and whether it's too big or too small - and you can't re-submit an answer immediately, of course). Re-reading (or at least calmly reading) the input text solved the problem and provided the answer.

I haven't written the UTs yet. I want to test my exposed stuff, not the inner guts. I also want to find a way to not change my code once I'm done with Part 1 and start coding on Part 2. I suppose having a `door.Solve(parts []int) (door.Result, error)` func would do the trick. This also calls for a `common` package with some stuff in it (at least "is my int in this slice ?)

## December 03
