```java
    /**
     * Constructor for sub-classes.
     *
     * @param state the {@link IdleStateEvent} which triggered the event.
     * @param first {@code true} if its the first idle event for the {@link IdleStateEvent}.
     */
    protected IdleStateEvent(IdleState state, boolean first) {
        this.state = ObjectUtil.checkNotNull(state, "state");
        this.first = first;
    }
```
