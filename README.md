# Get lerna changed

This library helps to get changed lerna bundle. Actually, it runs:

```bash
lerna changed
```
And gets changed project name
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