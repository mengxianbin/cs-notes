[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Spring](https://mengxianbin.github.io/cs-notes/site/Architecture/Spring) /
[Boot](https://mengxianbin.github.io/cs-notes/site/Architecture/Spring/Boot) /
[Configurations](https://mengxianbin.github.io/cs-notes/site/Architecture/Spring/Boot/Configurations) /
[Read](https://mengxianbin.github.io/cs-notes/site/Architecture/Spring/Boot/Configurations/Read) /
[ConfigurationProperties](https://mengxianbin.github.io/cs-notes/site/Architecture/Spring/Boot/Configurations/Read/ConfigurationProperties)

```yml
info:
    address: addr
    company: comp
```

```java
@Component 
@ConfigurationProperties(prefix ="info")
public class InfoConfig2
{
    private String address;
    private String company;

```
