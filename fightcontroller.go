package main

import (
    // "fmt"
    // "math"
)


type FightController struct {
    gamemodel *AntiGame
}


func NewFightController(hole [2]int) *FightController {
    return &FightController{NewAntiGame(hole)}
}


func (s *FightController) Start(player1 Player, player2 Player, isOpponentStarts bool){
    s.gamemodel.Start(player1, player2)

    var move string 
    for !s.gamemodel.IsEndGame() {
    // for {
        // fmt.Println(s.gamemodel.GetAvaliableMoves())
        move = s.gamemodel.CurrentPlayer().GetMove(s.gamemodel)
        if move == "pass" {
            s.gamemodel.PassMove()
        } else {
            i, j := textToCoord(move)
            s.gamemodel.Move(i, j)
        }
        // println(move)
        // fmt.Println("Move:", s.gamemodel.MoveNum())
        // fmt.Println("Moves:", s.gamemodel.GetAvaliableMoves())
        // s.gamemodel.Dump()
        // println()
        // println()
    }
// if s.gamemodel.CurrentPlayer().Point() < s.gamemodel.AnotherPlayer().Point() {
//     return s.gamemodel.CurrentPlayer().MovesMatrix()
// } else {
//     return s.gamemodel.AnotherPlayer().MovesMatrix()
// }
    // printMatrix()
    // opponent always should make finish move
    // var moves = s.gamemodel.MoveNum()
    // if !isOpponentStarts {
    //     moves--
    // }
    // if moves % 2 == 0 {
    //     var mockMove string
    //     fmt.Scanln(&mockMove)
    // }
}