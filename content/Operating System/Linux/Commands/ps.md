[Home](https://mengxianbin.github.io) /
[cs-note](https://mengxianbin.github.io/cs-note/content) /
[Operating System](https://mengxianbin.github.io/cs-note/content/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-note/content/Operating%20System/Linux) /
[Commands](https://mengxianbin.github.io/cs-note/content/Operating%20System/Linux/Commands) /
[ps](https://mengxianbin.github.io/cs-note/content/Operating%20System/Linux/Commands/ps)


```
NAME
       ps - report a snapshot of the current processes.

SYNOPSIS
       ps [options]

DESCRIPTION
       ps displays information about a selection of the active processes.  If you want a repetitive update of the selection and the displayed information, use top(1) instead.

       This version of ps accepts several kinds of options:

       1   UNIX options, which may be grouped and must be preceded by a dash.
       2   BSD options, which may be grouped and must not be used with a dash.
       3   GNU long options, which are preceded by two dashes.

EXAMPLES
       To see every process on the system using standard syntax:
          ps -e
          ps -ef
          ps -eF
          ps -ely

       To see every process on the system using BSD syntax:
          ps ax
          ps axu              
```
