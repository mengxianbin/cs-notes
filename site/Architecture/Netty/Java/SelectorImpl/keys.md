[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java) /
[SelectorImpl](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl) /
[keys](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl/keys)

```java
    @Override
    public final Set<SelectionKey> keys() {
        ensureOpen();
        return publicKeys;
    }

    @Override
    public final Set<SelectionKey> selectedKeys() {
        ensureOpen();
        return publicSelectedKeys;
    }
```
