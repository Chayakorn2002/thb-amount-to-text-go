# THB Amount to Text Go

This repository contains a Go project designed to convert numerical Thai Baht (THB) amounts into Thai text representation, which can be useful in various applications such as financial reporting, or receipt printing. 

The core conversion logic lies in `DecimalToBahtText` function, located in `decimal_to_baht_text.go`' This function accepts a `decimal.Decimal` from `github.com/shopspring/decimal` package as input and returns the Thai Baht amount as a `string` (textual representation).

## Usage

This project offers two methods for execution and interaction, through both command-line and HTTP service integration.

**Clone the Repository:**

Clone the repository to your local development environment:

```bash
git clone https://github.com/Chayakorn2002/thb-amount-to-text-go.git

cd thb-amount-to-text-go
```

### Option 1: Command-Line Interface Execution

1.  **Run the Application:**
    
    The project includes a `Makefile` for easy execution. You can use `make try` command to run the application.

    ```bash
    make try
    ```

    Upon execution, the terminal will prompt for an amount in THB input. Enter the desired numerical amount, and the program will output its corresponding Thai text representation.

### Option 2: HTTP Server

1.  **Start the HTTP Server**

    Run `make start` to start the Echo HTTP server, which exposes the `GET /convert` endpoint.

    ```bash
    make start
    ```

2.  **Try curling the endpoint**

    Once the application is running, you can test the conversion functionality by sending a GET request to the local server. The application will be accessible at `http://localhost:8080`. You can specify the amount you desire using the `amount` query parameter.

    ```bash
    curl --location 'http://localhost:8080/convert?amount=100.50'
    ```

    Replace `100.50` with the desired Thai Baht amount you want to convert to text.

### Option 3: Importing as a Go Module

1.  **Import the module**

    import and utilize the DecimalToBahtText function from the utils package of `github.com/Chayakorn2002/thb-amount-to-text-go` module.

    ```go
    package main

    import (
        "fmt"

        "github.com/Chayakorn2002/thb-amount-to-text-go/utils"
        "github.com/shopspring/decimal"
    )

    func main() {
        decimal, err := decimal.NewFromString("123456789.99")
        if err != nil {
            panic(err)
        }

        amountTxt, err := utils.DecimalToBahtText(decimal)
        if err != nil {
            panic(err)
        }

        fmt.Println(amountTxt)
    }
    ```
