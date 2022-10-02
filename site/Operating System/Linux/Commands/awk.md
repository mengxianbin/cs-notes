[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Commands](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands) /
[awk](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands/awk)

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
    Fields
        The variable NF is set to the total number of fields in the input record.
```

```
    Built-in Variables

       FILENAME    The  name  of the current input file.  If no files are specified on the command line, the value of FILE‐
                  NAME is “-”.  However, FILENAME is undefined inside the BEGIN rule (unless set by getline).

       FS          The input field separator, a space by default.  See Fields, above.

       NF          The number of fields in the current input record.

       NR          The total number of input records seen so far.

       OFMT        The output format for numbers, "%.6g", by default.

       OFS         The output field separator, a space by default.

       ORS         The output record separator, by default a newline.

       RS          The input record separator, by default a newline.

       RT          The  record terminator.  Gawk sets RT to the input text that matched the character or regular expression
                   specified by RS.

```

```
   Octal and Hexadecimal Constants
       You may use C-style octal and hexadecimal constants in your AWK program source code.  For example, the  octal  value
       011 is equal to decimal 9, and the hexadecimal value 0x11 is equal to decimal 17.

   String Constants
       String  constants in AWK are sequences of characters enclosed between double quotes (like "value").  Within strings,
       certain escape sequences are recognized, as in C.  These are:

       \\   A literal backslash.

       \a   The “alert” character; usually the ASCII BEL character.

       \b   Backspace.

       \f   Form-feed.

       \n   Newline.

       \r   Carriage return.

       \t   Horizontal tab.

       \v   Vertical tab.

       \xhex digits
            The character represented by the string of hexadecimal digits following the \x.  Up to two following  hexadeci‐
            mal digits are considered part of the escape sequence.  E.g., "\x1B" is the ASCII ESC (escape) character.

       \ddd The  character  represented  by the 1-, 2-, or 3-digit sequence of octal digits.  E.g., "\033" is the ASCII ESC
            (escape) character.

       \c   The literal character c.

       In compatibility mode, the characters represented by octal and hexadecimal escape sequences  are  treated  literally
       when used in regular expression constants.  Thus, /a\52b/ is equivalent to /a\*b/.
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

```sh
# build-in variables

awk -F\' 'BEGIN{printf "%4s %4s %4s %4s %4s %4s %4s %4s %4s\n", "FILENAME","ARGC","FNR","FS","NF","NR","OFS","ORS","RS";
                printf "---------------------------------------------\n"}
                 {printf "%4s %4s %4s %4s %4s %4s %4s %4s %4s\n", FILENAME,ARGC,FNR,FS,NF,NR,OFS,ORS,RS}' log.txt

awk '{print NR,FNR,$1,$2,$3}' log.txt

awk '{print $1,$2,$5}' OFS=" $ "  log.txt
```

```sh
# regular expression

awk '$2 ~ /th/ {print $2,$4}' log.txt

awk '$2 !~ /th/ {print $2,$4}' log.txt
awk '!/th/ {print $2,$4}' log.txt
```

```sh
# case

awk 'BEGIN{IGNORECASE=1} /this/' log.txt
```

```sh
# awk script

# cat script.awk

BEGIN {}
{}
END {}
```

```sh
# file size
ls -l *.txt | awk '{sum+=$5} END {print sum}'
```

```sh
# filter line length
awk 'length>80' log.txt
```

```sh
seq 9 | sed 'H;g' | awk -v RS='' '{for(i=1;i<=NF;i++)printf("%dx%d=%d%s", i, NR, i*NR, i==NR?"\n":"\t")}'
```

---

> REF

- <https://www.runoob.com/linux/linux-comm-awk.html>
