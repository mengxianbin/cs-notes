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
