package buxfer

import (
	"time"
)

type AccountsResponseRoot struct {
	Response AccountsResponse `json:"response"`
}

type AccountsResponse struct {
	Accounts []Account `json:"accounts"`
	Status   string    `json:"status"`
}

type Account struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Bank       string    `json:"bank"`
	Balance    float64   `json:"balance"`
	Currency   string    `json:"currency"`
	LastSynced time.Time `json:"lastsynced"`
}

type AddTransactionResponseRoot struct {
	Response AddTransactionResponse `json:"response"`
}

type AddTransactionResponse struct {
	Status           string `json:"status"`
	TransactionAdded bool   `json:"transactionAdded"`
	ParseStatus      string `json:"parseStatus"`
}

type AddTransactionParameters struct {
	Description   string
	Amount        float64
	AccountID     string
	FromAccountID string
	ToAccountID   string
	Date          time.Time // YYYY-MM-DD
	Tags          string    // comma-separated
	Type          string    // income, expense, transfer, refund, sharedBill, paidForFriend, loan
	Status        string    // cleared, pending

	// type=sharedBill only
	Payers      []Payer
	Sharers     []Sharer
	IsEvenSplit bool

	// type=loan only
	LoanedBy   string
	BorrowedBy string

	// type=paidForFriend only
	PaidBy  string
	PaidFor string
}

type BudgetsResponseRoot struct {
	Response BudgetsResponse `json:"response"`
}

type BudgetsResponse struct {
	Status  string   `json:"status"`
	Budgets []Budget `json:"budgets"`
}

type Budget struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Limit         string   `json:"limit"`
	Remaining     float64  `json:"remaining"`
	Period        string   `json:"period"`
	CurrentPeriod string   `json:"currentPeriod"`
	Tags          string   `json:"tags"`
	Keywords      []string `json:"keywords"`
}

type ContactsResponseRoot struct {
	Response ContactsResponse `json:"response"`
}

type ContactsResponse struct {
	Status   string    `json:"status"`
	Contacts []Contact `json:"contacts"`
}

type Contact struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
}

type GroupsResponseRoot struct {
	Response GroupsResponse `json:"response"`
}

type GroupsResponse struct {
	Status string  `json:"status"`
	Groups []Group `json:"groups"`
}

type Group struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Consolidated bool          `json:"consolidated"`
	Members      []GroupMember `json:"members"`
}

type GroupMember struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Balance float64 `json:"balance"`
}

type LoansResponseRoot struct {
	Response LoansResponse `json:"response"`
}

type LoansResponse struct {
	Status string `json:"status"`
	Loans  []Loan `json:"loans"`
}

type Loan struct {
	Entity      string  `json:"entity"`
	Type        string  `json:"type"`
	Balance     float64 `json:"balance"`
	Description string  `json:"description"`
}

type LoginResponseRoot struct {
	Response LoginResponse `json:"response"`
}

type LoginResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

type RemindersResponseRoot struct {
	Response RemindersResponse `json:"response"`
}

type RemindersResponse struct {
	Status    string     `json:"status"`
	Reminders []Reminder `json:"reminders"`
}

type Reminder struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	StartDate string  `json:"startDate"`
	Period    string  `json:"period"`
	Amount    float64 `json:"amount"`
	AccountID string  `json:"accountId"`
}

type TagsResponseRoot struct {
	Response TagsResponse `json:"response"`
}

type TagsResponse struct {
	Status string `json:"status"`
	Tags   []Tag  `json:"tag"`
}

type Tag struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parentId"`
}

type TransactionsResponseRoot struct {
	Response TransactionsResponse `json:"response"`
}

type TransactionsResponse struct {
	Status          string        `json:"status"`
	NumTransactions int           `json:"numTransactions"`
	Transactions    []Transaction `json:"transactions"`
}

type Transaction struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	AccountID   string  `json:"accountId"`
	Tags        string  `json:"tags"`
	ExtraInfo   string  `json:"extraInfo"`
}

type TransactionsParameters struct {
	AccountID   string    `json:"accountId,omitempty"`
	AccountName string    `json:"accountName,omitempty"`
	BudgetID    string    `json:"budgetId,omitempty"`
	BudgetName  string    `json:"budgetName,omitempty"`
	ContactID   string    `json:"contactId,omitempty"`
	ContactName string    `json:"contactName,omitempty"`
	EndDate     time.Time `json:"endDate,omitempty"`
	GroupID     string    `json:"groupId,omitempty"`
	GroupName   string    `json:"groupName,omitempty"`
	Month       time.Time `json:"month,omitempty"`
	StartDate   time.Time `json:"startDate,omitempty"`
	TagID       string    `json:"tagId,omitempty"`
	TagName     string    `json:"tagName,omitempty"`
}

type UploadStatementResponseRoot struct {
	Response UploadStatementResponse `json:"response"`
}

type UploadStatementResponse struct {
	Status   string  `json:"status"`
	Uploaded bool    `json:"uploaded"`
	Balance  float64 `json:"balance"`
}

type StatementParameters struct {
	AccountID  string
	Statement  string
	DateFormat string // "MM/DD/YYYY" or "DD/MM/YYYY"
}
