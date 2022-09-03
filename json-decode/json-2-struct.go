package json_decode

//
//import (
//	"encoding/json"
//	"fmt"
//	"os"
//)
//
//type StructName []struct {
//	Cddkfs   string     `json:"cddkfs"`
//	Children []Children `json:"children"`
//	ID       string     `json:"id"`
//	Text     string     `json:"text"`
//	URL      string     `json:"url"`
//	Cdsffz   string     `json:"cdsffz"`
//}
//type Children1 struct {
//	Cddkfs  string `json:"cddkfs"`
//	ID      string `json:"id"`
//	Text    string `json:"text"`
//	IconCls string `json:"iconCls"`
//	URL     string `json:"url"`
//	Cdsffz  string `json:"cdsffz"`
//}
//type Children struct {
//	Cddkfs   string      `json:"cddkfs"`
//	ID       string      `json:"id"`
//	Text     string      `json:"text"`
//	IconCls  string      `json:"iconCls,omitempty"`
//	URL      string      `json:"url"`
//	Cdsffz   string      `json:"cdsffz"`
//	Children []Children1 `json:"children,omitempty"`
//}
//
//func main() {
//	var info StructName
//	file, _ := os.Open("C:\\Users\\xx\\Desktop\\1.txt")
//	defer file.Close()
//	encoder := json.NewDecoder(file)
//	err := encoder.Decode(&info)
//	if err != nil {
//		fmt.Println("解码失败", err.Error())
//	}
//	for _, s := range info {
//		if len(s.Children) != 0 {
//			for _, child := range s.Children {
//				if len(child.Children) != 0 {
//					for _, children1 := range child.Children {
//						fmt.Println(s.Text + "--" + child.Text + "--" + children1.Text)
//					}
//				} else {
//					fmt.Println(s.Text + "--" + child.Text)
//				}
//			}
//		} else {
//			fmt.Println(s.Text)
//		}
//		fmt.Println(s.Text)
//	}
//
//}
