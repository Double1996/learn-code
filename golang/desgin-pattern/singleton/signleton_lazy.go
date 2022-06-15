package singleton

import "sync"

var (
	lazySinleton *Singleton
	once := &sync.Once{}		
)

func GetLazyInstance() *Singleton {
	 if lazySinleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	 }
	 return lazySingleton	
}