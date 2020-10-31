package main

import (
    "fmt"
    // "bufio"
    "os"
    "strconv"
    "math/rand"
    "time"
    // "io/ioutil"
)

const DEBUG = false

func main() {
    rand.Seed(time.Now().Unix())

    var whoFirst string
    var holeStr string
    var player1, player2 Player

    fmt.Scanln(&holeStr)
    if (len(holeStr) == 0) {
        return
    }
    i, j := textToCoord(holeStr)
    fmt.Scanln(&whoFirst)
    // writeFile(holeStr)
    // writeFile(whoFirst)


    if whoFirst == "black" {
        player1 = NewMCTSPlayer("Player1")
        player2 = NewOpponentPlayer("Player2")
    } else {
        player1 = NewOpponentPlayer("Player1")
        player2 = NewMCTSPlayer("Player2")
    }
    // player1 = NewAIPlayer("Player1")
    // player2 = NewRandomPlayer("Player2")
    // player1 = NewOpponentPlayer("Player1")
    // player2 = NewOpponentPlayer("Player2")

    startFight([2]int{i, j}, player1, player2)

    // wait until this process is killed
    for {}
}


func startFight(hole [2]int, player1 Player, player2 Player){
    model := NewAntiGame(hole)
    model.Start(player1, player2)

    var move string 
    for !model.IsEndGame() {
        move = model.CurrentPlayer().GetMove(model)
        if move == "pass" {
            model.PassMove()
        } else {
            i, j := textToCoord(move)
            model.Move(i, j)
        }
        if DEBUG {
            fmt.Println(move)
            fmt.Println("Move:", model.MoveNum())
            fmt.Println("Moves:", model.GetAvaliableMoves())
            model.Dump()
        }
    }
}

func writeFile(x interface{}){
    str := fmt.Sprintf("%v", x) + "\n"
    f, _ := os.OpenFile("./log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()
    f.Write([]byte(str))
}

func coordToText(coord [2]int) string {
    i := string(coord[1] + 'A')
    j := strconv.Itoa(coord[0] + 1)
    return i + j
}


func textToCoord(text string) (int, int) {
    j := int(text[0] - 'A')
    i := int(text[1] - 49)
    return i, j
}