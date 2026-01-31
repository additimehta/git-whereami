package git

import (
	"bytes"
	"os/exec"
	"strings"
)

func runGit(args ...string) (string, error)