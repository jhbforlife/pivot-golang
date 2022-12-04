# Objective

Add a GitHub Actions Workflow to test your `calculator` package as you add new functionality.

# Learning Goals

1. Understand Continuous Integration using GitHub Actions (GHA)
2. Learn how to write simple GHA workflows for your Go packages
3. Experience the CI workflow by extending your existing package with new behavior that is tested during your PR

# Acceptance Criteria - Part 1

1. Create a new branch off of your `main` branch called `add-ci` where you will do your work
2. Add a “Default” workflow to your repository designed to test your Go code (see [example workflow](https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go))
3. Your workflow must test your existing `calculator` package **as is** before you add new behavior to the package
4. Upon adding your workflow, verify that your `calculator` package’s tests are successfully executed (the “Actions” tab in your repo must have at least one successful workflow run)
5. Open a Pull Request with passing Checks (they’ll be passing if your tests pass during your workflow execution)