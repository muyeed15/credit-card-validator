# Credit Card Validator

A simple credit card validation system implemented in Go using the Luhn Algorithm. This project checks whether a credit card number is valid based on the Luhn checksum formula.

## Overview

The Luhn Algorithm, also known as the "modulus 10" or "mod 10" algorithm, is used to validate various identification numbers, including credit card numbers. This implementation checks the validity of a credit card number by calculating the checksum and verifying if it meets the criteria.

## Features

- Validate credit card numbers with lengths of 15 or 16 digits.
- Check if the card number is valid according to the Luhn Algorithm.

## Files

- `luhn.go`: Contains the `CheckLuhn` function that implements the Luhn Algorithm for validation.
- `main.go`: Provides a simple command-line interface to enter a credit card number and check its validity.

## Usage

To run the program:

1. Clone the repository:

   ```bash
   git clone https://github.com/muyeed15/credit-card-validator.git
   ```

2. Navigate to the project directory:

   ```bash
   cd credit-card-validator
   ```

3. Run the Go program:

   ```bash
   go run main.go
   ```

4. Enter a credit card number when prompted and receive validation feedback.

## Example

```bash
 <--------- Simple Credit Card Validation System (Luhn Algorithm) --------->
Enter your credit card number: 1234567812345670

Your card is valid!
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

The Luhn Algorithm was developed by Hans Peter Luhn, a researcher at IBM.
