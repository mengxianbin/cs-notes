[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[JVM](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM) /
[ClassLoader](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/ClassLoader) /
[Launcher](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/ClassLoader/Launcher) /
[Launcher](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/ClassLoader/Launcher/Launcher)

```java
    public ClassLoader getClassLoader() {
        return this.loader;
    }
```

```java
    public Launcher() {
        Launcher.ExtClassLoader var1;
        try {
            var1 = Launcher.ExtClassLoader.getExtClassLoader();
        } catch (IOException var10) {
            throw new InternalError("Could not create extension class loader", var10);
        }

        try {
            this.loader = Launcher.AppClassLoader.getAppClassLoader(var1);
        } catch (IOException var9) {
            throw new InternalError("Could not create application class loader", var9);
        }
        ...
    }
```
