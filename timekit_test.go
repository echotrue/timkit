package timkit

import (
	"testing"
	"time"
)

var (
	tk  *TimeKit
	l   *time.Location
	err error
)

func Location(t *testing.B) {
	l, err = time.LoadLocation("Local")
	if err != nil {
		t.Errorf("Test failed：%s .\n", err.Error())
	}
}
func init() {
	t := new(testing.B)
	Location(t)
	tk = NewTimeKit(time.Date(2021, time.Month(1), 2, 15, 4, 5, 0, l))
}

func BenchmarkTimeKit_AddCenturies(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	t := NewTimeKit(time.Now())

	for i := 0; i < b.N; i++ {
		t.AddCenturies(2)
	}
}

func BenchmarkTimeKit_StartOfWeek(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	t := NewTimeKit(time.Now())

	for i := 0; i < b.N; i++ {
		t.StartOfWeek()
	}
}

func TestTimeKit_AddCenturies(t *testing.T) {
	expected := NewTimeKit(time.Date(2221, 1, 2, 15, 4, 5, 0, l))
	tk.AddCenturies(2)

	if tk.String() != expected.String() {
		t.Errorf("AddCenturies = %+v ,expected %+v", tk, expected)
	}

}

func TestTimeKit_AddYears(t *testing.T) {
	expected := NewTimeKit(time.Date(2023, 1, 2, 15, 4, 5, 0, l))
	tk.AddYears(2)

	if tk.String() != expected.String() {
		t.Errorf("AddYears = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_AddQuarter(t *testing.T) {
	expected := NewTimeKit(time.Date(2021, 4, 2, 15, 4, 5, 0, l))
	tk.AddQuarters(1)

	if tk.String() != expected.String() {
		t.Errorf("AddQuarter = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_AddMonths(t *testing.T) {
	expected := NewTimeKit(time.Date(2021, 8, 2, 15, 4, 5, 0, l))
	tk.AddMonths(7)

	if tk.String() != expected.String() {
		t.Errorf("AddMonths = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_AddWeek(t *testing.T) {
	expected := NewTimeKit(time.Date(2021, 1, 16, 15, 4, 5, 0, l))
	tk.AddWeeks(2)

	if tk.String() != expected.String() {
		t.Errorf("AddWeek = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_AddDays(t *testing.T) {
	expected := NewTimeKit(time.Date(2021, 1, 11, 15, 4, 5, 0, l))
	tk.AddDays(3)

	if tk.String() != expected.String() {
		t.Errorf("AddDays = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_AddHours(t *testing.T) {
	expected := NewTimeKit(time.Date(2021, 1, 9, 6, 4, 5, 0, l))
	tk.AddHours(16)

	if tk.String() != expected.String() {
		t.Errorf("AddHours = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_AddMinutes(t *testing.T) {
	expected := NewTimeKit(time.Date(2021, 1, 8, 15, 20, 5, 0, l))
	tk.AddMinutes(16)

	if tk.String() != expected.String() {
		t.Errorf("AddMinutes = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_AddSeconds(t *testing.T) {
	expected := NewTimeKit(time.Date(2021, 1, 8, 15, 4, 28, 0, l))

	tk.AddSeconds(23)

	if tk.String() != expected.String() {
		t.Errorf("AddSeconds = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_AddWeekdays(t *testing.T) {
	expected := NewTimeKit(time.Date(2021, 1, 8, 15, 4, 5, 0, l))

	tk.AddWeekdays(5)

	if tk.String() != expected.String() {
		t.Errorf("AddWeekdays = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_StartOfCentury(t *testing.T) {
	expected := NewTimeKit(time.Date(2000, time.Month(1), 1, 0, 0, 0, 0, l))
	tk.StartOfCentury()
	if tk.String() != expected.String() {
		t.Errorf("StartOfCentury = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_StartOfYear(t *testing.T) {
	expected := NewTimeKit(time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, l))
	tk.StartOfYear()
	if tk.String() != expected.String() {
		t.Errorf("StartOfYear = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_StartOfQuarter(t *testing.T) {
	expected := NewTimeKit(time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, l))
	tk.StartOfQuarter()
	if tk.String() != expected.String() {
		t.Errorf("StartOfQuarter = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_StartOfMonth(t *testing.T) {
	expected := NewTimeKit(time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, l))
	tk.StartOfMonth()
	if tk.String() != expected.String() {
		t.Errorf("StartOfMonth = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_StartOfWeek(t *testing.T) {
	expected := NewTimeKit(time.Date(2020, time.Month(12), 28, 0, 0, 0, 0, l))
	tk.StartOfWeek()
	if tk.String() != expected.String() {
		t.Errorf("StartOfWeek = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_StartOfDay(t *testing.T) {
	expected := NewTimeKit(time.Date(2021, time.Month(1), 2, 0, 0, 0, 0, l))
	tk.StartOfDay()
	if tk.String() != expected.String() {
		t.Errorf("StartOfDay = %+v ,expected %+v", tk, expected)
	}
}

func TestTimeKit_DiffInMonths(t *testing.T) {
	ntk := NewTimeKit(time.Date(2020, 9, 18, 10, 30, 40, 0, l))
	expected := tk.DiffInMonths(ntk, true)
	if expected != 3 {
		t.Errorf("DiffInMonths = %+v ,expected %+v", expected, 3)
	}
}
