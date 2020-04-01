[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Concurrent](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent) /
[Summary](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Summary) /
[DeadLock](https://mengxianbin.github.io/cs-notes/site/Language/Java/Concurrent/Summary/DeadLock)

## 死锁的条件

* 互斥
* 占有且等待
* 不可抢占
* 循环等待

以上同时满足

## 解决方案：破坏以上任意一个条件即可

* 占有且等待： 一次申请全部所需资源
* 不可抢占： 所需其一申请失败，就释放所有已有
* 循环等待： 按序申请资源

---
