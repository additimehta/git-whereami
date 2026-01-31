package git

import (
	"bytes"
	"os/exec"
	"strings"
)

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return strings.TrimSpace(out.String()), err
}


func CurrentBranch() (string, error) {
	return runGit("branch", "--show-current")
}

func MergeBase(base string) (string, error) {
	return runGit("merge-base", base, "HEAD")
}

func AheadBehind(base string) (ahead, behind string, err error) {
	out, err := runGit("rev-list", "--left-right", "--count", base+"...HEAD")
	if err != nil {
		return "", "", err
	}
	parts := strings.Split(out, "\t")
	return parts[1], parts[0], nil
}