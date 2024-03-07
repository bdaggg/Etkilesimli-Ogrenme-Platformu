package seeders

import "EOP/model"

type SystemModuleSeeder struct{}

func (SystemModuleSeeder) Seed() (any, any) {
	Seed := []model.SystemModule{
		{
			ID:       1,
			Path:     "customer",
			Name:     "Müşteri Menüsü",
			GroupNum: 1,
			Actions: []model.SystemModuleAction{
				{ID: 1, Name: "Müşteri Ekleme", Key: "create"},
				{ID: 2, Name: "Müşteri Görüntüleme", Key: "read"},
				{ID: 3, Name: "Müşteri Güncelleme", Key: "update"},
				{ID: 4, Name: "Müşteri Silme", Key: "delete"},
			},
		},
		{
			ID:       2,
			Path:     "customer-type",
			Name:     "Müşteri Tipleri",
			GroupNum: 1,
			Actions: []model.SystemModuleAction{
				{ID: 5, Name: "Müşteri Tipi Ekleme", Key: "create"},
				{ID: 6, Name: "Müşteri Tipi Görüntüleme", Key: "read"},
				{ID: 7, Name: "Müşteri Tipi Güncelleme", Key: "update"},
				{ID: 8, Name: "Müşteri Tipi Silme", Key: "delete"},
			},
		},
		{
			ID:       3,
			Path:     "report",
			Name:     "Raporlar",
			GroupNum: 2,
			Actions: []model.SystemModuleAction{
				{ID: 9, Name: "Rapor Görüntüleme", Key: "read"},
				{ID: 10, Name: "Rapor Oluşturma", Key: "create"},
				{ID: 11, Name: "Rapor Güncelleme", Key: "update"},
				{ID: 12, Name: "Rapor Silme", Key: "delete"},
			},
		},
	}
	return model.SystemModule{}, Seed
}

type SystemRoleSeeder struct{}

func (SystemRoleSeeder) Seed() (any, any) {
	Seed := []model.SystemRole{
		{
			ID:   1,
			Name: "Admin",
			Privileges: []model.SystemModuleAction{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
		},
		{
			ID:   2,
			Name: "Varsayılan Kullanıcı",
			Privileges: []model.SystemModuleAction{
				{ID: 1},
			},
		},
	}
	return model.SystemRole{}, Seed
}
