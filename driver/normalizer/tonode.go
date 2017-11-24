package normalizer

import (
	"gopkg.in/bblfsh/sdk.v1/uast"
)

// ToNode is an instance of `uast.ObjectToNode`, defining how to transform an
// into a UAST (`uast.Node`).
//
// https://godoc.org/gopkg.in/bblfsh/sdk.v1/uast#ObjectToNode
var ToNode = &uast.ObjectToNode{
	TopLevelIsRootNode: true,
	InternalTypeKey:    "kind",
	OffsetKey:          "pos",
	EndOffsetKey:       "end",
	LineKey:            "line",
	ColumnKey:          "col",
	TokenKeys: map[string]bool{
		"text": true,
	},
	IsNode: func(v map[string]interface{}) bool {
		_, ok := v["kind"].(string)
		return ok
	},
}
