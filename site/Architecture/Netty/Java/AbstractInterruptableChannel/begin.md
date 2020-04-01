[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java) /
[AbstractInterruptableChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/AbstractInterruptableChannel) /
[begin](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/AbstractInterruptableChannel/begin)

```java
    /**
     * Marks the beginning of an I/O operation that might block indefinitely.
     *
     * <p> This method should be invoked in tandem with the {@link #end end}
     * method, using a {@code try}&nbsp;...&nbsp;{@code finally} block as
     * shown <a href="#be">above</a>, in order to implement asynchronous
     * closing and interruption for this channel.  </p>
     */
    protected final void begin() {
        if (interruptor == null) {
            interruptor = new Interruptible() {
                    public void interrupt(Thread target) {
                        synchronized (closeLock) {
                            if (closed)
                                return;
                            closed = true;
                            interrupted = target;
                            try {
                                AbstractInterruptibleChannel.this.implCloseChannel();
                            } catch (IOException x) { }
                        }
                    }};
        }
        blockedOn(interruptor);
        Thread me = Thread.currentThread();
        if (me.isInterrupted())
            interruptor.interrupt(me);
    }
```
