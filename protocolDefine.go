package wjbastetLibWSPackage

type WJBP_RequestType uint16

// byte 1 and 2 used for requestid; int16; it will reserved even don't need requestid,
const dWJBP_LengthBasicData = 5 // 2 byte for requestid, 3 byte for command
const dWJBP_OffsetRequest = 0
const dWJBP_OffsetCommand1 = 2 // 3
const dWJBP_OffsetCommand2 = 3 // 4
const dWJBP_OffsetCommand3 = 4 // 4

const dWJBP_ServerLengthBasicData = 21 // 8 bytes for data pointer proxy, 8 byte for pointer server, 2 byte for requestid, 3 byte for command
const dWJBP_ServerOffsetRequest = 16
const dWJBP_ServerOffsetCommand1 = 18 // cmd 1
const dWJBP_ServerOffsetCommand2 = 19 // cmd 2
const dWJBP_ServerOffsetCommand3 = 20 // cmd 3
