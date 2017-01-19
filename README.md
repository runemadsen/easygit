# easygit

This is a set of helper functions for git2go to make the API more accessible.

## Usage

```go
import "github.com/runemadsen/easygit"

func main() {

  branchNames := easygit.ListBranches("path/to/repo")

  currentBranch := easyGit.CurrentBranch("path/to/repo")

  err := easygit.DeleteBranch("path/to/repo", "mybranch")
}
```
