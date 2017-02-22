# easygit

This is a set of helper functions for git2go to make the API more accessible.

## Usage

```go
import "github.com/runemadsen/easygit"

func main() {

  // Add all files to index. Similar to 'git add .'
  err := easygit.AddAll("path/to/repo")

  // Commit files in the index. If repo is commit, it will also create HEAD
  err := easygit.Commit("path/to/repo", "My commit message", "Name", "Email")

  // List all local branches. Similar to 'git branch'
  branchNames := easygit.ListBranches("path/to/repo")

  // Get the current local branch
  currentBranch := easyGit.CurrentBranch("path/to/repo")

  // Deletes a branch
  err := easygit.DeleteBranch("path/to/repo", "mybranch")

  // Creates a branch from another branch.
  err := easygit.CreateBranch("path/to/repo", "master", "newbranch")

  // Pushes a branch to a HTTPS remote. Similar to 'git push origin master'
  err := easygit.PushBranch("path/to/repo", "origin", "master", "user", "password")
  fmt.Println(err.(*git.GitError).Code) // if -7, it was wrong credentials

  // Checks out a branch. Similar to 'git checkout slave'
  err := easygit.CheckoutBranch("path/to/repo", "mybranch")
}
```
