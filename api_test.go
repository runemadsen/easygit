package git2gobindings

import (
  "testing"
  "os"
  "io/ioutil"
  "os/exec"
)

func TestListBranches(t *testing.T) {
  ListBranches()
}

func TestMain(m *testing.M) {
  os.RemoveAll("testrepo")
  os.Mkdir("testrepo", os.ModePerm)
  os.Chdir("testrepo")
  ioutil.WriteFile("first.txt", []byte("first"), os.ModePerm)
  exec.Command("git", "init").Output()
  exec.Command("git", "add", ".").Output()
  exec.Command("git", "commit", "-m", "first commit").Output()
  exec.Command("git", "checkout", "-b", "slave").Output()
  exec.Command("git", "checkout", "master").Output()
  os.Chdir("../")
	os.Exit(m.Run())
}
