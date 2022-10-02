
```java
    volatile Object result;       // Either the result or boxed AltResult

    final boolean internalComplete(Object r) { // CAS from null to r
        return UNSAFE.compareAndSwapObject(this, RESULT, null, r);
    }
```
