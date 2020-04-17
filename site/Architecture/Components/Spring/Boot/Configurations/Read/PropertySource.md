[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Spring](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring) /
[Boot](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot) /
[Configurations](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot/Configurations) /
[Read](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot/Configurations/Read) /
[PropertySource](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot/Configurations/Read/PropertySource)

* config/a.properties

```properties
db.username = username
```

* @PropertySource + @Value
* @PropertySource + @ConfigurationProperties

```java
@Component
@PropertySource("config/a.properties")
class A {
    @Value("db.username")
    private String username;
}
```

---

* @PropertySource not support yml

---
