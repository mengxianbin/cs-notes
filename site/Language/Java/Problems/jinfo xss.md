[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Problems](https://mengxianbin.github.io/cs-notes/site/Language/Java/Problems) /
[jinfo xss](https://mengxianbin.github.io/cs-notes/site/Language/Java/Problems/jinfo%20xss)

# Xss Default Value ?

> JDK 1.8.0_152

## Windows 10

```bat
> jinfo -flag ThreadStackSize PID
-XX:ThreadStackSize=0
```

```bat
> java -XX:+PrintFlagsFinal -version | findstr ThreadStackSize
     intx CompilerThreadStackSize                   = 0                                   {pd product}
     intx ThreadStackSize                           = 0                                   {pd product}
     intx VMThreadStackSize                         = 0                                   {pd product}
```

## Ubuntu 20.04

```shell
$ jinfo -flag ThreadStackSize PID
-XX:ThreadStackSize=1024
```

```shell
$ java -XX:+PrintFlagsFinal -version | grep ThreadStackSize
     intx CompilerThreadStackSize                  = 1024                                   {pd product} {default}
     intx ThreadStackSize                          = 1024                                   {pd product} {default}
     intx VMThreadStackSize                        = 1024                                   {pd product} {default}
```

---

## Stack Overflow

...

## Official Document

- https://docs.oracle.com/javase/8/docs/technotes/tools/unix/java.html
- https://docs.oracle.com/en/java/javase/11/tools/java.html#GUID-3B1CE181-CD30-4178-9602-230B800D4FAE

---
