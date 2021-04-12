## why in locked block

- wait: read a condition in thread A
- notify: write the same condition in thread B
- reading and writing of the same condition must be exclusive
    - ensure thread-safe
        - keep the Happen-Before rule

- notice
    - the waiting thread and the notifying thread must use the same lock

#### some others like this

- condition: await & signal/signalAll

#### ref

- https://blog.csdn.net/lengxiao1993/article/details/52296220
- https://juejin.cn/post/6844904088505679879

---

## why in loop

- https://blog.csdn.net/u013256816/article/details/106654315

---

## why use notifyAll instead of notify

- 'notify' maybe cause deadlock
    - when the waked thread go to wait other then notify
        - all threads will keep waiting

#### concept

- object
    - moniter pool
    - waiting pool

#### ref

- https://blog.csdn.net/u013256816/article/details/106654315
- https://blog.csdn.net/meism5/article/details/90238268
- https://www.cnblogs.com/YuyuanNo1/p/11549781.html

---
