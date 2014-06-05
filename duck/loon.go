package duck

type Loon struct {
	IsSwimming bool
}

func (l *Loon) Quack() string {
	return "Quuuaack!"
}

func (l *Loon) Swim() {
	l.IsSwimming = true
}

func (l Loon) SwimWithoutPointer() {
	l.IsSwimming = true
}

func (l *Loon) Walk() {
	l.IsSwimming = false
}

type BabyLoon struct {
	Loon
}

func (l *BabyLoon) Quack() string {
	return "Quuueeck!"
}
