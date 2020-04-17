```java
    private static final AtomicIntegerFieldUpdater<DefaultChannelConfig> AUTOREAD_UPDATER =
            AtomicIntegerFieldUpdater.newUpdater(DefaultChannelConfig.class, "autoRead");
    
    @SuppressWarnings("FieldMayBeFinal")
    private volatile int autoRead = 1;
```
