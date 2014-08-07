/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : git.go

* Purpose :

* Creation Date : 08-07-2014

* Last Modified : Thu 07 Aug 2014 08:44:13 PM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package ezgit

import (
	"fmt"
	"os/exec"
	"strings"
)

type GitRepo struct {
	path   string
	bin    string
	prefix string
}

func NewGitRepo(path string, bin string) *GitRepo {
	return &GitRepo{
		path:   path,
		bin:    bin,
		prefix: fmt.Sprintf("cd %s; %s", path, bin),
	}
}

func osexec(cmd string) (stdOut string, stdErr error) {
	out, stdErr := exec.Command("/bin/sh", "-c", cmd).Output()
	stdOut = strings.TrimSpace(strings.Trim(string(out), "\n"))
	return
}

func (git *GitRepo) Commit(comment string, files []string) error {
	var fs string
	for _, v := range files {
		fs += v + " "
	}
	cmd := fmt.Sprintf("%s commit -m '%s' %s", git.prefix, comment, fs)
	_, err := osexec(cmd)
	return err
}

func (git *GitRepo) Push() error {
	cmd := fmt.Sprintf("%s push", git.prefix)
	_, err := osexec(cmd)
	return err
}

func (git *GitRepo) Add(files []string) error {
	var fs string
	for _, v := range files {
		fs += v + " "
	}
	cmd := fmt.Sprintf("%s add %s", git.prefix, fs)
	_, err := osexec(cmd)
	return err
}