# golang字节数据到结构体转换的快捷方法


## 特性

- 基于tag结构体标签
- 支持嵌套结构体

## 版权声明

**绝对原创**，如有使用，请加star，或评论

## 例子

```
// -------------------- 结构体定义 ---------------
type UK_Version struct {
	Major uint8 `ctype:"order=1,size=1"`
	Minor uint8 `ctype:"order=2,size=1"`
}

type UK_CosDEVINFO struct {
	Version          UK_Version `ctype:"order=1,size=2"`
	Manufacturer     string     `ctype:"order=2,size=64"`
	Issuer           string     `ctype:"order=3,size=64"`
	Label            string     `ctype:"order=4,size=32"`
	SerialNumber     string     `ctype:"order=5,size=32"`
	HWVersion        UK_Version `ctype:"order=6,size=2"`
	FirmwareVersion  UK_Version `ctype:"order=7,size=2"`
	AlgSymCap        uint32     `ctype:"order=8,size=4"`
	AlgAsymCap       uint32     `ctype:"order=9,size=4"`
	AlgHashCap       uint32     `ctype:"order=10,size=4"`
	DevAuthAlgId     uint32     `ctype:"order=11,size=4"`
	TotalSpace       uint32     `ctype:"order=12,size=4"`
	FreeSpace        uint32     `ctype:"order=13,size=4"`
	MaxECCBufferSize uint32     `ctype:"order=14,size=4"`
	MaxBufferSize    uint32     `ctype:"order=15,size=4"`
}

# ----------- 测试结果 ---------------
❯ go build && ./cstructparse
2021/12/09 23:22:50 字节数组: 010053696e6f43697068657200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000053696e6f4369706865720000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007a61796b3030310000000000000000000000000000000000000000000000000030303138303931373441384534353338353533353333333100000000000000000100010003170000000107000f000000010700000000020000f70100000800000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030000000a04f8f77fc7f
2021/12/09 23:22:50 结构体: main.UK_CosDEVINFO{Version:main.UK_Version{Major:0x1, Minor:0x0}, Manufacturer:"SinoCipher", Issuer:"SinoCipher", Label:"zayk001", SerialNumber:"001809174A8E453855353331", HWVersion:main.UK_Version{Major:0x1, Minor:0x0}, FirmwareVersion:main.UK_Version{Major:0x1, Minor:0x0}, AlgSymCap:0x1703, AlgAsymCap:0x70100, AlgHashCap:0xf, DevAuthAlgId:0x701, TotalSpace:0x20000, FreeSpace:0x1f700, MaxECCBufferSize:0x800, MaxBufferSize:0x1000}
```
