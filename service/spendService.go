package service

import (
	"android-be/model"
	"android-be/repository"
	"time"

	"github.com/google/uuid"
)

type SpendService struct {
	repo *repository.Database
}

func NewSpendService(rp *repository.Database) *SpendService {
	return &SpendService{
		repo: rp,
	}
}

func MappingCategory(category string) string {
	var mapping = map[string]string{
		"Ăn uống": "food",
		//		"Đồ dùng học tập":       "ruler",
		//		"Di chuyển":             "move",
		//		"May mặc":               "clother",
		//		"Thuê nhà":              "house",
		//		"Điện nước":             "electric",
		//		"Hóa đơn tiện ích khác": "bill",
		//		"Học phí":               "fee",
		//		"Sức khỏe":              "helth",
		//		"Thể thao":              "spot",
		//		"Lương":                 "salary",
		//		"Gia đình":              "family",
		//		"Thu nhập khác":         "other",
	}

	if mapping[category] != "" {
		return mapping[category]
	} else {
		return category
	}
}

func MappingDBCate(category string) string {
	var mapping = map[string]string{
		"food": "Ăn uống",
		//		"ruler":    "Đồ dùng học tập",
		//		"move":     "Di chuyển",
		//		"clother":  "May mặc",
		//		"house":    "Thuê nhà",
		//		"electric": "Điện nước",
		//		"bill":     "Hóa đơn tiện ích khác",
		//		"fee":      "Học phí",
		//		"helth":    "Sức khỏe",
		//		"spot":     "Thể thao",
		//		"salary":   "Lương",
		//		"family":   "Gia đình",
		//"other":    "Thu nhập khác",
	}

	if mapping[category] != "" {
		return mapping[category]
	} else {
		return category
	}
}

func (s *SpendService) ListSpend(uid string) ([]model.Spending, error) {
	sps, err := s.repo.GetListSpendByUid(uid)
	res := []model.Spending{}
	for _, v := range sps {
		v.Category = MappingDBCate(v.Category)
		res = append(res, v)
	}
	return res, err
}

func (s *SpendService) Create(spend *model.Spending) error {
	spend.Id = uuid.NewString()
	spend.Timestamp = time.Now().UnixMilli()
	spend.Category = MappingCategory(spend.Category)
	err := s.repo.InsertSpend(spend)
	return err
}

func (s *SpendService) GetBill(id string) (model.Spending, error) {
	sp, err := s.repo.GetSpend(id)
	if err != nil {
		return sp, err
	}

	sp.Category = MappingDBCate(sp.Category)
	return sp, err
}

func (s *SpendService) UpdateSpend(spend *model.Spending) error {
	spend.Category = MappingCategory(spend.Category)
	err := s.repo.UpdateSpend(spend)
	return err
}

func (s *SpendService) Delete(id string) error {
	err := s.repo.DeleteSpend(id)
	return err
}

func (s *SpendService) GetInWeek(uid string) ([]model.Spending, error) {
	var part = time.Now().AddDate(0, 0, -7).UnixMilli()
	sps, err := s.repo.GetSpendInWeek(part, uid)
	res := []model.Spending{}
	for _, v := range sps {
		v.Category = MappingDBCate(v.Category)
		res = append(res, v)
	}
	return res, err
}
