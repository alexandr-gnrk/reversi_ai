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


func main() {
    rand.Seed(time.Now().Unix())

    // model := NewAntiGame([2]int{0, 0})
    // model.Start(NewAIPlayer("Player1"), NewAIPlayer("Player2"))
    // mcts := NewMCTS(time.Millisecond * 200, model)
    // mcts.FindNextMove()
    
    var whoFirst string
    var holeStr string
    var player1, player2 Player


    fmt.Scanln(&holeStr)
    if (len(holeStr) == 0) {
        return
    }
    fmt.Scanln(&whoFirst)
    i, j := textToCoord(holeStr)
    // writeFile(holeStr)
    // writeFile(whoFirst)


    var isOpponentStarts bool = false
    if whoFirst == "black" {
        player1 = NewAIPlayer("Player1")
        player2 = NewOpponentPlayer("Player2")
    } else {
        isOpponentStarts = true
        player1 = NewOpponentPlayer("Player1")
        player2 = NewAIPlayer("Player2")
    }
    // player1 = NewAIPlayer("Player1")
    // player2 = NewAIPlayer("Player2")
    // player1 = NewOpponentPlayer("Player1")
    // player2 = NewOpponentPlayer("Player2")
    // controller := NewFightController([2]int{i, j})
    controller := NewFightController([2]int{i, j})
    controller.Start(player1, player2, isOpponentStarts)
// var cumMatrix [8][8]float64
// for i:=0; i < 100000; i++ {
//     x, y := rand.Intn(8), rand.Intn(8)
//     player1 = NewAIPlayer("Player1")
//     player2 = NewAIPlayer("Player2")
//     controller = NewFightController([2]int{x, y})
//     cumMatrix = addMatrix(
//         cumMatrix,
//         controller.Start(player1, player2, isOpponentStarts))
// }
// printMatrix(normMatrix(cumMatrix))
    // wait until this process is killed
    for {}
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