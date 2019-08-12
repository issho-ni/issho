package kazoku

import (
	"io"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalGQL implements graphql.Marshaler for kazoku.UserAccount_Role.
func (r UserAccount_Role) MarshalGQL(w io.Writer) {
	graphql.MarshalString(UserAccount_Role_name[int32(r)]).MarshalGQL(w)
}

// UnmarshalGQL implements graphql.Unmarshaler for kazoku.UserAccount_Role.
func (r *UserAccount_Role) UnmarshalGQL(v interface{}) error {
	s, err := graphql.UnmarshalString(v)
	*r = UserAccount_Role(UserAccount_Role_value[s])
	return err
}
