[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java) /
[SelectorImpl](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl) /
[ensureOpen](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl/ensureOpen)

```java
    private void ensureOpen() {
        if (!isOpen())
            throw new ClosedSelectorException();
    }
```
