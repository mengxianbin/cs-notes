[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java) /
[SelectorImpl](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/SelectorImpl) /
[implRegister](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/SelectorImpl/implRegister)

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
