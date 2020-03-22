```java
/**
 * An {@link ExecutorService} for running {@link ForkJoinTask}s.
 * A {@code ForkJoinPool} provides the entry point for submissions
 * from non-{@code ForkJoinTask} clients, as well as management and
 * monitoring operations.
 *
 * <p>A {@code ForkJoinPool} differs from other kinds of {@link
 * ExecutorService} mainly by virtue of employing
 * <em>work-stealing</em>: all threads in the pool attempt to find and
 * execute tasks submitted to the pool and/or created by other active
 * tasks (eventually blocking waiting for work if none exist). This
 * enables efficient processing when most tasks spawn other subtasks
 * (as do most {@code ForkJoinTask}s), as well as when many small
 * tasks are submitted to the pool from external clients.  Especially
 * when setting <em>asyncMode</em> to true in constructors, {@code
 * ForkJoinPool}s may also be appropriate for use with event-style
 * tasks that are never joined.
 *
 * <p>A static {@link #commonPool()} is available and appropriate for
 * most applications. The common pool is used by any ForkJoinTask that
 * is not explicitly submitted to a specified pool. Using the common
 * pool normally reduces resource usage (its threads are slowly
 * reclaimed during periods of non-use, and reinstated upon subsequent
 * use).
 *
 * <p>For applications that require separate or custom pools, a {@code
 * ForkJoinPool} may be constructed with a given target parallelism
 * level; by default, equal to the number of available processors.
 * The pool attempts to maintain enough active (or available) threads
 * by dynamically adding, suspending, or resuming internal worker
 * threads, even if some tasks are stalled waiting to join others.
 * However, no such adjustments are guaranteed in the face of blocked
 * I/O or other unmanaged synchronization. The nested {@link
 * ManagedBlocker} interface enables extension of the kinds of
 * synchronization accommodated.
 *
 * <p>In addition to execution and lifecycle control methods, this
 * class provides status check methods (for example
 * {@link #getStealCount}) that are intended to aid in developing,
 * tuning, and monitoring fork/join applications. Also, method
 * {@link #toString} returns indications of pool state in a
 * convenient form for informal monitoring.
 *
 * <p>As is the case with other ExecutorServices, there are three
 * main task execution methods summarized in the following table.
 * These are designed to be used primarily by clients not already
 * engaged in fork/join computations in the current pool.  The main
 * forms of these methods accept instances of {@code ForkJoinTask},
 * but overloaded forms also allow mixed execution of plain {@code
 * Runnable}- or {@code Callable}- based activities as well.  However,
 * tasks that are already executing in a pool should normally instead
 * use the within-computation forms listed in the table unless using
 * async event-style tasks that are not usually joined, in which case
 * there is little difference among choice of methods.
 *
 * <table BORDER CELLPADDING=3 CELLSPACING=1>
 * <caption>Summary of task execution methods</caption>
 *  <tr>
 *    <td></td>
 *    <td ALIGN=CENTER> <b>Call from non-fork/join clients</b></td>
 *    <td ALIGN=CENTER> <b>Call from within fork/join computations</b></td>
 *  </tr>
 *  <tr>
 *    <td> <b>Arrange async execution</b></td>
 *    <td> {@link #execute(ForkJoinTask)}</td>
 *    <td> {@link ForkJoinTask#fork}</td>
 *  </tr>
 *  <tr>
 *    <td> <b>Await and obtain result</b></td>
 *    <td> {@link #invoke(ForkJoinTask)}</td>
 *    <td> {@link ForkJoinTask#invoke}</td>
 *  </tr>
 *  <tr>
 *    <td> <b>Arrange exec and obtain Future</b></td>
 *    <td> {@link #submit(ForkJoinTask)}</td>
 *    <td> {@link ForkJoinTask#fork} (ForkJoinTasks <em>are</em> Futures)</td>
 *  </tr>
 * </table>
 *
 * <p>The common pool is by default constructed with default
 * parameters, but these may be controlled by setting three
 * {@linkplain System#getProperty system properties}:
 * <ul>
 * <li>{@code java.util.concurrent.ForkJoinPool.common.parallelism}
 * - the parallelism level, a non-negative integer
 * <li>{@code java.util.concurrent.ForkJoinPool.common.threadFactory}
 * - the class name of a {@link ForkJoinWorkerThreadFactory}
 * <li>{@code java.util.concurrent.ForkJoinPool.common.exceptionHandler}
 * - the class name of a {@link UncaughtExceptionHandler}
 * </ul>
 * If a {@link SecurityManager} is present and no factory is
 * specified, then the default pool uses a factory supplying
 * threads that have no {@link Permissions} enabled.
 * The system class loader is used to load these classes.
 * Upon any error in establishing these settings, default parameters
 * are used. It is possible to disable or limit the use of threads in
 * the common pool by setting the parallelism property to zero, and/or
 * using a factory that may return {@code null}. However doing so may
 * cause unjoined tasks to never be executed.
 *
 * <p><b>Implementation notes</b>: This implementation restricts the
 * maximum number of running threads to 32767. Attempts to create
 * pools with greater than the maximum number result in
 * {@code IllegalArgumentException}.
 *
 * <p>This implementation rejects submitted tasks (that is, by throwing
 * {@link RejectedExecutionException}) only when the pool is shut down
 * or internal resources have been exhausted.
 *
 * @since 1.7
 * @author Doug Lea
 */
@sun.misc.Contended
public class ForkJoinPool extends AbstractExecutorService { ... }
```