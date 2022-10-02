[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Commands](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands) /
[shopt](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands/shopt)

> bash built-in
> not POSIX-compliant

---

> manual

```man
    Set and unset shell options.

    Change the setting of each shell option OPTNAME.  Without any option
    arguments, list each supplied OPTNAME, or all shell options if no
    OPTNAMEs are given, with an indication of whether or not each is set.

    Options:
      -o        restrict OPTNAMEs to those defined for use with `set -o'
      -p        print each shell option with an indication of its status
      -q        suppress output
      -s        enable (set) each OPTNAME
      -u        disable (unset) each OPTNAME

    Exit Status:
    Returns success if OPTNAME is enabled; fails if an invalid option is
    given or OPTNAME is disabled.
```

---
