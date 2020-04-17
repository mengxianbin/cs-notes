[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Channel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel) /
[NioServerSocketChannel](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/NioServerSocketChannel) /
[isActive](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Channel/NioServerSocketChannel/isActive)

```java
    @Override
    public boolean isActive() {
        // As java.nio.ServerSocketChannel.isBound() will continue to return true even after the channel was closed
        // we will also need to check if it is open.
        return isOpen() && javaChannel().socket().isBound();
    }
```
