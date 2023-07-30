package main

import "fmt"

type Packages struct {
	List []List
}

type List struct {
	BussinesId string
}

func main() {
	packages := &Packages{}
	var list []List
	// list = append(list, List{BussinesId: "1"})
	// list = append(list, List{BussinesId: ""})
	// list = append(list, List{BussinesId: "3"})
	// list[0].BussinesId = "1"
	// list[1].BussinesId = "2"
	// list[2].BussinesId = "3"
	packages.List = list
	fmt.Println(packages)

	var ids []string

	// data := []string{}
	savedId := []string{"8", "23", "3", "5", "1"}

	fmt.Println(fmt.Sprintf("Leng %v %v %v", len(packages.List), len(ids), len(savedId)))
	// fmt.Println(savedId[4])
	if len(packages.List) != len(ids) {
		fmt.Println("not")
		return
	}
	for i := range packages.List {
		var ss = false
		for j := range savedId {
			fmt.Println(fmt.Sprintf("==%v==%v===", ids[i], savedId[j]))
			if ids[i] == savedId[j] {
				ss = true

				fmt.Println(ids[i], savedId[j])
				continue
			}

		}

		if !ss {
			fmt.Println("CREATE ")
			fmt.Println(ids[i], packages.List[i].BussinesId)
		}
	}
}
