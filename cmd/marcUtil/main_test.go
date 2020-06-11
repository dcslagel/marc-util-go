package main

// https://gobyexample.com/command-line-flags
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestLeader(t *testing.T) {
	file, err := os.Open("../../examples/loc-header-only.mrc")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mh := marcheader{}
	want := marcheader{}
	want.recordlength = "*****"
	want.recordstatus = "n"
	want.recordtype = "a"
	want.biblevel = "m"
	want.controltype = "#"
	want.characterencoding = "#"
	want.indicatorcount = "2"
	want.subfieldcodecount = "2"
	want.dataaddress = "*****"
	want.encodinglevel = "#"
	want.catalogingform = "a"
	want.multipartresourcereclevel = "#"
	want.lengthoffield = "4"
	want.lengthofstartcharpos = "5"
	want.lengthofimplementationportion = "0"
	want.undefined = "0"

	for scanner.Scan() {
		fmt.Println(len(scanner.Text()))
		headerstring := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal("Shouldn't see an error scanning a string")
		}

		mh.Parser(headerstring)

		if mh.recordlength != want.recordlength {
			t.Errorf("mh.recordlength = %q recordlength %q", mh.recordlength, want.recordlength)
		}
		if mh.recordstatus != want.recordstatus {
			t.Errorf("mh.recordstatus = %q recordstatus %q", mh.recordstatus, want.recordstatus)
		}
		if mh.recordtype != want.recordtype {
			t.Errorf("mh.recordtype = %q recordtype %q", mh.recordtype, want.recordtype)
		}
		if mh.biblevel != want.biblevel {
			t.Errorf("mh.biblevel = %q biblevel %q", mh.biblevel, want.biblevel)
		}
		if mh.controltype != want.controltype {
			t.Errorf("mh.controltype = %q controltype %q", mh.controltype, want.controltype)
		}
		if mh.characterencoding != want.characterencoding {
			t.Errorf("mh.characterencoding = %q characterencoding %q", mh.characterencoding, want.characterencoding)
		}
		if mh.indicatorcount != want.indicatorcount {
			t.Errorf("mh.indicatorcount = %q indicatorcount %q", mh.indicatorcount, want.indicatorcount)
		}
		if mh.subfieldcodecount != want.subfieldcodecount {
			t.Errorf("mh.subfieldcodecount = %q subfieldcodecount %q", mh.subfieldcodecount, want.subfieldcodecount)
		}
		if mh.dataaddress != want.dataaddress {
			t.Errorf("mh.dataaddress = %q dataaddress %q", mh.dataaddress, want.dataaddress)
		}
		if mh.encodinglevel != want.encodinglevel {
			t.Errorf("mh.encodinglevel = %q encodinglevel %q", mh.encodinglevel, want.encodinglevel)
		}
		if mh.catalogingform != want.catalogingform {
			t.Errorf("mh.catalogingform = %q catalogingform %q", mh.catalogingform, want.catalogingform)
		}
		if mh.multipartresourcereclevel != want.multipartresourcereclevel {
			t.Errorf("mh.multipartresourcereclevel = %q multipartresourcereclevel %q", mh.multipartresourcereclevel, want.multipartresourcereclevel)
		}
		if mh.lengthoffield != want.lengthoffield {
			t.Errorf("mh.lengthoffield = %q lengthoffield %q", mh.lengthoffield, want.lengthoffield)
		}
		if mh.lengthofstartcharpos != want.lengthofstartcharpos {
			t.Errorf("mh.lengthofstartcharpos = %q lengthofstartcharpos %q", mh.lengthofstartcharpos, want.lengthofstartcharpos)
		}
		if mh.lengthofimplementationportion != want.lengthofimplementationportion {
			t.Errorf("mh.lengthofimplementationportion = %q lengthofimplementationportion %q", mh.lengthofimplementationportion, want.lengthofimplementationportion)
		}
		if mh.undefined != want.undefined {
			t.Errorf("mh.undefined = %q undefined %q", mh.undefined, want.undefined)
		}
	}
}
