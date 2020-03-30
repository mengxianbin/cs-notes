```java
    private void processSelectedKeys() {
        if (selectedKeys != null) {
            processSelectedKeysOptimized();
        } else {
            processSelectedKeysPlain(selector.selectedKeys());
        }
    }
```

```java
    private void processSelectedKeysOptimized() {
        for (int i = 0; i < selectedKeys.size; ++i) {
            final SelectionKey k = selectedKeys.keys[i];
            // null out entry in the array to allow to have it GC'ed once the Channel close
            // See https://github.com/netty/netty/issues/2363
            selectedKeys.keys[i] = null;

            final Object a = k.attachment();

            if (a instanceof AbstractNioChannel) {
                processSelectedKey(k, (AbstractNioChannel) a);
            } else {
                @SuppressWarnings("unchecked")
                NioTask<SelectableChannel> task = (NioTask<SelectableChannel>) a;
                processSelectedKey(k, task);
            }
            ...
        }
    }
```

```java
    private void processSelectedKeysPlain(Set<SelectionKey> selectedKeys) {
        // check if the set is empty and if so just return to not create garbage by
        // creating a new Iterator every time even if there is nothing to process.
        // See https://github.com/netty/netty/issues/597
        if (selectedKeys.isEmpty()) {
            return;
        }

        Iterator<SelectionKey> i = selectedKeys.iterator();
        for (;;) {
            final SelectionKey k = i.next();
            final Object a = k.attachment();
            i.remove();

            if (a instanceof AbstractNioChannel) {
                processSelectedKey(k, (AbstractNioChannel) a);
            } else {
                @SuppressWarnings("unchecked")
                NioTask<SelectableChannel> task = (NioTask<SelectableChannel>) a;
                processSelectedKey(k, task);
            }
            ...
        }
    }
```

```java
private void processSelectedKey(SelectionKey k, AbstractNioChannel ch) {
    ...
    int readyOps = k.readyOps();
    // We first need to call finishConnect() before try to trigger a read(...) or write(...) as otherwise
    // the NIO JDK channel implementation may throw a NotYetConnectedException.
    if ((readyOps & SelectionKey.OP_CONNECT) != 0) {
        // remove OP_CONNECT as otherwise Selector.select(..) will always return without blocking
        // See https://github.com/netty/netty/issues/924
        int ops = k.interestOps();
        ops &= ~SelectionKey.OP_CONNECT;
        k.interestOps(ops);

        unsafe.finishConnect();
    }

    // Process OP_WRITE first as we may be able to write some queued buffers and so free memory.
    if ((readyOps & SelectionKey.OP_WRITE) != 0) {
        // Call forceFlush which will also take care of clear the OP_WRITE once there is nothing left to write
        ch.unsafe().forceFlush();
    }

    // Also check for readOps of 0 to workaround possible JDK bug which may otherwise lead
    // to a spin loop
    if ((readyOps & (SelectionKey.OP_READ | SelectionKey.OP_ACCEPT)) != 0 || readyOps == 0) {
        unsafe.read();
    }
    ...
}
```

```java
    private static void processSelectedKey(SelectionKey k, NioTask<SelectableChannel> task) {
        ...
        task.channelReady(k.channel(), k);
        ...
    }
```


