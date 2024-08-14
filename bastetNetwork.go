package wjbastetLibWSPackage

import (
	gatlingWSProtocol "github.com/gatlinglab/libGatlingWS/modProtocol"
)

type CWJBWSP_ParseData1 struct {
	RequestID uint16
	CMD1      byte
	CMD2      byte
	CMD3      byte
}

type CWJBWSP_Parser1 struct {
	parseData CWJBWSP_ParseData1
	sock      gatlingWSProtocol.IWJSocket
}

func WJBWSP_CreateParser1(socket gatlingWSProtocol.IWJSocket) *CWJBWSP_Parser1 {
	return &CWJBWSP_Parser1{sock: socket}
}

func (pInst *CWJBWSP_Parser1) DataParse(data []byte, result *CWJBWSP_ParseData1) int {
	datalen := len(data)
	if datalen < dWJBP_LengthBasicData {
		return -1
	}
	result.RequestID = uint16(data[dWJBP_OffsetRequest])<<8 | uint16(data[dWJBP_OffsetRequest+1])
	result.CMD1 = data[dWJBP_OffsetCommand1]
	result.CMD2 = data[dWJBP_OffsetCommand2]
	result.CMD3 = data[dWJBP_OffsetCommand3]

	// the rest data is data[WJBP_LengthBasicData:]
	return dWJBP_LengthBasicData
}
func (pInst *CWJBWSP_Parser1) DataParseDefault(data []byte) (*CWJBWSP_ParseData1, int) {
	iRet := pInst.DataParse(data, &pInst.parseData)

	return &pInst.parseData, iRet
	// the rest data is data[WJBP_LengthBasicData:]
}
func (pInst *CWJBWSP_Parser1) CommandSend(cmd1, cmd2, cmd3 byte, requestid uint16) error {
	_, err := pInst.DataSend(cmd1, cmd2, cmd3, requestid, nil)
	return err
}
func (pInst *CWJBWSP_Parser1) CommandSend2(cmd1, cmd2, cmd3 byte) error {
	_, err := pInst.DataSend(cmd1, cmd2, cmd3, 0, nil)
	return err
}
func (pInst *CWJBWSP_Parser1) CommandSend3(cmd3 byte, parseData *CWJBWSP_ParseData1) error {
	_, err := pInst.DataSend(parseData.CMD1, parseData.CMD2, cmd3, parseData.RequestID, nil)
	return err
}

func (pInst *CWJBWSP_Parser1) DataSend(cmd1, cmd2, cmd3 byte, requestid uint16, data []byte) (int, error) {
	datalen := 0
	if data != nil {
		datalen = len(data)
	}
	dataSend := make([]byte, dWJBP_LengthBasicData+datalen)
	if requestid != 0 {
		dataSend[0] = byte(requestid >> 8)
		dataSend[1] = byte(requestid)
	}
	dataSend[dWJBP_OffsetCommand1] = cmd1
	dataSend[dWJBP_OffsetCommand2] = cmd2
	dataSend[dWJBP_OffsetCommand3] = cmd3

	if data != nil {
		copy(dataSend[dWJBP_LengthBasicData:], data)
	}

	err := pInst.sock.WriteBinary(dataSend)

	return datalen, err
}
func (pInst *CWJBWSP_Parser1) DataSend2(cmd1, cmd2, cmd3 byte, data []byte) (int, error) {
	return pInst.DataSend(cmd1, cmd2, cmd3, 0, data)
}
func (pInst *CWJBWSP_Parser1) DataSend3(cmd3 byte, data []byte, parseData *CWJBWSP_ParseData1) (int, error) {
	return pInst.DataSend(parseData.CMD1, parseData.CMD2, cmd3, parseData.RequestID, data)
}
