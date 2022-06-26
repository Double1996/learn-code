package adapter

import "fmt"


// AWS Client asw sdk 
type ICreateServer interface {
	CreateSever(cpu, mem float64) error
}

// RunInstance 启动实例
func (c *AwsClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aws client run success, cpu: %f, mem: %f", cpu, mem)
	return nil
}

// AwsClientAdapter 适配器
type AwsClientAdapter struct {
	Client AwsClient
}


// CreateServer  启动实例
func (a *AwsClientAdapter) CreateServer()  error {
	a.Client.RunInstance(cpu, mem)
	return nil
}


type AliyunClient struct{}


// CreateServer 启动实例
func (c *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aws client run success, cpu： %d, mem: %d", cpu, mem)
	return nil
}

// Aliyun 适配器
type AliyunClientAdapter struct {
	Client AliyunClient
}

func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.CreateServer(int(cpu), int(mem))
	return nil
}






