[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Handler](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler) /
[timeout](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/timeout) /
[IdleStateEvent](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/timeout/IdleStateEvent) /
[constructor](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/timeout/IdleStateEvent/constructor)

```java
    /**
     * Constructor for sub-classes.
     *
     * @param state the {@link IdleStateEvent} which triggered the event.
     * @param first {@code true} if its the first idle event for the {@link IdleStateEvent}.
     */
    protected IdleStateEvent(IdleState state, boolean first) {
        this.state = ObjectUtil.checkNotNull(state, "state");
        this.first = first;
    }
```
