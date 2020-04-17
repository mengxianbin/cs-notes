[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[ByteBuf](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/ByteBuf) /
[methods](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/ByteBuf/methods) /
[ensureWritable](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/ByteBuf/methods/ensureWritable)

## ByteBuf

> io.netty.buffer

```java
public abstract ByteBuf ensureWritable(int minWritableBytes);
```

## AbstractByteBuf

> io.netty.buffer

```java
    @Override
    public ByteBuf ensureWritable(int minWritableBytes) {
        ensureWritable0(checkPositiveOrZero(minWritableBytes, "minWritableBytes"));
        return this;
    }
```

```java
    final void ensureWritable0(int minWritableBytes) {
        final int writerIndex = writerIndex();
        final int targetCapacity = writerIndex + minWritableBytes;
        if (targetCapacity <= capacity()) {
            ensureAccessible();
            return;
        }
        if (checkBounds && targetCapacity > maxCapacity) {
            ensureAccessible();
            throw new IndexOutOfBoundsException(String.format(
                    "writerIndex(%d) + minWritableBytes(%d) exceeds maxCapacity(%d): %s",
                    writerIndex, minWritableBytes, maxCapacity, this));
        }

        // Normalize the target capacity to the power of 2.
        final int fastWritable = maxFastWritableBytes();
        int newCapacity = fastWritable >= minWritableBytes ? writerIndex + fastWritable
                : alloc().calculateNewCapacity(targetCapacity, maxCapacity);

        // Adjust to the new capacity.
        capacity(newCapacity);
    }
```

```java
    public abstract ByteBufAllocator alloc();
```

## ByteBufAllocator

> io.netty.buffer

```java
public interface ByteBufAllocator { ... }
```

## AbstractByteBufAllocator

> package io.netty.buffer;

```java
    @Override
    public int calculateNewCapacity(int minNewCapacity, int maxCapacity) {
        checkPositiveOrZero(minNewCapacity, "minNewCapacity");
        if (minNewCapacity > maxCapacity) {
            throw new IllegalArgumentException(String.format(
                    "minNewCapacity: %d (expected: not greater than maxCapacity(%d)",
                    minNewCapacity, maxCapacity));
        }
        final int threshold = CALCULATE_THRESHOLD; // 4 MiB page

        if (minNewCapacity == threshold) {
            return threshold;
        }

        // If over threshold, do not double but just increase by threshold.
        if (minNewCapacity > threshold) {
            int newCapacity = minNewCapacity / threshold * threshold;
            if (newCapacity > maxCapacity - threshold) {
                newCapacity = maxCapacity;
            } else {
                newCapacity += threshold;
            }
            return newCapacity;
        }

        // Not over threshold. Double up to 4 MiB, starting from 64.
        int newCapacity = 64;
        while (newCapacity < minNewCapacity) {
            newCapacity <<= 1;
        }

        return Math.min(newCapacity, maxCapacity);
    }
```

---
