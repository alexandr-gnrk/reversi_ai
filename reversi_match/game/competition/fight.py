from ..controller.matchcontroller import MatchController

def fight():
    black_hole = text_coord_to_tuple(input())
    who_first = input() # black or white
    controller = MatchController(black_hole)
    
    # import pygame
    # from ..view.guiview import GUIView
    # pygame.init()
    # screen = pygame.display.set_mode((600, 600))
    # view = GUIView(screen)
    # controller.gamemodel.attach(view)
    
    controller.start(who_first)
    
    
def text_coord_to_tuple(coord):
    j = ord(coord[0]) - ord('A')
    i = int(coord[1]) - 1
    return (i, j)