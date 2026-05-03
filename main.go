package main

import (
	"fmt"
	"weekly/internals/empat"
)

func main() {
	// nums := []int{}
	// err := satu.ProcessNumber(nums)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// 	urlChn := make(chan string)

	// 	var wg sync.WaitGroup
	// 	urls := []string{"https:google.com", "https://koda.com", "https://rick", "https://morty", "https://dumbways"}

	// 	go RenderFetch(urlChn)

	// 	for _, url := range urls {
	// 		wg.Go(func() {
	// 			dua.WebFetcher(url, urlChn)
	// 		})
	// 	}
	// 	wg.Wait()
	// 	close(urlChn)
	// }

	// func RenderFetch(urlChan chan string) {
	// 	for url := range urlChan {
	// 		log.Println(url)
	// 	}

	// nopal := tiga.NewUser(12, "Nopal")
	// nopal2 := tiga.NewUser(11, "Nopal 2")
	// nopal3 := tiga.NewUser(10, "Nopal 3")
	// nopal4 := tiga.NewUser(10, "Nopal 4")
	// userManager := tiga.UserManager{
	// 	Users: make(map[int]*tiga.User),
	// }
	// fmt.Println(userManager)
	// userManager.AddUser(nopal2)
	// userManager.AddUser(nopal)
	// userManager.AddUser(nopal3)
	// userManager.AddUser(nopal4)

	// user, err := userManager.GetUser(11)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("User Id %d ditemukan : %s \n", user.ID, user.Nama)
	// userManager.GetAllUsers()

	c := empat.Circle{
		Radius: 7,
	}
	r := empat.Rectangle{
		Width:  10,
		Height: 10,
	}
	var circle empat.Shape = &c
	var rectangle empat.Shape = &r
	circleArea := circle.Area()
	rectangleArea := rectangle.Area()
	fmt.Printf("Luas dari circle adalah : %.2f\n", circleArea)
	fmt.Printf("Luas dari rectangle adalah : %.2f\n", rectangleArea)
}
