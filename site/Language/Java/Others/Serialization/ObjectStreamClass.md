[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Others](https://mengxianbin.github.io/cs-notes/site/Language/Java/Others) /
[Serialization](https://mengxianbin.github.io/cs-notes/site/Language/Java/Others/Serialization) /
[ObjectStreamClass](https://mengxianbin.github.io/cs-notes/site/Language/Java/Others/Serialization/ObjectStreamClass)

```java
    /** class-defined writeObject method, or null if none */
    private Method writeObjectMethod;
    /** class-defined readObject method, or null if none */
    private Method readObjectMethod;
    /** class-defined readObjectNoData method, or null if none */
    private Method readObjectNoDataMethod;
    /** class-defined writeReplace method, or null if none */
    private Method writeReplaceMethod;
    /** class-defined readResolve method, or null if none */
    private Method readResolveMethod;
```

```java
    /**
     * Invokes the readObject method of the represented serializable class.
     * Throws UnsupportedOperationException if this class descriptor is not
     * associated with a class, or if the class is externalizable,
     * non-serializable or does not define readObject.
     */
    void invokeReadObject(Object obj, ObjectInputStream in)
        throws ClassNotFoundException, IOException,
               UnsupportedOperationException
    {...}
```

```java
    /**
     * Invokes the readResolve method of the represented serializable class and
     * returns the result.  Throws UnsupportedOperationException if this class
     * descriptor is not associated with a class, or if the class is
     * non-serializable or does not define readResolve.
     */
    Object invokeReadResolve(Object obj)
        throws IOException, UnsupportedOperationException
    {...}
```
