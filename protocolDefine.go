package wjbastetLibWSPackage

type WJBP_RequestType uint16

// byte 1 and 2 used for requestid; int16; it will reserved even don't need requestid,
const WJBP_LengthBasicData = 5 // 2 byte for requestid, 3 byte for command
const WJBP_OffsetCommand1 = 2  // 3
const WJBP_OffsetCommand2 = 3  // 4
const WJBP_OffsetCommand3 = 4  // 4



