package main

import "fmt"

type Currency struct {
	Name   string
	Symbol string
}

type Price struct {
	Value    float64
	Currency Currency
}

func (p *Price) readable() string {
	if p.Value == float64(int(p.Value)) {
		return fmt.Sprintf("%s%.0f", p.Currency.Symbol, p.Value)
	}

	return fmt.Sprintf("%s%.1f", p.Currency.Symbol, p.Value)
}

type Product struct {
	Name  string
	Price Price
}

func (p *Product) readable() string {
	return fmt.Sprintf("%s", p.Name)
}

type Earnings struct {
	Product Product
	Total   Price
}

func (e *Earnings) readable() string {
	return fmt.Sprintf("%s: %s", e.Product.readable(), e.Total.readable())
}

type Expense struct {
	Name  string
	Total Price
}

type Shop struct {
	Earnings []Earnings
	Expenses []Expense
}

func (s *Shop) Sell(p Product, amount int) {
	if amount < 1 {
		return
	}

	for i, product := range s.Earnings {
		if product.Product == p {
			s.Earnings[i].Total.Value += p.Price.Value * float64(amount)
			return
		}
	}
}

func (s *Shop) Income(print bool) map[Currency]float64 {
	income := make(map[Currency]float64)
	for _, earnings := range s.Earnings {
		if print {
			fmt.Printf("%s\n", earnings.readable())
		}
		if _, ok := income[earnings.Product.Price.Currency]; ok {
			income[earnings.Product.Price.Currency] += earnings.Total.Value
		} else {
			income[earnings.Product.Price.Currency] = earnings.Total.Value
		}
	}
	return income
}

func (s *Shop) Outcome() map[Currency]float64 {
	outcome := make(map[Currency]float64)
	for _, expense := range s.Expenses {
		if _, ok := outcome[expense.Total.Currency]; ok {
			outcome[expense.Total.Currency] += expense.Total.Value
		} else {
			outcome[expense.Total.Currency] = expense.Total.Value
		}
	}
	return outcome
}

func main() {
	dollar := Currency{"USD", "$"}
	// Write your code here
	shop := initShop(dollar)
	printRevenue(shop)

	shop.Expenses = append(shop.Expenses, getExpense("Staff expenses", dollar))
	shop.Expenses = append(shop.Expenses, getExpense("Other expenses", dollar))

	printNetIncome(shop)
}

func initShop(currency Currency) Shop {

	bubblegum := Product{"Bubblegum", Price{2.0, currency}}
	toffee := Product{"Toffee", Price{0.2, currency}}
	iceCream := Product{"Ice cream", Price{5.0, currency}}
	milkChocolate := Product{"Milk Chocolate", Price{4.0, currency}}
	doughnut := Product{"Doughnut", Price{2.5, currency}}
	pancake := Product{"Pancake", Price{3.2, currency}}

	shop := Shop{
		[]Earnings{
			{bubblegum, Price{0.00, currency}},
			{toffee, Price{0.00, currency}},
			{iceCream, Price{0.00, currency}},
			{milkChocolate, Price{0.00, currency}},
			{doughnut, Price{0.00, currency}},
			{pancake, Price{0.00, currency}},
		},
		[]Expense{},
	}

	shop.Sell(bubblegum, 101)
	shop.Sell(toffee, 590)
	shop.Sell(iceCream, 450)
	shop.Sell(milkChocolate, 420)
	shop.Sell(doughnut, 430)
	shop.Sell(pancake, 25)

	return shop
}

func getExpense(name string, currency Currency) Expense {
	expense := Expense{name, Price{0.0, currency}}
	fmt.Printf("%s: ", expense.Name)
	if _, err := fmt.Scan(&expense.Total.Value); err != nil {
		panic(err)
	}
	return expense
}

func printRevenue(shop Shop) {
	fmt.Println("Earned amount:")
	income := shop.Income(true)

	output := "\nIncome:"
	for k, v := range income {
		output += " " + (&Price{v, k}).readable()
	}
	fmt.Println(output)
}

func printNetIncome(shop Shop) {
	income := shop.Income(false)
	outcome := shop.Outcome()

	output := "\nNet income:"
	net := make(map[Currency]float64)
	for currency, v := range income {
		net[currency] += v
	}
	for currency, v := range outcome {
		net[currency] -= v
	}
	for k, v := range net {
		output += " " + (&Price{v, k}).readable()
	}
	fmt.Println(output)
}

func printProducts(shop Shop) {
	fmt.Println("Prices:")
	for _, earnings := range shop.Earnings {
		fmt.Printf("%s: %s\n", earnings.Product.readable(), earnings.Product.Price.readable())
	}
}
