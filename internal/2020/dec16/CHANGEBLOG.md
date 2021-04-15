## December 16

Obviously, the problem isn't just about checking if numbers are in a slice. The easy way to implement part 1 is to keep a record of every number we've seen, and check for that number. Using a `map[int]struct{}` makes the search O(1). But it's already clear Part2 will ask that a ticket contains fields from different sections. So best keep that in mind whilst writing efficient code for part 1.

Let's have a look at the input:
- first, unintelligible fields are in the first section, and each of them has a `[min1, max1] or [min2, max2]` set of ranges.
- second, comes my ticket
- finally, we've got the surrounding passengers, graciously parsed by the providers of the exercise.

Each ticket has 20 numbers, which matches the number of available fields (how odd, I really wonder what Part2 will be about).

I'm not happy with my field parser. I guess it has to be a little sticky at some point, though. With a bit more diversity in the example (say some fields with not exactly 2 ranges of valid values), I'd've considered writing more generic code (with loops).

Part 1 wasn't too difficult, but I did have to write a `for`/`for`/`if`/`for`/`for`/`if` loop, which isn't something I'm very fond of.

Hurray, my guesses at what monstrosities could hide in Part 2 were partially correct.

Let's try to break down what must be done:
- first, remove invalid tickets (could re-use functions written for Part 1)
- second, we need to represent "ticket value #3 can be fieldA, or fieldC, or fieldD, or ..."

This second part (of the second part, that is) is the most interesting. I'll assume there's only one solution... My first go at the code produced undeterministic results. After a quick inspection, it turns out I currently have several choices for each field. Which means that my lazy approach is not strong enough. Here are the checks I didn't implement:
- numbers in a ticket must correspond to different fields;
- if a field is matched by only one "column", then that column is no longer eligible for other fields.

Now is both too late and late enough for the rest.

It's important, also, to take a few things into account. We have 18 fields, and eligible values are in the range [1,1000] (approx). This means the answer to Part2 is, potentially, 10^18*3? Using `int`s  might not be enough for this one, let's use `uint64`.


My first implementation of the Part2 wouldn't work, and it took me quite some time to understand why. The reason was that, in order to determine whether a ticket was valid, I'd sum the invalid fields, and check if that sum would be 0. One of the cheeky tickets had an invalid field of value 0, which wasn't caught by my code. I updated the func to return a value and a boolean in order to circumvent this conundrum.

The current state of things is that I don't like my code as it is. I feel i've written `func`s to wrap `for`s. This is due to my naive implementation. Of course, other implementations are possible, smarter ones. I've decided to use maps to have `find` operations in O(1).

Most of the code I've written for this exercice has the following (severe) limitations:
- I expect a properly ordered input. Can't swap neighbours' tickets with mine, for instance
- I assume the existence of a solution, and its unicity.
- I expect aforesaid solution to be computable using the information from tickets.

The following scenario wouldn't return great results:
- field A: eligible for fields {1,2}
  and no other field eligible for 1.
  My code can't solve this. 
