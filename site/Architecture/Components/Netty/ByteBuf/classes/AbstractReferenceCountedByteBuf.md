[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[ByteBuf](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/ByteBuf) /
[classes](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/ByteBuf/classes) /
[AbstractReferenceCountedByteBuf](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/ByteBuf/classes/AbstractReferenceCountedByteBuf)

```java
    @Override
    public ByteBuf retain() {
        return updater.retain(this);
    }

    @Override
    public ByteBuf retain(int increment) {
        return updater.retain(this, increment);
    }

    @Override
    public boolean release() {
        return handleRelease(updater.release(this));
    }

    @Override
    public boolean release(int decrement) {
        return handleRelease(updater.release(this, decrement));
    }
```
