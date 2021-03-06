package main

import (
    "math/rand"
)



type Minimax struct {
    player Player
    model *AntiGame
}

func NewMinimax(player Player, model *AntiGame) *Minimax {
    return &Minimax{player, model}
}


func (s *Minimax) GetBestMove(depth int) *Move {
    var score int
    var choseMove *Move
    var bestScore int = MAXINT
    model := s.model.Copy()

    moves := model.GetAvaliableMoves()

    for _, move := range moves {
        model.Move(move)
        score = s.minimax(model, depth, MININT, MAXINT, true)

        if score < bestScore {
            bestScore = score
            choseMove = move
        }
        model = s.model.Copy()
    }
    if bestScore == MAXINT {
        choseMove = moves[rand.Intn(len(moves))]
    }

    return choseMove
}


func (s *Minimax) minimax(srcModel *AntiGame, depth int, alpha int, beta int, isMinimizing bool) int {
    if depth == 0 || srcModel.IsEndGame() {
        return s.countScore(srcModel)
    }

    var score int
    var bestScore int
    model := srcModel.Copy()
    if isMinimizing {
        bestScore = MAXINT
        for _, move := range model.GetAvaliableMoves() {
            model.Move(move)
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
            model.Move(move)
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


func (s *Minimax) countScore(model *AntiGame) int {
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


func (s *Minimax) mapInd(ind int) int {
    if ind > 3 {
        return 7 - ind
    }
    return ind
}

func (s *Minimax) Min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func (s *Minimax) Max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}