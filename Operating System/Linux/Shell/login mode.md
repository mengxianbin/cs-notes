> bash

``` sh
➜  ~ bash -l -c 'shopt | grep login'
login_shell     on
➜  ~ bash -c 'shopt | grep login'
login_shell     off
```

``` sh
➜  ~ bash -l -c 'shopt login_shell; echo $?'
login_shell     on
0
➜  ~ bash -c 'shopt login_shell; echo $?'
login_shell     off
1
```

``` sh
bash
shopt -q login_shell && echo 'Login shell' || echo 'Non-login shell'
Non-login shell

bash -l # --login
shopt -q login_shell && echo 'Login shell' || echo 'Non-login shell'
Login shell
```

---

> zsh

``` sh
➜  ~ zsh -c ' [[ -o login ]] && echo "Login shell" || echo "Not login shell" '
Not login shell
➜  ~ zsh -l -c ' [[ -o login ]] && echo "Login shell" || echo "Not login shell" '
Login shell
```

```sh
➜  .m2 zsh -c ' setopt | grep -q login; echo $? '
1
➜  .m2 zsh -l -c ' setopt | grep -q login; echo $? '
0
```

---
