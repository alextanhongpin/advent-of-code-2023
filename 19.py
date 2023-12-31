# Usage:
# python main.py < data.txt
import re
from typing import TypedDict
from dataclasses import dataclass
import sys
from copy import deepcopy

class Part(TypedDict):
    x: int
    m: int
    a: int
    s: int

@dataclass
class Rule:
    def __init__(self, rule: str):
        cond, next = rule.split(":")
        self.rule = cond
        self.next = next

    def __call__(self, part: Part):
        return eval(self.rule, {}, part)

    def __repr__(self):
        return f"Rule({self.rule}:{self.next})"

    def calibrate(self, part: Part):
        key, val = re.split(r'<|>', self.rule)
        val = int(val)
        lhs = deepcopy(part)
        rhs = deepcopy(part)
        if '<' in self.rule:
            lhs[key][1] = val-1
            rhs[key][0] = val
        else:
            lhs[key][0] = val+1
            rhs[key][1] = val
        return lhs, rhs


@dataclass
class Workflow:
    def __init__(self, wf: str):
        (name, rest,) = wf.split("{")
        rest = rest[:-1]

        self.name = name
        *rules, fallback = rest.split(",")
        self.rules = list(map(Rule, rules))
        self.fallback = fallback

    def __call__(self, part: Part):
        for rule in self.rules:
            if rule(part):
                return rule.next

        return self.fallback

    def calibrate(self, part: Part):
        result = []
        for rule in self.rules:
            lhs, rhs = rule.calibrate(part)
            part = rhs
            result.append((lhs, rule.next))

        result.append((part, self.fallback))
        return result

    def __repr__(self):
        return f"Workflow({self.name},rules={self.rules},else={self.fallback})"


@dataclass
class Workflows:
    def __init__(self, wfs):
        self.wfs = {wf.name: wf for wf in wfs}
        assert len(wfs) == len(self.wfs), 'unique name'

    def __call__(self, part):
        name = "in"

        while True:
            wf = self.wfs[name]
            out = wf(part)
            match out:
                case 'R' | 'A':
                    return out == 'A'
                case next_wf:
                    name = next_wf

    def calibrate(self, part: Part):
        result = []
        q = [(part, self.wfs['in'])]
        while len(q):
            inp, wf = q.pop()
            opt = wf.calibrate(inp)
            for out, next_wf in opt:
                match next_wf:
                    case 'A':
                        result.append(out)
                    case 'R':
                        continue
                    case next_wf:
                        q.append((out, self.wfs[next_wf]))
        return result


def calculate_combinations(parts: list[Part]) -> int:
    total = 0
    for part in parts:
        n = 1
        for lo, hi in part.values():
            # The +1 is because the range is inclusive.
            # Say you have value from 1-5, the range is 5-1+1 = 5.
            n *= hi - lo + 1
        total += n
    return total


section = '\n'.join([line.strip() for line in sys.stdin] )
top, btm = section.split('\n\n')

keys = list('xmas')
parts: list[Part] = [dict(zip(keys, map(int, re.findall(r"\d+", line)))) for line in btm.split('\n')]
wfs: Workflows = Workflows([Workflow(line) for line in top.split('\n')])


print('Part 1:', sum(sum(part.values()) if wfs(part) else 0 for part in parts)) 
# Part 1: 19114, 434147



inp = {k: [1, 4000] for k in keys}
print('Part 2:', calculate_combinations(wfs.calibrate(inp)))
# Part 2: 167409079868000, 136146366355609
