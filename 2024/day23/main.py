import os 
from collections import defaultdict

INPUT_DIR = './input.txt'


def parse_input():
    lookup = defaultdict(list)
    with open(INPUT_DIR, 'r') as f:
        content = f

    return graph


def main():
    graph = parse_input()
    ans1 = solve_part1(graph)

    
if __name__ == '__main__':
    main()