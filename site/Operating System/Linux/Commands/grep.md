[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Commands](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands) /
[grep](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands/grep)

* -c, --count
* -i, --ignore-case
* -v, --invert-match
* -m NUM, --max-count=NUM
* -o, --only-matching
* -w, --word-regexp
* -P, --perl-regexp

---

* \\ Escape

---

```sh
   Context Line Control

       -A NUM, --after-context=NUM
              Print  NUM  lines  of  trailing  context  after  matching lines.
              Places  a  line  containing  a  group  separator  (--)   between
              contiguous  groups  of  matches.  With the -o or --only-matching
              option, this has no effect and a warning is given.

       -B NUM, --before-context=NUM
              Print NUM  lines  of  leading  context  before  matching  lines.
              Places   a  line  containing  a  group  separator  (--)  between
              contiguous groups of matches.  With the  -o  or  --only-matching
              option, this has no effect and a warning is given.

       -C NUM, -NUM, --context=NUM
              Print  NUM  lines of output context.  Places a line containing a
              group separator (--) between contiguous groups of matches.  With
              the  -o  or  --only-matching  option,  this  has no effect and a
              warning is given.
```

---
