package wjbastetLibWSPackage

type WJBP_RequestType uint16

// byte 1 and 2 used for requestid; int16; it will reserved even don't need requestid,
const dWJBP_LengthBasicData = 5 // 2 byte for requestid, 3 byte for command
const dWJBP_OffsetRequest = 0
const dWJBP_OffsetCommand1 = 2 // 3
const dWJBP_OffsetCommand2 = 3 // 4
const dWJBP_OffsetCommand3 = 4 // 4

const dWJBP_ServerLengthBasicData = 13 // 8 bytes for data pointer 2 byte for requestid, 3 byte for command
const dWJBP_ServerOffsetRequest = 8
const dWJBP_ServerOffsetCommand1 = 10 // cmd 1
const dWJBP_ServerOffsetCommand2 = 11 // cmd 2
const dWJBP_ServerOffsetCommand3 = 12 // cmd 3
