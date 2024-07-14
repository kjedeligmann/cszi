package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Завдання 1. Шифрування та дешифрування тексту за допомогою шифра Цезаря.\n\n")

	engInput := []English{
		{"Keep your friends close, but your enemies closer.", 15},
		{"Cryptography is the practice and study of techniques for secure communication in the presence of third parties called adversaries.", -7},
	}
	ukrInput := []Ukrainian{
		{"Украй простий приклад симетричного шифрування — шифр підстановки.", 18},
		{"Наука про математичні методи забезпечення конфіденційності, цілісності і автентичності інформації.", -5},
	}

	for _, input := range engInput {
		fmt.Printf("Вхідний текст: %s\n", input.Plaintext)
		fmt.Printf("Ключ: %d\n", input.Key)
		fmt.Printf("Вихідний текст: %s\n\n", input.encrypt())
	}
	for _, input := range ukrInput {
		fmt.Printf("Вхідний текст: %s\n", input.Plaintext)
		fmt.Printf("Ключ: %d\n", input.Key)
		fmt.Printf("Вихідний текст: %s\n\n", input.encrypt())
	}

	fmt.Printf("Завдання 2. Дешифрування тексту українською за допомогою логарифмічної функції правдоподібності.\n\n")

	cryptogram := "3.	Іцєгбдтьґдц, бдгцямач очс ьюмґциацж іцєгчо (дм стльцж ґеимґацж), хмофсц оцсмкдй стлье ґдмдцґдциае чаєбгямзчк вгб дтьґд вбочсбяютаал, їб ябфт недц оцьбгцґдмаб сюл хюмяе. Вчґюл очсьгцддл имґдбдабпб мамючхе о сто’лдбяе ґдбючддч, ямщфт оґч дмьч іцєгц ґдмюц нчюйі-ятаі ютпьб хюмяацяц сбґочситаця ємжчозтя. Ьюмґциач іцєгц хнтгтпюц вбвеюлгачґдй, о бґабоабяе, е оцпюлсч пбюбобюбябь. Ямщфт оґч іцєгц хмюцімюцґй нтххмжцґацяц втгтс ьгцвдбмамючхбя х оцьбгцґдмааля имґдбдабпб мамючхе сб оцамжбсе вбючмюємочдабпб іцєге. Оцамжчс вбюлпмо е дбяе, їбн оцьбгцґдбоеомдц гчхач іцєгц (амвгцьюмс, мюємочдц вчсґдмабоьц) сюл гчхацж имґдца вбочсбяютаал. Е вбючмюємочдабяе іцєгч Очфтатгм, мюпбгцдя іцєгеомаал оцьбгцґдбоеу ьюкибот ґюбоб, льт ьтгеу вчсґдмабоьбк ючдтг о хмютфабґдч очс дбпб, льм ючдтгм ьюкибобпб ґюбом оцьбгцґдбоеудйґл. Е ґтгтсцач дцґлим очґчяґбдцж гбьчо, Имгюйх Нтннчсф вбьмхмо, їб вбючмюємочдач іцєгц зйбпб дцве хмюціцюцґй имґдьбоб нтххмжцґацяц втгтс имґдбдаця мамючхбя."

	freqs := countLetters(cryptogram)
	logLikelihoods, maxLL := llForKeys(freqs)
	decrypted := (&Ukrainian{
		Plaintext: cryptogram,
		Key:       maxLL,
	}).decrypt()

	fmt.Printf("Криптограма:\n%s\n", cryptogram)

	fmt.Println("Частотні характеристики криптограми:")
	for idx, freq := range freqs {
		fmt.Printf("%q: %d; ", ukrU[idx], freq)
	}
	fmt.Println("\nЗначення логарифмічної функції правдоподібності для ключів:")
	for idx, ll := range logLikelihoods {
		fmt.Printf("%d: %.1f; ", idx, ll)
	}
	fmt.Printf("\nНайімовірніший ключ: %d\n", maxLL)

	fmt.Printf("Криптограма, розшифрована за допомогою ключа %d:\n%s\n", maxLL, decrypted)

	// Trying out go-echarts
	createHTML(logLikelihoods)
}
