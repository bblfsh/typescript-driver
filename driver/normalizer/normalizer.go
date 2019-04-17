package normalizer

import (
	. "github.com/bblfsh/sdk/v3/uast/transformer"
	"github.com/bblfsh/sdk/v3/uast/transformer/positioner"
)

var Preprocess = Transformers([][]Transformer{
	{Mappings(Preprocessors...)},
}...)

var Normalize = Transformers([][]Transformer{
	{Mappings(Normalizers...)},
}...)

// Preprocessors is a block of AST preprocessing rules rules.
var Preprocessors = []Mapping{
	// ObjectToNode defines how to normalize common fields of native AST
	// (like node type, token, positional information).
	//
	// https://godoc.org/github.com/bblfsh/sdk/v3/uast#ObjectToNode
	ObjectToNode{
		InternalTypeKey: "kind",
		OffsetKey:       "pos",
		EndOffsetKey:    "end",
		LineKey:         "line",
		ColumnKey:       "col",
	}.Mapping(),
}

// Code is a special block of transformations that are applied at the end
// and can access original source code file. It can be used to improve or
// fix positional information.
//
// https://godoc.org/github.com/bblfsh/sdk/v3/uast/transformer/positioner
var PreprocessCode = []CodeTransformer{
	positioner.FromOffset(),
}

// Normalizers is the main block of normalization rules to convert native AST to semantic UAST.
var Normalizers []Mapping
