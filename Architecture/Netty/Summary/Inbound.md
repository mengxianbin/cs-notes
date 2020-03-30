# Inbound Events

* fire: unsafe
* handle: Channel

---

* Default handler: TailContext, Empty Implementation

* Flow in pipeline: head -> tail

* call fireXXX to transfer event to next handler
    * if not, the event stops

---

1. Context.fire
1. Context.findContextInbound
1. nextContext.invoke
1. nextHandler.xxx
1. nextContext.fire

---
