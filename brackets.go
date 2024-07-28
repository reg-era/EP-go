package ep

func BracketProcess(str string) string {
	test := ""
	for i := 0; i < len(str); i++ {
		if str[i] == '(' || str[i] == ')' || str[i] == '[' || str[i] == ']' || str[i] == '{' || str[i] == '}' {
			test += string(str[i])
		}
		if len(test) >= 2 && IsBrackets(test[len(test)-2], test[len(test)-1]) {
			test = test[:len(test)-2]
		}
	}
	if len(test) == 0 {
		return "Ok"
	} else {
		return "Error"
	}
}

func Bracket(arg []string) string {
	res := ""
	for i := 0; i < len(arg); i++ {
		res += BracketProcess(arg[i]) + "\n"
	}
	return "sdfdsfsdfds"
}

func IsBrackets(f, l byte) bool {
	if (f == '(' && l == ')') || (f == '[' && l == ']') || (f == '{' && l == '}') {
		return true
	} else {
		return false
	}
}
