
## Java 9

* Optional.ifPresentOrElse

---

## Java 8

```java
package org.springframework.data.util;
...
public interface Optionals {
...
	public static <T> void ifPresentOrElse(Optional<T> optional, Consumer<? super T> consumer, Runnable runnable) {

		Assert.notNull(optional, "Optional must not be null!");
		Assert.notNull(consumer, "Consumer must not be null!");
		Assert.notNull(runnable, "Runnable must not be null!");

		if (optional.isPresent()) {
			optional.ifPresent(consumer);
		} else {
			runnable.run();
		}
	}
```

```xml
		<dependency>
			<groupId>org.springframework.data</groupId>
			<artifactId>spring-data-commons</artifactId>
			<version>${springdata.commons}</version>
		</dependency>
```
