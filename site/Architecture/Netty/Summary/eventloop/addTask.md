[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Summary](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Summary) /
[eventloop](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Summary/eventloop) /
[addTask](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Summary/eventloop/addTask)

```java
if (inEventLoop()) {
    ...
} else {
    eventLoop.execute(...);
}
```
