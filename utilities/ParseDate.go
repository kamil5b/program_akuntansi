package utilities

import "time"

//Turn date to string
func DateToString(time.Time) string {
	str := time.Now().String()[:23]
	tahun := str[0:4]
	bulan := str[5:7]
	tanggal := str[8:10]
	waktu := str[10:]
	strHasil := tanggal + "/" + bulan + "/" + tahun + waktu
	return strHasil
}

//Turn string to time.Time
func ParsingDate(input string) (time.Time, error) {
	if len(input) > 19 {
		tmp := input
		input = tmp[:19]
	} else if len(input) == 10 {
		input += " 00:00:00"
	} else if len(input) == 11 {
		input += "00:00:00"
	}
	t, err := time.Parse("02/01/2006 15:04:05", input)
	i := 0
	for err != nil {
		switch i {
		case 0:
			t, err = time.Parse("02-01-2006 15:04:05", input)
		case 1:
			t, err = time.Parse("2006/01/02 15:04:05", input)
		case 2:
			t, err = time.Parse("2006-01-02 15:04:05", input)
		case 3:
			t, err = time.Parse("01/02/2006 15:04:05", input)
		case 4:
			t, err = time.Parse("01-02-2006 15:04:05", input)
		case 5:
			t, err = time.Parse("2006/02/01 15:04:05", input)
		case 6:
			t, err = time.Parse("2006-02-01 15:04:05", input)
		case 7:
			t, err = time.Parse("02/01/2006T15:04:05", input)
		case 8:
			t, err = time.Parse("02-01-2006T15:04:05", input)
		case 9:
			t, err = time.Parse("2006/01/02T15:04:05", input)
		case 10:
			t, err = time.Parse("2006-01-02T15:04:05", input)
		case 11:
			t, err = time.Parse("01/02/2006T15:04:05", input)
		case 12:
			t, err = time.Parse("01-02-2006T15:04:05", input)
		case 13:
			t, err = time.Parse("2006/02/01T15:04:05", input)
		case 14:
			t, err = time.Parse("2006-02-01T15:04:05", input)
		case 15:
			t, err = time.Parse("January 02 2006 15:04:05", input)
		case 16:
			t, err = time.Parse("Jan 02 2006 15:04:05", input)
		case 17:
			t, err = time.Parse("02 January 2006 15:04:05", input)
		case 18:
			t, err = time.Parse("02 Jan 2006 15:04:05", input)
		case 19:
			t, err = time.Parse("2006 January 02 15:04:05", input)
		case 20:
			t, err = time.Parse("2006 Jan 02 15:04:05", input)
		case 21:
			t, err = time.Parse("2006 January 02T15:04:05", input)
		case 22:
			t, err = time.Parse("2006 Jan 02T15:04:05", input)
		default:
			return t, err
		}
		i++
	}
	return t, err
}
