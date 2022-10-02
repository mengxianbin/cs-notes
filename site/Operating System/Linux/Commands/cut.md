[Home](https://mengxianbin.github.io) /
[cs-notes](https://mengxianbin.github.io/cs-notes/site) /
[Operating System](https://mengxianbin.github.io/cs-notes/site/Operating%20System) /
[Linux](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux) /
[Commands](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands) /
[cut](https://mengxianbin.github.io/cs-notes/site/Operating%20System/Linux/Commands/cut)

```
NAME
       cut - remove sections from each line of files
```

```
       -b, --bytes=LIST
              select only these bytes

       -c, --characters=LIST
              select only these characters

       -d, --delimiter=DELIM
              use DELIM instead of TAB for field delimiter

       -f, --fields=LIST
              select  only these fields;  also print any line that contains no delimiter character, unless the -s option is    
              specified

       --complement
              complement the set of selected bytes, characters or fields
```

```
       Use one, and only one of -b, -c or -f.  Each LIST is made up of one range, or many ranges separated by commas.   Se‐    
       lected input is written in the same order that it is read, and is written exactly once.  Each range is one of:

       N      N'th byte, character or field, counted from 1

       N-     from N'th byte, character or field, to end of line

       N-M    from N'th to M'th (included) byte, character or field

       -M     from first to M'th (included) byte, character or field
```
