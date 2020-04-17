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
