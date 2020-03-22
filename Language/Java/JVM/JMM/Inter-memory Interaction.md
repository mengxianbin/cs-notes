- System modeled by constraints between actions
    - Use/assign actions correspond to thread computations
    - load/store actions correspond to thread fill/flush actions
    - read/write actions are main memory actions

---

## Operation type

* Lock: A variable that acts on main memory, which identifies a variable as a thread-exclusive state.
* Unlock: Acts on the main memory variable, which releases a locked variable, and the released variable can be locked by other threads.

* ...

---
