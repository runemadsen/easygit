package easygit

import (
	"strings"

	"github.com/libgit2/git2go"
)

func ListBranches(repoPath string) ([]string, error) {

	repo, err := git.OpenRepository(repoPath)
	if err != nil {
		return nil, err
	}

	iter, err := repo.NewBranchIterator(git.BranchLocal)
	if err != nil {
		return nil, err
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

func PushBranch(repoPath string, remoteName string, branch string, user string, pass string) error {

	repo, err := git.OpenRepository(repoPath)
	if err != nil {
		return nil
	}

	remote, err := repo.Remotes.Lookup(remoteName)
	if err != nil {
		return nil
	}

	err = remote.Push([]string{"refs/heads/" + branch}, &git.PushOptions{
		RemoteCallbacks: git.RemoteCallbacks{
			CredentialsCallback: func(url string, username_from_url string, allowed_types git.CredType) (git.ErrorCode, *git.Cred) {
				_, creds := git.NewCredUserpassPlaintext(user, pass)
				return git.ErrOk, &creds
			},
		},
	})

	return err
}

func CreateBranch(repoPath string, from string, to string) error {

	repo, err := git.OpenRepository(repoPath)
	if err != nil {
		return err
	}

	fromBranch, err := repo.LookupBranch(from, git.BranchLocal)
	if err != nil {
		return err
	}

	fromCommit, err := repo.LookupCommit(fromBranch.Target())
	if err != nil {
		return err
	}

	_, err = repo.CreateBranch(to, fromCommit, false)
	if err != nil {
		return err
	}

	return nil
}

// CheckoutBranch

// PushBranch

func CurrentBranch(repoPath string) (string, error) {

	repo, repoErr := git.OpenRepository(repoPath)
	if repoErr != nil {
		return "", repoErr
	}

	head, headErr := repo.Head()
	if repoErr != nil {
		return "", headErr
	}

	return strings.Split(head.Name(), "/")[2], nil
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
