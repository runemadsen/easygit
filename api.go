package easygit

import (
	"github.com/libgit2/git2go"
	"strings"
)

func ListBranches(repoPath string) []string {

	repo, _ := git.OpenRepository(repoPath)

	iter, _ := repo.NewReferenceIterator()
	nameIter := iter.Names()

	var branches []string

	name, err := nameIter.Next()
	for err == nil {
		split := strings.Split(name, "/")
		if(split[1] == "heads") {
			branches = append(branches, split[2])
		}
		name, err = nameIter.Next()
	}

	return branches
}
