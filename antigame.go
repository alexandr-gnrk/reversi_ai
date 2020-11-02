package main

import (
    "strconv"
)

type Cell struct {
    i int
    j int
}

func NewCell(i int, j int) *Cell {
    return &Cell{i, j}
}

func (s *Cell) IsEqual(cell *Cell) bool {
    if s.i == cell.i && s.j == cell.j {
        return true
    }
    return false
}

func (s *Cell) Unpack() (int, int) {
    return s.i, s.j
}

func (s *Cell) ToText() string {
    i := string(s.j + 'A')
    j := strconv.Itoa(s.i + 1)
    return i + j
}

func (s *Cell) FromText(text string) *Cell {
    s.j = int(text[0] - 'A')
    s.i = int(text[1] - 49)
    return s
}

type Move struct {
    cell *Cell
}

func NewMove(cell *Cell) *Move {
    return &Move{cell}
}


func (s *Move) IsPass() bool {
    if s.cell == nil {
        return true
    }
    return false
}


func (s *Move) IsEqual(move *Move) bool {
    if (s.cell == nil && move.Cell() != nil) {
        return false
    }
    if (s.cell == nil && move.Cell() == nil) || s.cell.IsEqual(move.Cell()) {
        return true
    }
    return false
}

func (s *Move) Cell() *Cell {
    return s.cell
}

func (s *Move) Unpack() (int, int) {
    if s.cell == nil {
        panic("trying unpack \"pass\" Move")
    }
    return s.cell.Unpack()
}

func (s *Move) ToText() string {
    if s.cell == nil {
        return "pass"
    }
    return s.cell.ToText()
}

func (s *Move) FromText(text string) *Move {
    if text == "pass" {
        s.cell = nil
        return s
    }
    s.cell = (&Cell{}).FromText(text)
    return s
}

func (s *Move) Copy() *Move {
    return &Move{s.cell}
}


type AntiGame struct {
    dimension int
    currentPlayer Player
    anotherPlayer Player
    board [8][8]Color
    hole *Cell
    moveNum int
    isGameOver bool
}

func NewAntiGame(hole *Cell) *AntiGame {
    res := &AntiGame{}
    res.dimension = 8
    res.currentPlayer = nil
    res.anotherPlayer = nil
    res.hole = hole
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
    s.board[s.hole.i][s.hole.j] = HOLE
}

func (s *AntiGame) CurrentPlayer() Player { return s.currentPlayer }
func (s *AntiGame) AnotherPlayer() Player { return s.anotherPlayer }
func (s *AntiGame) MoveNum() int { return s.moveNum }
func (s *AntiGame) Copy() *AntiGame {
    cp := NewAntiGame(s.hole)
    cp.currentPlayer = s.currentPlayer.GetShadow()
    cp.anotherPlayer = s.anotherPlayer.GetShadow()
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

func (s *AntiGame) GetAvaliableMoves() []*Move {
    avaliable := make([]*Move, 0)
    for i := 0; i < s.dimension; i++ {
        for j := 0; j < s.dimension; j++ {
            if s.isAvaliableCell(i, j) {
                avaliable = append(avaliable, NewMove(NewCell(i, j)))
            }
        }
    }
    // add pass move if there are not other moves
    if len(avaliable) == 0 {
        avaliable = append(avaliable, NewMove(nil))
    }
    return avaliable
}

func (s *AntiGame) isAvaliableCell(i int, j int) bool {
    if s.board[i][j] == HOLE || s.board[i][j] != EMPTY {
        return false
    }

    for _, direction := range DIRECTIONS {
        if s.isLineBounded(i, j, direction[0], direction[1]) {
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

func (s *AntiGame) Move(move *Move) {
    s.currentPlayer.SetLastMove(move)
    if move.IsPass() {
        s.moveNum++
        s.changePlayer()
        return
    }

    s.moveNum++

    i, j := move.Unpack()
    s.updateLines(i, j)
    s.reverseCell(i, j)

    s.currentPlayer.IncPoint()

    s.changePlayer()
}


func (s *AntiGame) IsEndGame() bool {
    cond1 := s.GetAvaliableMoves()
    if len(cond1) == 1 && cond1[0].IsPass() {
        s.changePlayer()
        cond2 := s.GetAvaliableMoves()
        s.changePlayer()
        if len(cond2) == 1 && cond1[0].IsPass() {
            return true
        }
    }
    return false 
}


func (s *AntiGame) endGame() {
    s.isGameOver = true
}

func (s *AntiGame) IsIn(moves []*Move, searchMove *Move) bool {
    for _, move := range moves {
        if move.IsEqual(searchMove) {
            return true
        }
    }
    return false
}

func (s *AntiGame) GetWinner() Player {
    if s.currentPlayer.Point() < s.anotherPlayer.Point() {
        return s.CurrentPlayer()
    } else if s.currentPlayer.Point() > s.anotherPlayer.Point() {
        return s.AnotherPlayer()
    }
    return nil
}


func (s *AntiGame) Dump() {
    avaliable := s.GetAvaliableMoves()
    println("  A B C D E F G H")
    for i := 0; i < s.dimension; i++ {
        print(i+1, " ")
        for j := 0; j < s.dimension; j++ {
            if s.IsIn(avaliable, NewMove(NewCell(i, j))) {
                print("X", " ")
            } else {
                print(s.board[i][j], " ")
            }
        }
        println()
    }
}