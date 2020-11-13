# dconv

Format conversions.

```shell
$ go get github.com/miku/dconv/...
```

## Convert JSON to YAML

First use case, a bit more human editable JSON.

```yaml
$ curl -sL https://git.io/JkOEo | dconv -y
abstracts:
- content: Belgium Herbarium image of Meise Botanic Garden.
  lang: de
  mimetype: text/plain
  sha1: cd3c76f5fd94bcf260f9ad74f797d9e79a824b1d
contribs:
- index: 0
  raw_name: Meise Botanic Garden
  role: author
...
```
