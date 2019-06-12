# If JSON has "dependency"

This library returns if JSON has certain dependency.
## Installation

Use the "go get".

```bash
go get github.com/malikov0216/if-json-has
```

## Usage

```go
import (
   "fmt"
   hasJSON "github.com/malikov0216/if-json-has"
)

func main () {
   if hasDep := hasJSON.FileRead("source to json", "project name"); hasDep {
   	//do-something-here
   }
}
```

## License
[MIT](https://choosealicense.com/licenses/mit/)