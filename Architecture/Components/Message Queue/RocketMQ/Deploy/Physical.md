```plantuml
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

```plantuml
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

