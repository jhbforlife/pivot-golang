# Assignment: Writing Tests

## Objective

Create a calculator library and write the tests that ensure all of its operations account for failure cases.

## Learning Goals

1. Create a library that you import into an executable (all within your homework repository)
2. Use the standard library’s `testing` package
3. Write tests from the “outside in”

## Acceptance Criteria

1. You have a package named `calculator` that exposes functions with the following signatures
    1. `Add(a, b int) int`
    2. `Subtract(a, b int) int`
    3. `Multiply(a, b int) int`
    4. `Divide(a, b int) (int, error)` — Note the special case for this function, handle “divide by zero” errors gracefully by returning an error
2. You have an executable that uses your package (i.e. it imports your package `github.com/<your-username>/<your-pivottechschool-repo>/calculator`)
3. You have a `calculator_test` package (inside a file `caltulator_test.go` to go along with your `calculator.go` file)
4. You make use of the “Table-Driven Testing” testing technique to test multiple scenarios
5. You use subtests to clearly delineate which scenarios are being tested
6. Provide a link to a GitHub pull request (PR) against your repository’s `main` branch with your solution for evaluation

## Resources

- [Recording from 10/31/22 class](https://www.notion.so/f20d8d35d0ed424ba30de523b31a81ea)

## Examples of Implementation
### [/product-server/writing-tests](https://github.com/jhbforlife/pivot-golang/writing-tests)  
run `go run main.go`

input   `2 + 2`  
output  `Result: 4`

input   `123 - 124`  
output  `Result: -1`

input   `1000*1000`  
output  `Result: 1000000`

input   `90 / 10`  
output  `Result: 9`

input   `1/0`  
output  `calculator: cannot divide by 0, please try again`

### [/product-server/writing-tests/calculator](https://github.com/jhbforlife/pivot-golang/writing-tests/calculator)  
run `go test -v`  
output  
```=== RUN   TestCalculator
=== RUN   TestCalculator/123+321
=== RUN   TestCalculator/900+99
=== RUN   TestCalculator/2+3
=== RUN   TestCalculator/321-123
=== RUN   TestCalculator/24-6
=== RUN   TestCalculator/1-2
=== RUN   TestCalculator/10*100
=== RUN   TestCalculator/42*0
=== RUN   TestCalculator/31*2
=== RUN   TestCalculator/121/11
=== RUN   TestCalculator/81/9
=== RUN   TestCalculator/100/0
--- PASS: TestCalculator (0.00s)
    --- PASS: TestCalculator/123+321 (0.00s)
    --- PASS: TestCalculator/900+99 (0.00s)
    --- PASS: TestCalculator/2+3 (0.00s)
    --- PASS: TestCalculator/321-123 (0.00s)
    --- PASS: TestCalculator/24-6 (0.00s)
    --- PASS: TestCalculator/1-2 (0.00s)
    --- PASS: TestCalculator/10*100 (0.00s)
    --- PASS: TestCalculator/42*0 (0.00s)
    --- PASS: TestCalculator/31*2 (0.00s)
    --- PASS: TestCalculator/121/11 (0.00s)
    --- PASS: TestCalculator/81/9 (0.00s)
    --- PASS: TestCalculator/100/0 (0.00s)
PASS
ok  	github.com/jhbforlife/pivot-golang/writing-tests/calculator	0.257s```