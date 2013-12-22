package skiplist

import (
	"bufio"
	"encoding/binary"
	"io"
	"log"
)

type RecordPersister interface {
	Persist(m *Map, f io.Writer) error
	Merge(m *Map, f io.Reader) error
}

type StringStringRecord struct {
	Key string
	Val string
}

type Int64Int64Record struct {
	Key int64
	Val int64
}

func (m *Map) Persist(f io.Writer, rp RecordPersister) error {
	rp.Persist(m, f)
	return nil
}

func (m *Map) Merge(f io.Reader, rp RecordPersister) error {
	rp.Merge(m, f)
	return nil
}

func (ss StringStringRecord) Persist(m *Map, f io.Writer) error {
	e := m.head[0]
	buf := bufio.NewWriter(f)
	defer buf.Flush()
	for e != nil {
		ss.Key = e.key.(string)
		ss.Val = e.val.(string)
		keyBytes := []byte(ss.Key)
		var keySize uint16 = uint16(len(keyBytes))
		err := binary.Write(buf, binary.LittleEndian, keySize)
		if err != nil {
			log.Println("binary.Write failed:", err)
		}
		err = binary.Write(buf, binary.LittleEndian, keyBytes)
		if err != nil {
			log.Println("binary.Write failed:", err)
		}
		valBytes := []byte(ss.Val)
		var valSize uint16 = uint16(len(valBytes))
		err = binary.Write(buf, binary.LittleEndian, valSize)
		if err != nil {
			log.Println("binary.Write failed:", err)
		}
		err = binary.Write(buf, binary.LittleEndian, valBytes)
		if err != nil {
			log.Println("binary.Write failed:", err)
		}
		e = e.next[0]
	}
	return nil
}

func (ss StringStringRecord) Merge(m *Map, f io.Reader) error {
	buf := bufio.NewReader(f)
	var err error = nil
	for err == nil {
		var size uint16
		err = binary.Read(buf, binary.LittleEndian, &size)
		if err != nil {
			break
		}
		b := make([]byte, size)
		binary.Read(buf, binary.LittleEndian, &b)
		ss.Key = string(b)
		binary.Read(buf, binary.LittleEndian, &size)
		b = make([]byte, size)
		err = binary.Read(buf, binary.LittleEndian, &b)
		ss.Val = string(b)
		m.Put(ss.Key, ss.Val)
	}
	return nil
}

func (ss Int64Int64Record) Persist(m *Map, f io.Writer) error {
	e := m.head[0]
	buf := bufio.NewWriter(f)
	defer buf.Flush()
	for e != nil {
		ss.Key = e.key.(int64)
		ss.Val = e.val.(int64)
		err := binary.Write(buf, binary.LittleEndian, ss)
		if err != nil {
			return err
		}
		e = e.next[0]
	}
	return nil
}

func (ss Int64Int64Record) Merge(m *Map, f io.Reader) error {
	buf := bufio.NewReader(f)
	var err error = nil
	for err == nil {
		err = binary.Read(buf, binary.LittleEndian, &ss)
		m.Put(ss.Key, ss.Val)
	}
	return nil
}
