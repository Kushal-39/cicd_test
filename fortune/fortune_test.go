package fortune

import "testing"

func TestRandomFortune(t *testing.T){
	f:= []string{"a","b","c"}
	for i:=0;i<10;i++{
		r:= RandomFortune(f)
		found:=false
		for _,val := range f{
			if r==val{
				found = true
				break
			}
		}
		if !found{
			t.Errorf("RandomFortune returned invalid value:%s",r)
		}
	}
}