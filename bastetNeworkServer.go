package wjbastetLibWSPackage

import (
	"encoding/binary"

	gatlingWSProtocol "github.com/gatlinglab/libGatlingWS/modProtocol"
)

type CWJBWSP_ServerParseData1 struct {
	PointerProxy  uint64
	PointerServer uint64
	RequestID     uint16
	CMD1          byte
	CMD2          byte
	CMD3          byte
}

type CWJBWSP_ServerParser1 struct {
	parseData CWJBWSP_ServerParseData1
	sock      gatlingWSProtocol.IWJSocket
}

func WJBWSP_CreateServerParser1(socket gatlingWSProtocol.IWJSocket) *CWJBWSP_ServerParser1 {
	return &CWJBWSP_ServerParser1{sock: socket}
}

// func (pInst *CWJBWSP_ServerParser1) SetProxyDataPointer(dataPoint uint64) {
// 	pInst.pointerProxy = dataPoint
// 	binary.LittleEndian.PutUint64(pInst.baseData, dataPoint)
// }
// func (pInst *CWJBWSP_ServerParser1) SetServerDataPointer(dataPoint uint64) {
// 	pInst.pointerServer = dataPoint
// 	binary.LittleEndian.PutUint64(pInst.baseData, dataPoint)
// }

func (pInst *CWJBWSP_ServerParser1) DataParse(data []byte, result *CWJBWSP_ServerParseData1) int {
	datalen := len(data)
	if datalen < dWJBP_ServerLengthBasicData {
		return -1
	}
	//copy(pInst.baseData, data[:dWJBP_ServerOffsetRequest])
	result.PointerProxy = binary.LittleEndian.Uint64(data)
	result.PointerServer = binary.LittleEndian.Uint64(data[8:])
	result.RequestID = uint16(data[dWJBP_ServerOffsetRequest])<<8 | uint16(data[dWJBP_ServerOffsetRequest+1])
	result.CMD1 = data[dWJBP_ServerOffsetCommand1]
	result.CMD2 = data[dWJBP_ServerOffsetCommand2]
	result.CMD3 = data[dWJBP_ServerOffsetCommand3]

	// the rest data is data[WJBP_LengthBasicData:]
	return dWJBP_ServerLengthBasicData
}
func (pInst *CWJBWSP_ServerParser1) DataParseDefault(data []byte) (*CWJBWSP_ServerParseData1, int) {
	iRet := pInst.DataParse(data, &pInst.parseData)

	return &pInst.parseData, iRet
	// the rest data is data[WJBP_LengthBasicData:]
}
func (pInst *CWJBWSP_ServerParser1) CommandSend(point1, point2 uint64, cmd1, cmd2, cmd3 byte, requestid uint16) error {
	_, err := pInst.DataSend(point1, point2, cmd1, cmd2, cmd3, requestid, nil)
	return err
}
func (pInst *CWJBWSP_ServerParser1) CommandSend2(point1, point2 uint64, cmd1, cmd2, cmd3 byte) error {
	_, err := pInst.DataSend(point1, point2, cmd1, cmd2, cmd3, 0, nil)
	return err
}
func (pInst *CWJBWSP_ServerParser1) CommandSend3(cmd3 byte, parseData *CWJBWSP_ServerParseData1) error {
	_, err := pInst.DataSend(parseData.PointerProxy, parseData.PointerServer, parseData.CMD1, parseData.CMD2, cmd3, parseData.RequestID, nil)
	return err
}

func (pInst *CWJBWSP_ServerParser1) DataSend(point1, point2 uint64, cmd1, cmd2, cmd3 byte, requestid uint16, data []byte) (int, error) {
	datalen := 0
	if data != nil {
		datalen = len(data)
	}
	dataSend := make([]byte, dWJBP_ServerLengthBasicData+datalen)
	binary.LittleEndian.PutUint64(dataSend, point1)
	//data1 := make([]byte, 8)
	binary.LittleEndian.PutUint64(dataSend[8:], point2)

	if requestid != 0 {
		dataSend[dWJBP_ServerOffsetRequest] = byte(requestid >> 8)
		dataSend[dWJBP_ServerOffsetRequest+1] = byte(requestid)
	}
	dataSend[dWJBP_ServerOffsetCommand1] = cmd1
	dataSend[dWJBP_ServerOffsetCommand2] = cmd2
	dataSend[dWJBP_ServerOffsetCommand3] = cmd3

	if data != nil {
		copy(dataSend[dWJBP_ServerLengthBasicData:], data)
	}

	err := pInst.sock.WriteBinary(dataSend)

	return datalen, err
}
func (pInst *CWJBWSP_ServerParser1) DataSend2(point1, point2 uint64, cmd1, cmd2, cmd3 byte, data []byte) (int, error) {
	return pInst.DataSend(point1, point2, cmd1, cmd2, cmd3, 0, data)
}
func (pInst *CWJBWSP_ServerParser1) DataSend3(cmd3 byte, data []byte, parseData *CWJBWSP_ServerParseData1) (int, error) {
	return pInst.DataSend(parseData.PointerProxy, parseData.PointerServer, parseData.CMD1, parseData.CMD2, cmd3, parseData.RequestID, data)
}
