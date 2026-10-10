# GoCalc
A simple CLI calculator tool because why not? 

## QuickStart
Run the code using the following: 
```powershell
git clone https://github.com/Priestytheplushie/GoCalc.git 
cd GoCalc 
go run main.go
``` 

## Usage

GoCalc has four-function and basic scientific operations, plus statistics, percentages, and probability. Type `exit` or `quit` to leave.

### Basic Operations

**Binary operations** take the form `[number] [operator] [number]`:

| Operator | Description | Example |
|----------|-------------|---------|
| `+` `-` `*` `/` | Basic arithmetic | `5 + 3` |
| `^` | Exponent | `2 ^ 3` |
| `%` | Modulo | `10 % 3` |

**Unary operations** take the form `[function] [number]`:

| Function | Description | Example |
|----------|-------------|---------|
| `sqrt` | Square root | `sqrt 16` |
| `abs` | Absolute value | `abs -5` |
| `sin`, `cos`, `tan` | Trigonometry (input in **degrees**) | `sin 90` |

**Memory:** use `ans` anywhere a number is expected to reuse your last result (e.g., `ans / 2`).

Example Usage:

```text
<GoCalc> 6 + 6
Result: 12

<GoCalc> ans / 2
Result: 6

<GoCalc> sqrt ans
Result: 2.449489742783178

<GoCalc> sin ans
Result: 0.042738639294777336

<GoCalc> ans ^ 3
Result: 7.806602622968441e-05
```

### Statistics & Probability

**Statistics** commands accept a list of numbers separated by spaces or commas:

| Command | Description | Example |
|---------|-------------|---------|
| `stats` | Count, sum, mean, median, min, max | `stats 10, 20, 30, 40` |
| `mean` | Average | `mean 4 8 15` |
| `median` | Middle value | `median 3 1 2` |
| `sum` | Total | `sum 1 2 3` |
| `min` / `max` | Smallest / largest value | `max 7 2 9` |

**Probability and combinatorics:**

| Command | Description | Example |
|---------|-------------|---------|
| `fact [n]` | Factorial (n up to 170) | `fact 5` → `120` |
| `ncr [n] [r]` | Combinations (n choose r) | `ncr 5 2` → `10` |
| `npr [n] [r]` | Permutations | `npr 5 2` → `20` |
| `prob [fav] [total]` | Chance of success as a percentage | `prob 3 10` → `30%` |

**Percentages and tips:**

| Command | Description | Example |
|---------|-------------|---------|
| `pct [num] [percent]` | Percent of a number | `pct 200 15` → `30` |
| `tip [bill] [tip%]` | Tip amount and total | `tip 50 20` → `Tip: $10 \| Total: $60` |

```text
<GoCalc> stats 10, 20, 30, 40
--- Statistical Summary ---
Count:  4
Sum:    100
Mean:   25
Median: 25
Min:    10
Max:    40

<GoCalc> ncr 5 2
Result (NCR): 10

<GoCalc> prob 3 10
Probability: 30% (Odds: 3 out of 10)
```

> [!WARNING]
> Statistics and probability features are still lightly tested, so double-check important results, and by lightly tested, I mean I just ran 1 test for the README.md, so be careful before using

## Development 
This project was again, made for a school weekened project, and as such will not be maintained, it's also pretty lame, I orignally wanted to make a Go PumpkinMC plugin, but the Go Docs were outdated, and I hit strange compiler errors, so I made this thing instead. Have fun ig