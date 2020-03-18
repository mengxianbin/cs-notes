[Home](https://mengxianbin.github.io) /
[cs-note](https://mengxianbin.github.io/cs-note) /
[Operate System](https://mengxianbin.github.io/cs-note/content/Operate%20System) /
[Linux](https://mengxianbin.github.io/cs-note/content/Operate%20System/Linux) /
[Commands](https://mengxianbin.github.io/cs-note/content/Operate%20System/Linux/Commands) /

```sh
# lsof -i:8000
COMMAND   PID USER   FD   TYPE   DEVICE SIZE/OFF NODE NAME
nodejs  26993 root   10u  IPv4 37999514      0t0  TCP *:8000 (LISTEN)
```

```man
NAME
       lsof - list open files

OPTIONS
       In the absence of any options, lsof lists all open files belonging to all active processes.
       
       -i [i]   selects the listing of files any of whose Internet address matches the address specified in i.  If  no
                address  is specified, this option selects the listing of all Internet and x.25 (HP-UX) network files.
```
