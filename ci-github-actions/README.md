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

# Acceptance Criteria - Part 2

1. Create a new branch off of your `add-ci` branch called `add-pow` for what comes next
2. Add a new `func Pow(x, y float64) float64` to your package that returns `x**y`
    1. For example, calling on your `calculator.Pow(2, 3)` should return the value `8`.
3. You may use the `math` package inside of your function to calculate and return the result.
4. You must have a `TestPow` function in your test suite to test your new function.
5. Open a Pull Request with passing Checks (they’ll be passing if your tests pass during your workflow execution)

# Resources

1. [GitHub Actions Documentation](https://docs.github.com/en/actions)
2. [Building and testing Go](https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go)
3. [Nov 21, 2022 class recording](https://www.notion.so/MONDAY-158d3ddfe6e64e949733d153f8318b62)