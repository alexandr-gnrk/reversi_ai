package main


import (
    "fmt"
    "time"
    "math/rand"
)


type Player interface {
    Name() string
    GetMove(model *AntiGame) *Move
    IncPoint()
    DecPoint()
    Point() int
    SetColor(color Color)
    Color() Color
    GetShadow() Player
    Real() Player
    IsEqual(p Player) bool
    LastMove() *Move
    SetLastMove(move *Move)
}


type AIPlayer struct {
    name string
    color Color
    point int
    lastMove *Move
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
    lastMove *Move
}


func NewAIPlayer(name string) *AIPlayer {
    return &AIPlayer{name, BLACK, 2, &Move{}}
}
func (s *AIPlayer) Name() string { return s.name }
func (s *AIPlayer) IncPoint() { s.point++ }
func (s *AIPlayer) DecPoint() { s.point-- }
func (s *AIPlayer) Point() int { return s.point }
func (s *AIPlayer) SetColor(color Color) { s.color = color }
func (s *AIPlayer) Color() Color { return s.color }
func (s *AIPlayer) Real() Player { return s }
func (s *AIPlayer) GetShadow() Player { return NewShadowPlayer(s) }
func (s *AIPlayer) IsEqual(p Player) bool { return s == p.Real() }
func (s *AIPlayer) LastMove() *Move { return s.lastMove }
func (s *AIPlayer) SetLastMove(move *Move) { s.lastMove = move }
func (s *AIPlayer) GetMove(model *AntiGame) *Move {
    moves := model.GetAvaliableMoves()
    move := moves[rand.Intn(len(moves))]
    fmt.Println(move.ToText())
    return move
}


func NewMinimaxPlayer(name string) *MinimaxPlayer {
    return &MinimaxPlayer{NewAIPlayer(name)}
}
func (s *MinimaxPlayer) GetMove(model *AntiGame) *Move  {
    move := NewMinimax(s, model).GetBestMove(3)
    fmt.Println(move.ToText())
    return move
}


func NewMCTSPlayer(name string) *MCTSPlayer {
    return &MCTSPlayer{NewAIPlayer(name)}
}
func (s *MCTSPlayer) GetMove(model *AntiGame) *Move  {
    mcts := NewMCTS(time.Millisecond * 300, model)
    move := mcts.FindNextMove()
    fmt.Println(move.ToText())
    return move
}

func NewOpponentPlayer(name string) *OpponentPlayer {
    return &OpponentPlayer{NewAIPlayer(name)}
}

func (s *OpponentPlayer) GetMove(model *AntiGame) *Move {
    var move string
    fmt.Scanln(&move)
    return (&Move{}).FromText(move)
}


func NewShadowPlayer(real Player) *ShadowPlayer {
    return &ShadowPlayer{real, real.Point(), real.LastMove()}
}

func (s *ShadowPlayer) IncPoint() { s.point++ }
func (s *ShadowPlayer) DecPoint() { s.point-- }
func (s *ShadowPlayer) Point() int { return s.point }
func (s *ShadowPlayer) Real() Player { return s.Player }
func (s *ShadowPlayer) LastMove() *Move { return s.lastMove }
func (s *ShadowPlayer) SetLastMove(move *Move) { s.lastMove = move }