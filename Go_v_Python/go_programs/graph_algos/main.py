import random
from collections import deque

graph = {}
def createRandomGraph(n=1000,p=0.1): #n nodes, p is some random threshold 
    global graph
    for u in range(n):
        for v in range(u+1,n):
            if random.random() < p: #generate random num between 0 and 1 
                if u not in graph:
                    graph[u] = [v]
                else:
                    graph[u].append(v)
                if v not in graph:
                    graph[v] = [u]
                else:
                    graph[v].append(u)

# implement iterative bfs
def bfs(start=0,end=999):
    # return shortest path between start and end nodes
    global graph
    queue = deque([(0,start)]) # use deque data structure - same in golang https://pkg.go.dev/github.com/gammazero/deque 
    visited = set({start}) # keep track of visited set of nodes
    while(queue):
        steps, node = queue.popleft() #popleft steps and the current node
        if node == end:
            return steps
        for neigh in graph[node]: # traverse all the neighbours of this node
            if neigh not in visited: # add if not already visited
                queue.append((steps+1,neigh))
                visited.add(neigh)
    return -1 # path not found between start and end nodes


if __name__=="__main__":
    n,p = 10000,0.1
    start,end = 0,999
    createRandomGraph(n,p)
    steps = bfs(start,end)
    print(f"shortest distance between node {start} and node {end} is {steps}")
