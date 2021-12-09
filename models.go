package main

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
