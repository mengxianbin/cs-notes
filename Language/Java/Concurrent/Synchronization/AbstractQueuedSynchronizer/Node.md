# Fields

```java
volatile Node prev;
volatile Node next;

/**
    * The thread that enqueued this node.  Initialized on
    * construction and nulled out after use.
    */
volatile Thread thread;
```
