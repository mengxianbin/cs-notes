```java
    private SelectionKey findKey(Selector sel) {
        assert Thread.holdsLock(keyLock);
        if (keys == null)
            return null;
        for (int i = 0; i < keys.length; i++)
            if ((keys[i] != null) && (keys[i].selector() == sel))
                return keys[i];
        return null;

    }
```
