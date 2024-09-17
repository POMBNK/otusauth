package auth

import "github.com/POMBNK/gateway/internal/entity"

func (r *SignInForm) ToUser() (entity.UserToSignIn, error) {
	return entity.UserToSignIn{
		Username: r.Username,
		Password: r.Password,
	}, nil
}

func (r *SignUpForm) ToUser() (entity.UserToSignUp, error) {
	u := entity.UserToSignUp{
		LastName:  r.LastName,
		FirstName: r.FirstName,
		Phone:     r.Phone,
		Username:  r.Username,
		Password:  r.Password,
	}

	if r.Email != nil {
		email, err := r.Email.MarshalJSON()
		if err != nil {
			return entity.UserToSignUp{}, err
		}
		pemail := string(email)
		u.Email = &pemail
	}

	return u, nil
}
