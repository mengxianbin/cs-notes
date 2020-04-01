[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Netty](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty) /
[Java](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java) /
[SelectorProvider](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorProvider) /
[provider](https://mengxianbin.github.io/cs-notes/site/Architecture/Netty/Java/SelectorProvider/provider)

```java
    public static SelectorProvider provider() {
        synchronized (lock) {
            if (provider != null)
                return provider;
            return AccessController.doPrivileged(
                new PrivilegedAction<>() {
                    public SelectorProvider run() {
                            if (loadProviderFromProperty())
                                return provider;
                            if (loadProviderAsService())
                                return provider;
                            provider = sun.nio.ch.DefaultSelectorProvider.create();
                            return provider;
                        }
                    });
        }
    }
```

```java
    private static boolean loadProviderFromProperty() {
        String cn = System.getProperty("java.nio.channels.spi.SelectorProvider");
        if (cn == null)
            return false;
        try {
            @SuppressWarnings("deprecation")
            Object tmp = Class.forName(cn, true,
                                       ClassLoader.getSystemClassLoader()).newInstance();
            provider = (SelectorProvider)tmp;
            return true;
        } catch (ClassNotFoundException x) {
            throw new ServiceConfigurationError(null, x);
        } catch (IllegalAccessException x) {
            throw new ServiceConfigurationError(null, x);
        } catch (InstantiationException x) {
            throw new ServiceConfigurationError(null, x);
        } catch (SecurityException x) {
            throw new ServiceConfigurationError(null, x);
        }
    }
```

```java
    private static boolean loadProviderAsService() {

        ServiceLoader<SelectorProvider> sl =
            ServiceLoader.load(SelectorProvider.class,
                               ClassLoader.getSystemClassLoader());
        Iterator<SelectorProvider> i = sl.iterator();
        for (;;) {
            try {
                if (!i.hasNext())
                    return false;
                provider = i.next();
                return true;
            } catch (ServiceConfigurationError sce) {
                if (sce.getCause() instanceof SecurityException) {
                    // Ignore the security exception, try the next provider
                    continue;
                }
                throw sce;
            }
        }
    }
```
