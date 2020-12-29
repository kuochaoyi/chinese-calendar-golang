package animal

// Animal 生肖
type Animal struct {
	order int64
}

var animalAlias = [...]string{
	"鼠", "牛", "虎", "兔", "龍", "蛇",
	"馬", "羊", "猴", "雞", "狗", "猪",
}

// NewAnimal 創建生肖物件
func NewAnimal(order int64) *Animal {
	if !isSupported(order) {
		return nil
	}
	return &Animal{order: order}
}

// Alias 返回生消名稱(鼠牛虎...)
func (animal *Animal) Alias() string {
	return animalAlias[(animal.order-1)%12]
}

func isSupported(order int64) bool {
	return 1 <= order && order <= 12
}
