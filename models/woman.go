package models

type Woman struct {
	Man
}

func (w *Woman) Breathe()    { w.breathe = true }
func (w *Woman) Eat()        { w.eat = true }
func (w *Woman) Think()      { w.think = true }
func (w *Woman) Sex() string { return "Female" }
func (w *Woman) IsAlive() bool {
	if w.breathe && w.eat && w.think {
		w.live = true
	}
	return w.live
}
