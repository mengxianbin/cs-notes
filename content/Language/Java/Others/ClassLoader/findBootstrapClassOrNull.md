[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[Others](https://mengxianbin.github.io/cs-notes/content/Language/Java/Others) /
[ClassLoader](https://mengxianbin.github.io/cs-notes/content/Language/Java/Others/ClassLoader) /
[findBootstrapClassOrNull](https://mengxianbin.github.io/cs-notes/content/Language/Java/Others/ClassLoader/findBootstrapClassOrNull)

```java
    /**
     * Returns a class loaded by the bootstrap class loader;
     * or return null if not found.
     */
    private Class<?> findBootstrapClassOrNull(String name)
    {
        if (!checkName(name)) return null;

        return findBootstrapClass(name);
    }

    // return null if not found
    private native Class<?> findBootstrapClass(String name);
```
