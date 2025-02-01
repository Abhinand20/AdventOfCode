import os 
from collections import defaultdict

INPUT_DIR = './input.txt'


def parse_input():
    graph = defaultdict(list)
    with open(INPUT_DIR, 'r') as f:
        for l in f.readlines():
            n1, n2 = l.rstrip('\n').split('-')  
            graph[n1].append(n2)
            graph[n2].append(n1)
    return graph

valid_cycles = set()

def find_cycles(graph, path):
    if len(path) == 4:
        if path[0] == path[-1]:
            ans = sorted(path[:-1])
            valid_cycles.add(','.join(ans))
        return
    curr = path[-1]
    for ngbs in graph[curr]:
        if len(path) > 1 and path[-1] == ngbs:
            continue
        path.append(ngbs)
        find_cycles(graph, path)
        path.pop()
    

def solve_part1(graph):
    for n in graph.keys():
        if n.startswith('t'):
            find_cycles(graph, [n])
    return len(valid_cycles)


longest_len = -1
ans = ""
def dfs(graph, path, visited):
    global longest_len
    global ans
    if len(path) > 1 and path[-1] == path[0]:
        l = len(path) - 1
        longest_len = max(longest_len, l)
        if longest_len == l:
            ans = ','.join(sorted(path[:-1]))
        return
    curr = path[-1]
    if curr in visited:
        return
    visited.add(curr)
    for ngb in graph[curr]:
        if path[-1] != ngb:
            dfs(graph, path + [ngb], visited)

def solve_part2(graph):
    visited = set()
    for n in graph.keys():
        if n not in visited:
            dfs(graph, [n], visited)
    return ans

def main():
    graph = parse_input()
    ans1 = solve_part1(graph)
    print(ans1)
    ans2 = solve_part2(graph)
    print(ans2)

    
if __name__ == '__main__':
    main()