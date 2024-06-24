package main

import (
	"fmt"
)

const Strikethrough = "\033[9m"
const Reset = "\033[0m"

// Реализовать паттерн «адаптер» на любом примере.

// Причин для применения адаптеров может быть много, однако, суть адаптера заключается в том, что он позволяет использовать код (метод)
// для случаев, для которых он изначально не был предназначен.
// Как пример, в проекте используется структура (old)PrinterIml, которая реализует интерфейс (old)Printer
// эта структура позволяет выдавать Имя и возраст пользователя методом Print.
// Но вот появился новый интерфейс принтера iPrinter500+, который позволяет выдавать на 50% больше информации, однако для этого у него уже другой метод
// PrintDetails. Если мы хотим использовать новый метод, но по-старому, то мы можем воспользоваться адаптером
// Ниже приведена реализация

// OldPrinter - старый интерфейс принтера
type OldPrinter interface {
	Print(user User)
}

// User - структура пользователя
type User struct {
	Name       string
	Age        int
	Occupation string
}

// OldPrinterImpl - старый принтер, который реализует OldPrinter
type OldPrinterImpl struct{}

func (p *OldPrinterImpl) Print(user User) {
	fmt.Printf("Old Printer - Name: %s, Age: %d\n", user.Name, user.Age)
}

// NewPrinter - новый интерфейс принтера
type NewPrinter interface {
	PrintDetails(user User)
}

// NewPrinterImpl - новый принтер, который реализует NewPrinter
type NewPrinterImpl struct{}

func (p *NewPrinterImpl) PrintDetails(user User) {
	fmt.Printf("New Printer - User: %s, Age: %d, Род деятельности: %s\n", user.Name, user.Age, user.Occupation)
}

// PrinterAdapter - адаптер, который позволяет использовать NewPrinter через интерфейс OldPrinter
type PrinterAdapter struct {
	newPrinter NewPrinter
}

func (a *PrinterAdapter) Print(user User) {
	a.newPrinter.PrintDetails(user)
}

func main() {
	// Старый принтер
	oldPrinter := &OldPrinterImpl{}
	user := User{Name: "John Doe", Age: 30, Occupation: "без" + Strikethrough + "з" + Reset + "работный"}
	oldPrinter.Print(user)

	// Новый принтер через адаптер
	newPrinter := &NewPrinterImpl{}
	adapter := &PrinterAdapter{newPrinter: newPrinter}
	adapter.Print(user)
}
