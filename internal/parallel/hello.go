package parallel

import (
	"fmt"

	mutex "github.com/KsenyaNikitchenko/mutex/pkg"
)

func Say(n int) {
	m := mutex.NewMutex(n)

	for i := 0; i < n; i++ {
		go func() {
			defer m.Unlock()
			fmt.Println("Hello, 世界")
		}()
	}

	m.Wait()
	fmt.Println("Работа всех потоков завершена")
}
