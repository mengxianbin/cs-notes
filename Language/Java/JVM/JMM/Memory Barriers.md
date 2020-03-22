## Catogories

### LoadLoad Barriers

The sequence: Load1; LoadLoad; Load2

### StoreStore Barriers

The sequence: Store1; StoreStore; Store2

### LoadStore Barriers

The sequence: Load1; LoadStore; Store2

### StoreLoad Barriers

The sequence: Store1; StoreLoad; Load2

---

## Volatile

### volatile write

1. StoreStore Barrier
1. Volatile Store (Write)
1. StoreLoad Barrier

### volatile read

1. Volatile Load (Read)
1. LoadLoad Barrier
1. LoadStore Barrier

---
