package php7

import (
	"github.com/i582/php-parser/internal/position"
	"github.com/i582/php-parser/internal/scanner"
	"github.com/i582/php-parser/pkg/ast"
	"github.com/i582/php-parser/pkg/cfg"
	"github.com/i582/php-parser/pkg/errors"
	"github.com/i582/php-parser/pkg/token"
)

// Parser structure
type Parser struct {
	Lexer          *scanner.Lexer
	currentToken   *token.Token
	rootNode       ast.Vertex
	errHandlerFunc func(*errors.Error)
	builder        *position.Builder
}

// NewParser creates and returns new Parser
func NewParser(lexer *scanner.Lexer, config cfg.Config) *Parser {
	return &Parser{
		Lexer:          lexer,
		errHandlerFunc: config.ErrorHandlerFunc,
		builder:        position.NewBuilder(),
	}
}

func (p *Parser) Lex(lval *yySymType) int {
	t := p.Lexer.Lex()

	p.currentToken = t
	lval.token = t

	return int(t.ID)
}

func (p *Parser) Error(msg string) {
	if p.errHandlerFunc == nil {
		return
	}

	p.errHandlerFunc(errors.NewError(msg, p.currentToken.Position))
}

// Parse the php7 Parser entrypoint
func (p *Parser) Parse() int {
	p.rootNode = nil

	return yyParse(p)
}

// GetRootNode returns root node
func (p *Parser) GetRootNode() ast.Vertex {
	return p.rootNode
}

// helpers

func lastNode(nn []ast.Vertex) ast.Vertex {
	if len(nn) == 0 {
		return nil
	}
	return nn[len(nn)-1]
}
