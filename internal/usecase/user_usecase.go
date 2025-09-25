package usecase

import (
	"sabj/internal/domain"
)

type UserUsecase struct {
	Repo domain.UserRepository
}

func (uc *UserUsecase) CreateUser(user *domain.User) error {
	// ここでは、バリデーションなどのビジネスロジックを記述できますが、
	// 今回はそのままリポジトリのCreateメソッドを呼び出します。

	err := uc.Repo.Create(user)
	return err
}
