[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[EventLoop](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop) /
[SingleThreadEventExecutor](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop/SingleThreadEventExecutor) /
[addTask](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/EventLoop/SingleThreadEventExecutor/addTask)

```java
    private void execute(Runnable task, boolean immediate) {
        boolean inEventLoop = inEventLoop();
        addTask(task);
        if (!inEventLoop) {
            startThread();
            ...
```

```java
    /**
     * Add a task to the task queue, or throws a {@link RejectedExecutionException} if this instance was shutdown
     * before.
     */
    protected void addTask(Runnable task) {
        ObjectUtil.checkNotNull(task, "task");
        if (!offerTask(task)) {
            reject(task);
        }
    }
```

```java
    final boolean offerTask(Runnable task) {
        if (isShutdown()) {
            reject();
        }
        return taskQueue.offer(task);
    }
```
