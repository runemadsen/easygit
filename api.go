package easygit

import (
	"github.com/libgit2/git2go"
	"strings"
)

func ListBranches(repoPath string) ([]string, error) {

	repo, repoErr := git.OpenRepository(repoPath)
	if repoErr != nil {
		return nil, repoErr
	}

	iter, iterErr := repo.NewReferenceIterator()
	if iterErr != nil {
		return nil, iterErr
	}

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

	return branches, nil
}
