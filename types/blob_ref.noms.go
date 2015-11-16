// This file was generated by nomdl/codegen.

package types

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
)

// RefOfBlob

type RefOfBlob struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfBlob(target ref.Ref) RefOfBlob {
	return RefOfBlob{target, &ref.Ref{}}
}

func (r RefOfBlob) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfBlob) Ref() ref.Ref {
	return EnsureRef(r.ref, r)
}

func (r RefOfBlob) Equals(other Value) bool {
	return other != nil && __typeForRefOfBlob.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfBlob) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfBlob) ChildValues() []Value {
	return nil
}

// A Noms Value that describes RefOfBlob.
var __typeForRefOfBlob Type

func (m RefOfBlob) Type() Type {
	return __typeForRefOfBlob
}

func init() {
	__typeForRefOfBlob = MakeCompoundType(RefKind, MakePrimitiveType(BlobKind))
	RegisterRef(__typeForRefOfBlob, builderForRefOfBlob)
}

func builderForRefOfBlob(r ref.Ref) Value {
	return NewRefOfBlob(r)
}

func (r RefOfBlob) TargetValue(cs chunks.ChunkSource) Blob {
	return ReadValue(r.target, cs).(Blob)
}

func (r RefOfBlob) SetTargetValue(val Blob, cs chunks.ChunkSink) RefOfBlob {
	return NewRefOfBlob(WriteValue(val, cs))
}