[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Basic](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic) /
[Math](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic/Math) /
[Random](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic/Math/Random) /
[initialScramble](https://mengxianbin.github.io/cs-notes/site/Language/Java/Basic/Math/Random/initialScramble)

```java
    private static long initialScramble(long seed) {
        return (seed ^ multiplier) & mask;
    }
```

```java
    private static final long multiplier = 0x5DEECE66DL;
    private static final long mask = (1L << 48) - 1;
```
