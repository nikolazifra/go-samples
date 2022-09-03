package samples

import (
	"testing"
)

func TestInput(t *testing.T) {

	var u UserInput = &NumericInput{}
	u.Add('1')
	u.Add('a')
	u.Add('.')
	u.Add('0')
	want := "10"
	if u.GetValue() != want {
		t.Errorf("GetValue want %v but got %v", want, u.GetValue())
	}

}
