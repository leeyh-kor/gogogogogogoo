package banking

import "errors"

type BankAccount struct {
	name    string
	account int
}

func NewAccount(name string) *BankAccount {
	account := BankAccount{name: name, account: 0}
	return &account
}

func (b BankAccount) Deposit(amount int) { //b는 복사본이기 때문에.. b.acccount += 이 원본에 적용되지 않음
	b.account += amount
}

func (b *BankAccount) Deposit2(amount int) { //포인터를 활용해서 원본에 접근하기 때문에 임마는 원본이 바뀜
	b.account += amount
}

func (b BankAccount) Balance() int { // b~ 는 receiver라고 부름
	return b.account // 이렇게 단순히 보여만 주는 함수에서는 원본 불러와서 괜한 위험 x
}

func (b *BankAccount) Withdraw(amount int) error { // nill => none null 과 같은것
	if b.account < amount {
		return errors.New("Can't withdraw you are poor")
	}
	b.account -= amount
	return nil
}
