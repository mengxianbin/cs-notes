[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Summary](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Summary) /
[pipeline](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Summary/pipeline)

* <https://netty.io/4.1/api/io/netty/channel/ChannelPipeline.html>

```
A list of ChannelHandlers which handles or intercepts inbound events and outbound operations of a Channel.
ChannelPipeline implements an advanced form of the Intercepting Filter pattern to give a user full control
 over how an event is handled and how the ChannelHandlers in a pipeline interact with each other.
```

* inbound: event <- socket.read
* outbound: operation -> socket.write

---
