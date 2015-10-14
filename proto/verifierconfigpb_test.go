// Code generated by protoc-gen-gogo.
// source: verifierconfig.proto
// DO NOT EDIT!

package proto

import testing "testing"
import math_rand "math/rand"
import time "time"
import github_com_andres_erbsen_protobuf_proto "github.com/andres-erbsen/protobuf/proto"
import github_com_andres_erbsen_protobuf_jsonpb "github.com/andres-erbsen/protobuf/jsonpb"
import fmt "fmt"
import go_parser "go/parser"

// discarding unused import gogoproto "gogoproto"

func TestVerifierConfigProto(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierConfig(popr, false)
	data, err := github_com_andres_erbsen_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &VerifierConfig{}
	if err := github_com_andres_erbsen_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	littlefuzz := make([]byte, len(data))
	copy(littlefuzz, data)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
	if len(littlefuzz) > 0 {
		fuzzamount := 100
		for i := 0; i < fuzzamount; i++ {
			littlefuzz[popr.Intn(len(littlefuzz))] = byte(popr.Intn(256))
			littlefuzz = append(littlefuzz, byte(popr.Intn(256)))
		}
		// shouldn't panic
		_ = github_com_andres_erbsen_protobuf_proto.Unmarshal(littlefuzz, msg)
	}
}

func TestVerifierConfigMarshalTo(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierConfig(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &VerifierConfig{}
	if err := github_com_andres_erbsen_protobuf_proto.Unmarshal(data, msg); err != nil {
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

func BenchmarkVerifierConfigProtoMarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*VerifierConfig, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedVerifierConfig(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := github_com_andres_erbsen_protobuf_proto.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkVerifierConfigProtoUnmarshal(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := github_com_andres_erbsen_protobuf_proto.Marshal(NewPopulatedVerifierConfig(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &VerifierConfig{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := github_com_andres_erbsen_protobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestVerifierConfigJSON(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierConfig(popr, true)
	marshaler := github_com_andres_erbsen_protobuf_jsonpb.Marshaller{}
	jsondata, err := marshaler.MarshalToString(p)
	if err != nil {
		t.Fatal(err)
	}
	msg := &VerifierConfig{}
	err = github_com_andres_erbsen_protobuf_jsonpb.UnmarshalString(jsondata, msg)
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
func TestVerifierConfigProtoText(t *testing.T) {
	t.Skip()
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierConfig(popr, true)
	data := github_com_andres_erbsen_protobuf_proto.MarshalTextString(p)
	msg := &VerifierConfig{}
	if err := github_com_andres_erbsen_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestVerifierConfigProtoCompactText(t *testing.T) {
	t.Skip()
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierConfig(popr, true)
	data := github_com_andres_erbsen_protobuf_proto.CompactTextString(p)
	msg := &VerifierConfig{}
	if err := github_com_andres_erbsen_protobuf_proto.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestVerifierConfigStringer(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierConfig(popr, false)
	s1 := p.String()
	s2 := fmt.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestVerifierConfigSize(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierConfig(popr, true)
	size2 := github_com_andres_erbsen_protobuf_proto.Size(p)
	data, err := github_com_andres_erbsen_protobuf_proto.Marshal(p)
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
	size3 := github_com_andres_erbsen_protobuf_proto.Size(p)
	if size3 != size {
		t.Errorf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkVerifierConfigSize(b *testing.B) {
	popr := math_rand.New(math_rand.NewSource(616))
	total := 0
	pops := make([]*VerifierConfig, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedVerifierConfig(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestVerifierConfigGoString(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierConfig(popr, false)
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
func TestVerifierConfigVerboseEqual(t *testing.T) {
	popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
	p := NewPopulatedVerifierConfig(popr, false)
	data, err := github_com_andres_erbsen_protobuf_proto.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &VerifierConfig{}
	if err := github_com_andres_erbsen_protobuf_proto.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}

//These tests are generated by github.com/andres-erbsen/protobuf/plugin/testgen
