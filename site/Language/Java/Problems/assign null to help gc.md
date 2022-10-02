[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Problems](https://mengxianbin.github.io/cs-notes/site/Language/Java/Problems) /
[assign null to help gc](https://mengxianbin.github.io/cs-notes/site/Language/Java/Problems/assign%20null%20to%20help%20gc)

# Will null-assigning help GC ?

> https://blog.csdn.net/dfdsggdgg/article/details/52463882

---

## Follow the Above Test

- My JDK Version: 1.8.0_152
- My VM Options: `-Xmx5M`

```java
import java.util.ArrayList;
import java.util.List;

public class Test {
    public static void main(String[] args) {
        boolean willAssignNull = true;

        List<String> list = new ArrayList<>();
        for (int i = 0; i < 100000; i++) {
            String a = new String("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");
            list.add(a);
        }
        if (willAssignNull) {
            list = null;
        }

        List<String> list2 = new ArrayList<>();
        for (int i = 0; i < 100000; i++) {
            String a = new String("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");
            list2.add(a);
        }
        if (willAssignNull) {
            list2 = null;
        }

        List<String> list3 = new ArrayList<>();
        for (int i = 0; i < 100000; i++) {
            String a = new String("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");
            list3.add(a);
        }
        if (willAssignNull) {
            list3 = null;
        }

        List<String> list4 = new ArrayList<>();
        for (int i = 0; i < 100000; i++) {
            String a = new String("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");
            list4.add(a);
        }
        if (willAssignNull) {
            list4 = null;
        }

        List<String> list5 = new ArrayList<>();
        for (int i = 0; i < 100000; i++) {
            String a = new String("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa");
            list5.add(a);
        }
        if (willAssignNull) {
            list5 = null;
        }
    }
}
```

## Summary

Yes, it's true.

---
