## December 12

This time we're moving on a grid. I'll assume a few things:
- the grid is infinite - we can travel by boat for 1984657543 units eastwards
- we won't have to deal with "difficult" values - angles that aren't multiples of 90, for instance.

Part 1 went rather smoothly. I did, however, ran into an obvious bump: I initially assumed the NSEW+F distances were in degrees too, as in latitude/longitude, due to an excess of these units at work. This made me write something wrong when I tried to close the Earth on the antimeridian. Since the example didn't cover this trap, I submitted a wrong answer and got rejected. But that's part of the exercise, so it's fine.

Part 2 will have me recode a few things from Part 1, but hopefully, not too much. I've decided not to encode the waypoint's position, but merely its relative position to the ship. This makes things way simpler.

My first thoughts here were to use complex numbers as they make rotation very easy. Sadly, Go doesn't provide nice tools for complex numbers on an integer-based grid. So I went with my own `(x,y)` structure, implementing rotations, and such operations.
