package converter

import (
	"crud_application/src/model"
	"crud_application/src/model/repository/entity"
)

func ConverterEntityToDomain(
	UserEntity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		UserEntity.Email,
		UserEntity.Password,
		UserEntity.Name,
		UserEntity.Age,
	)

	domain.SetID(UserEntity.Id.Hex())
	return domain
}
