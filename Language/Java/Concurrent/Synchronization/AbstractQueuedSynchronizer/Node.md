# Fields

```java
/** Marker to indicate a node is waiting in shared mode */
static final Node SHARED = new Node();
/** Marker to indicate a node is waiting in exclusive mode */
static final Node EXCLUSIVE = null;

/** waitStatus value to indicate thread has cancelled */
static final int CANCELLED =  1;
/** waitStatus value to indicate successor's thread needs unparking */
static final int SIGNAL    = -1;
/** waitStatus value to indicate thread is waiting on condition */
static final int CONDITION = -2;
/**
 * waitStatus value to indicate the next acquireShared should
 * unconditionally propagate
 */
static final int PROPAGATE = -3;

/**
 * Status field, taking on only the values:
 *   SIGNAL:     The successor of this node is (or will soon be)
 *               blocked (via park), so the current node must
 *               unpark its successor when it releases or
 *               cancels. To avoid races, acquire methods must
 *               first indicate they need a signal,
 *               then retry the atomic acquire, and then,
 *               on failure, block.
 *   CANCELLED:  This node is cancelled due to timeout or interrupt.
 *               Nodes never leave this state. In particular,
 *               a thread with cancelled node never again blocks.
 *   CONDITION:  This node is currently on a condition queue.
 *               It will not be used as a sync queue node
 *               until transferred, at which time the status
 *               will be set to 0. (Use of this value here has
 *               nothing to do with the other uses of the
 *               field, but simplifies mechanics.)
 *   PROPAGATE:  A releaseShared should be propagated to other
 *               nodes. This is set (for head node only) in
 *               doReleaseShared to ensure propagation
 *               continues, even if other operations have
 *               since intervened.
 *   0:          None of the above
 *
 * The values are arranged numerically to simplify use.
 * Non-negative values mean that a node doesn't need to
 * signal. So, most code doesn't need to check for particular
 * values, just for sign.
 *
 * The field is initialized to 0 for normal sync nodes, and
 * CONDITION for condition nodes.  It is modified using CAS
 * (or when possible, unconditional volatile writes).
 */
volatile int waitStatus;

volatile Node prev;
volatile Node next;

/**
 * The thread that enqueued this node.  Initialized on
 * construction and nulled out after use.
 */
volatile Thread thread;
```
