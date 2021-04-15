## December 07

After a day of rest, I've grabbed my cup of tea and here I am back to opening doors. It seems we've landed in an area
surrounded by glimmeringly green trees. CI is up and running, and I've added shiny tags to my project.

Today's puzzle is going to be about tree parsing. We'll have to care about infinite
loops (`big blue bag contains 2 big blue bags` or the like). Studying 3 years in Grenoble gives "2 shiny chartreuse
bags" a very specific meaning. This means thinking about design is paramount here. Paramounter than in previous
exercises, would I dare say, were it possible.

Let's analyse a bit the input: each container contains some amount of other containers. Right now, I'll keep in mind
that

- no container contains itself
- no container A contains a container B that ultimately would contain container A.
- each line has the pattern `${bagColour} bags contains [${childBagCount} ${childBagColour} bag(s)[,]?]+.`, except for
  leaves, which have the "special colour" `no other`.

No rule about "X contains Y" only appearing after a line reading "Y contains Z", since the first line has an unseen
container, and the last line has obviously a previously seen container.

We'll end up with a graph, but we don't know yet its roots nor its leaves. We'll build the graph as we parse the input
line. It could be that a node has several parents/children (a brown bag could be both in a red bag and a black bag).
Basically, we could build our tree with

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

The choice of holding parents rather than children is only based on the fact that the Part1 exercise asks for a count of
ancestors.

The option of retro-engineering is also present, where I make the assumption that whoever wrote the exercise made it
sure that no invalid output would be generated, and that, maybe, the bags are already sorted (that is, the first line of
the file is a bag with no parent, and the last lines of the file are only leaves), and the tree already exists.

Implementation-wise, I spent way too much trying to squeeze two regexps into one and hope it'd work. I also, for the
moment, left aside any loops (A includes B, which includes A). I finally finished Part1 (which had me check the
implementation of the built-in `delete` func). Almost surprisingly, Part2 is quite similar. I'm happy I wrote both
a `parents` and a `children` structure, I don't have to erase my code :). However, I don't want to count the children of
a bag several times, so I'll add a pointer to an int in each bag that will "remember" how many bags it can hold.
