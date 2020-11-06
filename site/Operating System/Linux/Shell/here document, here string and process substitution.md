[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Shell](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Shell) /
[here document, here string and process substitution](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Shell/here%20document%2C%20here%20string%20and%20process%20substitution)

## here document

```sh
wc << EOF
> EOF
```

## here string

```sh
wc <<< "abc"
```

## process substitution

```sh
wc < <(echo abc)
```
