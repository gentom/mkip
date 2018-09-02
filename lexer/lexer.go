package lexer

type Lexer struct {
	input        string
	position     int  // current position for input
	readPosition int  // next position of current position
	ch           byte // character that is currently under inspection
}
