package banking

type BankAccount struct {
	name    string
	account int
}

func NewAccount(name string) *BankAccount {
	account := BankAccount{name: name, account: 0}
	return &account
}

func (b BankAccount) Deposit(amount int) { // b~ 는 receiver
	b.account += amount
}

func (b BankAccount) Balance() int {
	return b.account
}
