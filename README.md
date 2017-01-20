# easygit

This is a set of helper functions for git2go to make the API more accessible.

## Usage

```go
import "github.com/runemadsen/easygit"

func main() {

  branchNames := easygit.ListBranches("path/to/repo")

  currentBranch := easyGit.CurrentBranch("path/to/repo")

  err := easygit.DeleteBranch("path/to/repo", "mybranch")

  err := easygit.CreateBranch("path/to/repo", "master", "slave")

  err := easygit.PushBranch("path/to/repo", "origin", "master", "user", "password")
}
```
