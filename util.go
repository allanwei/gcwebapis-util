package util

import (
	"fmt"
	"math"
	"net/url"
	"reflect"
	"time"
)

//Decoder ...
func Decoder(Str string) (string, error) {
	u, err := url.QueryUnescape(Str)
	if err != nil {
		return "", err
	}
	return u, nil

}

//TimeSpan ....
type TimeSpan struct {
	H int
	M int
	S int
}

//ParseTimeSpanWithMinutes ....
func ParseTimeSpanWithMinutes(i int) TimeSpan {
	return TimeSpan{H: 0, M: i, S: 0}
}

//GetHoursFromMin ....
func (d TimeSpan) GetHoursFromMin() (float64, error) {
	str := fmt.Sprintf("%dm", d.M)
	h, err := time.ParseDuration(str)
	if err != nil {
		return 0, err
	}
	return h.Hours(), nil
}

//GetTimes ....
func GetTimes() {
	h, err := time.ParseDuration("360m")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("Start: ?", h.Hours())
	h, err = time.ParseDuration("840m")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("End: ?", h.Hours())
	h, err = time.ParseDuration("840m")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("Start: ?", h.Hours())
	h, err = time.ParseDuration("1320m")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("End: ?", h.Hours())
	h, err = time.ParseDuration("1320m")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("Start: ?", h.Hours())
	h, err = time.ParseDuration("360m")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("End: ?", h.Hours())

}

//Strings ...
func Strings(input []string) []string {
	u := make([]string, 0)
	m := make(map[string]bool)
	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}

//Subtract ...
func Subtract(endtime, starttime string) (Duration *time.Duration, err error) {
	s, err := TimeParseRFC3339(starttime)
	if err != nil {
		return nil, err
	}
	e, err := TimeParseRFC3339(endtime)
	if err != nil {
		return nil, err
	}
	d := e.Sub(s)
	return &d, nil

}

//TimeParseRFC3339 ...
func TimeParseRFC3339(timestr string) (time.Time, error) {
	za := GetTimeZoneAttach()
	var ts string
	if len(timestr) < 20 {
		ts = fmt.Sprintf("%s%s", timestr, za)
	} else {
		ts = timestr
	}
	v1, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		return time.Now(), err
	}
	return v1, nil

}

//StrutForScan ...
func StrutForScan(u interface{}, columns []string) []interface{} {
	val := reflect.ValueOf(u).Elem()
	v := make([]interface{}, val.NumField())
	typeOfT := val.Type()
	for i := 0; i < len(columns); i++ {
		for j := 0; j < val.NumField(); j++ {
			valueField := val.Field(j)
			name := typeOfT.Field(j).Name
			if columns[i] == name {
				v[i] = valueField.Addr().Interface()
				break
			}
		}
	}

	return v
}

//TimeDurationToHHMMString ...
func TimeDurationToHHMMString(duration float64, unit string) (rect string) {
	s := fmt.Sprintf("%fm", duration)
	h, err := time.ParseDuration(s)
	if err != nil {
		return
	}
	sdate := "2016-01-02T00:00:00Z"
	t, _ := TimeParseRFC3339(sdate)
	a := t.Add(h)
	so := a.Format(time.RFC3339)
	rect = so[11:16]
	return
}

//TimeGetDate ...
func TimeGetDate(t time.Time) time.Time {
	s := t.Format("2006-01-02T00:00:00")
	n, _ := TimeParseRFC3339(s)
	return n
}

//SubstractTimeFromString ...
func SubstractTimeFromString(end, start string) time.Duration {
	endt, _ := TimeParseRFC3339(end)
	stt, _ := TimeParseRFC3339(start)
	d := endt.Sub(stt)
	return d
}

//GetSQLTimeString ...
func GetSQLTimeString(t time.Time) string {
	out := t.Format(time.RFC3339)
	o := out[0:19]
	return o
}

//TimeStringGetDate ...
func TimeStringGetDate(timestr string) *string {
	v1, _ := TimeParseRFC3339(timestr)
	s := v1.Format("2006-01-02T00:00:00")
	s1, _ := TimeParseRFC3339(s)
	out := GetSQLTimeString(s1)

	return &out
}

//GetTimeZoneAttach ...
func GetTimeZoneAttach() string {
	t := time.Now()
	_, offset := t.Zone()
	offcha := "-"
	if offset > 0 {
		offcha = "+"
	}
	oh := float64(offset / 3600)
	hs := math.Abs(oh)
	za := fmt.Sprintf("%s0%d:00", offcha, int(hs))
	return za
}
