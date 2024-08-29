import (
	"bytes"
	"encoding/binary"
	"errors"
)
var errInvalidTCPBBRInfo = errors.New("invalid TCP BBR info")

func parseInfo(b []byte) (tcpopt.Option, error) {
	if len(b) < sizeofTCPBBRInfo {
		return nil, errInvalidTCPBBRInfo
	}
	var info tcpBBRInfo
	buf := bytes.NewReader(b[:sizeofTCPBBRInfo])
	if err := binary.Read(buf, binary.LittleEndian, &info); err != nil {
		return nil, err
	}
	return &info, nil
}
