```java
public interface CommandLineRunner
{
    /**
     * Callback used to run the bean.
     * @param args incoming main method arguments
     * @throws Exception on error
     */
    void run (String... args)throws Exception;
}
```
```java
@Component
public class MyBean implements CommandLineRunner
{
    public void run(String... args)
    {
        // Do something...
    }
}
```

```java
@Bean
public
CommandLineRunner init()
{
    return (String... strings) -> {

    };
}
```
