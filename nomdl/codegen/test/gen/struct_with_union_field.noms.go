// This file was generated by nomdl/codegen.

package gen

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __genPackageInFile_struct_with_union_field_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("StructWithUnionField",
			[]types.Field{
				types.Field{"a", types.MakePrimitiveTypeRef(types.Float32Kind), false},
			},
			types.Choices{
				types.Field{"b", types.MakePrimitiveTypeRef(types.Float64Kind), false},
				types.Field{"c", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"d", types.MakePrimitiveTypeRef(types.BlobKind), false},
				types.Field{"e", types.MakePrimitiveTypeRef(types.ValueKind), false},
				types.Field{"f", types.MakeCompoundTypeRef(types.SetKind, types.MakePrimitiveTypeRef(types.UInt8Kind)), false},
			},
		),
	}, []ref.Ref{})
	__genPackageInFile_struct_with_union_field_CachedRef = types.RegisterPackage(&p)
}

// StructWithUnionField

type StructWithUnionField struct {
	_a           float32
	__unionIndex uint32
	__unionValue types.Value
	ref          *ref.Ref
}

func NewStructWithUnionField() StructWithUnionField {
	return StructWithUnionField{
		_a:           float32(0),
		__unionIndex: 0,
		__unionValue: types.Float64(0),
		ref:          &ref.Ref{},
	}
}

type StructWithUnionFieldDef struct {
	A            float32
	__unionIndex uint32
	__unionValue types.Value
}

func (def StructWithUnionFieldDef) New() StructWithUnionField {
	return StructWithUnionField{
		_a:           def.A,
		__unionIndex: def.__unionIndex,
		__unionValue: def.__unionValue,
		ref:          &ref.Ref{},
	}
}

func (s StructWithUnionField) Def() (d StructWithUnionFieldDef) {
	d.A = s._a
	d.__unionIndex = s.__unionIndex
	d.__unionValue = s.__unionValue
	return
}

var __typeRefForStructWithUnionField types.TypeRef

func (m StructWithUnionField) TypeRef() types.TypeRef {
	return __typeRefForStructWithUnionField
}

func init() {
	__typeRefForStructWithUnionField = types.MakeTypeRef(__genPackageInFile_struct_with_union_field_CachedRef, 0)
	types.RegisterStruct(__typeRefForStructWithUnionField, builderForStructWithUnionField, readerForStructWithUnionField)
}

func builderForStructWithUnionField(values []types.Value) types.Value {
	i := 0
	s := StructWithUnionField{ref: &ref.Ref{}}
	s._a = float32(values[i].(types.Float32))
	i++
	s.__unionIndex = uint32(values[i].(types.UInt32))
	i++
	s.__unionValue = values[i]
	i++
	return s
}

func readerForStructWithUnionField(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(StructWithUnionField)
	values = append(values, types.Float32(s._a))
	values = append(values, types.UInt32(s.__unionIndex))
	values = append(values, s.__unionValue)
	return values
}

func (s StructWithUnionField) Equals(other types.Value) bool {
	return other != nil && __typeRefForStructWithUnionField.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s StructWithUnionField) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s StructWithUnionField) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeRefForStructWithUnionField.Chunks()...)
	chunks = append(chunks, s.__unionValue.Chunks()...)
	return
}

func (s StructWithUnionField) A() float32 {
	return s._a
}

func (s StructWithUnionField) SetA(val float32) StructWithUnionField {
	s._a = val
	s.ref = &ref.Ref{}
	return s
}

func (s StructWithUnionField) B() (val float64, ok bool) {
	if s.__unionIndex != 0 {
		return
	}
	return float64(s.__unionValue.(types.Float64)), true
}

func (s StructWithUnionField) SetB(val float64) StructWithUnionField {
	s.__unionIndex = 0
	s.__unionValue = types.Float64(val)
	s.ref = &ref.Ref{}
	return s
}

func (def StructWithUnionFieldDef) B() (val float64, ok bool) {
	if def.__unionIndex != 0 {
		return
	}
	return float64(def.__unionValue.(types.Float64)), true
}

func (def StructWithUnionFieldDef) SetB(val float64) StructWithUnionFieldDef {
	def.__unionIndex = 0
	def.__unionValue = types.Float64(val)
	return def
}

func (s StructWithUnionField) C() (val string, ok bool) {
	if s.__unionIndex != 1 {
		return
	}
	return s.__unionValue.(types.String).String(), true
}

func (s StructWithUnionField) SetC(val string) StructWithUnionField {
	s.__unionIndex = 1
	s.__unionValue = types.NewString(val)
	s.ref = &ref.Ref{}
	return s
}

func (def StructWithUnionFieldDef) C() (val string, ok bool) {
	if def.__unionIndex != 1 {
		return
	}
	return def.__unionValue.(types.String).String(), true
}

func (def StructWithUnionFieldDef) SetC(val string) StructWithUnionFieldDef {
	def.__unionIndex = 1
	def.__unionValue = types.NewString(val)
	return def
}

func (s StructWithUnionField) D() (val types.Blob, ok bool) {
	if s.__unionIndex != 2 {
		return
	}
	return s.__unionValue.(types.Blob), true
}

func (s StructWithUnionField) SetD(val types.Blob) StructWithUnionField {
	s.__unionIndex = 2
	s.__unionValue = val
	s.ref = &ref.Ref{}
	return s
}

func (def StructWithUnionFieldDef) D() (val types.Blob, ok bool) {
	if def.__unionIndex != 2 {
		return
	}
	return def.__unionValue.(types.Blob), true
}

func (def StructWithUnionFieldDef) SetD(val types.Blob) StructWithUnionFieldDef {
	def.__unionIndex = 2
	def.__unionValue = val
	return def
}

func (s StructWithUnionField) E() (val types.Value, ok bool) {
	if s.__unionIndex != 3 {
		return
	}
	return s.__unionValue, true
}

func (s StructWithUnionField) SetE(val types.Value) StructWithUnionField {
	s.__unionIndex = 3
	s.__unionValue = val
	s.ref = &ref.Ref{}
	return s
}

func (def StructWithUnionFieldDef) E() (val types.Value, ok bool) {
	if def.__unionIndex != 3 {
		return
	}
	return def.__unionValue, true
}

func (def StructWithUnionFieldDef) SetE(val types.Value) StructWithUnionFieldDef {
	def.__unionIndex = 3
	def.__unionValue = val
	return def
}

func (s StructWithUnionField) F() (val SetOfUInt8, ok bool) {
	if s.__unionIndex != 4 {
		return
	}
	return s.__unionValue.(SetOfUInt8), true
}

func (s StructWithUnionField) SetF(val SetOfUInt8) StructWithUnionField {
	s.__unionIndex = 4
	s.__unionValue = val
	s.ref = &ref.Ref{}
	return s
}

func (def StructWithUnionFieldDef) F() (val SetOfUInt8Def, ok bool) {
	if def.__unionIndex != 4 {
		return
	}
	return def.__unionValue.(SetOfUInt8).Def(), true
}

func (def StructWithUnionFieldDef) SetF(val SetOfUInt8Def) StructWithUnionFieldDef {
	def.__unionIndex = 4
	def.__unionValue = val.New()
	return def
}

// SetOfUInt8

type SetOfUInt8 struct {
	s   types.Set
	ref *ref.Ref
}

func NewSetOfUInt8() SetOfUInt8 {
	return SetOfUInt8{types.NewSet(), &ref.Ref{}}
}

type SetOfUInt8Def map[uint8]bool

func (def SetOfUInt8Def) New() SetOfUInt8 {
	l := make([]types.Value, len(def))
	i := 0
	for d, _ := range def {
		l[i] = types.UInt8(d)
		i++
	}
	return SetOfUInt8{types.NewSet(l...), &ref.Ref{}}
}

func (s SetOfUInt8) Def() SetOfUInt8Def {
	def := make(map[uint8]bool, s.Len())
	s.s.Iter(func(v types.Value) bool {
		def[uint8(v.(types.UInt8))] = true
		return false
	})
	return def
}

func (s SetOfUInt8) Equals(other types.Value) bool {
	return other != nil && __typeRefForSetOfUInt8.Equals(other.TypeRef()) && s.Ref() == other.Ref()
}

func (s SetOfUInt8) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SetOfUInt8) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, s.TypeRef().Chunks()...)
	chunks = append(chunks, s.s.Chunks()...)
	return
}

// A Noms Value that describes SetOfUInt8.
var __typeRefForSetOfUInt8 types.TypeRef

func (m SetOfUInt8) TypeRef() types.TypeRef {
	return __typeRefForSetOfUInt8
}

func init() {
	__typeRefForSetOfUInt8 = types.MakeCompoundTypeRef(types.SetKind, types.MakePrimitiveTypeRef(types.UInt8Kind))
	types.RegisterValue(__typeRefForSetOfUInt8, builderForSetOfUInt8, readerForSetOfUInt8)
}

func builderForSetOfUInt8(v types.Value) types.Value {
	return SetOfUInt8{v.(types.Set), &ref.Ref{}}
}

func readerForSetOfUInt8(v types.Value) types.Value {
	return v.(SetOfUInt8).s
}

func (s SetOfUInt8) Empty() bool {
	return s.s.Empty()
}

func (s SetOfUInt8) Len() uint64 {
	return s.s.Len()
}

func (s SetOfUInt8) Has(p uint8) bool {
	return s.s.Has(types.UInt8(p))
}

type SetOfUInt8IterCallback func(p uint8) (stop bool)

func (s SetOfUInt8) Iter(cb SetOfUInt8IterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(uint8(v.(types.UInt8)))
	})
}

type SetOfUInt8IterAllCallback func(p uint8)

func (s SetOfUInt8) IterAll(cb SetOfUInt8IterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(uint8(v.(types.UInt8)))
	})
}

func (s SetOfUInt8) IterAllP(concurrency int, cb SetOfUInt8IterAllCallback) {
	s.s.IterAllP(concurrency, func(v types.Value) {
		cb(uint8(v.(types.UInt8)))
	})
}

type SetOfUInt8FilterCallback func(p uint8) (keep bool)

func (s SetOfUInt8) Filter(cb SetOfUInt8FilterCallback) SetOfUInt8 {
	out := s.s.Filter(func(v types.Value) bool {
		return cb(uint8(v.(types.UInt8)))
	})
	return SetOfUInt8{out, &ref.Ref{}}
}

func (s SetOfUInt8) Insert(p ...uint8) SetOfUInt8 {
	return SetOfUInt8{s.s.Insert(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfUInt8) Remove(p ...uint8) SetOfUInt8 {
	return SetOfUInt8{s.s.Remove(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfUInt8) Union(others ...SetOfUInt8) SetOfUInt8 {
	return SetOfUInt8{s.s.Union(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfUInt8) Subtract(others ...SetOfUInt8) SetOfUInt8 {
	return SetOfUInt8{s.s.Subtract(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfUInt8) Any() uint8 {
	return uint8(s.s.Any().(types.UInt8))
}

func (s SetOfUInt8) fromStructSlice(p []SetOfUInt8) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfUInt8) fromElemSlice(p []uint8) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = types.UInt8(v)
	}
	return r
}
