```sh
# find . -name "*.c"
# find . -type f
# find . -ctime -20
# find /var/log -type f -mtime +7 -ok rm {} \;
# find . -type f -perm 644 -exec ls -l {} \;
# find / -type f -size 0 -exec ls -l {} \;
```
