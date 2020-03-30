```java
    @Override
    protected void run() {
        int selectCnt = 0;
        for (;;) {
            ...
            strategy = select(curDeadlineNanos);
            ...
            processSelectedKeys();
            ...
        }
    }
```
