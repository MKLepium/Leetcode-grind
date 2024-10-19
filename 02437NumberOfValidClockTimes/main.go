package main

func checkHours(time string, wildcard0 bool, wildcard1 bool) int {
	if !wildcard0 && !wildcard1 {
		return 1
	}

	possibleChoicesHours := 0
	if wildcard0 && wildcard1 {
		possibleChoicesHours = 24
	} else if wildcard0 && !wildcard1 {
		if time[1] > '3' {
			possibleChoicesHours = 2
		} else {
			possibleChoicesHours = 3
		}
	} else {
		if time[0] < '2' {
			possibleChoicesHours = 10
		} else {
			possibleChoicesHours = 4
		}
	}
	return possibleChoicesHours
}

func checkMinutes(_ string, wildcard3 bool, wildcard4 bool) int {
	if !wildcard3 && !wildcard4 {
		return 1
	}
	possibleChoicesMinutes := 0
	if wildcard3 && wildcard4 {
		possibleChoicesMinutes = 60
	} else if !wildcard3 && wildcard4 {
		possibleChoicesMinutes = 10
	} else {
		possibleChoicesMinutes = 6
	}
	return possibleChoicesMinutes
}


func countTime(time string) int {
    /*
	Valid Times
	0. 0, 1, (or 2 if 2. is less than 5)
	1. 0-4, (5-9 if 1. is less than 2)
	2. :
	3. 6
	4. 10
	*/
	var wildcard0 bool = false;
	var wildcard1 bool = false;
	var wildcard3 bool = false;
	var wildcard4 bool = false;
	var validClockTimes int;

	var possibleChoicesHours int;
	var possibleChoicesMinutes int;

	if time[0] == '?' {
		wildcard0 = true;
	}
	if time[1] == '?' {
		wildcard1 = true;
	}
	if time[3] == '?' {
		wildcard3 = true;
	}
	if time[4] == '?' {
		wildcard4 = true;
	}

	if !wildcard0 && !wildcard1 && !wildcard3 && !wildcard4 {
		return 1
	}



	// Choices for Hours
	possibleChoicesHours = checkHours(time, wildcard0, wildcard1)

	// choices for minutes
	possibleChoicesMinutes = checkMinutes(time, wildcard3, wildcard4)


	if possibleChoicesHours == 0 || possibleChoicesMinutes == 0 {
		return 0
	}

	if possibleChoicesHours == 1 {
		return possibleChoicesMinutes
	}

	if possibleChoicesMinutes == 1 {
		return possibleChoicesHours
	}

	validClockTimes = possibleChoicesHours * possibleChoicesMinutes

	return validClockTimes
}
