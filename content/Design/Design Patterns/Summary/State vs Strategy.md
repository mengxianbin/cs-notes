[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Design](https://mengxianbin.github.io/cs-notes/content/Design) /
[Design Patterns](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns) /
[Summary](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Summary) /
[State vs Strategy](https://mengxianbin.github.io/cs-notes/content/Design/Design%20Patterns/Summary/State%20vs%20Strategy)

# State vs Strategy

## Applicability

* State
    * a state is associated one or more other states
* Strategy
    * strategies are independent

## Focus

* State
    * state switching
* Strategy
    * independent works

## Context Role Responsibility

* State
    * the context cooperates with states to switch states
* Strategy
    * the context switches strategies ifself

## Visibility

* State
    * states are internal
* Strategy
    * strategy maybe exposed

## Complexity

* State
    * states are coupled
    * state count is usually limited
* Strategy
    * simple
    * extensible
    * readable

---
