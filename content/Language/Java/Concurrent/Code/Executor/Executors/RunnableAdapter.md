[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent) /
[Code](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code) /
[Executor](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code/Executor) /
[Executors](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code/Executor/Executors) /
[RunnableAdapter](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code/Executor/Executors/RunnableAdapter)

```java
    /**
     * A callable that runs given task and returns given result
     */
    static final class RunnableAdapter<T> implements Callable<T> {
        final Runnable task;
        final T result;
        RunnableAdapter(Runnable task, T result) {
            this.task = task;
            this.result = result;
        }
        public T call() {
            task.run();
            return result;
        }
    }
```