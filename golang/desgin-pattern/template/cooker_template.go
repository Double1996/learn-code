package template

type Cooker struct {
	open()
	fire()
	cooke()
	outfire()
  close()
}


type CookMenu struct {
	
}


func (CookMenu) open() {
	fmt.Println("打开开关")
}

func (CookMenu) fire() {
	fmt.Println("开火")
}

// 做菜，交给具体的子类实现
func (CookMenu) cooke() {
}

func (CookMenu) outfire() {
	fmt.Println("关火")
}

func (CookMenu) close() {
	fmt.Println("关闭开关")
}


func doCook(cooker Cooker) {
	cook.open()
	cook.fire()
	cook.cooke()
	cook.outfire()
	cook.close()
}

type XiHongShi struct {
	CookMenu
}

func (*XiHongShi) cooke() {
	fmt.Println("做西红柿")
}

type ChaoJiDan struct {
	CookMenu
}

func (ChaoJiDan) cooke() {
	fmt.Println("做炒鸡蛋")
}