[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java) /
[SelectionKey](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/SelectionKey) /
[selector](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/SelectionKey/selector)

```java
    /**
     * Returns the selector for which this key was created.  This method will
     * continue to return the selector even after the key is cancelled.
     *
     * @return  This key's selector
     */
    public abstract Selector selector();
```
