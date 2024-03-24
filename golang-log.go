// 이 코드가 속한 패키기지가 어딘지 알려줌
// 패키지란 코드를 묶는 단위이다 항상 패키지로 시작하며
// 패키지 명은 아무거나 적어도 되는데
// main은 프로그램 시작점을 포함하는 패키지라는 의미가 있음, 시작 부분은 1군데여야함
// 따라서 go 프로그램은 main 패키지 하나와 여러개의 다른 패키지로 구성됨
// main 함수를 가지고 있는 애다라는 의미
package main

// fmt 라는 패키지를 가져오겠다는 의미
// fmt안의 기능(코드)를 가져와 쓰겠다
import "fmt"
// 패키지 전역변수
var g int = 10
// 상수 -> 메모리가 아닌 코드에 박혀있다, 따라서 값 변경 불가
// type을 안 정할 수 있고 사용될 때 결정됨 const pi = 3.14, var k int = 100 * pi 는 컴파일러에 의해 곱셈이 이미 실행되서 값이 코드에 박혀버림. 즉, 정수결과값 314를 k에 대입하는것이 되어버려 오류 X

// c의 typedef와 같은 역할을 하는 type이 있음
// type ColorType int -> int가 ColorType라는 이름으로도 사용가능
const con int = 10
const con2 float64 = 10.11112
const (
	Red int = iota //0
	Blue int = iota //1
	Green int = iota //2
)
const (
	// 반복되는 type, 값 할당은 생략도 가능
	C1 uint = iota + 1 // 1 = 0 + 1
	C2 // 2 = 1 + 1
	C3 // 3 = 2 + 1
	C4 uint8 = 1 << iota // 1 << 3
)
// 함수명 중 main만 유일하게 프로그램 시작점이라는 의미를 가짐
func main() {
	// go는 변수 선언만 해두면 무조건 0, string 변수면 ""로 초기화가 자동으로 된다
	// fmt라는 외부 패키지안에 포함된 Println이라는 함수(기능)를 쓰겠다
	// print 하고 line 한 줄 띄워라
	fmt.Println("hello world")
	c := sum(3,6)
	fmt.Println("함수 테스트",c)
	//fmt.Printf("%b",c) %b는 go에서만 쓰는 binary 표현

	d, check := divide(9, 3)
	fmt.Println("멀티 반환 함수 테스트1",d, check)
	d2, check2 := divide2(9, 3)
	fmt.Println("멀티 반환 함수 테스트2",d2, check2)
}
// c언어와 다르게 함수 정의 위치는 상관 없음
func sum(a int, b int) int {
	// go에서는 무조건 연산하는 변수끼리의 type이 같아야한다.
	// type이 맞지 않으면 (Ex, int와 int32 일때도 문제가 생김)연산시 형 변환이 필요 이 외는 c언어와 같다
	// var a float = 3.01 -> int32(a) = 3
	// 버퍼오버플로우, 메모리 잃음 현상등 주의 
	// Ex) int16 타입의 변수값 0b1 1000 0001, int8에 대입연산하면
	// 0b1000 0001 값이 할당된다. 즉, c언어와 동일!
	return a+b
}
// 멀티 반환 함수, 반환 값이 여러개인 경우
// go에서는 type이 같은 변수의 경우 공간절약을 위해 타입을 한번만 적도록 할 수 있다.
func divide(a , b int) (int, bool) {
	if b == 0{
		return 0, false
	}
	return a/b, true
}
// 출력 값에 이름 지정해주기 가능
func divide2(a , b int) (result int, check bool) {
	if b == 0{
		result = 0
		check = false
		return
	}
	result = a/b
	check = true
	return
}
// if문
// if 초기문; 조건문 {
//	문장
// }
// 초기문은 말 그대로 변수같은거 초기화임
/*
if filename, success := Uploadfile(); success{
	filename, success := Uploadfile();는 초기문
	success가 조건문
}else
{

}
*/

// switch문
// c와는 다르게 case에 여러 값이 올 수 있고 그 중 하나만 true여도 실행됨
// if와 마찬가지로 초기문 사용가능
/*
day := "thursday"

switch day{
case "monday", "tuesday": -> 요기 실행
	~~~
case "wenesday", "thursday", "friday":
	~~~
}
*/
// go에서는 반복문은 for만 있음 while문이 없다 do..while도 사용법은 c와 같음
/*
for 초기문; 조건문; 후처리 {
	코드 블록
}
or 초기문, 후처리 안쓸거면
for 조건문 {
	코드 블록
}
//무한 루프
for 조건문 {
	코드 블록
}
or
for {
	코드 블록
}
i -> iterator의 약자
for i := 0; i < 10; i++
{
	fmt.Print(i, ",")
}

<아무 이름>: -> 레이블이라 부름 break에서 사용가능
OuterFor:
	for ; a < 9; a++
	{
		for ; b < 9; b++
		{	
			if a*b == 45
				break OuterFor
		}
	}
OuterFor: 에서 부터 처음 만나는 for문 전체를 break해버림
goto문도 go에 존재
되도록 위 두개는 안쓰는걸 추천 -> 스택이 꼬일수있음
*/

// 배열
// var 변수명 [요소 개수]타입
// var t [5]float64
// 배열 선언시 개수는 항상 상수여야 -> c언어와 같음(정확히는 MSVC 컴파일러를 사용하는 윈도우에서 사용법과 같음)
// 배열 끼리 대입연산자 사용가능
/*
a := [3]int{1,2,3}
b := [3]int{100,200,300}
a = b (값 복사, 단 크기가 같아야함)
*/
/*
var t [5]float64 = [5]float64{23.0, 24.14, 21.1, 55.5, 29.9}
dats := [3]string{"a","as","dsa"}
var s = [5]int{1:10, 3:30} -> 특정 인덱스의 값을 지정해줌, s[1] = 10인거임. 그리고 초기값을 안적은 인덱스는 어떤 배열형태든 0으로 초기화
x := [...]int{10, 20, 30} -> [...]은 길이가 뒤에 나오는 값의 길이와 똑같아짐. 여기서는 3으로 정해짐. c언어에서 []안에 아무것도 안적어주는 것과 같음
[]int {10, 20, 30} -> 길이가 늘어날 수 있는 배열(동적 배열)인데 slice라고 함
len(t)하면 5가 나옴 (배열길이)

range 키워드는 배열에서 두개의 값을 반환해줌(slice, string, map등의 자료구조에 사용 가능, 각 요소를 순환해줌)
첫번째는 index, 두번째는 인덱스에 해당하는 값
for i, v := range t {
	fmt.Println(i, v)
}
예를 들어 i를 사용하지 않는다면? i 대신 _로 대체(빈칸 지시자)
for _, v := range t {
	fmt.Println(v)
}
2차원 배열 순회 -> 주소값을 받아온다는 개념으로 이해하면 될듯?
a := [2][5]int{
	{1,2,3,4,5},
	{5,6,7,8,9}, -> 콤파 찍어줘야한대요 만약 닫혀지는 }가 지금처럼 다른 줄에서 닫히면 {5,6,7,8,9}} 이렇게 하면 , 안찍어줘도 됨
}

for _, arr := range a {
	for _, v := range arr {
		fmt.Print(v, "")
	}
	fmt.Println()
}
*/

// 구조체
/*
type 타입명 struct {
	필드명 타입
}

type Student struct{
	Name string
	Class int
}

var a Student
a.Name = "이지지"
a.Class = 2
fmt.Print(a)
var a Student = {
	"이지지",
 		2,
 } == var a Student = {"이지지", 2} 

 var a Student = {Class: 2} -> 특정 변수만 초기화하고 싶은 경우 
 구조체도 대입연산자로 모든 필드 값(멤버변수)이 복사된다
*/

// 포인터 -> 모든 프로그래밍언어에서 사용하고 있으나 그걸 노출 시킨건 c/c++언어 이것보단 덜 노출 된게 golang
/*
포인터는 메모리 주소를 값으로 갖는 타입

var a int
var p *int
p = &a // a의 메모리 주소를 포인터 변수 p에 대입

*p = 20

포인터 변수의 기본값은 nil(null)
if p != nil
메모리 주소값은 8byte (64bit 운영체제)
구조체를 보면
var data Data
var p * Data = &data
p.value = 999나 (*p).value가 같다

var data Data
var p *Data = &data
는
var p *Data = &Data{} 로 대체 가능(Data 구조체를 만들어 주소를 반환합니다)
*/

/*
인스턴스(instance)는 메모리에 할당된 데이터의 실체

var data Data
var p *Data = &data
or
var p *Data = &Data{}

Data 인스턴스 하나가 만들어졌고 포인터 변수 p가 가리킵니다. 인스턴스 갯수 1개

var p1 *Data = &Data{}
var p2 *Data = p1
var p3 *Data = p1

한 Data 인스턴스를 세 포인터 변수가 가리킵니다. -> 인스턴스 갯수 1개

var data1 *Data
var data2 *Data = data1
var data3 *Data = data1

data1,2,3는 값만 같을 뿐 서로 다른 인스턴스
*/

/*
new() 내장함수 -> 같은 문법임

p1 := &Data{} // &를 사용하는 초기화 -> p1 := &Data{class:"1"} 이런식으로 필드값도 초기화 가능 
var p2 = new(Data) // new()를 사용하는 초기화 -> 필드값 초기화 불가능
*/

/* 인스턴스는 언제 사라지나?

func TestFunc()
{
	u := &User{} //u 포인터 변수를 선언하고 인스턴스를 생성
	u.Age = 30
	fmt.Println(u)
} // 내부 변수 u는 사라지고 인스턴스도 사라집니다


//c/c++에서는 newus함수가 끝나면서 지역변수(스택메모리에서 지역변수 pop, 함수 pop해버림)가 사라지기 때문에 에러가 생기지만 golang은 escape analasying이라고 컴파일러가 분석해서 이렇게 함수가 끝날 반환값도 사라지는 경우 stack이 아니라 heap 에다 만들어버림 -> 따라서 쓰임이 다하면 사라짐

func newus(name string, age int) *User
{
	var u := User{name, age}
	return &u
}
func main()
{
	userpointer := newus("AAA", 23)
	fmt.Println(userpointer)
}

*/

/* 문자열

파이썬에서 ''' Line1
Line2
Line3 ''' 을 쓰면 개행문자 \n을 따로 써주지 않아도 그대로 출력되는 것 처럼 golang에서는 백쿼터  ` `로 묶어주면 똑같은 효과를 일으킴
" "로 묶으면 \n 등 써줘야 개행이 됨

len()은 문자열 바이트 반환
str := "hello 월드" -> 한글은 3바이트로 표현된다, 따라서 len(str) = 12
str[7] = 이상한 값이 나옴(월을 나타내는 3바이트 중 중간의 1바이트 값을 나타냄)
for i := 0; i < len(str); i++
{
	fmt.Printf("타입: %T 값:%d 문자값: %c\n", str[i], str[i], str[i])
}
따라서 한글을 순수 인덱스로 접근하면 utf-8은 출력 불가능

해결방법
str := "hello 월드"
arr := []rune(str)

for i := 0; i < len(arr); i++
{
	fmt.Printf("타입: %T 값:%d 문자값: %c\n", arr[i], arr[i], arr[i])
}

[]rune -> slice인데 동적배열이다 (rune은 int32의 별칭타입, 1칸이 4바이트), 즉 arr := []rune(str)는 타입 변환을 해준 것
한글자당 4바이트씩 차지
'H' -> 4바이트를 차지
len(arr) -> 배열의 갯수가 나옴 여기서는 값이 8

아니면 range를 이용하면 한 글자씩 잘 출력

for _, v := range str
{
	fmt.Printf("타입: %T 값:%d 문자값: %c\n", v, v, v)
}

문자열은 + 연산자로 붙일 수 있음(+만 지원)
문자열 비교 가능
"Hello" == "World" -> false가 뜸
>,<, <=, >= 대소 비교 가능, 사전식 비교 - 대문자가 더 작다 ('A'-'Z': 65-90인 아스키 코드 참고)

str2 = str 가능 (string 은 포인터형태로 가리키고 있음, 즉 string은 16byte 자료형(struct, data 주소를 가리키는 변수와 len을 나타내는 변수 각각 8byte 멤버 변수가 있음)임)

문자열 일부만 수정할 수없다!, 문자열 불변
var str string = "Hello World"
str = "How are you?" 이건 가능
str[2] = 'a' 이건 불가능

굳이 가능하게 하려면

var str string = "Hello World"
var slice []byte = []byte(str) // 타입 변환시 값의 복사가 일어남

slice[2] = 'a' // 영어들이고 변경가능

fmt.Println(str) //str은 바뀌지 않았다
fmt.Printf("%s\n", slice)

str이 가르키는 메모리와 slice가 가르키는 메모리는 다르다!!!
문자열은 기본적으로 포인터를 이용해 읽어오는 c언어 방식과 같다

*/

/* 슬라이스
슬라이스는 Go에서 제공하는 동적 배열 타입

정적 - static -> compile tume, build time 즉, 코드를 기계어로 바꿀때 결정된다, 즉 실행도중 절대 바뀌지 않는 값 ex const
동적 - dynamic -> runtime 즉, 프로그램이 실행 중 결정된다

슬라이스 선언
var (변수명은 아무거나, 여기서는 slice라고 하겠음) []int
or
slice := []int{} //주의 var array = [...]int{1,2,3}는 사이즈 고정된 배열

make()를 이용한 초기화
var slice = make([]int, 3) // make는 go의 내장함수, slice,map 타입을 만들때 주로 사용, 3개의 요소를 가진 슬라이스를 만듦
// slice := []int{0,0,0}과 같다

slice[1] = 5

슬라이스 요소 추가 - append() -> 슬라이스에 요소를 추가한 새로운 슬라이스를 만들어서 반환

slice2 := append(slice, 4) or
slice = append(slice, 4)

여러개 한번에 추가
slice = append(slice, 3, 4, 5, 6, 7)

사실은 슬라이스는 Go에서 제공하는 배열을 가리키는 포인터 타입이다~ - https://www.youtube.com/watch?v=z-_6o7WYkiE&list=PLy-g2fnSzUTBHwuXkWQ834QHDZwLx6v6j&index=25 16:40~
type SliceHeader struct {
	Data unitptr // 실제 배열을 가리키는 포인터
	Len int // 요소 개수
	Cap int // 실제 배열의 길이
}
slice는 24바이트 구조체
var slice = make([]int, 3, 5) //길이 5인 배열을 만들고 길이 3만 쓰고있다

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println(array) // 3번째 인덱스값 안 변함
	fmt.Println(slice) // 3번째 인덱스 200
}

C/C++은 array가 변하는데 Go는 안 변함(배열도 변수처럼 대입연산으로 복사함) 인자로 넘긴 값이 복사됨
slice만 포인터 주소값과 길이만 복사가 되는 것임! 무조건 24바이트만 복사됨
근데 append가 함수 내에서 사용될 때는 값이 안바뀔수 있다. - https://www.youtube.com/watch?v=zcUqV5xk-So&list=PLy-g2fnSzUTBHwuXkWQ834QHDZwLx6v6j&index=26 0 ~ 13:00
append에선 새로운 슬라이스를 생성하기도 하기 때문이다! 주의!!! 
해결 방법은 
1. 함수 인자를 포인터로 받고 주소 넘겨주기
2. 새로운 슬라이스를 반환하도록 return 해주기

슬라이싱

array[startindex:endindex] startindex <= ~ < endindex , 파이썬의 슬라이싱과 다르다
array[시작인덱스: 끝인덱스: 최대인덱스(cap사이즈)]
*/

/* method (메서드) == 함수
매서드는 타입에 속한 함수

함수 -> 독립적
메서드 -> 타입 종속 (타입안에 함수가 있다~)

type Rabbit struct {
	...
}

메서드 선언
func (r Rabbit) info() int {
	return r.width * r.height
}
r Rabbit -> 리시버, 이걸 적어서 애가 어떤 타입에 속하는지 알려주는것
info() -> 메서드명
리시버는 모든 패키지 지역 타입이 가능하다. 구조체, 별칭 타입 등

단순하게 보면 메서드는 func f(r Rabbit, a int) 와 func (r Rabbit) f(a int) 동격이다. 기능은 같다. 근데 method는 c++의 class같은 느낌이랄까

이렇게도 가능, 이 함수 ff는 MyInt 리시버에 포함되어있는 함수가 된다.
type MyInt int

func (m MyInt) ff(){

}

예제
type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{ 100 } // a *account

	// a.balance는 100

	withdrawFunc(a, 30)

	a.withdrawMethod(30) //a라는 객체를 리시버로 받는 메서드를 호출	
}

객체(Object)란 데이터(State)와 기능(Function)을 묶은 것이다. -> 그냥 관리를 쉽게하기 위해 의존성을 낮추고 결합도를 높여주는거임(class 같은거임)
절차 중심에서 관계 중심으로

클래스간 상호 작용 - https://www.youtube.com/watch?v=-ijeABV8vLU&list=PLy-g2fnSzUTBHwuXkWQ834QHDZwLx6v6j&index=27
22분 정도부터

포인터 타입 메서드 vs 값 타입 메서드

type account struct {
	balance int
	firstname string
	lastname string
}

func (a1 *account) widthdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) widthdrawValue (amount int) {
	a2.balance -= amount
}

func (a2 account) widthdrawValue2 (amount int) account {
	a2.balance -= amount
	return a2
}
func main() {
	var mainA *account = &account{100, "joe", "park"}
	mainA.withdrawPointer(30) //(&mainA).withdrawValue(30) 이어야하는거 아냐? -> go에서는 귀차나서 지가 알아서 변환 해줌
	fmt.Println(mainA.balance)

	mainA.widthdrawValue(20) // (*mainA).withdrawValue(20) 이어야하는거 아냐? -> go에서는 귀차나서 지가 알아서 변환 해줌
	fmt.Println(mainA.balance)

	*mainA = mainA.widthdrawValue(20)
	fmt.Println(mainA.balance)
}

출력:
70
70 //mainA값이 복사된 a2의 balance의 값을 깍음 mainA는 변화 X
50

49분 ~
class의 상속같은 느낌의 사용법(상속이랑은 다르다, go는 상속이 없고 그냥 임베디드 필드하면 상속처럼 쓰는 것으로 보임)
*/

/*
인터페이스

구체화된 객체(Concrete object)가 아닌 추상화된 상호작용으로 관계를 표현


*/