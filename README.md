# math-skills

A Go program that reads a list of integers from a file and computes four statistical measures: **average**, **median**, **variance**, and **standard deviation**.

## Usage

```bash
go run . <path-to-file>
```

The file should contain integers separated by whitespace (spaces or newlines). Non-integer tokens are skipped with a warning on stderr.

### Error cases

| Situation | Output |
|---|---|
| No argument given | `Usage: go run . <file>` |
| File not found | OS error message |
| File empty or no valid numbers | `Error: no valid numbers found in file` |
| Non-integer token in file | Warning on stderr, token skipped |

### Example

Given a file `data.txt`:

```
189
113
121
121
118
```

Run:

```bash
go run . data.txt
```

Output:

```
Average: 132
Median: 121
Variance: 821
Standard Deviation: 29
```

## How it works

### 1. Reading and sorting (`SortList`)

Reads the file passed as the first argument, parses all whitespace-separated tokens as integers, and returns them sorted in ascending order along with the count. Exits with an error if no argument is given, the file cannot be read, or no valid integers are found. Non-integer tokens are skipped with a warning.

### 2. Average (`Average`)

Computes the **arithmetic mean** of all values, rounded to the nearest integer:

```
Average = round(sum / count)
```

### 3. Median (`Median`)

Finds the **middle value** of the sorted list:

- If the count is **odd**: the middle element.
- If the count is **even**: the mean of the two middle elements, rounded up if the sum is odd.

### 4. Variance (`Variance`)

Measures how spread out the values are from the average:

```
Variance = round(sum((x - average)²) / count)
```

### 5. Standard Deviation (`StanDev`)

The square root of the variance, rounded to the nearest integer:

```
Standard Deviation = round(sqrt(Variance))
```

## Project structure

```
math-skills/
├── main.go
└── functions/
    ├── SortList.go
    ├── Average.go
    ├── Median.go
    ├── Variance.go
    └── StanDev.go
```
