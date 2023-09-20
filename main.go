// NG Challenge 5 : Struct & Method 1
package main

import (
	"fmt"
	"ngc5/hero"
	"ngc5/person"
)

func main() {
	// Studi kasus
	p := person.Person{
		Name: "Bambang",
		Age:  45,
		Job:  "Gambler",
	}

	p.GetInfo()

	// Menjalankan AddYear 5 kali
	for i := 0; i < 5; i++ {
		p.AddYear()
	}

	// Menampilkan informasi setelah AddYear dijalankan
	fmt.Println("\nAddYear 5 kali:")
	p.GetInfo()

	// NG Challenge 5 : Struct & Method 2
	// Membuat slice dari struct Person
	people := []person.Person{
		{Name: "Mount", Age: 24, Job: "Dancer"},
		{Name: "Rashford", Age: 26, Job: "Football Player"},
		{Name: "Martial", Age: 27, Job: "Artist"},
	}

	// Menampilkan informasi setiap orang dalam slice
	for _, p := range people {
		p.GetInfo()
		fmt.Println() // Menambahkan baris kosong untuk pemisahan informasi
	}

	// NG Challenge 5 : struct HERO 2
	hero1 := hero.Hero{
		Name:           "Juggernaut",
		BaseAttack:     100,
		Defence:        50,
		CriticalDamage: 50,
		HealthPoint:    150,
		Weapon: hero.Weapon{
			Attack: 30,
		},
	}

	hero2 := hero.Hero{
		Name:           "Underlord",
		BaseAttack:     80,
		Defence:        70,
		CriticalDamage: 30,
		HealthPoint:    180,
		Weapon: hero.Weapon{
			Attack: 25,
		},
	}

	Battle(hero1, hero2)
}

// Fungsi Battle untuk melakukan pertarungan antara dua Hero
func Battle(attacker hero.Hero, defender hero.Hero) {
	// Hero pertama menyerang Hero kedua
	attacker.IsAttackedBy(defender)
	// Hero kedua menyerang balik Hero pertama
	defender.IsAttackedBy(attacker)

	// Cetak status setelah pertarungan
	fmt.Printf("%s attacking %s.\n", attacker.Name, defender.Name)
	fmt.Printf("%s received damage of %d. Health Point: %d\n", defender.Name, attacker.CountDamage()-defender.Defence, defender.HealthPoint)
	fmt.Printf("%s counterattack  %s.\n", defender.Name, attacker.Name)
	fmt.Printf("%s received damage of %d. Health Point: %d\n", attacker.Name, defender.CountDamage()-attacker.Defence, attacker.HealthPoint)
}
