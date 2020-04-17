[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Spring](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring) /
[Boot](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot) /
[Start](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot/Start) /
[Runner](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot/Start/Runner) /
[CommandLineRunner](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot/Start/Runner/CommandLineRunner)

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
