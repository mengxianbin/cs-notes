```java
    private void ensureOpen() {
        if (!isOpen())
            throw new ClosedSelectorException();
    }
```
