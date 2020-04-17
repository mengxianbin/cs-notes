```java
    /**
     * Retrieves this key's interest set.
     *
     * <p> It is guaranteed that the returned set will only contain operation
     * bits that are valid for this key's channel. </p>
     *
     * @return  This key's interest set
     *
     * @throws  CancelledKeyException
     *          If this key has been cancelled
     */
    public abstract int interestOps();
```

---

```java
    /**
     * Sets this key's interest set to the given value.
     *
     * <p> This method may be invoked at any time.  If this method is invoked
     * while a selection operation is in progress then it has no effect upon
     * that operation; the change to the key's interest set will be seen by the
     * next selection operation.
     *
     * @param  ops  The new interest set
     *
     * @return  This selection key
     *
     * @throws  IllegalArgumentException
     *          If a bit in the set does not correspond to an operation that
     *          is supported by this key's channel, that is, if
     *          {@code (ops & ~channel().validOps()) != 0}
     *
     * @throws  CancelledKeyException
     *          If this key has been cancelled
     */
    public abstract SelectionKey interestOps(int ops);
```

* Usages
    * register - AbstractSelectableChannel
    * register - SelectorImpl
    * processSelectedKey - NioEventLoop
    * doWrite - AbstractNioMessageChannel
    * doBeginRead - AbstractNioChannel
    * removeReadOp - AbstractNioChannel.AbstractUnsafe
        * read - AbstractNioMessageChannel.NioMessageUnsafe
        * read - AbstractNioByteChannel.NioByteUnsafe

---
