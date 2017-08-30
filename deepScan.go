package main

import  "fmt"
//удаление слов из текта(глубокая проверка)
func findCens(text,cens string) (string){// поиск строки равной cens
	var bann string
	var lastNumberSymbolCens int = len(cens)
	for i := 0 ; i < len(text) ; i ++{
		if string(text[i]) == string(cens[0]) && string(text[i:i+lastNumberSymbolCens]) == cens{//если совпали первые символы и совпали строки то добавить в банн 
			bann = string(text[i : i + lastNumberSymbolCens])
		}
	}
	return(bann)
}
func replace(text,cens string) (string) {
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
			m3 := text[i] //создаем переменную m3 со значением i-го элемента text
			amendedText += string(m3) /* преобразуем переменную m3 в тип string и добавляем в
	исправленный текст*/
			}
		}
	return (amendedText)
}
func relis(text, cens string)(string){
	var bann string = findCens(text,cens)// присвоить банн результат функции финдстринг
	var amendedText string = replace(text,cens)// присвоить амендеттекст результат функции реплейс
	for bann == cens {
		text = amendedText
		amendedText = replace(text,cens)
		bann = findCens(text,cens)
	}
	return(text)
}
func main() {
	var text string = "12a345aaasdsdsd678asd90"
	var cens string = "asd"
	fmt.Println(relis(text,cens))
}