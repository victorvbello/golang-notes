package human

type AndroidAdapter struct {
	android *Android
}

func NewAndroidAdapter(android *Android) *AndroidAdapter {
	return &AndroidAdapter{android: android}
}

func (aa *AndroidAdapter) GetName() string {
	return aa.android.ID
}

func (aa *AndroidAdapter) Walk(steps int) {
	for i := 0; i < steps; i++ {
		aa.android.MoveForward()
	}
}

func (aa *AndroidAdapter) Talk(words int) {
	var txt string
	for i := 0; i < words; i++ {
		txt += aa.android.ThinkWord()
	}
	aa.android.ReproduceSound(aa.android.TextToSound(txt))
}
