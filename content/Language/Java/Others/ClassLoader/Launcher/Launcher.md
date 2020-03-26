[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[Others](https://mengxianbin.github.io/cs-notes/content/Language/Java/Others) /
[ClassLoader](https://mengxianbin.github.io/cs-notes/content/Language/Java/Others/ClassLoader) /
[Launcher](https://mengxianbin.github.io/cs-notes/content/Language/Java/Others/ClassLoader/Launcher) /
[Launcher](https://mengxianbin.github.io/cs-notes/content/Language/Java/Others/ClassLoader/Launcher/Launcher)

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
