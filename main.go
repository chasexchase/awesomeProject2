package main

// На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).

// Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.

// Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.

// Sample Input:

// 12 мин. 13 сек.
// Sample Output:

// Fri May 15 19:28:18 UTC 2020

import (
	"fmt"
	"time"
)

func main() {
	const (
		now      = 1589570165
		UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	)
	var min, sec int64
	fmt.Scanf("%d мин. %d", &min, &sec)
	dur := sec + (min * 60)
	time1 := dur + now
	tm := time.Unix(time1, 0).UTC()
	fmt.Println(tm.Format(time.UnixDate))
}

// _____________________________________________________________________________________________________________________________

// На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).

// Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.

// Sample Input:

// 13.03.2018 14:00:15,12.03.2018 14:00:15
// Sample Output:

// 24h0m0s

// import (
// 	"fmt"
// 	"os"
// 	"bufio"
// 	"strings"
// 	"time"
// )
// func main() {
// 	const(
// 		DateTime   = "02.01.2006 15:04:05"
// 	)
// scanner := bufio.NewScanner(os.Stdin)
// scanner.Scan()
// input := scanner.Text()
// splitedTime := strings.Split(input, ",")   //splitedTime[0] && splitedTime[1] соответственно
// //Не забываем удалить перенос каретки, невидимый глазу, а то не пропарсится!
// splitedTime[1] = strings.TrimRight(splitedTime[1], "\n")
// t1, err := time.Parse(DateTime, splitedTime[0])
// 			if err != nil {
// 			panic(err)
// 		}
// t2, err := time.Parse(DateTime, splitedTime[1])
// 			if err != nil {
// 			panic(err)
// 		}
// var res time.Duration

// res = time.Since(t2) - time.Since(t1)

// fmt.Println(res.Round(time.Second))
// fmt.Println(time.Since(t1).Round(time.Second) - time.Since(t2).Round(time.Second))
// }

// _____________________________________________________________________________________________________________________________
// На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:

// 2020-05-15 08:00:00
// Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на стандартный вывод в том же формате.

// Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день, а затем вывести на стандартный вывод в том же формате.

// Sample Input:

// 2020-05-15 08:00:00
// Sample Output:

// 2020-05-15 08:00:00

// import (
// 	"fmt"
// 	"time"
// 	"os"
// 	"bufio"
// )

// func main() {
// 	const(
// 		DateTime   = "2006-01-02 15:04:05"
// 	)
// scanner := bufio.NewScanner(os.Stdin)
// scanner.Scan()
// input := scanner.Text()
// 	t, err := time.Parse(DateTime, input)
// 			if err != nil {
// 			panic(err)
// 		}
// 	if t.Hour() >= 13{
// 		new_dt := t.AddDate(0,0,1)
// 		fmt.Print(new_dt.Format(DateTime))
// 		} else {
// 	fmt.Print(input)
// 	}
// }
// _____________________________________________________________________________________________________________________________

// На стандартный ввод подается строковое представление даты и времени в следующем формате:

// 1986-04-16T05:20:00+06:00
// Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:

// Wed Apr 16 05:20:00 +0600 1986

// Для более эффективной работы рекомендуется ознакомиться с документацией о содержащихся в модуле time константах.

// Sample Input:

// 1986-04-16T05:20:00+06:00
// Sample Output:

// Wed Apr 16 05:20:00 +0600 1986

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	const(
// 		RFC3339     = "2006-01-02T15:04:05Z07:00"
// 	)
// 	var abc string
// 	fmt.Scanln(&abc)
// 	zxc, err := time.Parse(RFC3339, abc)
// 			if err != nil {
// 			panic(err)
// 		}
// 	fmt.Print(zxc.Format("Mon Jan _2 15:04:05 MST 2006"))
// }
// _____________________________________________________________________________________________________________________________

// _____________________________________________________________________________________________________________________________

// import (
// 	"net/http"
// 	"io/ioutil"
// 	"fmt"
// 	"encoding/json"
// )
// type users []struct {
// 	GlobalID int `json:"global_id"`
// }

// func main(){

// const URL = "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"

// rs, err := http.Get(URL)
// if err != nil {
//    panic(err)
// }
// // defer rs.Body.Close()
// //Данные для парсинга будут находится в теле файла : rs.Body
// //Это касается впринципе любой задачи где нужно скачивать файлы

// byteValue, err := ioutil.ReadAll(rs.Body)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
// var globalID users
// var sum int
// json.Unmarshal(byteValue, &globalID)

// 	for i:=0; i<len(globalID); i++{
// 		sum += globalID[i].GlobalID
// 	}
// 	fmt.Print(sum)
// }
// _____________________________________________________________________________________________________________________________

// На стандартный ввод подаются данные о студентах университетской группы в формате JSON:

// {
//     "ID":134,
//     "Number":"ИЛМ-1274",
//     "Year":2,
//     "Students":[
//         {
//             "LastName":"Вещий",
//             "FirstName":"Лифон",
//             "MiddleName":"Вениаминович",
//             "Birthday":"4апреля1970года",
//             "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
//             "Phone":"+7(948)709-47-24",
//             "Rating":[1,2,3]
//         },
//         {
//             // ...
//         }
//     ]
// }
// В сведениях о каждом студенте содержится информация о полученных им оценках (Rating). Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы. Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:

// {
//     "Average": 14.1
// }
// Как вы понимаете, для декодирования используется функция Unmarshal, а для кодирования MarshalIndent (префикс - пустая строка, отступ - 4 пробела).

// Если у вас возникли проблемы с чтением / записью данных, то этот комментарий для вас: в уроках об интерфейсах и работе с файлами мы рассказывали, что стандартный ввод / вывод - это файлы, к которым можно обратиться через os.Stdin и os.Stdout соответственно, они удовлетворяют интерфейсам io.Reader и io.Writer, из них можно читать, в них можно писать.

// Один из способов чтения, использовать ioutil.ReadAll:

// data, err := ioutil.ReadAll(os.Stdin)
// if err != nil {
//     // ...
// }

// // data - тип []byte
// Задачу можно выполнить и другими способами, в частности использовать bufio. Буквально в следующем шаге (через один, на самом деле) будет рассказано о еще одном способе чтения / записи, можете забежать немного вперед, а затем вернуться к задаче.

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

// type (
// 	Student struct {
// 		Rating []int
// 	}

// 	Group struct {
// 		Students []Student
// 	}

// 	Rating struct {
// 		Average float32
// 	}
// )

// func main () {
//     jsonfile, err := os.Open("text.json")
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     byteValue, err := ioutil.ReadAll(jsonfile)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     var group Group
//     json.Unmarshal(byteValue, &group)
//     var counter1, counter2 float32
//     for i:=0; i<len(group.Students); i++{
//         counter1++
//         for range group.Students[i].Rating {
//             counter2 ++
//         }
//     }
//     a := &Rating {float32(counter2)/float32(counter1)}
//     data, err := json.MarshalIndent(a, "", "    ")
//     if err != nil {
// 	fmt.Println(err)
// 	return
// }
//     fmt.Println(string(data))
//     // w := os.Stdout
//     // w.WriteString(string(data))
// }

// _____________________________________________________________________________________________________________________________

// import (
// 	"bufio"
// 	"encoding/csv"
// 	"fmt"
// 	"os"
// )
// func main () {
//     csvFile, _ := os.Open("/Users/a1234/Desktop/task.data.txt")
//     rd:=bufio.NewReader(csvFile)
//     reader:=csv.NewReader(rd)
//     reader.Comma = ';'
//     record, _ := reader.Read()
//     for num, item := range record{
//         if item == "0" {
//             fmt.Println(num+1)
//             break
//         }
//     }
// }
// _____________________________________________________________________________________________________________________________

// Ранее в рамках этого курса при решении задач требовалось прочитать что-то со стандартного ввода и вывести результат соответственно в стандартный вывод. При этом кто-то использовал пакет fmt, а кто-то - bufio + os. Все эти пакеты имеют свои особенности, поэтому в этой задаче мы попробуем решить знакомую нам проблему с помощью пакетов, которые кто-то мог до этого момента и не применять: bufio + os + strconv.

// Задача состоит в следующем: на стандартный ввод подаются целые числа в диапазоне 0-100, каждое число подается на стандартный ввод с новой строки (количество чисел не известно). Требуется прочитать все эти числа и вывести в стандартный вывод их сумму.

// Несколько подсказок: для чтения вы можете использовать типы bufio.Reader и bufio.Scanner, а для записи - bufio.Writer. При чтении помните об ошибке io.EOF. Конвертирование данных из строки в целое число и обратно осуществляется функциями Atoi и Itoa из пакета strconv соответственно. Чтение производится из стандартного ввода (os.Stdin), а запись - в стандартный вывод (os.Stdout).

// Все указанные в тексте задачи пакеты (strconv, bufio, os, io), уже импортированы (другие использовать нельзя), package main объявлен.

// Sample Input:

// 33
// 47
// 12
// 79
// 15
// Sample Output:

// 186

// import (
// 	"fmt"
// 	"strings"
//     "strconv"
//     "bufio"
//     "os"
//     "io"
// )
// func main () {
//     var x int
//     scanner := bufio.NewScanner(os.Stdin)

//     // Аналогичное в цикле
//     for scanner.Scan() { // возвращает true, пока файл не будет прочитан до конца
//     txt := scanner.Text()
//     txtInt, err := strconv.Atoi(txt)
// 	    if err != nil {
// 		    panic(err)
// 	        }
//      x =+ txtInt
//     }

// io.WriteString(os.Stdout,strconv.Itoa(x))
// }

// _____________________________________________________________________________________________________________________________
// Давайте используем ваши знания структур, методов и интерфейсов на практике и реализуем объект, удовлетворяющий интерфейсу fmt.Stringer. Назовем его "Батарейка".

// Во-первых, вы должны объявить новый тип, удовлетворяющий интерфейсу fmt.Stringer.

// Ваш тип должен предусматривать, что на печати он будет выглядеть так: [      XXXX]: где пробелы - "опустошенная" емкость батареи, а X - "заряженная".

// Во-вторых, на стандартный ввод вы получаете строку, состоящую ровно из 10 цифр: 0 или 1 (порядок 0/1 случайный). Ваша задача считать эту строку любым возможным способом и создать на основе этой строки объект объявленного вами на первом этапе типа: надеюсь, вы понимаете, что строка символизирует емкость батарейки: 0 - это "опустошенная" часть, а 1 - "заряженная".

// В-третьих, созданный вами объект должен называться batteryForTest (использование этого имени обязательно).

// В вашем распоряжении фактически весь файл, НО завершающая фигурная скобка функции main() вам не видна, но она присутствует. Перед этой скобкой присутствует функция (которая вам тоже не видна), принимающая в качестве аргумента объект типа fmt.Stringer - batteryForTest, и направляющая его на стандартный вывод, поэтому вам не требуется выводить что-то на печать самостоятельно.

// Удачи!

// Sample Input:

// 1000010011
// Sample Output:

// [      XXXX]

// type Battery struct {
// 	Power string
// }

// func (a Battery) asd() string {
//     count := strings.Count(a.Power, "1")
//     a1 := strings.Repeat(" ", 10-count)
//     a1 += strings.Repeat("X", count)
//     return fmt.Sprintf("[%v]", a1)
// }

// func main (){
// var test string
// fmt.Scan(&test)
// batteryForTest:=Battery{Power:test}
// fn:=batteryForTest.asd()
// fn(batteryForTest)
// }

// _____________________________________________________________________________________________________________________________

// Пришло время для задач, где вы сможете применить полученные знания на практике.

// Обязательные условия выполнения: данные со стандартного ввода читаются функцией readTask(), которая возвращает 3 значения типа пустой интерфейс. Эта функция использует пакеты encoding/json, fmt, и os - не удаляйте их из импорта. Скорее всего, вам понадобится пакет "fmt", но вы можете использовать любой другой пакет для записи в стандартный вывод не удаляя fmt.

// Итак, вы получаете 3 значения типа пустой интерфейс: если все удачно, то первые 2 значения будут float64. Третье значение в идеальном случае будет строкой, которая может иметь значения: "+", "-", "*", "/" (определенная математическая операция). Но такие идеальные случаи будут не всегда, вы можете получить значения других типов, а также строка в третьем значении может не относится к одной из указанных математических операций.

// Результат выполнения программы должен быть таков:

// в штатной ситуации вы должны напечатать в стандартный вывод результат выполнения математической операции с точностью до 4 цифры после запятой (fmt.Printf(%.4f));
// если первое или второе значение не является типом float64, вы должны напечатать сообщение об ошибке вида: value=полученное_значение: тип_значения (например: value=true: bool)
// если третье значение имеет неверный тип или передан знак, не относящийся к указанным выше математическим операциям, сообщение об ошибке должно иметь вид: неизвестная операция
// Гарантируется, что ошибка в аргументах может быть только одна, поэтому если вы при проверке первого значения увидели, что оно содержит ошибку - печатайте сообщение об ошибке и завершайте работу программы, проверка остальных аргументов значения уже не имеет, а проверяющая система воспримет 2 сообщения об ошибке как нарушение условия выполнения задания.

// Удачи!

// import (
// 	"encoding/json" // пакет используется для проверки ответа, не удаляйте его
// 	"fmt"           // пакет используется для проверки ответа, не удаляйте его
// 	"os"            // пакет используется для проверки ответа, не удаляйте его
// )

// func printError(value interface{}) {
//     fmt.Print("value=true: bool")
// }

// func main() {
// 	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
//     										 // все полученные значения имеют тип пустого интерфейса
// var v1, v2 float64
// if v1, ok := value1.(float64); ok {
// 		if v2, ok := value2.(float64); ok {
// 			if a, ok := operation.(string); ok {
// 				switch a {
// 					case "+":
// 						result := fmt.Sprintf("%.4f", v1+v2)
// 						fmt.Printf(result)
// 					case "-":
// 						result:= fmt.Sprintf("%.4f", v1-v2)
// 						fmt.Printf(result)
// 					case "/":
// 						result:= fmt.Sprintf("%.4f", v1/v2)
// 						fmt.Printf(result)
// 					case "*":
// 						result:= fmt.Sprintf("%.4f", v1*v2)
// 						fmt.Printf(result)
// 					default: fmt.Print("неизвестная операция")
// 					return
// 			}
// 				return
// 			}
// 			} else	{
// 				printError(v2)
// 				return
// 			}
// 	} else {
// 	printError(v1)
// 	return
// }
// }
// _____________________________________________________________________________________________________________________________

// var global int64 = 1

// type data struct {
// 	info    string
// 	counter int
// 	add     int32
// }

// func (data *data) save() (b uint64) {
// 	global--
// 	fmt.Println("saving...")
// 	return uint64(data.counter)
// }

// func main() {
// 	if err, data := create("", 0); err != nil {
// 		defer func() {
// 			data.save()
// 			fmt.Println("saved")
// 			fmt.Println("global:", global)
// 		}()
// 		panic(err)
// 	}
// }
// func create(info string, count int) (error, data) {
// 	if len(info) == 0 {
// 		return errors.New("len is 0"), data{info: "", counter: 0, add: 0}
// 	}
// 	data := data{
// 		info:    info,
// 		counter: count,
// 		add:     1,
// 	}
// 	global++
// 	return nil, data
// }

// _____________________________________________________________________________________________________________________________

// Используем анонимные функции на практике.

// Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint, которую внутри функции main в дальнейшем можно будет вызвать по имени fn.

// Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0. Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

// Пакет main объявлять не нужно!
// Вводить и выводить что-либо не нужно!
// Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.

// Sample Input:

// 727178
// Sample Output:

// 28

//func main() {
//	var x uint
//	fn := func(a uint) uint {
//		var s = strconv.FormatUint(uint64(a), 10)
//		var z string
//		for _, v := range s {
//			i, _ := strconv.Atoi(string(v))
//			if i != 0 {
//				if i%2 == 0 {
//					q := strconv.Itoa(i)
//					z += q
//				}
//			}
//		}
//		g, _ := strconv.Atoi(z)
//		p := 100
//		if g == 0 {
//			return uint(p)
//		}
//
//		return uint(g)
//	}
//
//}
// _____________________________________________________________________________________________________________________________

//func counter(start int) (func() int, func()) {
//	// если значение мутирует, то мы получим изменение и в этом замыкании
//	ctr := func() int {
//		return start
//	}
//
//	incr := func() {
//		start++
//	}
//
//	// и в ctr, и в incr сохраняется указатель на start
//	// мы создали замыкания но еще не вызывали
//	return ctr, incr
//}
//
//func main() {
//	// ctr, incr и ctr1, incr1 различаются своим состоянием
//	ctr, incr := counter(100)
//	ctr1, incr1 := counter(100)
//	fmt.Println("counter - ", ctr())
//	fmt.Println("counter1 - ", ctr1())
//	// увеличиваем на 1
//	incr()
//	fmt.Println("counter - ", ctr())
//	fmt.Println("counter1- ", ctr1())
//	// увеличиваем до 2
//	incr1()
//	incr1()
//	fmt.Println("counter - ", ctr())
//	fmt.Println("counter1- ", ctr1())
//}

//func outer(name string) func() {
//	// переменная
//	text := "Modified " + name
//
//	// замыкание. у функции есть доступ к переменной text даже
//	// после выхода за пределы блока
//	foo := func() {
//		fmt.Println(text)
//	}
//
//	// возвращаем замыкание
//	return foo
//}
//
//func main() {
//	// foo это наше замыкание
//	foo := outer("hello")
//
//	// вызов замыкания
//	foo()
//}

//func outer(name string) {
//	// переменная из внешней функции
//	text := "Modified " + name
//
//	// foo это внутренняя функция и из нее есть доступ к переменной `text`
//	// у замыкания есть доступ к этим переменным даже после выхода из блока
//	foo := func() {
//		fmt.Println(text)
//	}
//
//	// вызываем замыкание
//	foo()
//}
//
//func main() {
//	outer("hello")
//}

//func printMessage(message string) {
//	fmt.Println(message)
//}
//
//func getPrintMessage() func(string) {
//	// Возвращаем анонимную функцию
//	return func(message string) {
//		fmt.Println(message)
//	}
//}
//
//func main() {
//	// Именованная функция
//	printMessage("Hello function!")
//
//	// Анонимная фукция объявляется и вызывается
//	func(message string) {
//		fmt.Println(message)
//	}("Hello anonymous function!")
//
//	// Получаем анонимную функцию и вызываем ее
//	printfunc := getPrintMessage()
//	printfunc("Hello anonymous function using caller!")
//
//}

//func add(x int, y int) int      { return x + y }
//func subtract(x int, y int) int { return x - y }
//func multiply(x int, y int) int { return x * y }
//
//func selectFn(n int) func(int, int) int {
//	if n == 1 {
//		return add
//	} else if n == 2 {
//		return subtract
//	} else {
//		return multiply
//	}
//}
//
//func main() {
//
//	f := selectFn(1)
//	fmt.Println(f(3, 4)) // 7
//
//	f = selectFn(3)
//	fmt.Println(f(3, 4)) // 12
//}

//func main() {
//	var a, b int
//	fmt.Scan(&a, &b)
//	fmt.Println(sum(up10, del2, a, b))
//}
//func sum(fa func(int) int, fb func(int) int, a, b int) int {
//	return fa(a) + fb(b)
//}
//func up10(x int) int {
//	return x * 10
//}
//func del2(c int) int {
//	return c / 2
//}

//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Scan()
//	s := strings.Split(strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), " ", ""), ",", "."), ";")
//	fmt.Print(s)
//	result1, err := strconv.ParseFloat(s[0], 64)
//	if err != nil {
//		return
//	}
//	result2, err := strconv.ParseFloat(s[1], 64)
//	if err != nil {
//		return
//	}
//	fmt.Printf("%.4f", result1/result2)
//}

//func adding(x, y string) int64 {
//	var xrune []rune
//	var yrune []rune
//	var a, b string
//	for _, i := range x {
//		xrune = append(xrune, i)
//	}
//	for _, i := range xrune {
//		if unicode.IsDigit(i) == true {
//			a = a + string(i)
//		}
//	}
//	for _, i := range y {
//		yrune = append(yrune, i)
//	}
//	for _, i := range yrune {
//		if unicode.IsDigit(i) == true {
//			b = b + string(i)
//		}
//	}
//	res1, err := strconv.Atoi(a)
//	if err != nil {
//		panic(err)
//	}
//	res2, err := strconv.Atoi(b)
//	if err != nil {
//		panic(err)
//	}
//	res := int64(res1 + res2)
//
//	return res
//}
//
//func adding(firstDirtyNumber, secondDirtyNumber string) int64 {
//	firstNumber, err := strconv.ParseInt(cleaner(firstDirtyNumber), 10, 64)
//	if err != nil {
//		panic(err)
//	}
//
//	secondNumber, err := strconv.ParseInt(cleaner(secondDirtyNumber), 10, 64)
//	if err != nil {
//		panic(err)
//	}
//
//	return firstNumber + secondNumber
//}
//
//func cleaner(s string) string {
//	tmp := make([]rune, 0)
//
//	for _, r := range []rune(s) {
//		if unicode.IsDigit(r) {
//			tmp = append(tmp, r)
//		}
//	}
//
//	return string(tmp)
//}

//func main() {
//	var friendsOfDima []string
//	friendsOfSemyon := make([]string, 3)
//	friendsOfDima = append(friendsOfDima, "Костя", "Семён", "Таня")
//	friendsOfSemyon = append(friendsOfSemyon, "Валера", "Таня", "Дима")
//	friends := map[string][]string{
//		"Dima":   friendsOfDima,
//		"Semyon": friendsOfSemyon,
//		"Костя":  nil,
//	}
//	_, value := friends["Костя"]
//	fmt.Print(value, " ")
//	delete(friends, "Dima")
//	fmt.Print(friendsOfSemyon[3])
//}

//func main() {
//	groupCity := map[int][]string{
//		10:   []string{"Село", "Деревня", "ПГТ"},  // города с населением 10-99 тыс. человек
//		100:  []string{"Город", "Станица"},        // города с населением 100-999 тыс. человек
//		1000: []string{"Мегаполис", "Миллионник"}, // города с населением 1000 тыс. человек и более
//	}
//
//	cityPopulation := map[string]int{
//		"Город":     101,
//		"Станица":   102,
//		"Село":      103,
//		"Мегаполис": 104,
//	}
//
//	for key, _ := range cityPopulation {
//		isInMap := false
//		for _, value := range groupCity[100] {
//			if key == value {
//				isInMap = true
//				break
//			}
//		}
//		if isInMap == false {
//			delete(cityPopulation, key)
//		}
//	}
//}

//func main() {
//	a := make(map[int]int)
//	c := 0
//	for i := 0; i <= 9; i++ {
//		fmt.Scan(&c)
//		if v, ok := a[c]; ok {
//			fmt.Print(v, " ")
//		} else {
//			a[c] = work(c)
//			fmt.Print(v, " ")
//		}
//	}
//}

//func main() {
//	var number, digit int
//	var result string
//	fmt.Println("Введите число:")
//	fmt.Scanf("%d\n", &number)
//	for number > 0 {
//		digit = number % 10
//		number = int(number / 10)
//		result = fmt.Sprint(digit*digit) + result
//	}
//	fmt.Print(result)
//}

//func main() {
//	var x string
//	fmt.Scan(&x)
//	if len(x) <= 1000 {
//		y := []rune(x)
//		for i := 0; i <= len(y)-1; i++ {
//			if usl(y[i]) == true && y[0] <= y[i] {
//				y[0], y[i] = y[i], y[0]
//			}
//		}
//		fmt.Println(string(y[0]))
//	}
//}
//func usl(c rune) bool {
//	if unicode.IsDigit(c) == true {
//		return true
//	} else {
//		return false
//	}
//}

//func main() {
//	var x string
//	fmt.Scan(&x)
//	if len(x) <= 1000 {
//		for idx, elem := range x {
//			if pop(elem) == true && idx == len(x)-1 {
//				fmt.Printf("%s", string(elem))
//			} else if pop(elem) == true {
//				fmt.Printf("%s*", string(elem))
//			}
//		}
//	}
//}
//func pop(c rune) bool {
//	if unicode.Is(unicode.Latin, c) == true {
//		return true
//	} else {
//		return false
//	}
//}

//func main() {
//	fmt.Println(divide(15, 5))
//	fmt.Println(divide(4, 0))
//	fmt.Println("Program has been finished")
//}
//func divide(x, y float64) float64 {
//	if y == 0 {
//		panic("division by zero!")
//	}
//	return x / y
//}

//func main() {
//	var x, y int
//	fmt.Scan(&x, &y)
//	result, error := divide(x, y)
//	if error == nil {
//		fmt.Println(result)
//	} else {
//		fmt.Println("ошибка")
//	}
//}
//func divide(a int, b int) (int, error) {
//	return a / b, nil
//}

//func main() { // не решено
//	var x string
//	fmt.Scan(&x)
//	y := 0
//	for i := range x {
//		if pop(string(x[i])) == true && len(x) >= 5 {
//			y = 1
//		}
//	}
//	if y == 1 {
//		fmt.Println("Ok")
//	} else {
//		fmt.Println("Wrong password")
//	}
//}
//func pop(c string) bool {
//	if strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", c) == true {
//		return true
//	} else {
//		return false
//	}
//}

//func main() {
//	var x, y string
//	fmt.Scan(&x)
//	for i := range x {
//		if strings.Count(x, string(x[i])) <= 1 {
//			y += string(x[i])
//		}
//	}
//	fmt.Print(y)
//}

//func main() {
//	var x string
//	fmt.Scan(&x)
//	for idx, elem := range x {
//		if idx%2 != 0 {
//			fmt.Printf("%c", elem)
//		}
//	}
//}

//func main() {
//	a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
//	a = strings.TrimSuffix(a, "\n")
//	a = strings.ReplaceAll(a, " ", "")
//	a = strings.ToLower(a)
//	b := Reverse(a)
//	if a == b {
//		fmt.Println("Палиндром")
//	} else {
//		fmt.Println("Нет")
//	}
//	fmt.Println(a)
//	fmt.Println(b)
//}
//
//func Reverse(s string) string {
//	n := len(s)
//	runes := make([]rune, n)
//	for _, runee := range s {
//		n--
//		runes[n] = runee
//	}
//	return string(runes[n:])
//}

//func main() {
//	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
//	strings.TrimSuffix(text, "\n")
//	textRunes := []rune(text)
//	if unicode.IsUpper(textRunes[0]) && string(textRunes[len(textRunes)-2]) == "." {
//		fmt.Println("Right")
//	} else {
//		fmt.Println("Wrong")
//	}
//}

//func main() {
//	var txt string = "Слово"
//	txtRunes := []rune(txt)
//	fmt.Printf("\nСрез рун: %v", txtRunes)
//	fmt.Printf("\nИз рун получаем снова текст: %v", string(txtRunes))
//	lenTxtRunes := len(txtRunes)
//	fmt.Printf("\nДлина текста (среза): %v", lenTxtRunes)
//
//	for i, v := range txtRunes {
//		fmt.Printf("\nСимвол %d: %s", i, string(v))
//	}
//
//	txtRunes[lenTxtRunes-1] = rune('а')
//	fmt.Printf("\nЗаменили последнюю букву на \"а\": %v", string(txtRunes))
//
//	txtRunes = append(txtRunes[:4])
//	fmt.Printf("\nОтрезали последнюю букву: %v", string(txtRunes))
//
//	txtRunes = append(txtRunes, 'а', 'р', 'ь')
//	fmt.Printf("\nДобавили три буквы в конец: %v", string(txtRunes))
//}

//import "fmt"
//
//type zxc struct {
//	On    bool
//	Ammo  int
//	Power int
//}
//
//func (testStruct *zxc) Shoot() bool { // if true Ammo - 1
//	if testStruct.On == false || testStruct.Ammo == 0 {
//		return false
//	} else if testStruct.Ammo != 0 && testStruct.On == true {
//		testStruct.Ammo--
//	}
//	return true
//}
//
//func (testStruct *zxc) RideBike() bool { // if true Power - 1
//	if testStruct.On == false || testStruct.Power == 0 {
//		return false
//	} else if testStruct.Power != 0 && testStruct.On == true {
//		testStruct.Power--
//	}
//	return true
//}
//
//func main() {
//	testStruct := &zxc{false, 2, 1}
//
//	fmt.Print(testStruct.Ammo)
//	fmt.Print(testStruct.Power)
//}

//import (
//	"fmt"
//	"math"
//)
//
//type Vector struct {
//	x float64
//	y float64
//	z float64
//}
//
//func createVector(x float64, y float64, z float64) Vector {
//	return Vector{x, y, z}
//}
//
//func length(v Vector) int {
//	return int(math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z))
//}
//
//func main() {
//	a := createVector(6, 3, 2)
//	b := createVector(1, 2, 4)
//	fmt.Print(length(a), " ")
//	fmt.Print(length(b))
//}

//type Circle struct {
//	x, y, r float64
//}
//
//type Person struct {
//	Name string
//}
//
//type Android struct {
//	Person
//	Model string
//}
//
//func (p *Person) Talk() {
//	fmt.Println("Hi, my name is", p.Name)
//}
//
//func main() {
//	c := Circle{1, 2, 3}
//	fmt.Println("Call func with linked arguments do return: ", test1(&c))
//	fmt.Println("after", c)
//	fmt.Println(c.area())
//
//	v := Circle{1, 2, 3}
//	fmt.Println("Call func with not linked arguments do return: ", test2(v))
//	fmt.Println("after", v)
//	fmt.Println(v.area())
//
//}
//
//func test1(circle *Circle) *Circle {
//	circle.r *= 2
//	return circle
//
//}
//
//func test2(circle Circle) Circle {
//	circle.r *= 2
//	return circle
//
//}
//
//func (c *Circle) area() float64 {
//	return math.Pi * c.r * c.r
//}

//func main() {
//	var a, b, result int
//	result = 0
//	fmt.Scan(&a, &b)
//	for result <= a {
//		result = a % 10
//		switch {
//		case result != b:
//			result += result
//			a = a / 10
//		case result == b:
//
//		}
//	}
//	fmt.Println(result)
//}

//func main() {
//	var n, i uint32
//	var fib []uint32
//	_, _ = fmt.Scan(&n)
//	fib = append(fib, 0, 0)
//
//	var a, b uint32 = 1, 1
//	for i = 0; i < n; i++ {
//		a, b = b, a+b
//		fib = append(fib, a)
//	}
//	k := check(fib, n)
//	fmt.Println(k)
//}
//
//func check(nums []uint32, n uint32) int {
//	for k, v := range nums {
//		if v == n {
//			return k
//		}
//	}
//	return -1
//}

//var x int
//fmt.Scan(&x)
//if 1 <= x && x <= 100 {
//array := [100]int{}
//var a, i int
//for i = 0; i < x; i++ {
//fmt.Scan(&a)
//array[i] = a
//}
//for i = 1; i < x; i++ {
//if array[0] > array[i] {
//array[0], array[i] = array[i], array[0]
//}
//}
//pointer := 1
//for i = 1; i <= x; i++ {
//if array[0] == array[i] {
//pointer += 1
//}
//}
//fmt.Println(pointer)
//}

//var x int
//fmt.Scan(&x)
//if 1 <= x && x <= 100 {
//array := [100]int{}
//var a, i int
//for i = 0; i < x; i++ {
//fmt.Scan(&a)
//array[i] = a
//}
//var pointer int
//for i = 0; i < x; i++ {
//if array[i] > 0 {
//pointer++
//}
//}
//fmt.Println(pointer)
//}

//array := [5]int{}
//var a int
//for i := 0; i < 5; i++ {
//fmt.Scan(&a)
//array[i] = a
//if array[0] < array[i] {
//array[0], array[i] = array[i], array[0]
//}
//}
//fmt.Println(array[0])

//var x int
//fmt.Scan(&x)
//if x >= 4 {
//slice := make([]int, x, x)
//for i := 0; i < x; i++ {
//fmt.Scan(&slice[i])
//}
//fmt.Println(slice[3])
//}

//workArray := [10]uint8{}
//indexArray := [6]uint8{}
//inputArray := [16]uint8{}
//outputStr := ""
//
//_, _ = fmt.Scan(&inputArray)
//for i := range inputArray {
//_, _ = fmt.Scan(&inputArray[i])
//}
//for i := 0; i < 10; i++ {
//workArray[i] = inputArray[i]
//}
//for j := 0; j < 6; j++ {
//indexArray[j] = inputArray[j+10]
//}
//workArray[indexArray[0]], workArray[indexArray[1]] = workArray[indexArray[1]], workArray[indexArray[0]]
//workArray[indexArray[2]], workArray[indexArray[3]] = workArray[indexArray[3]], workArray[indexArray[2]]
//workArray[indexArray[4]], workArray[indexArray[5]] = workArray[indexArray[5]], workArray[indexArray[4]]
//for _, v := range workArray {
//outputStr += fmt.Sprintf("%d ", v)
//}
//fmt.Print(outputStr + "type ok")

//
//var x, p, y int // x vklad, p percent, y itog
//i := 0
//fmt.Scan(&x, &p, &y)
//for x < y {
//x = x + (x * p / 100)
//i++
//}
//fmt.Println(i)

//var a int
//for fmt.Scan(&a); a <= 100; fmt.Scan(&a) {
//if a < 10 {
//continue
//}
//fmt.Println(a)
//}

//var (
//	a int
//	n = 0 // counter
//	x = 0 // max
//)
//for fmt.Scan(&a); a != 0; fmt.Scan(&a) {
//if a != 0 {
//if a > x {
//n = 1
//x = a
//} else if a == x {
//n++
//}
//} else {
//fmt.Println(n)
//}
//}
//fmt.Println(n)

//var n, i, a, sum int
//fmt.Scan(&n)
//for i = 1; i <= n; i++ {
//fmt.Scan(&a)
//if a > 9 && a < 100 && a%8 == 0 {
//sum += a
//}
//
//}
//fmt.Println(sum)

//var a, b, c, d, e, f int
//fmt.Scanf("%1d%1d%1d%1d%1d%1d", &a, &b, &c, &d, &e, &f)
//
//if (d + e + f) == (a + b + c) {
//fmt.Print("YES")
//} else {
//fmt.Println("NO")
//}

//var i, a, b int
//fmt.Scan(&i)
//a = (i % 10) + ((i % 100) / 10) + ((i % 1000) / 100)
//b = (i%10000)/1000 + (i%100000)/10000 + (i%1000000)/100000
//if a == b {
//fmt.Println("YES")
//} else {
//fmt.Println("NO")
//}

//if i < 10000 && i > 1000 {
//a = i / 1000
//fmt.Println(a)
//} else if i < 1000 && i > 100 {
//b = i / 100
//fmt.Println(b)
//} else if i < 100 && i > 10 {
//c = i / 10
//fmt.Println(c)
//} else if i > 10000 {
//fmt.Println("NO")
//} else if i == 10000 {
//x = i / 10000
//fmt.Println(x)
//} else {
//fmt.Println(i)
//}

//var i, a, b, c int
//fmt.Scan(&i)
//a = i / 100
//b = i % 100 / 10
//c = i % 10
//if a != b && b != c && c != a {
//fmt.Println("YES")
//} else {
//fmt.Println("NO")
//}

//var i int
//fmt.Scan(&i)
//switch {
//case i > 0:
//fmt.Println("Число положительное")
//case i < 0:
//fmt.Println("Число отрицательное")
//default:
//fmt.Println("Ноль")
//}

//var name string
//var age int
//fmt.Println("vvedite imya i vosrat")
//fmt.Scan(&name, &age)
//
//fmt.Println(name, age)

//import "github.com/mitchellh/mapstructure"
//
//type Point struct {
//	X int `mapstructure:"xx"`
//	Y int `mapstructure:"yy"`
//}
//
//func (p Point) method() {
//	fmt.Println("call Point method")
//}
//func main() {
//	pointsMap := map[string]int{
//		"xx": 123,
//		"yy": 345,
//	}
//	for k, v := range pointsMap {
//		fmt.Println(k)
//		fmt.Println(v)
//	}
//	p1 := Point{}
//	mapstructure.Decode(pointsMap, &p1)
//	fmt.Println(p1)
//
//}

//pointsMap := map[string]Point{
//	"b": {
//		Y: 13,
//		X: 15,
//	},
//}
//otherPpointsMap := make(map[int]Point) // zamena figurnyh skobok
//pointsMap["a"] = Point{
//	X: 1,
//	Y: 2,
//}
////fmt.Println(pointsMap)
////fmt.Println(pointsMap["a"])
//otherPpointsMap[1] = Point{1, 2}
////fmt.Println(otherPpointsMap)
////fmt.Println(otherPpointsMap[1])
//
//var oneMorePointsMap map[string]Point
//if oneMorePointsMap == nil {
//	oneMorePointsMap = map[string]Point{
//		"a": {1, 2},
//	}
//} else {
//	oneMorePointsMap["a"] = Point{1, 2}
//}
////fmt.Println(oneMorePointsMap["a"].X)
////fmt.Println(oneMorePointsMap["a"].Y)
////oneMorePointsMap["a"].method()
//
//key := "a"
//value, ok := oneMorePointsMap[key]
//if ok {
//	fmt.Printf("key=%s exist in map\n", key)
//	fmt.Println(value)
//} else {
//	fmt.Printf("key=%s does not exist in map\n", key)
//	fmt.Println(value)
//}

//arr := []string{"a", "b", "c"}
//for i, l := range arr {
//fmt.Println(i)
//fmt.Println(l)
//
//}

//var snil []int
//fmt.Println(snil, len(snil), cap(snil))
//if snil == nil {
//fmt.Println("slice is nil!")
//}
//	slices
//s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//t := s[:5]
//fmt.Println(t)
//q := s[5:]
//fmt.Println(q)
//w := s[:]
//fmt.Println(w)

//	//animalsArray := [4]string{
//	//	"dog",
//	//	"gir",
//	//	"cat",
//	//	"cat1",
//	//}
//	//var a []string = animalsArray[0:2]
//	//fmt.Println(a)
//	//var b []string = animalsArray[1:3]
//	//fmt.Println(b)
//	//
//	//b[0] = "123"
//	//fmt.Println(a)
//	//fmt.Println(animalsArray)
//
//	//animalsSlice := []string{
//	//	"dog",
//	//	"gir",
//	//	"cat",
//	//	"cat1",
//	//}
//
//}
//
//letters := []string{"a", "b", "c"}
//letters[0] = "1"
//letters = append(letters, "d")
//letters = append(letters, "e", "f")
//fmt.Println(letters)

//createSlice := make([]string, 3)
//fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
//fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
//createSlice[0] = "1"
//createSlice[1] = "2"
//createSlice[2] = "3"
//createSlice = append(createSlice, "4")
//fmt.Println(createSlice)
//fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
//fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
//createSlice = append(createSlice, "5", "6", "7")
//fmt.Println(createSlice)
//fmt.Println(fmt.Sprintf("len: %d", len(createSlice)))
//fmt.Println(fmt.Sprintf("cap: %d", cap(createSlice)))
//
//func arrays() {
//	// arrays
//	var a [2]string
//	a[0] = "hello"
//	a[1] = "world"
//	//a[2] = "ok" - error
//	fmt.Println(a)
//	fmt.Println(a[1])
//
//	numbers := [...]int{1, 2, 3}
//	numbers[2] = 4
//}
//var x int
//fmt.Scan(&x)
//if x >= 4 {
//slice := make([]int, x, x)
//for i := 0; i < x; i++ {
//fmt.Scan(&slice[i])
//}
//fmt.Println(slice[3])
//}
