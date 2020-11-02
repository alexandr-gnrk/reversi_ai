package main

import (
    "time"
    "math/rand"
)

// MonteCarloTreeSearch
type MCTS struct {
    calcTime time.Duration
    model *AntiGame
    opponent Player
}

func NewMCTS(calcTime time.Duration, model *AntiGame) *MCTS {
    return &MCTS{calcTime, model, model.AnotherPlayer()}
}



func (s *MCTS) FindNextMove() *Move {
    startTime := time.Now()
    moves := s.model.GetAvaliableMoves()
    len := len(moves)
    var wins []int = make([]int, len)
    
    type WinResult struct {
        moveNum int
        winner Player
    }

    ch := make(chan WinResult)
    done := make(chan struct{})
    for i := 0; i < len; i++ {
        go func(moveNum int, move *Move, c chan WinResult, done chan struct{}) {
            for {
                select {
                case ch <- WinResult{moveNum, s.randomPlay(s.model, move)}:
                case <- done:
                    return
                }
            } 
        }(i, moves[i], ch, done)
    }

    for time.Since(startTime) < s.calcTime {
        select {
        case res := <- ch:
            if res.winner != nil && !res.winner.IsEqual(s.opponent) {
                wins[res.moveNum]++
            }
        }
    }
    done <- struct{}{}
    
    maxWins := wins[0]
    maxMove := moves[0]
    allWins := 0
    for i, move := range moves {
        allWins += wins[i]
        if wins[i] > maxWins {
            maxWins = wins[i]
            maxMove = move
        }
    }
    // writeFile(allWins)
    // writeFile(wins)
    return maxMove
}


func (s *MCTS) randomPlay(srcModel *AntiGame , move *Move) Player {
    model := srcModel.Copy()
    model.Move(move)

    // simulate random game
    for !model.IsEndGame() {
        moves := model.GetAvaliableMoves()
        move := moves[rand.Intn(len(moves))]
        model.Move(move)
    }

    return model.GetWinner()
}


// func (s *MCTS) FindNextMove1() *Move {
//     startTime := time.Now()
//     moves := s.model.GetAvaliableMoves()
//     len := len(moves)
//     var wins []int = make([]int, len) 
//     i := 0
//     for time.Since(startTime) < s.calcTime {
//         winner := s.randomPlay(moves[i])
//         if winner != nil && !winner.IsEqual(s.opponent) {
//             wins[i]++
//         }
//         i++
//         if i == len {
//             i = 0
//         }
//     }
    
//     maxWins := wins[0]
//     maxMove := moves[0]
//     for i, move := range moves {
//         if wins[i] > maxWins {
//             maxWins = wins[i]
//             maxMove = move
//         }
//     }
//     writeFile(wins)
//     return maxMove
// }


// func (s *MCTS) randomPlay(move *Move) Player {
//     model := s.model.Copy()
//     model.Move(move)

//     // simulate random game
//     for !model.IsEndGame() {
//         moves := model.GetAvaliableMoves()
//         move := moves[rand.Intn(len(moves))]
//         model.Move(move)
//     }

//     return model.GetWinner()
// }
