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
but that's not (yet) available. Go Generics would solve this for me with this:
```go
func ReadLines(type Element) (inputFilePath string, objectBuilder func(line string)(Element, error)) ([]Element, error)
```

Anyways. Part 1 wasn't very difficult, but I had to make sure I was calling the right `regexp` func. I wonder how much effort I should put in testing against invalid input files. Obviously, the website is expecting an integer anser, therefore it wouldn't generate bad stuff, right?

I read the rest of Part2 too fast, and I submitted a wrong answer (the website tells you it's wrong, and whether it's too big or too small - and you can't re-submit an answer immediately, of course). Re-reading (or at least calmly reading) the input text solved the problem and provided the answer.

I haven't written the UTs yet. I want to test my exposed stuff, not the inner guts. I also want to find a way to not change my code once I'm done with Part 1 and start coding on Part 2. I suppose having a `door.Solve(parts []int) (door.Result, error)` func would do the trick. This also calls for a `common` package with some stuff in it (at least "is my int in this slice ?)

## December 03
Morning thoughts - maybe I can use a method instead of a function in that `ReadLines()` - a method that can store results internally in the object.

Added support for Parts. And then I realised I really needed to change some things to support "Part1 but not Part2" resulting in something different than "Part1 and Part2 returned 0". So I used pointers, and then I had to point to numbers, and I finally added a `common` package for that kind of stuff.

After quite some refactoring on the parsing of the arguments, etc. I started working on the solution for Part1. It's not that tricky, we only have to make a few checks, and use the modulo operator. The code I'm writing is becoming more and more similar, I've refactored my `decXX.go` files so that they all share the same mechanics.

I've pondered a bit, and decided that parsing each line to determine where trees are isn't interesting, because, in the end, all that matters is one character per line. So I'll keep them as lines of strings so far. However, I could use slices of bytes rather than strings, for each line. That'd make accessing the bits a bit cleaner. And I'd use a byte to store the const tree character, which is better than a string.

Part2 seems to want to re-use Part1 (as always), and also seems to re-user the `product([]int) int` func we needed for Day1. Maybe that one'll join the `common` package. I've written a short version that computes the result. I don't know yet whether I want to make that version more generic (I didn't call `product`, for instance). But these vectors seem quite hardcoded to me.

The only question open for Day03, for me, would be "can we compute the result with fewer checks?". Currently I parse the file 5 times (once per toboggan slide). But I could first determine where trees stand, and then loop on that, and perform a check (with modulo, this time) for each check we need to do. The benefit, here, would be that I wouldn't have to cast a tree twice to a string. However, I (well, not _really_ me, but the CPU) would have to perform some maths (check if a point is on a line, basically), and benchmarks would be necessary to know which option is better.
 
 Oh, I've also broken the Makefile. I have to check how to pass optional parameters to a command from a make target. To validate this day, I had to `make build && ./adventofcode.out -parts 1,2 -door 3`. This last option was mandatory, as I finished a tad too late.
 
 