[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Message Queue](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Message%20Queue) /
[RocketMQ](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Message%20Queue/RocketMQ) /
[Deploy](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Message%20Queue/RocketMQ/Deploy) /
[Physical](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Message%20Queue/RocketMQ/Deploy/Physical)

```puml
@startuml

cloud NameServer as ns

cloud Producers as ps
cloud Brokers as bs
cloud Consumers as cs

ns -- ps
ns -- bs
ns -- cs

ps - bs
bs - cs

@enduml
```

```puml
@startuml

cloud Brokers as bs {
    node Master1 as m1
    node Master2 as m2
    node Slave1 as s1
    node Slave2 as s2

    m1 - s1
    m2 - s2

    m1 -[hidden]- m2
}

@enduml
```

