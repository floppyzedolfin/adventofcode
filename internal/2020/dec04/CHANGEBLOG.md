## December 04

I've reworked the Makefile. I can now run `make run -- -door 3 -parts 1,2`. I've also re-done the parsing thing, using a
shiny interface. I'm happier with this cleaner version that is definitely not the one I'd use had there been only 5 days
in December. I'm starting to feel the need for a cleaner version of my `if parts.Contain(door.Prima)`. I'm thinking
using a `map[door.Part]func()` for computing, and a `map[door.Part]{resultType}` for responses. This means changing
from `if + if` to `for`, everywhere. I personally think it better - the underlying question being "_is 2 a lot?_"

Just read the Part 1 item. So much for writing a parser that creates objects now an object can be spread over several
lines... Well, at least we know when an object is complete - that only happens when we know a new object is needed,
which happens on the first line, and on any empty line. The parser for this one will be slightly different, but the
current implementation should allow for it.

I have realised I'm updating previous code, which makes seeing improvements trickier. From now on, I'll try to stop
updating whatever happened in the past (in the `2020/decXX` files).

I was about to declare a struct with Height being an `int`, but then I saw a line where its value was declared in
inches. `string` it is, then. Please, use coherent units. If you don't know which unit you should use, refer
to [standards](https://en.wikipedia.org/wiki/International_System_of_Units). Funfact: I had doubts, when I created
the `input` file, that I'd mess with the original input file. Indeed, for git-reasons, I wanted the file to finish on an
empty line, which, as it would turn out, would have a specific meaning for our parsing. Thankfully, it had no impact
whatsoever.

After the first line writing, it appears the passportID, which I expected to be an int, is in fact a string. Also, there
is absolutely no need for a countryID storage, so I decided to drop it. Consider it a _performance optimization_, or
what have you.

Part two required, in my opinion, a lot of verbose and unintelligent code. My first attempt failed miserably because I
forgot that regexp.FindStringSubmatch returns not only the matches, but also, as its first parameter, the input string.
This is nice, as it lets you check for `res[1]` to find the contents of your first captured expression, but it also
means the length of `res` is not the number of captured expressions !

I'm happy with the way things are in the `dec04.go` file. The code I've written will easily be copy-pastable for days to
come. This probably means I'm missing something and that I could factorise some lines. On a (very) small note, changing
from `if/if` to `for` on a map made my code undeterministic. So far, this is fine, but I might want to rewrite
the `String()` func on results - I've written a specific key sorter for this. For this day, I've written my
function-related constants within the functions. I didn't want them dangling around.

Stuff I want to do in the common sections: prolly rename that `common` package to something cleaner. Add some automatic
runs on github.
