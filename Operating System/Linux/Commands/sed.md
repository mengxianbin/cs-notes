* -e script, --expression=script
* -n, --quiet, --silent
* -r, --regexp-extended
* -i[SUFFIX], --in-place[=SUFFIX]

---

```
       s/regexp/replacement/
              Attempt to match regexp against the pattern space.  If successful, replace that portion matched with replacement.  The replacement may contain the special char-     
              acter  &  to  refer to that portion of the pattern space which matched, and the special escapes \1 through \9 to refer to the corresponding matching sub-expres-     
              sions in the regexp.

       g G    Copy/append hold space to pattern space.              
       p      Print the current pattern space.       

       a \
       text   Append text, which has each embedded newline preceded by a backslash.

       i \
       text   Insert text, which has each embedded newline preceded by a backslash.

       c \
       text   Replace the selected lines with text, which has each embedded newline preceded by a backslash.

       d      Delete pattern space.  Start next cycle.
```

---
