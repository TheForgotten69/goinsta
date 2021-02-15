# Contribution

Welcome programmer!

If you want to contribute to Goinsta API you must follow a simple instructions.

- **Test your code after making pull request**. The title says it all.
- **Include jokes if you can**. This instruction is optional.

## Tests

You need at least one goinsta exported object

```go
package main

import (
 "fmt"
 "github.com/TheForgotten69/goinsta/v2"
 "github.com/TheForgotten69/goinsta/v2/utilities"
)

func main() {
 inst := goinsta.New("user", "password")
 err := inst.Login()
 if err != nil {
  fmt.Fatal(err)
 }
 fmt.Print(utilities.ExportAsBase64String(inst))
}
```

Then you can use the output generated above to run your tests in the cli

```bash
INSTAGRAM_BASE64_USERNAME=BASE64_OUTPUT go test ./...
```
