```man
OPTIONS
       -k --keep
              Keep (don't delete) input files during compression or decompression.
```


```man
OPTIONS
       -c --stdout --to-stdout
              Write  output  on standard output; keep original files unchanged.  If there are several input files, the output consists of a sequence of independently compressed members. To obtain better compression, concatenate all input files before com‐
              pressing them.
ADVANCED USAGE
       Multiple compressed files can be concatenated. In this case, gunzip will extract all members at once. For example:

             gzip -c file1  > foo.gz
             gzip -c file2 >> foo.gz

       Then

             gunzip -c foo

       is equivalent to

             cat file1 file2

       In case of damage to one member of a .gz file, other members can still be recovered (if the damaged member is removed). However, you can get better compression by compressing all members at once:

             cat file1 file2 | gzip > foo.gz

       compresses better than

             gzip -c file1 file2 > foo.gz

       If you want to recompress concatenated files to get better compression, do:

             gzip -cd old.gz | gzip > new.gz

       If a compressed file consists of several members, the uncompressed size and CRC reported by the --list option applies to the last member only. If you need the uncompressed size for all members, you can use:

             gzip -cd file.gz | wc -c

       If  you  wish to create a single archive file with multiple members so that members can later be extracted independently, use an archiver such as tar or zip. GNU tar supports the -z option to invoke gzip transparently. gzip is designed as a comple‐       
       ment to tar, not as a replacement.
```
