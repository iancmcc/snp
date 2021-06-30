# snp: command-line Snappy

```
Usage of ./snp:
  -b    Output compressed data as binary instead of a base64-encoded string
  -d    Decode from stdin
```

With no arguments, `snp` compresses stdin using Snappy and outputs
a base64-encoded result. When `-b` is passed, it will output the compressed
data as binary.

When `-d` is passed, `snp` will decode the input, automatically decoding base64
first if detected.

## Example

```
# Note base64 output by default

$ echo "Hi There" | ./snp
CSBIaSBUaGVyZQo=

# Note automatic base64 decoding

$ echo "Hi There" | ./snp | ./snp -d
Hi There

# Also decodes when it's straight binary

$ echo "Hi There" | ./snp -b | ./snp -d
Hi There
```
