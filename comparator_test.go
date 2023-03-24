package comparator

import "testing"

func TestGetIndext(t *testing.T) {
	
	// First test with default alphabet
	c := New()

	cases := []struct {
		in string
		want int
	}{
		{"Bob",1},
		{"Jaque",9},
		{"Toor",19},
		{"Siegmunt",18},
		{"Fortcale",5},
	}
	for _, v := range cases {
		got := c.GetIndex(v.in[0:1])
		if got != v.want {
			t.Errorf("c.GetIndex(%v) == %d, want %d", v.in, got, v.want)
		}
	}

	// Second test with redact alphabet
	c.SetABC([]string{"Z","I","A","abc"})
	
	cases = []struct {
		in string
		want int
	}{
		{"alex",3},
		{"Zoo",0},
		{"Ivan",1},
		{"Andrey",2},
		{"Sack",4},
		{"bob",3},
	}

	for _, v := range cases {
		got := c.GetIndex(v.in[0:1])
		if got != v.want {
			t.Errorf("c.GetIndex(%v) == %d, want %d \nWith custom alphbet [\"Z\",\"I\",\"A\",\"abc\"]", v.in, got, v.want)
		}
	}
}

func TestCompare(t *testing.T) {
	
	// First test with default alphabet
	c := New()

	cases := []struct {
		in []string
		want int
	}{
		{[]string{"Bob","Oleg"}, -1},
		{[]string{"Jaque","Mark"},-1},
		{[]string{"Toor","Alex"},1},
		{[]string{"Siegmunt","Sack"},1},
		{[]string{"Fortcale","Fortcale"},0},
	}
	for _, v := range cases {
		got := c.Compare(v.in[0], v.in[1])
		if got != v.want {
			t.Errorf("c.Compare(%v,%v) == %d, want %d", v.in[0], v.in[1], got, v.want)
		}
	}

	// Second test with redact alphabet
	c.SetABC([]string{"Z","I","A","abc"})
	
	cases = []struct {
		in []string
		want int
	}{
		{[]string{"Bob","Oleg"}, -1},
		{[]string{"Jaque","Mark"},0},
		{[]string{"Is","Zoo"},1},
		{[]string{"are","bound"},0},
		{[]string{"Fortcale","Fortcale"},0},
	}

	for _, v := range cases {
		got := c.Compare(v.in[0], v.in[1])
		if got != v.want {
			t.Errorf("c.Compare(%v,%v) == %d, want %d \nWith custom alphbet [\"Z\",\"I\",\"A\",\"abc\"]", v.in[0], v.in[1], got, v.want)
		}
	}
}

func TestLess(t *testing.T) {
	
	c := New()

	cases := []struct {
		in []string
		want bool
	}{
		{[]string{"Bob","Oleg"}, true},
		{[]string{"Jaque","Mark"}, true},
		{[]string{"Toor","Alex"}, false},
		{[]string{"Siegmunt","Sack"}, false},
		{[]string{"Fortcale","Fortcale"}, false},
	}
	for _, v := range cases {
		got := c.Less(v.in[0], v.in[1])
		if got != v.want {
			t.Errorf("c.Less(%v,%v) == %t, want %t", v.in[0], v.in[1], got, v.want)
		}
	}
}
func TestMore(t *testing.T) {
	
	c := New()

	cases := []struct {
		in []string
		want bool
	}{
		{[]string{"Bob","Oleg"}, false},
		{[]string{"Jaque","Mark"}, false},
		{[]string{"Toor","Alex"}, true},
		{[]string{"Siegmunt","Sack"}, true},
		{[]string{"Fortcale","Fortcale"}, false},
	}
	for _, v := range cases {
		got := c.More(v.in[0], v.in[1])
		if got != v.want {
			t.Errorf("c.More(%v,%v) == %t, want %t", v.in[0], v.in[1], got, v.want)
		}
	}

}
func TestEqual(t *testing.T) {
	
	c := New()

	cases := []struct {
		in []string
		want bool
	}{
		{[]string{"Bob","Oleg"}, false},
		{[]string{"Jaque","Mark"}, false},
		{[]string{"Toor","Alex"}, false},
		{[]string{"Siegmunt","Sack"}, false},
		{[]string{"Fortcale","Fortcale"}, true},
	}
	for _, v := range cases {
		got := c.Equal(v.in[0], v.in[1])
		if got != v.want {
			t.Errorf("c.Equal(%v,%v) == %t, want %t", v.in[0], v.in[1], got, v.want)
		}
	}

}
func TestLessEqual(t *testing.T) {
	
	c := New()

	cases := []struct {
		in []string
		want bool
	}{
		{[]string{"Bob","Oleg"}, true},
		{[]string{"Jaque","Mark"}, true},
		{[]string{"Toor","Alex"}, false},
		{[]string{"Siegmunt","Sack"}, false},
		{[]string{"Fortcale","Fortcale"}, true},
	}
	for _, v := range cases {
		got := c.LessEqual(v.in[0], v.in[1])
		if got != v.want {
			t.Errorf("c.LessEqual(%v,%v) == %t, want %t", v.in[0], v.in[1], got, v.want)
		}
	}
}
func TestMoreEqual(t *testing.T) {
	
	c := New()

	cases := []struct {
		in []string
		want bool
	}{
		{[]string{"Bob","Oleg"}, false},
		{[]string{"Jaque","Mark"}, false},
		{[]string{"Toor","Alex"}, true},
		{[]string{"Siegmunt","Sack"}, true},
		{[]string{"Fortcale","Fortcale"}, true},
	}
	for _, v := range cases {
		got := c.MoreEqual(v.in[0], v.in[1])
		if got != v.want {
			t.Errorf("c.MoreEqual(%v,%v) == %t, want %t", v.in[0], v.in[1], got, v.want)
		}
	}

}