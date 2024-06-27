package ReplacingTime_1736

func maximumTime(time string) string {
	timeByte := []byte(time)
	if timeByte[3] == '?' {
		timeByte[3] = '5'
	}
	if timeByte[4] == '?' {
		timeByte[4] = '9'
	}
	if timeByte[0] == '?' {
		if int(timeByte[1]-'0') > 3 && int(timeByte[1]-'0') < 10 {
			timeByte[0] = '1'
		} else {
			timeByte[0] = '2'
		}
	}
	if timeByte[1] == '?' {
		timeByte[1] = '9'
	}
	if timeByte[0] == '2' && timeByte[1] == '9' {
		timeByte[1] = '3'
	}
	return string(timeByte)
}
