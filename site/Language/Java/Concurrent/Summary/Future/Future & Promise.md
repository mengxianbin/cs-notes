[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent) /
[Summary](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Summary) /
[Future](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Summary/Future) /
[Future & Promise](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Summary/Future/Future%20&%20Promise)

# Future & Promise

> decouple computing from results
> compose tasks flexibly
> parallelize

- future: the result
- promise: computing

---

- future <- functional programming
- promise <- callback hell

---

## JDK Future

> Since 1.5

```java
package java.util.concurrent;
/**
 * A Future represents the result of an asynchronous computation....
 */
public interface Future<V> {
    boolean cancel(boolean mayInterruptIfRunning);
    boolean isCancelled();
    boolean isDone();
    V get() throws InterruptedException, ExecutionException;
    V get(long timeout, TimeUnit unit) throws InterruptedException, ExecutionException, TimeoutException;
}
```

---

## JavaScript Promise

> ES6

- **prototype.then**
- **prototype.catch**
- **prototype.finally**
- all
- any
- race
- ...

## Java Third Callback Implements

#### Netty

```java
package io.netty.util.concurrent;

/**
 * The result of an asynchronous operation.
 */
public interface Future<V> extends java.util.concurrent.Future<V> {
    ...
```

- `boolean isSuccess();`
- `boolean isCancellable();`
- `Throwable cause();`
- **addListener**
- **addListeners**
- removeListener
- removeListeners
- ...

#### Guava

```java
package com.google.common.util.concurrent;
/**
 * A {@link Future} that accepts completion listeners....
 */
public interface ListenableFuture<V> extends Future<V> {
    void addListener(Runnable listener, Executor executor);
}
```

## Java Third Promise Implements

#### Netty

```java
package io.netty.util.concurrent;

/**
 * Special {@link Future} which is writable.
 */
public interface Promise<V> extends Future<V> {
    Promise<V> setSuccess(V result);
    boolean trySuccess(V result);
    Promise<V> setFailure(Throwable cause);
    boolean tryFailure(Throwable cause);
    ...
```

#### Guava

```java
package com.google.common.util.concurrent;

/**
 * A {@link ListenableFuture} whose result can be set by a {@link #set(Object)}, {@link
 * #setException(Throwable)} or {@link #setFuture(ListenableFuture)} call. It can also, like any
 * other {@code Future}, be {@linkplain #cancel cancelled}....
 */
public final class SettableFuture<V> extends AbstractFuture.TrustedFuture<V> {

  public boolean set(@Nullable V value) {
    return super.set(value);
  }
  public boolean setException(Throwable throwable) {
    return super.setException(throwable);
  }
  public boolean setFuture(ListenableFuture<? extends V> future) {
    return super.setFuture(future);
  }

  ...
}
```

## JDK CompletableFuture

---
