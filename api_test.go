package git2gobindings

import (
	"fmt"
  "testing"
  "os"
)

func TestListBranches(t *testing.T) {
  ListBranches()
}

func TestMain(m *testing.M) {
  fmt.Println("setup repo")
	os.Exit(m.Run())
}
