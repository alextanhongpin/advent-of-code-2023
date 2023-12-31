# Usage:
# python main.py <filename: data.txt>
import sys


file_name = sys.argv[1:]

lines = []
with open(file_name[0], "r") as f:
    lines = [line.strip() for line in f]


grid = {}
dim = 0
for y, line in enumerate(lines):
    dim = max(dim, y)
    for x, char in enumerate(line):
        grid[x, y] = char
dim += 1


def dfs(q = []):
    seen = {}

    while q:
        x, y, dx, dy = q.pop()
        x = x + dx
        y = y + dy

        # Out of bound.
        if (x, y) not in grid:
            continue

        if (x, y, dx, dy) in seen:
            continue

        seen[x, y, dx, dy] = True

        match grid[x, y]:
            case '|':
                if dx != 0:
                    q.append((x, y, 0, 1))
                    q.append((x, y, 0, -1))
                if dy != 0:
                    q.append((x, y, dx, dy))
            case '-':
                if dx != 0:
                    q.append((x, y, dx, dy))
                if dy != 0:
                    q.append((x, y, -1, 0))
                    q.append((x, y, 1, 0))
            case '\\':
                if dx != 0:
                    q.append((x, y, 0, dx))
                if dy != 0:
                    q.append((x, y, dy, 0))
            case '/':
                if dx != 0:
                    q.append((x, y, 0, -dx))
                if dy != 0:
                    q.append((x, y, -dy, 0))
            case '.':
                q.append((x, y, dx, dy))
            case rest:
                raise ValueError(f"Unknown char: {rest}")
    return seen

def score(seen: set):
    unique = set()
    for x, y, dx, dy in seen:
        unique.add((x, y))
    return len(unique)

print('Part 1:', score(dfs([(-1, 0, 1, 0)])))


max_score = 0
for x in range(dim):
    # Move downwards.
    max_score = max(max_score, score(dfs([(x, 0, 0, 1)])))
    # Move upwards.
    max_score = max(max_score, score(dfs([(x, dim-1, 0, -1)])))

for y in range(dim):
    # Move right.
    max_score = max(max_score, score(dfs([(0, y, 1, 0)])))
    # Move left.
    max_score = max(max_score, score(dfs([(dim-1, y, -1, 0)])))

print('Part 2:', max_score)
