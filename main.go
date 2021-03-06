package main

import (
	"excel-maker/model"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/tealeg/xlsx"
)

type MyMainWindow struct {
	*walk.MainWindow
}

var idMap map[int]string

func initIdMap() {
	idMap = make(map[int]string)
	idMap[1] = "one"
	idMap[2] = "two"
	idMap[3] = "three"
	idMap[4] = "four"
	idMap[5] = "five"
	idMap[6] = "six"
	idMap[7] = "seven"
	idMap[8] = "eight"
	idMap[9] = "nine"
	idMap[10] = "ten"
	idMap[11] = "eleven"
	idMap[12] = "twelve"
	idMap[13] = "thirteen"
	idMap[14] = "fourteen"
	idMap[15] = "fifteen"
	idMap[16] = "sixteen"
	idMap[17] = "seveteen"
	idMap[18] = "eighteen"
}

func main() {

	mw := new(MyMainWindow)
	initIdMap()
	var db *walk.DataBinder
	invoice := new(model.Invoice)
	var widgetlit []Widget
	var widgetlit2 []Widget
	Hsnolist := make([]*walk.LineEdit, 18)
	ItemNolist := make([]*walk.LineEdit, 18)
	Deslist := make([]*walk.LineEdit, 18)
	Quantitylist := make([]*walk.LineEdit, 18)
	UnitPricelist := make([]*walk.LineEdit, 18)
	Netlist := make([]*walk.LineEdit, 18)

	Namelist := make([]*walk.LineEdit, 18)
	Standardlist := make([]*walk.LineEdit, 18)
	Codelist := make([]*walk.LineEdit, 18)
	Weightlist := make([]*walk.LineEdit, 18)
	Cartonslist := make([]*walk.LineEdit, 18)
	Palletlist := make([]*walk.LineEdit, 18)

	widgetlit = append(widgetlit, Label{
		Text: "商品海关编号(HsNo)",
	},
		Label{
			Text: "项目号(ItemNo)",
		},
		Label{
			Text: "商品描述(Description)",
		},
		Label{
			Text: "商品重量(Quantity)",
		},
		Label{
			Text: "净重(NET WT)",
		},
		Label{
			Text: "单价(UnitPrice)",
		},
	)
	widgetlit2 = append(widgetlit2, Label{
		Text: "名称(productName)",
	},
		Label{
			Text: "规格(Standard)",
		},
		Label{
			Text: "材料代号(Code)",
		},
		Label{
			Text: "重量(weight)",
		},
		Label{
			Text: "总箱数(TotalCartons)",
		},
		Label{
			Text: "编号(Pallet Number)",
		},
	)
	editlist := make([]LineEdit, 108)
	editlist1 := make([]LineEdit, 108)
	for i, edit := range editlist {
		switch i % 6 {
		case 0:
			edit.AssignTo = &Hsnolist[i/6]
		case 1:
			edit.AssignTo = &ItemNolist[i/6]
		case 2:
			edit.AssignTo = &Deslist[i/6]
		case 3:
			edit.AssignTo = &Quantitylist[i/6]
		case 4:
			edit.AssignTo = &Netlist[i/6]
		case 5:
			edit.AssignTo = &UnitPricelist[i/6]
		}
		widgetlit = append(widgetlit, edit)
	}
	for i, edit := range editlist1 {
		switch i % 6 {
		case 0:
			edit.AssignTo = &Namelist[i/6]
		case 1:
			edit.AssignTo = &Standardlist[i/6]
		case 2:
			edit.AssignTo = &Codelist[i/6]
		case 3:
			edit.AssignTo = &Weightlist[i/6]
		case 4:
			edit.AssignTo = &Cartonslist[i/6]
		case 5:
			edit.AssignTo = &Palletlist[i/6]
		}
		widgetlit2 = append(widgetlit2, edit)
	}
	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "太仓巨仁光伏材料有限公司",
		MinSize:  Size{1800, 600},
		Layout:   VBox{},
		Children: []Widget{
			HSplitter{
				HandleWidth: 5,
				DataBinder: DataBinder{
					AssignTo:       &db,
					DataSource:     invoice,
					ErrorPresenter: ToolTipErrorPresenter{},
				},
				Children: []Widget{
					Composite{
						Layout: Grid{Columns: 2},
						Children: []Widget{
							Label{
								MaxSize: Size{10, 10},
								Text:    "发票号(INVOICE NO):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("InvoiceNo"),
							},

							Label{
								MaxSize: Size{10, 10},
								Text:    "日期(Date):",
							},
							DateEdit{
								MaxSize: Size{10, 10},
								Date:    Bind("Date"),
							},

							Label{
								MaxSize: Size{10, 10},
								Text:    "公司名(TO):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("To"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "合同号(ContractNo):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("ContractNo"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "地址(Address):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("Address"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "支付方式(PAYMENT TERMS):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("PaymentTerms"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "进出口交换比率(Terms of trade):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("TermsOfTrade"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "箱号(NO.OF Pallet):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("Pallet"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "毛重(GROSS WT):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("Gross"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "量度(MEASUREMENT):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("Measurement"),
							},
							VSpacer{
								ColumnSpan: 2,
								Size:       8,
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "客户姓名(Customer name):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("Customer"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "订单号(Order Number):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("Order"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "生产订单(Production Order):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("Production"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "配送地址(Shipping address):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("Shipping"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "签字(Warehouse signature):",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("Signature"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "Qc signature:",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("Qcsignature"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "customer sign back:",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("Signback"),
							},
							Label{
								MaxSize: Size{10, 10},
								Text:    "driver signature:",
							},
							LineEdit{
								MaxSize: Size{10, 10},
								Text:    Bind("DriverSignature"),
							},
						},
					},
					Composite{
						Layout:   Grid{Columns: 6},
						Children: widgetlit,
					},
					Composite{
						Layout:   Grid{Columns: 6},
						Children: widgetlit2,
					},
				},
			},
			PushButton{
				Text: "生成表格",
				OnClicked: func() {
					if err := db.Submit(); err != nil {
						fmt.Println(err)
						return
					}
					goods := make([]model.Good, 0)
					pallets := make([]model.Packing, 0)
					var total float64
					for i, hsno := range Hsnolist {
						if hsno.Text() != "" {
							quantity, err := strconv.ParseFloat(Quantitylist[i].Text(), 32)
							unitprice, err1 := strconv.ParseFloat(UnitPricelist[i].Text(), 32)
							netwt, err3 := strconv.ParseFloat(Netlist[i].Text(), 32)
							if err != nil || err1 != nil || err3 != nil {
								walk.MsgBox(mw, "失败", "请输入正确的单价和重量", walk.MsgBoxIconInformation)
								return
							}
							good := model.Good{
								HsNo:        hsno.Text(),
								ItemNo:      ItemNolist[i].Text(),
								Description: Deslist[i].Text(),
								Quantity:    quantity,
								NetWt:       netwt,
								UnitPrice:   unitprice,
							}
							good.Amount = good.Quantity * good.UnitPrice
							total += good.Amount
							goods = append(goods, good)
						}
					}
					var count int
					for i, name := range Namelist {
						if name.Text() != "" {
							count++
							var pallet model.Packing
							weight, err := strconv.ParseFloat(Weightlist[i].Text(), 32)
							cartons, err1 := strconv.ParseFloat(Cartonslist[i].Text(), 32)
							datas, err2 := strconv.ParseFloat(Palletlist[i].Text(), 32)
							if err != nil || err1 != nil || err2 != nil {
								walk.MsgBox(mw, "失败", "请输入正确的单价和重量", walk.MsgBoxIconInformation)
								return
							}
							pallet.Id = idMap[count]
							pallet.Name = name.Text()
							pallet.Standard = Standardlist[i].Text()
							pallet.Code = Codelist[i].Text()
							pallet.Weight = weight
							pallet.Cartons = cartons
							pallet.Pallet = datas
							pallets = append(pallets, pallet)
						}
					}
					invoice.Goods = goods
					invoice.Total = total
					mw.doExcel(*invoice)
					mw.doPallertExcel(pallets, *invoice)
				},
			},
		},
	}.Run()
}
func (mv *MyMainWindow) doPallertExcel(pallets []model.Packing, invoice model.Invoice) {
	fmt.Println(pallets)
	excelFileName := "document/packinglist.xlsx"

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	setTitle(xlFile, invoice)

	for k := 0; k < len(pallets)-7; k++ {
		row := &xlsx.Row{Sheet: xlFile.Sheets[0]}
		cells := make([]*xlsx.Cell, 7)
		for i := 0; i < 7; i++ {
			cells[i] = xlsx.NewCell(row)
			cells[i].SetStyle(xlFile.Sheets[0].Rows[7].Cells[0].GetStyle())
		}
		row.Cells = cells
		rows := xlFile.Sheets[0].Rows[7:]
		xlFile.Sheets[0].Rows = append(xlFile.Sheets[0].Rows[:7], row)
		xlFile.Sheets[0].Rows = append(xlFile.Sheets[0].Rows, rows...)
	}
	var weighttotal float64
	var cartontotal float64
	var pellettotal float64
	for k, v := range pallets {
		weighttotal += v.Weight
		cartontotal += v.Cartons
		pellettotal += v.Pallet
		for i, cell := range xlFile.Sheets[0].Rows[7+k].Cells {
			cell.GetStyle().Font = *xlsx.NewFont(10, "Times New Roman")
			switch i {
			case 0:
				cell.SetString(v.Id)
			case 1:
				cell.SetString(v.Name)
			case 2:
				cell.SetString(v.Standard)
			case 3:
				cell.SetString(v.Code)
			case 4:
				cell.SetString(strconv.FormatFloat(v.Weight, 'f', 2, 64))
			case 5:
				cell.SetString(strconv.FormatFloat(v.Cartons, 'f', 2, 64))
			case 6:
				cell.SetString(strconv.FormatFloat(v.Pallet, 'f', 2, 64))
			}
		}
	}
	var index int
	if len(pallets) > 7 {
		index = len(pallets)
	} else {
		index = 7
	}
	xlFile.Sheets[0].Rows[7+index].Cells[4].SetString(strconv.FormatFloat(weighttotal, 'f', -1, 64))
	xlFile.Sheets[0].Rows[7+index].Cells[5].SetString(strconv.FormatFloat(cartontotal, 'f', -1, 64))
	xlFile.Sheets[0].Rows[7+index].Cells[6].SetString(strconv.FormatFloat(pellettotal, 'f', -1, 64))
	err = xlFile.Save("result/packinglist.xlsx")
	if err != nil {
		walk.MsgBox(mv, "保存失败", err.Error(), walk.MsgBoxIconInformation)
	}
}

func getValue(invoice model.Invoice, key string) string {
	invoiceval := reflect.ValueOf(&invoice).Elem()
	for i := 0; i < invoiceval.NumField()-1; i++ {
		value := invoiceval.Field(i)
		name := invoiceval.Type().Field(i).Name
		if name == key {
			if value.Kind() == reflect.String {
				return value.String()
			} else if value.Kind() == reflect.Float64 {
				return strconv.FormatFloat(value.Float(), 'E', -1, 64)
			} else {
				date, ok := value.Interface().(time.Time)
				if ok {
					return strings.Split(date.Format("2006-01-02 15:04:05"), " ")[0]
				}
			}
		}
	}
	return ""
}

func setTitle(xlFile *xlsx.File, invoice model.Invoice) {
	for i, sheet := range xlFile.Sheets {
		if i == 1 {
			break
		}
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				if strings.ContainsAny(text, "?") {
					strs := strings.Split(text, "?")
					for i := 0; i < len(strs)/2; i++ {
						text = strings.Replace(text, "?"+strs[i*2+1]+"?", getValue(invoice, strs[i*2+1]), -1)
					}
					cell.SetString(text)
				}
			}
		}
	}
}
func makeNewRow(xlFile *xlsx.File, index int, invoice model.Invoice) {
	for k := 0; k < len(invoice.Goods)-index; k++ {
		row := &xlsx.Row{Sheet: xlFile.Sheets[0]}
		cells := make([]*xlsx.Cell, 6)
		for i := 0; i < 6; i++ {
			cells[i] = xlsx.NewCell(row)
			cells[i].SetStyle(xlFile.Sheets[0].Rows[10].Cells[0].GetStyle())
		}
		row.Cells = cells
		rows := xlFile.Sheets[0].Rows[10:]
		xlFile.Sheets[0].Rows = append(xlFile.Sheets[0].Rows[:10], row)
		xlFile.Sheets[0].Rows = append(xlFile.Sheets[0].Rows, rows...)
	}
}

func makeNewRow1(xlFile *xlsx.File, index int, invoice model.Invoice) {
	for k := 0; k < len(invoice.Goods)-index; k++ {
		row := &xlsx.Row{Sheet: xlFile.Sheets[0]}
		cells := make([]*xlsx.Cell, 7)
		for i := 0; i < 7; i++ {
			cells[i] = xlsx.NewCell(row)
			cells[i].SetStyle(xlFile.Sheets[0].Rows[10].Cells[0].GetStyle())
			if i == 2 {
				cells[i].Merge(1, 0)
			}
		}
		row.Cells = cells
		rows := xlFile.Sheets[0].Rows[10:]
		xlFile.Sheets[0].Rows = append(xlFile.Sheets[0].Rows[:10], row)
		xlFile.Sheets[0].Rows = append(xlFile.Sheets[0].Rows, rows...)
	}
}
func setCell(xlFile *xlsx.File, k int, v model.Good, typeRow int) {
	for i, cell := range xlFile.Sheets[0].Rows[11+k].Cells {
		cell.GetStyle().Font = *xlsx.NewFont(10, "Times New Roman")
		k = i
		if typeRow == 2 {
			if i > 3 {
				k = i - 1
			} else if i == 3 {
				continue
			}
		}
		switch k {
		case 0:
			cell.SetString(v.HsNo)
		case 1:
			cell.SetString(v.ItemNo)
		case 2:
			cell.SetString(v.Description)
		case 3:
			cell.SetString(strconv.FormatFloat(v.Quantity, 'f', 2, 64))
		case 4:
			cell.SetString(strconv.FormatFloat(v.UnitPrice, 'f', 2, 64))
		case 5:
			cell.SetString(strconv.FormatFloat(v.Amount, 'f', 2, 64))
		}
	}
}
func (mv *MyMainWindow) doExcel(invoice model.Invoice) {
	excelFileName := "document/invoice.xlsx"
	//第二章表
	excelFileName2 := "document/oversea.xlsx"

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	xlFile2, err1 := xlsx.OpenFile(excelFileName2)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	setTitle(xlFile, invoice)
	setTitle(xlFile2, invoice)
	if len(invoice.Goods) > 3 {
		makeNewRow(xlFile, 3, invoice)
		if len(invoice.Goods) > 4 {
			makeNewRow1(xlFile2, 4, invoice)
		}
	}

	if len(invoice.Goods) > 8 {
		for k := 0; k < len(invoice.Goods)-8; k++ {
			row := &xlsx.Row{Sheet: xlFile.Sheets[1]}
			row.SetHeightCM(2)
			cells := make([]*xlsx.Cell, 8)
			for i := 0; i < 8; i++ {
				cells[i] = xlsx.NewCell(row)
				cells[i].SetStyle(xlFile.Sheets[1].Rows[11].Cells[0].GetStyle())
				if i == 3 || i == 6 || i == 7 {
					cells[i].Merge(0, 6+k)
				}
			}
			row.Cells = cells
			rows := xlFile.Sheets[1].Rows[11:]
			xlFile.Sheets[1].Rows = append(xlFile.Sheets[1].Rows[:11], row)
			xlFile.Sheets[1].Rows = append(xlFile.Sheets[1].Rows, rows...)
		}
	}
	var quantitytotal float64
	var nettotal float64
	for k, v := range invoice.Goods {
		quantitytotal += v.Quantity
		nettotal += v.NetWt
		setCell(xlFile, k, v, 1)
		setCell(xlFile2, k, v, 2)
		for i, cell := range xlFile.Sheets[1].Rows[10+k].Cells {
			cell.GetStyle().Font = *xlsx.NewFont(10, "Times New Roman")
			switch i {
			case 0:
				cell.SetString(v.HsNo)
			case 1:
				cell.SetString(v.ItemNo)
			case 2:
				cell.SetString(v.Description)
			case 4:
				cell.SetString(strconv.FormatFloat(v.Quantity, 'f', 2, 64))
			case 5:
				cell.SetString(strconv.FormatFloat(v.NetWt, 'f', 2, 64))
			}
		}
	}
	xlFile.Sheets[1].Rows[10].Cells[3].SetString(invoice.Pallet)
	xlFile.Sheets[1].Rows[10].Cells[6].SetString(invoice.Gross)
	xlFile.Sheets[1].Rows[10].Cells[7].SetString(invoice.Measurement)
	var index, index1 int
	if len(invoice.Goods) > 4 {
		index = len(invoice.Goods)
	} else {
		index = 5
	}
	if len(invoice.Goods) > 8 {
		index1 = len(invoice.Goods)
	} else {
		index1 = 8
	}
	if len(invoice.Goods) > 3 {
		index = len(invoice.Goods)
	} else {
		index = 3
	}
	xlFile2.Sheets[0].Rows[11+index].Cells[4].SetString(strconv.FormatFloat(quantitytotal, 'f', -1, 64))
	xlFile2.Sheets[0].Rows[11+index].Cells[6].SetString(strconv.FormatFloat(invoice.Total, 'f', -1, 64))
	xlFile.Sheets[0].Rows[11+index].Cells[1].SetString(strconv.FormatFloat(invoice.Total, 'f', -1, 64))
	xlFile.Sheets[1].Rows[10+index1].Cells[4].SetString(strconv.FormatFloat(quantitytotal, 'f', -1, 64))
	xlFile.Sheets[1].Rows[10+index1].Cells[5].SetString(strconv.FormatFloat(nettotal, 'f', -1, 64))
	xlFile.Sheets[1].Rows[10+index1].Cells[6].SetString(invoice.Gross)
	xlFile.Sheets[1].Rows[10+index1].Cells[7].SetString(invoice.Measurement)
	err = xlFile.Save("result/invoice.xlsx")
	if err != nil {
		walk.MsgBox(mv, "保存失败", err.Error(), walk.MsgBoxIconInformation)
	}
	err = xlFile2.Save("result/oversea.xlsx")
	if err != nil {
		walk.MsgBox(mv, "保存失败", err.Error(), walk.MsgBoxIconInformation)
	}

	walk.MsgBox(mv, "保存成功", "successful", walk.MsgBoxIconInformation)

}
