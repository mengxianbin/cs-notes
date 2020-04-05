```puml
@startuml

actor Client as c
boundary Server as s

c -> s: FIN
s -> c: ACK
s -> c: FIN
c -> s: ACK

@enduml
```

* can not send FIN and ACK together

---
