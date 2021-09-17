```
log4j-core-2.6.2.jar
```

## AbstractOutputStreamAppender

> org.apache.logging.log4j.core.appender

```java
    protected void directEncodeEvent(final LogEvent event) {
        getLayout().encode(event, manager);
        if (this.immediateFlush || event.isEndOfBatch()) {
            manager.flush();
        }
    }
```

## PatternLayout

> org.apache.logging.log4j.core.layout

```java
        @Override
        public StringBuilder toSerializable(final LogEvent event, final StringBuilder buffer) {
            final int len = formatters.length;
            for (int i = 0; i < len; i++) {
                formatters[i].format(event, buffer);
            }
            if (replace != null) { // creates temporary objects
                String str = buffer.toString();
                str = replace.format(str);
                buffer.setLength(0);
                buffer.append(str);
            }
            return buffer;
        }
```

## FileLocationPatternConverter

> org.apache.logging.log4j.core.pattern

```java
    @Override
    public void format(final LogEvent event, final StringBuilder output) {
        final StackTraceElement element = event.getSource();

        if (element != null) {
            output.append(element.getFileName());
        }
    }
```

## LineLocationPatternConverter

> org.apache.logging.log4j.core.pattern

```java
    @Override
    public void format(final LogEvent event, final StringBuilder output) {
        final StackTraceElement element = event.getSource();

        if (element != null) {
            output.append(element.getLineNumber());
        }
    }
```

## MutableLogEvent

> org.apache.logging.log4j.core.impl

```java
    @Override
    public StackTraceElement getSource() {
        if (source != null) {
            return source;
        }
        if (loggerFqcn == null || !includeLocation) {
            return null;
        }
        source = Log4jLogEvent.calcLocation(loggerFqcn);
        return source;
    }
```

## Log4jLogEvent

> org.apache.logging.log4j.core.impl

```java
    public static StackTraceElement calcLocation(final String fqcnOfLogger) {
        if (fqcnOfLogger == null) {
            return null;
        }
        // LOG4J2-1029 new Throwable().getStackTrace is faster than Thread.currentThread().getStackTrace().
        final StackTraceElement[] stackTrace = new Throwable().getStackTrace();
        StackTraceElement last = null;
        for (int i = stackTrace.length - 1; i > 0; i--) {
            final String className = stackTrace[i].getClassName();
            if (fqcnOfLogger.equals(className)) {
                return last;
            }
            last = stackTrace[i];
        }
        return null;
    }
```

---
