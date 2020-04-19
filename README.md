# Reversing a string in Go

A while back,
I got a whiteboard interview question about how I would reverse a string in Go.
The problem wasn't very formally stated,
and in retrospect,
I am amazed I got that job.

The problem is a bit weirder than you'd think at first.
You can't just reverse the bytes of a string.
Go uses UTF-8 encoding,
which has multi-byte characters,
and the order of the multi-byte character's bytes matters.

You can't just cast a string to `[]rune` and reverse the array of runes.
There's more than one way to represent some characters.
Unicode code points can "combine" into the same character,
both visually and officially.

    "\u00e9" == "\u0065\u0301"

at least visually and officially,
but if you reverse the rune order of "\u0065\u0301" inside a string,
you get the wrong thing reversed.

You have to [normalize](https://blog.golang.org/normalization) the UTF-8 string before reversing,
and even then, I bet there's some corner cases that will trip you up.

### Programs

* `mkam` - generate one of two different byte-level encodings of "Amélie"
   1. `mkam  0 > amelie0` gets you the 8-byte version
   2. `mkam  1 > amelie1` gets you the 7-byte version
* `v1` and `v2` illustrate two ways to reverse the array of runes
* `prenormalizing` - tries two ways of normalizing unicode strings,
and shows you the differences between the two ways.

To see everything:

```go
$ go build mkam.go
$ ./mkam 0 > amelie0
$ ./mkam 1 > amelie1
$ go build v1.go
$ ./v1 $(cat amelie0)
"Amélie"
"eiĺemA"
```

Ha ha ha! Naive rune-reversal moves the accent from the "e" to the "l".
But not always:

```go
$ ./v1 $(cat amelie1)
"Amélie"
"eilémA"
```

To try normalizing the input string before reversing the runes:

```go
$ go build prenormalizing.go
$ ./prenormalizing $(cat amelie0)
NFC normalized:
"Amélie"
"eilémA"
NFD normalized:
"Amélie"
"eiĺemA"
$ ./prenormalizing $(cat amelie1)
NFC normalized:
"Amélie"
"eilémA"
NFD normalized:
"Amélie"
"eiĺemA"
```

NFC normalizing the string before rune-array-reversing works.
But there are certainly cases where you'd want to use NFD normalizing.

This ends up being a subtle problem,
because it depends on knowing Unicode,
particulars of UTF-8 encoding,
and some small knowledge of Go.
If you ask it in an interview,
you should probably have an expected answer based on the level
of experience you're hiring for.
You should be prepared for a lengthy discourse if you interview
someone who's well-acquainted with Unicode.
You might get a really unique solution once in a while.
