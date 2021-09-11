package chart

import "encoding/json"

func GetBarString(cost []float64, names []string) string {
	o := NewBarStruct(cost, names)
	s, _ := json.Marshal(o)
	return HeadShortCode + string(s) + FootShortCode
}

func NewBarStruct(cost []float64, names []string) *PieStruct {
	var ds = []Datasets{
		{
			Data:            cost,
			BackgroundColor: colorArr,
		},
	}
	return &PieStruct{
		Type: "bar",
		Data: Data{
			Labels:   names,
			Datasets: ds,
		},
		Options: Options{
			Plugins: Plugins{
				Title: Title{
					Display: true,
					Text:    "kingram的持仓收益柱状图(单位：usdt)",
				},
				Legend: Legend{
					Display:  false,
					Position: "right",
				},
			},
			MaintainAspectRatio: false,
		},
	}
}
