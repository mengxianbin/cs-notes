[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Theory](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Theory) /
[Interrupt and Exception](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Theory/Interrupt%20and%20Exception) /
[Outer Interruption Flow](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Theory/Interrupt%20and%20Exception/Outer%20Interruption%20Flow)

* 每条指令执行结束后，CPU 检查是否有外部中断信号

* 若有外部中断信号，需要保护被中断进程的 CPU 环境

* 根据中断信号类型，转入相应的中断处理程序

* 恢复原进程的 CPU 环境，并退出中断，返回原进程继续往下执行

---
