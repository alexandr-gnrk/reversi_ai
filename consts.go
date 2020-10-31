package main

type Color int
const (
    EMPTY Color = iota
    WHITE
    BLACK
    HOLE
)

type Winner int
const (
    FIRST Winner = iota
    SECOND
    TIE
)

const (
    MAXINT = 1 << (32 - 1) - 1
    MININT = -MAXINT - 1
) 


var DIRECTIONS = [8][2]int{
    {-1, -1}, {-1, 0}, {-1, 1},
    { 0, -1}, { 0, 1},
    { 1, -1}, { 1, 0}, { 1, 1}}

// [0.67224, 0.97957, 0.91347, 0.91027,
// 0.97718, 0.99614, 0.93145, 0.95975,
// 0.91317, 0.93529, 0.93713, 0.92579,
// 0.90997, 0.96679, 0.92338]

// 90 %
var MASK = [8][8]int{
    {251, -55, 10, 13},
    {-55, -73,  -10, -3},
    {10, -10,  -14, -2},
    {13, -3,   -2},
}