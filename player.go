package main


import (
    "fmt"
    // "math/rand"
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
        ai := NewAI(s, model)
        move = coordToText(ai.GetBestMove(3))
        // moves := model.GetAvaliableMoves()
        // movePos := moves[rand.Intn(len(moves))]
        // move = coordToText(movePos)
        // s.movesMatrix[movePos[0]][movePos[1]] = 1
    }
    fmt.Println(move)
    return move
}


type OpponentPlayer struct { 
    *AIPlayer 
}

func NewOpponentPlayer(name string) *OpponentPlayer {
    return &OpponentPlayer{NewAIPlayer(name)}
}

func (s *OpponentPlayer) GetMove(model *AntiGame) string {
    var move string
    fmt.Scanln(&move)
    return move
}


type ShadowPlayer struct {
    Player
    point int
    passNext bool
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


func printMatrix(matrix [8][8]float64) {
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            // fmt.Print(matrix[i][j], " ")
            fmt.Printf("%.5f ", matrix[i][j])
        }
        fmt.Println()
    }
}


func addMatrix(matrix1 [8][8]float64, matrix2 [8][8]float64) [8][8]float64 {
    var matrix3 [8][8]float64
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            matrix3[i][j] = matrix1[i][j] + matrix2[i][j]
        }
    }
    return matrix3
}


func normMatrix(matrix [8][8]float64) [8][8]float64 {
    var max float64 = matrix[0][0]
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            if matrix[i][j] > max {
                max = matrix[i][j]
            }
        }
    }
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            matrix[i][j] /= max
        }
    }
    return matrix
}



