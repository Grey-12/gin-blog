package models

type Tag struct {
	Model

	Name string `json:"name"`
	CreateBy string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tag []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tag)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
