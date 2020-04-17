[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java) /
[SelectionKey](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/SelectionKey) /
[attach](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Java/SelectionKey/attach)

```java
    private volatile Object attachment;

    private static final AtomicReferenceFieldUpdater<SelectionKey,Object>
        attachmentUpdater = AtomicReferenceFieldUpdater.newUpdater(
            SelectionKey.class, Object.class, "attachment"
        );
```

```java
    /**
     * Attaches the given object to this key.
     *
     * <p> An attached object may later be retrieved via the {@link #attachment()
     * attachment} method.  Only one object may be attached at a time; invoking
     * this method causes any previous attachment to be discarded.  The current
     * attachment may be discarded by attaching {@code null}.  </p>
     *
     * @param  ob
     *         The object to be attached; may be {@code null}
     *
     * @return  The previously-attached object, if any,
     *          otherwise {@code null}
     */
    public final Object attach(Object ob) {
        return attachmentUpdater.getAndSet(this, ob);
    }
```

```java
    /**
     * Retrieves the current attachment.
     *
     * @return  The object currently attached to this key,
     *          or {@code null} if there is no attachment
     */
    public final Object attachment() {
        return attachment;
    }
```
