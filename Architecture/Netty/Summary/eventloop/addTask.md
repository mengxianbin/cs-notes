```java
if (inEventLoop()) {
    ...
} else {
    eventLoop.execute(...);
}
```
