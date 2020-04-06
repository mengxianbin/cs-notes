[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Network](https://mengxianbin.github.io/cs-notes/site/Network) /
[TCP](https://mengxianbin.github.io/cs-notes/site/Network/TCP) /
[terminate](https://mengxianbin.github.io/cs-notes/site/Network/TCP/terminate)

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
