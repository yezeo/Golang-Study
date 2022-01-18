# Chapter 27. 객체 지향 설계 원칙 SOLID

## 27-I. 객체 지향 설계 5가지 원칙 SOLID

- `SOLID`: 객체 지향 설계 5가지 원칙
    - SRP 단일 책임 원칙
    - OCP 개방-폐쇄 원칙
    - LSP 리스코프 치환 원칙
    - ISP 인터페이스 분리 원칙
    - DIP 의존 관계 역전 원칙


- 나쁜 설계: 상호 결집도 매우 높음, 응집도 낮음
    - 경직성 rigidity: 모듈 간 결합도 높음
    - 부서지기 쉬움 fragility
    - 부동성 immobility
    - **하나의 모듈이 스스로 완성 X, 다른 모듈에 의존적인 관계를 가지고 있는 경우** :(
    

- 좋은 설계: 상호 결합도 낮음, 응집도 높음

---

## 27-II. 단일 책임 원칙, Single Responsibility Principle

- 코드 재사용성 높여 줌

```go
type FianceReport struct { // 회계 보고서
    report string
}

func (r *FianceReport) SendReport(email string) { // 보고서 전송

}
```

`FianceReport`: 회계 보고서 객체, 하지만 보고서 전송 책임까지 가지고 있으므로 단일 책임 원칙 **위배**

만약 다른 종류의 Report가 있다면 SendReport() 메서드를 새로 만들어야 함

▶ FianceReport는 Report 인터페이스 구현, ReportSender는 Report 인터페이스 이용하는 관계 형성하면 단일 책임 원칙에 입각한 설계가 됨.

```go
type Report interface {
    // Report() 메서드를 포함한 Report 인터페이스
    Report() string
}

type FianceReport struct {
    // 경제 보고서를 담당하는 FianceReport
    report string
}

func (r *FianceReport) Report() string {
    // Report 인터페이스를 구현
    return r.report 
}

type ReportSender struct {
    // 보고서 전송 담당
}

func (s *ReportSender) SendReport(report Report) {
    // Report 인터페이스 객체를 인수로 받음
}
```

---

## 27-III. 개방-폐쇄 원칙, Open-closed principle

확장에는 열려 있고, 변경에는 닫혀 있다

- 상호 결합도를 줄여 새 기능을 추가할 때 기존 구현을 변경하지 않아도 되도록 함

```go
func SenReport(r *Report, method SendType, receiver string) {
    switch method {
    case Email: //이메일 전송
    case Fax: // 팩스 전송
    case PDF: // pdf 파일 생성
    case Printer: // 프린팅
    }
}
```

전송 방식 추가 시 새로운 case 만들어 구현 추가해 줌 -> 기존 함수 구현 변경, 따라서 원칙 위배. 단일 책임 원칙도 위배됨.

```go
type ReportSender interface { Send (r *Report) }

type EmailSender struct {}
func (e *EmailSender) Send(r *Report) { // 이메일 전송 }

type FaxSender struct {}
func (f *FaxSender) Send(r *Report { // 팩스 전송 }
```

새로운 전송 방식을 추가할 때 새로운 객체를 추가해 주면 됨. 기존 구현은 변경할 필요 없음.

---


## 27-IV. 리스코프 치환 원칙 Liskov Substitution principle, LSP

- 예상치 못한 작동 예방 가능

상위 타입을 인수로 받는 함수에 하위 타입의 인스턴스를 넣어도 잘 동작해야 함

```go
type Report interface {
    Report() string
}

type MarketingReport {}

func (m *MarketingReport) Report() string {...}

func SendReport(r Report) {
    if _, ok := r.(*MarketingReporT); ok { // r이 마케팅 보고서일 경우 패닉
        painc("Can't send MarketingReport")
    }
}

var report = &MarketingReport{}
SendReport(report) // 패닉 발생
```

호출자 입장에서 MarketingReport 인스턴스도 전송이 잘 될 것이라 예상하지만, 실제로는 패닉 발생

상위 타입 Report에 대해 작동하는 SendReport() 함수는 하위 타입인 MarketingReport에 대해서도 똑같이 작동해야 하지만, 위 코드는 그렇지 않음 -> 원칙 위배

---

## 27-V. 인터페이스 분리 원칙 Interface Segregation Principle

- 클라이언트는 자신이 이용하지 않는 메서드에 의존하지 않아야 함
    - 인터페이스 분리 -> 불필요한 메서드들과 의존 관계가 끊어짐 -> 더 가볍게 인터페이스 이용 가능

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

Report 메소드는 메서드를 총 4개 포함하지만 SendReport()는 그 중 하나만 이용. 따라서 이용자에게 불필요한 메서드들을 인터페이스가 포함 중 -> 인터페이스 분리 원칙을 위반

```go
type Report interface { Report() string }
type WrittenInfo interface { 
    Pages() int 
    Author() string
    WrittenDate() time.Time
}

func SendReport(r Report) {
    send(r.Report())
}
```

많은 메소드들을 포함하는 커다란 인터페이스보다는 적은 수의 메서드를 가진 인터페이스 여러 개로 이루어진 객체가 더 좋음

인터페이스 분리 -> 불필요한 메서드들과 의존 관계 X -> 더 가볍게 인터페이스 이용 가능

---


## 27-VI. 의존 관계 역전 원칙, Dependency Inversion Principle

- 상위 계층이 하위 계층에 의존하는 전통적인 의존 관계를 반전(역전)시킴으로써 상위 계층이 하위 계층의 구현으로부터 독립되게 할 수 있다.
    - 원칙1: 상위 모듈은 하위 모듈에 의존 X, 둘 다 추상 모듈에 의존해야 함
    - 원칙2: 추상 모듈은 구체화된 모듈에 의존 X, 구체화된 모듈은 추상 모듈에 의존해야 함

---

## 연습 문제

1. S

2. 3번,  5번

3.

```go
type Attacker interface { Name() string }
type DamageTaker interface { DealDamage(att Attacker, damage int )}
type Player struct { name string }
type Monster struct { hp int }
func (p *Player) Nam() string { return p.name }
func (p *Player) Attack(dt DamageTaker) { dt.DealDamage(p, 100) }
func (m *Monster) DealDamage (attacker Attacker, damage int) {
    m.hp -= damage
    if m.hp < 0 { fmt.Println(attackr.Name(), "가 나를 죽였다") }
}
```
