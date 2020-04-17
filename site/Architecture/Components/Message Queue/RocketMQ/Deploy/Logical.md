[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Message Queue](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Message%20Queue) /
[RocketMQ](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Message%20Queue/RocketMQ) /
[Deploy](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Message%20Queue/RocketMQ/Deploy) /
[Logical](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Message%20Queue/RocketMQ/Deploy/Logical)

```puml
@startuml

usecase Brokers as bs

rectangle {
    cloud ProducerGroupB as pgb
    cloud ProducerGroupA as pga
    cloud pgc

    pga -[hidden]- pgc
    pgc -[hidden]- pgb
    hide pgc
}

rectangle {
    cloud ConsumerGroupA as cga
    cloud cgc
    cloud ConsumerGroupB as cgb

    cga -[hidden]- cgc
    cgc -[hidden]- cgb
    hide cgc
}


pga -[norank]> bs: topic1
pgb -[norank]> bs: topic2, topic3
pgc -> bs
bs -[norank]> cga: topic1, topic3
bs -[norank]> cgb: topic2
bs -> cgc


cga -[hidden]-- cgb

@enduml
```

---
