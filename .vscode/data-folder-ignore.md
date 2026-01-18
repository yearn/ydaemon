# Data Folder Git Configuration

## Problem

The `/data` folder contains files that are tracked in git and required for production deployments, but we don't want local changes to show up in git status during development.

## Solution

We use `git update-index --skip-worktree` to tell git to ignore local changes to these files without removing them from the repository.

## Setup (Per Developer)

Each developer needs to run this command once on their local machine:

```bash
git ls-files data/ | while read file; do git update-index --skip-worktree "$file"; done
```

This marks all files in the `data/` folder so that git will ignore local modifications.

## How It Works

- Files remain in the repository (deployments work normally)
- Local changes to these files won't appear in `git status`
- The setting is local to each developer's machine
- Does NOT propagate via git push/pull

## Undoing (If Needed)

To stop ignoring local changes and track them again:

```bash
git ls-files data/ | while read file; do git update-index --no-skip-worktree "$file"; done
```

## Checking Status

To see which files are marked with skip-worktree:

```bash
git ls-files -v | grep '^S'
```

Files marked with `S` are being skipped.

## When This Gets Disabled

The `--skip-worktree` flag may be cleared by git in these situations:

- Merge conflicts involving these files
- Some git operations that affect file status

If you notice changes in the `data/` folder appearing in git status again, simply re-run the setup command above.

## Why Not .gitignore?

`.gitignore` only works for untracked files. Since the `data/` folder files are already tracked and needed for production, we can't use `.gitignore`.
