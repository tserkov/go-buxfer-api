# Buxfer API (in Go!)

## Example
```golang
package main

import (
	"fmt"
	
	"github.com/tserkov/go-buxfer-api"
)

func main() {
	// Acquire an authorization token first
	token, err := buxfer.Login("my_username", "my_password")

	// Then use the token to call an endpoint, like accounts
	accounts, _ := buxfer.GetAccounts(token)

    for _, account := range accounts {
		fmt.Printf(
			"I have %s%.2f in my %s account!\n",
			account.Currency,
			account.Balance,
			account.Name,
		)
    }
}
```

## Methods
### `Login(username string, password string)`
get a unique 'token' which must included on all future requests

### `AddTransaction(token string, transaction AddTransactionParameters)`
add a transaction

### `Accounts(token string)`
list of accounts with balances

### `Budgets(token string)`
list of your budgets

### `Contacts(token string)`
list of your contacts

### `Groups(token string)`
list of your groups

### `Loans(token string)`
list of your loans

### `Reminders(token string)`
list of your reminders

### `Tags(token string)`
list of your transaction tags

### `Transactions(token string, params TransactionParameters)`
list of transactions, 25 at a time

### `UploadStatement(token string, statement StatementParameters)`
upload a bank or credit card statement
