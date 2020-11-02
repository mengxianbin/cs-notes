## Addition

* sys prop
    * -Dname=value

* env var
    * win SET
    * *nix EXPORT

## Retrival

* sys prop
    * System.getProperty(String): String
    * System.getProperty(): Properties

* env var
    * System.getenv(String): String
    * System.getenv(): `Map<String, String>`

## Runtime manipulate

* sys prop
    * writable
        * System.setProperty
        * System.getProperties().load

* env var
    * read-only

## Access

* sys prop
    * only current process

* env var
    * all processes on the machine

---
