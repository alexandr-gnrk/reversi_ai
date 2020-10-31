
# Reversi AI
This is an AI implementation for game [Anti Reversi](https://en.wikipedia.org/wiki/Reversi) with black hole.

## Setup
Clone the repository and change the working directory:

    git clone https://github.com/alexandr-gnrk/reversi_ai.git
    cd reversi

Build the game:

    go build

Test AI winrate:
    
    ./path_to_tester --command ./reversi_ai

## AI types

The game supports three types of AI players:
- Random moves
- [Minimax](https://en.wikipedia.org/wiki/Minimax) with [alpha–beta pruning](https://en.wikipedia.org/wiki/Alpha%E2%80%93beta_pruning)
- [Monte Carlo tree search](https://en.wikipedia.org/wiki/Monte_Carlo_tree_search)
