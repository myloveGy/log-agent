package tail

type Config struct {
	Filename string `json:"filename" ini:"filename" yaml:"filename"` // 文件名称
	Path     string `json:"path" ini:"path" yaml:"path"`             // 路径
	Rule     bool   `json:"rule" ini:"rule" yaml:"rule"`             // 使用规则
	Type     string `json:"type" ini:"type" yaml:"type"`             // 内容格式
	Driver   string `json:"driver" ini:"driver" yaml:"driver"`       // 写入处理注册名称
	Name     string `json:"name" ini:"name" yaml:"name"`             // 表名称
}
