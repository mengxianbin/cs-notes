```java
    static class ExtClassLoader extends URLClassLoader {
        private static volatile Launcher.ExtClassLoader instance;

        public static Launcher.ExtClassLoader getExtClassLoader() throws IOException {
            if (instance == null) {
                Class var0 = Launcher.ExtClassLoader.class;
                synchronized(Launcher.ExtClassLoader.class) {
                    if (instance == null) {
                        instance = createExtClassLoader();
                    }
                }
            }

            return instance;
        }

        private static Launcher.ExtClassLoader createExtClassLoader() throws IOException {
            try {
                return (Launcher.ExtClassLoader)AccessController.doPrivileged(new PrivilegedExceptionAction<Launcher.ExtClassLoader>() {
                    public Launcher.ExtClassLoader run() throws IOException {
                        File[] var1 = Launcher.ExtClassLoader.getExtDirs();
                        int var2 = var1.length;

                        for(int var3 = 0; var3 < var2; ++var3) {
                            MetaIndex.registerDirectory(var1[var3]);
                        }

                        return new Launcher.ExtClassLoader(var1);
                    }
                });
            } catch (PrivilegedActionException var1) {
                throw (IOException)var1.getException();
            }
        }

        public ExtClassLoader(File[] var1) throws IOException {
            super(getExtURLs(var1), (ClassLoader)null, Launcher.factory);
            SharedSecrets.getJavaNetAccess().getURLClassPath(this).initLookupCache(this);
        }

        ...
    }
```
