## December 14

It's been 19 days since I last was connected to the Interwebs. Let's face it, I've idly spent my time. I had the option of refactoring everything I'd done so far, but then this whole repo wouldn't look like something I'd code on my spare time. Once more unto the breach it is, then. This whole calendar makes me think wherever an issue appears, a software engineer is suited to fix it. Clearly not the right approach to life. Planes can't fly, captains can't helm, buses go amok, what have you. Why can't people leave me alone on my vacation? I'm sure at some point my shoelace will be loose and I'll need to implement a minimax algorithm to decide whether the loop should be done using the right hand or the left.

This time, the ferry can't dock. What kind of a service do these people provide!

We'll note that the first bit of the bitmask is actually that of the highest index. This time - again? my memory starts to become faulty after 3 weeks of not looking at all at this code; I'm pretty sure we've already dealt with them - bits manipulations are in order.

The input file has two kind of lines:
- those that start with mask: these reset the value of the mask;
- those that start with mem: these set the value of an entry.

Using some dark voodoo magic (`awk -F'[\\[\\]]' 'BEGIN{min=10000000; max=0}/mem/{if($2<min){min=$2}; if ($2>max){max=$2}}END{print "min: " min " max: " max}' 2020/dec14/input`) we notice that the range of memory addresses remains in the uint16 range. Let's nonetheless use a uint64 as the description reads memory addresses are in a 36-bit range. Also, we notice the first line seems to always be a mask. But just to be sure, I've written a little function that would take care of unset masks.

There's nothing much to say about this exercise - at least about Part1, so far. The highlight of my work might be that LUT I decided to use to avoid computing powers of 2 all the time.

Haven't had a glance at Part2, but sleep is due.

Part2 had me recode quite a lot of my work in order to have inheritance. I'd've expected the output for Part2 to be way bigger - at least bigger than Part1.

A dear friend recommended me to benchmark my LUT. Turns out it's really not efficient.
```bash
[~/go/src/github.com/floppyzedolfin/adventofcode/2020/dec14/mask][adventofcode][dec14]
$ go test -bench=. -benchtime=5s
goos: linux
goarch: amd64
pkg: github.com/floppyzedolfin/adventofcode/2020/dec14/mask
BenchmarkPowersOf2Array-8       461666686               12.9 ns/op
BenchmarkPowerOf2Function-8     1000000000               0.557 ns/op
PASS
ok      github.com/floppyzedolfin/adventofcode/2020/dec14/mask  7.910s
```
Using the function is 24 times faster than using the LUT. So much for that shot at optimization... 
