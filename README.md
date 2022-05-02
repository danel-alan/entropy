# Entropy

## build

```shell
// requires docker and go
$ make build

// or

$ go run ./cmd/<app>/main.go --help
```

## Example output

```json
{
    "entropy_detail": [
        7.50,
        4.42,
        4.64,
        4.42,
        4.6,
        4.24,
        4.35,
        0.14,
        0.16
    ],
    "summary": {
        "low_entropy_blocks": 2,
        "high_entropy_blocks": 1
    }
}
```

## entropy-cli
### Usage 

```shell
$ entropy-cli --help

usage: entropy [flags] [path ...]
  -high float
        threshold for counting blocks as high entropy (default 7)
  -low float
        threshold for counting blocks as low entropy (default 2)
  -size uint
        block size to analize (default 1024)
```

## entropy-rest

### Usage 

```shell
$ entropy-rest --help

usage: entropy-rest [flags]
  -def_size uint
        default size for a block (default 1024)
  -high float
        threshold for counting blocks as high entropy (default 7)
  -low float
        threshold for counting blocks as low entropy (default 2)
  -port int
        port of the application (default 8080)
```

### Curl example
```shell
curl --location --request POST 'http://<host:port>/api/entropy' \
--form 'block_size="<size>"' \
--form 'file=@"<path to file>"'
```

## entropy-html

### Usage 

```shell
$ entropy-rest --help

usage: entropy-html [flags]
  -def_size uint
        default size for a block (default 1024)
  -high float
        threshold for counting blocks as high entropy (default 7)
  -low float
        threshold for counting blocks as low entropy (default 2)
  -port int
        port of the application (default 8080)
```
