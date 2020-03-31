```java
private static final FastThreadLocal<EventExecutor> mappings = new FastThreadLocal<EventExecutor>();
```

```java
    /**
     * Set the current {@link EventExecutor} that is used by the {@link Thread}.
     */
    private static void setCurrentEventExecutor(EventExecutor executor) {
        mappings.set(executor);
    }
```
