## December 09

Someone told me I look weird in a suit. I agree with them. Let's focus on the real activity here: hacking an airborne
plane. This problem reminisces about one of the first. We're looking for numbers that [do not] sum up to a given number.
I think this time I'll try a new approach, at least for Part1. Indeed, we only need to store information about the last
25 numbers. I think we can even trim that down to a total of 300 (25*25 - 25)/2. Let me give you an example of a dataset
of 10 input digits (a through h):

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

This is the list of all sums of 8 numbers (a-h). However, a very short glance will let you know that the upper (or
lower) half is also present on the other side of the diagonal, due to the symmetry property of + on integers, and that
the diaganal is hit only when we add an item to itself.

In our scenario, we won't be doing such shenanigans. So we can trim our table down to this: (using the bottom-half is
easier)

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

Now let's say we want to check the next line, which happens to be line #253. This line reads a value (V) which must be
the sum of two of the previous 8 items. We check in the table to know if we can find that value we are looking for by
scanning it. If we don't find it, that's it, case closed. If we do find it, however, this line #253 will override
whatever was on line #253-8 = 245.

Since 253 = 3 + 31*8, 253 = 3 mod 8, so we'll update the line/column 3 (which would be d, here, since we start counting
at 0). The new table will be:

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

we can find the index in the slice with the following formula, where `l` is the line (starts at 0 for a) and `c` is the
column (starts at 0 for a) (with `l > c` because we are in the lower half):

```
sIndex := c + l*(l-1)/2
```

etc.

In the end, I spent way to much time writing stuff for this optimisation. I decided to scrap it, and go for the "write
code that works" approach.

Part Two is about finding a set of numbers that add up to your vulnerability. Obviously, we need only look in the
numbers before the vulnerability - otherwise, we'd be adding the vulnerability to some numbers and hope that doesn't
exceed the vulnerability.

For this part, we'll need to increase the contiguous range sum by adding the "next" numbers from the list, and we can
reduce it by removing the first numbers we've read. The code will basically be:

```go
while !equal() {
  if smaller() {
    lastElementPosition++
  } else {
    firstElementPosition++
  }
}
```

It doesn't make sense to me to keep track of the highest/lowest values in
the `[firstElementPosition:lastElementPosition]` range we're considering. However, when performing `a+b+c` and
then `a+b+c+d`, we can save some time and store the result of `a+b+c`, and then either add `d` or remove `a`. The func
that then looks for the minimum and the maximum can be slightly optimized. Indeed, we know that the minimum is going to
be in one of the first 25 (preambleLength) values, and the maximum is going to be in one of the last 25 (preambleLength)
values, per construction.
