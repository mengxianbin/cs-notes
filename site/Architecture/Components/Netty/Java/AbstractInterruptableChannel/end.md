[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java) /
[AbstractInterruptableChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/AbstractInterruptableChannel) /
[end](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/AbstractInterruptableChannel/end)

```java
    /**
     * Marks the end of an I/O operation that might block indefinitely.
     *
     * <p> This method should be invoked in tandem with the {@link #begin
     * begin} method, using a {@code try}&nbsp;...&nbsp;{@code finally} block
     * as shown <a href="#be">above</a>, in order to implement asynchronous
     * closing and interruption for this channel.  </p>
     *
     * @param  completed
     *         {@code true} if, and only if, the I/O operation completed
     *         successfully, that is, had some effect that would be visible to
     *         the operation's invoker
     *
     * @throws  AsynchronousCloseException
     *          If the channel was asynchronously closed
     *
     * @throws  ClosedByInterruptException
     *          If the thread blocked in the I/O operation was interrupted
     */
    protected final void end(boolean completed)
        throws AsynchronousCloseException
    {
        blockedOn(null);
        Thread interrupted = this.interrupted;
        if (interrupted != null && interrupted == Thread.currentThread()) {
            this.interrupted = null;
            throw new ClosedByInterruptException();
        }
        if (!completed && closed)
            throw new AsynchronousCloseException();
    }
```
