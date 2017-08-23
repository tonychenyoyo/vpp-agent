// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.
package memif

import "reflect"

var Types = map[string]reflect.Type{
	"MemifCreate":      reflect.TypeOf((*MemifCreate)(nil)).Elem(),
	"MemifCreateReply": reflect.TypeOf((*MemifCreateReply)(nil)).Elem(),
	"MemifDelete":      reflect.TypeOf((*MemifDelete)(nil)).Elem(),
	"MemifDeleteReply": reflect.TypeOf((*MemifDeleteReply)(nil)).Elem(),
	"MemifDetails":     reflect.TypeOf((*MemifDetails)(nil)).Elem(),
	"MemifDump":        reflect.TypeOf((*MemifDump)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"NewMemifCreate":      reflect.ValueOf(NewMemifCreate),
	"NewMemifCreateReply": reflect.ValueOf(NewMemifCreateReply),
	"NewMemifDelete":      reflect.ValueOf(NewMemifDelete),
	"NewMemifDeleteReply": reflect.ValueOf(NewMemifDeleteReply),
	"NewMemifDetails":     reflect.ValueOf(NewMemifDetails),
	"NewMemifDump":        reflect.ValueOf(NewMemifDump),
}

var Variables = map[string]reflect.Value{}

var Consts = map[string]reflect.Value{
	"VlAPIVersion": reflect.ValueOf(VlAPIVersion),
}
