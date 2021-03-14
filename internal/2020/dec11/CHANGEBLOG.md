## December 11

Today seems to be the day we pay tribute to the late John Conway. The problem is about convoluting matrices, but with
the slight trick that some spots on the map are to remain empty.

I've written the code storing the seats to update as a slice between the scanning and the updating, I think it makes the
code work a bit nicer than storing the whole new status.

Part two requires some re-working of the internal wiring. I'll use a parameter for my `WaitingArea` to help me here.

I've also seen the coverage of my code significantly shrink. There are two main reasons for this:
- I've stopped covering the `String` method from the `decXX` packages, as I haven't brought anything new there since dec05
- I've started splitting my code in packages, which means the part of code that is not `String` is also reduced.

However, my UTs cover the inner packages too, but this doesn't appear in `go test --cover ./...`'s output. Also, for some obscure reason, it seems I can no longer run tests in parallel (with `t.Run`), as it messes with the computations.
