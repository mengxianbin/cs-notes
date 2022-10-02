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
