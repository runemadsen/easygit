package easygit

import (
	"github.com/libgit2/git2go"
)

func ListBranches(repoPath string) ([]string, error) {

	repo, repoErr := git.OpenRepository(repoPath)
	if repoErr != nil {
		return nil, repoErr
	}

	iter, iterErr := repo.NewBranchIterator(git.BranchLocal)
	if iterErr != nil {
		return nil, iterErr
	}

	var branches []string

	branch, _, err := iter.Next()
	for err == nil {
		name, _ := branch.Name()
		branches = append(branches, name)
		branch, _, err = iter.Next()
	}

	return branches, nil
}

func DeleteBranch(repoPath string, branchName string) error {

	repo, repoErr := git.OpenRepository(repoPath)
	if repoErr != nil {
		return repoErr
	}

	branch, branchErr := repo.LookupBranch(branchName, git.BranchLocal)
	if branchErr != nil {
		return branchErr
	}
	return branch.Delete()
}
