* Direct
    * 分配、释放慢
    * 读写快
* Heap
    * 分配、释放快
    * 读写慢

---

## Data Stream

* Heap
    * Local IO -> Direct -> Heap -> Direct -> Local IO
* Direct
    * Local IO -> Direct -> Local IO    

---

## More About Direct Memory

* 避免在直接内存和堆内存间复制
* 不受堆内存大小限制，受物理内存大小限制

---

## Direct Scenes

* 大量数据，生命周期长
* 频繁 IO 读写，如 网络并发

---
