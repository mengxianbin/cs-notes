[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Collection](https://mengxianbin.github.io/cs-notes/site/Language/Java/Collection) /
[HashMap](https://mengxianbin.github.io/cs-notes/site/Language/Java/Collection/HashMap) /
[loadFactor](https://mengxianbin.github.io/cs-notes/site/Language/Java/Collection/HashMap/loadFactor)

```java
    /**
     * The load factor used when none specified in constructor.
     */
    static final float DEFAULT_LOAD_FACTOR = 0.75f;
```

```java
/**
 ...
 * <p>As a general rule, the default load factor (.75) offers a good
 * tradeoff between time and space costs.  Higher values decrease the
 * space overhead but increase the lookup cost (reflected in most of
 * the operations of the {@code HashMap} class, including
 * {@code get} and {@code put}).  The expected number of entries in
 * the map and its load factor should be taken into account when
 * setting its initial capacity, so as to minimize the number of
 * rehash operations.  If the initial capacity is greater than the
 * maximum number of entries divided by the load factor, no rehash
 ...
```

