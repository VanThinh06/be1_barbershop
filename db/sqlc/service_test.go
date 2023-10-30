package db

// import (
// 	"context"
// 	"testing"

// 	// "github.com/google/uuid"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// 	"gopkg.in/guregu/null.v4"
// )

// func TestCreateService(t *testing.T) {
// 	arg := CreateServiceParams{
// 		ServiceCategoryID: uuid.MustParse("a9e5df3d-76f2-429b-ade2-dfe7a129f2f9"),
// 		Work:              "Cut ABC",
// 		Timer:             null.IntFrom(120),
// 		Price:             110,
// 		Image:             null.StringFrom("imageservice.png"),
// 		Description:       null.StringFrom("Cut blablablablablablablablablablablablablablablablabl"),
// 	}
// 	service, err := testQueries.CreateService(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, service)
// }
