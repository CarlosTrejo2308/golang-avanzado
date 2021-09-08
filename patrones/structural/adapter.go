package main

// Adapter: Adapta una interfaz a como lo pide otra

type Payment interface {
	Pay()
}

type CashPayment struct{}

func (c *CashPayment) Pay() {
	println("Pago con efectivo")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

// Redefinimos el metodo, hay que adaptarlo
func (b *BankPayment) Pay(banckAccount int) {
	println("Pago con tarjeta:", banckAccount)
}

// Creando un adaptador de bank a payment
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

// Adaptamos el metodo pay de bank a payment
func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	/*
		No se puede usar directamente el metodo pay de bankpayment,
		bank := &BankPayment{}
		ProcessPayment(bank)
	*/

	bpa := &BankPaymentAdapter{
		BankPayment: &BankPayment{},
		bankAccount: 123,
	}
	ProcessPayment(bpa)

}
