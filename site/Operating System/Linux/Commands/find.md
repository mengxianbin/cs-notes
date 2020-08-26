[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Commands](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands) /
[find](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands/find)

```sh
# find . -name "*.c"
# find . -type f
# find . -ctime -20
# find /var/log -type f -mtime +7 -ok rm {} \;
# find . -type f -perm 644 -exec ls -l {} \;
# find / -type f -size 0 -exec ls -l {} \;
```
