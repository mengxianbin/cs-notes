# ReentrantLock

## Constructor

```java
/**
    * Creates an instance of {@code ReentrantLock}.
    * This is equivalent to using {@code ReentrantLock(false)}.
    */
public ReentrantLock() {
    sync = new NonfairSync();
}

/**
    * Creates an instance of {@code ReentrantLock} with the
    * given fairness policy.
    *
    * @param fair {@code true} if this lock should use a fair ordering policy
    */
public ReentrantLock(boolean fair) {
    sync = fair ? new FairSync() : new NonfairSync();
}
```

## Methods

```java
    public void lock() {
        sync.lock();
    }
```

```java
    public void lockInterruptibly() throws InterruptedException {
        sync.acquireInterruptibly(1);
    }
```

```java
    public boolean tryLock() {
        return sync.nonfairTryAcquire(1);
    }
```

```java
    public boolean tryLock(long timeout, TimeUnit unit)
            throws InterruptedException {
        return sync.tryAcquireNanos(1, unit.toNanos(timeout));
    }
```

```java
    public void unlock() {
        sync.release(1);
    }
```
