[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/content) /
[Language](https://mengxianbin.github.io/cs-notes/content/Language) /
[Java](https://mengxianbin.github.io/cs-notes/content/Language/Java) /
[Others](https://mengxianbin.github.io/cs-notes/content/Language/Java/Others) /
[Serialization](https://mengxianbin.github.io/cs-notes/content/Language/Java/Others/Serialization) /
[ObjectInputStream](https://mengxianbin.github.io/cs-notes/content/Language/Java/Others/Serialization/ObjectInputStream)

```java
    /**
     * Reads (or attempts to skip, if obj is null or is tagged with a
     * ClassNotFoundException) instance data for each serializable class of
     * object in stream, from superclass to subclass.  Expects that passHandle
     * is set to obj's handle before this method is called.
     */
    private void readSerialData(Object obj, ObjectStreamClass desc)
    {
        ObjectStreamClass.ClassDataSlot[] slots = desc.getClassDataLayout();
        for (int i = 0; i < slots.length; i++) {
            ObjectStreamClass slotDesc = slots[i].desc;
        ...
        slotDesc.invokeReadObject(obj, this);
        ...
    }
```

```java
    /**
     * Reads and returns "ordinary" (i.e., not a String, Class,
     * ObjectStreamClass, array, or enum constant) object, or null if object's
     * class is unresolvable (in which case a ClassNotFoundException will be
     * associated with object's handle).  Sets passHandle to object's assigned
     * handle.
     */
    private Object readOrdinaryObject(boolean unshared)
        throws IOException
    {
        ...
        readSerialData(obj, desc);
        ...
        Object rep = desc.invokeReadResolve(obj);
        ...
        handles.setObject(passHandle, obj = rep);
        ...
        return obj;
    }
```

```java
    /**
     * Underlying readObject implementation.
     * @param type a type expected to be deserialized; non-null
     * @param unshared true if the object can not be a reference to a shared object, otherwise false
     */
    private Object readObject0(Class<?> type, boolean unshared) throws IOException {
        ...
        return checkResolve(readOrdinaryObject(unshared));
        ...
    }
```

```java
    /**
     * Internal method to read an object from the ObjectInputStream of the expected type.
     * Called only from {@code readObject()} and {@code readString()}.
     * Only {@code Object.class} and {@code String.class} are supported.
     *
     * @param type the type expected; either Object.class or String.class
     * @return an object of the type
     * @throws  IOException Any of the usual Input/Output related exceptions.
     * @throws  ClassNotFoundException Class of a serialized object cannot be
     *          found.
     */
    private final Object readObject(Class<?> type)
        throws IOException, ClassNotFoundException
    {
        ...
        Object obj = readObject0(type, false);
        ...
    }
```
