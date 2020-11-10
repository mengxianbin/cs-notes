[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns) /
[Behavioral Patterns](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns) /
[Others](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns/Others) /
[Mediator](https://mengxianbin.github.io/cs-notes/site/Design/Design%20Patterns/Behavioral%20Patterns/Others/Mediator)


```plantuml
@startuml

interface Mediator {
    + requestA(a CommunicatorA)
    + requestB(b CommunicatorB)
    + requestC(c CommunicatorC)
}

interface CommunicatorA {
    + requestA(mediator: Mediator)
    + responseA1()
    + responseA2()
}

interface CommunicatorB {
    + requestB(mediator: Mediator)
    + responseB1()
}

interface CommunicatorC {
    + requestC(mediator: Mediator)
    + responseC1()
    + responseC2()
    + responseC3()
}

Mediator -- CommunicatorA
Mediator -- CommunicatorB
Mediator -- CommunicatorC

@enduml
```
