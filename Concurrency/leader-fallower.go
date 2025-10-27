for i := 0; i < 4; i++ {
	go func(id int) {
		for {
			select {
			case msg := <-requests:
				fmt.Printf("Worker %d handling %s\n", id, msg)
			}
		}
	}(i)
}
