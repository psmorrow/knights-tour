# Knight's Tour
A simple heuristic brute force solution for the knight's tour problem.

A knight's tour is a sequence of moves of a knight on a chessboard such that the knight visits every square only once.

A brute-force search for a knight's tour is impractical on all but the smallest boards; for example, on an 8 × 8 board there are approximately 4×10^51 possible move sequences, and it is well beyond the capacity to perform operations on such a large set.

To build: `go build -o knights knights.go`

[![Sonarcloud Status](https://sonarcloud.io/api/project_badges/measure?project=psmorrow_knights-tour&metric=alert_status)](https://sonarcloud.io/dashboard?id=psmorrow_knights-tour)

## Things to consider:

### Divide and conquer algorithm
By dividing the board into smaller pieces, constructing tours on each piece, and patching the pieces together, one can construct tours on most rectangular boards in linear time - that is, in a time proportional to the number of squares on the board.
