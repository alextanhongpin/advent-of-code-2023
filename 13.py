# Usage:
# python main.py < data.txt
import sys

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


def reflection_score(mat, diff=0):
    for m, mul in [(mat, 100), (rotate_90_deg_clockwise(mat), 1)]:
        n = len(m)
        for mid in range(1, n):
            top = m[:mid]
            btm = m[mid:]
            lim = min(len(top), len(btm))

            top = top[-lim:]
            btm = btm[:lim]

            assert len(top) > 0
            assert len(btm) > 0
            assert len(top) == len(btm)

            delta = 0
            for i, j in zip(top, reversed(btm)):
                delta += sum(t != b for t, b in zip(i, j))
            if delta == diff:
                return mid * mul


total = sum([reflection_score(mat) for mat in matrices])
print("Part 1:", total)  # 35360

total = sum([reflection_score(mat, diff=1) for mat in matrices])
print("Part 2:", total)  # 36755
