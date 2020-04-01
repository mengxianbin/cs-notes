[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Basic](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic) /
[IO](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic/IO) /
[Buffer](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic/IO/Buffer)

# java.nio.Buffer

## Fields

```java
    // Cached unsafe-access object
    static final Unsafe UNSAFE = Unsafe.getUnsafe();

    // Invariants: mark <= position <= limit <= capacity
    private int mark = -1;
    private int position = 0;
    private int limit;
    private int capacity;
```

## Methods

```java
    /**
     * Flips this buffer.  The limit is set to the current position and then
     * the position is set to zero.  If the mark is defined then it is
     * discarded.
     *
     * <p> After a sequence of channel-read or <i>put</i> operations, invoke
     * this method to prepare for a sequence of channel-write or relative
     * <i>get</i> operations.  For example:
     *
     * <blockquote><pre>
     * buf.put(magic);    // Prepend header
     * in.read(buf);      // Read data into rest of buffer
     * buf.flip();        // Flip buffer
     * out.write(buf);    // Write header + data to channel</pre></blockquote>
     *
     * <p> This method is often used in conjunction with the {@link
     * java.nio.ByteBuffer#compact compact} method when transferring data from
     * one place to another.  </p>
     *
     * @return  This buffer
     */
    public Buffer flip() {
        limit = position;
        position = 0;
        mark = -1;
        return this;
    }
```

```java
    /**
     * Returns this buffer's position.
     *
     * @return  The position of this buffer
     */
    public final int position() {
        return position;
    }
```

```java
    /**
     * Returns this buffer's capacity.
     *
     * @return  The capacity of this buffer
     */
    public final int capacity() {
        return capacity;
    }
```

```java
    /**
     * Returns this buffer's limit.
     *
     * @return  The limit of this buffer
     */
    public final int limit() {
        return limit;
    }
```

```java
    /**
     * Sets this buffer's mark at its position.
     *
     * @return  This buffer
     */
    public Buffer mark() {
        mark = position;
        return this;
    }
```

```java
    /**
     * Clears this buffer.  The position is set to zero, the limit is set to
     * the capacity, and the mark is discarded.
     *
     * <p> Invoke this method before using a sequence of channel-read or
     * <i>put</i> operations to fill this buffer.  For example:
     *
     * <blockquote><pre>
     * buf.clear();     // Prepare buffer for reading
     * in.read(buf);    // Read data</pre></blockquote>
     *
     * <p> This method does not actually erase the data in the buffer, but it
     * is named as if it did because it will most often be used in situations
     * in which that might as well be the case. </p>
     *
     * @return  This buffer
     */
    public Buffer clear() {
        position = 0;
        limit = capacity;
        mark = -1;
        return this;
    }
```

```java
    /**
     * Returns the number of elements between the current position and the
     * limit.
     *
     * @return  The number of elements remaining in this buffer
     */
    public final int remaining() {
        return limit - position;
    }
```

```java
    /**
     * Resets this buffer's position to the previously-marked position.
     *
     * <p> Invoking this method neither changes nor discards the mark's
     * value. </p>
     *
     * @return  This buffer
     *
     * @throws  InvalidMarkException
     *          If the mark has not been set
     */
    public Buffer reset() {
        int m = mark;
        if (m < 0)
            throw new InvalidMarkException();
        position = m;
        return this;
    }
```

```java
    /**
     * Rewinds this buffer.  The position is set to zero and the mark is
     * discarded.
     *
     * <p> Invoke this method before a sequence of channel-write or <i>get</i>
     * operations, assuming that the limit has already been set
     * appropriately.  For example:
     *
     * <blockquote><pre>
     * out.write(buf);    // Write remaining data
     * buf.rewind();      // Rewind buffer
     * buf.get(array);    // Copy data into array</pre></blockquote>
     *
     * @return  This buffer
     */
    public Buffer rewind() {
        position = 0;
        mark = -1;
        return this;
    }
```

## See Also

* java.nio.ByteBuffer

---
