package singleton

// 饿汉模式单例
type Singleton struct {}

var singleton *Singleton 

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton{
	return singleton
}