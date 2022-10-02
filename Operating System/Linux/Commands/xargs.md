```man
NAME
       xargs - build and execute command lines from standard input

SYNOPSIS
       xargs  [-0prtx]  [-E  eof-str]  [-e[eof-str]]  [--eof[=eof-str]]  [--null]  [-d  delimiter]  [--delimiter  delimiter]  [-I  replace-str] [-i[replace-str]]
       [--replace[=replace-str]] [-l[max-lines]] [-L max-lines] [--max-lines[=max-lines]] [-n max-args] [--max-args=max-args]  [-s  max-chars]  [--max-chars=max-
       chars]  [-P  max-procs]  [--max-procs=max-procs]  [--interactive]  [--verbose]  [--exit] [--no-run-if-empty] [--arg-file=file] [--show-limits] [--version]
       [--help] [command [initial-arguments]]

OPTIONS
       -I replace-str
              Replace  occurrences  of  replace-str  in  the initial-arguments with names read from standard input.  Also, unquoted blanks do not terminate input
              items; instead the separator is the newline character.  Implies -x and -L 1.

       --replace[=replace-str]
       -i[replace-str]
              This option is a synonym for -Ireplace-str if replace-str is specified, and for -I{} otherwise.  This option is deprecated; use -I instead.
```
