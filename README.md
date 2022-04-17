# github.com/pascal71/simplemath

## Usage

### Initialize your module

```
$ go mod init example.com/my-golib-demo
```

### Get the simplemath module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/pascal71/simplemath@v0.1.0
```

```go
package main

import (
    "fmt"

    "github.com/pascal71/simplemath"
)

func main() {
    fmt.Println(simplemath.Add(3,4))
}
```

## Testing

```
$ go test
```

## Tagging

```
$ git tag v0.1.0
$ git push origin --tags
```

