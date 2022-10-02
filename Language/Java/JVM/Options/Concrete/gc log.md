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
