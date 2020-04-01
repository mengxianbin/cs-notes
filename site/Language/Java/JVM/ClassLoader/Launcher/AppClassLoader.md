[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[JVM](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM) /
[ClassLoader](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/ClassLoader) /
[Launcher](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/ClassLoader/Launcher) /
[AppClassLoader](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/ClassLoader/Launcher/AppClassLoader)

```java
    static class AppClassLoader extends URLClassLoader {
        final URLClassPath ucp = SharedSecrets.getJavaNetAccess().getURLClassPath(this);

        public static ClassLoader getAppClassLoader(final ClassLoader var0) throws IOException {
            final String var1 = System.getProperty("java.class.path");
            final File[] var2 = var1 == null ? new File[0] : Launcher.getClassPath(var1);
            return (ClassLoader)AccessController.doPrivileged(new PrivilegedAction<Launcher.AppClassLoader>() {
                public Launcher.AppClassLoader run() {
                    URL[] var1x = var1 == null ? new URL[0] : Launcher.pathToURLs(var2);
                    return new Launcher.AppClassLoader(var1x, var0);
                }
            });
        }

        AppClassLoader(URL[] var1, ClassLoader var2) {
            super(var1, var2, Launcher.factory);
            this.ucp.initLookupCache(this);
        }
        
        ...
    }
```
