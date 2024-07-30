package fazz

import "fmt"

type skillSet struct {
	name string
}

type Bio struct {
	fullName     string
	age          int
	gender       string
	isSmoker     bool
	cigarVariant []string
	skillSet     []skillSet
}

func Survey() {
	respondent := []Bio{
		{
			fullName:     "Ilyas",
			age:          20,
			gender:       "laki-laki",
			isSmoker:     true,
			cigarVariant: []string{"Esse", "Sampoerna"},
			skillSet: []skillSet{
				{
					name: "Javasript",
				},
				{
					name: "HTML",
				},
				{
					name: "PHP",
				},
			},
		},
		{
			fullName: "Fulan",
			age:      16,
			gender:   "perempuan",
			isSmoker: false,
		},
		{
			fullName:     "Somat",
			age:          40,
			gender:       "laki-laki",
			isSmoker:     true,
			cigarVariant: []string{"Esse", "Sampoerna", "Magnum"},
		},
		{
			fullName:     "Somat",
			age:          40,
			gender:       "laki-laki",
			isSmoker:     true,
			cigarVariant: []string{"Esse", "Sampoerna", "Magnum"},
		},
		{
			fullName:     "Somat",
			age:          40,
			gender:       "laki-laki",
			isSmoker:     true,
			cigarVariant: []string{"Esse", "Sampoerna", "Magnum"},
			skillSet: []skillSet{
				{
					name: "Javasript",
				},
				{
					name: "HTML",
				},
				{
					name: "PHP",
				},
			},
		},
	}
	// for _, x := range respondent {
	// 	fmt.Printf("Name : %s , Age : %d , isSmoker : %t , cigar variant : ", x.name, x.age, x.isSmoker)
	// 	for _, y := range x.cigarVariant {
	// 		fmt.Printf("%s ", y)
	// 	}
	// 	fmt.Println("")
	// }
	fmt.Println(respondent[0].skillSet[0].name)
	fmt.Println(respondent[4].skillSet[2].name)
}
