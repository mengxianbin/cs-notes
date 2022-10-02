[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Collection](https://mengxianbin.github.io/cs-notes/site/Language/Java/Collection) /
[CopyOnWriteArrayList](https://mengxianbin.github.io/cs-notes/site/Language/Java/Collection/CopyOnWriteArrayList) /
[CopyOnWriteArrayList](https://mengxianbin.github.io/cs-notes/site/Language/Java/Collection/CopyOnWriteArrayList/CopyOnWriteArrayList)

```java
public boolean add(E e) {
    final ReentrantLock lock = this.lock;
    lock.lock();
    try {
        Object[] elements = getArray();
        int len = elements.length;
        Object[] newElements = Arrays.copyOf(elements, len + 1);
        newElements[len] = e;
        setArray(newElements);
        return true;
    } finally {
        lock.unlock();
    }
}

final void setArray(Object[] a) {
    array = a;
}

@SuppressWarnings("unchecked")
private E get(Object[] a, int index) {
    return (E) a[index];
}
```
