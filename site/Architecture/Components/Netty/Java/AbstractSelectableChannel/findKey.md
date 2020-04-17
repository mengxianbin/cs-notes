[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java) /
[AbstractSelectableChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/AbstractSelectableChannel) /
[findKey](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/AbstractSelectableChannel/findKey)

```java
    private SelectionKey findKey(Selector sel) {
        assert Thread.holdsLock(keyLock);
        if (keys == null)
            return null;
        for (int i = 0; i < keys.length; i++)
            if ((keys[i] != null) && (keys[i].selector() == sel))
                return keys[i];
        return null;

    }
```
