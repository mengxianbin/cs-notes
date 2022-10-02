[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Commands](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands) /
[sudo](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands/sudo)

> sudo, sudoedit — execute a command as another user

```man
     -i [command]
                 The -i (simulate initial login) option runs the shell specified by the password database
                 entry of the target user as a login shell.  This means that login-specific resource files
                 such as .profile or .login will be read by the shell.  If a command is specified, it is
                 passed to the shell for execution via the shell’s -c option.  If no command is specified,
                 an interactive shell is executed.  sudo attempts to change to that user’s home directory
                 before running the shell.  The security policy shall initialize the environment to a mini-
                 mal set of variables, similar to what is present when a user logs in.  The Command
                 Environment section in the sudoers(5) manual documents how the -i option affects the envi-
                 ronment in which a command is run when the sudoers policy is in use.
```

```man
     -U user     The -U (other user) option is used in conjunction with the -l
                 option to specify the user whose privileges should be listed.
                 The security policy may restrict listing other users’ privi-
                 leges.  The sudoers policy only allows root or a user with
                 the ALL privilege on the current host to use this option.

     -u user     The -u (user) option causes sudo to run the specified command
                 as a user other than root.  To specify a uid instead of a
                 user name, #uid.  When running commands as a uid, many shells
                 require that the ‘#’ be escaped with a backslash (‘\’).
                 Security policies may restrict uids to those listed in the
                 password database.  The sudoers policy allows uids that are
                 not in the password database as long as the targetpw option
                 is not set.  Other security policies may not support this.
```

```man
     -l[l] [command]
                 If no command is specified, the -l (list) option will list
                 the allowed (and forbidden) commands for the invoking user
                 (or the user specified by the -U option) on the current host.
                 If a command is specified and is permitted by the security
                 policy, the fully-qualified path to the command is displayed
                 along with any command line arguments.  If command is speci-
                 fied but not allowed, sudo will exit with a status value of
                 1.  If the -l option is specified with an l argument (i.e.
                 -ll), or if -l is specified multiple times, a longer list
                 format is used.
```

```man
     -S, --stdin
                 Write the prompt to the standard error and read the password from the standard input instead of using the terminal device.
```
