[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Language](https://mengxianbin.github.io/cs-notes/site/Language) /
[Batch](https://mengxianbin.github.io/cs-notes/site/Language/Batch) /
[netsh](https://mengxianbin.github.io/cs-notes/site/Language/Batch/netsh)

# netsh (Network Shell)

- netsh interface portproxy show all
- netsh interface portproxy add ...
- netsh interface portproxy delete ...

---

> netsh interface portproxy add v4tov4 /?

```
       listenport     - IPv4 port on which to listen.
       connectaddress - IPv4 address to which to connect.
       connectport    - IPv4 port to which to connect.
       listenaddress  - IPv4 address on which to listen.
       protocol       - Protocol to use.  Currently only TCP is supported.
```

---
