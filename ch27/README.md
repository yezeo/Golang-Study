# Chapter 27 객체지향 설계 원칙 SOLID

## 27.1 객체지향 설계 5가지 원칙 SOLID

- 단일 책임 원칙 single responsibility principle, SRP
- 개방-폐쇄 원칙 open-closed principle, OCP
- 리스코프 치환 원칙 liskov substitution principle, LSP
- 인터페이스 분리 원칙 interface segregation principle, ISP
- 의존 관계 역전 원칙 dependency inversion principle, DIP

반드시 지켜야 하는 의무사항은 아니지만 이 원칙들에 입각해서 설계를 하면 더 좋은 설계를 할 수 있다.

**왜 설계를 잘해야 하는가?**

설계 : 프로그램 코드를 이루는 각 모듈 간 의존 관계를 정의하는 것

현대 프로그래밍은 과거에 비해서 매우 복잡하다. 과거 한두 명 → 현재 수십~수백 명이 각자의 맡은 바 코드를 구현한다. 맡은 바 코드를 모듈 단위로 볼 수 있는데, 이러한 모듈이 모여 프로그램을 이루다 보니까 설계를 잘하지 않으면 많은 문제가 발생할 수 있다.

**나쁜 설계**

- 경직성 rigidity
    
    모듈 간의 결합도가 너무 높아서 코드를 변경하기 어려운 구조. 때론 모듈 간 의존 관계가 거미줄처럼 얽혀 있어서 어디부터 손대야 할지 모를 정도로 복잡한 구조를 갖기도 한다.
    
- 부서지기 쉬움 fragility
    
    한 부분을 건드렸더니 다른 부분까지 망가지는 경우. 언제 어떤 부분이 망가질지 모르기 때문에 프로그램을 변경하기가 매우 힘들다.
    
- 부동성 immobility
    
    코드 일부분을 현재 어플리케이션에서 분리해서 다른 프로젝트에도 쓰고 싶지만 모듈 간 결합도가 너무 높아서 옮길 수 없는 경우. 코드 재사용률이 급격히 감소하므로 (같거나) 비슷한 기능을 매번 새로 구현해야 한다.
    

⇒ 상호 결합도가 매우 높고 응집도가 낮다.

‘상호 결합도가 높다’ : 모듈이 서로 강하게 결합되어 있어서 떼어낼 수 없다는 뜻. 경직성 증가, 그로 인해 한 모듈의 수정이 다른 모듈로 전파되어 예기치 못한 문제가 생기고 코드 재사용성을 낮추게 된다.

‘응집도가 낮다’ : 하나의 모듈이 스스로 자립하지 못한다는 뜻. 하나의 모듈이 스스로 완성되지 못하고 다른 모듈에 의존적인 관계를 가지고 있는 경우.

**좋은 설계**

좋은 설계 : 나쁜 설계 요소가 없는 설계

⇒ 상호 결합도가 낮고 응집도가 높은 설계.

상호 결합도가 낮다 → 모듈을 쉽게 떼어내서 다른 곳에 사용할 수 있다

모듈 간 독립성이 있다 → 한 부분을 변경하더라도 다른 모듈에 문제를 발생시키지 않는다

⇒ 자연스럽게 모듈 완성도가 높아져서 응집도가 높아진다

## 27.2 단일 책임 원칙

**정의**

- “모든 객체는 책임을 하나만 져야 한다.”

**이점**

- 코드 재사용성을 높여준다.

```go
type FinanceReport struct {
	report string
}

func (r *FinanceReport) SendReport(email string) {
	...
}
```

FinanceReport는 단일 책임 원칙을 위배했다! → 회계 보고서라는 책임 + 보고서를 전속하는 책임

```go
type MarketingReport struct {
	report string
}

func (r *MarketingReport) SendReport(email string) {
	...
}
```

MarketingReport는 FinanceReport의 SendReport() 메서드 사용불가! 

구현이 비슷한 메서드를 MarketingReport 객체 안에 또 만들기? → 보고서 종류가 늘어난다면?

**단일 책임 원칙에 입각한 설계**

![0](.\img\0.png)

FinanceReport는 경제 보고서만 책임, ReportSender 보고서 전송이라는 책임 하나만

향후 다른 보고서가 나오더라도 Report 인터페이스만 구현하면 ReportSender를 그대로 이용할 수 있다

```go
type Report interface {				//Report() 메서드를 포함한 Report 인터페이스
	Report() string
}

type FinanceReport struct {		//경제 보고서를 담당하는 FinanceReport
	report string
}

func (r *FinanceReport) Report() string {	//Report 인터페이스를 구현
	return r.report
}

type ReportSender struct {								//보고서 전송을 담당
	...
}

func (r *ReportSender) SendReport(report Report) {
	//Report 인터페이스 객체를 인수로 받음
	...
}
```

## 27.3 개방-폐쇄 원칙

**정의**

- “확장에는 열려 있고, 변경에는 닫혀 있다.”

**이점**

- 상호 결합도를 줄여 새 기능을 추가할 때 기존 구현을 변경하지 않아도 된다.

⇒ ‘프로그램에 기능을 추가할 때 기존 코드의 변경을 최소화해야 한다’

```go
func SendReport(r *report, method SendType, receiver string) {
	switch method {
		case Email:
			//이메일 전송
		case Fax:
			//팩스 전송
		case PDF:
			//pdf 파일 생성
		case Printer:
			//프린팅
		...
	}
}
```

전송 방식 추가하려면 새로운 case를 만들어 구현을 추가해주면 된다. = 기존 SendReport() 함수 구현 변경 ⇒ 개방-폐쇄 원칙 위배

**개방-폐쇄 원칙 입각한 코드**

```go
type ReportSender interface {
	Send(r *Report)
}

type EmailSender struct {
}

func (e *EmailSender) Send(r *Report) {
	//이메일 전송
}

type FaxSender struct {
}

func (f *FaxSender) Send(r *Report) {
	//팩스 전송
}
```

EmailSender, FaxSender 모두 ReportSender 인터페이스를 구현한 객체

새로운 전송 방식 추가 → ReportSender 구현한 새로운 객체 추가 ( = 기존 구현 변경 X)

잘못된 예제는 단일 책임 원칙도 위반, SOLID 원칙은 상호 연결되어 있어서 하나만 잘 지켜도 나머지가 저절로 지켜지는 경우가 많다

## 27.4 리스코프 치환 원칙

**정의**

- “*q(x)*를 타입 T의 객체 *x*에 대해 증명할 수 있는 속성이라 하자. 그렇다면 S가 T의 하위 타입이라면 *q(y)*는 타입 S의 객체 *y*에 대해 증명할 수 있어야 한다.”

**이점**

- 예상치 못한 작동을 예방할 수 있다.

```go
type T interface {				//Something() 메서드를 포함한 인터페이스
	Something()
}

type S struct {
}

func (s *S) Something() {	//T 인터페이스 구현
}

type U struct {
}

func (u *U) Something() {	//T 인터페이스 구현
}

func q(t T){
	...
}

var y = &S{}
var u = &U{}
q(y)
q(u)                      //둘 다 잘 동작해야 한다
```

T 인터페이스, S 객체와 U 객체가 구현

함수 q()는 인터페이스 T를 인수로 받는다 ⇒ S 객체 인스턴스인 y와 U 객체 인스턴스인 u 모두에 대해서 잘 동작해야 한다 

실제론 그렇지 않는 경우가 발생 : 함수를 호출하는 호출자와 함수 구현자의 계약 관계 발생

```go
type Report interface {
	Report() string
}

func SendReport(r Report)
```

SendReport() 함수의 호출자가 함수 호출하는 의도? Report 전송할 것 ⇒ 호출자와 함수 간 계약이 성립한다

함수 호출했는데 Report가 전송되지 않고 다른 일이 발생된다면 호출자가 예상하지 못한 버그가 발생!

```go
type Report interface {
	Report() string
}

func MarketingReport{
}

func (m *MarketingReport) Report() string{
	...
}

func SendReport(r Report){
	if _, ok := r.(*MarketingReport); ok {
		panic("Can't send MarketingReport")
	}
	...
}

var report = &MarketingReport{}
SendReport(report)
```

Report 인터페이스, MarketingReport 객체

SendReport() 함수는 Report 인터페이스를 인수로 받는다

MarketingReport는 Report 인터페이스 구현 → SendReport()의 인수로 사용 가능

호출자 입장에서는 당연히 MarketingReport 인스턴스도 전송이 잘 될거라 예상 → 실제로는 패닉 발생

상위 타입에 대해서 작동하는 함수가 하위 타입에 대해서도 똑같이 작동해야 하지만 이코드는 그렇지 못하기 때문에 리스코프 치환 원칙을 위배한 코드가 된다.

리스코프 치환 원칙에 입각한 코드 = 함수 계약 관계를 준수하는 코드

Go 언어보다는 상속을 지원하는 다른 언어에서 더 큰 문제를 발생시킨다

## 27.5 인터페이스 분리 원칙

**정의**

- “클라이언트는 자신이 이용하지 않는 메서드에 의존하지 않아야 한다.”

**이점**

- 인터페이스를 분리하면 불필요한 메서드들과 의존 관계가 끊어져 더 가볍게 인터페이스를 이용할 수 있다.

```go
type Report interface {
	Report() string
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func SendReport(r Report) {
	send(r.Report())
}
```

Report 인터페이스는 메서드를 총 4개 포함

SendReport()는 Report 인터페이스가 포함한 4개 메서드 중 Report() 메서드만 사용

즉, 인터페이스 이용자에게 불필요한 메서드들을 인터페이스가 포함하고 있다 ⇒ 인터페이스 분리 원칙 위반한 코드

**인터페이스 분리 원칙에 입각한 코드**

```go
type Report interface {
	Report() string
}

type WrittenInfo interface {
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func SendReport(r Report) {
	send(r.Report())
}
```

Report 인터페이스는 메서드 하나만 가진다. SendReport() 함수가 필요한 유일한 메서드인 Report를 포함한 인터페이스와 관계를 맺고, 불필요한 메서드와는 관계 맺지 않는다. 즉, 많은 메서드들을 포함하는 커다란 인터페이스보다는 적은 수의 메서드를 가진 인터페이스 여러 개로 이뤄진 객체가 더 좋다!

이처럼 인터페이스를 분리해 불필요한 메서드들과 의존 관계를 끊으면 더 가볍게 인터페이스를 이용할 수 있다

## 27.6 의존 관계 역전 원칙

**정의**

- “상위 계층이 하위 계층에 의존하는 전통적인 의존 관계를 반전(역전)시킴으로써 상위 계층이 하위 계층의 구현으로부터 독립되게 할 수 있다.”

원칙

- 원칙 1 : “상위 모듈은 하위 모듈에 의존해서는 안 된다. 둘 다 추상 모듈에 의존해야 한다.”
- 원칙 2 : “추상 모듈은 구체화된 모듈에 의존해서는 안 된다. 구체화된 모듈은 추상 모듈에 의존해야 한다.”

**이점**

- 구체화된 모듈이 아닌 추상 모듈에 의존함으로써 확장성이 증가한다.
- 상호 결합도가 낮아져서 다른 프로그램으로 이식성이 증가한다.

### 27.6.1 원칙 1 뜯어보기

“상위 모듈은 하위 모듈에 의존해서는 안 된다. 둘 다 추상 모듈에 의존해야 한다.”

대개 해결책을 찾을 때 위에서 아래로 내려가며 사고하는 경향이 있다. = 탑다운 방식

예) 키보드로 받은 입력을 네트워크로 ‘전송’하는 객체

![1](.\img\1.png)

상위 모듈인 전송의 동작 : 하위 모듈 키보드에서 값을 읽어서 하위 모듈 네트워크로 값을 쓰는 방식으로 구현 → 상위모듈인 전송과 하위모듈인 키보드, 네트워크 간 결합도가 높아진다

제대로 객체지향 설계 하려면 탑다운 X → 전통적인 방식에서 역전! → ‘둘 다 추상화 모듈에 의존해야 한다’

**의존 관계 역전 원칙에 입각한 설계**

![2](.\img\2.png)

키보드는 입력이라는 추상 모듈 구현, 네트워크는 출력이라는 추상 모듈 구현

전송 모듈은 구체화된 객체인 키보드와 네트워크가 아닌 추상화된 입력과 출력 모듈을 사용

즉, 키보드, 네트워크, 전송 모두 추상 모듈에 의존하는 관계가 된다

어떤 이득을 얻게 되나? 

- 각 의존 관계를 떨어뜨리면 각 모듈은 본연의 기능에 충실할 수 있다.
- 서로 결합도가 낮아짐으로써 독립적이게 된다
- 서로 독립성이 유지되기 때문에 다른 애플리케이션에도 사용 가능 = 코드 재사용성이 높아진다

### 27.6.2 원칙 2 뜯어보기

“추상 모듈은 구체화된 모듈에 의존해서는 안 된다. 구체화된 모듈은 추상 모듈에 의존해야 한다.”

예) 메일이 수신되면 알람을 울린다고 가정

```go
type Mail struct{
	alarm Alarm
}

type Alarm struct{
}

func (m *Mail) OnRecv() {
	m.alarm.Alarm()
}
```

메일 객체는 알람 객체를 소유하고 있고, 메일 수신 시 호출되는 OnRecv() 메서드에서 소유한 알람 객체를 사용해 알람 울린다

![3](.\img\3.png)

메일이라는 구체화된 모듈이 알람이라는 구체화된 모듈에 의존 ⇒ 의존 관계 역전 원칙에 위배

**의존 관계 역전 원칙에 입각한 설계**
![4](.\img\4.png)

메일은 Event라는 인터페이스 구현, 알람은 EventListener라는 인터페이스 구현

구체화된 객체를 통한 관계 X ⇒ 추상화된 객체를 통한 관계

어떤 구체화된 모듈도, 추상화된 모듈도 구체화된 모듈에 의존적이지 않다

```go
type Event interface{
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct{
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {	//Event 인터페이스 구현
	m.listener = listener 
}

func (m *Mail) OnRecv() {			//등록된 listener의 OnFire() 호출
	m.listener.OnFire()
}

type Alarm struct{
}

func (a *Alarm) OnFire() {		//EventListener 인터페이스 구현
	//알람
	fmt.Println("알람! 알람!")
}

var mail = &Mail{}
var listener EventListener = &Alarm{}

mail.Register(listener)
mail.OnRecv()								//알람이 울리게 된다
```

Event 인터페이스는 Register() 메서드를 가지고, Mail 객체는 이를 구현하여 Register() 메서드가 호출되면 EventListener를 등록

OnRecv() 메서드가 호출 → 등록된 EventListener 객체의 OnFire() 메서드 호출

Alarm 객체는 EventListener 인터페이스 구현하여 OnFire() 메서드가 호출될 때 알람이 울리도록 구현

mail 인스턴스에 Alarm 인스턴스를 등록하면 메일 수신 시 알람이 울리게 된다

의존 관계를 역전하면 다양한 EventListener를 만들어서 등록할 수 있다. 즉 개방-폐쇄 원칙 역시 지켜지게 된다. 다양한 Event도 손쉽게 코드를 추가할 수 있다.

## 27.7 학습 마무리

SOLID 5가지 원칙들이 서로가 연결되어 있다. 즉 한 가지 원칙만 잘 지켜도 나머지 원칙들이 저절로 지켜진다.

공통 목적 : ‘결합도는 낮게, 응집도는 높게’

## 연습문제

1. 단일 책임 원칙
2. 3 (의무 아님), 5 (확장에 개방, 변경에 폐쇄)
3. Player와 Monster가 서로 구체화된 객체로 의존한다. → 서로 추상화된 객체로 의존하도록 변경

```go
type Attacker interface{
	Name() string
}

type DamageTaker interface {
	DealDamage(att Attacker, damage int)
}

type Player struct{
	name string
}

type Monster struct{
	hp int
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Attack(dt DamageTAker) {
	dt.DealDamage(p, 100)
}

func (m *Monster) DealDamage(attacker Attacker, damage int) {
	m.hp -= damage
	if m.hp < 0 {
		fmt.Println(attacker.Name(), "가 나를 죽였다")
	}
}
```