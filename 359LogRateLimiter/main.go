package main

type Logger struct {
	seenVars map[string]int;
}

func Constructor() Logger {
	var Logger Logger;
	Logger.seenVars = make(map[string]int);
	return Logger;
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if message == "" {
		return false;
	}

	val, ok := this.seenVars[message];
	if !ok {
		this.seenVars[message] = timestamp;
		return true;
	}

	if (val + 10) <= timestamp {
		this.seenVars[message] = timestamp;
		return true;
	}
	return false;
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
