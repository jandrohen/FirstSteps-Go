package models

type Man struct {
	age     int
	height  float32
	weight  float32
	breathe bool
	think   bool
	eat     bool
	live    bool
}

func (m *Man) Breathe()    { m.breathe = true }
func (m *Man) Eat()        { m.eat = true }
func (m *Man) Think()      { m.think = true }
func (m *Man) Sex() string { return "Male" }
func (m *Man) IsAlive() bool {
	if m.breathe && m.eat && m.think {
		m.live = true
	}
	return m.live
}
