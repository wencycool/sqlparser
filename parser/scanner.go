// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2013 The ql Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSES/QL-LICENSE file.

// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package parser

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ruiaylin/sqlparser/ast"
	"github.com/ruiaylin/sqlparser/util/charset"
	"github.com/ruiaylin/sqlparser/util/stringutil"
)

type lexer struct {
	c            int
	col          int
	errs         []error
	expr         ast.ExprNode
	i            int
	lcol         int
	line         int
	ncol         int
	nline        int
	sc           int
	src          string
	val          []byte
	ungetBuf     []byte
	stmtStartPos int
	stringLit    []byte

	// record token's offset of the input
	tokenEndOffset   int
	tokenStartOffset int
}

// NewLexer builds a new lexer.
func NewLexer(src string) (l *lexer) {
	l = &lexer{
		src:   src,
		nline: 1,
		ncol:  0,
	}
	l.next()
	return
}

func (l *lexer) Errors() []error {
	return l.errs
}

func (l *lexer) Expr() ast.ExprNode {
	return l.expr
}

func (l *lexer) unget(b byte) {
	l.ungetBuf = append(l.ungetBuf, b)
	l.i--
	l.ncol--
	l.tokenEndOffset--
}

func (l *lexer) next() int {
	if un := len(l.ungetBuf); un > 0 {
		nc := l.ungetBuf[0]
		l.ungetBuf = l.ungetBuf[1:]
		l.c = int(nc)
		return l.c
	}

	if l.c != 0 {
		l.val = append(l.val, byte(l.c))
	}
	l.c = 0
	if l.i < len(l.src) {
		l.c = int(l.src[l.i])
		l.i++
	}
	switch l.c {
	case '\n':
		l.lcol = l.ncol
		l.nline++
		l.ncol = 0
	default:
		l.ncol++
	}
	l.tokenEndOffset++
	return l.c
}

func (l *lexer) Errorf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	err := fmt.Errorf("line %d column %d near \"%s\"%s", l.line, l.col, l.val, s)
	l.errs = append(l.errs, err)
}

func (l *lexer) stmtText() string {
	endPos := l.i
	if l.src[l.i-1] == '\n' {
		endPos = l.i - 1 // trim new line
	}
	if l.src[l.stmtStartPos] == '\n' {
		l.stmtStartPos++
	}

	text := l.src[l.stmtStartPos:endPos]

	l.stmtStartPos = l.i
	return text
}

func (l *lexer) Lex(lval *yySymType) (r int) {
	defer func() {
		lval.offset = l.tokenStartOffset
		l.tokenStartOffset = l.tokenEndOffset
	}()
	const (
		INITIAL = iota
		S1
		S2
		S3
		S4
	)

	c0, c := 0, l.c

yystate0:

	l.val = l.val[:0]
	c0, l.line, l.col = l.c, l.nline, l.ncol

	switch yyt := l.sc; yyt {
	default:
		panic(fmt.Errorf(`invalid start condition %d`, yyt))
	case 0: // start condition: INITIAL
		goto yystart1
	case 1: // start condition: S1
		goto yystart1411
	case 2: // start condition: S2
		goto yystart1417
	case 3: // start condition: S3
		goto yystart1423
	case 4: // start condition: S4
		goto yystart1426
	}

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = l.next()
yystart1:
	switch {
	default:
		goto yystate3 // c >= '\x01' && c <= '\b' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '\x1f' || c == '$' || c == '%%' || c >= '(' && c <= ',' || c == ';' || c >= '[' && c <= '^' || c == '{' || c >= '}' && c <= 'ÿ'
	case c == '!':
		goto yystate6
	case c == '"':
		goto yystate8
	case c == '#':
		goto yystate9
	case c == '&':
		goto yystate11
	case c == '-':
		goto yystate15
	case c == '.':
		goto yystate17
	case c == '/':
		goto yystate22
	case c == '0':
		goto yystate27
	case c == ':':
		goto yystate36
	case c == '<':
		goto yystate38
	case c == '=':
		goto yystate43
	case c == '>':
		goto yystate44
	case c == '?':
		goto yystate47
	case c == '@':
		goto yystate48
	case c == 'A' || c == 'a':
		goto yystate67
	case c == 'B' || c == 'b':
		goto yystate132
	case c == 'C' || c == 'c':
		goto yystate176
	case c == 'D' || c == 'd':
		goto yystate315
	case c == 'E' || c == 'e':
		goto yystate469
	case c == 'F' || c == 'f':
		goto yystate511
	case c == 'G' || c == 'g':
		goto yystate557
	case c == 'H' || c == 'h':
		goto yystate591
	case c == 'I' || c == 'i':
		goto yystate638
	case c == 'J' || c == 'j':
		goto yystate691
	case c == 'K' || c == 'k':
		goto yystate695
	case c == 'L' || c == 'l':
		goto yystate710
	case c == 'M' || c == 'm':
		goto yystate791
	case c == 'N' || c == 'n':
		goto yystate862
	case c == 'O' || c == 'o':
		goto yystate886
	case c == 'P' || c == 'p':
		goto yystate908
	case c == 'Q' || c == 'q':
		goto yystate951
	case c == 'R' || c == 'r':
		goto yystate961
	case c == 'S' || c == 's':
		goto yystate1051
	case c == 'T' || c == 't':
		goto yystate1201
	case c == 'U' || c == 'u':
		goto yystate1264
	case c == 'V' || c == 'v':
		goto yystate1321
	case c == 'W' || c == 'w':
		goto yystate1350
	case c == 'X' || c == 'x':
		goto yystate1379
	case c == 'Y' || c == 'y':
		goto yystate1385
	case c == 'Z' || c == 'z':
		goto yystate1399
	case c == '\'':
		goto yystate14
	case c == '\n':
		goto yystate5
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate4
	case c == '\x00':
		goto yystate2
	case c == '_':
		goto yystate1407
	case c == '`':
		goto yystate1408
	case c == '|':
		goto yystate1409
	case c >= '1' && c <= '9':
		goto yystate34
	}

yystate2:
	c = l.next()
	goto yyrule1

yystate3:
	c = l.next()
	goto yyrule352

yystate4:
	c = l.next()
	switch {
	default:
		goto yyrule2
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate5
	}

yystate5:
	c = l.next()
	switch {
	default:
		goto yyrule2
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate5
	}

yystate6:
	c = l.next()
	switch {
	default:
		goto yyrule352
	case c == '=':
		goto yystate7
	}

yystate7:
	c = l.next()
	goto yyrule34

yystate8:
	c = l.next()
	goto yyrule13

yystate9:
	c = l.next()
	switch {
	default:
		goto yyrule3
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate10
	}

yystate10:
	c = l.next()
	switch {
	default:
		goto yyrule3
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate10
	}

yystate11:
	c = l.next()
	switch {
	default:
		goto yyrule352
	case c == '&':
		goto yystate12
	case c == '^':
		goto yystate13
	}

yystate12:
	c = l.next()
	goto yyrule27

yystate13:
	c = l.next()
	goto yyrule28

yystate14:
	c = l.next()
	goto yyrule14

yystate15:
	c = l.next()
	switch {
	default:
		goto yyrule352
	case c == '-':
		goto yystate16
	}

yystate16:
	c = l.next()
	goto yyrule6

yystate17:
	c = l.next()
	switch {
	default:
		goto yyrule352
	case c >= '0' && c <= '9':
		goto yystate18
	}

yystate18:
	c = l.next()
	switch {
	default:
		goto yyrule10
	case c == 'E' || c == 'e':
		goto yystate19
	case c >= '0' && c <= '9':
		goto yystate18
	}

yystate19:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '+' || c == '-':
		goto yystate20
	case c >= '0' && c <= '9':
		goto yystate21
	}

yystate20:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate21
	}

yystate21:
	c = l.next()
	switch {
	default:
		goto yyrule10
	case c >= '0' && c <= '9':
		goto yystate21
	}

yystate22:
	c = l.next()
	switch {
	default:
		goto yyrule352
	case c == '*':
		goto yystate23
	case c == '/':
		goto yystate26
	}

yystate23:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate24
	case c >= '\x01' && c <= ')' || c >= '+' && c <= 'ÿ':
		goto yystate23
	}

yystate24:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate24
	case c == '/':
		goto yystate25
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '.' || c >= '0' && c <= 'ÿ':
		goto yystate23
	}

yystate25:
	c = l.next()
	goto yyrule5

yystate26:
	c = l.next()
	switch {
	default:
		goto yyrule4
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate26
	}

yystate27:
	c = l.next()
	switch {
	default:
		goto yyrule9
	case c == '.':
		goto yystate18
	case c == '8' || c == '9':
		goto yystate29
	case c == 'B' || c == 'b':
		goto yystate30
	case c == 'E' || c == 'e':
		goto yystate19
	case c == 'X' || c == 'x':
		goto yystate32
	case c >= '0' && c <= '7':
		goto yystate28
	}

yystate28:
	c = l.next()
	switch {
	default:
		goto yyrule9
	case c == '.':
		goto yystate18
	case c == '8' || c == '9':
		goto yystate29
	case c == 'E' || c == 'e':
		goto yystate19
	case c >= '0' && c <= '7':
		goto yystate28
	}

yystate29:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate18
	case c == 'E' || c == 'e':
		goto yystate19
	case c >= '0' && c <= '9':
		goto yystate29
	}

yystate30:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '0' || c == '1':
		goto yystate31
	}

yystate31:
	c = l.next()
	switch {
	default:
		goto yyrule12
	case c == '0' || c == '1':
		goto yystate31
	}

yystate32:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate33
	}

yystate33:
	c = l.next()
	switch {
	default:
		goto yyrule11
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate33
	}

yystate34:
	c = l.next()
	switch {
	default:
		goto yyrule9
	case c == '.':
		goto yystate18
	case c == 'E' || c == 'e':
		goto yystate19
	case c >= '0' && c <= '9':
		goto yystate35
	}

yystate35:
	c = l.next()
	switch {
	default:
		goto yyrule9
	case c == '.':
		goto yystate18
	case c == 'E' || c == 'e':
		goto yystate19
	case c >= '0' && c <= '9':
		goto yystate35
	}

yystate36:
	c = l.next()
	switch {
	default:
		goto yyrule352
	case c == '=':
		goto yystate37
	}

yystate37:
	c = l.next()
	goto yyrule32

yystate38:
	c = l.next()
	switch {
	default:
		goto yyrule352
	case c == '<':
		goto yystate39
	case c == '=':
		goto yystate40
	case c == '>':
		goto yystate42
	}

yystate39:
	c = l.next()
	goto yyrule29

yystate40:
	c = l.next()
	switch {
	default:
		goto yyrule30
	case c == '>':
		goto yystate41
	}

yystate41:
	c = l.next()
	goto yyrule38

yystate42:
	c = l.next()
	goto yyrule35

yystate43:
	c = l.next()
	goto yyrule31

yystate44:
	c = l.next()
	switch {
	default:
		goto yyrule352
	case c == '=':
		goto yystate45
	case c == '>':
		goto yystate46
	}

yystate45:
	c = l.next()
	goto yyrule33

yystate46:
	c = l.next()
	goto yyrule37

yystate47:
	c = l.next()
	goto yyrule40

yystate48:
	c = l.next()
	switch {
	default:
		goto yyrule39
	case c == '@':
		goto yystate49
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate66
	}

yystate49:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == 'G' || c == 'g':
		goto yystate51
	case c == 'L' || c == 'l':
		goto yystate58
	case c == 'S' || c == 's':
		goto yystate60
	case c >= 'A' && c <= 'F' || c >= 'H' && c <= 'K' || c >= 'M' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'k' || c >= 'm' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate50
	}

yystate50:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate50
	}

yystate51:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate50
	case c == 'L' || c == 'l':
		goto yystate52
	}

yystate52:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate50
	case c == 'O' || c == 'o':
		goto yystate53
	}

yystate53:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate50
	case c == 'B' || c == 'b':
		goto yystate54
	}

yystate54:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate50
	case c == 'A' || c == 'a':
		goto yystate55
	}

yystate55:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate50
	case c == 'L' || c == 'l':
		goto yystate56
	}

yystate56:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate50
	case c == '.':
		goto yystate57
	}

yystate57:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate50
	}

yystate58:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate50
	case c == 'O' || c == 'o':
		goto yystate59
	}

yystate59:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate50
	case c == 'C' || c == 'c':
		goto yystate54
	}

yystate60:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate50
	case c == 'E' || c == 'e':
		goto yystate61
	}

yystate61:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate50
	case c == 'S' || c == 's':
		goto yystate62
	}

yystate62:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate50
	case c == 'S' || c == 's':
		goto yystate63
	}

yystate63:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate50
	case c == 'I' || c == 'i':
		goto yystate64
	}

yystate64:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate50
	case c == 'O' || c == 'o':
		goto yystate65
	}

yystate65:
	c = l.next()
	switch {
	default:
		goto yyrule248
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate50
	case c == 'N' || c == 'n':
		goto yystate56
	}

yystate66:
	c = l.next()
	switch {
	default:
		goto yyrule249
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate66
	}

yystate67:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'E' || c >= 'G' && c <= 'K' || c == 'M' || c >= 'O' && c <= 'R' || c == 'T' || c >= 'W' && c <= 'Z' || c == '_' || c == 'a' || c == 'e' || c >= 'g' && c <= 'k' || c == 'm' || c >= 'o' && c <= 'r' || c == 't' || c >= 'w' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate69
	case c == 'C' || c == 'c':
		goto yystate71
	case c == 'D' || c == 'd':
		goto yystate76
	case c == 'F' || c == 'f':
		goto yystate85
	case c == 'L' || c == 'l':
		goto yystate89
	case c == 'N' || c == 'n':
		goto yystate94
	case c == 'S' || c == 's':
		goto yystate102
	case c == 'U' || c == 'u':
		goto yystate106
	case c == 'V' || c == 'v':
		goto yystate119
	}

yystate68:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate69:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate70
	}

yystate70:
	c = l.next()
	switch {
	default:
		goto yyrule41
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate71:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate72
	}

yystate72:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate73
	}

yystate73:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate74
	}

yystate74:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate75
	}

yystate75:
	c = l.next()
	switch {
	default:
		goto yyrule303
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate76:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate77
	case c == 'M' || c == 'm':
		goto yystate82
	}

yystate77:
	c = l.next()
	switch {
	default:
		goto yyrule42
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate78
	}

yystate78:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate79
	}

yystate79:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate80
	}

yystate80:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate81
	}

yystate81:
	c = l.next()
	switch {
	default:
		goto yyrule43
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate82:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate83
	}

yystate83:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate84
	}

yystate84:
	c = l.next()
	switch {
	default:
		goto yyrule44
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate85:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate86
	}

yystate86:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate87
	}

yystate87:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate88
	}

yystate88:
	c = l.next()
	switch {
	default:
		goto yyrule45
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate89:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate90
	case c == 'T' || c == 't':
		goto yystate91
	}

yystate90:
	c = l.next()
	switch {
	default:
		goto yyrule46
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate91:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate92
	}

yystate92:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate93
	}

yystate93:
	c = l.next()
	switch {
	default:
		goto yyrule47
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate94:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'B' || c == 'C' || c >= 'E' && c <= 'X' || c == 'Z' || c == '_' || c == 'b' || c == 'c' || c >= 'e' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate95
	case c == 'D' || c == 'd':
		goto yystate100
	case c == 'Y' || c == 'y':
		goto yystate101
	}

yystate95:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate96
	}

yystate96:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'Y' || c == 'y':
		goto yystate97
	}

yystate97:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Y' || c == '_' || c >= 'a' && c <= 'y':
		goto yystate68
	case c == 'Z' || c == 'z':
		goto yystate98
	}

yystate98:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate99
	}

yystate99:
	c = l.next()
	switch {
	default:
		goto yyrule48
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate100:
	c = l.next()
	switch {
	default:
		goto yyrule49
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate101:
	c = l.next()
	switch {
	default:
		goto yyrule50
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate102:
	c = l.next()
	switch {
	default:
		goto yyrule52
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate103
	}

yystate103:
	c = l.next()
	switch {
	default:
		goto yyrule51
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate104
	}

yystate104:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate105
	}

yystate105:
	c = l.next()
	switch {
	default:
		goto yyrule53
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate106:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate107
	}

yystate107:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate108
	}

yystate108:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate109
	}

yystate109:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate110
	}

yystate110:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate111
	}

yystate111:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate112
	}

yystate112:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate113
	}

yystate113:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate114
	}

yystate114:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate115
	}

yystate115:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate116
	}

yystate116:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate117
	}

yystate117:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate118
	}

yystate118:
	c = l.next()
	switch {
	default:
		goto yyrule54
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate119:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate120
	}

yystate120:
	c = l.next()
	switch {
	default:
		goto yyrule55
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate121
	}

yystate121:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate122
	}

yystate122:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate123
	}

yystate123:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'W' || c == 'w':
		goto yystate124
	}

yystate124:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate125
	}

yystate125:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate126
	}

yystate126:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate127
	}

yystate127:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate128
	}

yystate128:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate129
	}

yystate129:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate130
	}

yystate130:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate131
	}

yystate131:
	c = l.next()
	switch {
	default:
		goto yyrule56
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate132:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'H' || c == 'J' || c == 'K' || c == 'M' || c == 'N' || c >= 'P' && c <= 'S' || c >= 'U' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'h' || c == 'j' || c == 'k' || c == 'm' || c == 'n' || c >= 'p' && c <= 's' || c >= 'u' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate136
	case c == 'I' || c == 'i':
		goto yystate145
	case c == 'L' || c == 'l':
		goto yystate158
	case c == 'O' || c == 'o':
		goto yystate161
	case c == 'T' || c == 't':
		goto yystate169
	case c == 'Y' || c == 'y':
		goto yystate173
	case c == '\'':
		goto yystate133
	}

yystate133:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '0' || c == '1':
		goto yystate134
	}

yystate134:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '0' || c == '1':
		goto yystate134
	case c == '\'':
		goto yystate135
	}

yystate135:
	c = l.next()
	goto yyrule12

yystate136:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate137
	case c == 'T' || c == 't':
		goto yystate140
	}

yystate137:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate138
	}

yystate138:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate139
	}

yystate139:
	c = l.next()
	switch {
	default:
		goto yyrule57
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate140:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'W' || c == 'w':
		goto yystate141
	}

yystate141:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate142
	}

yystate142:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate143
	}

yystate143:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate144
	}

yystate144:
	c = l.next()
	switch {
	default:
		goto yyrule58
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate145:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'M' || c >= 'O' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'm' || c >= 'o' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate146
	case c == 'N' || c == 'n':
		goto yystate150
	case c == 'T' || c == 't':
		goto yystate157
	}

yystate146:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate147
	}

yystate147:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate148
	}

yystate148:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate149
	}

yystate149:
	c = l.next()
	switch {
	default:
		goto yyrule322
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate150:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate151
	case c == 'L' || c == 'l':
		goto yystate154
	}

yystate151:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate152
	}

yystate152:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'Y' || c == 'y':
		goto yystate153
	}

yystate153:
	c = l.next()
	switch {
	default:
		goto yyrule336
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate154:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate155
	}

yystate155:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate156
	}

yystate156:
	c = l.next()
	switch {
	default:
		goto yyrule59
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate157:
	c = l.next()
	switch {
	default:
		goto yyrule317
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate158:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate159
	}

yystate159:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate160
	}

yystate160:
	c = l.next()
	switch {
	default:
		goto yyrule339
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate161:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate162
	case c == 'T' || c == 't':
		goto yystate167
	}

yystate162:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate163
	}

yystate163:
	c = l.next()
	switch {
	default:
		goto yyrule346
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate164
	}

yystate164:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate165
	}

yystate165:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate166
	}

yystate166:
	c = l.next()
	switch {
	default:
		goto yyrule347
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate167:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate168
	}

yystate168:
	c = l.next()
	switch {
	default:
		goto yyrule60
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate169:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate170
	}

yystate170:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate171
	}

yystate171:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate172
	}

yystate172:
	c = l.next()
	switch {
	default:
		goto yyrule61
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate173:
	c = l.next()
	switch {
	default:
		goto yyrule62
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate174
	}

yystate174:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate175
	}

yystate175:
	c = l.next()
	switch {
	default:
		goto yyrule348
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate176:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'G' || c >= 'I' && c <= 'N' || c == 'P' || c == 'Q' || c == 'S' || c == 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'g' || c >= 'i' && c <= 'n' || c == 'p' || c == 'q' || c == 's' || c == 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate177
	case c == 'H' || c == 'h':
		goto yystate185
	case c == 'O' || c == 'o':
		goto yystate202
	case c == 'R' || c == 'r':
		goto yystate275
	case c == 'U' || c == 'u':
		goto yystate283
	}

yystate177:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate178
	}

yystate178:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c == 'D' || c >= 'F' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c == 'd' || c >= 'f' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate179
	case c == 'E' || c == 'e':
		goto yystate183
	case c == 'T' || c == 't':
		goto yystate184
	}

yystate179:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate180
	}

yystate180:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate181
	}

yystate181:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate182
	}

yystate182:
	c = l.next()
	switch {
	default:
		goto yyrule301
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate183:
	c = l.next()
	switch {
	default:
		goto yyrule63
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate184:
	c = l.next()
	switch {
	default:
		goto yyrule64
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate185:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate186
	case c == 'E' || c == 'e':
		goto yystate196
	}

yystate186:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate187
	}

yystate187:
	c = l.next()
	switch {
	default:
		goto yyrule334
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate188
	case c == 'S' || c == 's':
		goto yystate193
	}

yystate188:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate189
	}

yystate189:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate190
	}

yystate190:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate191
	}

yystate191:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate192
	}

yystate192:
	c = l.next()
	switch {
	default:
		goto yyrule65
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate193:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate194
	}

yystate194:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate195
	}

yystate195:
	c = l.next()
	switch {
	default:
		goto yyrule66
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate196:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate197
	}

yystate197:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate198
	}

yystate198:
	c = l.next()
	switch {
	default:
		goto yyrule67
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate199
	}

yystate199:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate200
	}

yystate200:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate201
	}

yystate201:
	c = l.next()
	switch {
	default:
		goto yyrule68
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate202:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'K' || c >= 'O' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'k' || c >= 'o' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate203
	case c == 'L' || c == 'l':
		goto yystate209
	case c == 'M' || c == 'm':
		goto yystate221
	case c == 'N' || c == 'n':
		goto yystate244
	case c == 'U' || c == 'u':
		goto yystate272
	}

yystate203:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate204
	}

yystate204:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate205
	}

yystate205:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate206
	}

yystate206:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate207
	}

yystate207:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate208
	}

yystate208:
	c = l.next()
	switch {
	default:
		goto yyrule69
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate209:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate210
	case c == 'U' || c == 'u':
		goto yystate217
	}

yystate210:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate211
	}

yystate211:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate212
	}

yystate212:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate213
	case c == 'I' || c == 'i':
		goto yystate214
	}

yystate213:
	c = l.next()
	switch {
	default:
		goto yyrule70
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate214:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate215
	}

yystate215:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate216
	}

yystate216:
	c = l.next()
	switch {
	default:
		goto yyrule71
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate217:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate218
	}

yystate218:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate219
	}

yystate219:
	c = l.next()
	switch {
	default:
		goto yyrule72
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate220
	}

yystate220:
	c = l.next()
	switch {
	default:
		goto yyrule73
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate221:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c == 'N' || c == 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c == 'n' || c == 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate222
	case c == 'P' || c == 'p':
		goto yystate231
	}

yystate222:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate223
	case c == 'I' || c == 'i':
		goto yystate226
	}

yystate223:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate224
	}

yystate224:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate225
	}

yystate225:
	c = l.next()
	switch {
	default:
		goto yyrule74
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate226:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate227
	}

yystate227:
	c = l.next()
	switch {
	default:
		goto yyrule75
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate228
	}

yystate228:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate229
	}

yystate229:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate230
	}

yystate230:
	c = l.next()
	switch {
	default:
		goto yyrule76
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate231:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate232
	case c == 'R' || c == 'r':
		goto yystate235
	}

yystate232:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate233
	}

yystate233:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate234
	}

yystate234:
	c = l.next()
	switch {
	default:
		goto yyrule77
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate235:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate236
	}

yystate236:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate237
	}

yystate237:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate238
	}

yystate238:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate239
	case c == 'I' || c == 'i':
		goto yystate241
	}

yystate239:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate240
	}

yystate240:
	c = l.next()
	switch {
	default:
		goto yyrule78
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate241:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate242
	}

yystate242:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate243
	}

yystate243:
	c = l.next()
	switch {
	default:
		goto yyrule79
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate244:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'M' || c >= 'O' && c <= 'R' || c == 'T' || c == 'U' || c >= 'W' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'm' || c >= 'o' && c <= 'r' || c == 't' || c == 'u' || c >= 'w' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate245
	case c == 'N' || c == 'n':
		goto yystate251
	case c == 'S' || c == 's':
		goto yystate261
	case c == 'V' || c == 'v':
		goto yystate268
	}

yystate245:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate246
	}

yystate246:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate247
	}

yystate247:
	c = l.next()
	switch {
	default:
		goto yyrule80
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate248
	}

yystate248:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'W' || c == 'w':
		goto yystate249
	}

yystate249:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate250
	}

yystate250:
	c = l.next()
	switch {
	default:
		goto yyrule81
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate251:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate252
	}

yystate252:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate253
	}

yystate253:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate254
	}

yystate254:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate255
	}

yystate255:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate256
	}

yystate256:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate257
	}

yystate257:
	c = l.next()
	switch {
	default:
		goto yyrule82
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate258
	}

yystate258:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate259
	}

yystate259:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate260
	}

yystate260:
	c = l.next()
	switch {
	default:
		goto yyrule83
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate261:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate262
	}

yystate262:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate263
	}

yystate263:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate264
	}

yystate264:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate265
	}

yystate265:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate266
	}

yystate266:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate267
	}

yystate267:
	c = l.next()
	switch {
	default:
		goto yyrule84
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate268:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate269
	}

yystate269:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate270
	}

yystate270:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate271
	}

yystate271:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate272:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate273
	}

yystate273:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate274
	}

yystate274:
	c = l.next()
	switch {
	default:
		goto yyrule86
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate275:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate276
	case c == 'O' || c == 'o':
		goto yystate280
	}

yystate276:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate277
	}

yystate277:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate278
	}

yystate278:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate279
	}

yystate279:
	c = l.next()
	switch {
	default:
		goto yyrule87
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate280:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate281
	}

yystate281:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate282
	}

yystate282:
	c = l.next()
	switch {
	default:
		goto yyrule88
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate283:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate284
	}

yystate284:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Q' || c == 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'q' || c == 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate285
	case c == 'R' || c == 'r':
		goto yystate289
	case c == 'T' || c == 't':
		goto yystate311
	}

yystate285:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate286
	}

yystate286:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate287
	}

yystate287:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate288
	}

yystate288:
	c = l.next()
	switch {
	default:
		goto yyrule89
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate289:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate290
	}

yystate290:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate291
	}

yystate291:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate292
	}

yystate292:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate293
	}

yystate293:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'S' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 's' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate294
	case c == 'T' || c == 't':
		goto yystate298
	case c == 'U' || c == 'u':
		goto yystate307
	}

yystate294:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate295
	}

yystate295:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate296
	}

yystate296:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate297
	}

yystate297:
	c = l.next()
	switch {
	default:
		goto yyrule90
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate298:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate299
	}

yystate299:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate300
	}

yystate300:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate301
	}

yystate301:
	c = l.next()
	switch {
	default:
		goto yyrule92
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate302
	}

yystate302:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate303
	}

yystate303:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate304
	}

yystate304:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate305
	}

yystate305:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'P' || c == 'p':
		goto yystate306
	}

yystate306:
	c = l.next()
	switch {
	default:
		goto yyrule313
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate307:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate308
	}

yystate308:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate309
	}

yystate309:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate310
	}

yystate310:
	c = l.next()
	switch {
	default:
		goto yyrule93
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate311:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate312
	}

yystate312:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate313
	}

yystate313:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate314
	}

yystate314:
	c = l.next()
	switch {
	default:
		goto yyrule91
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate315:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'B' || c == 'C' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'N' || c == 'P' || c == 'Q' || c == 'S' || c == 'T' || c >= 'V' && c <= 'X' || c == 'Z' || c == '_' || c == 'b' || c == 'c' || c >= 'f' && c <= 'h' || c >= 'j' && c <= 'n' || c == 'p' || c == 'q' || c == 's' || c == 't' || c >= 'v' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate316
	case c == 'D' || c == 'd':
		goto yystate388
	case c == 'E' || c == 'e':
		goto yystate390
	case c == 'I' || c == 'i':
		goto yystate433
	case c == 'O' || c == 'o':
		goto yystate445
	case c == 'R' || c == 'r':
		goto yystate450
	case c == 'U' || c == 'u':
		goto yystate453
	case c == 'Y' || c == 'y':
		goto yystate463
	}

yystate316:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate317
	case c == 'Y' || c == 'y':
		goto yystate342
	}

yystate317:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate318
	case c == 'E' || c == 'e':
		goto yystate324
	}

yystate318:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate319
	}

yystate319:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate320
	}

yystate320:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate321
	}

yystate321:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate322
	}

yystate322:
	c = l.next()
	switch {
	default:
		goto yyrule94
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate323
	}

yystate323:
	c = l.next()
	switch {
	default:
		goto yyrule95
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate324:
	c = l.next()
	switch {
	default:
		goto yyrule329
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate325
	case c == '_':
		goto yystate329
	}

yystate325:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate326
	}

yystate326:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate327
	}

yystate327:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate328
	}

yystate328:
	c = l.next()
	switch {
	default:
		goto yyrule332
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate329:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'E' || c >= 'G' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'e' || c >= 'g' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate330
	case c == 'F' || c == 'f':
		goto yystate333
	case c == 'S' || c == 's':
		goto yystate339
	}

yystate330:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate331
	}

yystate331:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate332
	}

yystate332:
	c = l.next()
	switch {
	default:
		goto yyrule96
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate333:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate334
	}

yystate334:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate335
	}

yystate335:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate336
	}

yystate336:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate337
	}

yystate337:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate338
	}

yystate338:
	c = l.next()
	switch {
	default:
		goto yyrule97
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate339:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate340
	}

yystate340:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate341
	}

yystate341:
	c = l.next()
	switch {
	default:
		goto yyrule98
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate342:
	c = l.next()
	switch {
	default:
		goto yyrule99
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'P' && c <= 'Z' || c >= 'a' && c <= 'm' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate343
	case c == 'O' || c == 'o':
		goto yystate347
	case c == '_':
		goto yystate362
	}

yystate343:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate344
	}

yystate344:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate345
	}

yystate345:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate346
	}

yystate346:
	c = l.next()
	switch {
	default:
		goto yyrule100
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate347:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate68
	case c == 'F' || c == 'f':
		goto yystate348
	}

yystate348:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'V' || c == 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'v' || c == 'x' || c == 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate349
	case c == 'W' || c == 'w':
		goto yystate354
	case c == 'Y' || c == 'y':
		goto yystate358
	}

yystate349:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate350
	}

yystate350:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate351
	}

yystate351:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate352
	}

yystate352:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate353
	}

yystate353:
	c = l.next()
	switch {
	default:
		goto yyrule102
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate354:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate355
	}

yystate355:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate356
	}

yystate356:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate357
	}

yystate357:
	c = l.next()
	switch {
	default:
		goto yyrule101
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate358:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate359
	}

yystate359:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate360
	}

yystate360:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate361
	}

yystate361:
	c = l.next()
	switch {
	default:
		goto yyrule103
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate362:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'L' || c >= 'N' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'l' || c >= 'n' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate363
	case c == 'M' || c == 'm':
		goto yystate367
	case c == 'S' || c == 's':
		goto yystate382
	}

yystate363:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate364
	}

yystate364:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate365
	}

yystate365:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate366
	}

yystate366:
	c = l.next()
	switch {
	default:
		goto yyrule104
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate367:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate368
	}

yystate368:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate369
	case c == 'N' || c == 'n':
		goto yystate378
	}

yystate369:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate370
	}

yystate370:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate371
	}

yystate371:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate372
	}

yystate372:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate373
	}

yystate373:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate374
	}

yystate374:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate375
	}

yystate375:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate376
	}

yystate376:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate377
	}

yystate377:
	c = l.next()
	switch {
	default:
		goto yyrule105
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate378:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate379
	}

yystate379:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate380
	}

yystate380:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate381
	}

yystate381:
	c = l.next()
	switch {
	default:
		goto yyrule106
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate382:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate383
	}

yystate383:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate384
	}

yystate384:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate385
	}

yystate385:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate386
	}

yystate386:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate387
	}

yystate387:
	c = l.next()
	switch {
	default:
		goto yyrule107
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate388:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate389
	}

yystate389:
	c = l.next()
	switch {
	default:
		goto yyrule108
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate390:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'B' || c == 'D' || c == 'E' || c >= 'G' && c <= 'K' || c >= 'M' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c == 'b' || c == 'd' || c == 'e' || c >= 'g' && c <= 'k' || c >= 'm' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate391
	case c == 'C' || c == 'c':
		goto yystate399
	case c == 'F' || c == 'f':
		goto yystate404
	case c == 'L' || c == 'l':
		goto yystate409
	case c == 'S' || c == 's':
		goto yystate427
	}

yystate391:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate392
	}

yystate392:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate393
	}

yystate393:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate394
	}

yystate394:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate395
	}

yystate395:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate396
	}

yystate396:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate397
	}

yystate397:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate398
	}

yystate398:
	c = l.next()
	switch {
	default:
		goto yyrule109
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate399:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate400
	}

yystate400:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate401
	}

yystate401:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate402
	}

yystate402:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate403
	}

yystate403:
	c = l.next()
	switch {
	default:
		goto yyrule323
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate404:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate405
	}

yystate405:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate406
	}

yystate406:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate407
	}

yystate407:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate408
	}

yystate408:
	c = l.next()
	switch {
	default:
		goto yyrule110
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate409:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate410
	case c == 'E' || c == 'e':
		goto yystate424
	}

yystate410:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'Y' || c == 'y':
		goto yystate411
	}

yystate411:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate412
	case c == '_':
		goto yystate414
	}

yystate412:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate413
	}

yystate413:
	c = l.next()
	switch {
	default:
		goto yyrule111
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate414:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate415
	}

yystate415:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate416
	}

yystate416:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'Y' || c == 'y':
		goto yystate417
	}

yystate417:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate418
	}

yystate418:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'W' || c == 'w':
		goto yystate419
	}

yystate419:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate420
	}

yystate420:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate421
	}

yystate421:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate422
	}

yystate422:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate423
	}

yystate423:
	c = l.next()
	switch {
	default:
		goto yyrule112
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate424:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate425
	}

yystate425:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate426
	}

yystate426:
	c = l.next()
	switch {
	default:
		goto yyrule113
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate427:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate428
	}

yystate428:
	c = l.next()
	switch {
	default:
		goto yyrule114
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate429
	}

yystate429:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate430
	}

yystate430:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate431
	}

yystate431:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate432
	}

yystate432:
	c = l.next()
	switch {
	default:
		goto yyrule115
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate433:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c == 'T' || c == 'U' || c >= 'W' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c == 't' || c == 'u' || c >= 'w' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate434
	case c == 'V' || c == 'v':
		goto yystate444
	}

yystate434:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'b' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate435
	case c == 'T' || c == 't':
		goto yystate439
	}

yystate435:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate436
	}

yystate436:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate437
	}

yystate437:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate438
	}

yystate438:
	c = l.next()
	switch {
	default:
		goto yyrule117
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate439:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate440
	}

yystate440:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate441
	}

yystate441:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate442
	}

yystate442:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate443
	}

yystate443:
	c = l.next()
	switch {
	default:
		goto yyrule118
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate444:
	c = l.next()
	switch {
	default:
		goto yyrule119
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate445:
	c = l.next()
	switch {
	default:
		goto yyrule120
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate446
	}

yystate446:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate447
	}

yystate447:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate448
	}

yystate448:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate449
	}

yystate449:
	c = l.next()
	switch {
	default:
		goto yyrule326
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate450:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate451
	}

yystate451:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'P' || c == 'p':
		goto yystate452
	}

yystate452:
	c = l.next()
	switch {
	default:
		goto yyrule116
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate453:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate454
	case c == 'P' || c == 'p':
		goto yystate456
	}

yystate454:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate455
	}

yystate455:
	c = l.next()
	switch {
	default:
		goto yyrule121
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate456:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate457
	}

yystate457:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate458
	}

yystate458:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate459
	}

yystate459:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate460
	}

yystate460:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate461
	}

yystate461:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate462
	}

yystate462:
	c = l.next()
	switch {
	default:
		goto yyrule122
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate463:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate464
	}

yystate464:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate465
	}

yystate465:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate466
	}

yystate466:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate467
	}

yystate467:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate468
	}

yystate468:
	c = l.next()
	switch {
	default:
		goto yyrule123
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate469:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c == 'M' || c >= 'O' && c <= 'R' || c >= 'T' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'k' || c == 'm' || c >= 'o' && c <= 'r' || c >= 't' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate470
	case c == 'N' || c == 'n':
		goto yystate473
	case c == 'S' || c == 's':
		goto yystate486
	case c == 'X' || c == 'x':
		goto yystate491
	}

yystate470:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate471
	}

yystate471:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate472
	}

yystate472:
	c = l.next()
	switch {
	default:
		goto yyrule124
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate473:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'B' || c == 'C' || c == 'E' || c == 'F' || c >= 'H' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c == 'b' || c == 'c' || c == 'e' || c == 'f' || c >= 'h' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate474
	case c == 'D' || c == 'd':
		goto yystate478
	case c == 'G' || c == 'g':
		goto yystate479
	case c == 'U' || c == 'u':
		goto yystate484
	}

yystate474:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate475
	}

yystate475:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate476
	}

yystate476:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate477
	}

yystate477:
	c = l.next()
	switch {
	default:
		goto yyrule125
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate478:
	c = l.next()
	switch {
	default:
		goto yyrule126
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate479:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate480
	}

yystate480:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate481
	}

yystate481:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate482
	}

yystate482:
	c = l.next()
	switch {
	default:
		goto yyrule127
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate483
	}

yystate483:
	c = l.next()
	switch {
	default:
		goto yyrule128
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate484:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate485
	}

yystate485:
	c = l.next()
	switch {
	default:
		goto yyrule130
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate486:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate487
	}

yystate487:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate488
	}

yystate488:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'P' || c == 'p':
		goto yystate489
	}

yystate489:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate490
	}

yystate490:
	c = l.next()
	switch {
	default:
		goto yyrule131
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate491:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'O' || c >= 'Q' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'h' || c >= 'j' && c <= 'o' || c >= 'q' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate492
	case c == 'I' || c == 'i':
		goto yystate497
	case c == 'P' || c == 'p':
		goto yystate501
	case c == 'T' || c == 't':
		goto yystate506
	}

yystate492:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate493
	}

yystate493:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate494
	}

yystate494:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate495
	}

yystate495:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate496
	}

yystate496:
	c = l.next()
	switch {
	default:
		goto yyrule129
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate497:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate498
	}

yystate498:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate499
	}

yystate499:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate500
	}

yystate500:
	c = l.next()
	switch {
	default:
		goto yyrule132
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate501:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate502
	}

yystate502:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate503
	}

yystate503:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate504
	}

yystate504:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate505
	}

yystate505:
	c = l.next()
	switch {
	default:
		goto yyrule133
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate506:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate507
	}

yystate507:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate508
	}

yystate508:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate509
	}

yystate509:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate510
	}

yystate510:
	c = l.next()
	switch {
	default:
		goto yyrule134
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate511:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'H' || c == 'J' || c == 'K' || c == 'M' || c == 'N' || c == 'P' || c == 'Q' || c == 'S' || c == 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'h' || c == 'j' || c == 'k' || c == 'm' || c == 'n' || c == 'p' || c == 'q' || c == 's' || c == 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate512
	case c == 'I' || c == 'i':
		goto yystate516
	case c == 'L' || c == 'l':
		goto yystate527
	case c == 'O' || c == 'o':
		goto yystate531
	case c == 'R' || c == 'r':
		goto yystate547
	case c == 'U' || c == 'u':
		goto yystate550
	}

yystate512:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate513
	}

yystate513:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate514
	}

yystate514:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate515
	}

yystate515:
	c = l.next()
	switch {
	default:
		goto yyrule308
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate516:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Q' || c >= 'S' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'q' || c >= 's' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate517
	case c == 'R' || c == 'r':
		goto yystate521
	case c == 'X' || c == 'x':
		goto yystate524
	}

yystate517:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate518
	}

yystate518:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate519
	}

yystate519:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate520
	}

yystate520:
	c = l.next()
	switch {
	default:
		goto yyrule135
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate521:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate522
	}

yystate522:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate523
	}

yystate523:
	c = l.next()
	switch {
	default:
		goto yyrule136
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate524:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate525
	}

yystate525:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate526
	}

yystate526:
	c = l.next()
	switch {
	default:
		goto yyrule137
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate527:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate528
	}

yystate528:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate529
	}

yystate529:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate530
	}

yystate530:
	c = l.next()
	switch {
	default:
		goto yyrule325
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate531:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c == 'S' || c == 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c == 's' || c == 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate532
	case c == 'U' || c == 'u':
		goto yystate539
	}

yystate532:
	c = l.next()
	switch {
	default:
		goto yyrule138
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c == 'D' || c >= 'F' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c == 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate533
	case c == 'E' || c == 'e':
		goto yystate535
	}

yystate533:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate534
	}

yystate534:
	c = l.next()
	switch {
	default:
		goto yyrule139
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate535:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate536
	}

yystate536:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate537
	}

yystate537:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate538
	}

yystate538:
	c = l.next()
	switch {
	default:
		goto yyrule140
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate539:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate540
	}

yystate540:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate541
	}

yystate541:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate542
	}

yystate542:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate543
	}

yystate543:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate544
	}

yystate544:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'W' || c == 'w':
		goto yystate545
	}

yystate545:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate546
	}

yystate546:
	c = l.next()
	switch {
	default:
		goto yyrule141
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate547:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate548
	}

yystate548:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate549
	}

yystate549:
	c = l.next()
	switch {
	default:
		goto yyrule142
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate550:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate551
	}

yystate551:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate552
	}

yystate552:
	c = l.next()
	switch {
	default:
		goto yyrule143
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate553
	}

yystate553:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate554
	}

yystate554:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate68
	case c == 'X' || c == 'x':
		goto yystate555
	}

yystate555:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate556
	}

yystate556:
	c = l.next()
	switch {
	default:
		goto yyrule144
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate557:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate558
	case c == 'L' || c == 'l':
		goto yystate565
	case c == 'R' || c == 'r':
		goto yystate570
	}

yystate558:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate559
	}

yystate559:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate560
	}

yystate560:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate561
	}

yystate561:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate562
	}

yystate562:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate563
	}

yystate563:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate564
	}

yystate564:
	c = l.next()
	switch {
	default:
		goto yyrule235
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate565:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate566
	}

yystate566:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate567
	}

yystate567:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate568
	}

yystate568:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate569
	}

yystate569:
	c = l.next()
	switch {
	default:
		goto yyrule236
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate570:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'D' || c >= 'F' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c >= 'f' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate571
	case c == 'E' || c == 'e':
		goto yystate575
	case c == 'O' || c == 'o':
		goto yystate581
	}

yystate571:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate572
	}

yystate572:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate573
	}

yystate573:
	c = l.next()
	switch {
	default:
		goto yyrule145
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate574
	}

yystate574:
	c = l.next()
	switch {
	default:
		goto yyrule146
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate575:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate576
	}

yystate576:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate577
	}

yystate577:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate578
	}

yystate578:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate579
	}

yystate579:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate580
	}

yystate580:
	c = l.next()
	switch {
	default:
		goto yyrule147
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate581:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate582
	}

yystate582:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'P' || c == 'p':
		goto yystate583
	}

yystate583:
	c = l.next()
	switch {
	default:
		goto yyrule148
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate584
	}

yystate584:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate585
	}

yystate585:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate586
	}

yystate586:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate587
	}

yystate587:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate588
	}

yystate588:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate589
	}

yystate589:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate590
	}

yystate590:
	c = l.next()
	switch {
	default:
		goto yyrule149
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate591:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'D' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c >= 'f' && c <= 'h' || c >= 'j' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate592
	case c == 'E' || c == 'e':
		goto yystate599
	case c == 'I' || c == 'i':
		goto yystate601
	case c == 'O' || c == 'o':
		goto yystate613
	}

yystate592:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c == 'T' || c == 'U' || c >= 'W' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c == 't' || c == 'u' || c >= 'w' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate593
	case c == 'V' || c == 'v':
		goto yystate595
	}

yystate593:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate594
	}

yystate594:
	c = l.next()
	switch {
	default:
		goto yyrule150
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate595:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate596
	}

yystate596:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate597
	}

yystate597:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate598
	}

yystate598:
	c = l.next()
	switch {
	default:
		goto yyrule151
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate599:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate68
	case c == 'X' || c == 'x':
		goto yystate600
	}

yystate600:
	c = l.next()
	switch {
	default:
		goto yyrule152
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate601:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate602
	}

yystate602:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate603
	}

yystate603:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate604
	}

yystate604:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'P' || c == 'p':
		goto yystate605
	}

yystate605:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate606
	}

yystate606:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate607
	}

yystate607:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate608
	}

yystate608:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate609
	}

yystate609:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate610
	}

yystate610:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate611
	}

yystate611:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'Y' || c == 'y':
		goto yystate612
	}

yystate612:
	c = l.next()
	switch {
	default:
		goto yyrule153
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate613:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate614
	}

yystate614:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate615
	}

yystate615:
	c = l.next()
	switch {
	default:
		goto yyrule154
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate616
	}

yystate616:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate617
	case c == 'S' || c == 's':
		goto yystate632
	}

yystate617:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate618
	}

yystate618:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate619
	case c == 'N' || c == 'n':
		goto yystate628
	}

yystate619:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate620
	}

yystate620:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate621
	}

yystate621:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate622
	}

yystate622:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate623
	}

yystate623:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate624
	}

yystate624:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate625
	}

yystate625:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate626
	}

yystate626:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate627
	}

yystate627:
	c = l.next()
	switch {
	default:
		goto yyrule155
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate628:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate629
	}

yystate629:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate630
	}

yystate630:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate631
	}

yystate631:
	c = l.next()
	switch {
	default:
		goto yyrule156
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate632:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate633
	}

yystate633:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate634
	}

yystate634:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate635
	}

yystate635:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate636
	}

yystate636:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate637
	}

yystate637:
	c = l.next()
	switch {
	default:
		goto yyrule157
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate638:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c == 'E' || c >= 'H' && c <= 'M' || c >= 'O' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c == 'e' || c >= 'h' && c <= 'm' || c >= 'o' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate639
	case c == 'F' || c == 'f':
		goto yystate648
	case c == 'G' || c == 'g':
		goto yystate653
	case c == 'N' || c == 'n':
		goto yystate658
	case c == 'S' || c == 's':
		goto yystate679
	}

yystate639:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate640
	}

yystate640:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate641
	}

yystate641:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate642
	}

yystate642:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate643
	}

yystate643:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate68
	case c == 'F' || c == 'f':
		goto yystate644
	}

yystate644:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate645
	}

yystate645:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate646
	}

yystate646:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate647
	}

yystate647:
	c = l.next()
	switch {
	default:
		goto yyrule158
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate648:
	c = l.next()
	switch {
	default:
		goto yyrule159
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate649
	}

yystate649:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate650
	}

yystate650:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate651
	}

yystate651:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate652
	}

yystate652:
	c = l.next()
	switch {
	default:
		goto yyrule160
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate653:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate654
	}

yystate654:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate655
	}

yystate655:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate656
	}

yystate656:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate657
	}

yystate657:
	c = l.next()
	switch {
	default:
		goto yyrule162
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate658:
	c = l.next()
	switch {
	default:
		goto yyrule168
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'M' || c >= 'O' && c <= 'R' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'm' || c >= 'o' && c <= 'r' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate659
	case c == 'N' || c == 'n':
		goto yystate662
	case c == 'S' || c == 's':
		goto yystate665
	case c == 'T' || c == 't':
		goto yystate669
	}

yystate659:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate660
	}

yystate660:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate68
	case c == 'X' || c == 'x':
		goto yystate661
	}

yystate661:
	c = l.next()
	switch {
	default:
		goto yyrule163
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate662:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate663
	}

yystate663:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate664
	}

yystate664:
	c = l.next()
	switch {
	default:
		goto yyrule164
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate665:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate666
	}

yystate666:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate667
	}

yystate667:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate668
	}

yystate668:
	c = l.next()
	switch {
	default:
		goto yyrule165
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate669:
	c = l.next()
	switch {
	default:
		goto yyrule349
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate670
	case c == 'O' || c == 'o':
		goto yystate678
	}

yystate670:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate671
	case c == 'R' || c == 'r':
		goto yystate674
	}

yystate671:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate672
	}

yystate672:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate673
	}

yystate673:
	c = l.next()
	switch {
	default:
		goto yyrule350
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate674:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'U' || c >= 'W' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'u' || c >= 'w' && c <= 'z':
		goto yystate68
	case c == 'V' || c == 'v':
		goto yystate675
	}

yystate675:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate676
	}

yystate676:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate677
	}

yystate677:
	c = l.next()
	switch {
	default:
		goto yyrule166
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate678:
	c = l.next()
	switch {
	default:
		goto yyrule167
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate679:
	c = l.next()
	switch {
	default:
		goto yyrule169
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate680
	case c == 'O' || c == 'o':
		goto yystate684
	}

yystate680:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate681
	}

yystate681:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate682
	}

yystate682:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate683
	}

yystate683:
	c = l.next()
	switch {
	default:
		goto yyrule161
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate684:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate685
	}

yystate685:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate686
	}

yystate686:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate687
	}

yystate687:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate688
	}

yystate688:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate689
	}

yystate689:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate690
	}

yystate690:
	c = l.next()
	switch {
	default:
		goto yyrule170
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate691:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate692
	}

yystate692:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate693
	}

yystate693:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate694
	}

yystate694:
	c = l.next()
	switch {
	default:
		goto yyrule171
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate695:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate696
	}

yystate696:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'Y' || c == 'y':
		goto yystate697
	}

yystate697:
	c = l.next()
	switch {
	default:
		goto yyrule172
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate698
	case c == '_':
		goto yystate699
	}

yystate698:
	c = l.next()
	switch {
	default:
		goto yyrule174
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate699:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate700
	}

yystate700:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate701
	}

yystate701:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate702
	}

yystate702:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate703
	}

yystate703:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate704
	}

yystate704:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate705
	}

yystate705:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate706
	}

yystate706:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate707
	}

yystate707:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Y' || c == '_' || c >= 'a' && c <= 'y':
		goto yystate68
	case c == 'Z' || c == 'z':
		goto yystate708
	}

yystate708:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate709
	}

yystate709:
	c = l.next()
	switch {
	default:
		goto yyrule173
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate710:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'B' || c == 'D' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'N' || c >= 'P' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c == 'b' || c == 'd' || c >= 'f' && c <= 'h' || c >= 'j' && c <= 'n' || c >= 'p' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate711
	case c == 'C' || c == 'c':
		goto yystate724
	case c == 'E' || c == 'e':
		goto yystate728
	case c == 'I' || c == 'i':
		goto yystate743
	case c == 'O' || c == 'o':
		goto yystate749
	case c == 'T' || c == 't':
		goto yystate787
	}

yystate711:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate712
	}

yystate712:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate713
	}

yystate713:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate714
	}

yystate714:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate715
	}

yystate715:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate716
	}

yystate716:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate717
	}

yystate717:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate718
	}

yystate718:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate719
	}

yystate719:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate720
	}

yystate720:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate721
	}

yystate721:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate722
	}

yystate722:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate723
	}

yystate723:
	c = l.next()
	switch {
	default:
		goto yyrule175
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate724:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate725
	}

yystate725:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate726
	}

yystate726:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate727
	}

yystate727:
	c = l.next()
	switch {
	default:
		goto yyrule186
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate728:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'E' || c >= 'G' && c <= 'M' || c >= 'O' && c <= 'U' || c >= 'W' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'e' || c >= 'g' && c <= 'm' || c >= 'o' && c <= 'u' || c >= 'w' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate729
	case c == 'F' || c == 'f':
		goto yystate734
	case c == 'N' || c == 'n':
		goto yystate736
	case c == 'V' || c == 'v':
		goto yystate740
	}

yystate729:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate730
	}

yystate730:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate731
	}

yystate731:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate732
	}

yystate732:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate733
	}

yystate733:
	c = l.next()
	switch {
	default:
		goto yyrule176
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate734:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate735
	}

yystate735:
	c = l.next()
	switch {
	default:
		goto yyrule177
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate736:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate737
	}

yystate737:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate738
	}

yystate738:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate739
	}

yystate739:
	c = l.next()
	switch {
	default:
		goto yyrule178
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate740:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate741
	}

yystate741:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate742
	}

yystate742:
	c = l.next()
	switch {
	default:
		goto yyrule179
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate743:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c == 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c == 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate744
	case c == 'M' || c == 'm':
		goto yystate746
	}

yystate744:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate745
	}

yystate745:
	c = l.next()
	switch {
	default:
		goto yyrule180
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate746:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate747
	}

yystate747:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate748
	}

yystate748:
	c = l.next()
	switch {
	default:
		goto yyrule181
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate749:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'M' || c >= 'O' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'm' || c >= 'o' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate750
	case c == 'N' || c == 'n':
		goto yystate765
	case c == 'W' || c == 'w':
		goto yystate775
	}

yystate750:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate751
	case c == 'K' || c == 'k':
		goto yystate764
	}

yystate751:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate752
	case c == 'T' || c == 't':
		goto yystate762
	}

yystate752:
	c = l.next()
	switch {
	default:
		goto yyrule182
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate753
	}

yystate753:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate754
	}

yystate754:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate755
	}

yystate755:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate756
	}

yystate756:
	c = l.next()
	switch {
	default:
		goto yyrule314
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate757
	}

yystate757:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate758
	}

yystate758:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate759
	}

yystate759:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate760
	}

yystate760:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'P' || c == 'p':
		goto yystate761
	}

yystate761:
	c = l.next()
	switch {
	default:
		goto yyrule315
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate762:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate763
	}

yystate763:
	c = l.next()
	switch {
	default:
		goto yyrule183
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate764:
	c = l.next()
	switch {
	default:
		goto yyrule184
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate765:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate766
	}

yystate766:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate767
	case c == 'T' || c == 't':
		goto yystate771
	}

yystate767:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate768
	}

yystate768:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate769
	}

yystate769:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate770
	}

yystate770:
	c = l.next()
	switch {
	default:
		goto yyrule341
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate771:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate772
	}

yystate772:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate68
	case c == 'X' || c == 'x':
		goto yystate773
	}

yystate773:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate774
	}

yystate774:
	c = l.next()
	switch {
	default:
		goto yyrule345
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate775:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate776
	case c == '_':
		goto yystate778
	}

yystate776:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate777
	}

yystate777:
	c = l.next()
	switch {
	default:
		goto yyrule185
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate778:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'P' || c == 'p':
		goto yystate779
	}

yystate779:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate780
	}

yystate780:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate781
	}

yystate781:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate782
	}

yystate782:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate783
	}

yystate783:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate784
	}

yystate784:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate785
	}

yystate785:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'Y' || c == 'y':
		goto yystate786
	}

yystate786:
	c = l.next()
	switch {
	default:
		goto yyrule187
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate787:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate788
	}

yystate788:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate789
	}

yystate789:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate790
	}

yystate790:
	c = l.next()
	switch {
	default:
		goto yyrule188
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate791:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'D' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c >= 'f' && c <= 'h' || c >= 'j' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate792
	case c == 'E' || c == 'e':
		goto yystate799
	case c == 'I' || c == 'i':
		goto yystate815
	case c == 'O' || c == 'o':
		goto yystate852
	}

yystate792:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate68
	case c == 'X' || c == 'x':
		goto yystate793
	}

yystate793:
	c = l.next()
	switch {
	default:
		goto yyrule189
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate794
	}

yystate794:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate795
	}

yystate795:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate796
	}

yystate796:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'W' || c == 'w':
		goto yystate797
	}

yystate797:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate798
	}

yystate798:
	c = l.next()
	switch {
	default:
		goto yyrule190
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate799:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate800
	}

yystate800:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate801
	}

yystate801:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate802
	}

yystate802:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate803
	}

yystate803:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'H' || c >= 'J' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'h' || c >= 'j' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate804
	case c == 'I' || c == 'i':
		goto yystate808
	case c == 'T' || c == 't':
		goto yystate811
	}

yystate804:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate805
	}

yystate805:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate806
	}

yystate806:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate807
	}

yystate807:
	c = l.next()
	switch {
	default:
		goto yyrule340
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate808:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate809
	}

yystate809:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate810
	}

yystate810:
	c = l.next()
	switch {
	default:
		goto yyrule321
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate811:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate812
	}

yystate812:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate68
	case c == 'X' || c == 'x':
		goto yystate813
	}

yystate813:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate814
	}

yystate814:
	c = l.next()
	switch {
	default:
		goto yyrule343
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate815:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate816
	case c == 'N' || c == 'n':
		goto yystate825
	}

yystate816:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate817
	}

yystate817:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate818
	}

yystate818:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate819
	}

yystate819:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate820
	}

yystate820:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate821
	}

yystate821:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate822
	}

yystate822:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate823
	}

yystate823:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate824
	}

yystate824:
	c = l.next()
	switch {
	default:
		goto yyrule191
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate825:
	c = l.next()
	switch {
	default:
		goto yyrule192
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate826
	case c == '_':
		goto yystate847
	}

yystate826:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate827
	}

yystate827:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate828
	}

yystate828:
	c = l.next()
	switch {
	default:
		goto yyrule193
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate829
	}

yystate829:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate830
	case c == 'S' || c == 's':
		goto yystate841
	}

yystate830:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate831
	}

yystate831:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate832
	}

yystate832:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate833
	}

yystate833:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate834
	}

yystate834:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate835
	}

yystate835:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate836
	}

yystate836:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate837
	}

yystate837:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate838
	}

yystate838:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate839
	}

yystate839:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate840
	}

yystate840:
	c = l.next()
	switch {
	default:
		goto yyrule194
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate841:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate842
	}

yystate842:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate843
	}

yystate843:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate844
	}

yystate844:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate845
	}

yystate845:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate846
	}

yystate846:
	c = l.next()
	switch {
	default:
		goto yyrule195
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate847:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate848
	}

yystate848:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate849
	}

yystate849:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'W' || c == 'w':
		goto yystate850
	}

yystate850:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate851
	}

yystate851:
	c = l.next()
	switch {
	default:
		goto yyrule196
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate852:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate853
	case c == 'N' || c == 'n':
		goto yystate855
	}

yystate853:
	c = l.next()
	switch {
	default:
		goto yyrule197
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate854
	}

yystate854:
	c = l.next()
	switch {
	default:
		goto yyrule198
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate855:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate856
	}

yystate856:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate857
	}

yystate857:
	c = l.next()
	switch {
	default:
		goto yyrule199
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate858
	}

yystate858:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate859
	}

yystate859:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate860
	}

yystate860:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate861
	}

yystate861:
	c = l.next()
	switch {
	default:
		goto yyrule200
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate862:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'N' || c >= 'P' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'n' || c >= 'p' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate863
	case c == 'O' || c == 'o':
		goto yystate873
	case c == 'U' || c == 'u':
		goto yystate876
	}

yystate863:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate864
	case c == 'T' || c == 't':
		goto yystate867
	}

yystate864:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate865
	}

yystate865:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate866
	}

yystate866:
	c = l.next()
	switch {
	default:
		goto yyrule201
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate867:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate868
	}

yystate868:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate869
	}

yystate869:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate870
	}

yystate870:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate871
	}

yystate871:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate872
	}

yystate872:
	c = l.next()
	switch {
	default:
		goto yyrule202
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate873:
	c = l.next()
	switch {
	default:
		goto yyrule302
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c == 'U' || c == 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c == 'u' || c == 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate874
	case c == 'W' || c == 'w':
		goto yystate875
	}

yystate874:
	c = l.next()
	switch {
	default:
		goto yyrule203
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate875:
	c = l.next()
	switch {
	default:
		goto yyrule316
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate876:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate877
	case c == 'M' || c == 'm':
		goto yystate881
	}

yystate877:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate878
	}

yystate878:
	c = l.next()
	switch {
	default:
		goto yyrule307
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate879
	}

yystate879:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate68
	case c == 'F' || c == 'f':
		goto yystate880
	}

yystate880:
	c = l.next()
	switch {
	default:
		goto yyrule278
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate881:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate882
	}

yystate882:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate883
	}

yystate883:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate884
	}

yystate884:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate885
	}

yystate885:
	c = l.next()
	switch {
	default:
		goto yyrule324
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate886:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'M' || c == 'O' || c == 'Q' || c == 'S' || c == 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'm' || c == 'o' || c == 'q' || c == 's' || c == 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'F' || c == 'f':
		goto yystate887
	case c == 'N' || c == 'n':
		goto yystate892
	case c == 'P' || c == 'p':
		goto yystate895
	case c == 'R' || c == 'r':
		goto yystate900
	case c == 'U' || c == 'u':
		goto yystate904
	}

yystate887:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate68
	case c == 'F' || c == 'f':
		goto yystate888
	}

yystate888:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate889
	}

yystate889:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate890
	}

yystate890:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate891
	}

yystate891:
	c = l.next()
	switch {
	default:
		goto yyrule204
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate892:
	c = l.next()
	switch {
	default:
		goto yyrule205
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate893
	}

yystate893:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'Y' || c == 'y':
		goto yystate894
	}

yystate894:
	c = l.next()
	switch {
	default:
		goto yyrule206
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate895:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate896
	}

yystate896:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate897
	}

yystate897:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate898
	}

yystate898:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate899
	}

yystate899:
	c = l.next()
	switch {
	default:
		goto yyrule207
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate900:
	c = l.next()
	switch {
	default:
		goto yyrule209
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate901
	}

yystate901:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate902
	}

yystate902:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate903
	}

yystate903:
	c = l.next()
	switch {
	default:
		goto yyrule208
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate904:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate905
	}

yystate905:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate906
	}

yystate906:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate907
	}

yystate907:
	c = l.next()
	switch {
	default:
		goto yyrule210
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate908:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'N' || c == 'P' || c == 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'n' || c == 'p' || c == 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate909
	case c == 'O' || c == 'o':
		goto yystate916
	case c == 'R' || c == 'r':
		goto yystate920
	}

yystate909:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate910
	}

yystate910:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate911
	}

yystate911:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'W' || c == 'w':
		goto yystate912
	}

yystate912:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate913
	}

yystate913:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate914
	}

yystate914:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate915
	}

yystate915:
	c = l.next()
	switch {
	default:
		goto yyrule211
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate916:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'W' || c == 'w':
		goto yystate917
	}

yystate917:
	c = l.next()
	switch {
	default:
		goto yyrule212
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate918
	}

yystate918:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate919
	}

yystate919:
	c = l.next()
	switch {
	default:
		goto yyrule213
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate920:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'h' || c >= 'j' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate921
	case c == 'I' || c == 'i':
		goto yystate932
	case c == 'O' || c == 'o':
		goto yystate944
	}

yystate921:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate922
	case c == 'P' || c == 'p':
		goto yystate928
	}

yystate922:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate923
	}

yystate923:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate924
	}

yystate924:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate925
	}

yystate925:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate926
	}

yystate926:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate927
	}

yystate927:
	c = l.next()
	switch {
	default:
		goto yyrule327
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate928:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate929
	}

yystate929:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate930
	}

yystate930:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate931
	}

yystate931:
	c = l.next()
	switch {
	default:
		goto yyrule214
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate932:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'U' || c >= 'W' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'u' || c >= 'w' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate933
	case c == 'V' || c == 'v':
		goto yystate937
	}

yystate933:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate934
	}

yystate934:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate935
	}

yystate935:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'Y' || c == 'y':
		goto yystate936
	}

yystate936:
	c = l.next()
	switch {
	default:
		goto yyrule215
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate937:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate938
	}

yystate938:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate939
	}

yystate939:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate940
	}

yystate940:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate941
	}

yystate941:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate942
	}

yystate942:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate943
	}

yystate943:
	c = l.next()
	switch {
	default:
		goto yyrule216
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate944:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate945
	}

yystate945:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate946
	}

yystate946:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate947
	}

yystate947:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate948
	}

yystate948:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate949
	}

yystate949:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate950
	}

yystate950:
	c = l.next()
	switch {
	default:
		goto yyrule217
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate951:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate952
	}

yystate952:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate953
	case c == 'I' || c == 'i':
		goto yystate958
	}

yystate953:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate954
	}

yystate954:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate955
	}

yystate955:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate956
	}

yystate956:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate957
	}

yystate957:
	c = l.next()
	switch {
	default:
		goto yyrule218
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate958:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate959
	}

yystate959:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate960
	}

yystate960:
	c = l.next()
	switch {
	default:
		goto yyrule219
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate961:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'D' || c >= 'F' && c <= 'H' || c == 'J' || c == 'K' || c == 'M' || c == 'N' || c >= 'P' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c >= 'f' && c <= 'h' || c == 'j' || c == 'k' || c == 'm' || c == 'n' || c >= 'p' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate962
	case c == 'E' || c == 'e':
		goto yystate965
	case c == 'I' || c == 'i':
		goto yystate1021
	case c == 'L' || c == 'l':
		goto yystate1025
	case c == 'O' || c == 'o':
		goto yystate1029
	case c == 'T' || c == 't':
		goto yystate1047
	}

yystate962:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate963
	}

yystate963:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate964
	}

yystate964:
	c = l.next()
	switch {
	default:
		goto yyrule237
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate965:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'B' || c == 'C' || c == 'E' || c >= 'H' && c <= 'K' || c >= 'M' && c <= 'O' || c == 'Q' || c == 'R' || c == 'T' || c == 'U' || c >= 'W' && c <= 'Z' || c == '_' || c == 'b' || c == 'c' || c == 'e' || c >= 'h' && c <= 'k' || c >= 'm' && c <= 'o' || c == 'q' || c == 'r' || c == 't' || c == 'u' || c >= 'w' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate966
	case c == 'D' || c == 'd':
		goto yystate969
	case c == 'F' || c == 'f':
		goto yystate976
	case c == 'G' || c == 'g':
		goto yystate984
	case c == 'L' || c == 'l':
		goto yystate988
	case c == 'P' || c == 'p':
		goto yystate998
	case c == 'S' || c == 's':
		goto yystate1010
	case c == 'V' || c == 'v':
		goto yystate1016
	}

yystate966:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate967
	case c == 'L' || c == 'l':
		goto yystate968
	}

yystate967:
	c = l.next()
	switch {
	default:
		goto yyrule238
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate968:
	c = l.next()
	switch {
	default:
		goto yyrule328
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate969:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate970
	}

yystate970:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate971
	}

yystate971:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate972
	}

yystate972:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate973
	}

yystate973:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate974
	}

yystate974:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate975
	}

yystate975:
	c = l.next()
	switch {
	default:
		goto yyrule220
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate976:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate977
	}

yystate977:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate978
	}

yystate978:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate979
	}

yystate979:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate980
	}

yystate980:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate981
	}

yystate981:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate982
	}

yystate982:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate983
	}

yystate983:
	c = l.next()
	switch {
	default:
		goto yyrule244
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate984:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate985
	}

yystate985:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate68
	case c == 'X' || c == 'x':
		goto yystate986
	}

yystate986:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'P' || c == 'p':
		goto yystate987
	}

yystate987:
	c = l.next()
	switch {
	default:
		goto yyrule242
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate988:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate989
	}

yystate989:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate990
	}

yystate990:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate991
	}

yystate991:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate992
	}

yystate992:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate993
	}

yystate993:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate994
	}

yystate994:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate995
	}

yystate995:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate996
	}

yystate996:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate997
	}

yystate997:
	c = l.next()
	switch {
	default:
		goto yyrule239
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate998:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate999
	case c == 'L' || c == 'l':
		goto yystate1006
	}

yystate999:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1000
	}

yystate1000:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1001
	}

yystate1001:
	c = l.next()
	switch {
	default:
		goto yyrule240
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1002
	}

yystate1002:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate1003
	}

yystate1003:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1004
	}

yystate1004:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1005
	}

yystate1005:
	c = l.next()
	switch {
	default:
		goto yyrule241
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1006:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1007
	}

yystate1007:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1008
	}

yystate1008:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1009
	}

yystate1009:
	c = l.next()
	switch {
	default:
		goto yyrule243
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1010:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1011
	}

yystate1011:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1012
	}

yystate1012:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1013
	}

yystate1013:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1014
	}

yystate1014:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1015
	}

yystate1015:
	c = l.next()
	switch {
	default:
		goto yyrule300
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1016:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1017
	}

yystate1017:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1018
	}

yystate1018:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1019
	}

yystate1019:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1020
	}

yystate1020:
	c = l.next()
	switch {
	default:
		goto yyrule247
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1021:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate1022
	}

yystate1022:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate1023
	}

yystate1023:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1024
	}

yystate1024:
	c = l.next()
	switch {
	default:
		goto yyrule221
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1025:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1026
	}

yystate1026:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate1027
	}

yystate1027:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1028
	}

yystate1028:
	c = l.next()
	switch {
	default:
		goto yyrule245
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1029:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'T' || c == 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 't' || c == 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1030
	case c == 'U' || c == 'u':
		goto yystate1036
	case c == 'W' || c == 'w':
		goto yystate1039
	}

yystate1030:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1031
	}

yystate1031:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate1032
	}

yystate1032:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1033
	}

yystate1033:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1034
	}

yystate1034:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate1035
	}

yystate1035:
	c = l.next()
	switch {
	default:
		goto yyrule222
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1036:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1037
	}

yystate1037:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate1038
	}

yystate1038:
	c = l.next()
	switch {
	default:
		goto yyrule223
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1039:
	c = l.next()
	switch {
	default:
		goto yyrule224
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate1040
	}

yystate1040:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate68
	case c == 'F' || c == 'f':
		goto yystate1041
	}

yystate1041:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1042
	}

yystate1042:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1043
	}

yystate1043:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate1044
	}

yystate1044:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1045
	}

yystate1045:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1046
	}

yystate1046:
	c = l.next()
	switch {
	default:
		goto yyrule225
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1047:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1048
	}

yystate1048:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1049
	}

yystate1049:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate1050
	}

yystate1050:
	c = l.next()
	switch {
	default:
		goto yyrule246
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1051:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c == 'D' || c == 'F' || c == 'G' || c == 'J' || c == 'K' || c == 'N' || c == 'R' || c == 'S' || c >= 'V' && c <= 'X' || c == 'Z' || c == '_' || c == 'a' || c == 'b' || c == 'd' || c == 'f' || c == 'g' || c == 'j' || c == 'k' || c == 'n' || c == 'r' || c == 's' || c >= 'v' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1052
	case c == 'E' || c == 'e':
		goto yystate1058
	case c == 'H' || c == 'h':
		goto yystate1095
	case c == 'I' || c == 'i':
		goto yystate1101
	case c == 'L' || c == 'l':
		goto yystate1106
	case c == 'M' || c == 'm':
		goto yystate1110
	case c == 'O' || c == 'o':
		goto yystate1117
	case c == 'P' || c == 'p':
		goto yystate1120
	case c == 'Q' || c == 'q':
		goto yystate1124
	case c == 'T' || c == 't':
		goto yystate1153
	case c == 'U' || c == 'u':
		goto yystate1176
	case c == 'Y' || c == 'y':
		goto yystate1195
	}

yystate1052:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate1053
	}

yystate1053:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1054
	}

yystate1054:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate1055
	}

yystate1055:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1056
	}

yystate1056:
	c = l.next()
	switch {
	default:
		goto yyrule226
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1057
	}

yystate1057:
	c = l.next()
	switch {
	default:
		goto yyrule227
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1058:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'K' || c >= 'M' && c <= 'Q' || c >= 'U' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'k' || c >= 'm' && c <= 'q' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1059
	case c == 'L' || c == 'l':
		goto yystate1075
	case c == 'R' || c == 'r':
		goto yystate1079
	case c == 'S' || c == 's':
		goto yystate1089
	case c == 'T' || c == 't':
		goto yystate1094
	}

yystate1059:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1060
	}

yystate1060:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1061
	}

yystate1061:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate1062
	}

yystate1062:
	c = l.next()
	switch {
	default:
		goto yyrule251
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate1063
	}

yystate1063:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate1064
	}

yystate1064:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1065
	}

yystate1065:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1066
	}

yystate1066:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1067
	}

yystate1067:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1068
	}

yystate1068:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1069
	}

yystate1069:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1070
	}

yystate1070:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1071
	}

yystate1071:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1072
	}

yystate1072:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1073
	}

yystate1073:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate1074
	}

yystate1074:
	c = l.next()
	switch {
	default:
		goto yyrule252
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1075:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1076
	}

yystate1076:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1077
	}

yystate1077:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1078
	}

yystate1078:
	c = l.next()
	switch {
	default:
		goto yyrule253
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1079:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1080
	}

yystate1080:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1081
	}

yystate1081:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1082
	}

yystate1082:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1083
	}

yystate1083:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Y' || c == '_' || c >= 'a' && c <= 'y':
		goto yystate68
	case c == 'Z' || c == 'z':
		goto yystate1084
	}

yystate1084:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1085
	}

yystate1085:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate1086
	}

yystate1086:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1087
	}

yystate1087:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1088
	}

yystate1088:
	c = l.next()
	switch {
	default:
		goto yyrule228
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1089:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1090
	}

yystate1090:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1091
	}

yystate1091:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1092
	}

yystate1092:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1093
	}

yystate1093:
	c = l.next()
	switch {
	default:
		goto yyrule229
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1094:
	c = l.next()
	switch {
	default:
		goto yyrule254
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1095:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1096
	case c == 'O' || c == 'o':
		goto yystate1099
	}

yystate1096:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1097
	}

yystate1097:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1098
	}

yystate1098:
	c = l.next()
	switch {
	default:
		goto yyrule255
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1099:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'W' || c == 'w':
		goto yystate1100
	}

yystate1100:
	c = l.next()
	switch {
	default:
		goto yyrule256
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1101:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate1102
	}

yystate1102:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1103
	}

yystate1103:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1104
	}

yystate1104:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate1105
	}

yystate1105:
	c = l.next()
	switch {
	default:
		goto yyrule304
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1106:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1107
	}

yystate1107:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1108
	}

yystate1108:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'P' || c == 'p':
		goto yystate1109
	}

yystate1109:
	c = l.next()
	switch {
	default:
		goto yyrule257
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1110:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1111
	}

yystate1111:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1112
	}

yystate1112:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1113
	}

yystate1113:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1114
	}

yystate1114:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1115
	}

yystate1115:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1116
	}

yystate1116:
	c = l.next()
	switch {
	default:
		goto yyrule320
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1117:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate1118
	}

yystate1118:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1119
	}

yystate1119:
	c = l.next()
	switch {
	default:
		goto yyrule230
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1120:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1121
	}

yystate1121:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1122
	}

yystate1122:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1123
	}

yystate1123:
	c = l.next()
	switch {
	default:
		goto yyrule231
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1124:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1125
	}

yystate1125:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate1126
	}

yystate1126:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1127
	case c == 'N' || c == 'n':
		goto yystate1145
	}

yystate1127:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1128
	}

yystate1128:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1129
	case c == 'L' || c == 'l':
		goto yystate1132
	}

yystate1129:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate1130
	}

yystate1130:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1131
	}

yystate1131:
	c = l.next()
	switch {
	default:
		goto yyrule311
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1132:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1133
	}

yystate1133:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate1134
	}

yystate1134:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate68
	case c == 'F' || c == 'f':
		goto yystate1135
	}

yystate1135:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1136
	}

yystate1136:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate1137
	}

yystate1137:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1138
	}

yystate1138:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate1139
	}

yystate1139:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate1140
	}

yystate1140:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1141
	}

yystate1141:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1142
	}

yystate1142:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'W' || c == 'w':
		goto yystate1143
	}

yystate1143:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1144
	}

yystate1144:
	c = l.next()
	switch {
	default:
		goto yyrule310
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1145:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1146
	}

yystate1146:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate1147
	}

yystate1147:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1148
	}

yystate1148:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1149
	}

yystate1149:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1150
	}

yystate1150:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate1151
	}

yystate1151:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1152
	}

yystate1152:
	c = l.next()
	switch {
	default:
		goto yyrule312
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1153:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1154
	case c == 'R' || c == 'r':
		goto yystate1172
	}

yystate1154:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c == 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c == 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1155
	case c == 'T' || c == 't':
		goto yystate1157
	}

yystate1155:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1156
	}

yystate1156:
	c = l.next()
	switch {
	default:
		goto yyrule232
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1157:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c == 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c == 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1158
	case c == 'U' || c == 'u':
		goto yystate1170
	}

yystate1158:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate1159
	}

yystate1159:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'P' || c == 'p':
		goto yystate1160
	}

yystate1160:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1161
	}

yystate1161:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1162
	}

yystate1162:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1163
	}

yystate1163:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1164
	}

yystate1164:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1165
	}

yystate1165:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1166
	}

yystate1166:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1167
	}

yystate1167:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1168
	}

yystate1168:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1169
	}

yystate1169:
	c = l.next()
	switch {
	default:
		goto yyrule233
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1170:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1171
	}

yystate1171:
	c = l.next()
	switch {
	default:
		goto yyrule234
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1172:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1173
	}

yystate1173:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate1174
	}

yystate1174:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'P' || c == 'p':
		goto yystate1175
	}

yystate1175:
	c = l.next()
	switch {
	default:
		goto yyrule259
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1176:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate1177
	case c == 'M' || c == 'm':
		goto yystate1194
	}

yystate1177:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate1178
	case c == 'S' || c == 's':
		goto yystate1182
	}

yystate1178:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1179
	}

yystate1179:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1180
	}

yystate1180:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1181
	}

yystate1181:
	c = l.next()
	switch {
	default:
		goto yyrule258
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1182:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1183
	}

yystate1183:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1184
	}

yystate1184:
	c = l.next()
	switch {
	default:
		goto yyrule260
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1185
	}

yystate1185:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1186
	}

yystate1186:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate1187
	}

yystate1187:
	c = l.next()
	switch {
	default:
		goto yyrule261
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate1188
	}

yystate1188:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1189
	}

yystate1189:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1190
	}

yystate1190:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate1191
	}

yystate1191:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1192
	}

yystate1192:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate68
	case c == 'X' || c == 'x':
		goto yystate1193
	}

yystate1193:
	c = l.next()
	switch {
	default:
		goto yyrule262
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1194:
	c = l.next()
	switch {
	default:
		goto yyrule263
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1195:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1196
	}

yystate1196:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate1197
	}

yystate1197:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1198
	}

yystate1198:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1199
	}

yystate1199:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1200
	}

yystate1200:
	c = l.next()
	switch {
	default:
		goto yyrule264
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1201:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'D' || c == 'F' || c == 'G' || c >= 'J' && c <= 'N' || c == 'P' || c == 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c == 'f' || c == 'g' || c >= 'j' && c <= 'n' || c == 'p' || c == 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1202
	case c == 'E' || c == 'e':
		goto yystate1207
	case c == 'H' || c == 'h':
		goto yystate1210
	case c == 'I' || c == 'i':
		goto yystate1213
	case c == 'O' || c == 'o':
		goto yystate1234
	case c == 'R' || c == 'r':
		goto yystate1235
	}

yystate1202:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate1203
	}

yystate1203:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1204
	}

yystate1204:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1205
	}

yystate1205:
	c = l.next()
	switch {
	default:
		goto yyrule265
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1206
	}

yystate1206:
	c = l.next()
	switch {
	default:
		goto yyrule266
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1207:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate68
	case c == 'X' || c == 'x':
		goto yystate1208
	}

yystate1208:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1209
	}

yystate1209:
	c = l.next()
	switch {
	default:
		goto yyrule344
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1210:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1211
	}

yystate1211:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1212
	}

yystate1212:
	c = l.next()
	switch {
	default:
		goto yyrule267
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1213:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate1214
	case c == 'N' || c == 'n':
		goto yystate1221
	}

yystate1214:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1215
	}

yystate1215:
	c = l.next()
	switch {
	default:
		goto yyrule330
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1216
	}

yystate1216:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1217
	}

yystate1217:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1218
	}

yystate1218:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate1219
	}

yystate1219:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'P' || c == 'p':
		goto yystate1220
	}

yystate1220:
	c = l.next()
	switch {
	default:
		goto yyrule331
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1221:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'Y' || c == 'y':
		goto yystate1222
	}

yystate1222:
	c = l.next()
	switch {
	default:
		goto yyrule318
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'H' || c >= 'J' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'h' || c >= 'j' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate1223
	case c == 'I' || c == 'i':
		goto yystate1227
	case c == 'T' || c == 't':
		goto yystate1230
	}

yystate1223:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1224
	}

yystate1224:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1225
	}

yystate1225:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate1226
	}

yystate1226:
	c = l.next()
	switch {
	default:
		goto yyrule338
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1227:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1228
	}

yystate1228:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1229
	}

yystate1229:
	c = l.next()
	switch {
	default:
		goto yyrule319
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1230:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1231
	}

yystate1231:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'W' || c == 'Y' || c == 'Z' || c == '_' || c >= 'a' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate68
	case c == 'X' || c == 'x':
		goto yystate1232
	}

yystate1232:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1233
	}

yystate1233:
	c = l.next()
	switch {
	default:
		goto yyrule342
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1234:
	c = l.next()
	switch {
	default:
		goto yyrule268
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1235:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'H' || c >= 'J' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'h' || c >= 'j' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1236
	case c == 'I' || c == 'i':
		goto yystate1250
	case c == 'U' || c == 'u':
		goto yystate1257
	}

yystate1236:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1237
	case c == 'N' || c == 'n':
		goto yystate1242
	}

yystate1237:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1238
	}

yystate1238:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1239
	}

yystate1239:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1240
	}

yystate1240:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate1241
	}

yystate1241:
	c = l.next()
	switch {
	default:
		goto yyrule269
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1242:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1243
	}

yystate1243:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1244
	}

yystate1244:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1245
	}

yystate1245:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1246
	}

yystate1246:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1247
	}

yystate1247:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1248
	}

yystate1248:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1249
	}

yystate1249:
	c = l.next()
	switch {
	default:
		goto yyrule270
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1250:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate1251
	case c == 'M' || c == 'm':
		goto yystate1256
	}

yystate1251:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate1252
	}

yystate1252:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1253
	}

yystate1253:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1254
	}

yystate1254:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1255
	}

yystate1255:
	c = l.next()
	switch {
	default:
		goto yyrule271
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1256:
	c = l.next()
	switch {
	default:
		goto yyrule272
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1257:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1258
	case c == 'N' || c == 'n':
		goto yystate1259
	}

yystate1258:
	c = l.next()
	switch {
	default:
		goto yyrule309
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1259:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1260
	}

yystate1260:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1261
	}

yystate1261:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1262
	}

yystate1262:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1263
	}

yystate1263:
	c = l.next()
	switch {
	default:
		goto yyrule273
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1264:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'M' || c == 'O' || c == 'Q' || c == 'R' || c >= 'U' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'm' || c == 'o' || c == 'q' || c == 'r' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1265
	case c == 'N' || c == 'n':
		goto yystate1269
	case c == 'P' || c == 'p':
		goto yystate1300
	case c == 'S' || c == 's':
		goto yystate1308
	case c == 'T' || c == 't':
		goto yystate1314
	}

yystate1265:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1266
	}

yystate1266:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1267
	}

yystate1267:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1268
	}

yystate1268:
	c = l.next()
	switch {
	default:
		goto yyrule282
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1269:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'H' || c == 'J' || c >= 'M' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'h' || c == 'j' || c >= 'm' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1270
	case c == 'I' || c == 'i':
		goto yystate1279
	case c == 'K' || c == 'k':
		goto yystate1285
	case c == 'L' || c == 'l':
		goto yystate1290
	case c == 'S' || c == 's':
		goto yystate1294
	}

yystate1270:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1271
	}

yystate1271:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate1272
	}

yystate1272:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate1273
	}

yystate1273:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1274
	}

yystate1274:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1275
	}

yystate1275:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1276
	}

yystate1276:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1277
	}

yystate1277:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate1278
	}

yystate1278:
	c = l.next()
	switch {
	default:
		goto yyrule274
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1279:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c == 'P' || c >= 'R' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c == 'p' || c >= 'r' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1280
	case c == 'Q' || c == 'q':
		goto yystate1282
	}

yystate1280:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1281
	}

yystate1281:
	c = l.next()
	switch {
	default:
		goto yyrule275
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1282:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate1283
	}

yystate1283:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1284
	}

yystate1284:
	c = l.next()
	switch {
	default:
		goto yyrule276
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1285:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1286
	}

yystate1286:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1287
	}

yystate1287:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'W' || c == 'w':
		goto yystate1288
	}

yystate1288:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1289
	}

yystate1289:
	c = l.next()
	switch {
	default:
		goto yyrule277
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1290:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1291
	}

yystate1291:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1292
	}

yystate1292:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate1293
	}

yystate1293:
	c = l.next()
	switch {
	default:
		goto yyrule279
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1294:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1295
	}

yystate1295:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate1296
	}

yystate1296:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1297
	}

yystate1297:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1298
	}

yystate1298:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate1299
	}

yystate1299:
	c = l.next()
	switch {
	default:
		goto yyrule305
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1300:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'O' || c >= 'Q' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate1301
	case c == 'P' || c == 'p':
		goto yystate1305
	}

yystate1301:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1302
	}

yystate1302:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1303
	}

yystate1303:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1304
	}

yystate1304:
	c = l.next()
	switch {
	default:
		goto yyrule280
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1305:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1306
	}

yystate1306:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1307
	}

yystate1307:
	c = l.next()
	switch {
	default:
		goto yyrule281
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1308:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1309
	case c == 'I' || c == 'i':
		goto yystate1311
	}

yystate1309:
	c = l.next()
	switch {
	default:
		goto yyrule283
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1310
	}

yystate1310:
	c = l.next()
	switch {
	default:
		goto yyrule284
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1311:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1312
	}

yystate1312:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate1313
	}

yystate1313:
	c = l.next()
	switch {
	default:
		goto yyrule285
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1314:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate68
	case c == 'C' || c == 'c':
		goto yystate1315
	}

yystate1315:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate68
	case c == '_':
		goto yystate1316
	}

yystate1316:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate1317
	}

yystate1317:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1318
	}

yystate1318:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1319
	}

yystate1319:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1320
	}

yystate1320:
	c = l.next()
	switch {
	default:
		goto yyrule250
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1321:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1322
	case c == 'E' || c == 'e':
		goto yystate1344
	}

yystate1322:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1323
	case c == 'R' || c == 'r':
		goto yystate1327
	}

yystate1323:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate68
	case c == 'U' || c == 'u':
		goto yystate1324
	}

yystate1324:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1325
	}

yystate1325:
	c = l.next()
	switch {
	default:
		goto yyrule286
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1326
	}

yystate1326:
	c = l.next()
	switch {
	default:
		goto yyrule287
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1327:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'D' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c == 'a' || c >= 'd' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate1328
	case c == 'C' || c == 'c':
		goto yystate1334
	case c == 'I' || c == 'i':
		goto yystate1338
	}

yystate1328:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1329
	}

yystate1329:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1330
	}

yystate1330:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1331
	}

yystate1331:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1332
	}

yystate1332:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'Y' || c == 'y':
		goto yystate1333
	}

yystate1333:
	c = l.next()
	switch {
	default:
		goto yyrule337
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1334:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate1335
	}

yystate1335:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1336
	}

yystate1336:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1337
	}

yystate1337:
	c = l.next()
	switch {
	default:
		goto yyrule335
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1338:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1339
	}

yystate1339:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c == 'A' || c >= 'C' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate68
	case c == 'B' || c == 'b':
		goto yystate1340
	}

yystate1340:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1341
	}

yystate1341:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1342
	}

yystate1342:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1343
	}

yystate1343:
	c = l.next()
	switch {
	default:
		goto yyrule288
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1344:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1345
	}

yystate1345:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1346
	}

yystate1346:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1347
	}

yystate1347:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1348
	}

yystate1348:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1349
	}

yystate1349:
	c = l.next()
	switch {
	default:
		goto yyrule289
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1350:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'D' || c == 'F' || c == 'G' || c >= 'I' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c == 'f' || c == 'g' || c >= 'i' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1351
	case c == 'E' || c == 'e':
		goto yystate1358
	case c == 'H' || c == 'h':
		goto yystate1370
	case c == 'R' || c == 'r':
		goto yystate1375
	}

yystate1351:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1352
	}

yystate1352:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1353
	}

yystate1353:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1354
	}

yystate1354:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1355
	}

yystate1355:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate68
	case c == 'G' || c == 'g':
		goto yystate1356
	}

yystate1356:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate68
	case c == 'S' || c == 's':
		goto yystate1357
	}

yystate1357:
	c = l.next()
	switch {
	default:
		goto yyrule290
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1358:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1359
	}

yystate1359:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate1360
	}

yystate1360:
	c = l.next()
	switch {
	default:
		goto yyrule291
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'D' || c == 'd':
		goto yystate1361
	case c == 'O' || c == 'o':
		goto yystate1364
	}

yystate1361:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1362
	}

yystate1362:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'Y' || c == 'y':
		goto yystate1363
	}

yystate1363:
	c = l.next()
	switch {
	default:
		goto yyrule292
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1364:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate68
	case c == 'F' || c == 'f':
		goto yystate1365
	}

yystate1365:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'X' || c == 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate68
	case c == 'Y' || c == 'y':
		goto yystate1366
	}

yystate1366:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1367
	}

yystate1367:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1368
	}

yystate1368:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1369
	}

yystate1369:
	c = l.next()
	switch {
	default:
		goto yyrule293
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1370:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1371
	}

yystate1371:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1372
	case c == 'R' || c == 'r':
		goto yystate1373
	}

yystate1372:
	c = l.next()
	switch {
	default:
		goto yyrule294
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1373:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1374
	}

yystate1374:
	c = l.next()
	switch {
	default:
		goto yyrule295
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1375:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1376
	}

yystate1376:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1377
	}

yystate1377:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1378
	}

yystate1378:
	c = l.next()
	switch {
	default:
		goto yyrule296
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1379:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1383
	case c == '\'':
		goto yystate1380
	}

yystate1380:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate1381
	}

yystate1381:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate1382
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate1381
	}

yystate1382:
	c = l.next()
	goto yyrule11

yystate1383:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1384
	}

yystate1384:
	c = l.next()
	switch {
	default:
		goto yyrule297
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1385:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1386
	}

yystate1386:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate68
	case c == 'A' || c == 'a':
		goto yystate1387
	}

yystate1387:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1388
	}

yystate1388:
	c = l.next()
	switch {
	default:
		goto yyrule333
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'V' || c >= 'X' && c <= 'Z' || c >= 'a' && c <= 'v' || c >= 'x' && c <= 'z':
		goto yystate68
	case c == 'W' || c == 'w':
		goto yystate1389
	case c == '_':
		goto yystate1393
	}

yystate1389:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1390
	}

yystate1390:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1391
	}

yystate1391:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'J' || c >= 'L' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate68
	case c == 'K' || c == 'k':
		goto yystate1392
	}

yystate1392:
	c = l.next()
	switch {
	default:
		goto yyrule298
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1393:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'L' || c >= 'N' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate68
	case c == 'M' || c == 'm':
		goto yystate1394
	}

yystate1394:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1395
	}

yystate1395:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate68
	case c == 'N' || c == 'n':
		goto yystate1396
	}

yystate1396:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate68
	case c == 'T' || c == 't':
		goto yystate1397
	}

yystate1397:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate68
	case c == 'H' || c == 'h':
		goto yystate1398
	}

yystate1398:
	c = l.next()
	switch {
	default:
		goto yyrule299
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1399:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate68
	case c == 'E' || c == 'e':
		goto yystate1400
	}

yystate1400:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate68
	case c == 'R' || c == 'r':
		goto yystate1401
	}

yystate1401:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate68
	case c == 'O' || c == 'o':
		goto yystate1402
	}

yystate1402:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate68
	case c == 'F' || c == 'f':
		goto yystate1403
	}

yystate1403:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate68
	case c == 'I' || c == 'i':
		goto yystate1404
	}

yystate1404:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1405
	}

yystate1405:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate68
	case c == 'L' || c == 'l':
		goto yystate1406
	}

yystate1406:
	c = l.next()
	switch {
	default:
		goto yyrule306
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1407:
	c = l.next()
	switch {
	default:
		goto yyrule351
	case c == '$' || c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate68
	}

yystate1408:
	c = l.next()
	goto yyrule15

yystate1409:
	c = l.next()
	switch {
	default:
		goto yyrule352
	case c == '|':
		goto yystate1410
	}

yystate1410:
	c = l.next()
	goto yyrule36

	goto yystate1411 // silence unused label error
yystate1411:
	c = l.next()
yystart1411:
	switch {
	default:
		goto yyrule16
	case c == '"':
		goto yystate1413
	case c == '\\':
		goto yystate1415
	case c == '\x00':
		goto yystate2
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate1412
	}

yystate1412:
	c = l.next()
	switch {
	default:
		goto yyrule16
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate1412
	}

yystate1413:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c == '"':
		goto yystate1414
	}

yystate1414:
	c = l.next()
	goto yyrule18

yystate1415:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate1416
	}

yystate1416:
	c = l.next()
	goto yyrule17

	goto yystate1417 // silence unused label error
yystate1417:
	c = l.next()
yystart1417:
	switch {
	default:
		goto yyrule20
	case c == '\'':
		goto yystate1419
	case c == '\\':
		goto yystate1421
	case c == '\x00':
		goto yystate2
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate1418
	}

yystate1418:
	c = l.next()
	switch {
	default:
		goto yyrule20
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate1418
	}

yystate1419:
	c = l.next()
	switch {
	default:
		goto yyrule23
	case c == '\'':
		goto yystate1420
	}

yystate1420:
	c = l.next()
	goto yyrule22

yystate1421:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate1422
	}

yystate1422:
	c = l.next()
	goto yyrule21

	goto yystate1423 // silence unused label error
yystate1423:
	c = l.next()
yystart1423:
	switch {
	default:
		goto yystate1424 // c >= '\x01' && c <= '\b' || c >= '\n' && c <= '\x1f' || c >= '!' && c <= 'ÿ'
	case c == '\t' || c == ' ':
		goto yystate1425
	case c == '\x00':
		goto yystate2
	}

yystate1424:
	c = l.next()
	goto yyrule8

yystate1425:
	c = l.next()
	switch {
	default:
		goto yyrule7
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate1425
	}

	goto yystate1426 // silence unused label error
yystate1426:
	c = l.next()
yystart1426:
	switch {
	default:
		goto yyrule24
	case c == '\x00':
		goto yystate2
	case c == '`':
		goto yystate1428
	case c >= '\x01' && c <= '_' || c >= 'a' && c <= 'ÿ':
		goto yystate1427
	}

yystate1427:
	c = l.next()
	switch {
	default:
		goto yyrule24
	case c >= '\x01' && c <= '_' || c >= 'a' && c <= 'ÿ':
		goto yystate1427
	}

yystate1428:
	c = l.next()
	switch {
	default:
		goto yyrule26
	case c == '`':
		goto yystate1429
	}

yystate1429:
	c = l.next()
	goto yyrule25

yyrule1: // \0
	{
		return 0
	}
yyrule2: // [ \t\n\r]+

	goto yystate0
yyrule3: // #.*

	goto yystate0
yyrule4: // \/\/.*

	goto yystate0
yyrule5: // \/\*([^*]|\*+[^*/])*\*+\/

	goto yystate0
yyrule6: // --
	{
		l.sc = S3
		goto yystate0
	}
yyrule7: // [ \t]+.*
	{
		{
			l.sc = 0
		}
		goto yystate0
	}
yyrule8: // [^ \t]
	{
		{
			l.sc = 0
			l.c = '-'
			n := len(l.val)
			l.unget(l.val[n-1])
			return '-'
		}
		goto yystate0
	}
yyrule9: // {int_lit}
	{
		return toInt(l, lval, string(l.val))
	}
yyrule10: // {float_lit}
	{
		return toFloat(l, lval, string(l.val))
	}
yyrule11: // {hex_lit}
	{
		return toHex(l, lval, string(l.val))
	}
yyrule12: // {bit_lit}
	{
		return toBit(l, lval, string(l.val))
	}
yyrule13: // \"
	{
		l.sc = S1
		goto yystate0
	}
yyrule14: // '
	{
		l.sc = S2
		goto yystate0
	}
yyrule15: // `
	{
		l.sc = S4
		goto yystate0
	}
yyrule16: // [^\"\\]*
	{
		l.stringLit = append(l.stringLit, l.val...)
		goto yystate0
	}
yyrule17: // \\.
	{
		l.stringLit = append(l.stringLit, l.val...)
		goto yystate0
	}
yyrule18: // \"\"
	{
		l.stringLit = append(l.stringLit, '"')
		goto yystate0
	}
yyrule19: // \"
	{
		l.stringLit = append(l.stringLit, '"')
		l.sc = 0
		return l.str(lval, "\"")
	}
yyrule20: // [^'\\]*
	{
		l.stringLit = append(l.stringLit, l.val...)
		goto yystate0
	}
yyrule21: // \\.
	{
		l.stringLit = append(l.stringLit, l.val...)
		goto yystate0
	}
yyrule22: // ''
	{
		l.stringLit = append(l.stringLit, '\'')
		goto yystate0
	}
yyrule23: // '
	{
		l.stringLit = append(l.stringLit, '\'')
		l.sc = 0
		return l.str(lval, "'")
	}
yyrule24: // [^`]*
	{
		l.stringLit = append(l.stringLit, l.val...)
		goto yystate0
	}
yyrule25: // ``
	{
		l.stringLit = append(l.stringLit, '`')
		goto yystate0
	}
yyrule26: // `
	{
		l.sc = 0
		lval.ident = string(l.stringLit)
		l.stringLit = l.stringLit[0:0]
		return identifier
	}
yyrule27: // "&&"
	{
		return andand
	}
yyrule28: // "&^"
	{
		return andnot
	}
yyrule29: // "<<"
	{
		return lsh
	}
yyrule30: // "<="
	{
		return le
	}
yyrule31: // "="
	{
		return eq
	}
yyrule32: // ":="
	{
		return assignmentEq
	}
yyrule33: // ">="
	{
		return ge
	}
yyrule34: // "!="
	{
		return neq
	}
yyrule35: // "<>"
	{
		return neqSynonym
	}
yyrule36: // "||"
	{
		return oror
	}
yyrule37: // ">>"
	{
		return rsh
	}
yyrule38: // "<=>"
	{
		return nulleq
	}
yyrule39: // "@"
	{
		return at
	}
yyrule40: // "?"
	{
		return placeholder
	}
yyrule41: // {abs}
	{
		lval.ident = string(l.val)
		return abs
	}
yyrule42: // {add}
	{
		return add
	}
yyrule43: // {adddate}
	{
		lval.ident = string(l.val)
		return addDate
	}
yyrule44: // {admin}
	{
		lval.ident = string(l.val)
		return admin
	}
yyrule45: // {after}
	{
		lval.ident = string(l.val)
		return after
	}
yyrule46: // {all}
	{
		return all
	}
yyrule47: // {alter}
	{
		return alter
	}
yyrule48: // {analyze}
	{
		return analyze
	}
yyrule49: // {and}
	{
		return and
	}
yyrule50: // {any}
	{
		lval.ident = string(l.val)
		return any
	}
yyrule51: // {asc}
	{
		return asc
	}
yyrule52: // {as}
	{
		return as
	}
yyrule53: // {ascii}
	{
		lval.ident = string(l.val)
		return ascii
	}
yyrule54: // {auto_increment}
	{
		lval.ident = string(l.val)
		return autoIncrement
	}
yyrule55: // {avg}
	{
		lval.ident = string(l.val)
		return avg
	}
yyrule56: // {avg_row_length}
	{
		lval.ident = string(l.val)
		return avgRowLength
	}
yyrule57: // {begin}
	{
		lval.ident = string(l.val)
		return begin
	}
yyrule58: // {between}
	{
		return between
	}
yyrule59: // {binlog}
	{
		lval.ident = string(l.val)
		return binlog
	}
yyrule60: // {both}
	{
		return both
	}
yyrule61: // {btree}
	{
		lval.ident = string(l.val)
		return btree
	}
yyrule62: // {by}
	{
		return by
	}
yyrule63: // {case}
	{
		return caseKwd
	}
yyrule64: // {cast}
	{
		lval.item = string(l.val)
		return cast
	}
yyrule65: // {character}
	{
		return character
	}
yyrule66: // {charset}
	{
		lval.ident = string(l.val)
		return charsetKwd
	}
yyrule67: // {check}
	{
		return check
	}
yyrule68: // {checksum}
	{
		lval.ident = string(l.val)
		return checksum
	}
yyrule69: // {coalesce}
	{
		lval.ident = string(l.val)
		return coalesce
	}
yyrule70: // {collate}
	{
		return collate
	}
yyrule71: // {collation}
	{
		lval.ident = string(l.val)
		return collation
	}
yyrule72: // {column}
	{
		return column
	}
yyrule73: // {columns}
	{
		lval.ident = string(l.val)
		return columns
	}
yyrule74: // {comment}
	{
		lval.ident = string(l.val)
		return comment
	}
yyrule75: // {commit}
	{
		lval.ident = string(l.val)
		return commit
	}
yyrule76: // {committed}
	{
		lval.ident = string(l.val)
		return committed
	}
yyrule77: // {compact}
	{
		lval.ident = string(l.val)
		return compact
	}
yyrule78: // {compressed}
	{
		lval.ident = string(l.val)
		return compressed
	}
yyrule79: // {compression}
	{
		lval.ident = string(l.val)
		return compression
	}
yyrule80: // {concat}
	{
		lval.ident = string(l.val)
		return concat
	}
yyrule81: // {concat_ws}
	{
		lval.ident = string(l.val)
		return concatWs
	}
yyrule82: // {connection}
	{
		lval.ident = string(l.val)
		return connection
	}
yyrule83: // {connection_id}
	{
		lval.ident = string(l.val)
		return connectionID
	}
yyrule84: // {constraint}
	{
		return constraint
	}
yyrule85: // {convert}
	{
		lval.item = string(l.val)
		return convert
	}
yyrule86: // {count}
	{
		lval.ident = string(l.val)
		return count
	}
yyrule87: // {create}
	{
		return create
	}
yyrule88: // {cross}
	{
		return cross
	}
yyrule89: // {curdate}
	{
		lval.item = string(l.val)
		return curDate
	}
yyrule90: // {current_date}
	{
		lval.item = string(l.val)
		return currentDate
	}
yyrule91: // {curtime}
	{
		lval.ident = string(l.val)
		return curTime
	}
yyrule92: // {current_time}
	{
		lval.item = string(l.val)
		return currentTime
	}
yyrule93: // {current_user}
	{
		lval.item = string(l.val)
		return currentUser
	}
yyrule94: // {database}
	{
		lval.item = string(l.val)
		return database
	}
yyrule95: // {databases}
	{
		return databases
	}
yyrule96: // {date_add}
	{
		lval.ident = string(l.val)
		return dateAdd
	}
yyrule97: // {date_format}
	{
		lval.ident = string(l.val)
		return dateFormat
	}
yyrule98: // {date_sub}
	{
		lval.ident = string(l.val)
		return dateSub
	}
yyrule99: // {day}
	{
		lval.ident = string(l.val)
		return day
	}
yyrule100: // {dayname}
	{
		lval.ident = string(l.val)
		return dayname
	}
yyrule101: // {dayofweek}
	{
		lval.ident = string(l.val)
		return dayofweek
	}
yyrule102: // {dayofmonth}
	{
		lval.ident = string(l.val)
		return dayofmonth
	}
yyrule103: // {dayofyear}
	{
		lval.ident = string(l.val)
		return dayofyear
	}
yyrule104: // {day_hour}
	{
		lval.item = string(l.val)
		return dayHour
	}
yyrule105: // {day_microsecond}
	{
		lval.item = string(l.val)
		return dayMicrosecond
	}
yyrule106: // {day_minute}
	{
		lval.item = string(l.val)
		return dayMinute
	}
yyrule107: // {day_second}
	{
		lval.item = string(l.val)
		return daySecond
	}
yyrule108: // {ddl}
	{
		return ddl
	}
yyrule109: // {deallocate}
	{
		lval.ident = string(l.val)
		return deallocate
	}
yyrule110: // {default}
	{
		return defaultKwd
	}
yyrule111: // {delayed}
	{
		return delayed
	}
yyrule112: // {delay_key_write}
	{
		lval.ident = string(l.val)
		return delayKeyWrite
	}
yyrule113: // {delete}
	{
		return deleteKwd
	}
yyrule114: // {desc}
	{
		return desc
	}
yyrule115: // {describe}
	{
		return describe
	}
yyrule116: // {drop}
	{
		return drop
	}
yyrule117: // {disable}
	{
		lval.ident = string(l.val)
		return disable
	}
yyrule118: // {distinct}
	{
		return distinct
	}
yyrule119: // {div}
	{
		return div
	}
yyrule120: // {do}
	{
		lval.ident = string(l.val)
		return do
	}
yyrule121: // {dual}
	{
		return dual
	}
yyrule122: // {duplicate}
	{
		lval.item = string(l.val)
		return duplicate
	}
yyrule123: // {dynamic}
	{
		lval.ident = string(l.val)
		return dynamic
	}
yyrule124: // {else}
	{
		return elseKwd
	}
yyrule125: // {enable}
	{
		lval.ident = string(l.val)
		return enable
	}
yyrule126: // {end}
	{
		lval.ident = string(l.val)
		return end
	}
yyrule127: // {engine}
	{
		lval.ident = string(l.val)
		return engine
	}
yyrule128: // {engines}
	{
		lval.ident = string(l.val)
		return engines
	}
yyrule129: // {execute}
	{
		lval.ident = string(l.val)
		return execute
	}
yyrule130: // {enum}
	{
		return enum
	}
yyrule131: // {escape}
	{
		lval.ident = string(l.val)
		return escape
	}
yyrule132: // {exists}
	{
		return exists
	}
yyrule133: // {explain}
	{
		return explain
	}
yyrule134: // {extract}
	{
		lval.item = string(l.val)
		return extract
	}
yyrule135: // {fields}
	{
		lval.ident = string(l.val)
		return fields
	}
yyrule136: // {first}
	{
		lval.ident = string(l.val)
		return first
	}
yyrule137: // {fixed}
	{
		lval.ident = string(l.val)
		return fixed
	}
yyrule138: // {for}
	{
		return forKwd
	}
yyrule139: // {force}
	{
		return force
	}
yyrule140: // {foreign}
	{
		return foreign
	}
yyrule141: // {found_rows}
	{
		lval.ident = string(l.val)
		return foundRows
	}
yyrule142: // {from}
	{
		return from
	}
yyrule143: // {full}
	{
		lval.ident = string(l.val)
		return full
	}
yyrule144: // {fulltext}
	{
		return fulltext
	}
yyrule145: // {grant}
	{
		return grant
	}
yyrule146: // {grants}
	{
		lval.ident = string(l.val)
		return grants
	}
yyrule147: // {greatest}
	{
		lval.ident = string(l.val)
		return greatest
	}
yyrule148: // {group}
	{
		return group
	}
yyrule149: // {group_concat}
	{
		lval.ident = string(l.val)
		return groupConcat
	}
yyrule150: // {hash}
	{
		lval.ident = string(l.val)
		return hash
	}
yyrule151: // {having}
	{
		return having
	}
yyrule152: // {hex}
	{
		lval.ident = string(l.val)
		return hex
	}
yyrule153: // {high_priority}
	{
		return highPriority
	}
yyrule154: // {hour}
	{
		lval.ident = string(l.val)
		return hour
	}
yyrule155: // {hour_microsecond}
	{
		lval.item = string(l.val)
		return hourMicrosecond
	}
yyrule156: // {hour_minute}
	{
		lval.item = string(l.val)
		return hourMinute
	}
yyrule157: // {hour_second}
	{
		lval.item = string(l.val)
		return hourSecond
	}
yyrule158: // {identified}
	{
		lval.ident = string(l.val)
		return identified
	}
yyrule159: // {if}
	{
		lval.item = string(l.val)
		return ifKwd
	}
yyrule160: // {ifnull}
	{
		lval.ident = string(l.val)
		return ifNull
	}
yyrule161: // {isnull}
	{
		lval.ident = string(l.val)
		return isNull
	}
yyrule162: // {ignore}
	{
		return ignore
	}
yyrule163: // {index}
	{
		return index
	}
yyrule164: // {inner}
	{
		return inner
	}
yyrule165: // {insert}
	{
		return insert
	}
yyrule166: // {interval}
	{
		return interval
	}
yyrule167: // {into}
	{
		return into
	}
yyrule168: // {in}
	{
		return in
	}
yyrule169: // {is}
	{
		return is
	}
yyrule170: // {isolation}
	{
		lval.ident = string(l.val)
		return isolation
	}
yyrule171: // {join}
	{
		return join
	}
yyrule172: // {key}
	{
		return key
	}
yyrule173: // {key_block_size}
	{
		lval.ident = string(l.val)
		return keyBlockSize
	}
yyrule174: // {keys}
	{
		return keys
	}
yyrule175: // {last_insert_id}
	{
		lval.ident = string(l.val)
		return lastInsertID
	}
yyrule176: // {leading}
	{
		return leading
	}
yyrule177: // {left}
	{
		lval.item = string(l.val)
		return left
	}
yyrule178: // {length}
	{
		lval.ident = string(l.val)
		return length
	}
yyrule179: // {level}
	{
		lval.ident = string(l.val)
		return level
	}
yyrule180: // {like}
	{
		return like
	}
yyrule181: // {limit}
	{
		return limit
	}
yyrule182: // {local}
	{
		lval.ident = string(l.val)
		return local
	}
yyrule183: // {locate}
	{
		lval.ident = string(l.val)
		return locate
	}
yyrule184: // {lock}
	{
		return lock
	}
yyrule185: // {lower}
	{
		lval.ident = string(l.val)
		return lower
	}
yyrule186: // {lcase}
	{
		lval.ident = string(l.val)
		return lcase
	}
yyrule187: // {low_priority}
	{
		return lowPriority
	}
yyrule188: // {ltrim}
	{
		lval.ident = string(l.val)
		return ltrim
	}
yyrule189: // {max}
	{
		lval.ident = string(l.val)
		return max
	}
yyrule190: // {max_rows}
	{
		lval.ident = string(l.val)
		return maxRows
	}
yyrule191: // {microsecond}
	{
		lval.ident = string(l.val)
		return microsecond
	}
yyrule192: // {min}
	{
		lval.ident = string(l.val)
		return min
	}
yyrule193: // {minute}
	{
		lval.ident = string(l.val)
		return minute
	}
yyrule194: // {minute_microsecond}
	{
		lval.item = string(l.val)
		return minuteMicrosecond
	}
yyrule195: // {minute_second}
	{
		lval.item = string(l.val)
		return minuteSecond
	}
yyrule196: // {min_rows}
	{
		lval.ident = string(l.val)
		return minRows
	}
yyrule197: // {mod}
	{
		return mod
	}
yyrule198: // {mode}
	{
		lval.ident = string(l.val)
		return mode
	}
yyrule199: // {month}
	{
		lval.ident = string(l.val)
		return month
	}
yyrule200: // {monthname}
	{
		lval.ident = string(l.val)
		return monthname
	}
yyrule201: // {names}
	{
		lval.ident = string(l.val)
		return names
	}
yyrule202: // {national}
	{
		lval.ident = string(l.val)
		return national
	}
yyrule203: // {not}
	{
		return not
	}
yyrule204: // {offset}
	{
		lval.ident = string(l.val)
		return offset
	}
yyrule205: // {on}
	{
		return on
	}
yyrule206: // {only}
	{
		lval.ident = string(l.val)
		return only
	}
yyrule207: // {option}
	{
		return option
	}
yyrule208: // {order}
	{
		return order
	}
yyrule209: // {or}
	{
		return or
	}
yyrule210: // {outer}
	{
		return outer
	}
yyrule211: // {password}
	{
		lval.ident = string(l.val)
		return password
	}
yyrule212: // {pow}
	{
		lval.ident = string(l.val)
		return pow
	}
yyrule213: // {power}
	{
		lval.ident = string(l.val)
		return power
	}
yyrule214: // {prepare}
	{
		lval.ident = string(l.val)
		return prepare
	}
yyrule215: // {primary}
	{
		return primary
	}
yyrule216: // {privileges}
	{
		lval.ident = string(l.val)
		return privileges
	}
yyrule217: // {procedure}
	{
		return procedure
	}
yyrule218: // {quarter}
	{
		lval.ident = string(l.val)
		return quarter
	}
yyrule219: // {quick}
	{
		lval.ident = string(l.val)
		return quick
	}
yyrule220: // {redundant}
	{
		lval.ident = string(l.val)
		return redundant
	}
yyrule221: // {right}
	{
		return right
	}
yyrule222: // {rollback}
	{
		lval.ident = string(l.val)
		return rollback
	}
yyrule223: // {round}
	{
		lval.ident = string(l.val)
		return round
	}
yyrule224: // {row}
	{
		lval.ident = string(l.val)
		return row
	}
yyrule225: // {row_format}
	{
		lval.ident = string(l.val)
		return rowFormat
	}
yyrule226: // {schema}
	{
		lval.item = string(l.val)
		return schema
	}
yyrule227: // {schemas}
	{
		return schemas
	}
yyrule228: // {serializable}
	{
		lval.item = string(l.val)
		return serializable
	}
yyrule229: // {session}
	{
		lval.ident = string(l.val)
		return session
	}
yyrule230: // {some}
	{
		lval.ident = string(l.val)
		return some
	}
yyrule231: // {space}
	{
		lval.ident = string(l.val)
		return space
	}
yyrule232: // {start}
	{
		lval.ident = string(l.val)
		return start
	}
yyrule233: // {statsPersistent}
	{
		lval.ident = string(l.val)
		return statsPersistent
	}
yyrule234: // {status}
	{
		lval.ident = string(l.val)
		return status
	}
yyrule235: // {get_lock}
	{
		lval.ident = string(l.val)
		return getLock
	}
yyrule236: // {global}
	{
		lval.ident = string(l.val)
		return global
	}
yyrule237: // {rand}
	{
		lval.ident = string(l.val)
		return rand
	}
yyrule238: // {read}
	{
		return read
	}
yyrule239: // {release_lock}
	{
		lval.ident = string(l.val)
		return releaseLock
	}
yyrule240: // {repeat}
	{
		lval.item = string(l.val)
		return repeat
	}
yyrule241: // {repeatable}
	{
		lval.ident = string(l.val)
		return repeatable
	}
yyrule242: // {regexp}
	{
		return regexpKwd
	}
yyrule243: // {replace}
	{
		lval.item = string(l.val)
		return replace
	}
yyrule244: // {references}
	{
		return references
	}
yyrule245: // {rlike}
	{
		return rlike
	}
yyrule246: // {rtrim}
	{
		lval.ident = string(l.val)
		return rtrim
	}
yyrule247: // {reverse}
	{
		lval.ident = string(l.val)
		return reverse
	}
yyrule248: // {sys_var}
	{
		lval.item = string(l.val)
		return sysVar
	}
yyrule249: // {user_var}
	{
		lval.item = string(l.val)
		return userVar
	}
yyrule250: // {utc_date}
	{
		lval.item = string(l.val)
		return utcDate
	}
yyrule251: // {second}
	{
		lval.ident = string(l.val)
		return second
	}
yyrule252: // {second_microsecond}
	{
		lval.item = string(l.val)
		return secondMicrosecond
	}
yyrule253: // {select}
	{
		return selectKwd
	}
yyrule254: // {set}
	{
		return set
	}
yyrule255: // {share}
	{
		return share
	}
yyrule256: // {show}
	{
		return show
	}
yyrule257: // {sleep}
	{
		lval.ident = string(l.val)
		return sleep
	}
yyrule258: // {subdate}
	{
		lval.ident = string(l.val)
		return subDate
	}
yyrule259: // {strcmp}
	{
		lval.item = string(l.val)
		return strcmp
	}
yyrule260: // {substr}
	{
		lval.ident = string(l.val)
		return substring
	}
yyrule261: // {substring}
	{
		lval.ident = string(l.val)
		return substring
	}
yyrule262: // {substring_index}
	{
		lval.ident = string(l.val)
		return substringIndex
	}
yyrule263: // {sum}
	{
		lval.ident = string(l.val)
		return sum
	}
yyrule264: // {sysdate}
	{
		lval.item = string(l.val)
		return sysDate
	}
yyrule265: // {table}
	{
		return tableKwd
	}
yyrule266: // {tables}
	{
		lval.ident = string(l.val)
		return tables
	}
yyrule267: // {then}
	{
		return then
	}
yyrule268: // {to}
	{
		return to
	}
yyrule269: // {trailing}
	{
		return trailing
	}
yyrule270: // {transaction}
	{
		lval.ident = string(l.val)
		return transaction
	}
yyrule271: // {triggers}
	{
		lval.ident = string(l.val)
		return triggers
	}
yyrule272: // {trim}
	{
		lval.ident = string(l.val)
		return trim
	}
yyrule273: // {truncate}
	{
		lval.item = string(l.val)
		return truncate
	}
yyrule274: // {uncommitted}
	{
		lval.ident = string(l.val)
		return uncommitted
	}
yyrule275: // {union}
	{
		return union
	}
yyrule276: // {unique}
	{
		return unique
	}
yyrule277: // {unknown}
	{
		lval.ident = string(l.val)
		return unknown
	}
yyrule278: // {nullif}
	{
		lval.ident = string(l.val)
		return nullIf
	}
yyrule279: // {unlock}
	{
		return unlock
	}
yyrule280: // {update}
	{
		return update
	}
yyrule281: // {upper}
	{
		lval.ident = string(l.val)
		return upper
	}
yyrule282: // {ucase}
	{
		lval.ident = string(l.val)
		return ucase
	}
yyrule283: // {use}
	{
		return use
	}
yyrule284: // {user}
	{
		lval.ident = string(l.val)
		return user
	}
yyrule285: // {using}
	{
		return using
	}
yyrule286: // {value}
	{
		lval.ident = string(l.val)
		return value
	}
yyrule287: // {values}
	{
		return values
	}
yyrule288: // {variables}
	{
		lval.ident = string(l.val)
		return variables
	}
yyrule289: // {version}
	{
		lval.ident = string(l.val)
		return version
	}
yyrule290: // {warnings}
	{
		lval.ident = string(l.val)
		return warnings
	}
yyrule291: // {week}
	{
		lval.ident = string(l.val)
		return week
	}
yyrule292: // {weekday}
	{
		lval.ident = string(l.val)
		return weekday
	}
yyrule293: // {weekofyear}
	{
		lval.ident = string(l.val)
		return weekofyear
	}
yyrule294: // {when}
	{
		return when
	}
yyrule295: // {where}
	{
		return where
	}
yyrule296: // {write}
	{
		return write
	}
yyrule297: // {xor}
	{
		return xor
	}
yyrule298: // {yearweek}
	{
		lval.ident = string(l.val)
		return yearweek
	}
yyrule299: // {year_month}
	{
		lval.item = string(l.val)
		return yearMonth
	}
yyrule300: // {restrict}
	{
		lval.item = string(l.val)
		return restrict
	}
yyrule301: // {cascade}
	{
		lval.item = string(l.val)
		return cascade
	}
yyrule302: // {no}
	{
		lval.ident = string(l.val)
		return no
	}
yyrule303: // {action}
	{
		lval.ident = string(l.val)
		return action
	}
yyrule304: // {signed}
	{
		lval.ident = string(l.val)
		return signed
	}
yyrule305: // {unsigned}
	{
		return unsigned
	}
yyrule306: // {zerofill}
	{
		return zerofill
	}
yyrule307: // {null}
	{
		lval.item = nil
		return null
	}
yyrule308: // {false}
	{
		return falseKwd
	}
yyrule309: // {true}
	{
		return trueKwd
	}
yyrule310: // {calc_found_rows}
	{
		lval.ident = string(l.val)
		return calcFoundRows
	}
yyrule311: // {sql_cache}
	{
		lval.ident = string(l.val)
		return sqlCache
	}
yyrule312: // {sql_no_cache}
	{
		lval.ident = string(l.val)
		return sqlNoCache
	}
yyrule313: // {current_ts}
	{
		lval.item = string(l.val)
		return currentTs
	}
yyrule314: // {localtime}
	{
		return localTime
	}
yyrule315: // {localts}
	{
		return localTs
	}
yyrule316: // {now}
	{
		lval.ident = string(l.val)
		return now
	}
yyrule317: // {bit}
	{
		lval.ident = string(l.val)
		return bitType
	}
yyrule318: // {tiny}
	{
		lval.item = string(l.val)
		return tinyIntType
	}
yyrule319: // {tinyint}
	{
		lval.item = string(l.val)
		return tinyIntType
	}
yyrule320: // {smallint}
	{
		lval.item = string(l.val)
		return smallIntType
	}
yyrule321: // {mediumint}
	{
		lval.item = string(l.val)
		return mediumIntType
	}
yyrule322: // {bigint}
	{
		lval.item = string(l.val)
		return bigIntType
	}
yyrule323: // {decimal}
	{
		lval.item = string(l.val)
		return decimalType
	}
yyrule324: // {numeric}
	{
		lval.item = string(l.val)
		return numericType
	}
yyrule325: // {float}
	{
		lval.item = string(l.val)
		return floatType
	}
yyrule326: // {double}
	{
		lval.item = string(l.val)
		return doubleType
	}
yyrule327: // {precision}
	{
		lval.item = string(l.val)
		return precisionType
	}
yyrule328: // {real}
	{
		lval.item = string(l.val)
		return realType
	}
yyrule329: // {date}
	{
		lval.ident = string(l.val)
		return dateType
	}
yyrule330: // {time}
	{
		lval.ident = string(l.val)
		return timeType
	}
yyrule331: // {timestamp}
	{
		lval.ident = string(l.val)
		return timestampType
	}
yyrule332: // {datetime}
	{
		lval.ident = string(l.val)
		return datetimeType
	}
yyrule333: // {year}
	{
		lval.ident = string(l.val)
		return yearType
	}
yyrule334: // {char}
	{
		lval.item = string(l.val)
		return charType
	}
yyrule335: // {varchar}
	{
		lval.item = string(l.val)
		return varcharType
	}
yyrule336: // {binary}
	{
		lval.item = string(l.val)
		return binaryType
	}
yyrule337: // {varbinary}
	{
		lval.item = string(l.val)
		return varbinaryType
	}
yyrule338: // {tinyblob}
	{
		lval.item = string(l.val)
		return tinyblobType
	}
yyrule339: // {blob}
	{
		lval.item = string(l.val)
		return blobType
	}
yyrule340: // {mediumblob}
	{
		lval.item = string(l.val)
		return mediumblobType
	}
yyrule341: // {longblob}
	{
		lval.item = string(l.val)
		return longblobType
	}
yyrule342: // {tinytext}
	{
		lval.item = string(l.val)
		return tinytextType
	}
yyrule343: // {mediumtext}
	{
		lval.item = string(l.val)
		return mediumtextType
	}
yyrule344: // {text}
	{
		lval.ident = string(l.val)
		return textType
	}
yyrule345: // {longtext}
	{
		lval.item = string(l.val)
		return longtextType
	}
yyrule346: // {bool}
	{
		lval.ident = string(l.val)
		return boolType
	}
yyrule347: // {boolean}
	{
		lval.ident = string(l.val)
		return booleanType
	}
yyrule348: // {byte}
	{
		lval.item = string(l.val)
		return byteType
	}
yyrule349: // {int}
	{
		lval.item = string(l.val)
		return intType
	}
yyrule350: // {integer}
	{
		lval.item = string(l.val)
		return integerType
	}
yyrule351: // {ident}
	{
		lval.ident = string(l.val)
		return handleIdent(lval)
	}
yyrule352: // .
	{
		return c0
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	return int(unicode.ReplacementChar)
}

func (l *lexer) npos() (line, col int) {
	if line, col = l.nline, l.ncol; col == 0 {
		line--
		col = l.lcol + 1
	}
	return
}

func (l *lexer) str(lval *yySymType, pref string) int {
	l.sc = 0
	// TODO: performance issue.
	s := string(l.stringLit)
	l.stringLit = l.stringLit[0:0]
	v, err := stringutil.Unquote(pref + s)
	if err != nil {
		v = strings.TrimSuffix(s, pref)
	}
	lval.ident = v
	return stringLit
}

func (l *lexer) trimIdent(idt string) string {
	idt = strings.TrimPrefix(idt, "`")
	idt = strings.TrimSuffix(idt, "`")
	return idt
}

func handleIdent(lval *yySymType) int {
	s := lval.ident
	// A character string literal may have an optional character set introducer and COLLATE clause:
	// [_charset_name]'string' [COLLATE collation_name]
	// See https://dev.mysql.com/doc/refman/5.7/en/charset-literal.html
	if !strings.HasPrefix(s, "_") {
		return identifier
	}
	cs, _, err := charset.GetCharsetInfo(s[1:])
	if err != nil {
		return identifier
	}
	lval.item = cs
	return underscoreCS
}
