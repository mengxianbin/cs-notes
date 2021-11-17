```
NAME
       sed - stream editor for filtering and transforming text

SYNOPSIS
       sed [OPTION]... {script-only-if-no-other-script} [input-file]...

DESCRIPTION
       Sed  is a stream editor.  A stream editor is used to perform basic text transformations on an input stream (a file or input from a pipeline).  While
       in some ways similar to an editor which permits scripted edits (such as ed), sed works by making only one pass over  the  input(s),  and  is  conse‐
       quently more efficient.  But it is sed's ability to filter text in a pipeline which particularly distinguishes it from other types of editors.  
```

---

```
       -n, --quiet, --silent

              suppress automatic printing of pattern space

       -e script, --expression=script

              add the script to the commands to be executed

       -f script-file, --file=script-file

              add the contents of script-file to the commands to be executed

       -i[SUFFIX], --in-place[=SUFFIX]

              edit files in place (makes backup if SUFFIX supplied)

       -E, -r, --regexp-extended

              use extended regular expressions in the script (for portability use POSIX -E).

       -s, --separate

              consider files as separate rather than as a single, continuous long stream.
```

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
