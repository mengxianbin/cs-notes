```java
        @Override
        public final void flush() {
            assertEventLoop();

            ChannelOutboundBuffer outboundBuffer = this.outboundBuffer;
            if (outboundBuffer == null) {
                return;
            }

            outboundBuffer.addFlush();
            flush0();
        }
```

```java
        @SuppressWarnings("deprecation")
        protected void flush0() {
            ...
            final ChannelOutboundBuffer outboundBuffer = this.outboundBuffer;
            ...
            doWrite(outboundBuffer);
            ...
        }
```
