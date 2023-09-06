package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID         string    `valid:"uuid"`
	ResourceID string    `valid: "notnull"`
	FilePath   string    `valid: "notnull"`
	CreatedAt  time.Time `valid: "-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{}
}

//	É a variavel que vai ser atribuida
//
// coloca depois do func um (video *Video)
//
//	É a struct principal, por isso o *
func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)
	if err != nil {
		return err
	}
	return nil
}
