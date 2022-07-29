
import (
	"sync".
	"fmt"
	"time"
)

wg := &sync.WaitGroup{}
func main (
	workNumber := 100
	ch := make(chan bool, 3)
	for i := 0;  i < workNumber; i++ {
		go Work1(ch, i)
	}
	wg.Wait() 
)


func Work1(chan bool, i int) {
	 defer wg.Done()
	 wg.Add(1)
		
	 ch <- true
	 fmt.Print("done worker %d", i)
	 time.sleep(time.Second)
	 <-ch
}