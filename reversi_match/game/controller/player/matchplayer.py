import random
import copy
import sys
from .player import Player


    # EMPTY = 0
    # BLACK = 1
    # WHITE = 2
    # HOLE = 3

class AIPlayer(Player):
    def __init__(self, name):
        super().__init__(name)
        self.mask = [[4, 4, 4, 4, 4, 4, 4, 4], [4, 3, 3, 3, 3, 3, 3, 4], [4, 3, 2, 2, 2, 2, 3, 4], [4, 3, 2, 1, 1, 2, 3, 4], [4, 3, 2, 1, 1, 2, 3, 4], [4, 3, 2, 2, 2, 2, 3, 4], [4, 3, 3, 3, 3, 3, 3, 4], [4, 4, 4, 4, 4, 4, 4, 4]]
        
    def get_move(self, model):
        if self.pass_next:
            self.pass_next = False
            move = 'pass'
        else:
            # get list of available movements and choose random
            # moves = model.get_available_moves()
            # move = random.choice(moves)
            # move = tuple_to_text_coord(move)
            move = self.best_move(model)
            if move is None:
                move = 'pass'
            else:
                move = tuple_to_text_coord(move)
        print(move)
        return move

    def best_move(self, src_model, depth=2):
        best_score = sys.maxsize
        choose_move = None
        model = src_model.deepcopy()

        for move in model.get_available_moves():
            model.move(*move)
            score = self.minimax(model, depth, -sys.maxsize, sys.maxsize, False)
            # print('score:', score)
            if score < best_score:
                best_score = score
                choose_move = move
            model = src_model.deepcopy()

        # print('best_score:', best_score)
        return choose_move


    def minimax(self, src_model, depth, alpha, beta, is_minimizing):
        # check is win or depth
        if depth == 0 or src_model.is_end_game():
            return self.count_score(src_model)

        model = src_model.deepcopy()
        if is_minimizing:
            best_score = sys.maxsize
            for move in model.get_available_moves():
                model.move(*move)
                score = self.minimax(model, depth - 1, alpha, beta, False)
                best_score = min(score, best_score)
                beta = min(beta, best_score)
                if beta <= alpha:
                    break
                model = src_model.deepcopy()
        else:
            best_score = -sys.maxsize
            for move in model.get_available_moves():
                model.move(*move)
                score = self.minimax(model, depth - 1, alpha, beta, True)
                best_score = max(score, best_score)
                alpha = max(alpha, best_score)
                if beta <= alpha:
                    break
                model = src_model.deepcopy()
                

        # print(best_score)
        return best_score


    def count_score(self, model):

        length = len(model.board)
        score = 0
        for i in range(length):
            for j in range(length):
                if model.board[i][j] == self.color:
                    score += 1*self.mask[i][j]
        return score

    def ij_to_coeff(self, i, j):
        i = i-3 if i>3 else 4-i
        j = j-3 if j>3 else 4-j
        
        return max(i, j)

class SimplePlayer(Player):
    def __init__(self, name):
        super().__init__(name)
        
    def get_move(self, model):
        if self.pass_next:
            self.pass_next = False
            move = 'pass'
        else:
            # get list of available movements and choose random
            moves = model.get_available_moves()
            move = random.choice(moves)
            move = tuple_to_text_coord(move)
        return move

class OpponentPlayer(Player):
    def __init__(self, name):
        super().__init__(name)

    def get_move(self, model):
        # print(model)
        # import cProfile
        # cProfile.run('self.best_move(model)')
        # move = self.best_move(model)
        # return tuple_to_text_coord(move)
        return input()


def tuple_to_text_coord(coord):
    i = chr(coord[1] + ord('A'))
    j = str(coord[0] + 1)
    return i + j