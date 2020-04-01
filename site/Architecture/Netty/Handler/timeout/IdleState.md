[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Handler](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Handler) /
[timeout](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Handler/timeout) /
[IdleState](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Handler/timeout/IdleState)

```java
package io.netty.handler.timeout;

import io.netty.channel.Channel;


/**
 * An {@link Enum} that represents the idle state of a {@link Channel}.
 */
public enum IdleState {
    /**
     * No data was received for a while.
     */
    READER_IDLE,
    /**
     * No data was sent for a while.
     */
    WRITER_IDLE,
    /**
     * No data was either received or sent for a while.
     */
    ALL_IDLE
}

```
