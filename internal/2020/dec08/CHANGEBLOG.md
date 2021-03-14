## December 08

Life has its ways and I'm having trouble keeping up with the speed of time. The map (we can safely assume it's a
travelling path, now) is no longer linear (a line per day).We're off for an adventure over the woods.

What seemed like a pleasant flight has become debugging some assembler code. Yummie. Part1, however, doesn't look that
nasty. We'll store instructions, and already executed instructions (_via_ a flag, for instance), and that should be
enough.

Part1 is done, I added a few checks after going through an infinite loop myself. Part2 seems a bit trickier.
Brute-forcing the line to change would be an O(n²) (as we have to execute the program each time). There might be a
cleaner solution, but I'm going all-in with the `run` inside the loop.

I find this code cleaner than "yesterday's", it doesn't have pointers to maps, or the like. Nevertheless, trying to find
loops and fix them in low-level languages, despite being rather fun for an hour, would not be my dream job.
