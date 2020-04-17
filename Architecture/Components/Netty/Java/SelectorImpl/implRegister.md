```java
    /**
     * Register the key in the selector.
     *
     * The default implementation checks if the selector is open. It should
     * be overridden by selector implementations as needed.
     */
    protected void implRegister(SelectionKeyImpl ski) {
        ensureOpen();
    }
```
