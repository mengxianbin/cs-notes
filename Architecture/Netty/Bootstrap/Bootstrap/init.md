```java
    @Override
    @SuppressWarnings("unchecked")
    void init(Channel channel) {
        ChannelPipeline p = channel.pipeline();
        p.addLast(config.handler());

        setChannelOptions(channel, options0().entrySet().toArray(EMPTY_OPTION_ARRAY), logger);
        setAttributes(channel, attrs0().entrySet().toArray(EMPTY_ATTRIBUTE_ARRAY));
    }
```
