# Git Workflow Guide

This project follows a feature-branch workflow.

## 1. Branches

-   **`main`**: The stable, production-ready code.
-   **`develop`** (Optional): Integration branch for testing before merging to main.
-   **`feature/*`**: Create a new branch for every new feature or fix.

## 2. How to Work

### Starting a New Feature
Always create a new branch from `main` (or `develop`):

```bash
git checkout main
git pull origin main
git checkout -b feature/my-new-feature
```

### Making Changes
1.  Write your code.
2.  Stage changes: `git add .`
3.  Commit: `git commit -m "Add new feature"`

### Saving to GitHub
Push your feature branch to GitHub:

```bash
git push -u origin feature/my-new-feature
```

### Merging
1.  Go to GitHub.
2.  Open a **Pull Request (PR)** from `feature/my-new-feature` to `main`.
3.  Review and merge.
