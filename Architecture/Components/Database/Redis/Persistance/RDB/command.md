| command    | save         | bgsave                |
|------------|--------------|-----------------------|
| IO type    | sync         | async                 |
| blocking   | yes          | yes (when fork)       |
| complexity | O(n)         | O(n)                  |
| pro        | save memory  | not block client      |
| con        | block client | cost memory when fork |

---
