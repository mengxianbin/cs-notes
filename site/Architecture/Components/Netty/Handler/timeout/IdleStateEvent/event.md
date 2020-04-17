[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Handler](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler) /
[timeout](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/timeout) /
[IdleStateEvent](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/timeout/IdleStateEvent) /
[event](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Handler/timeout/IdleStateEvent/event)

```java
    public static final IdleStateEvent FIRST_READER_IDLE_STATE_EVENT =
            new DefaultIdleStateEvent(IdleState.READER_IDLE, true);
    public static final IdleStateEvent READER_IDLE_STATE_EVENT =
            new DefaultIdleStateEvent(IdleState.READER_IDLE, false);
    public static final IdleStateEvent FIRST_WRITER_IDLE_STATE_EVENT =
            new DefaultIdleStateEvent(IdleState.WRITER_IDLE, true);
    public static final IdleStateEvent WRITER_IDLE_STATE_EVENT =
            new DefaultIdleStateEvent(IdleState.WRITER_IDLE, false);
    public static final IdleStateEvent FIRST_ALL_IDLE_STATE_EVENT =
            new DefaultIdleStateEvent(IdleState.ALL_IDLE, true);
    public static final IdleStateEvent ALL_IDLE_STATE_EVENT =
            new DefaultIdleStateEvent(IdleState.ALL_IDLE, false);
```
