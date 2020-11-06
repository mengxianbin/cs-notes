[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Basic](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic) /
[Network](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic/Network) /
[InetAddressCachePolicy](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic/Network/InetAddressCachePolicy)

> JDK 1.8

```java
public final class InetAddressCachePolicy {
    static {
        Integer var0 = (Integer)AccessController.doPrivileged(new PrivilegedAction<Integer>() {
            public Integer run() {
                String var1;
                try {
                    var1 = Security.getProperty("networkaddress.cache.ttl");
                    if (var1 != null) {
                        return Integer.valueOf(var1);
                    }
                } catch (NumberFormatException var3) {
                }

                try {
                    var1 = System.getProperty("sun.net.inetaddr.ttl");
                    if (var1 != null) {
                        return Integer.decode(var1);
                    }
                } catch (NumberFormatException var2) {
                }

                return null;
            }
        });
        if (var0 != null) {
            cachePolicy = var0;
            if (cachePolicy < 0) {
                cachePolicy = -1;
            }

            propertySet = true;
        } else if (System.getSecurityManager() == null) {
            cachePolicy = 30;
        }

        var0 = (Integer)AccessController.doPrivileged(new PrivilegedAction<Integer>() {
            public Integer run() {
                String var1;
                try {
                    var1 = Security.getProperty("networkaddress.cache.negative.ttl");
                    if (var1 != null) {
                        return Integer.valueOf(var1);
                    }
                } catch (NumberFormatException var3) {
                }

                try {
                    var1 = System.getProperty("sun.net.inetaddr.negative.ttl");
                    if (var1 != null) {
                        return Integer.decode(var1);
                    }
                } catch (NumberFormatException var2) {
                }

                return null;
            }
        });
        if (var0 != null) {
            negativeCachePolicy = var0;
            if (negativeCachePolicy < 0) {
                negativeCachePolicy = -1;
            }

            propertyNegativeSet = true;
        }

    }
    ...
}
```
