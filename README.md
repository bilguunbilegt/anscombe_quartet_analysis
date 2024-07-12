# Anscombe's Quartet Analysis

This project performs linear regression analysis and generates scatter plots for Anscombe's Quartet dataset using Golang. The results are written to results.txt file and scatter plots are saved as PNG images.

## Project Structure

- `main.go`: The main program that performs linear regression analysis, generates scatter plots, and writes the results to a file.
- `main_test.go`: The test file containing unit tests for the main functionalities.
- `results.txt`: The output file where the results of the linear regression analysis are written.
- `anscombe_set_1.png`, `anscombe_set_2.png`, `anscombe_set_3.png`, `anscombe_set_4.png`: Scatter plot images for each dataset in Anscombe's Quartet.

## Dependencies

This project uses the following Go packages:
- `github.com/montanaflynn/stats`: For statistical analysis.
- `gonum.org/v1/plot`: For generating plots.

## Installation

1. Install Go on your machine if you haven't already. Follow the instructions on the (https://golang.org/doc/install).

2. Get the necessary Go packages:
    ```sh
    go get github.com/montanaflynn/stats
    go get gonum.org/v1/plot/...
    ```

## Usage

1. Clone the repository:
    ```
    git clone https://github.com/bilguunbilegt/anscombe_quartet_analysis.git
    cd anscombe_quartet_analysis
    ```

2. Run the program:
    ```
    go run main.go
    ```

3. Check the output:
   - The results of the linear regression analysis will be written to `results.txt`.
   - Scatter plot images will be saved as `anscombe_set_1.png`, `anscombe_set_2.png`, `anscombe_set_3.png`, and `anscombe_set_4.png`.

## Testing

To test the application, follow these steps:

1. Ensure you are in the project directory:
    ```
    cd anscombe_quartet_analysis
    ```

2. Run the tests:
    ```
    go test
    ```

3. The tests will run and provide output indicating success or failure of the unit tests.

### Tests Included

- **Data Quality Check**: Tests the `checkDataQuality` function to ensure it correctly identifies issues in the input data.
- **Linear Regression**: Tests the `linearRegression` function to verify that it returns the correct linear regression series.
- **R-squared Calculation**: Tests the `calculateRSquared` function to ensure it correctly computes the R-squared value.
- **Correlation Calculation**: Tests the `stats.Correlation` function to verify it returns the expected correlation value.

