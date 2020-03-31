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
