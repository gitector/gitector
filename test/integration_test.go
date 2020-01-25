package test

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"testing"
)

var binaryName = "gitector"
var scenarios = packr.NewBox("./scenarios")

func TestCanCompareBranch(t *testing.T) {
	tmpDir, _ := CreateTempDir()
	binaryPath := tmpDir + "/" + binaryName
	buildBinary(binaryPath)
	initGitFakeRepo(tmpDir)
	updateConfig(tmpDir)
	createCommit("this is too long message written inside commit written fo contain more than 72 characters master")
	createBranch("new-feature")
	createCommit("this is too long message written inside commit written fo contain more than 72 characters new-feature1")
	createCommit("this is too long message written inside commit written fo contain more than 72 characters new-feature2")
	createCommit("this is too long message written inside commit written fo contain more than 72 characters new-feature3")
	res, _ := exec.Command(binaryPath).Output()
	textRes := string(res[:])
	RemoveAllFilesFromDir(tmpDir)

	expectedOutput, _ := scenarios.FindString("scenario1.txt")
	assert.Equal(t, expectedOutput, textRes)
}

func initGitFakeRepo(path string) {
	printOutput(exec.Command("git", "init", path).Output())
	fmt.Println(os.Chdir(path))
}

func CreateTempDir() (string, error) {
	return ioutil.TempDir("", "gitector_test")
}

func createBranch(branchName string) {
	printOutput(exec.Command("git", "checkout", "-b", branchName).Output())
}

func switchBranch(branchName string) {
	printOutput(exec.Command("git", "checkout", branchName).Output())
}

func createCommit(message string) {
	filename := randName(10)
	printOutput(exec.Command("touch", filename).Output())
	printOutput(exec.Command("git", "add", filename).Output())
	printOutput(exec.Command("git", "commit", "-m", message).Output())
}

func updateConfig(path string) {
	newConfig := []byte(`
{
  "allowed_domains": [],
  "title_max_characters": 72,
  "ticket_regexp": "\\[\\d\\]"
}
`)
	ioutil.WriteFile(path + "/.gitector.json", newConfig, 0644)
}

func RemoveAllFilesFromDir(dir string) {
	_ = os.RemoveAll(dir)
}

func printOutput(data []byte, err error) {
	if err != nil {
		fmt.Printf("%s", data)
		fmt.Println(err)
	}
}

func randName(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func buildBinary(path string) {
	err := os.Chdir("..")
	if err != nil {
		fmt.Printf("could not switch dir %s: %v", binaryName, err.Error())
		os.Exit(1)
	}

	getDep := exec.Command("go", "build", "-o", path)
	if err := getDep.Run(); err != nil {
		fmt.Printf("could not make binary for %s: %v", binaryName, err.Error())
		os.Exit(1)
	}
}
