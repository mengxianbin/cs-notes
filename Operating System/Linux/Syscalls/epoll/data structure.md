* RB Tree
* List

---

```plantuml
@startuml

epoll_create -> RB Tree
epoll_ctl --> RB Tree
"RB Tree" ->[Trigger] List
List -up-> epoll_wait

@enduml
```

---
