[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Reactor](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Reactor) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Reactor/Netty) /
[Multiple Thread](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Reactor/Netty/Multiple%20Thread)

```puml
@startuml

storage ReactorThread {
    (Reactor)
    (Acceptor)
}

(Client) -> Reactor

Reactor -> (Reactor Thread Pool): Dispatcher

(Reactor Thread Pool) --> (Handlers): read\nencode
(Handlers) --> (Reactor Thread Pool): write\ndecode

@enduml
```
