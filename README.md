> Work in progress
# git-whereami

A CLI tool written in Go that tells you **where you are in your Git project**.

It shows:
- your current branch
- your base branch (usually `origin/main`)
- how many commits you're ahead/behind
- a visual commit graph of your work

---

## Example

```bash
$ git-whereami

Base branch   : origin/main
Current branch: feature-login
Ahead         : 2
Behind        : 1

You are here:

* 9f3a2c1 (HEAD -> feature-login) fix login bug
* 7ab12d4 add validation