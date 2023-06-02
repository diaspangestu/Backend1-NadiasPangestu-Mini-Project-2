package superadmin

import (
	"github.com/diaspangestu/Backend1-NadiasPangestu-Mini-Project-2/repositories"
)

type UsecaseSuperadminInterface interface {
	ApprovedAdminRegister(id uint) error
	RejectedAdminRegister(id uint) error
}

type UsecaseSuperadmin struct {
	superadminRepo repositories.Superadmin
}

func (uc UsecaseSuperadmin) ApprovedAdminRegister(id uint) error {
	err := uc.superadminRepo.ApprovedAdminRegister(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc UsecaseSuperadmin) RejectedAdminRegister(id uint) error {
	err := uc.superadminRepo.RejectedAdminRegister(id)
	if err != nil {
		return err
	}

	return nil
}
