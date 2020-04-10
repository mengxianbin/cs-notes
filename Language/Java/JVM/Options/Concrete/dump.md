* jmap -dump:format=b,file=e.bin pid
* jmap -histo pid

* jmap -histo:live pid
    * 会先触发 Full GC

---

* Tool
    * jvisualvm
    * Eclipse Memory Analysis Tool (MAT)

---

-XX:+HeapDumpBeforeFullGC
-XX:+HeapDumpOnOutOfMemoryError
-XX:HeapDumpPath=/tmp/heapdump.hprof

---
