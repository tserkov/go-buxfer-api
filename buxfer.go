package buxfer

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

/*
	@todo vendor go-querystring
*/

const (
	API_LOGIN_URL    = "https://www.buxfer.com/api/login?userid=%s&password=%s"
	API_ENDPOINT_URL = "https://www.buxfer.com/api/%s?token=%s"
)

func AddTransaction(token string, tx *TransactionParameters) (bool, error) {
	v := url.Values{}
	v.Set("description", stmt.Description)
	v.Set("amount", stmt.Statement)
	v.Set("accountId", stmt.AccountID)
	v.Set("fromAccountId", stmt.FromAccountID)
	v.Set("toAccountId", stmt.ToAccountID)
	v.Set("date", stmt.Date)
	v.Set("tags", stmt.Tags)
	v.Set("type", stmt.Type)
	v.Set("status", stmt.Status)

	if stmt.Type == "sharedBill" {
		v.Set("payers", stmt.Payers)
		v.Set("sharers", stmt.Sharers)
		v.Set("isEvenSplit", stmt.IsEvenSplit)
	} else if stmt.Type == "loan" {
		v.Set("loanedBy", stmt.LoanedBy)
		v.Set("borrowedBy", stmt.BorrowedBy)
	} else if stmt.Type == "paidForFriend" {
		v.Set("paidBy", stmt.PaidBy)
		v.Set("paidFor", stmt.PaidFor)
	}

	res := new(AddTransactionResponseRoot)
	err := post(fmt.Sprintf(API_ENDPOINT_URL, "add_transaction", token), v, &res)
	if err != nil {
		return false, err
	}

	return res.Response.TransactionAdded, nil
}

func Login(username, password string) (string, error) {
	res := new(LoginResponseRoot)
	err := get(fmt.Sprintf(API_LOGIN_URL, username, password), &res)
	if err != nil {
		return "", err
	}

	return res.Response.Token, nil
}

func GetAccounts(token string) ([]Account, error) {
	res := new(AccountsResponseRoot)
	err := get(fmt.Sprintf(API_ENDPOINT_URL, "accounts", token), &res)
	if err != nil {
		return nil, err
	}

	return res.Response.Accounts, nil
}

func GetBudgets(token string) ([]Budget, error) {
	res := new(BudgetsResponseRoot)
	err := get(fmt.Sprintf(API_ENDPOINT_URL, "budgets", token), &res)
	if err != nil {
		return nil, err
	}

	return res.Response.Budgets, nil
}

func GetContacts(token string) ([]Contact, error) {
	res := new(ContactsResponseRoot)
	err := get(fmt.Sprintf(API_ENDPOINT_URL, "contacts", token), &res)
	if err != nil {
		return nil, err
	}

	return res.Response.Contacts, nil
}

func GetGroups(token string) ([]Group, error) {
	res := new(GroupsResponseRoot)
	err := get(fmt.Sprintf(API_ENDPOINT_URL, "groups", token), &res)
	if err != nil {
		return nil, err
	}

	return res.Response.Groups, nil
}

func GetLoans(token string) ([]Loan, error) {
	res := new(LoansResponseRoot)
	err := get(fmt.Sprintf(API_ENDPOINT_URL, "loans", token), &res)
	if err != nil {
		return nil, err
	}

	return res.Response.Loans, nil
}

func GetReminders(token string) ([]Reminder, error) {
	res := new(RemindersResponseRoot)
	err := get(fmt.Sprintf(API_ENDPOINT_URL, "reminders", token), &res)
	if err != nil {
		return nil, err
	}

	return res.Response.Reminders, nil
}

func GetTags(token string) ([]Tag, error) {
	res := new(TagsResponseRoot)
	err := get(fmt.Sprintf(API_ENDPOINT_URL, "tags", token), &res)
	if err != nil {
		return nil, err
	}

	return res.Response.Tags, nil
}

func GetTransactions(token string, params *TransactionsParameters) ([]Transactions, error) {
	url := fmt.Sprintf(API_ENDPOINT_URL, "transactions", token)

	if params != nil {
		url = url + "&" + query.Values(params).Encode()
	}

	res := new(TransactionsResponseRoot)
	err := get(url, &res)
	if err != nil {
		return nil, err
	}

	return res.Response.Transactions, nil
}

func UploadStatement(token string, stmt *StatementParameters) (float64, error) {
	v := url.Values{}
	v.Set("accountId", stmt.AccountID)
	v.Set("statement", stmt.Statement)
	if stmt.DateFormat != "" {
		v.Set("dateFormat", stmt.DateFormat)
	}

	res := new(UploadStatementResponseRoot)
	err := post(fmt.Sprintf(API_ENDPOINT_URL, "upload_statement", token), v, &res)
	if err != nil {
		return 0.0, err
	}

	return res.Response.Balance, nil
}

var client = &http.Client{Timeout: 10 * time.Second}

func get(url string, dest *interface{}) (err error) {
	r, err := client.Get(url)
	if err != nil {
		return
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(dest)

	if dest.Response.Status != "OK" {
		return "", errors.New("Not OK") // @todo
	}
}

func post(url string, params url.Values, dest *interface{}) (err error) {
	r, err := client.PostForm(url, params)
	if err != nil {
		return
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(dest)

	if dest.Response.Status != "OK" {
		return "", errors.New("Not OK") // @todo
	}
}
