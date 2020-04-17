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
