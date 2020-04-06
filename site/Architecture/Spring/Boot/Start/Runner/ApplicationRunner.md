[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Spring](https://mengxianbin.github.io/cs-notes/site/Architecture/Spring) /
[Boot](https://mengxianbin.github.io/cs-notes/site/Architecture/Spring/Boot) /
[Start](https://mengxianbin.github.io/cs-notes/site/Architecture/Spring/Boot/Start) /
[Runner](https://mengxianbin.github.io/cs-notes/site/Architecture/Spring/Boot/Start/Runner) /
[ApplicationRunner](https://mengxianbin.github.io/cs-notes/site/Architecture/Spring/Boot/Start/Runner/ApplicationRunner)

```java
public interface ApplicationRunner 
{    
/**

     * Callback used to run the bean.

     * @param args incoming application arguments

     * @throws Exception on error

     */
    void run(ApplicationArguments args) throws Exception;
}
```
