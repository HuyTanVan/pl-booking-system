package db

import "context"

type CreateUserTxParams struct {
	CreateUserParams
	CreateEmailVerificationTokenParams
	AfterCreate func(user User, token EmailVerificationToken) error
}

type CreateUserTxResult struct {
	User  User
	Token EmailVerificationToken
}

func (store *Store) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}
		// if user is created successfully, then user object above has its ID
		arg.CreateEmailVerificationTokenParams.UserID = result.User.ID
		result.Token, err = q.CreateEmailVerificationToken(ctx, arg.CreateEmailVerificationTokenParams)
		if err != nil {
			return err
		}
		return arg.AfterCreate(result.User, result.Token)
	})

	return result, err
}
