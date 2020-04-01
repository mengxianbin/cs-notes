[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Algorithm](https://mengxianbin.github.io/cs-notes/site/Algorithm) /
[Problems](https://mengxianbin.github.io/cs-notes/site/Algorithm/Problems) /
[Shortest Path](https://mengxianbin.github.io/cs-notes/site/Algorithm/Problems/Shortest%20Path) /
[Dijkstra](https://mengxianbin.github.io/cs-notes/site/Algorithm/Problems/Shortest%20Path/Dijkstra) /
[Complexity](https://mengxianbin.github.io/cs-notes/site/Algorithm/Problems/Shortest%20Path/Dijkstra/Complexity)

Graph(V, E):

## Time

Time = O(V)*T(extract-min) + O(E)*T(decrease-key)

| data structure | T(extract-min) | T(decrease-key) | complexity  |
|----------------|----------------|-----------------|-------------|
| array          | O(V)           | O(1)            | O(V^2)      |
| binary heap    | O(lgV)         | O(lgV)          | O(ElgV)     |
| Fobinacci heap | O(lgV)         | O(1)            | O(E + VlgV) |

## Space

---
