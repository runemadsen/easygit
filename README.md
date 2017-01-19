# easygit

This is a set of helper functions for git2go to make the API more accessible.

## Usage

```go
import "github.com/runemadsen/easygit"

func main() {
  branches := easygit.ListBranches("path/to/repo")
  // Returns []string{ "master", "slave" }
}
```
