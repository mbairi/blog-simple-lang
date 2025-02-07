// Code generated from src/Lang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package BaseLang // Lang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type LangParser struct {
	*antlr.BaseParser
}

var LangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func langParserInit() {
	staticData := &LangParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'+'", "'-'", "'*'", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "OpenParen", "CloseParen", "Plus", "Minus", "Multiply", "Divide",
		"WhiteSpaces", "DecimalLiteral",
	}
	staticData.RuleNames = []string{
		"unit",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 8, 28, 2, 0, 7, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 9, 8,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 5, 0, 23, 8, 0, 10, 0, 12, 0, 26, 9, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 31,
		0, 8, 1, 0, 0, 0, 2, 3, 6, 0, -1, 0, 3, 4, 5, 1, 0, 0, 4, 5, 3, 0, 0, 0,
		5, 6, 5, 2, 0, 0, 6, 9, 1, 0, 0, 0, 7, 9, 5, 8, 0, 0, 8, 2, 1, 0, 0, 0,
		8, 7, 1, 0, 0, 0, 9, 24, 1, 0, 0, 0, 10, 11, 10, 5, 0, 0, 11, 12, 5, 6,
		0, 0, 12, 23, 3, 0, 0, 6, 13, 14, 10, 4, 0, 0, 14, 15, 5, 5, 0, 0, 15,
		23, 3, 0, 0, 5, 16, 17, 10, 3, 0, 0, 17, 18, 5, 3, 0, 0, 18, 23, 3, 0,
		0, 4, 19, 20, 10, 2, 0, 0, 20, 21, 5, 4, 0, 0, 21, 23, 3, 0, 0, 3, 22,
		10, 1, 0, 0, 0, 22, 13, 1, 0, 0, 0, 22, 16, 1, 0, 0, 0, 22, 19, 1, 0, 0,
		0, 23, 26, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 1, 1,
		0, 0, 0, 26, 24, 1, 0, 0, 0, 3, 8, 22, 24,
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

// LangParserInit initializes any static state used to implement LangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LangParserInit() {
	staticData := &LangParserStaticData
	staticData.once.Do(langParserInit)
}

// NewLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLangParser(input antlr.TokenStream) *LangParser {
	LangParserInit()
	this := new(LangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &LangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Lang.g4"

	return this
}

// LangParser tokens.
const (
	LangParserEOF            = antlr.TokenEOF
	LangParserOpenParen      = 1
	LangParserCloseParen     = 2
	LangParserPlus           = 3
	LangParserMinus          = 4
	LangParserMultiply       = 5
	LangParserDivide         = 6
	LangParserWhiteSpaces    = 7
	LangParserDecimalLiteral = 8
)

// LangParserRULE_unit is the LangParser rule.
const LangParserRULE_unit = 0

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBase returns the base token.
	GetBase() antlr.Token

	// SetBase sets the base token.
	SetBase(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IUnitContext

	// GetBracketContent returns the bracketContent rule contexts.
	GetBracketContent() IUnitContext

	// GetRight returns the right rule contexts.
	GetRight() IUnitContext

	// SetLeft sets the left rule contexts.
	SetLeft(IUnitContext)

	// SetBracketContent sets the bracketContent rule contexts.
	SetBracketContent(IUnitContext)

	// SetRight sets the right rule contexts.
	SetRight(IUnitContext)

	// Getter signatures
	OpenParen() antlr.TerminalNode
	CloseParen() antlr.TerminalNode
	AllUnit() []IUnitContext
	Unit(i int) IUnitContext
	DecimalLiteral() antlr.TerminalNode
	Divide() antlr.TerminalNode
	Multiply() antlr.TerminalNode
	Plus() antlr.TerminalNode
	Minus() antlr.TerminalNode

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	left           IUnitContext
	bracketContent IUnitContext
	base           antlr.Token
	right          IUnitContext
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_unit
	return p
}

func InitEmptyUnitContext(p *UnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_unit
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) GetBase() antlr.Token { return s.base }

func (s *UnitContext) SetBase(v antlr.Token) { s.base = v }

func (s *UnitContext) GetLeft() IUnitContext { return s.left }

func (s *UnitContext) GetBracketContent() IUnitContext { return s.bracketContent }

func (s *UnitContext) GetRight() IUnitContext { return s.right }

func (s *UnitContext) SetLeft(v IUnitContext) { s.left = v }

func (s *UnitContext) SetBracketContent(v IUnitContext) { s.bracketContent = v }

func (s *UnitContext) SetRight(v IUnitContext) { s.right = v }

func (s *UnitContext) OpenParen() antlr.TerminalNode {
	return s.GetToken(LangParserOpenParen, 0)
}

func (s *UnitContext) CloseParen() antlr.TerminalNode {
	return s.GetToken(LangParserCloseParen, 0)
}

func (s *UnitContext) AllUnit() []IUnitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnitContext); ok {
			len++
		}
	}

	tst := make([]IUnitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnitContext); ok {
			tst[i] = t.(IUnitContext)
			i++
		}
	}

	return tst
}

func (s *UnitContext) Unit(i int) IUnitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *UnitContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(LangParserDecimalLiteral, 0)
}

func (s *UnitContext) Divide() antlr.TerminalNode {
	return s.GetToken(LangParserDivide, 0)
}

func (s *UnitContext) Multiply() antlr.TerminalNode {
	return s.GetToken(LangParserMultiply, 0)
}

func (s *UnitContext) Plus() antlr.TerminalNode {
	return s.GetToken(LangParserPlus, 0)
}

func (s *UnitContext) Minus() antlr.TerminalNode {
	return s.GetToken(LangParserMinus, 0)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (s *UnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LangVisitor:
		return t.VisitUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LangParser) Unit() (localctx IUnitContext) {
	return p.unit(0)
}

func (p *LangParser) unit(_p int) (localctx IUnitContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewUnitContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IUnitContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, LangParserRULE_unit, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(8)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LangParserOpenParen:
		{
			p.SetState(3)
			p.Match(LangParserOpenParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(4)

			var _x = p.unit(0)

			localctx.(*UnitContext).bracketContent = _x
		}
		{
			p.SetState(5)
			p.Match(LangParserCloseParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LangParserDecimalLiteral:
		{
			p.SetState(7)

			var _m = p.Match(LangParserDecimalLiteral)

			localctx.(*UnitContext).base = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(22)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_unit)
				p.SetState(10)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}

				{
					p.SetState(11)
					p.Match(LangParserDivide)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(12)

					var _x = p.unit(6)

					localctx.(*UnitContext).right = _x
				}

			case 2:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_unit)
				p.SetState(13)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}

				{
					p.SetState(14)
					p.Match(LangParserMultiply)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(15)

					var _x = p.unit(5)

					localctx.(*UnitContext).right = _x
				}

			case 3:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_unit)
				p.SetState(16)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}

				{
					p.SetState(17)
					p.Match(LangParserPlus)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(18)

					var _x = p.unit(4)

					localctx.(*UnitContext).right = _x
				}

			case 4:
				localctx = NewUnitContext(p, _parentctx, _parentState)
				localctx.(*UnitContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_unit)
				p.SetState(19)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}

				{
					p.SetState(20)
					p.Match(LangParserMinus)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(21)

					var _x = p.unit(3)

					localctx.(*UnitContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *LangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *UnitContext = nil
		if localctx != nil {
			t = localctx.(*UnitContext)
		}
		return p.Unit_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LangParser) Unit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
