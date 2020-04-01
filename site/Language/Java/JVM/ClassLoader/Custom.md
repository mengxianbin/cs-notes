[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[JVM](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM) /
[ClassLoader](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/ClassLoader) /
[Custom](https://mengxianbin.github.io/cs-notes/site/Language/Java/JVM/ClassLoader/Custom)

```java
/**
 * Custom loader first, then the system loader.
 * 
 * @see java.net.URLClassLoader
 */
public class CostomClassLoader extends ClassLoader {

    private String classpath;

    public CostomClassLoader(String classpath) {
        super(null);
        this.classpath = classpath;
    }

    @Override
    protected Class<?> findClass(String name) throws ClassNotFoundException {
        String classFilePath = getClassFilePath(name);
        byte[] bytes = getClassBytes(classFilePath);
        Class<?> klass = defineClass(null, bytes, 0, bytes.length);
        return klass != null ? klass : super.findClass(name);
    }

    private byte[] getClassBytes(String name) {
        try (FileInputStream inputStream = new FileInputStream(name);
             FileChannel inputChannel = inputStream.getChannel();

             ByteArrayOutputStream outputStream = new ByteArrayOutputStream();
             WritableByteChannel outputChannel = Channels.newChannel(outputStream)) {

            ByteBuffer buffer = ByteBuffer.allocate(1024);
            while (inputChannel.read(buffer) > 0) {
                buffer.flip();
                outputChannel.write(buffer);
                buffer.clear();
            }

            return outputStream.toByteArray();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    private String getClassFilePath(String name) {
        return String.format("%s/%s.class", classpath, name.replace('.', '/'));
    }
}
```
