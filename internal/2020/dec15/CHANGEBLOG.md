## December 15

After reading some of the code I've written so far, I'd factorize lots of it weren't it "atomical". This `PartsSorter` func, for instance, needs to be changed.

Door 15 opens to a memory game. Option 1 for implementing a solution would be to use a single array and parse it from end to start everytime we get a new number.
I want to do things slightly differently, and keep track of numbers I've already seen.

This solved Part1, was very easy to adapt to Part2, but... it takes a tremendous time to complete the runs!

My first code used to increment the "last seen time" of each entry. I did that before I realised I had to keep track of time. Now that I have my clock available, I can use it to know the "age" of a number.

After a slight change (computing age only when I need it, rather than saving it, which cost an unnecessary loop), I've "improved" the performances of the executable, but computing the Part2 still takes a fair share of time. For this matter, I've commented out the tests provided there.
The code doesn't really seem parallelisable.
