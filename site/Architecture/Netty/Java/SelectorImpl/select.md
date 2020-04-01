[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java) /
[SelectorImpl](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl) /
[select](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl/select)

```java
    @Override
    public final int select() throws IOException {
        return lockAndDoSelect(null, -1);
    }
```

```java
    private int lockAndDoSelect(Consumer<SelectionKey> action, long timeout)
        throws IOException
    {
        synchronized (this) {
            ensureOpen();
            if (inSelect)
                throw new IllegalStateException("select in progress");
            inSelect = true;
            try {
                synchronized (publicSelectedKeys) {
                    return doSelect(action, timeout);
                }
            } finally {
                inSelect = false;
            }
        }
    }
```

```java
    /**
     * Selects the keys for channels that are ready for I/O operations.
     *
     * @param action  the action to perform, can be null
     * @param timeout timeout in milliseconds to wait, 0 to not wait, -1 to
     *                wait indefinitely
     */
    protected abstract int doSelect(Consumer<SelectionKey> action, long timeout)
        throws IOException;
```
