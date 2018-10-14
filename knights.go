// A knight's tour is a sequence of moves of a knight on a chessboard such that the knight visits every square only once.
// A brute-force search for a knight's tour is impractical on all but the smallest boards; for example, on an 8 × 8 board there are approximately 4×1051 possible move sequences, and it is well beyond the capacity of modern computers (or networks of computers) to perform operations on such a large set.
// The heuristics approach is described here: https://en.wikipedia.org/wiki/Knight%27s_tour

package main

import (
    "fmt"
    "os"
    "sort"
    "time"
)

type move struct {
    row int
    column int
    heuristic int
}

const squaresOnSide int = 6
const squaresOnBoard int = squaresOnSide * squaresOnSide

var possibleMoves = []move { {+1, -2, 0}, {+1, +2, 0}, {+2, -1, 0}, {+2, +1, 0}, {-1, -2, 0}, {-1, +2, 0}, {-2, +1, 0}, {-2, +1, 0} }
var heuristics [squaresOnSide][squaresOnSide]int
var heuristicMoves [squaresOnSide][squaresOnSide][]move
var board [squaresOnSide][squaresOnSide]int

func promptUser(promptText string, minValue int, maxValue int, errorText string) int {
    var value int

    // input 1 based
    fmt.Print(promptText)
    _, err := fmt.Scan(&value)
    if err != nil {
        fmt.Println("Invalid input:", err)
        os.Exit(1)
    }
    if value < minValue || value > maxValue {
        fmt.Println(errorText)
        os.Exit(2)
    }

    // adjust to 0 based
    value--

    return value
}

func calculateHeuristics() {
    // calculate heuristic values for each square
    for row := 0; row < squaresOnSide; row++ {
        for column := 0; column < squaresOnSide; column++ {
            var squareHeuristic int = 0
            for _, nextMove := range possibleMoves {
                if row+nextMove.row >= 0 && row+nextMove.row < squaresOnSide && column+nextMove.column >= 0 && column+nextMove.column < squaresOnSide {
                    squareHeuristic += 1
                }
            }
            heuristics[row][column] = squareHeuristic
        }
    }

    // calculate heuristic moves for each square
    for row := 0; row < squaresOnSide; row++ {
        for column := 0; column < squaresOnSide; column++ {
            var squareHeuristics []move
            for _, nextMove := range possibleMoves {
                if row+nextMove.row >= 0 && row+nextMove.row < squaresOnSide && column+nextMove.column >= 0 && column+nextMove.column < squaresOnSide {
                    squareHeuristics = append(squareHeuristics, move{ nextMove.row, nextMove.column, heuristics[row+nextMove.row][column+nextMove.column] })
                }
            }
            // sort possible move heuristics
            sort.SliceStable(squareHeuristics, func(i, j int) bool { return squareHeuristics[i].heuristic < squareHeuristics[j].heuristic })
            heuristicMoves[row][column] = squareHeuristics
        }
    }
}

func solveBoard(row int, column int, level int) bool {
    // is this move invalid? it is off the board or the space is occupied
    if row < 0 || row >= squaresOnSide || column < 0 || column >= squaresOnSide || board[row][column] != 0 {
        return false
    }

    // mark the position occupied by setting it to the level
    board[row][column] = level

    // is the board solved?
    if level == squaresOnBoard {
        return true
    }

    // calculate the next position (heuristics order of next moves)
    var heuristicMoves = heuristicMoves[row][column]
    for _, nextMove := range heuristicMoves {
        if solveBoard(row+nextMove.row, column+nextMove.column, level+1) {
            return true
        }
    }

    // next move failed, so this one does too, undo current move
    board[row][column] = 0
    return false
}

func printBoard() {
    // print out the board as a matrix displaying the move numbers at each row and column
    fmt.Println("Board");
    for row := 0; row < squaresOnSide; row++ {
        for column := 0; column < squaresOnSide; column++ {
            if column > 0 {
                fmt.Print("  ");
            }
            fmt.Print(fmt.Sprintf("%02d", board[row][column]))
        }
        fmt.Println("")
    }
    fmt.Println("")
}

func main() {
    fmt.Println("Knight's Tour")
    fmt.Println(fmt.Sprintf("%d x %d Board", squaresOnSide, squaresOnSide))
    fmt.Println("")

    // ask the user for the beginning board position, both row and column
    fmt.Println("Choose a beginning square on the board.")
    var row int = promptUser(fmt.Sprintf("Enter the row (1-%d): ", squaresOnSide), 1, squaresOnSide, fmt.Sprintf("Invalid row: You must enter a number between 1-%d inclusive.", squaresOnSide))
    var column int = promptUser(fmt.Sprintf("Enter the column (1-%d): ", squaresOnSide), 1, squaresOnSide, fmt.Sprintf("Invalid column: You must enter a number between 1-%d inclusive.", squaresOnSide))
    fmt.Println("")

    // start calulating the solution based upon the user's beginning position
    var start time.Time = time.Now()
    calculateHeuristics()
    solveBoard(row, column, 1)

    // output the solution and duration it took to calculate
    printBoard()
    fmt.Println("Elapsed time (ms): ", time.Since(start))
}
