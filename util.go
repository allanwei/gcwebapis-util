package util

import (
	"fmt"
	"net/url"
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
	s, err := time.Parse(starttime, starttime)
	if err != nil {
		return nil, err
	}
	e, err := time.Parse(endtime, endtime)
	if err != nil {
		return nil, err
	}
	d := e.Sub(s)
	return &d, nil

}
