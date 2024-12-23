import os
import re
import numpy as np

INPUT_FILE = 'input.txt'

def parse_input():
    parsed = [] # [[(xA,yA), (xB,yB), (c1, c2)]]
    with open(INPUT_FILE, 'r') as f:
        content = f.read()

    tokens = content.split("\n\n")
    for t in tokens:
        datum = t.split('\n')
        elem = []
        for d in datum:
            elem.append([int(f) for f in re.findall('\d+', d)])
        parsed.append(elem)
    return parsed

def solve_eq(coefs):
    a1, a2 = coefs[0][0], coefs[1][0]
    b1, b2 = coefs[0][1], coefs[1][1]
    vars = np.array([[a1, a2], [b1, b2]])
    const = np.array([c + 10000000000000 for c in coefs[2]])
    try:
        a = np.linalg.solve(vars, const)
        if all(np.round(val,2).is_integer() for val in a):
            return a
    except np.linalg.LinAlgError:
        return []
    return []



def solve_part1(parsed):
    ans = 0
    for p in parsed:
        a = solve_eq(p)
        if len(a) > 0:
            ans += 3 * a[0] + a[1]
    return ans

def main():
    parsed_input = parse_input()
    ans = solve_part1(parsed_input)
    print(ans)

if __name__ == '__main__':
    main()