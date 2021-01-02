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
 
 ## December 04
 I've reworked the Makefile. I can now run `make run -- -door 3 -parts 1,2`. I've also re-done the parsing thing, using a shiny interface. I'm happier with this cleaner version that is definitely not the one I'd use had there been only 5 days in December.
 I'm starting to feel the need for a cleaner version of my `if parts.Contain(door.Prima)`. I'm thinking using a `map[door.Part]func()` for computing, and a `map[door.Part]{resultType}` for responses. This means changing from `if + if` to `for`, everywhere. I personally think it better - the underlying question being "_is 2 a lot?_"
 
 Just read the Part 1 item. So much for writing a parser that creates objects now an object can be spread over several lines... Well, at least we know when an object is complete - that only happens when we know a new object is needed, which happens on the first line, and on any empty line. The parser for this one will be slightly different, but the current implementation should allow for it.
 
 I have realised I'm updating previous code, which makes seeing improvements trickier. From now on, I'll try to stop updating whatever happened in the past (in the `2020/decXX` files).
 
 I was about to declare a struct with Height being an `int`, but then I saw a line where its value was declared in inches. `string` it is, then. Please, use coherent units. If you don't know which unit you should use, refer to [standards](https://en.wikipedia.org/wiki/International_System_of_Units). Funfact: I had doubts, when I created the `input` file, that I'd mess with the original input file. Indeed, for git-reasons, I wanted the file to finish on an empty line, which, as it would turn out, would have a specific meaning for our parsing. Thankfully, it had no impact whatsoever. 
 
 After the first line writing, it appears the passportID, which I expected to be an int, is in fact a string. Also, there is absolutely no need for a countryID storage, so I decided to drop it. Consider it a _performance optimization_, or what have you.
 
 Part two required, in my opinion, a lot of verbose and unintelligent code. My first attempt failed miserably because I forgot that regexp.FindStringSubmatch returns not only the matches, but also, as its first parameter, the input string. This is nice, as it lets you check for `res[1]` to find the contents of your first captured expression, but it also means the length of `res` is not the number of captured expressions ! 
 
 I'm happy with the way things are in the `dec04.go` file. The code I've written will easily be copy-pastable for days to come. This probably means I'm missing something and that I could factorise some lines. On a (very) small note, changing from `if/if` to `for` on a map made my code undeterministic. So far, this is fine, but I might want to rewrite the `String()` func on results - I've written a specific key sorter for this. For this day, I've written my function-related constants within the functions. I didn't want them dangling around.
 
 Stuff I want to do in the common sections: prolly rename that `common` package to something cleaner. Add some automatic runs on github.
 
 ## December 05
 Being a best friend's best man, I couldn't write code neither the 5th nor the 6th. Well. Not before late PM. I've added some UTs over dec04, and I must admit that, since the instructions provide examples, lest we can do is use them. But do I really want to write UTs to cover all my code as I did in the previous examples? Yes. But will I do it for a "fun project" I work on on my free time? That depends on how much time I can dedicate to this. From here on, I won't guarantee a 100% coverage of the code, but I will guarantee a 100% coverage of the doc examples. 
 
 Unveiling the 5th box, and it seems we're going to meddle with dyadic expressions whereto developers so often refer as _binary numbers_. Someone decided to use letters to let people know where to sit in a plane, and it seems to me it's someone with an addiction to the Higher/Lower game. What's the motivation? Wouldn't storing the string of the seat `"127"` be both simpler and shorter?
 
 The first approach could be to store each seat that is occupied in a slice, as we've done for previous days, but since we're looking for unused seats, I'm thinking removing items from a list of all possibles during parsing could make sense.
 
 Ah. There's a new question that is raised on Dec05. What is that image that is starting to appear on the adventofcode page? At fist I though of a Christmas tree (till December 5th). By now, this makes absolutely no sense. Could it be we're drawing a map of the traveller's path, suggested a radiant acquaintance.
 
 Before I started writing code, I copy-pasted the input file, and checked it had less than `2^(nbRows+nbCols)`, which equals 1024 (7 rows, 3 cols). Working for 10 years as a software engineer has taught me never to fully trust a software engineer you don't already trust as a friend.
  
 Job's done. I really don't like my `availableSeats` func. I think it's wrong. At first (see paragraphs above) I wanted to store only free seats, but that didn't work with the Part1 exercise. If these two problems were distinct, I'd write a different func for both. I have at least 2 reasons for not writing the second func right now: it's late, and the problems aren't distinct at all.
 
## December 06
Wedding and work have put me behind schedule, but I am willing to do my best and catch up with the doors yet to open. I've got tea, clementines, biscuits, and internet. If today's exercise seems somewhat similar to those of yester days, I'll write some template/script to generate `dec##` folder. I should also focus on my CI about now.

A quick glance at the input shows the order of the answers isn't alphabetical (that is, `acb` is a possible line). It won't matter, but it's still interesting, and it might come in handy sometime. This exercise looks quite a lot like the passport one (dec04) - with data for "an entry" spreading over several lines, and with the empty line as an "entry separator". Our parser here will be very similar.

Part1 wasn't that difficult. Part2 seems a bit more twisted, but I'll use the answers from the first person as "the reference" whereto others could comply. As we explore the answers from others, I'll remove items from it. Turns out I had written a "lazy" implementation for Part1, where I wouldn't consider groups as lists of individuals, but rather as bunches of answers. As I have _decided_ to use the same parser for both parts, I'll readapt the code for Part1. It'll be slightly more complex.

I think my answers here are not time-optimal, nor memory-optimal. Since we parse each file line by line, I could do processing right there, rather than after the reading. The way I have writ my code allows for, in my opinion, distinction between the data structure and its use.

Something rather funny (if you have a broken sense of humour) happened to me on this one. I was writing tests, and decided to swap parts and write the Secunda first, and the Prima second. Turns out I realised one of my methods was altering the data structure, because I had taken the liberty of assigning the value of a field to a variable, when said field was a map, and said variable was to be updated with `delete` somewhere further. I fixed this by writing a func that basically copies a map.

## December 07
After a day of rest, I've grabbed my cup of tea and here I am back to opening doors. It seems we've landed in an area surrounded by glimmeringly green trees. CI is up and running, and I've added shiny tags to my project.

Today's puzzle is going to be about tree parsing. We'll have to care about infinite loops (`big blue bag contains 2 big blue bags` or the like). Studying 3 years in Grenoble gives "2 shiny chartreuse bags" a very specific meaning. This means thinking about design is paramount here. Paramounter than in previous exercises, would I dare say, were it possible.

Let's analyse a bit the input: each container contains some amount of other containers. Right now, I'll keep in mind that
- no container contains itself
- no container A contains a container B that ultimately would contain container A.
- each line has the pattern `${bagColour} bags contains [${childBagCount} ${childBagColour} bag(s)[,]?]+.`, except for leaves, which have the "special colour" `no other`.
 
No rule about "X contains Y" only appearing after a line reading "Y contains Z", since the first line has an unseen container, and the last line has obviously a previously seen container.

We'll end up with a graph, but we don't know yet its roots nor its leaves. We'll build the graph as we parse the input line. It could be that a node has several parents/children (a brown bag could be both in a red bag and a black bag). Basically, we could build our tree with 
```go
type bagName string
type bag struct {
	parents map[bagName]int // this bag can be found in these bags
}
```
and we'll hold the data with
```go
type bagIndex map[bagName]bag
```

The choice of holding parents rather than children is only based on the fact that the Part1 exercise asks for a count of ancestors.

The option of retro-engineering is also present, where I make the assumption that whoever wrote the exercise made it sure that no invalid output would be generated, and that, maybe, the bags are already sorted (that is, the first line of the file is a bag with no parent, and the last lines of the file are only leaves), and the tree already exists.

Implementation-wise, I spent way too much trying to squeeze two regexps into one and hope it'd work. I also, for the moment, left aside any loops (A includes B, which includes A). I finally finished Part1 (which had me check the implementation of the built-in `delete` func). Almost surprisingly, Part2 is quite similar. I'm happy I wrote both a `parents` and a `children` structure, I don't have to erase my code :). However, I don't want to count the children of a bag several times, so I'll add a pointer to an int in each bag that will "remember" how many bags it can hold.

## December 08
Life has its ways and I'm having trouble keeping up with the speed of time. The map (we can safely assume it's a travelling path, now) is no longer linear (a line per day).We're off for an adventure over the woods.

What seemed like a pleasant flight has become debugging some assembler code. Yummie. Part1, however, doesn't look that nasty. We'll store instructions, and already executed instructions (_via_ a flag, for instance), and that should be enough.

Part1 is done, I added a few checks after going through an infinite loop myself. Part2 seems a bit trickier. Brute-forcing the line to change would be an O(n²) (as we have to execute the program each time). There might be a cleaner solution, but I'm going all-in with the `run` inside the loop.

I find this code cleaner than "yesterday's", it doesn't have pointers to maps, or the like. Nevertheless, trying to find loops and fix them in low-level languages, despire being rather fun for an hour, would not be my dream job.

## December 09
Someone told me I look weird in a suit. I agree with them. Let's focus on the real activity here: hacking an airborne plane. This problem reminisces about one of the first. We're looking for numbers that [do not] sum up to a given number. I think this time I'll try a new approach, at least for Part1. Indeed, we only need to store information about the last 25 numbers. I think we can even trim that down to a total of 300 (25*25 - 25)/2. Let me give you an example of a dataset of 10 input digits (a through h):


|   |  a  |  b  |  c  |  d  |  e  |  f  |  g  |  h  |
|---| --- | --- | --- | --- | --- | --- | --- | --- |
| a | a+a | a+b | a+c | a+d | a+e | a+f | a+g | a+h |
| b | b+a | b+b | b+c | b+d | b+e | b+f | b+g | b+h |
| c | c+a | c+b | c+c | c+d | c+e | c+f | c+g | c+h |
| d | d+a | d+b | d+c | d+d | d+e | d+f | d+g | d+h |
| e | e+a | e+b | e+c | e+d | e+e | e+f | e+g | e+h |
| f | f+a | f+b | f+c | f+d | f+e | f+f | f+g | f+h |
| g | g+a | g+b | g+c | g+d | g+e | g+f | g+g | g+h |
| h | h+a | h+b | g+c | h+d | h+e | h+f | h+g | h+h |

This is the list of all sums of 8 numbers (a-h). However, a very short glance will let you know that the upper (or lower) half is also present on the other side of the diagonal, due to the symmetry property of + on integers, and that the diaganal is hit only when we add an item to itself. 

In our scenario, we won't be doing such shenanigans. So we can trim our table down to this: (using the bottom-half is easier)

|     | 0:a | 1:b | 2:c | 3:d | 4:e | 5:f | 6:g | 7:h |
|-----| --- | --- | --- | --- | --- | --- | --- | --- |
| 0:a |     |     |     |     |     |     |     |     |
| 1:b | b+a |     |     |     |     |     |     |     |
| 2:c | c+a | c+b |     |     |     |     |     |     |
| 3:d | d+a | d+b | d+c |     |     |     |     |     |
| 4:e | e+a | e+b | e+c | e+d |     |     |     |     |
| 5:f | f+a | f+b | f+c | f+d | f+e |     |     |     |
| 6:g | g+a | g+b | g+c | g+d | g+e | g+f |     |     |
| 7:h | h+a | h+b | g+c | h+d | h+e | h+f | h+g |     |

Now let's say we want to check the next line, which happens to be line #253. This line reads a value (V) which must be the sum of two of the previous 8 items. We check in the table to know if we can find that value we are looking for by scanning it. If we don't find it, that's it, case closed. If we do find it, however, this line #253 will override whatever was on line #253-8 = 245.

Since 253 = 3 + 31*8, 253 = 3 mod 8, so we'll update the line/column 3 (which would be d, here, since we start counting at 0). The new table will be:

|     | 0:a | 1:b | 2:c | 3:V | 4:e | 5:f | 6:g | 7:h |
|-----| --- | --- | --- | --- | --- | --- | --- | --- |
| 0:a |     |     |     |     |     |     |     |     |
| 1:b | b+a |     |     |     |     |     |     |     |
| 2:c | c+a | c+b |     |     |     |     |     |     |
| 3:V | V+a | V+b | V+c |     |     |     |     |     |
| 4:e | e+a | e+b | e+c | e+V |     |     |     |     |
| 5:f | f+a | f+b | f+c | f+V | f+e |     |     |     |
| 6:g | g+a | g+b | g+c | g+V | g+e | g+f |     |     |
| 7:h | h+a | h+b | g+c | h+V | h+e | h+f | h+g |     |

This updates 7 cells of the table. The 2D table can be flattened down to a 1D slice :

|     | 0:a | 1:b | 2:c | 3:d | 4:e | 5:f | 6:g | 7:h |
|-----| --- | --- | --- | --- | --- | --- | --- | --- |
| 0:a |     |     |     |     |     |     |     |     |
| 1:b |  0  |     |     |     |     |     |     |     |
| 2:c |  1  |  2  |     |     |     |     |     |     |
| 3:d |  3  |  4  |  5  |     |     |     |     |     |
| 4:e |  6  |  7  |  8  |  9  |     |     |     |     |
| 5:f | 10  | 11  | 12  | 13  | 14  |     |     |     |
| 6:g | 15  | 16  | 17  | 18  | 19  | 20  |     |     |
| 7:h | 21  | 22  | 23  | 24  | 25  | 26  | 27  |     |

we can find the index in the slice with the following formula, where `l` is the line (starts at 0 for a) and `c` is the column (starts at 0 for a) (with `l > c` because we are in the lower half):
```
sIndex := c + l*(l-1)/2
```
etc.

In the end, I spent way to much time writing stuff for this optimisation. I decided to scrap it, and go for the "write code that works" approach.

Part Two is about finding a set of numbers that add up to your vulnerability. Obviously, we need only look in the numbers before the vulnerability - otherwise, we'd be adding the vulnerability to some numbers and hope that doesn't exceed the vulnerability.

For this part, we'll need to increase the contiguous range sum by adding the "next" numbers from the list, and we can reduce it by removing the first numbers we've read. The code will basically be: 
```go
while !equal() {
	if smaller() {
		lastElementPosition++
	} else {
		firstElementPosition++
	}
}
```

It doesn't make sense to me to keep track of the highest/lowest values in the `[firstElementPosition:lastElementPosition]` range we're considering. However, when performing `a+b+c` and then `a+b+c+d`, we can save some time and store the result of `a+b+c`, and then either add `d` or remove `a`. The func that then looks for the minimum and the maximum can be slightly optimized. Indeed, we know that the minimum is going to be in one of the first 25 (preambleLength) values, and the maximum is going to be in one of the last 25 (preambleLength) values, per construction.

