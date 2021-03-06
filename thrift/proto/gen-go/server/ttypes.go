// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package server

import (
	"bytes"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - Field1
//  - Field2
//  - Field3
//  - Field4
//  - Field5
type Msg struct {
	Field1 string `thrift:"field1,1" json:"field1"`
	Field2 string `thrift:"field2,2" json:"field2"`
	Field3 string `thrift:"field3,3" json:"field3"`
	Field4 string `thrift:"field4,4" json:"field4"`
	Field5 string `thrift:"field5,5" json:"field5"`
}

func NewMsg() *Msg {
	return &Msg{}
}

func (p *Msg) GetField1() string {
	return p.Field1
}

func (p *Msg) GetField2() string {
	return p.Field2
}

func (p *Msg) GetField3() string {
	return p.Field3
}

func (p *Msg) GetField4() string {
	return p.Field4
}

func (p *Msg) GetField5() string {
	return p.Field5
}
func (p *Msg) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Msg) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Field1 = v
	}
	return nil
}

func (p *Msg) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Field2 = v
	}
	return nil
}

func (p *Msg) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Field3 = v
	}
	return nil
}

func (p *Msg) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Field4 = v
	}
	return nil
}

func (p *Msg) readField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.Field5 = v
	}
	return nil
}

func (p *Msg) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Msg"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Msg) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("field1", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:field1: ", p), err)
	}
	if err := oprot.WriteString(string(p.Field1)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.field1 (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:field1: ", p), err)
	}
	return err
}

func (p *Msg) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("field2", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:field2: ", p), err)
	}
	if err := oprot.WriteString(string(p.Field2)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.field2 (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:field2: ", p), err)
	}
	return err
}

func (p *Msg) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("field3", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:field3: ", p), err)
	}
	if err := oprot.WriteString(string(p.Field3)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.field3 (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:field3: ", p), err)
	}
	return err
}

func (p *Msg) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("field4", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:field4: ", p), err)
	}
	if err := oprot.WriteString(string(p.Field4)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.field4 (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:field4: ", p), err)
	}
	return err
}

func (p *Msg) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("field5", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:field5: ", p), err)
	}
	if err := oprot.WriteString(string(p.Field5)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.field5 (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:field5: ", p), err)
	}
	return err
}

func (p *Msg) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Msg(%+v)", *p)
}
