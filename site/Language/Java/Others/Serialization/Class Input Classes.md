[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Java](https://mengxianbin.github.io/cs-notes/site/Language/Java) /
[Others](https://mengxianbin.github.io/cs-notes/site/Language/Java/Others) /
[Serialization](https://mengxianbin.github.io/cs-notes/site/Language/Java/Others/Serialization) /
[Class Input Classes](https://mengxianbin.github.io/cs-notes/site/Language/Java/Others/Serialization/Class%20Input%20Classes)

* https://docs.oracle.com/javase/8/docs/platform/serialization/spec/input.html#a2971

```
3.7 The readResolve Method
For Serializable and Externalizable classes, the readResolve method allows a class to replace/resolve the object read from the stream before it is returned to the caller. By implementing the readResolve method, a class can directly control the types and instances of its own instances being deserialized. The method is defined as follows:

   ANY-ACCESS-MODIFIER Object readResolve()
                throws ObjectStreamException;
The readResolve method is called when ObjectInputStream has read an object from the stream and is preparing to return it to the caller. ObjectInputStream checks whether the class of the object defines the readResolve method. If the method is defined, the readResolve method is called to allow the object in the stream to designate the object to be returned. The object returned should be of a type that is compatible with all uses. If it is not compatible, a ClassCastException will be thrown when the type mismatch is discovered.

For example, a Symbol class could be created for which only a single instance of each symbol binding existed within a virtual machine. The readResolve method would be implemented to determine if that symbol was already defined and substitute the preexisting equivalent Symbol object to maintain the identity constraint. In this way the uniqueness of Symbol objects can be maintained across serialization.
```
