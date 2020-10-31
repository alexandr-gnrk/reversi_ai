package main

import (
    "time"
    "math"
    "fmt"
    "math/rand"
)

// MonteCarloTreeSearch
type MCTS struct {
    calcTime time.Duration
    tree *Tree
    opponent Player
}

func NewMCTS(calcTime time.Duration, model *AntiGame) *MCTS {
    return &MCTS{calcTime, NewTree(model), model.AnotherPlayer()}
}


func (s *MCTS) FindNextMove() [2]int {
    startTime := time.Now()
    for time.Since(startTime) < s.calcTime {
        promisingNode := s.selectPromisingNode()
        if !promisingNode.model.IsEndGame() {
            promisingNode.Expand()
        }
        exploringNode := promisingNode
        if len(exploringNode.Childs()) > 0 {
            exploringNode = promisingNode.RandomChild()
        }
        winner := s.randomPlayResult(exploringNode)
        s.backPropagate(exploringNode, winner)
    }

    if DEBUG {
        fmt.Print("Root: ")
        fmt.Println(s.tree.Root())
        fmt.Println("Childs: ")
        for _, node := range s.tree.Root().Childs() {
            fmt.Println("\t", node)
        }
    }

    
    // root could be a nill
    // all node is loosing
    s.tree.SetRoot(s.tree.Root().getMaxWinScoreChild())
    // if s.tree.Root() == nil {
    //     return s.tree.Root().RandomChild().Model().LastMove()
    // }
    return s.tree.Root().Model().LastMove()
}


func (s *MCTS) randomPlayResult(node *Node) Player {
    model := node.Model().Copy()
    winner := model.GetWinner()
    if model.IsEndGame() && (winner == nil || winner.IsEqual(s.opponent)) {
        node.Parent().SetWinScore(MININT)
        return model.GetWinner()
    }
    for !model.IsEndGame() {
        moves := model.GetAvaliableMoves()
        if len(moves) > 0 {
            move := moves[rand.Intn(len(moves))]
            model.Move(move[0], move[1])
        } else {
            model.PassMove()
        }
    }

    return model.GetWinner()
}

func (s *MCTS) selectPromisingNode() *Node {
    node := s.tree.Root()
    for len(node.Childs()) != 0 {
        node = s.findBestNodeByUCT(node)
    }
    return node
}


func (s *MCTS) backPropagate(node *Node, player Player) {
    // fmt.Println("Player:", player)
    for node != nil {
        node.IncVisitCount()
        // if player != nil && node.Model().CurrentPlayer().IsEqual(player) {
        if player != nil && !s.opponent.IsEqual(player) {
            node.AddWinScore(10)
        }
        node = node.Parent()
    }
}


func (s *MCTS) findBestNodeByUCT(node *Node) *Node {
    parentVisit := node.VisitCount()
    maxUCT := float64(MININT)
    var maxNode *Node
    for _, currNode := range node.Childs() {
        UCT := s.uctValue(parentVisit, currNode.WinScore(), currNode.VisitCount())
        if UCT > maxUCT {
            maxUCT = UCT
            maxNode = currNode
        }
    }
    // maxNode := node.Childs()[rand.Intn(len(node.Childs()))]
    return maxNode
}

func (s *MCTS) uctValue(parentVisit int, winScore float64, visit int) float64 {
    if visit == 0 {
        return float64(MAXINT)
    }
    return (winScore / float64(visit)) + 1.41 * math.Sqrt(math.Log(float64(parentVisit)) / float64(visit))
}