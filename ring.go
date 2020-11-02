package main

import (
    "fmt"
    "os"
    "math/rand"
    "time"
    "flag"
    "log"
    "runtime/pprof"
)

const DEBUG = false
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")


func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close()
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }

    rand.Seed(time.Now().Unix())
    var myColor, holeStr string
    var player1, player2 Player

    fmt.Scanln(&holeStr)
    // end if hole pos is incorrect
    if (len(holeStr) == 0) {
        return
    }
    fmt.Scanln(&myColor)

    player1 = NewOpponentPlayer("Player1")
    // player2 = NewOpponentPlayer("Player2")
    // player2 = NewMCTSPlayer("Player2")
    // player1 = NewMinimaxPlayer("Player1")
    player2 = NewMCTSPlayer("Player2")
    if myColor == "black" {
        player1, player2 = player2, player1 
    }

    startFight(
        (&Cell{}).FromText(holeStr), 
        player1, 
        player2)
}


func startFight(hole *Cell, player1 Player, player2 Player){
    model := NewAntiGame(hole)
    model.Start(player1, player2)

    for {
        move := model.CurrentPlayer().GetMove(model)
        model.Move(move)
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