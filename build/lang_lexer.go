// Code generated from src/Lang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package BaseLang

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type LangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LangLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func langlexerLexerInit() {
	staticData := &LangLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'+'", "'-'", "'*'", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "OpenParen", "CloseParen", "Plus", "Minus", "Multiply", "Divide",
		"WhiteSpaces", "DecimalLiteral",
	}
	staticData.RuleNames = []string{
		"OpenParen", "CloseParen", "Plus", "Minus", "Multiply", "Divide", "WhiteSpaces",
		"DecimalLiteral", "DecimalIntegerLiteral",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 8, 59, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 4, 6, 33,
		8, 6, 11, 6, 12, 6, 34, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 42, 8, 7, 10,
		7, 12, 7, 45, 9, 7, 1, 7, 3, 7, 48, 8, 7, 1, 8, 1, 8, 1, 8, 5, 8, 53, 8,
		8, 10, 8, 12, 8, 56, 9, 8, 3, 8, 58, 8, 8, 0, 0, 9, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 0, 1, 0, 4, 2, 0, 9, 9, 32, 32, 1, 0,
		48, 57, 1, 0, 49, 57, 2, 0, 48, 57, 95, 95, 62, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 1, 19, 1, 0, 0, 0, 3,
		21, 1, 0, 0, 0, 5, 23, 1, 0, 0, 0, 7, 25, 1, 0, 0, 0, 9, 27, 1, 0, 0, 0,
		11, 29, 1, 0, 0, 0, 13, 32, 1, 0, 0, 0, 15, 47, 1, 0, 0, 0, 17, 57, 1,
		0, 0, 0, 19, 20, 5, 40, 0, 0, 20, 2, 1, 0, 0, 0, 21, 22, 5, 41, 0, 0, 22,
		4, 1, 0, 0, 0, 23, 24, 5, 43, 0, 0, 24, 6, 1, 0, 0, 0, 25, 26, 5, 45, 0,
		0, 26, 8, 1, 0, 0, 0, 27, 28, 5, 42, 0, 0, 28, 10, 1, 0, 0, 0, 29, 30,
		5, 47, 0, 0, 30, 12, 1, 0, 0, 0, 31, 33, 7, 0, 0, 0, 32, 31, 1, 0, 0, 0,
		33, 34, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 36, 1,
		0, 0, 0, 36, 37, 6, 6, 0, 0, 37, 14, 1, 0, 0, 0, 38, 39, 3, 17, 8, 0, 39,
		43, 5, 46, 0, 0, 40, 42, 7, 1, 0, 0, 41, 40, 1, 0, 0, 0, 42, 45, 1, 0,
		0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 48, 1, 0, 0, 0, 45, 43,
		1, 0, 0, 0, 46, 48, 3, 17, 8, 0, 47, 38, 1, 0, 0, 0, 47, 46, 1, 0, 0, 0,
		48, 16, 1, 0, 0, 0, 49, 58, 5, 48, 0, 0, 50, 54, 7, 2, 0, 0, 51, 53, 7,
		3, 0, 0, 52, 51, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54,
		55, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 57, 49, 1, 0, 0,
		0, 57, 50, 1, 0, 0, 0, 58, 18, 1, 0, 0, 0, 6, 0, 34, 43, 47, 54, 57, 1,
		0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// LangLexerInit initializes any static state used to implement LangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LangLexerInit() {
	staticData := &LangLexerLexerStaticData
	staticData.once.Do(langlexerLexerInit)
}

// NewLangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLangLexer(input antlr.CharStream) *LangLexer {
	LangLexerInit()
	l := new(LangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LangLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Lang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LangLexer tokens.
const (
	LangLexerOpenParen      = 1
	LangLexerCloseParen     = 2
	LangLexerPlus           = 3
	LangLexerMinus          = 4
	LangLexerMultiply       = 5
	LangLexerDivide         = 6
	LangLexerWhiteSpaces    = 7
	LangLexerDecimalLiteral = 8
)
