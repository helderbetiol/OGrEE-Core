// Code generated by goyacc - DO NOT EDIT.

package main

import __yyfmt__ "fmt"

import (
	cmd "cli/controllers"
	"strconv"
	"strings"
)

func resMap(x *string) map[string]interface{} {
	resarr := strings.Split(*x, "=")
	res := make(map[string]interface{})
	attrs := make(map[string]string)

	for i := 0; i+1 < len(resarr); {
		if i+1 < len(resarr) {
			switch resarr[i] {
			case "id", "name", "category", "parentID",
				"description", "domain", "parentid", "parentId":
				res[resarr[i]] = resarr[i+1]

			default:
				attrs[resarr[i]] = resarr[i+1]
			}
			i += 2
		}
	}
	res["attributes"] = attrs
	return res
}

func replaceOCLICurrPath(x string) string {
	return strings.Replace(x, "_", cmd.State.CurrPath, 1)
}

type yySymType struct {
	yys     int
	n       int
	s       string
	sarr    []string
	ast     *ast
	wNode   *whileNode
	fNode   *forNode
	iNode   *ifNode
	rNode   *symbolReferenceNode
	cNode   *comparatorNode
	aNode   *arithNode
	bNode   *boolNode
	nNode   *numNode
	comNode *commonNode
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault      = 57424
	yyEofCode      = 57344
	TOK_AND        = 57415
	TOK_ATTR       = 57356
	TOK_ATTRSPEC   = 57390
	TOK_BASHTYPE   = 57365
	TOK_BLDG       = 57350
	TOK_BOOL       = 57359
	TOK_CD         = 57371
	TOK_CLR        = 57373
	TOK_CMDFLAG    = 57367
	TOK_CMDS       = 57401
	TOK_COMMA      = 57399
	TOK_CREATE     = 57360
	TOK_DELETE     = 57363
	TOK_DEREF      = 57404
	TOK_DEVICE     = 57353
	TOK_DIV        = 57417
	TOK_DOC        = 57370
	TOK_DONE       = 57423
	TOK_DOT        = 57400
	TOK_ELSE       = 57409
	TOK_EQUAL      = 57366
	TOK_EXIT       = 57369
	TOK_FI         = 57422
	TOK_FOR        = 57407
	TOK_GET        = 57361
	TOK_GREATER    = 57419
	TOK_GREP       = 57374
	TOK_IF         = 57406
	TOK_LBLOCK     = 57410
	TOK_LBRAC      = 57397
	TOK_LESS       = 57420
	TOK_LPAREN     = 57412
	TOK_LS         = 57375
	TOK_LSBLDG     = 57380
	TOK_LSDEV      = 57383
	TOK_LSOG       = 57377
	TOK_LSRACK     = 57382
	TOK_LSROOM     = 57381
	TOK_LSSITE     = 57379
	TOK_LSSUBDEV   = 57384
	TOK_LSSUBDEV1  = 57385
	TOK_LSTEN      = 57378
	TOK_MULT       = 57418
	TOK_NOT        = 57416
	TOK_NUM        = 57346
	TOK_OCBLDG     = 57386
	TOK_OCDEL      = 57358
	TOK_OCDEV      = 57387
	TOK_OCPSPEC    = 57395
	TOK_OCRACK     = 57388
	TOK_OCROOM     = 57389
	TOK_OCSDEV     = 57393
	TOK_OCSDEV1    = 57394
	TOK_OCSITE     = 57391
	TOK_OCTENANT   = 57392
	TOK_OR         = 57414
	TOK_PLUS       = 57357
	TOK_PWD        = 57372
	TOK_RACK       = 57352
	TOK_RBLOCK     = 57411
	TOK_RBRAC      = 57398
	TOK_ROOM       = 57351
	TOK_RPAREN     = 57413
	TOK_SEARCH     = 57364
	TOK_SELECT     = 57396
	TOK_SEMICOL    = 57405
	TOK_SITE       = 57349
	TOK_SLASH      = 57368
	TOK_SUBDEVICE  = 57354
	TOK_SUBDEVICE1 = 57355
	TOK_TEMPLATE   = 57402
	TOK_TENANT     = 57348
	TOK_THEN       = 57421
	TOK_TREE       = 57376
	TOK_UPDATE     = 57362
	TOK_VAR        = 57403
	TOK_WHILE      = 57408
	TOK_WORD       = 57347
	yyErrCode      = 57345

	yyMaxDepth = 200
	yyTabOfs   = -140
)

var (
	yyPrec = map[int]int{
		TOK_MULT:  0,
		TOK_OCDEL: 0,
		TOK_DIV:   0,
		TOK_PLUS:  0,
		TOK_EQUAL: 1,
	}

	yyXLAT = map[int]int{
		57344: 0,   // $end (119x)
		57422: 1,   // TOK_FI (118x)
		57423: 2,   // TOK_DONE (116x)
		57409: 3,   // TOK_ELSE (113x)
		57358: 4,   // TOK_OCDEL (111x)
		57405: 5,   // TOK_SEMICOL (108x)
		57347: 6,   // TOK_WORD (104x)
		57400: 7,   // TOK_DOT (62x)
		57346: 8,   // TOK_NUM (56x)
		57404: 9,   // TOK_DEREF (49x)
		57357: 10,  // TOK_PLUS (49x)
		57368: 11,  // TOK_SLASH (48x)
		57390: 12,  // TOK_ATTRSPEC (47x)
		57359: 13,  // TOK_BOOL (43x)
		57416: 14,  // TOK_NOT (43x)
		57450: 15,  // P1 (41x)
		57449: 16,  // P (38x)
		57366: 17,  // TOK_EQUAL (38x)
		57356: 18,  // TOK_ATTR (31x)
		57414: 19,  // TOK_OR (31x)
		57413: 20,  // TOK_RPAREN (30x)
		57411: 21,  // TOK_RBLOCK (26x)
		57415: 22,  // TOK_AND (25x)
		57412: 23,  // TOK_LPAREN (25x)
		57448: 24,  // ORIENTN (24x)
		57453: 25,  // WORDORNUM (23x)
		57454: 26,  // factor (20x)
		57459: 27,  // unary (20x)
		57458: 28,  // term (16x)
		57395: 29,  // TOK_OCPSPEC (15x)
		57455: 30,  // nex (14x)
		57419: 31,  // TOK_GREATER (14x)
		57420: 32,  // TOK_LESS (14x)
		57417: 33,  // TOK_DIV (13x)
		57418: 34,  // TOK_MULT (13x)
		57452: 35,  // REL (10x)
		57406: 36,  // TOK_IF (9x)
		57408: 37,  // TOK_WHILE (9x)
		57429: 38,  // EQAL (8x)
		57371: 39,  // TOK_CD (8x)
		57360: 40,  // TOK_CREATE (8x)
		57363: 41,  // TOK_DELETE (8x)
		57361: 42,  // TOK_GET (8x)
		57375: 43,  // TOK_LS (8x)
		57377: 44,  // TOK_LSOG (8x)
		57376: 45,  // TOK_TREE (8x)
		57362: 46,  // TOK_UPDATE (8x)
		57425: 47,  // BASH (7x)
		57426: 48,  // CLSD_STMT (7x)
		57433: 49,  // JOIN (7x)
		57434: 50,  // K (7x)
		57435: 51,  // NT_CREATE (7x)
		57436: 52,  // NT_DEL (7x)
		57437: 53,  // NT_GET (7x)
		57438: 54,  // NT_UPDATE (7x)
		57439: 55,  // OCCHOOSE (7x)
		57441: 56,  // OCDEL (7x)
		57442: 57,  // OCDOT (7x)
		57443: 58,  // OCGET (7x)
		57444: 59,  // OCLISYNTX (7x)
		57445: 60,  // OCSEL (7x)
		57446: 61,  // OCUPDATE (7x)
		57451: 62,  // Q (7x)
		57457: 63,  // stmnt (7x)
		57373: 64,  // TOK_CLR (7x)
		57370: 65,  // TOK_DOC (7x)
		57369: 66,  // TOK_EXIT (7x)
		57374: 67,  // TOK_GREP (7x)
		57380: 68,  // TOK_LSBLDG (7x)
		57383: 69,  // TOK_LSDEV (7x)
		57382: 70,  // TOK_LSRACK (7x)
		57381: 71,  // TOK_LSROOM (7x)
		57379: 72,  // TOK_LSSITE (7x)
		57384: 73,  // TOK_LSSUBDEV (7x)
		57385: 74,  // TOK_LSSUBDEV1 (7x)
		57378: 75,  // TOK_LSTEN (7x)
		57372: 76,  // TOK_PWD (7x)
		57396: 77,  // TOK_SELECT (7x)
		57430: 78,  // EXPR (6x)
		57447: 79,  // OPEN_STMT (6x)
		57431: 80,  // F (5x)
		57350: 81,  // TOK_BLDG (3x)
		57353: 82,  // TOK_DEVICE (3x)
		57352: 83,  // TOK_RACK (3x)
		57398: 84,  // TOK_RBRAC (3x)
		57351: 85,  // TOK_ROOM (3x)
		57349: 86,  // TOK_SITE (3x)
		57348: 87,  // TOK_TENANT (3x)
		57427: 88,  // CTRL (2x)
		57428: 89,  // E (2x)
		57432: 90,  // GETOBJS (2x)
		57456: 91,  // start (2x)
		57410: 92,  // TOK_LBLOCK (2x)
		57354: 93,  // TOK_SUBDEVICE (2x)
		57355: 94,  // TOK_SUBDEVICE1 (2x)
		57421: 95,  // TOK_THEN (2x)
		57440: 96,  // OCCR (1x)
		57401: 97,  // TOK_CMDS (1x)
		57399: 98,  // TOK_COMMA (1x)
		57397: 99,  // TOK_LBRAC (1x)
		57386: 100, // TOK_OCBLDG (1x)
		57387: 101, // TOK_OCDEV (1x)
		57388: 102, // TOK_OCRACK (1x)
		57389: 103, // TOK_OCROOM (1x)
		57391: 104, // TOK_OCSITE (1x)
		57392: 105, // TOK_OCTENANT (1x)
		57402: 106, // TOK_TEMPLATE (1x)
		57403: 107, // TOK_VAR (1x)
		57424: 108, // $default (0x)
		57345: 109, // error (0x)
		57365: 110, // TOK_BASHTYPE (0x)
		57367: 111, // TOK_CMDFLAG (0x)
		57407: 112, // TOK_FOR (0x)
		57393: 113, // TOK_OCSDEV (0x)
		57394: 114, // TOK_OCSDEV1 (0x)
		57364: 115, // TOK_SEARCH (0x)
	}

	yySymNames = []string{
		"$end",
		"TOK_FI",
		"TOK_DONE",
		"TOK_ELSE",
		"TOK_OCDEL",
		"TOK_SEMICOL",
		"TOK_WORD",
		"TOK_DOT",
		"TOK_NUM",
		"TOK_DEREF",
		"TOK_PLUS",
		"TOK_SLASH",
		"TOK_ATTRSPEC",
		"TOK_BOOL",
		"TOK_NOT",
		"P1",
		"P",
		"TOK_EQUAL",
		"TOK_ATTR",
		"TOK_OR",
		"TOK_RPAREN",
		"TOK_RBLOCK",
		"TOK_AND",
		"TOK_LPAREN",
		"ORIENTN",
		"WORDORNUM",
		"factor",
		"unary",
		"term",
		"TOK_OCPSPEC",
		"nex",
		"TOK_GREATER",
		"TOK_LESS",
		"TOK_DIV",
		"TOK_MULT",
		"REL",
		"TOK_IF",
		"TOK_WHILE",
		"EQAL",
		"TOK_CD",
		"TOK_CREATE",
		"TOK_DELETE",
		"TOK_GET",
		"TOK_LS",
		"TOK_LSOG",
		"TOK_TREE",
		"TOK_UPDATE",
		"BASH",
		"CLSD_STMT",
		"JOIN",
		"K",
		"NT_CREATE",
		"NT_DEL",
		"NT_GET",
		"NT_UPDATE",
		"OCCHOOSE",
		"OCDEL",
		"OCDOT",
		"OCGET",
		"OCLISYNTX",
		"OCSEL",
		"OCUPDATE",
		"Q",
		"stmnt",
		"TOK_CLR",
		"TOK_DOC",
		"TOK_EXIT",
		"TOK_GREP",
		"TOK_LSBLDG",
		"TOK_LSDEV",
		"TOK_LSRACK",
		"TOK_LSROOM",
		"TOK_LSSITE",
		"TOK_LSSUBDEV",
		"TOK_LSSUBDEV1",
		"TOK_LSTEN",
		"TOK_PWD",
		"TOK_SELECT",
		"EXPR",
		"OPEN_STMT",
		"F",
		"TOK_BLDG",
		"TOK_DEVICE",
		"TOK_RACK",
		"TOK_RBRAC",
		"TOK_ROOM",
		"TOK_SITE",
		"TOK_TENANT",
		"CTRL",
		"E",
		"GETOBJS",
		"start",
		"TOK_LBLOCK",
		"TOK_SUBDEVICE",
		"TOK_SUBDEVICE1",
		"TOK_THEN",
		"OCCR",
		"TOK_CMDS",
		"TOK_COMMA",
		"TOK_LBRAC",
		"TOK_OCBLDG",
		"TOK_OCDEV",
		"TOK_OCRACK",
		"TOK_OCROOM",
		"TOK_OCSITE",
		"TOK_OCTENANT",
		"TOK_TEMPLATE",
		"TOK_VAR",
		"$default",
		"error",
		"TOK_BASHTYPE",
		"TOK_CMDFLAG",
		"TOK_FOR",
		"TOK_OCSDEV",
		"TOK_OCSDEV1",
		"TOK_SEARCH",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:   {0, 1},
		1:   {91, 1},
		2:   {91, 3},
		3:   {91, 1},
		4:   {63, 1},
		5:   {63, 1},
		6:   {63, 1},
		7:   {63, 0},
		8:   {88, 1},
		9:   {88, 1},
		10:  {79, 7},
		11:  {79, 7},
		12:  {79, 9},
		13:  {79, 6},
		14:  {48, 1},
		15:  {48, 9},
		16:  {48, 6},
		17:  {78, 3},
		18:  {78, 1},
		19:  {49, 3},
		20:  {49, 1},
		21:  {38, 4},
		22:  {38, 4},
		23:  {38, 1},
		24:  {35, 3},
		25:  {35, 4},
		26:  {35, 4},
		27:  {35, 4},
		28:  {35, 3},
		29:  {35, 1},
		30:  {30, 3},
		31:  {30, 3},
		32:  {30, 1},
		33:  {28, 3},
		34:  {28, 3},
		35:  {28, 1},
		36:  {27, 2},
		37:  {27, 2},
		38:  {27, 1},
		39:  {26, 3},
		40:  {26, 1},
		41:  {26, 1},
		42:  {26, 1},
		43:  {50, 1},
		44:  {50, 1},
		45:  {50, 1},
		46:  {50, 1},
		47:  {51, 3},
		48:  {51, 4},
		49:  {53, 2},
		50:  {53, 3},
		51:  {54, 3},
		52:  {52, 2},
		53:  {89, 1},
		54:  {89, 1},
		55:  {89, 1},
		56:  {89, 1},
		57:  {89, 1},
		58:  {89, 1},
		59:  {89, 1},
		60:  {89, 1},
		61:  {24, 1},
		62:  {24, 1},
		63:  {24, 0},
		64:  {25, 1},
		65:  {25, 1},
		66:  {25, 4},
		67:  {25, 1},
		68:  {80, 4},
		69:  {80, 3},
		70:  {16, 1},
		71:  {16, 2},
		72:  {15, 3},
		73:  {15, 1},
		74:  {15, 4},
		75:  {15, 1},
		76:  {15, 2},
		77:  {15, 1},
		78:  {15, 2},
		79:  {15, 0},
		80:  {62, 2},
		81:  {62, 2},
		82:  {62, 2},
		83:  {62, 2},
		84:  {62, 2},
		85:  {62, 2},
		86:  {62, 2},
		87:  {62, 2},
		88:  {62, 2},
		89:  {62, 2},
		90:  {62, 2},
		91:  {62, 2},
		92:  {62, 3},
		93:  {62, 1},
		94:  {47, 1},
		95:  {47, 1},
		96:  {47, 1},
		97:  {47, 1},
		98:  {47, 1},
		99:  {47, 1},
		100: {47, 2},
		101: {47, 2},
		102: {47, 2},
		103: {47, 2},
		104: {47, 2},
		105: {47, 2},
		106: {47, 2},
		107: {47, 2},
		108: {47, 2},
		109: {59, 2},
		110: {59, 1},
		111: {59, 1},
		112: {59, 1},
		113: {59, 1},
		114: {59, 1},
		115: {59, 1},
		116: {96, 5},
		117: {96, 5},
		118: {96, 5},
		119: {96, 5},
		120: {96, 7},
		121: {96, 7},
		122: {96, 7},
		123: {96, 7},
		124: {96, 7},
		125: {96, 7},
		126: {96, 7},
		127: {96, 7},
		128: {56, 2},
		129: {61, 5},
		130: {58, 2},
		131: {90, 3},
		132: {90, 1},
		133: {55, 4},
		134: {57, 6},
		135: {57, 4},
		136: {57, 4},
		137: {57, 2},
		138: {60, 1},
		139: {60, 5},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [304][]uint16{
		// 0
		{133, 4: 163, 133, 161, 162, 9: 164, 183, 160, 15: 159, 190, 191, 36: 149, 150, 39: 165, 155, 158, 156, 166, 179, 175, 157, 176, 148, 50: 144, 151, 154, 152, 153, 187, 184, 188, 186, 146, 189, 185, 145, 142, 177, 182, 181, 178, 169, 172, 171, 170, 168, 173, 174, 167, 180, 192, 79: 147, 88: 143, 91: 141},
		{140},
		{139, 5: 442},
		{137},
		{136, 136, 136, 136, 5: 136},
		// 5
		{135, 135, 135, 135, 5: 135},
		{134, 134, 134, 134, 5: 134},
		{132},
		{131},
		{23: 413, 92: 412},
		// 10
		{23: 361},
		{97, 97, 97, 97, 5: 97},
		{96, 96, 96, 96, 5: 96},
		{95, 95, 95, 95, 5: 95},
		{94, 94, 94, 94, 5: 94},
		// 15
		{81: 350, 353, 352, 85: 351, 349, 348, 89: 357, 93: 354, 355},
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 346, 81: 350, 353, 352, 85: 351, 349, 348, 89: 347, 93: 354, 355},
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 15: 159, 340, 18: 61},
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 339},
		{70, 70, 70, 70, 5: 70, 7: 70, 70, 12: 70, 18: 70},
		// 20
		{61, 61, 61, 61, 198, 61, 161, 197, 61, 199, 12: 61, 15: 338, 18: 61},
		{67, 67, 67, 67, 5: 67, 7: 67, 67, 11: 336, 67, 18: 67},
		{7: 208, 97: 326, 106: 327, 325},
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 324},
		{6: 323},
		// 25
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 322},
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 321},
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 320},
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 319},
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 318},
		// 30
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 317},
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 316},
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 315},
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 314},
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 313},
		// 35
		{61, 61, 61, 61, 198, 61, 161, 197, 310, 199, 11: 160, 15: 159, 311},
		{47, 47, 47, 47, 5: 47},
		{46, 46, 46, 46, 5: 46},
		{45, 45, 45, 45, 5: 45},
		{44, 44, 44, 44, 5: 44},
		// 40
		{43, 43, 43, 43, 5: 43},
		{42, 42, 42, 42, 5: 42},
		{41, 41, 41, 41, 5: 41, 307, 39: 302, 303, 306, 304, 301, 309, 308, 305},
		{81: 230, 236, 234, 85: 232, 228, 226, 96: 224, 100: 229, 235, 233, 231, 227, 225},
		{30, 30, 30, 30, 5: 30},
		// 45
		{29, 29, 29, 29, 5: 29},
		{28, 28, 28, 28, 5: 28},
		{27, 27, 27, 27, 5: 27},
		{26, 26, 26, 26, 5: 26},
		{25, 25, 25, 25, 5: 25},
		// 50
		{7: 211},
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 200, 99: 201},
		{2, 2, 2, 2, 5: 2, 7: 193},
		{18: 194},
		{17: 195},
		// 55
		{6: 196},
		{1, 1, 1, 1, 5: 1},
		{7: 208},
		{63, 63, 63, 63, 5: 63, 7: 63, 63, 12: 63, 18: 63},
		{6: 207},
		// 60
		{10, 10, 10, 10, 5: 10},
		{6: 202, 90: 203},
		{84: 8, 98: 205},
		{84: 204},
		{7, 7, 7, 7, 5: 7},
		// 65
		{6: 202, 90: 206},
		{84: 9},
		{62, 62, 62, 62, 5: 62, 7: 62, 62, 12: 62, 18: 62},
		{64, 64, 64, 64, 5: 64, 7: 64, 64, 11: 209, 64, 18: 64},
		{61, 61, 61, 61, 198, 61, 161, 197, 61, 199, 12: 61, 15: 210, 18: 61},
		// 70
		{66, 66, 66, 66, 5: 66, 7: 66, 66, 12: 66, 18: 66},
		{18: 212},
		{17: 213},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 220},
		{6: 79},
		// 75
		{6: 78},
		{76, 76, 76, 76, 5: 76, 12: 76, 18: 76},
		{75, 75, 75, 75, 5: 75, 12: 75, 18: 75},
		{6: 221},
		{73, 73, 73, 73, 5: 73, 12: 73, 18: 73},
		// 80
		{11, 11, 11, 11, 5: 11},
		{4: 215, 6: 77, 10: 214, 24: 222},
		{6: 223},
		{74, 74, 74, 74, 5: 74, 12: 74, 18: 74},
		{31, 31, 31, 31, 5: 31},
		// 85
		{29: 297},
		{29: 293},
		{29: 289},
		{29: 285},
		{29: 279},
		// 90
		{29: 273},
		{29: 267},
		{29: 261},
		{29: 255},
		{29: 249},
		// 95
		{29: 243},
		{29: 237},
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 61, 15: 159, 238},
		{12: 239},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 240},
		// 100
		{12: 241},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 242},
		{13, 13, 13, 13, 5: 13},
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 61, 15: 159, 244},
		{12: 245},
		// 105
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 246},
		{12: 247},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 248},
		{14, 14, 14, 14, 5: 14},
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 61, 15: 159, 250},
		// 110
		{12: 251},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 252},
		{12: 253},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 254},
		{15, 15, 15, 15, 5: 15},
		// 115
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 61, 15: 159, 256},
		{12: 257},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 258},
		{12: 259},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 260},
		// 120
		{16, 16, 16, 16, 5: 16},
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 61, 15: 159, 262},
		{12: 263},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 264},
		{12: 265},
		// 125
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 266},
		{17, 17, 17, 17, 5: 17},
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 61, 15: 159, 268},
		{12: 269},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 270},
		// 130
		{12: 271},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 272},
		{18, 18, 18, 18, 5: 18},
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 61, 15: 159, 274},
		{12: 275},
		// 135
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 276},
		{12: 277},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 278},
		{19, 19, 19, 19, 5: 19},
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 61, 15: 159, 280},
		// 140
		{12: 281},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 282},
		{12: 283},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 284},
		{20, 20, 20, 20, 5: 20},
		// 145
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 61, 15: 159, 286},
		{12: 287},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 288},
		{21, 21, 21, 21, 5: 21},
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 61, 15: 159, 290},
		// 150
		{12: 291},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 292},
		{22, 22, 22, 22, 5: 22},
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 61, 15: 159, 294},
		{12: 295},
		// 155
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 296},
		{23, 23, 23, 23, 5: 23},
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 61, 15: 159, 298},
		{12: 299},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 300},
		// 160
		{24, 24, 24, 24, 5: 24},
		{40, 40, 40, 40, 5: 40},
		{39, 39, 39, 39, 5: 39},
		{38, 38, 38, 38, 5: 38},
		{37, 37, 37, 37, 5: 37},
		// 165
		{36, 36, 36, 36, 5: 36},
		{35, 35, 35, 35, 5: 35},
		{34, 34, 34, 34, 5: 34},
		{33, 33, 33, 33, 5: 33},
		{32, 32, 32, 32, 5: 32},
		// 170
		{50, 50, 50, 50, 5: 50},
		{49, 49, 49, 49, 5: 49, 8: 312},
		{48, 48, 48, 48, 5: 48},
		{51, 51, 51, 51, 5: 51},
		{52, 52, 52, 52, 5: 52},
		// 175
		{53, 53, 53, 53, 5: 53},
		{54, 54, 54, 54, 5: 54},
		{55, 55, 55, 55, 5: 55},
		{56, 56, 56, 56, 5: 56},
		{57, 57, 57, 57, 5: 57},
		// 180
		{58, 58, 58, 58, 5: 58},
		{59, 59, 59, 59, 5: 59},
		{60, 60, 60, 60, 5: 60},
		{3, 3, 3, 3, 5: 3, 7: 62},
		{12, 12, 12, 12, 5: 12},
		// 185
		{29: 332},
		{29: 330},
		{29: 328},
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 329},
		{4, 4, 4, 4, 5: 4},
		// 190
		{61, 61, 61, 61, 198, 61, 161, 197, 9: 199, 11: 160, 15: 159, 331},
		{5, 5, 5, 5, 5: 5},
		{6: 333},
		{17: 334},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 335},
		// 195
		{6, 6, 6, 6, 5: 6},
		{61, 61, 61, 61, 198, 61, 161, 197, 61, 199, 12: 61, 15: 337, 18: 61},
		{68, 68, 68, 68, 5: 68, 7: 68, 68, 12: 68, 18: 68},
		{69, 69, 69, 69, 5: 69, 7: 69, 69, 12: 69, 18: 69},
		{88, 88, 88, 88, 5: 88},
		// 200
		{18: 342, 80: 341},
		{89, 89, 89, 89, 5: 89},
		{17: 343},
		{4: 215, 6: 216, 8: 217, 10: 214, 13: 219, 24: 218, 344},
		{71, 71, 71, 71, 5: 71, 18: 342, 80: 345},
		// 205
		{72, 72, 72, 72, 5: 72},
		{91, 91, 91, 91, 5: 91},
		{18: 342, 80: 356},
		{4: 87, 6: 87, 87, 9: 87, 11: 87, 18: 87},
		{4: 86, 6: 86, 86, 9: 86, 11: 86, 18: 86},
		// 210
		{4: 85, 6: 85, 85, 9: 85, 11: 85, 18: 85},
		{4: 84, 6: 84, 84, 9: 84, 11: 84, 18: 84},
		{4: 83, 6: 83, 83, 9: 83, 11: 83, 18: 83},
		{4: 82, 6: 82, 82, 9: 82, 11: 82, 18: 82},
		{4: 81, 6: 81, 81, 9: 81, 11: 81, 18: 81},
		// 215
		{4: 80, 6: 80, 80, 9: 80, 11: 80, 18: 80},
		{90, 90, 90, 90, 5: 90},
		{4: 198, 6: 161, 197, 9: 199, 11: 160, 15: 159, 359, 18: 342, 80: 358},
		{93, 93, 93, 93, 5: 93},
		{18: 342, 80: 360},
		// 220
		{92, 92, 92, 92, 5: 92},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 367, 30: 366, 35: 365, 38: 364, 49: 363, 78: 362},
		{19: 377, 406},
		{19: 122, 122, 122, 380},
		{14: 383, 17: 382, 19: 120, 120, 120, 120},
		// 225
		{14: 117, 17: 117, 19: 117, 117, 117, 117},
		{4: 397, 10: 396, 14: 111, 17: 111, 19: 111, 111, 111, 111, 31: 395, 394},
		{4: 108, 10: 108, 14: 108, 17: 108, 19: 108, 108, 108, 108, 31: 108, 108, 391, 390},
		{4: 105, 10: 105, 14: 105, 17: 105, 19: 105, 105, 105, 105, 31: 105, 105, 105, 105},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 389},
		// 230
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 388},
		{4: 102, 10: 102, 14: 102, 17: 102, 19: 102, 102, 102, 102, 31: 102, 102, 102, 102},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 367, 30: 366, 35: 365, 38: 364, 49: 363, 78: 376},
		{4: 100, 10: 100, 14: 100, 17: 100, 19: 100, 100, 100, 100, 31: 100, 100, 100, 100},
		{4: 99, 10: 99, 14: 99, 17: 99, 19: 99, 99, 99, 99, 31: 99, 99, 99, 99},
		// 235
		{4: 98, 10: 98, 14: 98, 17: 98, 19: 98, 98, 98, 98, 31: 98, 98, 98, 98},
		{19: 377, 378},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 367, 30: 366, 35: 365, 38: 364, 49: 379},
		{4: 101, 10: 101, 14: 101, 17: 101, 19: 101, 101, 101, 101, 31: 101, 101, 101, 101},
		{19: 123, 123, 123, 380},
		// 240
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 367, 30: 366, 35: 365, 38: 381},
		{14: 383, 17: 382, 19: 121, 121, 121, 121},
		{17: 386},
		{17: 384},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 367, 30: 366, 35: 385},
		// 245
		{14: 118, 17: 118, 19: 118, 118, 118, 118},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 367, 30: 366, 35: 387},
		{14: 119, 17: 119, 19: 119, 119, 119, 119},
		{4: 103, 10: 103, 14: 103, 17: 103, 19: 103, 103, 103, 103, 31: 103, 103, 103, 103},
		{4: 104, 10: 104, 14: 104, 17: 104, 19: 104, 104, 104, 104, 31: 104, 104, 104, 104},
		// 250
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 393},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 392},
		{4: 106, 10: 106, 14: 106, 17: 106, 19: 106, 106, 106, 106, 31: 106, 106, 106, 106},
		{4: 107, 10: 107, 14: 107, 17: 107, 19: 107, 107, 107, 107, 31: 107, 107, 107, 107},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 17: 404, 23: 372, 26: 371, 368, 367, 30: 403},
		// 255
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 17: 400, 23: 372, 26: 371, 368, 367, 30: 401},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 399},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 398},
		{4: 109, 10: 109, 14: 109, 17: 109, 19: 109, 109, 109, 109, 31: 109, 109, 391, 390},
		{4: 110, 10: 110, 14: 110, 17: 110, 19: 110, 110, 110, 110, 31: 110, 110, 391, 390},
		// 260
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 367, 30: 402},
		{4: 397, 10: 396, 14: 112, 17: 112, 19: 112, 112, 112, 112},
		{4: 397, 10: 396, 14: 114, 17: 114, 19: 114, 114, 114, 114},
		{4: 397, 10: 396, 14: 116, 17: 116, 19: 116, 116, 116, 116},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 367, 30: 405},
		// 265
		{4: 397, 10: 396, 14: 115, 17: 115, 19: 115, 115, 115, 115},
		{2: 133, 4: 163, 6: 161, 162, 9: 164, 183, 160, 15: 159, 190, 191, 36: 149, 150, 39: 165, 155, 158, 156, 166, 179, 175, 157, 176, 409, 50: 144, 151, 154, 152, 153, 187, 184, 188, 186, 146, 189, 185, 145, 408, 177, 182, 181, 178, 169, 172, 171, 170, 168, 173, 174, 167, 180, 192, 79: 407},
		{2: 411},
		{1: 126, 126, 126},
		{2: 410},
		// 270
		{124, 124, 124, 124},
		{127, 127, 127},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 367, 30: 366, 35: 365, 38: 364, 49: 363, 78: 426},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 367, 30: 366, 35: 365, 38: 364, 49: 363, 78: 414},
		{19: 377, 415},
		// 275
		{95: 416},
		{3: 133, 163, 6: 161, 162, 9: 164, 183, 160, 15: 159, 190, 191, 36: 417, 419, 39: 165, 155, 158, 156, 166, 179, 175, 157, 176, 418, 50: 144, 151, 154, 152, 153, 187, 184, 188, 186, 146, 189, 185, 145, 408, 177, 182, 181, 178, 169, 172, 171, 170, 168, 173, 174, 167, 180, 192},
		{23: 413},
		{3: 423},
		{23: 420},
		// 280
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 367, 30: 366, 35: 365, 38: 364, 49: 363, 78: 421},
		{19: 377, 422},
		{2: 133, 4: 163, 6: 161, 162, 9: 164, 183, 160, 15: 159, 190, 191, 36: 417, 419, 39: 165, 155, 158, 156, 166, 179, 175, 157, 176, 409, 50: 144, 151, 154, 152, 153, 187, 184, 188, 186, 146, 189, 185, 145, 408, 177, 182, 181, 178, 169, 172, 171, 170, 168, 173, 174, 167, 180, 192},
		{1: 133, 4: 163, 6: 161, 162, 9: 164, 183, 160, 15: 159, 190, 191, 36: 417, 419, 39: 165, 155, 158, 156, 166, 179, 175, 157, 176, 424, 50: 144, 151, 154, 152, 153, 187, 184, 188, 186, 146, 189, 185, 145, 408, 177, 182, 181, 178, 169, 172, 171, 170, 168, 173, 174, 167, 180, 192},
		{1: 425},
		// 285
		{125, 125, 125, 125},
		{19: 377, 21: 427},
		{95: 428},
		{1: 133, 3: 133, 163, 6: 161, 162, 9: 164, 183, 160, 15: 159, 190, 191, 36: 149, 150, 39: 165, 155, 158, 156, 166, 179, 175, 157, 176, 431, 50: 144, 151, 154, 152, 153, 187, 184, 188, 186, 146, 189, 185, 145, 429, 177, 182, 181, 178, 169, 172, 171, 170, 168, 173, 174, 167, 180, 192, 79: 430},
		{1: 441, 3: 126},
		// 290
		{1: 440},
		{3: 432},
		{36: 433, 435, 79: 434},
		{92: 412},
		{1: 439},
		// 295
		{23: 436},
		{4: 370, 6: 374, 8: 373, 13: 375, 369, 23: 372, 26: 371, 368, 367, 30: 366, 35: 365, 38: 364, 49: 363, 78: 437},
		{19: 377, 438},
		{36: 433, 435, 79: 407},
		{128, 128, 128},
		// 300
		{129, 129, 129},
		{130, 130, 130},
		{133, 4: 163, 133, 161, 162, 9: 164, 183, 160, 15: 159, 190, 191, 36: 149, 150, 39: 165, 155, 158, 156, 166, 179, 175, 157, 176, 148, 50: 144, 151, 154, 152, 153, 187, 184, 188, 186, 146, 189, 185, 145, 142, 177, 182, 181, 178, 169, 172, 171, 170, 168, 173, 174, 167, 180, 192, 79: 147, 88: 143, 91: 443},
		{138},
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
	const yyError = 109

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
			println("OGREE: Unrecognised command!")
cmd.WarningLogger.Println("Unknown Command")			/*yylex.Error(msg)*/
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
	case 6:
		{
			x := yyS[yypt-0].comNode
			println("now calling exectute")
			x.execute()
		}
	case 43:
		{
			println("@State start")
		}
	case 47:
		{
			cmd.PostObj(cmd.EntityStrToInt(yyS[yypt-1].s), yyS[yypt-1].s, resMap(&yyS[yypt-0].s))
		}
	case 48:
		{
			yyVAL.s = yyS[yypt-0].s
			cmd.Disp(resMap(&yyS[yypt-0].s))
			cmd.PostObj(cmd.EntityStrToInt(yyS[yypt-2].s), yyS[yypt-2].s, resMap(&yyS[yypt-0].s))
		}
	case 49:
		{
			cmd.GetObject(yyS[yypt-0].s)
		}
	case 50:
		{ /*cmd.Disp(resMap(&$4)); */
			cmd.SearchObjects(yyS[yypt-1].s, resMap(&yyS[yypt-0].s))
		}
	case 51:
		{
			yyVAL.s = yyS[yypt-0].s /*cmd.Disp(resMap(&$4));*/
			cmd.UpdateObj(yyS[yypt-1].s, resMap(&yyS[yypt-0].s))
		}
	case 52:
		{
			println("@State NT_DEL")
			cmd.DeleteObj(yyS[yypt-0].s)
		}
	case 61:
		{
			yyVAL.s = yyS[yypt-0].s
		}
	case 62:
		{
			yyVAL.s = yyS[yypt-0].s
		}
	case 63:
		{
			yyVAL.s = ""
		}
	case 64:
		{
			yyVAL.s = yyS[yypt-0].s
			dCatchPtr = yyS[yypt-0].s
		}
	case 65:
		{
			x := strconv.Itoa(yyS[yypt-0].n)
			yyVAL.s = x
			dCatchPtr = yyS[yypt-0].n
		}
	case 66:
		{
			yyVAL.s = yyS[yypt-3].s + yyS[yypt-2].s + yyS[yypt-1].s + yyS[yypt-0].s
			dCatchPtr = yyS[yypt-3].s + yyS[yypt-2].s + yyS[yypt-1].s + yyS[yypt-0].s
		}
	case 67:
		{
			var x bool
			if yyS[yypt-0].s == "false" {
				x = false
			} else {
				x = true
			}
			dCatchPtr = x
		}
	case 68:
		{
			yyVAL.s = string(yyS[yypt-3].s + "=" + yyS[yypt-1].s + "=" + yyS[yypt-0].s)
			println("So we got: ", yyVAL.s)
		}
	case 69:
		{
			yyVAL.s = yyS[yypt-2].s + "=" + yyS[yypt-0].s
		}
	case 71:
		{
			yyVAL.s = "/" + yyS[yypt-0].s
		}
	case 72:
		{
			yyVAL.s = yyS[yypt-2].s + "/" + yyS[yypt-0].s
		}
	case 73:
		{
			yyVAL.s = yyS[yypt-0].s
		}
	case 74:
		{
			yyVAL.s = "../" + yyS[yypt-0].s
		}
	case 75:
		{
			yyVAL.s = yyS[yypt-0].s
		}
	case 76:
		{
			yyVAL.s = ".."
		}
	case 77:
		{
			yyVAL.s = "-"
		}
	case 78:
		{
			yyVAL.s = ""
		}
	case 79:
		{
			yyVAL.s = ""
		}
	case 80:
		{
			cmd.CD(yyS[yypt-0].s)
		}
	case 81:
		{
			cmd.LS(yyS[yypt-0].s)
		}
	case 82:
		{
			cmd.LSOBJECT(yyS[yypt-0].s, 0)
		}
	case 83:
		{
			cmd.LSOBJECT(yyS[yypt-0].s, 1)
		}
	case 84:
		{
			cmd.LSOBJECT(yyS[yypt-0].s, 2)
		}
	case 85:
		{
			cmd.LSOBJECT(yyS[yypt-0].s, 3)
		}
	case 86:
		{
			cmd.LSOBJECT(yyS[yypt-0].s, 4)
		}
	case 87:
		{
			cmd.LSOBJECT(yyS[yypt-0].s, 5)
		}
	case 88:
		{
			cmd.LSOBJECT(yyS[yypt-0].s, 6)
		}
	case 89:
		{
			cmd.LSOBJECT(yyS[yypt-0].s, 7)
		}
	case 90:
		{
			cmd.Tree("", yyS[yypt-0].n)
		}
	case 91:
		{
			cmd.Tree(yyS[yypt-0].s, 0)
		}
	case 92:
		{
			cmd.Tree(yyS[yypt-1].s, yyS[yypt-0].n)
		}
	case 93:
		{
			cmd.Execute()
		}
	case 96:
		{
			cmd.LSOG()
		}
	case 97:
		{
			cmd.PWD()
		}
	case 98:
		{
			cmd.Exit()
		}
	case 99:
		{
			cmd.Help("")
		}
	case 100:
		{
			cmd.Help("ls")
		}
	case 101:
		{
			cmd.Help("cd")
		}
	case 102:
		{
			cmd.Help("create")
		}
	case 103:
		{
			cmd.Help("gt")
		}
	case 104:
		{
			cmd.Help("update")
		}
	case 105:
		{
			cmd.Help("delete")
		}
	case 106:
		{
			cmd.Help(yyS[yypt-0].s)
		}
	case 107:
		{
			cmd.Help("tree")
		}
	case 108:
		{
			cmd.Help("lsog")
		}
	case 109:
		{
			yyVAL.comNode = &commonNode{COMMON, cmd.ShowClipBoard, "select", nil}
		}
	case 110:
		{
			yyVAL.comNode = yyS[yypt-0].comNode
		}
	case 111:
		{
			yyVAL.comNode = &commonNode{COMMON, cmd.ShowClipBoard, "select", nil}
		}
	case 112:
		{
			yyVAL.comNode = yyS[yypt-0].comNode
		}
	case 113:
		{
			yyVAL.comNode = &commonNode{COMMON, nil, "select", nil}
		}
	case 114:
		{
			yyVAL.comNode = &commonNode{COMMON, cmd.ShowClipBoard, "select", nil}
		}
	case 115:
		{
			yyVAL.comNode = yyS[yypt-0].comNode
			println("Alright")
		}
	case 116:
		{
			cmd.GetOCLIAtrributes(cmd.StrToStack(replaceOCLICurrPath(yyS[yypt-2].s)), cmd.TENANT, map[string]interface{}{"attributes": map[string]interface{}{"color": yyS[yypt-0].s}}, rlPtr)
		}
	case 117:
		{
			cmd.GetOCLIAtrributes(cmd.StrToStack(replaceOCLICurrPath(yyS[yypt-2].s)), cmd.TENANT, map[string]interface{}{"attributes": map[string]interface{}{"color": yyS[yypt-0].s}}, rlPtr)
		}
	case 118:
		{
			cmd.GetOCLIAtrributes(cmd.StrToStack(replaceOCLICurrPath(yyS[yypt-2].s)), cmd.SITE, map[string]interface{}{"attributes": map[string]interface{}{"orientation": yyS[yypt-0].s}}, rlPtr)
		}
	case 119:
		{
			cmd.GetOCLIAtrributes(cmd.StrToStack(replaceOCLICurrPath(yyS[yypt-2].s)), cmd.SITE, map[string]interface{}{"attributes": map[string]interface{}{"orientation": yyS[yypt-0].s}}, rlPtr)
		}
	case 120:
		{
			cmd.GetOCLIAtrributes(cmd.StrToStack(replaceOCLICurrPath(yyS[yypt-4].s)), cmd.BLDG, map[string]interface{}{"attributes": map[string]interface{}{"posXY": yyS[yypt-2].s, "size": yyS[yypt-0].s}}, rlPtr)
		}
	case 121:
		{
			cmd.GetOCLIAtrributes(cmd.StrToStack(replaceOCLICurrPath(yyS[yypt-4].s)), cmd.BLDG, map[string]interface{}{"attributes": map[string]interface{}{"posXY": yyS[yypt-2].s, "size": yyS[yypt-0].s}}, rlPtr)
		}
	case 122:
		{
			cmd.GetOCLIAtrributes(cmd.StrToStack(replaceOCLICurrPath(yyS[yypt-4].s)), cmd.ROOM, map[string]interface{}{"attributes": map[string]interface{}{"posXY": yyS[yypt-2].s, "size": yyS[yypt-0].s}}, rlPtr)
		}
	case 123:
		{
			cmd.GetOCLIAtrributes(cmd.StrToStack(replaceOCLICurrPath(yyS[yypt-4].s)), cmd.ROOM, map[string]interface{}{"attributes": map[string]interface{}{"posXY": yyS[yypt-2].s, "size": yyS[yypt-0].s}}, rlPtr)
		}
	case 124:
		{
			cmd.GetOCLIAtrributes(cmd.StrToStack(replaceOCLICurrPath(yyS[yypt-4].s)), cmd.RACK, map[string]interface{}{"attributes": map[string]interface{}{"posXY": yyS[yypt-2].s, "size": yyS[yypt-0].s}}, rlPtr)
		}
	case 125:
		{
			cmd.GetOCLIAtrributes(cmd.StrToStack(replaceOCLICurrPath(yyS[yypt-4].s)), cmd.RACK, map[string]interface{}{"attributes": map[string]interface{}{"posXY": yyS[yypt-2].s, "size": yyS[yypt-0].s}}, rlPtr)
		}
	case 126:
		{
			cmd.GetOCLIAtrributes(cmd.StrToStack(replaceOCLICurrPath(yyS[yypt-4].s)), cmd.DEVICE, map[string]interface{}{"attributes": map[string]interface{}{"slot": yyS[yypt-2].s, "sizeUnit": yyS[yypt-0].s}}, rlPtr)
		}
	case 127:
		{
			cmd.GetOCLIAtrributes(cmd.StrToStack(replaceOCLICurrPath(yyS[yypt-4].s)), cmd.DEVICE, map[string]interface{}{"attributes": map[string]interface{}{"slot": yyS[yypt-2].s, "sizeUnit": yyS[yypt-0].s}}, rlPtr)
		}
	case 128:
		{
			yyVAL.comNode = &commonNode{COMMON, cmd.DeleteObj, "delete", replaceOCLICurrPath(yyS[yypt-0].s)}
		}
	case 129:
		{
			println("Attribute Acquired")
			val := yyS[yypt-2].s + "=" + yyS[yypt-0].s
			cmd.UpdateObj(replaceOCLICurrPath(yyS[yypt-4].s), resMap(&val))
		}
	case 130:
		{
			yyVAL.comNode = &commonNode{COMMON, cmd.GetObject, "get", replaceOCLICurrPath(yyS[yypt-0].s)}
		}
	case 131:
		{
			x := make([]string, 0)
			x = append(x, cmd.State.CurrPath+"/"+yyS[yypt-2].s)
			x = append(x, yyS[yypt-0].sarr...)
			yyVAL.sarr = x
		}
	case 132:
		{
			yyVAL.sarr = []string{cmd.State.CurrPath + "/" + yyS[yypt-0].s}
		}
	case 133:
		{
			cmd.State.ClipBoard = &yyS[yypt-1].sarr
			println("Selection made!")
		}
	case 134:
		{
			dynamicMap[yyS[yypt-2].s] = varCtr
			dynamicSymbolTable[varCtr] = dCatchPtr
			varCtr += 1
			switch dCatchPtr.(type) {
			case string:
				x := dCatchPtr.(string)
				println("You want to assign", yyS[yypt-2].s, "with value of", x)
			case int:
				x := dCatchPtr.(int)
				println("You want to assign", yyS[yypt-2].s, "with value of", x)
			case bool:
				x := dCatchPtr.(bool)
				println("You want to assign", yyS[yypt-2].s, "with value of", x)
			case float64, float32:
				x := dCatchPtr.(float64)
				println("You want to assign", yyS[yypt-2].s, "with value of", x)
			}
		}
	case 135:
		{
			cmd.LoadFile(yyS[yypt-0].s)
		}
	case 136:
		{
			cmd.LoadFile(yyS[yypt-0].s)
		}
	case 137:
		{
			v := dynamicSymbolTable[dynamicMap[yyS[yypt-0].s]]
			switch v.(type) {
			case string:
				x := v.(string)
				println("So You want the value: ", x)
			case int:
				x := v.(int)
				println("So You want the value: ", x)
			case bool:
				x := v.(bool)
				println("So You want the value: ", x)
			case float64, float32:
				x := dCatchPtr.(float64)
				println("So You want the value: ", x)
			}
		}
	case 138:
		{
			yyVAL.comNode = &commonNode{COMMON, cmd.ShowClipBoard, "select", nil}
			println("So we haven't done anythig")
		}
	case 139:
		{
			x := yyS[yypt-2].s + "=" + yyS[yypt-0].s
			yyVAL.comNode = &commonNode{COMMON, cmd.UpdateSelection, "select", resMap(&x)}
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
