```java
    /**
     * Invoked by selection operations to process the cancelled-key set
     */
    protected final void processDeregisterQueue() throws IOException {
        assert Thread.holdsLock(this);
        assert Thread.holdsLock(publicSelectedKeys);

        Set<SelectionKey> cks = cancelledKeys();
        synchronized (cks) {
            if (!cks.isEmpty()) {
                Iterator<SelectionKey> i = cks.iterator();
                while (i.hasNext()) {
                    SelectionKeyImpl ski = (SelectionKeyImpl)i.next();
                    i.remove();

                    // remove the key from the selector
                    implDereg(ski);

                    selectedKeys.remove(ski);
                    keys.remove(ski);

                    // remove from channel's key set
                    deregister(ski);

                    SelectableChannel ch = ski.channel();
                    if (!ch.isOpen() && !ch.isRegistered())
                        ((SelChImpl)ch).kill();
                }
            }
        }
    }
```
