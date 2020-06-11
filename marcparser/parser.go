package marcparser

import (
	"fmt"
)

// var mheader = map[string]string

// https://www.loc.gov/marc/umb/um07to10.html#part9
type marcheader struct {
	recordlength                  string
	recordstatus                  string
	recordtype                    string
	biblevel                      string
	controltype                   string
	characterencoding             string
	indicatorcount                string
	subfieldcodecount             string
	dataaddress                   string
	encodinglevel                 string
	catalogingform                string
	multipartresourcereclevel     string
	lengthoffield                 string
	lengthofstartcharpos          string
	lengthofimplementationportion string
	undefined                     string
}

func Init() marcheader {
	mh := marcheader{}
	return mh
}

func (mheader *marcheader) Parser(headerstring string) {

	mheader.recordlength = headerstring[0:5]
	mheader.recordstatus = headerstring[5:6]
	mheader.recordtype = headerstring[6:7]
	mheader.biblevel = headerstring[7:8]
	mheader.controltype = headerstring[8:9]
	mheader.characterencoding = headerstring[9:10]
	mheader.indicatorcount = headerstring[10:11]
	mheader.subfieldcodecount = headerstring[11:12]
	mheader.dataaddress = headerstring[12:17]
	mheader.encodinglevel = headerstring[17:18]
	mheader.catalogingform = headerstring[18:19]
	mheader.multipartresourcereclevel = headerstring[19:20]
	mheader.lengthoffield = headerstring[20:21]
	mheader.lengthofstartcharpos = headerstring[21:22]
	mheader.lengthofimplementationportion = headerstring[22:23]
	mheader.undefined = headerstring[23:24]

}

func (mheader *marcheader) PrintHeader() {
	fmt.Println(mheader.recordlength)
	fmt.Println(mheader.recordstatus)
	fmt.Println(mheader.recordtype)
	fmt.Println(mheader.biblevel)
	fmt.Println(mheader.controltype)
	fmt.Println(mheader.characterencoding)
	fmt.Println(mheader.indicatorcount)
	fmt.Println(mheader.subfieldcodecount)
	fmt.Println(mheader.dataaddress)
	fmt.Println(mheader.encodinglevel)
	fmt.Println(mheader.catalogingform)
	fmt.Println(mheader.multipartresourcereclevel)
	fmt.Println(mheader.lengthoffield)
	fmt.Println(mheader.lengthofstartcharpos)
	fmt.Println(mheader.lengthofimplementationportion)
	fmt.Println(mheader.undefined)
}
