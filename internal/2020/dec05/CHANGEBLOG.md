## December 05

Being a best friend's best man, I couldn't write code neither the 5th nor the 6th. Well. Not before late PM. I've added
some UTs over dec04, and I must admit that, since the instructions provide examples, lest we can do is use them. But do
I really want to write UTs to cover all my code as I did in the previous examples? Yes. But will I do it for a "fun
project" I work on on my free time? That depends on how much time I can dedicate to this. From here on, I won't
guarantee a 100% coverage of the code, but I will guarantee a 100% coverage of the doc examples.

Unveiling the 5th box, and it seems we're going to meddle with dyadic expressions whereto developers so often refer as _
binary numbers_. Someone decided to use letters to let people know where to sit in a plane, and it seems to me it's
someone with an addiction to the Higher/Lower game. What's the motivation? Wouldn't storing the string of the
seat `"127"` be both simpler and shorter?

The first approach could be to store each seat that is occupied in a slice, as we've done for previous days, but since
we're looking for unused seats, I'm thinking removing items from a list of all possibles during parsing could make
sense.

Ah. There's a new question that is raised on Dec05. What is that image that is starting to appear on the adventofcode
page? At fist I though of a Christmas tree (till December 5th). By now, this makes absolutely no sense. Could it be
we're drawing a map of the traveller's path, suggested a radiant acquaintance.

Before I started writing code, I copy-pasted the input file, and checked it had less than `2^(nbRows+nbCols)`, which
equals 1024 (7 rows, 3 cols). Working for 10 years as a software engineer has taught me never to fully trust a software
engineer you don't already trust as a friend.

Job's done. I really don't like my `availableSeats` func. I think it's wrong. At first (see paragraphs above) I wanted
to store only free seats, but that didn't work with the Part1 exercise. If these two problems were distinct, I'd write a
different func for both. I have at least 2 reasons for not writing the second func right now: it's late, and the
problems aren't distinct at all.
