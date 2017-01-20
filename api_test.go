package easygit

import (
	"io/ioutil"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/libgit2/git2go"
)

func TestListBranches(t *testing.T) {

	repo := createTestRepo(t)
	seedTestRepo(t, repo)
	defer cleanupTestRepo(t, repo)

	branches, err := ListBranches(repo.Workdir())
	checkFatal(t, err)
	if branches[0] != "master" {
		fail(t)
	}
}

func TestCreateBranch(t *testing.T) {

	repo := createTestRepo(t)
	seedTestRepo(t, repo)
	defer cleanupTestRepo(t, repo)

	CreateBranch(repo.Workdir(), "master", "slave")

	branches, err := ListBranches(repo.Workdir())
	checkFatal(t, err)
	if branches[0] != "master" || branches[1] != "slave" {
		fail(t)
	}

}

func TestDeleteBranch(t *testing.T) {

	repo := createTestRepo(t)
	seedTestRepo(t, repo)
	defer cleanupTestRepo(t, repo)

	CreateBranch(repo.Workdir(), "master", "slave")

	err := DeleteBranch(repo.Workdir(), "slave")
	checkFatal(t, err)

	branches, err := ListBranches(repo.Workdir())
	checkFatal(t, err)
	if len(branches) != 1 {
		fail(t)
	}

}

func TestCurrentBranch(t *testing.T) {

	repo := createTestRepo(t)
	seedTestRepo(t, repo)
	defer cleanupTestRepo(t, repo)

	branch, err := CurrentBranch(repo.Workdir())
	checkFatal(t, err)
	if branch != "master" {
		fail(t)
	}
}

// Test setup
// ---------------------------------------------------

func createTestRepo(t *testing.T) *git.Repository {
	path, err := ioutil.TempDir("", "easygit")
	checkFatal(t, err)
	repo, err := git.InitRepository(path, false)
	checkFatal(t, err)

	tmpfile := "README"
	err = ioutil.WriteFile(path+"/"+tmpfile, []byte("foo\n"), 0644)
	checkFatal(t, err)
	return repo
}

func createBareTestRepo(t *testing.T) *git.Repository {
	path, err := ioutil.TempDir("", "git2go")
	checkFatal(t, err)
	repo, err := git.InitRepository(path, true)
	checkFatal(t, err)
	return repo
}

func cleanupTestRepo(t *testing.T, r *git.Repository) {
	var err error
	if r.IsBare() {
		err = os.RemoveAll(r.Path())
	} else {
		err = os.RemoveAll(r.Workdir())
	}
	checkFatal(t, err)
	r.Free()
}

func seedTestRepo(t *testing.T, repo *git.Repository) (*git.Oid, *git.Oid) {
	loc, err := time.LoadLocation("Europe/Berlin")
	checkFatal(t, err)
	sig := &git.Signature{
		Name:  "Rand Om Hacker",
		Email: "random@hacker.com",
		When:  time.Date(2013, 03, 06, 14, 30, 0, 0, loc),
	}

	idx, err := repo.Index()
	checkFatal(t, err)
	err = idx.AddByPath("README")
	checkFatal(t, err)
	err = idx.Write()
	checkFatal(t, err)
	treeID, err := idx.WriteTree()
	checkFatal(t, err)

	message := "This is a commit\n"
	tree, err := repo.LookupTree(treeID)
	checkFatal(t, err)
	commitID, err := repo.CreateCommit("HEAD", sig, sig, message, tree)
	checkFatal(t, err)

	return commitID, treeID
}

func fail(t *testing.T) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		t.Fatalf("Unable to get caller")
	}
	t.Fatalf("Fail at %v:%v; %v", file, line)
}

func checkFatal(t *testing.T, err error) {
	if err == nil {
		return
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		t.Fatalf("Unable to get caller")
	}
	t.Fatalf("Fail at %v:%v; %v", file, line, err)
}
