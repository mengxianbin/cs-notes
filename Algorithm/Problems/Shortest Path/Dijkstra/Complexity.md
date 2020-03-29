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
