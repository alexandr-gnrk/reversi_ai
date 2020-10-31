package main


import (
    "fmt"
    "time"
    "math/rand"
    // "../antigame"
)


type Player interface {
    Name() string
    GetMove(model *AntiGame) string
    IncPoint()
    DecPoint()
    Point() int
    SetColor(color Color)
    Color() Color
    SetPassNext(passNext bool)
    PassNext() bool
    MovesMatrix() [8][8]float64
    GetShadow() Player
    Real() Player
    IsEqual(p Player) bool
}



type AIPlayer struct {
    name string
    color Color
    point int
    passNext bool
    movesMatrix [8][8]float64
}
type MinimaxPlayer struct { 
    *AIPlayer 
}
type MCTSPlayer struct { 
    *AIPlayer 
}
type OpponentPlayer struct { 
    *AIPlayer 
}
type ShadowPlayer struct {
    Player
    point int
    passNext bool
}


func NewAIPlayer(name string) *AIPlayer {
    return &AIPlayer{name, BLACK, 2, false, [8][8]float64{}}
}
func (s *AIPlayer) Name() string { return s.name }
func (s *AIPlayer) IncPoint() { s.point++ }
func (s *AIPlayer) DecPoint() { s.point-- }
func (s *AIPlayer) Point() int { return s.point }
func (s *AIPlayer) SetColor(color Color) { s.color = color }
func (s *AIPlayer) Color() Color { return s.color }
func (s *AIPlayer) SetPassNext(passNext bool) { s.passNext = passNext }
func (s *AIPlayer) PassNext() bool { return s.passNext }
func (s *AIPlayer) MovesMatrix() [8][8]float64 { return s.movesMatrix }
func (s *AIPlayer) Real() Player { return s }
func (s *AIPlayer) GetShadow() Player { return NewShadowPlayer(s) }
func (s *AIPlayer) IsEqual(p Player) bool { return s == p.Real() }
func (s *AIPlayer) GetMove(model *AntiGame) string {
    var move string
    if s.PassNext() {
        s.SetPassNext(false)
        move = "pass"
    } else {
        moves := model.GetAvaliableMoves()
        movePos := moves[rand.Intn(len(moves))]
        move = coordToText(movePos)
        s.movesMatrix[movePos[0]][movePos[1]] = 1
    }
    fmt.Println(move)
    return move
}


func NewMinimaxPlayer(name string) *MinimaxPlayer {
    return &MinimaxPlayer{NewAIPlayer(name)}
}
func (s *MinimaxPlayer) GetMove(model *AntiGame) string {
    var move string
        if s.PassNext() {
        s.SetPassNext(false)
        move = "pass"
    } else {
        ai := NewMinimax(s, model)
        move = coordToText(ai.GetBestMove(3))
    }
    fmt.Println(move)
    return move
}


func NewMCTSPlayer(name string) *MCTSPlayer {
    return &MCTSPlayer{NewAIPlayer(name)}
}
func (s *MCTSPlayer) GetMove(model *AntiGame) string {
    var move string
    mcts := NewMCTS(time.Millisecond * 1800, model)
    if s.PassNext() {
        s.SetPassNext(false)
        move = "pass"
    } else {
        move = coordToText(mcts.FindNextMove())
    }
    fmt.Println(move)
    return move
}

func NewOpponentPlayer(name string) *OpponentPlayer {
    return &OpponentPlayer{NewAIPlayer(name)}
}

func (s *OpponentPlayer) GetMove(model *AntiGame) string {
    var move string
    fmt.Scanln(&move)
    return move
}


func NewShadowPlayer(real Player) *ShadowPlayer {
    return &ShadowPlayer{real, real.Point(), real.PassNext()}
}

func (s *ShadowPlayer) IncPoint() { s.point++ }
func (s *ShadowPlayer) DecPoint() { s.point-- }
func (s *ShadowPlayer) Point() int { return s.point }
func (s *ShadowPlayer) SetPassNext(passNext bool) { s.passNext = passNext }
func (s *ShadowPlayer) PassNext() bool { return s.passNext }
func (s *ShadowPlayer) Real() Player { return s.Player }