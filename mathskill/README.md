
# Math Statistics Calculator
This Go program calculates various statistical measures for a set of numbers provided in a text file.
Features

    Calculates the average (mean) of the numbers
    Calculates the median of the numbers
    Calculates the variance of the numbers
    Calculates the standard deviation of the numbers

Usage

    Compile the Go program:

go build -o math-stats your-program.go

    Run the program with the path to the input text file as an argument:

./math-stats data.txt

The program expects the input file to be a text file with one number per line. If the file is not found or is empty, the program will print an error message.
Input File Format
The input file should contain one number per line, with each number separated by a newline character. For example:

10.5
7.2
3.8
15.1
9.4

Output
The program will output the following statistical measures:

    Average (mean)
    Median
    Variance
    Standard Deviation

The output will be printed to the console in the following format:

Average: 9
Median: 9
Variance: 16
Standard Deviation: 4

Dependencies
This program uses the following Go packages:

    bufio
    fmt
    math
    os
    strconv
    strings
    math-skills/functions

The math-skills/functions package contains the implementation of the statistical functions used in this program.
License
This program is licensed under the MIT License.