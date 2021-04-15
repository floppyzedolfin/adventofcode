## December 17

I decided to somewhat reorganise the code. `adventofcode/pkg` contains common utilities (that were previously at root level); and `adventofcode/internal` contains specific pieces of code that aren't to be commonified.

Today is about Conway. Obviously. The man was a genius. We're now playing 3d-life. My first thoughts were that the evolution is symmetrical with regard to the `z=0` plane. So, no real need to compute whatever happens for z<0, as it will mirror whatever happens for z > 0.
We've already worked on something quite similar - when we were trying to seat passengers waiting for their plane, on December 11th. The main difference being that, this time, space isn't bounded. For this reason, I don't want to have a map of all the existing space, but I'll rely on something else to keep track of which places are busy and which aren't.

The solution I want to use relies on using the `map` structure in Go. I want to mark known positions and their neighbours. I want a `map[3dPoint]numberOfNeighbours`, and that map will help me iterate over time.

There is a slight twist here. The future cube's status depends on its current status. If a cube has 3 neighbours, it will be in active state. However, if it has 2, its future state could be active or inactive, depending on its current state. This means we need to take into account the previous state when building the next.

Turns out my first attempt was successful. Also turns out my nice symmetry trick needs adaptation for a 4th dimension. My code has explicit lists of what the neighbours of a cube are. That was one of the options (these lists are constant throughout time and space). The other implementation would have been something like this, which I found too nested:
```go
for [X] {
	for [Y] {
		for [Z] {
			for [W] {
				// process
			}
		}
	}
}
```

My first attempt didn't work, and I highly suspect my `Neighbour4D` list to have been wrong. I re-generated it with [real tools](https://www.gnu.org/software/gawk/manual/gawk.html), and the output matched the expected result. Next step, for this program, would be to think of reducing the memory consumption, but I'm happy I've got a solution that is `O(number of cubes)` (where the multiplying constant is at least `80x` for the 4D case, but, hey... it doesn't use that much memory space).
