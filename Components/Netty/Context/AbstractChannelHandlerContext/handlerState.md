```java
private volatile int handlerState = INIT;

private static final AtomicIntegerFieldUpdater<AbstractChannelHandlerContext> HANDLER_STATE_UPDATER =
            AtomicIntegerFieldUpdater.newUpdater(AbstractChannelHandlerContext.class, "handlerState");
```

```java
    final boolean setAddComplete() {
        for (;;) {
            int oldState = handlerState;
            if (oldState == REMOVE_COMPLETE) {
                return false;
            }
            // Ensure we never update when the handlerState is REMOVE_COMPLETE already.
            // oldState is usually ADD_PENDING but can also be REMOVE_COMPLETE when an EventExecutor is used that is not
            // exposing ordering guarantees.
            if (HANDLER_STATE_UPDATER.compareAndSet(this, oldState, ADD_COMPLETE)) {
                return true;
            }
        }
    }
```

```java
    final void setAddPending() {
        boolean updated = HANDLER_STATE_UPDATER.compareAndSet(this, INIT, ADD_PENDING);
        assert updated; // This should always be true as it MUST be called before setAddComplete() or setRemoved().
    }
```

```java
    final void setRemoved() {
        handlerState = REMOVE_COMPLETE;
    }
```
