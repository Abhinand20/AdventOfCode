import os 
import collections
from typing import Tuple
from collections import defaultdict
from dataclasses import dataclass

INPUT_DIR = './input.txt'


class Equation():
    def __init__(self, op1, op2, val, operation):
        self.op1 = op1
        self.op2 = op2
        self.val = val
        self.operation = operation
    
    def __str__(self):
        return self.op1 + self.operation + self.op2 + "=" + self.val
    
    def solve(self, known: collections.defaultdict) -> Tuple[bool, int]:
        if self.op1 not in known or self.op2 not in known:
            return False, -1
        ans = -1
        o1, o2 = known[self.op1], known[self.op2]
        if self.operation == 'AND':
            ans = o1 & o2
        elif self.operation == 'OR':
            ans = o1 | o2
        elif self.operation == 'XOR':
            ans = o1 ^ o2
        return True, ans


def parse_input():
    known = defaultdict(int)
    unknowns = defaultdict(Equation)
    with open(INPUT_DIR, 'r') as f:
        content = f.read()
    
    init, circuit = content.split("\n\n")
    temp = init.split('\n')
    for token in temp:
        k, v = token.split(": ")
        known[k] = int(v)
    temp = circuit.split('\n')
    for t in temp:
        lhs, rhs = t.split(" -> ")
        op1, operation, op2 = lhs.split()
        unknowns[rhs] = Equation(op1, op2, rhs, operation)
    return known, unknowns
    

def recur(k, eq, knowns, unknowns):
    solved, ans = eq.solve(knowns)
    if solved:
        knowns[k] = ans
        return
    
    recur(eq.op1, unknowns[eq.op1], knowns, unknowns)
    recur(eq.op2, unknowns[eq.op2], knowns, unknowns)
    _, ans = eq.solve(knowns)
    knowns[k] = ans

def solve_part1(knowns, unknowns):
    for k, eq in unknowns.items():
        if k not in knowns:
            recur(k, eq, knowns, unknowns)
    all_bits = []
    for k in knowns.keys():
        if k.startswith('z'):
            all_bits.append(k)
    all_bits.sort()
    ans = 0
    for i, k in enumerate(all_bits):
        ans += (1 << i) * knowns[k]
    return ans


def main():
    knowns, unknowns = parse_input()
    ans = solve_part1(knowns, unknowns)
    print(ans)

if __name__ == '__main__':
    main()