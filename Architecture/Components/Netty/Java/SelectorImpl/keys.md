```java
    @Override
    public final Set<SelectionKey> keys() {
        ensureOpen();
        return publicKeys;
    }

    @Override
    public final Set<SelectionKey> selectedKeys() {
        ensureOpen();
        return publicSelectedKeys;
    }
```
