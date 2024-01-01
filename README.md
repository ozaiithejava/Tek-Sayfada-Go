# Tek Sayfa Golang Syntax

![Golang Logo](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png)


#### Değişkenler:
```Go
package main

import "fmt"

func main() {
    // Boolean
    var isActive bool // Değer atanmadığında varsayılan değer "false" olur
    fmt.Println("isActive:", isActive)

    // Tam Sayılar
    var age int = 25
    var temperature = -10 // Tür belirtilmezse, derleyici değerden tipi çıkarır
    count := 100          // Kısa tanımlama şekli, sadece fonksiyon içinde kullanılabilir
    var distance int64    // Büyük tam sayıları temsil etmek için int64 kullanılabilir

    fmt.Println("Yaş:", age)
    fmt.Println("Sıcaklık:", temperature)
    fmt.Println("Sayı:", count)
    fmt.Println("Mesafe:", distance)

    // Ondalık Sayılar
    var price float32 = 29.99
    var pi = 3.14159265359 // Float64 türü, daha fazla hassasiyet sağlar

    fmt.Println("Fiyat:", price)
    fmt.Println("Pi:", pi)

    // Karmaşık Sayılar
    var complexNum complex64 = 1 + 2i
    imaginaryNum := 5i // Kısa tanımlama şekli ile karmaşık sayı
    fmt.Println("Karmaşık Sayı:", complexNum)
    fmt.Println("Sanal Sayı:", imaginaryNum)
}
```
#### Daha çok değisken:

```Go
package main

import "fmt"

func main() {
    // int türleri
    var myInt int = 42
    var myInt8 int8 = 127
    var myInt16 int16 = -32768
    var myInt32 int32 = -2147483648
    var myInt64 int64 = 9223372036854775807

    // uint türleri
    var myUint uint = 42
    var myUint8 uint8 = 255
    var myUint16 uint16 = 65535
    var myUint32 uint32 = 4294967295
    var myUint64 uint64 = 18446744073709551615

    // float türleri
    var myFloat32 float32 = 3.14
    var myFloat64 float64 = 3.14

    // complex türleri
    var myComplex64 complex64 = 1 + 2i
    var myComplex128 complex128 = 1 + 2i

    // Değerleri yazdırma
    fmt.Println("int:", myInt)
    fmt.Println("int8:", myInt8)
    fmt.Println("int16:", myInt16)
    fmt.Println("int32:", myInt32)
    fmt.Println("int64:", myInt64)

    fmt.Println("uint:", myUint)
    fmt.Println("uint8:", myUint8)
    fmt.Println("uint16:", myUint16)
    fmt.Println("uint32:", myUint32)
    fmt.Println("uint64:", myUint64)

    fmt.Println("float32:", myFloat32)
    fmt.Println("float64:", myFloat64)

    fmt.Println("complex64:", myComplex64)
    fmt.Println("complex128:", myComplex128)
}

```

#### If Else

if ve else yapısı belirli koşulların sağlanması veya sağlanmaması durumunda devreye giren kod bloklarıdır

```Go
package main

import "fmt"

func main() {
    // Basit if-else
    age := 25

    if age < 18 {
        fmt.Println("Yaşınız 18'den küçük.")
    } else {
        fmt.Println("Yaşınız 18 veya daha büyük.")
    }

    // Basit else-if
    score := 75

    if score >= 90 {
        fmt.Println("Notunuz A")
    } else if score >= 80 {
        fmt.Println("Notunuz B")
    } else if score >= 70 {
        fmt.Println("Notunuz C")
    } else {
        fmt.Println("Notunuz D veya daha düşük")
    }

    // Defer kullanımı (finally yerine)
    defer fmt.Println("Bu kod bloğu çalıştıktan sonra defer çalışacak")
}
```

#### switch case

Go dilinde switch ifadesi, belirli bir değeri kontrol etmek ve farklı durumları ele almak için kullanılır.

```Go
package main

import "fmt"

func main() {
    dayOfWeek := 3

    switch dayOfWeek {
    case 1:
        fmt.Println("Pazartesi")
    case 2:
        fmt.Println("Salı")
    case 3:
        fmt.Println("Çarşamba")
    case 4:
        fmt.Println("Perşembe")
    case 5:
        fmt.Println("Cuma")
    case 6:
        fmt.Println("Cumartesi")
    case 7:
        fmt.Println("Pazar")
    default:
        fmt.Println("Geçersiz gün")
    }

    // switch ifadesiyle tip kontrolü
    var x interface{} = 42

    switch x.(type) {
    case int:
        fmt.Println("x bir tam sayı")
    case float64:
        fmt.Println("x bir ondalık sayı")
    case string:
        fmt.Println("x bir karakter dizisi")
    default:
        fmt.Println("x belirtilen tipte değil")
    }
}
```

#### Panic / Defer / Recover

panic:

panic ifadesi, ciddi hataları temsil eder ve bir programın çalışmasını durdurur.

Bir fonksiyon içinde panic ifadesi kullanıldığında, o fonksiyonun çalışması anında sona erer.

panic oluşturulan bir hata mesajı veya değeri içerebilir.

```Go
panic("Bu bir hata!")
```

recover:

recover ifadesi, bir panic durumunu yakalamak ve programın çökmesini önlemek için kullanılır.

recover genellikle defer ile birlikte kullanılır.

Bir fonksiyon içinde recover çağrıldığında ve o fonksiyon bir panic durumunda çağrıldığında, recover çağrısı panic ifadesini yakalar ve programın normal akışına geri döndürür

```Go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Panic yakalandı:", r)
    }
}()
```

defer:

defer ifadesi, bir fonksiyonun sonunda belirtilen ifadelerin çalışmasını erteleyen bir mekanizmadır.

Genellikle temizlik işlemleri, kaynakların serbest bırakılması gibi işlemlerde kullanılır.

defer ifadeleri, fonksiyonun sonunda çalıştırılır, ancak bu ifadelerin sırası, tanımlandıkları sırayla değil, fonksiyonun sonunda çalıştırılacak şekilde belirlenir.

```Go
defer fmt.Println("Bu ifade fonksiyonun sonunda çalışacak")
```

##### Çalışır örnek:

```Go
package main

import "fmt"

func main() {
    performOperation()
}

func performOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Hata yakalandı:", r)
        }
    }()

    fmt.Println("Bir işlem yapılıyor.")

    // Bir hata durumu simüle edelim
    simulateError()
    
    // Bu kod asla çalışmaz, çünkü yukarıdaki hata nedeniyle fonksiyon durur
    fmt.Println("Bu kod asla çalışmaz")
}

func simulateError() {
    // Hata durumu yarat
    panic("Bu bir hata!")
}
```

#### Fonksiyonlar
Basit bir hello world ile başlayalım

```Go
package main

import "fmt"

func helloWorld() {
    fmt.Println("Hello, World!")
}

func main() {
    helloWorld()
}

```

Parametre Alan Fonskiyon

```Go
package main

import "fmt"

func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

func main() {
    greet("Alice")
    greet("Bob")
}
```

Birden Fazla Parametre Alan Fonskiyon

```Go
package main

import "fmt"

func getGreeting(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

func main() {
    greeting := getGreeting("Alice")
    fmt.Println(greeting)
}
```

Birden fazla return alan fonksiyon

```Go
package main

import "fmt"

func getFullName(firstName, lastName string) (string, string) {
    fullName := fmt.Sprintf("%s %s", firstName, lastName)
    return fullName, fmt.Sprintf("Hello, %s!", fullName)
}

func main() {
    fullName, greeting := getFullName("John", "Doe")
    fmt.Println(greeting)
    fmt.Println("Full Name:", fullName)
}
```

Public / Private fonksiyon

```Go
package main

import (
	"fmt"
	"github.com/yourusername/yourpackage/utils"
)

// ExportedFunction bir paket içinde ve diğer paketlerden erişilebilir
func ExportedFunction() {
	fmt.Println("Bu fonksiyon diğer paketlerden erişilebilir.")
}

// nonExportedFunction sadece bu paket içinden erişilebilir
func nonExportedFunction() {
	fmt.Println("Bu fonksiyon sadece bu paket içinden erişilebilir.")
}

func main() {
	ExportedFunction() // Diğer paketten erişilebilen fonksiyon
	utils.CallNonExported() // Bu paket içinden erişilebilen fonksiyon
}
```

#### Arrayler

```Go
package main

import "fmt"

func main() {
    // String tipinde 3 elemanlı bir dizi
    var names [3]string
    names[0] = "Alice"
    names[1] = "Bob"
    names[2] = "Charlie"

    fmt.Println("Names:", names)
}
```
Dizi uzunlugunu otomatik belirleyen cins

```Go
package main

import "fmt"

func main() {
    // Elemanları otomatik belirlenen bir dizi
    numbers := [...]int{1, 2, 3, 4, 5}

    fmt.Println("Numbers:", numbers)
}
```

Döngü ile diziyi gezmek

```Go
package main

import "fmt"

func main() {
    // String tipinde 3 elemanlı bir dizi
    var names [3]string
    names[0] = "Alice"
    names[1] = "Bob"
    names[2] = "Charlie"

    // Diziyi döngü ile gezmek
    for i := 0; i < len(names); i++ {
        fmt.Println(names[i])
    }
}
```

#### Dilimler

Dizilerden daha esnek yapılardır

```Go
package main

import "fmt"

func main() {
    // Elemanları otomatik belirlenen bir dilim
    numbers := []int{1, 2, 3, 4, 5}

    // Dilimi döngü ile gezmek
    for _, number := range numbers {
        fmt.Println(number)
    }
}
```

çok boyutlu diziler

```Go
package main

import "fmt"

func main() {
    // 2 boyutlu bir dizi
    matrix := [2][3]int{
        {1, 2, 3},
        {4, 5, 6},
    }

    fmt.Println("Matrix:", matrix)
}
```

#### Mapler

Basit mapler

```Go
package main

import "fmt"

func main() {
    // String tipinde anahtarlar ve int tipinde değerler içeren bir map
    ages := map[string]int{
        "Alice":   25,
        "Bob":     30,
        "Charlie": 22,
    }

    fmt.Println("Bob'un yaşı:", ages["Bob"])
}
```

Eleman ekleme / silme

```go
package main

import "fmt"

func main() {
    ages := map[string]int{
        "Alice":   25,
        "Bob":     30,
        "Charlie": 22,
    }

    // Yeni bir eleman eklemek
    ages["David"] = 28

    // Bir elemanı silmek
    delete(ages, "Bob")

    fmt.Println("Ages:", ages)
}
```

Değeri kontrol etmek

```Go
package main

import "fmt"

func main() {
    ages := map[string]int{
        "Alice":   25,
        "Bob":     30,
        "Charlie": 22,
    }

    // "Alice" anahtarının varlığını kontrol etme
    if age, exists := ages["Alice"]; exists {
        fmt.Println("Alice'in yaşı:", age)
    } else {
        fmt.Println("Alice bulunamadı.")
    }
}
```

Map içinde Map(nested)

```Go
package main

import "fmt"

func main() {
    users := map[string]map[string]string{
        "Alice": {
            "FirstName": "Alice",
            "LastName":  "Doe",
            "Age":       "25",
        },
        "Bob": {
            "FirstName": "Bob",
            "LastName":  "Smith",
            "Age":       "30",
        },
    }

    fmt.Println("Alice'in soyadı:", users["Alice"]["LastName"])
}
```

Map fonksiyonlari

```Go
package main

import "fmt"

func main() {
    colors := map[string]string{
        "Red":   "#FF0000",
        "Green": "#00FF00",
        "Blue":  "#0000FF",
    }

    // Map'in uzunluğunu öğrenme
    fmt.Println("Color sayısı:", len(colors))

    // Bir elemanı silme
    delete(colors, "Green")

    fmt.Println("Colors:", colors)
}
```

Parametre olarak kullanma

```Go
package main

import "fmt"

func printMap(data map[string]string) {
    for key, value := range data {
        fmt.Printf("%s: %s\n", key, value)
    }
}

func main() {
    colors := map[string]string{
        "Red":   "#FF0000",
        "Green": "#00FF00",
        "Blue":  "#0000FF",
    }

    printMap(colors)
}
```

Dongu ile gezme

```Go
package main

import "fmt"

func main() {
    scores := map[string]int{
        "Alice":   95,
        "Bob":     80,
        "Charlie": 92,
    }

    // Map üzerinde döngü ile gezme
    for name, score := range scores {
        fmt.Printf("%s'in notu: %d\n", name, score)
    }
}
```

Map ve struct kullanımı


```Go
package main

import "fmt"

// User struct'ında bir map kullanımı
type User struct {
    ID      int
    Details map[string]string
}

func main() {
    user := User{
        ID: 1,
        Details: map[string]string{
            "FirstName": "Alice",
            "LastName":  "Doe",
            "Age":       "25",
        },
    }

    fmt.Println("User:", user)
}
```

###### Örnek

```Go
package main

import "fmt"

func main() {
    operations := map[string]func(int, int) int{
        "add":      add,
        "subtract": subtract,
    }

    result := performOperation(10, 5, "add", operations)
    fmt.Println("Sonuç (Toplama):", result)

    result = performOperation(10, 5, "subtract", operations)
    fmt.Println("Sonuç (Çıkarma):", result)
}

func add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func performOperation(a, b int, operation string, operations map[string]func(int, int) int) int {
    if fn, exists := operations[operation]; exists {
        return fn(a, b)
    }
    return 0
}
```

#### Arayüzler(interface)

```Go
type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

Empty interface

```go
var emptyInterface interface{}
```

Embeded interface

```go
type Writer interface {
    Write([]byte) error
}

type Closer interface {
    Close() error
}

type WriterCloser interface {
    Writer
    Closer
}
```

Type Assertion ve Type Switch 


```Go
var myInterface interface{} = "Hello, Golang!"

// Type assertion
if str, ok := myInterface.(string); ok {
    fmt.Println("Bu bir string:", str)
}

// Type switch
switch value := myInterface.(type) {
case string:
    fmt.Println("Bu bir string:", value)
case int:
    fmt.Println("Bu bir integer:", value)
}
```

İnterface İmplementasyonu

```Go
type Logger interface {
    Log(message string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
    fmt.Println("Log:", message)
}
```

Boş Arayüz ve Reflection

```Go
import "fmt"

func main() {
    var value interface{} = 42

    // Reflection kullanarak tür bilgisini alma
    valueType := reflect.TypeOf(value)
    fmt.Println("Tür:", valueType)

    // Tür bilgisini kullanarak yeni bir değer oluşturma
    newValue := reflect.New(valueType).Elem().Interface()

    fmt.Println("Yeni Değer:", newValue)
}

```


#### Hata türleri

```Go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("Sıfıra bölme hatası")
    }
    return a / b, nil
}
```

Hata kontrolü

```Go
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Hata:", err)
} else {
    fmt.Println("Sonuç:", result)
}
```

Hatada özel ifadeler kullanma


```Go
package main

import (
    "fmt"
    "github.com/pkg/errors"
)

func customError() error {
    return errors.New("Özel bir hata oluştu")
}

func main() {
    err := customError()
    wrappedErr := errors.Wrap(err, "Ana hataya özel bilgi ekleme")
    fmt.Println(wrappedErr)
}
```

Özel Hata türleri

```Go
package main

import "fmt"

// CustomError özel bir hata türüdür
type CustomError struct {
    Message string
}

// Error metodu CustomError türünün error interface'ini implemente eder
func (e CustomError) Error() string {
    return e.Message
}

func main() {
    err := CustomError{Message: "Bu özel bir hata"}
    fmt.Println("Hata:", err)
}
```

Defer kullanarak Hata işleme

```Go
package main

import "fmt"

func main() {
    file, err := openFile("example.txt")
    if err != nil {
        fmt.Println("Dosya açma hatası:", err)
        return
    }
    defer file.Close()

    // Dosya üzerinde işlemler
}

func openFile(filename string) (*File, error) {
    // Dosya açma işlemi
    // ...

    return file, err
}
```

#### Pointer(obj *Obj)

Bir pointer, bir değişkenin bellek adresini tutan bir türdür.

```Go
var x int = 42
var ptr *int = &x
```

& operatörü, bir değişkenin bellek adresini alır.

*int, bir int tipindeki değerin adresini tutan bir pointer tipidir.

##### Kullanımı:

```Go
var x int = 42
var ptr *int = &x

fmt.Println("Değer:", x)
fmt.Println("Pointer ile Değer:", *ptr)
```
*ptr ile pointer'ın gösterdiği bellek adresindeki değeri alabiliyoruz!

Fonksiyonda kullanma:

```go
func double(x *int) {
    *x *= 2
}

func main() {
    value := 10
    double(&value)
    fmt.Println("Değer:", value)
}
```

Fonksiyonlara pointer geçirerek, fonksiyon içinde orijinal değeri değiştirebilirsiniz

Pointers, bellek yönetiminde önemli bir rol oynar. Bellek adreslerini doğru kullanmak, kaynak yönetimi ve performans açısından önemlidir.

Yeni değer oluşturma

```Go
var ptr *int = new(int)
*ptr = 5
```

```Go 
new()
``` 
fonksiyonu ile yeni bir değer oluşturup, pointer üzerinden bu değeri kullanabilirsiniz.

nil(null) ve güvenlik

```go
var ptr *int
fmt.Println("Nil Pointer:", ptr)
```

#### OOOP

```Go
// hello.go

package main

import "fmt"

// Selamla fonksiyonu bir selam mesajı yazdırır
func Selamla(isim string) {
    fmt.Println("Merhaba, " + isim + "!")
}
```

```go
// main.go

package main

import "./hello" // hello.go dosyasının bulunduğu klasörü düzeltmeye göre ayarlayın

func main() {
    // Selamla fonksiyonunu kullanma örneği
    hello.Selamla("Ahmet")
}
```

#### Struct 

```go
package main

import "fmt"

// Person struct'ı
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func main() {
    // Yeni bir Person örneği oluştur
    p := Person{FirstName: "John", LastName: "Doe", Age: 30}

    // Struct özelliklerine erişim
    fmt.Println(p.FirstName, p.LastName, p.Age)
}
```

#### Arena

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Bellek kullanımını izleme
	go func() {
		for {
			printMemoryUsage()
			time.Sleep(1 * time.Second)
		}
	}()

	// Büyük miktarda bellek tüketen işlem
	for i := 0; i < 1000000; i++ {
		_ = make([]byte, 1024)
	}

	// Bellek profili al
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("\nToplam Ayrılan Bellek: %v MB\n", memStats.TotalAlloc/1024/1024)
}

func printMemoryUsage() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Ayrılan Bellek: %v MB\n", memStats.Alloc/1024/1024)
}

```

#### Async

```Go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Goroutine kullanarak asenkron işlem
	go asyncFunction("Merhaba, ")

	// Ana program devam eder
	fmt.Println("Bu kısım hemen çalışır.")

	// Bir süre bekleyelim, aksi takdirde goroutine'in işlemesi bitmeden program sonlanabilir
	time.Sleep(2 * time.Second)
}

func asyncFunction(message string) {
	// Goroutine içinde asenkron olarak çalışacak işlem
	time.Sleep(1 * time.Second)
	fmt.Println(message + "Dünya!")
}
```

#### iota

İota, özel bir const ifadesidir ve bir blok içindeki her bir const değişkenine otomatik olarak artan sayısal değerleri atar. Eğer iota kullanılmışsa ve bir değer belirtilmemişse, iota'nın değeri sıfırdan başlar ve her bir const değişkeni için bir artar.

```go
package main

import "fmt"

const (
	Pazartesi = iota // 0
	Sali             // 1
	Carsamba         // 2
	Persembe         // 3
	Cuma             // 4
	Cumartesi        // 5
	Pazar            // 6
)

func main() {
	fmt.Println(Pazartesi, Sali, Carsamba, Persembe, Cuma, Cumartesi, Pazar)
}
```

