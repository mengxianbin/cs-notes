* /META-INF/spring.factories

```properties
# Initializers
org.springframework.context.ApplicationContextInitializer=\
...

# Application Listeners
org.springframework.context.ApplicationListener=\
...

# Auto Configuration Import Listeners
org.springframework.boot.autoconfigure.AutoConfigurationImportListener=\
...

# Auto Configuration Import Filters
org.springframework.boot.autoconfigure.AutoConfigurationImportFilter=\
...

# Auto Configure
org.springframework.boot.autoconfigure.EnableAutoConfiguration=\
...

# Failure analyzers
org.springframework.boot.diagnostics.FailureAnalyzer=\
...

# Template availability providers
org.springframework.boot.autoconfigure.template.TemplateAvailabilityProvider=\
...
```

---

* XXXAutoConfiguration.java
    * eg. JpaRepositoriesAutoConfiguration

```java
@Configuration
@ConditionalOnBean(DataSource.class)
@ConditionalOnClass(JpaRepository.class)
@ConditionalOnMissingBean({ JpaRepositoryFactoryBean.class, JpaRepositoryConfigExtension.class })
@ConditionalOnProperty(prefix = "spring.data.jpa.repositories", name = "enabled", havingValue = "true",
		matchIfMissing = true)
@Import(JpaRepositoriesAutoConfigureRegistrar.class)
@AutoConfigureAfter({ HibernateJpaAutoConfiguration.class, TaskExecutionAutoConfiguration.class })
public class JpaRepositoriesAutoConfiguration {
    ...
```

---
