[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent) /
[Code](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code) /
[Future](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Future) /
[CompletableFuture](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Future/CompletableFuture) /
[result](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Code/Future/CompletableFuture/result)


```java
    volatile Object result;       // Either the result or boxed AltResult

    final boolean internalComplete(Object r) { // CAS from null to r
        return UNSAFE.compareAndSwapObject(this, RESULT, null, r);
    }
```
