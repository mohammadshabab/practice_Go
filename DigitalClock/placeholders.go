package main

type placeholder [5]string

var zero = placeholder{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = placeholder{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}

var two = placeholder{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = placeholder{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = placeholder{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = placeholder{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = placeholder{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = placeholder{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = placeholder{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = placeholder{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}
var colon = placeholder{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

/*
var alarm = [...]placeholder{
	{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	},
	{
		"███",
		"█ █",
		"███",
		"█ █",
		"█ █",
	},
	{
		"█  ",
		"█  ",
		"█  ",
		"█  ",
		"███",
	},
	{
		"███",
		"█ █",
		"███",
		"█ █",
		"█ █",
	},
	{
		"██ ",
		"█ █",
		"██ ",
		"█ █",
		"█ █",
	},
	{
		"█ █",
		"███",
		"█ █",
		"█ █",
		"█ █",
	},
	{
		" █ ",
		" █ ",
		" █ ",
		"   ",
		" █ ",
	},
	{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	},
}
*/
var dot = placeholder{
	"   ",
	"   ",
	"   ",
	"   ",
	" ░ ",
}

var digits = [...]placeholder{
	zero, one, two, three, four, five, six, seven, eight, nine,
}

// var al = [...]alarm{
// 	a, l, a, r, m, exc,
// }

// for line := range digits[0] {

// 	for digit := range digits {
// 		fmt.Print(digits[digit][line], " ")
// 	}
// 	fmt.Println()
// }
