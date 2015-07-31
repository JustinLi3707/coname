// Code generated by protoc-gen-gogo.
// source: local.proto
// DO NOT EDIT!

package proto

import testing "testing"
import math_rand "math/rand"
import time "time"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
import fmt "fmt"
import go_parser "go/parser"
import proto1 "github.com/gogo/protobuf/proto"

// discarding unused import gogoproto "gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

func TestReplicaStateProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedReplicaState(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &ReplicaState{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestReplicaStateMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedReplicaState(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &ReplicaState{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkReplicaStateProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*ReplicaState, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedReplicaState(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := github_com_gogo_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkReplicaStateProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := github_com_gogo_protobuf_proto.Marshal(NewPopulatedReplicaState(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &ReplicaState{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_gogo_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestVerifierStateProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierState(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &VerifierState{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestVerifierStateMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierState(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &VerifierState{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkVerifierStateProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*VerifierState, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedVerifierState(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := github_com_gogo_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkVerifierStateProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := github_com_gogo_protobuf_proto.Marshal(NewPopulatedVerifierState(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &VerifierState{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_gogo_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestReplicaStateJSON(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedReplicaState(popr, true)
	marshaler := github_com_gogo_protobuf_jsonpb.Marshaller{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatal(err)
	}
	msg := &ReplicaState{}
	err = github_com_gogo_protobuf_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestVerifierStateJSON(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierState(popr, true)
	marshaler := github_com_gogo_protobuf_jsonpb.Marshaller{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatal(err)
	}
	msg := &VerifierState{}
	err = github_com_gogo_protobuf_jsonpb.UnmarshalString(jsondata, msg)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestReplicaStateProtoText(t *testing.T) {
	t.Skip()
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedReplicaState(popr, true)
	data := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &ReplicaState{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestReplicaStateProtoCompactText(t *testing.T) {
	t.Skip()
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedReplicaState(popr, true)
	data := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &ReplicaState{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestVerifierStateProtoText(t *testing.T) {
	t.Skip()
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierState(popr, true)
	data := github_com_gogo_protobuf_proto.MarshalTextString(p)
	msg := &VerifierState{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestVerifierStateProtoCompactText(t *testing.T) {
	t.Skip()
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierState(popr, true)
	data := github_com_gogo_protobuf_proto.CompactTextString(p)
	msg := &VerifierState{}
	if err := github_com_gogo_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestReplicaStateVerboseEqual(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedReplicaState(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &ReplicaState{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestVerifierStateVerboseEqual(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierState(popr, false)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &VerifierState{}
	if err := github_com_gogo_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestReplicaStateGoString(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedReplicaState(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestVerifierStateGoString(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierState(popr, false)
	s1 := p.GoString()
	s2 := fmt.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestReplicaStateSize(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedReplicaState(popr, true)
	size2 := github_com_gogo_protobuf_proto.Size(p)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Errorf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Errorf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := github_com_gogo_protobuf_proto.Size(p)
	if size3 != size {
		t.Errorf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkReplicaStateSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*ReplicaState, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedReplicaState(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestVerifierStateSize(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierState(popr, true)
	size2 := github_com_gogo_protobuf_proto.Size(p)
	data, err := github_com_gogo_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Errorf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Errorf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := github_com_gogo_protobuf_proto.Size(p)
	if size3 != size {
		t.Errorf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkVerifierStateSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*VerifierState, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedVerifierState(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestReplicaStateStringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedReplicaState(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestVerifierStateStringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierState(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}

//These tests are generated by github.com/gogo/protobuf/plugin/testgen
