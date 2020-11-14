# dcast (data cast)

Format conversions.

```shell
$ go get github.com/miku/dconv/...
```

## Convert JSON to YAML

First use case, a bit more human editable JSON.

```yaml
$ echo '{"a": ["b", "c"], "d": {"e": [1, 2, 3], "f": "g", "h": "i"}}' | dcast -y
a:
- b
- c
d:
  e:
  - 1
  - 2
  - 3
  f: g
  h: i

$ curl -sL fixtures/doc.json | dconv -y
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
