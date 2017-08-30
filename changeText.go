package main

import  "fmt"
//замена слов
func findCens(text,cens string) (int){// поиск строки равной cens
	var bann int 
	var lastNumberSymbolCens int = len(cens)
	for i := 0 ; i < len(text) ; i ++{
		if string(text[i]) == string(cens[0]) && string(text[i:i+lastNumberSymbolCens]) == cens{//если совпали первые символы и совпали строки то добавить в банн 
		bann = 1
		}
	}
	return(bann)
}
func replace(text string,cens string,change string) (string) {
	var amendedText string
	var lastNumberSymbolCens int = len(cens) // номер последнего символа в строке цензуры
		for i := 0 ; i < len(text) ; i++{ // цикл по text
		if string(text[i]) != string(cens[0]){ /*если и-тый символ текста не равен первому символу
	цензуры то добавлять этот симлолы в исправленный текст*/
			m1 := text[i] // создаем переменную m1 со значением i-го элемента text
			amendedText += string(m1) /* преобразуем переменную m1 в тип string и добавляем в
	исправленный текст*/
		} else if string(text[i:i+lastNumberSymbolCens]) != cens { /* если итый элемент text равен
	первому элементу cens идет проверка равентсва подстроки text длинной равной cens с cens*/  
			m2 := text[i] //создаем переменную m2 со значением i-го элемента text
			amendedText += string(m2) /* преобразуем переменную m2 в тип string и добавляем в
	исправленный текст*/
		} else { // иначе 
			i = i + lastNumberSymbolCens // меняем значение i на значение индекса после пропущенной подстроки 
			m3 := change //создаем переменную m3 со значением i-го элемента text
			amendedText += m3 /* преобразуем переменную m3 в тип string и добавляем в
	исправленный текст*/
			}
		}
	return (amendedText)
}
func relis(text string, cens string,change string)(string){
	var bann int = findCens(text,cens)// присвоить банн результат функции финдстринг
	var amendedText string = replace(text,cens,change)// присвоить амендеттекст результат функции реплейс
	for bann == 1 {
		text = amendedText
		amendedText = replace(text,cens,change)
		bann = findCens(text,cens)
	}
	return(text)
}
func main() {
	var text string = "12a345aaasdsdsd678asd90"
	var cens string = "asd"
	var change string = "rec"
	fmt.Println(relis(text,cens,change))
}