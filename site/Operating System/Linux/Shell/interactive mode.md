[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Shell](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Shell) /
[interactive mode](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Shell/interactive%20mode)

``` sh
[[ $- == *i* ]] && echo "Interactive shell" || echo "Non-interactive shell"
Interactive shell

bash -c ' [[ $- == *i* ]] && echo "Interactive shell" || echo "Non-interactive shell" '
Non-interactive shell

zsh -c ' [[ $- == *i* ]] && echo "Interactive shell" || echo "Non-interactive shell" '
Non-interactive shell
```

---
