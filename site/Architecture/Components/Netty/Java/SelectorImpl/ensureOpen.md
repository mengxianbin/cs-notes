[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java) /
[SelectorImpl](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/SelectorImpl) /
[ensureOpen](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/SelectorImpl/ensureOpen)

```java
    private void ensureOpen() {
        if (!isOpen())
            throw new ClosedSelectorException();
    }
```
