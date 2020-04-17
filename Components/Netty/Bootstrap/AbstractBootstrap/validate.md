```java
    /**
     * Validate all the parameters. Sub-classes may override this, but should
     * call the super method in that case.
     */
    public B validate() {
        if (group == null) {
            throw new IllegalStateException("group not set");
        }
        if (channelFactory == null) {
            throw new IllegalStateException("channel or channelFactory not set");
        }
        return self();
    }
```
