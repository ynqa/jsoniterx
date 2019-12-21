# jsoniterx

 [Extension](https://github.com/json-iterator/go/blob/dc11f49689fd1c9a6de20749def70bd889bf0d42/reflect_extension.go#L46-L56) for [json-iterator/go](https://github.com/json-iterator/go).

## Tags

### `time.Time`

|Name|description|
|--:|--:|
|format|A string for format. [more details](https://golang.org/pkg/time/#pkg-constants).|
|location|A string for [Location](https://golang.org/pkg/time/#Location).|

## Example

```go
package main

import (
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/ynqa/jsoniterx"
)

type Example struct {
	Time time.Time `json:"time" format:"2006-01-02 15:04:05" location:"Asia/Tokyo"`
}

func main() {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.RegisterExtension(jsoniterx.TimePlugin())

	var e Example
	str := `{"time": "2019-01-01 12:00:00"}`
	if err := json.Unmarshal([]byte(str), &e); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(e.Time)
}
```
