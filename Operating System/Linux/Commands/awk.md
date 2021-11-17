```
GAWK(1)                                               Utility Commands                                              /

NAME
       gawk - pattern scanning and processing language

SYNOPSIS
       gawk [ POSIX or GNU style options ] -f program-file [ -- ] file ...
       gawk [ POSIX or GNU style options ] [ -- ] program-text file ...
```

```
       -f program-file
       --file program-file
              Read the AWK program source from the file program-file, instead of from  the  first  command  line  argument.    
              Multiple  -f  (or  --file)  options may be used.  Files read with -f are treated as if they begin with an im‐    
              plicit @namespace "awk" statement.

       -F fs
       --field-separator fs
              Use fs for the input field separator (the value of the FS predefined variable).

       -v var=val
       --assign var=val
              Assign the value val to the variable var, before execution of the program begins.  Such variable  values  are    
              available to the BEGIN rule of an AWK program.
```

```
The variable NF is set to the total number of fields in the input record.
```

```
AWK PROGRAM EXECUTION
       An AWK program consists of a sequence of optional directives, pattern-action statements, and optional function defi‐    
       nitions.

              @include "filename"
              @load "filename"
              @namespace "name"
              pattern   { action statements }
              function name(parameter list) { statements }
```

---

- '{print $1}'
- '{print $NF}'

---

```
awk '{[pattern] action}' {filenames}
```

```sh
# separator
awk -F

awk -F, '{print $1,$2}' inputfile
awk 'BEGIN{FS=","} {print $1,$2}' inputfile
awk -F '[ ,]'  '{print $1,$2,$5}' inputfile
```

```sh
# varialbe
awk -v

awk -va=1 -vb=s '{print $1,$1+a,$1b}' inputfile
```

```sh
# scriptfile
awk -f

awk -f scriptfile inputfile
```

```sh
# condition
awk '$1>2' log.txt
awk '$1==2 {print $1,$3}' log.txt
awk '$1>2 && $2=="Are" {print $1,$2,$3}' log.txt
```


---

> REF

- <https://www.runoob.com/linux/linux-comm-awk.html>
