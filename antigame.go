package main

import (
    // "fmt"
    // "math"
)


type AntiGame struct {
    dimension int
    currentPlayer Player
    anotherPlayer Player
    board [8][8]Color
    holePos [2]int
    moveNum int
    isGameOver bool
}

var DIRECTIONS = [8][2]int{
    {-1, -1}, {-1, 0}, {-1, 1},
    { 0, -1}, { 0, 1},
    { 1, -1}, { 1, 0}, { 1, 1}}


func NewAntiGame(holePos [2]int) *AntiGame {
    res := &AntiGame{}
    res.dimension = 8
    res.currentPlayer = nil
    res.anotherPlayer = nil
    res.holePos = holePos
    res.moveNum = 0
    res.isGameOver = false
    return res
}


func (s *AntiGame) Start(player1 Player, player2 Player) {
    s.isGameOver = false
    s.currentPlayer = player1
    s.anotherPlayer = player2
    s.currentPlayer.SetColor(BLACK)
    s.anotherPlayer.SetColor(WHITE)
    s.initialPlacement()
}

func (s *AntiGame) initialPlacement() {
    s.board[s.dimension/2 - 1][s.dimension/2 - 1] = WHITE
    s.board[s.dimension/2 - 1][s.dimension/2] = BLACK
    s.board[s.dimension/2][s.dimension/2 - 1] = BLACK
    s.board[s.dimension/2][s.dimension/2] = WHITE
    s.board[s.holePos[0]][s.holePos[1]] = HOLE
}

func (s *AntiGame) CurrentPlayer() Player { return s.currentPlayer }
func (s *AntiGame) AnotherPlayer() Player { return s.anotherPlayer }
func (s *AntiGame) MoveNum() int { return s.moveNum }
func (s *AntiGame) Copy() *AntiGame {
    cp := NewAntiGame(s.holePos)
    cp.currentPlayer = s.currentPlayer.Copy()
    cp.anotherPlayer = s.anotherPlayer.Copy()
    cp.board = s.BoardCopy()
    cp.moveNum = s.moveNum
    cp.isGameOver = s.isGameOver
    return cp
}
func (s *AntiGame) BoardCopy() [8][8]Color {
    var cp [8][8]Color
    for i := range s.board {
        copy(cp[i][:], s.board[i][:])
    }
    return cp
}


func (s *AntiGame) changePlayer() {
    s.currentPlayer, s.anotherPlayer = s.anotherPlayer, s.currentPlayer
}

func (s *AntiGame) GetAvaliableMoves() [][2]int {
    avaliable := make([][2]int, 0)
    for i := 0; i < s.dimension; i++ {
        for j := 0; j < s.dimension; j++ {
            if s.isAvaliableCell(i, j) {
                avaliable = append(avaliable, [2]int{i, j})
            }
        }
    }
    return avaliable
}

func (s *AntiGame) isAvaliableCell(i int, j int) bool {
    if s.board[i][j] == HOLE || s.board[i][j] != EMPTY {
        return false
    }

    for _, direction := range DIRECTIONS {
        iDiff, jDiff := direction[0], direction[1]
        if s.isLineBounded(i, j, iDiff, jDiff) {
            return true
        }
    }
    return false
}

func (s *AntiGame) isCellExist(i int, j int) bool {
    if i >= 0 && j >= 0 && i < s.dimension && j < s.dimension && s.board[i][j] != HOLE {
        return true
    }
    return false
}

func (s *AntiGame) isLineBounded(i int, j int, iDiff int, jDiff int) bool {
    i += iDiff
    j += jDiff
    amount := 0

    for s.isCellExist(i, j) && s.board[i][j] == s.anotherPlayer.Color() {
        i += iDiff
        j += jDiff
        amount++
    }

    if s.isCellExist(i, j) && s.board[i][j] == s.currentPlayer.Color() && amount > 0 {
        return true
    }
    return false
}

func (s *AntiGame) reverseLine(i int, j int, iDiff int, jDiff int) {
    i += iDiff
    j += jDiff
    for s.board[i][j] == s.anotherPlayer.Color() {
        s.reverseCell(i, j)
        i += iDiff
        j += jDiff
    }
}

func (s *AntiGame) reverseCell(i int, j int) {
    s.board[i][j] = s.currentPlayer.Color()
    s.currentPlayer.IncPoint()
    s.anotherPlayer.DecPoint()
}

func (s *AntiGame) updateLines(i int, j int) {
    for _, direction := range DIRECTIONS {
        iDiff, jDiff := direction[0], direction[1]
        if s.isLineBounded(i, j, iDiff, jDiff) {
            s.reverseLine(i, j, iDiff, jDiff)
        }
    }
}

func (s *AntiGame) Move(i int, j int) {
    if !s.IsIn(s.GetAvaliableMoves(), [2]int{i, j}) {
        print("Incorrect move!\n")
        return
    }
    s.moveNum++
    s.updateLines(i, j)
    s.board[i][j] = s.currentPlayer.Color()

    s.currentPlayer.IncPoint()

    s.changePlayer()
    if len(s.GetAvaliableMoves()) == 0 {
        s.currentPlayer.SetPassNext(true)
        return
    }
    // if s.IsEndGame() {
    //     s.endGame()
    // }
}


func (s *AntiGame) PassMove() {
    s.moveNum++
    s.changePlayer()
}

func (s *AntiGame) IsEndGame() bool {
    cond1 := s.GetAvaliableMoves()
    if len(cond1) == 0 {
        s.changePlayer()
        cond2 := s.GetAvaliableMoves()
        s.changePlayer()
        if len(cond2) == 0 {
            return true
        }
    }
    return false 
}


func (s *AntiGame) endGame() {
    s.isGameOver = true
}

func (s *AntiGame) IsIn(moves [][2]int, searchMove [2]int) bool {
    for _, move := range moves {
        if move == searchMove { return true }
    }
    return false
}


func (s *AntiGame) Dump() {
    avaliable := s.GetAvaliableMoves()
    for i := 0; i < s.dimension; i++ {
        for j := 0; j < s.dimension; j++ {
            if s.IsIn(avaliable, [2]int{i, j}) {
                print("X", " ")
            } else {
                print(s.board[i][j], " ")
            }
        }
        println()
    }
}