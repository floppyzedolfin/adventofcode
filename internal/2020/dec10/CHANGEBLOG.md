## December 10

To be fully honest with the reader, we've increased the year counter, gifts have been unwrapped, above average beverages
were drunk.

This puzzle seems very much to be about sorting an input data set, and then count the differences between consecutive
values. Luckily enough, Go implements `sort` for integers, so that's that I don't have to write. It also seems we need
to use _all_ adapters, which means we might have to skip some in Part 2.

I'll make some wild assumptions here.

- First, the input dataset is valid and doesn't have gaps of 4-jolts. For instance, it will always have an adapter with
  joltage 1, 2 or 3 (the one we'll plug in the plane socket).
- Second, all adapters are different - the bag of adapters `{1,1}` wouldn't verify this.

Part 2 is a bit trickier. Differences of `3` split the dataset into a combination of what's left of the 3 and what's
right of it. Basically, if `c` is the combination counting function, and `x,3,y` is a sequence of "something, then a
difference of 3, then something else". Example: `{0,1,2,5,6,8}` has one difference of 3 (between 2 and 5).
> `c(x,3,y) = c(x) * c(y)`

The reasons for gaps of `3` to be unavoidable is that if we remove the number on the left (`2`) or that on the
right (`5`), the gap will expand, and it won't be a valid gap any longer.

In my example, `c({0,1,2,5,6,8}) = c({0,1,2})*c({5,6,8})`, and we've reduced the complexity of the problem.

A short analysis (`sort -n 2020/dec10/input | awk 'NR > 1 { print $0 - prev } { prev = $0 }' | sort -u`) of the provided
dataset shows we have no gaps of 2 jolts. This makes the problem way simpler, as we only need to count the number of
gaps of ones between each gap of 3 to compute the total number of combinations.

So, what we are looking at, now, a set of consecutive numbers bounded by gaps of 3(`{11,14,15,16,17,18,21}`), and we're
wondering how we can remove some and still have a valid dataset. When looking at gaps between these numbers, this
changes to removing some `1`s from `{3,1,1,1,1,3}` without creating a gap of length `3`.

If we think about it a bit more, removing an item from the list is the same thing as adding two consecutive gaps. This
means the problem is now equivalent to "how many `+` can we write between the `1` without having a number reach 4 or
more in the sequence `1 1 1 1` ?" In this example, `1+1 1 1` would be a valid change.

The number of different combinations is returned by the Tribonacci sequence (after skipping initial zeroes):
> | n | c(n) |
> |---|------|
> | 0 |    1 | 
> | 1 |    1 |
> | 2 |    2 |
> | 3 |    4 |
> | 4 |    7 |
> | 5 |   13 |
> | 6 |   24 |
> | 7 |   44 |
> etc.

See [here](https://oeis.org/A000073), and works from Emeric Deutsch, for details.
