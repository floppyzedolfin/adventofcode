## December 13

Part 1 wasn't really interesting. Nothing much to do.

Part 2, on the other side, seems way more complex. Brute forcing it using the largest bus number might be a solution, but I want to try something else.

We're looking for a number N such that
```
  N % a1 = 1
  N % a2 = 2
  N % a3 = 3 
  ...
  N % an = n 
```

After looking at the bus numbers in the input dataset, we notice that they are all prime numbers. This might point towards a "nice" implementation for a solution, but we'll have to call our dear Bézout. This case is very similar to the [Chinese remainder theorem](https://en.wikipedia.org/wiki/Chinese_remainder_theorem).

So I've worked that solution and implemented it.
I was very eager to know whether my code was operational, so I ran it against the provided input, and submitted that solution. Of course, it was wrong on the first go, but that didn't stop me from trying a second run after noticing I'd taken the opposite remainders at some part.

After a couple of such attempts, I decided to test my code against the provided examples, and, with logs, managed to find what was wrong. I fixed it, ran my code against the provided examples, and, hurray!, my unit tests were green.

I ran the code against the provided input, and it failed. I'm now supposing `int64` isn't enough to store the response. I'll try the `big` package tomorrow.

Then something unexpected happened. A new neighbour moved in, and they've foolishly decided they'd want an Internet connection. Why do people do this to me, really ?

Turns out their connecting to the Interwebs also, as a side effect, ended my subscription to my ISP. Of course, I wasn't notified I'd be thrown away, but I did my best to help them fix this as fast as possible. "48h max delay before it's fixed" quickly became "7 days", and then "7 days with the extra 3 days wait to get someone on the line to get a technician come to your building".

That's not really the worst thing that can happen to a human being. However, it's the worst thing that could happen to me, regarding this set of problems - I can't submit my answers, I can't push my code to get feedback from trusted peers, and I can't get more food for brains on a daily-basis. I'm trying to remain zen about this, "in the end, it'll work". Whoever said the journey was as important as the destination was clearly never left hanging on the phone listening to a recorded voice telling them to wait a little bit more.
