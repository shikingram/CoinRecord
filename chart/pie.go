package chart

import "encoding/json"

const HeadShortCode = `{{< chart >}}`

const FootShortCode = `{{< /chart >}}`

func getRandomColorArray(nums int) []string {
	var arr []string
	for i := 0; i < nums; i++ {
		arr = append(arr, RandomColor())
	}
	return arr
}

var colorArr []string

func GetPieString(cost []float64, names []string) string {
	o := NewPieStruct(cost, names)
	s, _ := json.Marshal(o)
	return HeadShortCode + string(s) + FootShortCode
}

func NewPieStruct(cost []float64, names []string) *PieStruct {
	colorArr = getRandomColorArray(len(names))
	var ds = []Datasets{
		{
			Data:            cost,
			BackgroundColor: colorArr,
		},
	}
	return &PieStruct{
		Type: "doughnut",
		Data: Data{
			Labels:   names,
			Datasets: ds,
		},
		Options: Options{
			Plugins: Plugins{
				Title: Title{
					Display: true,
					Text:    "kingram的持仓环形分布图(单位：usdt)",
				},
				Legend: Legend{
					Display:  true,
					Position: "right",
				},
			},
			MaintainAspectRatio: false,
		},
	}
}

type PieStruct struct {
	Type    string  `json:"type"`
	Data    Data    `json:"data"`
	Options Options `json:"options"`
}
type Datasets struct {
	Data            []float64 `json:"data"`
	BackgroundColor []string  `json:"backgroundColor"`
}
type Data struct {
	Labels   []string   `json:"labels"`
	Datasets []Datasets `json:"datasets"`
}
type Title struct {
	Display bool   `json:"display"`
	Text    string `json:"text"`
}
type Plugins struct {
	Title  Title  `json:"title"`
	Legend Legend `json:"legend"`
}
type Options struct {
	Plugins             Plugins `json:"plugins"`
	MaintainAspectRatio bool    `json:"maintainAspectRatio"`
}

type Legend struct {
	Display  bool   `json:"display"`
	Position string `json:"position"`
}
