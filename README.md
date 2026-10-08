# git-whereami

A small CLI tool written in Go to help you quickly understand where you are in a Git repository.

Instead of running several Git commands, see:
- your current branch
- your base branch
- how many commits you're ahead or behind
- a terminal commit graph of work on your branch that isn't in the base branch

## Install

```bash
brew install git-whereami
```

## Example

Illustrative output from a feature branch:

```text
$ git-whereami

Base branch   : main
Current branch: feature-login
Ahead         : 6
Behind        : 2

You are here:
* 9f3a2c1 (HEAD -> feature-login) Fix login bug
*   8e4b1c2 Merge feature-validation into feature-login
|\
| * 7ab12d4 Add input validation
| * 6b4d921 Add validation tests
* | 5c1e203 Create login endpoint
|/
* 2c1e4f0 Create login form
```

In this example, the feature branch has six commits that aren't on `main`, including a merge commit. The `*`, `|`, `/`, and `\` characters show how those commits connect.

The graph comes from Git's `log --graph --oneline --decorate` command. Currently, `git-whereami` only draws commits reachable from your current branch but not from the base branch. It does **not** draw the two commits you're behind on `main`.

## What I'd like to improve

- Show where the current branch diverged from the base branch.
- Explain what “ahead” and “behind” mean in plain language.
- Offer a more complete graph when you want to see branches and merges together.
