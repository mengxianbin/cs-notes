[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java) /
[AbstractInterruptableChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/AbstractInterruptableChannel) /
[close](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/AbstractInterruptableChannel/close)

```java
    /**
     * Closes this channel.
     *
     * <p> If the channel has already been closed then this method returns
     * immediately.  Otherwise it marks the channel as closed and then invokes
     * the {@link #implCloseChannel implCloseChannel} method in order to
     * complete the close operation.  </p>
     *
     * @throws  IOException
     *          If an I/O error occurs
     */
    public final void close() throws IOException {
        synchronized (closeLock) {
            if (closed)
                return;
            closed = true;
            implCloseChannel();
        }
    }
```
