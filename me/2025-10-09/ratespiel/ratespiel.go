package main //Hauptpacket meines Programmes

import (
	"fmt"       //für terminal
	"math/rand" //für Zufallszahlen
	"time"      //für Wirklichen Zufall, da Computer zu unterschiedlichen Zeiten Zufallszahlen hat.
) //importieren von Bibilotheken immer untereinander, ansonsten werden sie nicht unterstrichen und damit nicht erkannt.

func main() {
	fmt.Println("numberguessinggame") //Spielname
	GuessGame()                       //Funktion aufrufen, also hier die Spiellogik
}
func ReadNumber() int {
	var n int                                        //int = integer = ganze Zahl
	fmt.Print("Rate eine Zahl zwischen 1 und 100: ") //Eingabeaufforderung
	fmt.Scan(&n)                                     //liest die Eingabe des Spielers ein und speichert sie in der Variable n
	return n                                         //gibt die Zahl zurück, die der Spieler eingegeben hat

}
func GuessGame() {
	rand.Seed(time.Now().UnixNano())   //zufallszahl, ist aber veraltet, deswegen durchgestrichen. Funktioniert aber trotzdem
	secretNumber := rand.Intn(100) + 1 //1-100, weil im Programmieren von 0 gezählt wird

	for n := 0; n < 5; n++ { //3 Versuche, Schleife n++ für n plus 1
		guess := ReadNumber() //liest die Zahl des Spielers ein und speichert diese
		if guess == secretNumber {
			fmt.Println("Richtig geraten!")
			return
		} else if guess < secretNumber {
			fmt.Println("Zu niedrig")
		} else {
			fmt.Println("Zu hoch")
		}

	}

	fmt.Println("Zu viele falsche Versuche")

}
