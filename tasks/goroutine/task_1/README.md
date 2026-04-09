Результат при println всегда должен быть 100. 
Решить задачку через:
atomic
mutex
channel 

func main() {
    counter := 0

    for i := 0; i < 100; i++ {
        go func() {
            counter = counter + 1
        }()
    }

    fmt.Println(counter)
}