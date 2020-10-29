package main

import "math"
const (
    MAXINT = 1 << (32 - 1) - 1
    MININT = -MAXINT - 1
) 

// 0.66 0.99 0.90 0.90
// 0.97 0.98 0.92 0.96
// 0.90 0.93 0.93 0.90
// 0.89 0.96 0.91 0.00

// [0.67224, 0.97957, 0.91347, 0.91027,
// 0.97718, 0.99614, 0.93145, 0.95975,
// 0.91317, 0.93529, 0.93713, 0.92579,
// 0.90997, 0.96679, 0.92338]

// [-2.5076, 0.56570, -0.0953, -0.1273, 
// 0.54180, 0.73139, 0.08449, 0.36749, 
// -0.09830, 0.12289, 0.14129, 0.02789,
// -0.13030, 0.43789, 0.00379]

// var MASK = [8][8]int{
//     {251, -55, 10, 13},
//     {-55, -73,  -10, -40},
//     {10, -10,  -14, -2},
//     {13, -40,   -2},}
var MASK = [8][8]int{
    {251, -55, 10, 13},
    {-55, -73,  -10, -3},
    {10, -10,  -14, -2},
    {13, -3,   -2},
}

type AI struct {
    player Player
    model *AntiGame
}

func NewAI(player Player, model *AntiGame) *AI {
    return &AI{player, model}
}


func (s *AI) GetBestMove(depth int) [2]int {
    var score int
    var choseMove [2]int
    var bestScore int = MAXINT
    model := s.model.Copy()

    writeFile(model.GetAvaliableMoves())
    for _, move := range model.GetAvaliableMoves() {
        model.Move(move[0], move[1])
        score = s.minimax(model, depth, MININT, MAXINT, true)

        if score < bestScore {
            bestScore = score
            choseMove = move
        }
        model = s.model.Copy()
    }
    writeFile(choseMove)
    return choseMove
}


func (s *AI) minimax(srcModel *AntiGame, depth int, alpha int, beta int, isMinimizing bool) int {
    if depth == 0 || srcModel.IsEndGame() {
        return s.countScore(srcModel)
    }

    var score int
    var bestScore int
    model := srcModel.Copy()
    if isMinimizing {
        bestScore = MAXINT
        for _, move := range model.GetAvaliableMoves() {
            model.Move(move[0], move[1])
            score = s.minimax(model, depth - 1, alpha, beta, false)
            bestScore = s.Min(score, bestScore)
            beta = s.Min(beta, bestScore)
            if beta <= alpha {
                break
            }
            model = srcModel.Copy()
        }
    } else {
        bestScore = MININT
        for _, move := range model.GetAvaliableMoves() {
            model.Move(move[0], move[1])
            score = s.minimax(model, depth - 1, alpha, beta, true)
            bestScore = s.Max(score, bestScore)
            alpha = s.Max(beta, bestScore)
            if beta <= alpha {
                break
            }
            model = srcModel.Copy()
        }
    }
    return bestScore
}


func (s *AI) countScore(model *AntiGame) int {
    var score int = 0
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            if model.board[i][j] == s.player.Color() {
                score += MASK[s.mapInd(i)][s.mapInd(j)]
            }
        }
    }
    return score
}


func (s *AI) getCoeff(i int, j int) int {
    if i > 3 {
        i -= 3
    } else {
        i = 4 - i
    }

    if j > 3 {
        j -= 3
    } else {
        j = 4 -j
    }
    return int(math.Pow(2.0, float64(s.Max(i, j))))
}

func (s *AI) mapInd(ind int) int {
    if ind > 3 {
        return 7 - ind
    }
    return ind
}

func (s *AI) Min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func (s *AI) Max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}