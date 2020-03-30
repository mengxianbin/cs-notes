```java
    @Override
    public ChannelFuture write(Object msg) {
        return pipeline.write(msg);
    }
```
