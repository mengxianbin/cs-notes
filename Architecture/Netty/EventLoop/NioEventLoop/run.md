```java
    @Override
    protected void run() {
        ...
        for (;;) {
            int strategy;
            ...
            strategy = selectStrategy.calculateStrategy(selectNowSupplier, hasTasks());

            switch (strategy) {
            case SelectStrategy.CONTINUE:
                continue;

            case SelectStrategy.BUSY_WAIT:
                // fall-through to SELECT since the busy-wait is not supported with NIO

            case SelectStrategy.SELECT:
                ...
                try {
                    if (!hasTasks()) {
                        strategy = select(curDeadlineNanos);
                    }
                } finally {
                    ...
                }
                // fall through
            default:
            }
            ...
            strategy = select(curDeadlineNanos);
            ...
            if (ioRatio == 100) {
                try {
                    if (strategy > 0) {
                        processSelectedKeys();
                    }
                } finally {
                    // Ensure we always run tasks.
                    ranTasks = runAllTasks();
                }
            } else if (strategy > 0) {
                final long ioStartTime = System.nanoTime();
                try {
                    processSelectedKeys();
                } finally {
                    // Ensure we always run tasks.
                    final long ioTime = System.nanoTime() - ioStartTime;
                    ranTasks = runAllTasks(ioTime * (100 - ioRatio) / ioRatio);
                }
            } else {
                ranTasks = runAllTasks(0); // This will run the minimum number of tasks
            }
            ...
    }
```
