[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Design](https://mengxianbin.github.io/cs-notes/site/Design) /
[UML](https://mengxianbin.github.io/cs-notes/site/Design/UML) /
[Tools](https://mengxianbin.github.io/cs-notes/site/Design/UML/Tools) /
[PlantUML](https://mengxianbin.github.io/cs-notes/site/Design/UML/Tools/PlantUML)


* [Official](https://plantuml.com/)

* [Download](https://plantuml.com/zh/download)

* [GitHub](https://github.com/plantuml/plantuml)

* [Playground](https://www.planttext.com/)

* Example

```puml
@startuml test

Alice -> Bob : message 1
Alice -> Bob : message 2

newpage

Alice -> Bob : message 3
Alice -> Bob : message 4

newpage A title for the\nlast page

Alice -> Bob : message 5
Alice -> Bob : message 6

@enduml
```

---

## Layout

#### [Controlling Class Layout](https://isgb.otago.ac.nz/infosci/mark.george/templates/blob/8e98805c117c7b2e9b9f545c47b50366bb644e5e/plantuml/class-diagram-tips.md)

> 1

```puml
' force class diagram mode
class c1
 
' horizontal placement
c1 -> c2
 
' vertical placement
c2 --> c3
 
' left placement
c1 -left-> c4
```

> 2

```puml
class c4

c1 --> c2

' link will not affect class placement
' c4 --> c2
c4 -[norank]-> c2
```

> 3

```puml
@startuml

class c2

c3 --> c1

' place c2 and c4 closer together
c2 -[hidden]-> c4 
 
' repeated again to place c2 and c4 even closer
c2 -[hidden]-> c4

@enduml
```

> 4

```puml
@startuml

class c2

' place c2 and c4 closer
c2 -[hidden]- c4
       
' draw the actual link between the hidden links
c2 <-[norank]- c4
    
' place c2 and c4 even closer
c2 -[hidden]- c4

@enduml
```

##### General Layout Tips

> 5

```puml
@startuml

class c1 
c1 -> c2
c2 --> c3
c3 -left-> c4
 
' whoops, this flips the diagram on its head
c4 --> c2

@enduml
```

> 6

```puml
@startuml

class c2

' ranking maintained, so layout isn't dramatically affected
' arrow drawn at opposite end to account for reverse order
c2 <-- c4

@enduml
```

> 7

```puml
@startuml

c4 -up-> c2

@enduml
```

##### Layout Direction

> 8

```puml
@startuml

left to right direction
class c1
c1 -> c2
c2 --> c3
c3 -left-> c4
c2 <-- c4

@enduml
```

##### Spacing

> 9

```puml
@startuml

SkinParam {
    'NodeSep 45 ' horizontal spacing
    NodeSep 90 ' horizontal spacing
    'RankSep 45 ' vertical spacing
    RankSep 90 ' vertical spacing
}

class a

a -> b
a --> c

@enduml
```

---
