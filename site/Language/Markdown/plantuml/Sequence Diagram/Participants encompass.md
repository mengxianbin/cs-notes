[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Markdown](https://mengxianbin.github.io/cs-notes/site/Language/Markdown) /
[plantuml](https://mengxianbin.github.io/cs-notes/site/Language/Markdown/plantuml) /
[Sequence Diagram](https://mengxianbin.github.io/cs-notes/site/Language/Markdown/plantuml/Sequence%20Diagram) /
[Participants encompass](https://mengxianbin.github.io/cs-notes/site/Language/Markdown/plantuml/Sequence%20Diagram/Participants%20encompass)

```puml
@startuml

box "Internal Service" #LightBlue
participant Bob
participant Alice
end box
participant Other

Bob -> Alice : hello
Alice -> Other : hello

@enduml
```
