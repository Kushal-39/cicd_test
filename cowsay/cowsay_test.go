package cowsay

import (
	"strings"
	"testing"
)

func TestWrapping(t *testing.T){
	msg:="This is a text message to wrap properly"
	wrapped:= WrapText(msg,10)
	lines:=0
	for _, line:=range strings.Split(wrapped,"\n"){
		if len(line)>10{
			t.Errorf("line too long: %s",line)
		}
		lines++
	}
	if lines==0{
		t.Errorf("No lines returned")
	}
}