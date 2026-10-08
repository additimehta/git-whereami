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
Ahead         : 3
Behind        : 2

You are here:
* 9f3a2c1 (HEAD -> feature-login) Fix login bug
* 7ab12d4 Add validation
* 2c1e4f0 Create login form
```

The graph uses Git's `log --graph --oneline --decorate` output. **Currently it shows only commits reachable from your branch but not from the base branch**, not the full history of every branch. The ahead/behind counts can therefore include commits that aren't drawn in this graph.

## What I'd like to improve

- Show where the current branch diverged from the base branch.
- Explain what “ahead” and “behind” mean in plain language.
- Offer a more complete graph when you want to see branches and merges together.
