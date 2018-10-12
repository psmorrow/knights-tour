// A knight's tour is a sequence of moves of a knight on a chessboard such that the knight visits every square only once.
// A brute-force search for a knight's tour is impractical on all but the smallest boards; for example, on an 8 × 8 board there are approximately 4×1051 possible move sequences, and it is well beyond the capacity of modern computers (or networks of computers) to perform operations on such a large set.

package main

import (
    "fmt"
    //"math/rand"
    "os"
    "time"
)

type move struct {
    row int
    column int
}

const squaresOnSide int = 6
const squaresOnBoard int = squaresOnSide * squaresOnSide

var possibleMoves = []move { {+1, -2}, {+1, +2}, {+2, -1}, {+2, +1}, {-1, -2}, {-1, +2}, {-2, +1}, {-2, +1} }
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

    // calculate the next position (constant order of next moves)

    for _, nextMove := range possibleMoves {
        if solveBoard(row+nextMove.row, column+nextMove.column, level+1) {
            return true
        }
    }

    // calculate the next position (random order of next moves)

/*
    var orderMoves []int = rand.Perm(len(possibleMoves))
    for _, index := range orderMoves {
        var nextMove move = possibleMoves[index]
        if solveBoard(row+nextMove.row, column+nextMove.column, level+1) {
            return true
        }
    }
*/

    // next move failed, so this one does too, undo current move

    board[row][column] = 0
    return false
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
    solveBoard(row, column, 1)

    // output the solution and duration it took to calculate

    printBoard()
    fmt.Println("Elapsed time (ms): ", time.Since(start))
}
