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
