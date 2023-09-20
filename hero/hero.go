package hero

import (
	"math/rand"
)

// NG Challenge 5 : struct HERO 1
type Hero struct {
	Name           string
	BaseAttack     int
	Defence        int
	CriticalDamage int
	HealthPoint    int
	Weapon         Weapon
}

type Weapon struct {
	Attack int
}

func (h Hero) CountDamage() int {
	totalDamage := h.BaseAttack + h.Weapon.Attack
	if rand.Intn(2) == 1 {
		totalDamage += h.CriticalDamage
	}
	return totalDamage
}

// Method isAttackedBy untuk menghitung total damage yang diterima dan pengurangan HealthPoint
func (h *Hero) IsAttackedBy(attacker Hero) {
	totalDamageReceived := attacker.CountDamage() - h.Defence

	if totalDamageReceived > 0 {
		h.HealthPoint -= totalDamageReceived
	}
}
