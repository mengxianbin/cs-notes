[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty) /
[Bootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap) /
[AbstractBootstrap](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap/AbstractBootstrap) /
[localAddress](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Netty/Bootstrap/AbstractBootstrap/localAddress)

```java
    /**
     * @see #localAddress(SocketAddress)
     */
    public B localAddress(String inetHost, int inetPort) {
        return localAddress(SocketUtils.socketAddress(inetHost, inetPort));
    }
```

```java
    /**
     * The {@link SocketAddress} which is used to bind the local "end" to.
     */
    public B localAddress(SocketAddress localAddress) {
        this.localAddress = localAddress;
        return self();
    }
```