package lexer

type Lexer struct {

  input string
  position int
  currentPosition int
  ch byte

}


func New(input String)*Lexer {
  l:=&lexer{input : input}
    return l
}

func (l *Lexer) readChar() {
 if l.readPosition >= len(l.input) {
  l.ch = 0


 } // end of if

 else {
 l.ch = l.input[l.readPosition]
 l.readPosition += 1

 }


}