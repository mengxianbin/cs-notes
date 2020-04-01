[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Summary](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Summary) /
[Zero Copy](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Summary/Zero%20Copy) /
[classic](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Summary/Zero%20Copy/classic)


```puml
@startuml

rectangle "User Context" as uc1 {
    rectangle "User Space" as uc1us {
        
    }
    rectangle "Kernal Space" as uc1ks {
        
    }
    rectangle "Hardware" as uc1hd {
        
    }

    uc1us -[hidden]- uc1ks
    uc1ks -[hidden]- uc1hd
}

rectangle "Kernel Context" as kc1 {
    rectangle "User Space" as kc1us {

    }
    rectangle "Kernal Space" as kc1ks {
        (kernel buffer) as kb
    }
    rectangle "Hardware" as kc1hd {
        (hardware drive) as hd
    }

    kc1us -[hidden]- kc1ks
    kc1ks -[hidden]- kc1hd
}

rectangle "User Context" as uc2 {
    rectangle "User Space" as uc2us {
        (user buffer) as ub
    }
    rectangle "Kernal Space" as uc2ks {

    }
    rectangle "Hardware" as uc2hd {

    }
    uc2us -[hidden]- uc2ks
    uc2ks -[hidden]- uc2hd
}

rectangle "Kernel Context" as kc2 {
    rectangle "User Space" as kc2us {

    }
    rectangle "Kernal Space" as kc2ks {
        (socket buffer) as sb
    }
    rectangle "Hardware" as kc2hd {
        (protocol engine) as pe
    }
    kc2us -[hidden]- kc2ks
    kc2ks -[hidden]- kc2hd
}

uc1 -[hidden] kc1
kc1 -[hidden] uc2
uc2 -[hidden] kc2

hd -[norank]-> kb: DMA Copy
kb -[norank]-> ub: CPU Copy
ub -[norank]-> sb: CPU Copy
sb -[norank]-> pe: DMA Copy

@enduml
```

---

* do copy: 4 times
* switch context: 3 times

---
