package main


import (
    "math/rand"
)


type Tree struct {
    root *Node
}

func NewTree(model *AntiGame) *Tree {
    return &Tree{NewNode(model, nil)}
}

func (s *Tree) Root() *Node { return s.root }
func (s *Tree) SetRoot(newRoot *Node) { s.root = newRoot }


type Node struct {
    model *AntiGame
    parent *Node
    childs []*Node
    visitCount int
    winScore float64
}


func NewNode(model *AntiGame, parent *Node) *Node {
    return &Node{model, parent, make([]*Node, 0), 0, 0}
}

func (s *Node) AddNode(node *Node) {
    s.childs = append(s.childs, node)
}

func (s *Node) AddModel(model *AntiGame) {
    // add new node by model
    s.AddNode(NewNode(model, s))
}


func (s *Node) getMaxWinScoreChild() *Node {
    // searching child node with highest win score
    maxWinScore := s.childs[0].winScore
    var maxNode *Node = s.childs[0]
    for _, node := range s.childs {
        if node.winScore > maxWinScore {
            maxWinScore = node.winScore
            maxNode = node
        }
    }
    return maxNode
}

func (s *Node) Expand() {
    // add all possible new nodes
    moves := s.model.GetAvaliableMoves()
    if len(moves) == 0 {
        newModel := s.model.Copy()
        newModel.PassMove()
        s.AddModel(newModel)
        return
    }
    for _, move := range moves {
        newModel := s.model.Copy()
        newModel.Move(move[0], move[1])
        s.AddModel(newModel)
    }
}

func (s *Node) RandomChild() *Node {
    return s.childs[rand.Intn(len(s.childs))]
}


func (s *Node) Model() *AntiGame { return s.model }
func (s *Node) Parent() *Node { return s.parent }
func (s *Node) Childs() []*Node { return s.childs }
func (s *Node) VisitCount() int { return s.visitCount }
func (s *Node) WinScore() float64 { return s.winScore }
func (s *Node) IncVisitCount() { s.visitCount++ }
func (s *Node) AddWinScore(val float64) { s.winScore += val }
func (s *Node) SetWinScore(val float64) { s.winScore = val }
