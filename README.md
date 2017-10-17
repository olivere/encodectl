# Encoder

This code is just an example of transforming between different charsets.

```
$ go build
$ cat testdata/garcon.windows-1252.txt
Gar�on !
$ cat testdata/garcon.windows-1252.txt | ./encodectl -ie=windows-1252 -oe=utf-8
Garçon !
$ ./encodectl -ie=windows-1252 -oe=utf-8 -i testdata/garcon.windows-1252.txt -o outputfile
$ cat outputfile
Garçon !
```

# License

MIT
