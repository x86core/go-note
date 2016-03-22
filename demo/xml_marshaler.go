package main


import(
	"fmt"
	"encoding/xml"
	"bytes"
	"bufio"
	
)

type StringMap map[string]string

func main(){
	var enc *xml.Encoder
	
	
	var a = make(StringMap)
	
	a["a"] = "a"
	a["b"] = "b"
	a["c"] = "c"
	
	var startEle = xml.StartElement{
		Name: xml.Name{
			"start space",
			"start local",
		},
		Attr: []xml.Attr{
			xml.Attr{
				Name:xml.Name{
					"attr space",
					"attr local",
				},
				Value:"attr str",
			},
		},
	}
	
	b := bytes.NewBuffer(make([]byte,0))
	
	bw := bufio.NewWriter(b)
	enc = xml.NewEncoder(bw)
	
	a.MapToXml(enc ,startEle)
	
	
	fmt.Println(b)
	
}


func (s StringMap)MapToXml( e *xml.Encoder, start xml.StartElement) error {

	tokens := []xml.Token{start}
	
	for key, value := range s {
		t := xml.StartElement{Name:xml.Name{"", key}}
		tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
	}
	
	tokens = append(tokens, xml.EndElement{start.Name})
	
	
	for _, t := range tokens {
		err := e.EncodeToken(t)
		
		if err != nil {
			return err
		}
	}
	
	err := e.Flush()
	
	if err != nil {
		return err
	}
	
	return nil
	
}