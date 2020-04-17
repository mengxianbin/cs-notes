```java
    @Override
    public void execute(Runnable command) {
        next().execute(command);
    }
```
