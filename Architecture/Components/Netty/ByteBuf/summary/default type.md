## pooled or not

ByteBufUtil.static

```java
    static {

        String allocType = SystemPropertyUtil.get(
                "io.netty.allocator.type", PlatformDependent.isAndroid() ? "unpooled" : "pooled");
        allocType = allocType.toLowerCase(Locale.US).trim();
        ByteBufAllocator alloc;
      
        if ("unpooled".equals(allocType)) {
            alloc = UnpooledByteBufAllocator.DEFAULT;
        } else if ("pooled".equals(allocType)) {
            alloc = PooledByteBufAllocator.DEFAULT;
        } else {
            alloc = PooledByteBufAllocator.DEFAULT;
        }

        DEFAULT_ALLOCATOR = alloc;
    }
```

## direct or heap

PlatformDependent.static

```java
	static {
		if (!isAndroid()) {
			if (javaVersion() >= 9) {
				CLEANER = CleanerJava9.isSupported() ? new CleanerJava9() : NOOP;
			} else {
				CLEANER = CleanerJava6.isSupported() ? new CleanerJava6() : NOOP;
			}
		} else {
			CLEANER = NOOP;
		}
		DIRECT_BUFFER_PREFERRED = CLEANER != NOOP && !SystemPropertyUtil.getBoolean("io.netty.noPreferDirect", false);
	}
```

## summary

- pooled
- direct

## ref

> https://www.hicode.club/articles/2020/07/25/1595690469536.html

---
