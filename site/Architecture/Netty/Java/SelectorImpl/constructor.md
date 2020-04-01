[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java) /
[SelectorImpl](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl) /
[constructor](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl/constructor)

```java
    protected SelectorImpl(SelectorProvider sp) {
        super(sp);
        keys = ConcurrentHashMap.newKeySet();
        selectedKeys = new HashSet<>();
        publicKeys = Collections.unmodifiableSet(keys);
        publicSelectedKeys = Util.ungrowableSet(selectedKeys);
    }
```
