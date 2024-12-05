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

This was a fun one. I made an assumption that all I just needed the rule pairs, and that I wouldn't need to build a graph, which turned out to be true. The rules contained every pair combo so there was no transitive logic.

Because of this, for part 2 I was able to just the builtin `slice.SortedFunc` function, passing it a comparator that does an existence check in the rule table. If the rule exists, they are already in the correct order. Simples!
