// 将 jPOS 的配置文件转换成 goPOS 的配置文件
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

type Field struct {
	ID     int    `xml:"id,attr,omitempty"  json:"id"`
	Length int    `xml:"length,attr,omitempty"  json:"length"`
	Name   string `xml:"name,attr,omitempty"  json:"name"`
	Class  string `xml:"class,attr,omitempty" json:"type"`
	Pad    bool   `xml:"pad,attr,omitempty" json:"pad"`
}

type ISOPackager struct {
	Fields []*Field `xml:"isofield" json:"fields"`
}

// 修改 jpos Class 定义为 goPOS 的定义
func (isop *ISOPackager) fixClass() {
	for _, f := range isop.Fields {
		f.Class = strings.Replace(f.Class, "org.jpos.iso.I", "", -1)
	}
}

func main() {
	xb, err := ioutil.ReadFile("iso93binary.xml")
	if err != nil {
		panic(err)
	}
	j := &ISOPackager{}
	xml.Unmarshal(xb, j)
	j.fixClass()
	b, _ := json.MarshalIndent(j, "", "  ")
	fmt.Println(string(b))
}
