[Home](https://mengxianbin.github.io) /
[cs-note](https://mengxianbin.github.io/cs-note/content) /
[Operate System](https://mengxianbin.github.io/cs-note/content/Operate%20System) /
[Linux](https://mengxianbin.github.io/cs-note/content/Operate%20System/Linux) /
[Commands](https://mengxianbin.github.io/cs-note/content/Operate%20System/Linux/Commands) /

```man
The options are as follows:

-A      Display information about other users' processes, including those without controlling terminals.

-a      Display information about other users' processes as well as your own.  This will skip any processes which do not have a controlling terminal,
        unless the -x option is also specified.

-e      Identical to -A.

-f      Display the uid, pid, parent pid, recent CPU usage, process start time, controlling tty, elapsed CPU
        usage, and the associated command.  If the -u option is also used, display the user name rather then the
        numeric uid.  When -o or -O is used to add to the display following -f, the command field is not trun-
        cated as severely as it is in other formats.

-l      Display information associated with the following keywords: uid, pid, ppid, flags, cpu, pri, nice,
        vsz=SZ, rss, wchan, state=S, paddr=ADDR, tty, time, and command=CMD.

-u      Display the processes belonging to the specified usernames.

-x      When displaying processes matched by other options, include processes which do not have a controlling terminal.  This is the opposite of the -X option.  If both -X and -x
        are specified in the same command, then ps will use the one which was specified last.
```

```
state   The state is given by a sequence of characters, for example, ``RWNA''.  The first character indicates the run state of the process:

        I       Marks a process that is idle (sleeping for longer than about 20 seconds).
        R       Marks a runnable process.
        S       Marks a process that is sleeping for less than about 20 seconds.
        T       Marks a stopped process.
        U       Marks a process in uninterruptible wait.
        Z       Marks a dead process (a ``zombie'').
```
