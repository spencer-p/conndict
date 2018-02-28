# conndict

`conndict` is short for connected dictionaries.

## Background

This project attempts to answer this question:

> Starting from one word in the dictionary, you may look up any word that occurs
> in that word's definition (and so on, recursively). What word should you begin
> with to learn the most words?

This is evidently a graph traversal problem where words are vertices and some
word `w` being in the definition of another word `m` implies a connection `m ->
w`. Then we can answer this question with any graph traversal algorithm.

## What I currently have

I got an API key from Merriam-Webster for looking up definitions and figured out
how to parse their results. I then wrote a simple web page where you can start
with any word and click to expand words in its definition to explore the graph
manually. This can get pretty big pretty quickly.

For running it:

 1. Get an API key [here](https://www.dictionaryapi.com/) and put it in
	`secrets.go`
 1. `go build`
 1. `./conndict -serve`
 1. Navigate to `localhost:8000/view/[your starting word here]`

This is a lot of fun and really cool to look at, but it doesn't really answer
the original question.

## What I want to do

I want to know what the answer to the original question is. This will probably
entail dynamically building my own database of definitions (since
Merriam-Webster is slow). There's a lot of storage commitment, then writing the
code, and then letting that code run for however long it needs.

Since each query to Merriam-Webster takes about 1 second, and a baseline
estimate for words in English is 171,000, running the program alone could take
anywhere between two days and two weeks.
