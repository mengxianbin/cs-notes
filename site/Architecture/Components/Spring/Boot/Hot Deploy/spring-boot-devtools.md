[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Architecture](https://mengxianbin.github.io/cs-notes/site/Architecture) /
[Components](https://mengxianbin.github.io/cs-notes/site/Architecture/Components) /
[Spring](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring) /
[Boot](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot) /
[Hot Deploy](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot/Hot%20Deploy) /
[spring-boot-devtools](https://mengxianbin.github.io/cs-notes/site/Architecture/Components/Spring/Boot/Hot%20Deploy/spring-boot-devtools)

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-devtools</artifactId>
    <optional>true</optional>
</dependency>
```

```properites
# 热部署开关，false即不启用热部署
spring.devtools.restart.enabled: true

# 指定热部署的目录
#spring.devtools.restart.additional-paths: src/main/java

# 指定目录不更新
spring.devtools.restart.exclude: test

```
