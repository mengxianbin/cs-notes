[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[JVM](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM) /
[Options](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/Options) /
[Concrete](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/Options/Concrete) /
[dump](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/Options/Concrete/dump)

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
