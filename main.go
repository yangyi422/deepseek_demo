package main

func main() {

}

type Name struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

// 特性
type Properties struct {
	Name Name `json:"name"`
}

// 参数
type Parameters struct {
	Type       string     `json:"type"` // object
	Properties Properties `json:"properties"`
	Required   []string   `json:"required,omitempty"` // 必须字段
}

type Tools struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Paramters   Parameters `json:"parameters"`
}

func getClosingPrice(name string) string {
	switch name {
	case "青岛啤酒":
		return "50"
	case "贵州茅台":
		return "2000"
	default:
		return "没有该股票"
	}
}
