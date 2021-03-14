## December 03

Morning thoughts - maybe I can use a method instead of a function in that `ReadLines()` - a method that can store
results internally in the object.

Added support for Parts. And then I realised I really needed to change some things to support "Part1 but not Part2"
resulting in something different than "Part1 and Part2 returned 0". So I used pointers, and then I had to point to
numbers, and I finally added a `common` package for that kind of stuff.

After quite some refactoring on the parsing of the arguments, etc. I started working on the solution for Part1. It's not
that tricky, we only have to make a few checks, and use the modulo operator. The code I'm writing is becoming more and
more similar, I've refactored my `decXX.go` files so that they all share the same mechanics.

I've pondered a bit, and decided that parsing each line to determine where trees are isn't interesting, because, in the
end, all that matters is one character per line. So I'll keep them as lines of strings so far. However, I could use
slices of bytes rather than strings, for each line. That'd make accessing the bits a bit cleaner. And I'd use a byte to
store the const tree character, which is better than a string.

Part2 seems to want to re-use Part1 (as always), and also seems to re-user the `product([]int) int` func we needed for
Day1. Maybe that one'll join the `common` package. I've written a short version that computes the result. I don't know
yet whether I want to make that version more generic (I didn't call `product`, for instance). But these vectors seem
quite hardcoded to me.

The only question open for Day03, for me, would be "can we compute the result with fewer checks?". Currently I parse the
file 5 times (once per toboggan slide). But I could first determine where trees stand, and then loop on that, and
perform a check (with modulo, this time) for each check we need to do. The benefit, here, would be that I wouldn't have
to cast a tree twice to a string. However, I (well, not _really_ me, but the CPU) would have to perform some maths (
check if a point is on a line, basically), and benchmarks would be necessary to know which option is better.

Oh, I've also broken the Makefile. I have to check how to pass optional parameters to a command from a make target. To
validate this day, I had to `make build && ./adventofcode.out -parts 1,2 -door 3`. This last option was mandatory, as I
finished a tad too late.
