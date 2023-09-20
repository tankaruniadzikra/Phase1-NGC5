package person

import "fmt"

type Person struct {
	Name string
	Age  int
	Job  string
}

// Method GetInfo untuk menampilkan informasi Person
func (p Person) GetInfo() {
	fmt.Printf("Name : %s\nAge : %d\nJob : %s\n", p.Name, p.Age, p.Job)
}

// Method AddYear untuk menambahkan usia dan mengubah pekerjaan jika usia >= 50
func (p *Person) AddYear() {
	p.Age++
	if p.Age >= 50 {
		p.Job = "Retired"
	}
}
