# rinha-compiler-go

Interpretador em Golang feito para [rinha de compiler](https://github.com/aripiprazole/rinha-de-compiler).

## Features

- [x] Int, Str, Bool
- [x] Binary
- [x] Let
- [x] Function
- [x] If
- [x] Call
- [x] Print
- [x] First
- [x] Second
- [x] Tuple


## Build

```
docker build -t rinha .
```

## Run

```
docker run -v {json_ast_filename}:/var/rinha/source.rinha.json rinha
```

