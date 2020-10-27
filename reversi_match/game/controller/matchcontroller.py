from ..model.game import Game
from ..model.antigame import AntiGame
from .player.matchplayer import AIPlayer, OpponentPlayer, SimplePlayer
from .gamemode import GameMode
import random
import time


class MatchController():

    def __init__(self, black_hole=None):
        self.gamemodel = AntiGame(black_hole)


    def create_players(self, who_first):
        if who_first == 'black':
            player1 = AIPlayer('Player1')
            player2 = OpponentPlayer('Player2')
        else:
            player1 = OpponentPlayer('Player1')
            player2 = AIPlayer('Player2')

        # player1 = AIPlayer('Player1')
        # player2 = SimplePlayer('Player2')
        return player1, player2


    def start(self, who_first):
        # get parametrs from conslon and create game model 
        players = self.create_players(who_first)
        self.gamemodel.start(*players)

        # game loop
        while not self.gamemodel.is_game_over:
            move = self.gamemodel.current_player.get_move(self.gamemodel)
            if move == 'pass':
                self.gamemodel.pass_move()
            else:
                move_tuple = text_coord_to_tuple(move)
                self.gamemodel.move(*move_tuple)


def text_coord_to_tuple(coord):
    j = ord(coord[0]) - ord('A')
    i = int(coord[1]) - 1
    return (i, j)