## December 06

Wedding and work have put me behind schedule, but I am willing to do my best and catch up with the doors yet to open.
I've got tea, clementines, biscuits, and internet. If today's exercise seems somewhat similar to those of yester days,
I'll write some template/script to generate `dec##` folder. I should also focus on my CI about now.

A quick glance at the input shows the order of the answers isn't alphabetical (that is, `acb` is a possible line). It
won't matter, but it's still interesting, and it might come in handy sometime. This exercise looks quite a lot like the
passport one (dec04) - with data for "an entry" spreading over several lines, and with the empty line as an "entry
separator". Our parser here will be very similar.

Part1 wasn't that difficult. Part2 seems a bit more twisted, but I'll use the answers from the first person as "the
reference" whereto others could comply. As we explore the answers from others, I'll remove items from it. Turns out I
had written a "lazy" implementation for Part1, where I wouldn't consider groups as lists of individuals, but rather as
bunches of answers. As I have _decided_ to use the same parser for both parts, I'll readapt the code for Part1. It'll be
slightly more complex.

I think my answers here are not time-optimal, nor memory-optimal. Since we parse each file line by line, I could do
processing right there, rather than after the reading. The way I have writ my code allows for, in my opinion,
distinction between the data structure and its use.

Something rather funny (if you have a broken sense of humour) happened to me on this one. I was writing tests, and
decided to swap parts and write the Secunda first, and the Prima second. Turns out I realised one of my methods was
altering the data structure, because I had taken the liberty of assigning the value of a field to a variable, when said
field was a map, and said variable was to be updated with `delete` somewhere further. I fixed this by writing a func
that basically copies a map.
