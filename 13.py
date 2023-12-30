import sys
from copy import deepcopy


matrices: list[list[list[str]]] = []
lines: list[list[str]] = []
for line in sys.stdin:
    line = line.strip()
    if line == "":
        matrices.append(lines)
        lines = []
    else:
        lines.append(list(line))
if len(lines) > 0:
    matrices.append(lines)


def rotate_90_deg_clockwise(mat):
    return [list(row)[::-1] for row in zip(*mat)]


def reflection_score(mat):
    scores = []
    for m, mul in [(mat, 100), (rotate_90_deg_clockwise(mat), 1)]:
        n = len(m)
        for mid in range(n):
            valid = True
            epoch = 0
            for i, j in zip(range(n), range(1, n + 1)):
                if mid - i < 0 or mid + j >= n:
                    break
                epoch += 1
                valid = m[mid - i] == m[mid + j]
                if not valid:
                    break
            if valid and epoch:
                scores.append((mid + 1) * mul)
    return scores


total = sum([reflection_score(mat)[0] for mat in matrices])
print("Part 1:", total)


total = 0
for mat in matrices:
    old_score = reflection_score(mat)

    found = False
    r, c = len(mat), len(mat[0])
    for i in range(r):
        for j in range(c):
            mat2 = deepcopy(mat)
            # Invert.
            mat2[i][j] = "#" if mat2[i][j] == "." else "."

            new_score = reflection_score(mat2)
            if new_score is None:
                continue
            if (delta := set(new_score) - set(old_score)) and len(delta) > 0:
                total += list(delta)[0]
                found = True
                break
        if found:
            break
    if not found:
        total += old_score

print("Part 2:", total)
