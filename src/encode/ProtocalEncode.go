package encode

import (
	"encoding/binary"
	"bytes"
)

type Head struct  {
	Cmd byte
	Version byte
	BodyLen uint32
}

type Package struct  {
	Head Head
	Data []byte
}

func GetPackage(cmd, version uint16, data []byte) []byte {
	dataPackage := new(Package)
	dataPackage.Head.Cmd = resolveIntToByte(cmd)
	dataPackage.Head.Version = resolveIntToByte(version)
	dataPackage.Head.BodyLen = uint32(len(data))
	dataPackage.Data = data
	result := resolvePackageToByte(dataPackage)
	return result
}

func resolveIntToByte(value uint16) byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, value)
	return buf.Bytes()[0]
}

func resolvePackageToByte(pac *Package) []byte{
	var buf bytes.Buffer
	buf.WriteByte(pac.Head.Cmd)
	buf.WriteByte(pac.Head.Version)
	f := new(bytes.Buffer)
	binary.Write(f, binary.LittleEndian, pac.Head.BodyLen)
	buf.Write(buf.Bytes()[0:4])
	buf.Write(pac.Data)
	return buf.Bytes()
}
