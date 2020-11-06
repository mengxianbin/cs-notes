[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Commands](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands) /
[strace](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands/strace)

> strace - trace system calls and signals

---

```man
       -f          Trace child processes as they are created by currently traced processes as a result of the fork(2), vfork(2) and clone(2) system calls.  Note that -p PID -f will attach all threads of process PID if it is multi-threaded, not only thread       
                   with thread_id = PID.

       -e trace=syscall_set
                   file         Trace all system calls which take a file name as an argument.  You can think of this as an abbreviation for -e trace=open,stat,chmod,unlink,...  which is useful to seeing what files the process is referencing.  Furthermore,       
                                using the abbreviation will ensure that you don't accidentally forget to include a call like lstat(2) in the list.  Betchya woulda forgot that one.  The syntax without a preceding percent sign ("-e  trace=file")  is  depre‐       
                   process      Trace all system calls which involve process management.  This is useful for watching the fork, wait, and exec steps of a process.  The syntax without a preceding percent sign ("-e trace=process") is deprecated.
                   network      Trace all the network related system calls.  The syntax without a preceding percent sign ("-e trace=network") is deprecated.
                   signal       Trace all signal related system calls.  The syntax without a preceding percent sign ("-e trace=signal") is deprecated.
                   ipc          Trace all IPC related system calls.  The syntax without a preceding percent sign ("-e trace=ipc") is deprecated.
                   desc         Trace all file descriptor related system calls.  The syntax without a preceding percent sign ("-e trace=desc") is deprecated.
                   memory       Trace all memory mapping related system calls.  The syntax without a preceding percent sign ("-e trace=memory") is deprecated.
                   %statfs      Trace statfs, statfs64, statvfs, osf_statfs, and osf_statfs64 system calls.  The same effect can be achieved with -e trace=/^(.*_)?statv?fs regular expression.
                   %fstatfs     Trace fstatfs, fstatfs64, fstatvfs, osf_fstatfs, and osf_fstatfs64 system calls.  The same effect can be achieved with -e trace=/fstatv?fs regular expression.
                   %%statfs     Trace syscalls related to file system statistics (statfs-like, fstatfs-like, and ustat).  The same effect can be achieved with -e trace=/statv?fs|fsstat|ustat regular expression.

       -o filename
       --output=filename
                   Write the trace output to the file filename rather than to stderr.  filename.pid form is used if -ff option is supplied.  If the argument begins with '|' or '!', the rest of the argument is treated as a command and all output  is  piped       
                   to it.  This is convenient for piping the debugging output to a program without affecting the redirections of executed programs.  The latter is not compatible with -ff option currently.
```

``` sh
#examples:
strace -f -e trace=file -o /tmp/login_interactive /bin/bash -l
```

---
