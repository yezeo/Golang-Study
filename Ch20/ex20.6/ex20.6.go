package main

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker
	att.Attack()
}
