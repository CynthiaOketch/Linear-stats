# Linear stats

This Go program reads a set of numerical data from a file and computes two key statistics: the Linear Regression Line and the Pearson Correlation Coefficient. The program is designed to be simple, efficient, and accurate, producing results with the required precision.

## Features

- **Linear Regression Line Calculation**: Computes the best-fit line through the data points using the formula `y = mx + b`.

   ![alt text](<Screenshot from 2024-09-04 11-16-51.png>)
- **Pearson Correlation Coefficient Calculation**: Measures the linear correlation between two variables, providing a value between -1 and 1.

   ![alt text](<Screenshot from 2024-09-04 11-17-01.png>)

## File Input

The program reads data from a file passed as an argument. The file should contain a list of numbers, each on a new line. These numbers represent the y-values of a graph where the x-values are the line numbers (0, 1, 2, 3, ...).

### Example `data.txt`


## Installation

1. **Clone the Repository**:

   ```bash
   git clone https://learn.zone01kisumu.ke/git/coketch/linear-stats.git
   cd linear-stats

## Usage
To run the program, use the following command:
```go
    go run main.go data.txt
```
The output will be formatted as:

    Linear Regression Line: y = <m_value>x + <b_value>
    Pearson Correlation Coefficient: <correlation_value>
    
The m_value and b_value in the Linear Regression Line will have 6 decimal places.
The correlation_value for the Pearson Correlation Coefficient will have 10 decimal places.

## Contributing
Contributions are welcome! If you have suggestions or improvements, feel free to submit a Pull Request or open an issue.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
