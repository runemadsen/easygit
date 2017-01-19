package easygit

import (
  "testing"
  "os"
  "io/ioutil"
  "os/exec"
  "fmt"
  "sort"
)

func TestListBranches(t *testing.T) {
  branches, err := ListBranches("testrepo")
  sort.Strings(branches)
  if branches[0] != "master" || branches[1] != "slave" || err != nil {
    fmt.Println(branches)
    fmt.Println("TestListBranches failed")
    t.Fail()
  }
}

func TestDeleteBranch(t *testing.T) {
  err := DeleteBranch("testrepo", "xxxdeleteme")
  if err != nil {
    fmt.Println("TestDeleteBranch failed")
    t.Fail()
  }
}

func TestCurrentBranch(t *testing.T) {
  branch, err := CurrentBranch("testrepo")
  fmt.Println(branch)
  if branch != "master" || err != nil  {
    fmt.Println("CurrentBranch failed")
    t.Fail()
  }
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
  exec.Command("git", "checkout", "-b", "xxxdeleteme").Output()
  exec.Command("git", "checkout", "master").Output()
  os.Chdir("../")
	os.Exit(m.Run())
}
