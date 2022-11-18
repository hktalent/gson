# Overview

[![Go Reference](https://pkg.go.dev/badge/github.com/hktalent/gson.svg)](https://pkg.go.dev/github.com/hktalent/gson)

The tests is the doc.

A tiny JSON lib to read and alter a JSON value. The data structure is lazy, it's parse-on-read so that you can replace the parser with a faster one if performance is critical, use method `JSON.Raw` to do it.

# New features
- use github.com/json-iterator/go Improve performance

# Example
```go
obj := gson.NewFrom(`{"a": {"b": [1, 2]}}`)

fmt.Println(obj.Get("a.b.0").Int())

obj.Set("a.b.1", "ok").Set("c", 2)
obj.Del("c")
fmt.Println(">", obj.JSON("> ", "  "))

```