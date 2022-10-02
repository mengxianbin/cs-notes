[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[JVM](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM) /
[Options](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/Options) /
[Concrete](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/Options/Concrete) /
[gc log](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/Options/Concrete/gc%20log)

## GC Logging Flags in Java 8 and Earlier

- -XX:+PrintGC

- -XX:+PrintGCDetails

- -XX:+PrintGCDateStamps and -XX:+PrintGCTimeStamps

- -Xloggc

```sh
java -cp $CLASSPATH -Xloggc mypackage.MainClass # stdout
java -cp $CLASSPATH -Xloggc:/tmp/gc.log mypackage.MainClass
```

## GC Logging Flags in Java 9 and Later

- -Xlog

```sh
java -Xlog:logging=debug -version
java -cp $CLASSPATH -Xlog:gc*=debug:file=/tmp/gc.log mypackage.MainClass
java -cp $CLASSPATH -Xlog:gc*=debug:stdout -Xlog:gc*=debug:file=/tmp/gc.log mypackage.MainClass
```

## References

- https://www.baeldung.com/java-gc-logging-to-file

---
