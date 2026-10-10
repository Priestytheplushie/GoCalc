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
GoCalc supports basic scientific and four function functionality

### Basic Operations
GoCalc supports binary operations, such as 5+3, 2^3, and 10 % 3, including the operators +, -, *, / ^, %, as well as Unary operations, such as sqrt, abs, sin, including sqrt, abs, sin, cos, tan. Type `exit` or `quit` to leave. It also now supports `ans` and remembers the previous result. It also uses ANSI escape codes cuz why not?

```bash 
# Example Output
PS C:\Users\Priesty\Downloads\GoCalc> go run main.go         
Welcome to GoCalc
-----------------------------------------
Binary ops (e.g., 5 + 3, 2 ^ 3, 10 % 3): +, -, *, /, ^, %
Unary ops  (e.g., sqrt 16, abs -5, sin 90): sqrt, abs, sin, cos, tan
Memory:     Type 'ans' to use your last result
Type 'exit' to quit

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

<GoCalc> 
```

## Development 
This project was again, made for a school weekened project, and as such will not be maintained, it's also pretty lame, I orignally wanted to make a Go PumpkinMC plugin, but the Go Docs were outdated, and I hit strange compiler errors, so I made this thing instead. Have fun ig