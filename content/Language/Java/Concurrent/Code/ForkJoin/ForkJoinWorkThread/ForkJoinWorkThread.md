[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent) /
[Code](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code) /
[ForkJoin](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code/ForkJoin) /
[ForkJoinWorkThread](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code/ForkJoin/ForkJoinWorkThread) /
[ForkJoinWorkThread](https://mengxianbin.github.io/cs-notes/content/Language/Java/Concurrent/Code/ForkJoin/ForkJoinWorkThread/ForkJoinWorkThread)

```java
/**
 * A thread managed by a {@link ForkJoinPool}, which executes
 * {@link ForkJoinTask}s.
 * This class is subclassable solely for the sake of adding
 * functionality -- there are no overridable methods dealing with
 * scheduling or execution.  However, you can override initialization
 * and termination methods surrounding the main task processing loop.
 * If you do create such a subclass, you will also need to supply a
 * custom {@link ForkJoinPool.ForkJoinWorkerThreadFactory} to
 * {@linkplain ForkJoinPool#ForkJoinPool use it} in a {@code ForkJoinPool}.
 *
 * @since 1.7
 * @author Doug Lea
 */
public class ForkJoinWorkerThread extends Thread {
```