# Usage:
# python main.py < data.txt
import re
from dataclasses import dataclass
import sys
from copy import deepcopy


@dataclass
class Rule:
    def __init__(self, rule):
        cond, next = rule.split(":")
        self.rule = cond
        self.next = next

    def __call__(self, input):
        return eval(self.rule, {}, input)

    def __repr__(self):
        return f"Rule({self.rule}:{self.next})"

    def calibrate(self, input):
        key, val = re.split(r'<|>', self.rule)
        val = int(val)
        lhs = deepcopy(input)
        rhs = deepcopy(input)
        if '<' in self.rule:
            lhs[key][1] = val-1
            rhs[key][0] = val
        else:
            lhs[key][0] = val+1
            rhs[key][1] = val
        return lhs, rhs


@dataclass
class Workflow:
    def __init__(self, wf):
        (name, rest,) = wf.split("{")
        rest = rest.split("}")[0]

        self.name = name
        *ifs, els = rest.split(",")
        self.ifs = list(map(Rule, ifs))
        self.els = els

    def __call__(self, input):
        for rule in self.ifs:
            if rule(input):
                return rule.next
        return self.els

    def calibrate(self, input):
        result = []
        inp = input
        for rule in self.ifs:
            lhs, rhs = rule.calibrate(inp)
            inp = deepcopy(rhs)
            result.append((lhs, rule.next))

        result.append((inp, self.els))
        return result

    def __repr__(self):
        return f"Workflow({self.name},rules={self.ifs},else={self.els})"


@dataclass
class Workflows:
    def __init__(self, wfs):
        self.wfs = {wf.name: wf for wf in wfs}
        assert len(wfs) == len(self.wfs), 'unique name'

    def __call__(self, input):
        name = "in"

        while True:
            wf = self.wfs[name]
            out = wf(input)
            match out:
                case 'R' | 'A':
                    return out == 'A'
                case next_wf:
                    name = next_wf

    def calibrate(self, input):
        total = 0
        q = [(input, self.wfs['in'])]
        while len(q):
            inp, wf = q.pop()
            opt = wf.calibrate(inp)
            for out, next_wf in opt:
                match next_wf:
                    case 'A':
                        t = 1
                        for a, b in out.values():
                            assert b > a, 'invalid range'
                            # The +1 is because the range is inclusive.
                            t *= (b - a) + 1 
                        total += t
                    case 'R':
                        continue
                    case next_wf:
                        q.append((out, self.wfs[next_wf]))
        return total


workflows: list[Workflow] = []
inputs = []
is_wf = True
keys = list('xmas')
for line in sys.stdin:
    line = line.strip()
    if line == "":
        is_wf = False
        continue
    if is_wf:
        workflows.append(Workflow(line))
    else:
        digits = map(int, re.findall(r"\d+", line))
        inputs.append(dict(zip(keys, digits)))


wfs = Workflows(workflows)
print('Part 1:', sum(sum(input.values()) if wfs(input) else 0 for input in inputs)) 
# Part 1: 19114, 434147

inp = {k: [1, 4000] for k in keys}
print('Part 2:', wfs.calibrate(inp))
# Part 2: 167409079868000, 136146366355609
