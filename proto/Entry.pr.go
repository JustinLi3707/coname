// Copyright 2014-2015 The Dename Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package proto

import "fmt"

type Entry_PreserveEncoding struct {
	Entry
	PreservedEncoding []byte `json:"-"`
}

func (m *Entry_PreserveEncoding) UpdateEncoding() (err error) {
	m.PreservedEncoding, err = m.Entry.Marshal()
	return err
}

func (m *Entry_PreserveEncoding) Reset() {
	*m = Entry_PreserveEncoding{}
}

func (m *Entry_PreserveEncoding) Size() int {
	return len(m.PreservedEncoding)
}

func (m *Entry_PreserveEncoding) Marshal() ([]byte, error) {
	size := m.Size()
	data := make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Entry_PreserveEncoding) MarshalTo(data []byte) (int, error) {
	return copy(data, m.PreservedEncoding), nil
}

func (m *Entry_PreserveEncoding) Unmarshal(data []byte) error {
	m.PreservedEncoding = append([]byte{}, data...)
	return m.Entry.Unmarshal(data)
}

func NewPopulatedEntry_PreserveEncoding(r randyClient, easy bool) *Entry_PreserveEncoding {
	this := &Entry_PreserveEncoding{Entry: *NewPopulatedEntry(r, easy)}
	this.UpdateEncoding()
	return this
}

func (this *Entry_PreserveEncoding) VerboseEqual(that interface{}) error {
	if thatP, ok := that.(*Entry_PreserveEncoding); ok {
		return this.Entry.VerboseEqual(&thatP.Entry)
	}
	if thatP, ok := that.(Entry_PreserveEncoding); ok {
		return this.Entry.VerboseEqual(&thatP.Entry)
	}
	return fmt.Errorf("types don't match: %T != Entry_PreserveEncoding")
}

func (this *Entry_PreserveEncoding) Equal(that interface{}) bool {
	if thatP, ok := that.(*Entry_PreserveEncoding); ok {
		return this.Entry.Equal(&thatP.Entry)
	}
	if thatP, ok := that.(Entry_PreserveEncoding); ok {
		return this.Entry.Equal(&thatP.Entry)
	}
	return false
}

func (this *Entry_PreserveEncoding) GoString() string {
	if this == nil {
		return "nil"
	}
	return `proto.Entry_PreserveEncoding{Entry: ` + this.Entry.GoString() + `, PreservedEncoding: ` + fmt.Sprintf("%#v", this.PreservedEncoding) + `}`
}

func (this *Entry_PreserveEncoding) String() string {
	if this == nil {
		return "nil"
	}
	return `proto.Entry_PreserveEncoding{Entry: ` + this.Entry.String() + `, PreservedEncoding: ` + fmt.Sprintf("%v", this.PreservedEncoding) + `}`
}
