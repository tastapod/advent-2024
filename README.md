# Advent of Code 2024

This year I am using Go. I like the way [Advent of Code][aoc] lets me explore Go's APIs and paradigms.

Some of these would be more expressive in a functional language, and I miss that Go
does not have these kind of language features, like if-as-value. I might try introducing
some Map or Reduce-like behaviour, which I started doing [last year][fn].

[aoc]: https://adventofcode.com/
[fn]: https://github.com/tastapod/advent2023/blob/main/fn/fn.go

## Notes

### Day 2

I feel part 2 could be less clunky / more elegant in terms of how I am removing the term for each check.

### Day 3

I am liking how Go does regexp capture groups. This solution was neat (and quick!) and I like how the code came out.

### Day 4

I lost a chunk of time by not realising that `range` defaults to providing its index rather than its values, by which I mean:

```golang
for delta := range []int{-1, 0, 1} {
	println(delta)
}

Prints:
0
1
2
```

To get the actual values, you need:

```go
for _, delta := range []int{-1, 0, 1} {
	println(delta)
}

Prints:
-1
0
1
```

### Day 5

This was a fun one. I made an assumption that I just needed the literal rule pairs and that I wouldn't need to build a graph, which turned out to be true. The rules contain every pair combo so there is no transitive logic.

Because of this, for part 2 I was able to just use the builtin `slice.SortedFunc` function, passing it a comparator that does an existence check in the rule table. If the rule exists, the pair is in the correct order. Simples!

Spoiler! The solution is 'dumber' than I realised. You can think of the page numbers just as symbols. The only time you need to parse a value is when you are summing the middle values. The rest of the time, the rule checking is just an existence check for "before|after". So now I have dumbed down my solution.


### Day 6

This is weirdly algorithmic. I started with the body of a test and just kind of wrote the code in one hit. I was able to reuse the grid padding code from Day 4.

As I was writing the code to take a step, the state machine just kind of fell into place. I had a couple of quirks that my tests and tracing quickly helped with (I forgot to replace the initial position with a '.', for instance!) but part 1 fell together fairly straightforwardly.
