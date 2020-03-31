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
