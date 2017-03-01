package page_analyzer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/YoungZhou93/startGolang/web/crawler/common/page"
	"github.com/guotie/gogb2312"
	"github.com/tealeg/xlsx"
	"strconv"
	"strings"
)

type SaveAnalyer struct {
	results []string
}

func NewSaveAnalyer() *SaveAnalyer {
	return &SaveAnalyer{}
}

func (this *SaveAnalyer) Analyze(page *page.Page) {
	fmt.Println("开始分析")
	number := page.GetRequest().Label
	fmt.Println(number)
	quety := page.GetHtmlParser()
	title := ""
	name:=quety.Find(".post-title").Text();
	nameByte := []byte(name)
	outputname, _, _, _ := gogb2312.ConvertGB2312(nameByte)
	name=string(outputname)
	zz := ""
	jc := ""
	//临床表现
	quety.Find(".jb-xx-zz li").Each(func(i int, contentSelection *goquery.Selection) {
		title = contentSelection.Find("a").Text()
		titleByte := []byte(title)
		output, _, _, _ := gogb2312.ConvertGB2312(titleByte)
		if zz == "" {
			zz = zz + string(output)
		} else {
			zz = zz + "、" + string(output)
		}
	})
	//检查手段
	quety.Find(".jb-xx-jc li").Each(func(i int, contentSelection *goquery.Selection) {
		title = contentSelection.Find("a").Text()
		titleByte := []byte(title)
		output, _, _, _ := gogb2312.ConvertGB2312(titleByte)
		if jc == "" {
			jc = jc + string(output)
		} else {
			jc = jc + "、" + string(output)
		}
	})

	if zz=="" && jc==""{
		quety.Find(".zz-xx-zz li").Each(func(i int, contentSelection *goquery.Selection) {
			title = contentSelection.Find("a").Text()
			titleByte := []byte(title)
			output, _, _, _ := gogb2312.ConvertGB2312(titleByte)

			if zz == "" {
				zz = zz + string(output)
			} else {
				zz = zz + "、" + string(output)
			}

		})

		quety.Find(".zz-xx-jc li").Each(func(i int, contentSelection *goquery.Selection) {
			title = contentSelection.Find("a").Text()
			titleByte := []byte(title)
			output, _, _, _ := gogb2312.ConvertGB2312(titleByte)
			if jc == "" {
				jc = jc + string(output)
			} else {
				jc = jc + "、" + string(output)
			}

		})
	}
	result := zz + "#" + jc+"#"+name
	fmt.Println(result)
	this.results = append(this.results, number+"#"+result)

	fmt.Println("分析结束")
}

func (this *SaveAnalyer) AnalyzeFinsh() {
	count:=0;
	excelFileName := "d:\\数据.xlsx"
	xlFile, error := xlsx.OpenFile(excelFileName)
	if error != nil {
	}
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			length:=400
			if i >= 400 && i<500{
					fmt.Print("第")
					fmt.Print(i)
					fmt.Println("项")
					curr:=""
					if i-length>=len(this.results){
						curr=this.results[len(this.results)-1]
					}else {
						curr = this.results[i - length]
					}
					if len(curr)> 4 {
						fmt.Println(curr)
						strs := strings.Split(curr, "#")
						fmt.Print(i)
						fmt.Print("*")
						fmt.Println(strs[0])
						if strconv.Itoa(i) == strs[0] {
							fmt.Print(row.Cells[1].Value)
							fmt.Print("*")
							fmt.Println(strs[3])
							if strs[3]==row.Cells[1].Value {
								count++;
								if row.Cells[2].Value == "" || len(row.Cells[2].Value) == 0 {
									row.Cells[2].Value = strs[1]
								}
								if row.Cells[3].Value == "" || len(row.Cells[3].Value) == 0 {
									row.Cells[3].Value = strs[2]
								}
							}
						}
						if strconv.Itoa(i) != strs[0] {
							curr2:=""
							for j:=i;j>length-1;j--{
								if j-length>=len(this.results){
									curr2=this.results[len(this.results)-1]
								}else {
									curr2 = this.results[j - length]
								}

								if len(curr2)> 4 {
									strs2 := strings.Split(curr2, "#")
									fmt.Print(i)
									fmt.Print("*")
									fmt.Println(strs2[0])
									if strconv.Itoa(i) == strs2[0] {
										if strs[3]==row.Cells[1].Value {
											count++;
											if row.Cells[2].Value == "" || len(row.Cells[2].Value) == 0 {
												row.Cells[2].Value = strs2[1]
											}
											if row.Cells[3].Value == "" || len(row.Cells[3].Value) == 0 {
												row.Cells[3].Value = strs2[2]
											}
										}
										break
									}
								}

							}


						}
					}

			}
		}
	}
	xlFile.Save(excelFileName)
	fmt.Println(count);
}
