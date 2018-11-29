package models

type Files struct {
	Model
	Bucket	string
	Name    string
	Key 	string `gorm:"column:fkey"`
	Fsize 	int `json:"fsize"`
	Hash 	string `json:"hash"`
}

func AddFileRecord(file *Files) error{
	if err := db.Create(file).Error; err != nil {
		return err
	}
	return nil
}

func FindFileByName(name string) string {
	var file Files
	db.Model(Files{}).Where("name = ?", name).Select("fkey").First(&file)
	return file.Key
}


