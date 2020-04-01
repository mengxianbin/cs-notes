[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java) /
[SelectorImpl](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl) /
[register](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorImpl/register)

```java
    @Override
    protected final SelectionKey register(AbstractSelectableChannel ch,
                                          int ops,
                                          Object attachment)
    {
        if (!(ch instanceof SelChImpl))
            throw new IllegalSelectorException();
        SelectionKeyImpl k = new SelectionKeyImpl((SelChImpl)ch, this);
        k.attach(attachment);

        // register (if needed) before adding to key set
        implRegister(k);

        // add to the selector's key set, removing it immediately if the selector
        // is closed. The key is not in the channel's key set at this point but
        // it may be observed by a thread iterating over the selector's key set.
        keys.add(k);
        try {
            k.interestOps(ops);
        } catch (ClosedSelectorException e) {
            assert ch.keyFor(this) == null;
            keys.remove(k);
            k.cancel();
            throw e;
        }
        return k;
    }
```
