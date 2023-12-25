from math import prod
import networkx as nx

G = nx.Graph()

with open("day25.txt") as f:
    for line in f:
        line = line.replace(":", "").strip()
        v, *adj = line.split(' ')
        for a in adj:
            G.add_edge(v, a)

G.remove_edges_from(nx.minimum_edge_cut(G))

# part 1: 54
# part 2: 551196
print(prod([len(c) for c in nx.connected_components(G)]))
