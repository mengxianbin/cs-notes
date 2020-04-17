```java
        @Override
        public final void write(Object msg, ChannelPromise promise) {
            ...
            ChannelOutboundBuffer outboundBuffer = this.outboundBuffer;
            ...
            outboundBuffer.addMessage(msg, size, promise);
        }
```
