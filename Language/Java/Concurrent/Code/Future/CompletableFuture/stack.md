
```java
    /** Pushes the given completion (if it exists) unless done. */
    final void push(UniCompletion<?,?> c) {
        if (c != null) {
            while (result == null && !tryPushStack(c))
                lazySetNext(c, null); // clear on failure
        }
    }
```

```java
    volatile Completion stack;    // Top of Treiber stack of dependent actions

    final boolean casStack(Completion cmp, Completion val) {
        return UNSAFE.compareAndSwapObject(this, STACK, cmp, val);
    }

    /** Returns true if successfully pushed c onto stack. */
    final boolean tryPushStack(Completion c) {
        Completion h = stack;
        lazySetNext(c, h);
        return UNSAFE.compareAndSwapObject(this, STACK, h, c);
    }

    /** Unconditionally pushes c onto stack, retrying if necessary. */
    final void pushStack(Completion c) {
        do {} while (!tryPushStack(c));
    }
```
