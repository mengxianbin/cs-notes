```java
    protected SelectorImpl(SelectorProvider sp) {
        super(sp);
        keys = ConcurrentHashMap.newKeySet();
        selectedKeys = new HashSet<>();
        publicKeys = Collections.unmodifiableSet(keys);
        publicSelectedKeys = Util.ungrowableSet(selectedKeys);
    }
```
