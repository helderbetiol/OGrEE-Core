// Code generated by goyacc - DO NOT EDIT.

package main

import __yyfmt__ "fmt"

import (
	"cli/cmd"
	"cli/utils"
	"strings"
)

func resMap(x *string) map[string]string {
	resarr := strings.Split(*x, "=")
	res := make(map[string]string)

	for i := 0; i+1 < len(resarr); {
		if i+1 < len(resarr) {
			res[resarr[i]] = resarr[i+1]
			i += 2
		}
	}
	return res
}

type yySymType struct {
	yys int
	//n int
	s string
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault      = 57356
	yyEofCode      = 57344
	TOKEN_ATTR     = 57348
	TOKEN_BASHTYPE = 57350
	TOKEN_CMDFLAG  = 57352
	TOKEN_CRUDOP   = 57349
	TOKEN_DOC      = 57355
	TOKEN_ENTITY   = 57347
	TOKEN_EQUAL    = 57351
	TOKEN_EXIT     = 57354
	TOKEN_SLASH    = 57353
	TOKEN_WORD     = 57346
	yyErrCode      = 57345

	yyMaxDepth = 200
	yyTabOfs   = -20
)

var (
	yyPrec = map[int]int{}

	yyXLAT = map[int]int{
		57344: 0,  // $end (17x)
		57348: 1,  // TOKEN_ATTR (5x)
		57346: 2,  // TOKEN_WORD (5x)
		57360: 3,  // F (3x)
		57364: 4,  // P (2x)
		57357: 5,  // B (1x)
		57358: 6,  // D (1x)
		57359: 7,  // E (1x)
		57361: 8,  // K (1x)
		57362: 9,  // L (1x)
		57363: 10, // M (1x)
		57365: 11, // Q (1x)
		57367: 12, // start (1x)
		57350: 13, // TOKEN_BASHTYPE (1x)
		57352: 14, // TOKEN_CMDFLAG (1x)
		57349: 15, // TOKEN_CRUDOP (1x)
		57355: 16, // TOKEN_DOC (1x)
		57347: 17, // TOKEN_ENTITY (1x)
		57351: 18, // TOKEN_EQUAL (1x)
		57354: 19, // TOKEN_EXIT (1x)
		57353: 20, // TOKEN_SLASH (1x)
		57366: 21, // Z (1x)
		57356: 22, // $default (0x)
		57345: 23, // error (0x)
	}

	yySymNames = []string{
		"$end",
		"TOKEN_ATTR",
		"TOKEN_WORD",
		"F",
		"P",
		"B",
		"D",
		"E",
		"K",
		"L",
		"M",
		"Q",
		"start",
		"TOKEN_BASHTYPE",
		"TOKEN_CMDFLAG",
		"TOKEN_CRUDOP",
		"TOKEN_DOC",
		"TOKEN_ENTITY",
		"TOKEN_EQUAL",
		"TOKEN_EXIT",
		"TOKEN_SLASH",
		"Z",
		"$default",
		"error",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {12, 1},
		2:  {12, 1},
		3:  {12, 1},
		4:  {12, 1},
		5:  {8, 2},
		6:  {8, 4},
		7:  {7, 2},
		8:  {3, 4},
		9:  {3, 4},
		10: {10, 0},
		11: {21, 1},
		12: {4, 3},
		13: {4, 1},
		14: {11, 1},
		15: {5, 3},
		16: {5, 2},
		17: {5, 1},
		18: {6, 1},
		19: {9, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [27][]uint8{
		// 0
		{5: 27, 24, 8: 22, 25, 11: 23, 21, 28, 15: 26, 30, 19: 29},
		{20},
		{19},
		{18},
		{17},
		// 5
		{16},
		{7: 33, 17: 35, 21: 34},
		{6},
		{3, 2: 31},
		{2},
		// 10
		{1},
		{4, 14: 32},
		{5},
		{15},
		{2: 43, 4: 42},
		// 15
		{1: 37, 9, 36},
		{13},
		{18: 38},
		{2: 39},
		{10, 37, 3: 40, 10: 41},
		// 20
		{12},
		{11},
		{1: 37, 3: 46},
		{1: 7, 20: 44},
		{2: 43, 4: 45},
		// 25
		{1: 8},
		{14},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 23

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			println("@State start")
		}
	case 5:
		{
			println("@State K")
		}
	case 6:
		{
			yyVAL.s = yyS[yypt-0].s
			println("Finally: " + yyVAL.s)
			cmd.Disp(resMap(&yyS[yypt-0].s))
		}
	case 8:
		{
			yyVAL.s = string(yyS[yypt-3].s + "=" + yyS[yypt-1].s + "=" + yyS[yypt-0].s)
			println("So we got: ", yyVAL.s)
		}
	case 9:
		{
			yyVAL.s = yyS[yypt-3].s + "=" + yyS[yypt-1].s
			println("Taking the M")
			println("SUP DUDE: ", yyS[yypt-1].s)
		}
	case 17:
		{
			cmd.Execute()
		}
	case 18:
		{
			utils.Exit()
		}
	case 19:
		{
			cmd.Help()
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
