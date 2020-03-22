[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[JVM](https://mengxianbin.github.io/cs-notes/content/Language/Java/JVM) /
[JMM](https://mengxianbin.github.io/cs-notes/content/Language/Java/JVM/JMM) /
[Inter-memory Interaction](https://mengxianbin.github.io/cs-notes/content/Language/Java/JVM/JMM/Inter-memory%20Interaction)

- System modeled by constraints between actions
    - Use/assign actions correspond to thread computations
    - load/store actions correspond to thread fill/flush actions
    - read/write actions are main memory actions

---

## Operation type

* Lock: A variable that acts on main memory, which identifies a variable as a thread-exclusive state.
* Unlock: Acts on the main memory variable, which releases a locked variable, and the released variable can be locked by other threads.

* ...

---
