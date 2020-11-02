``` sh
[[ $- == *i* ]] && echo "Interactive shell" || echo "Non-interactive shell"
Interactive shell

bash -c ' [[ $- == *i* ]] && echo "Interactive shell" || echo "Non-interactive shell" '
Non-interactive shell

zsh -c ' [[ $- == *i* ]] && echo "Interactive shell" || echo "Non-interactive shell" '
Non-interactive shell
```

---
