[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Spring](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring) /
[Boot](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot) /
[Configurations](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot/Configurations) /
[Read](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot/Configurations/Read) /
[Value](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot/Configurations/Read/Value)

```yml
info:
    address: addr
    company: comp
```

```java
public class InfoConfig1
{
    @Value("${info.address}")    
    private String address;
```
