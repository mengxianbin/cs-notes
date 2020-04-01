[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java) /
[SelectorImpl](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl) /
[implRegister](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl/implRegister)

```java
    /**
     * Register the key in the selector.
     *
     * The default implementation checks if the selector is open. It should
     * be overridden by selector implementations as needed.
     */
    protected void implRegister(SelectionKeyImpl ski) {
        ensureOpen();
    }
```
