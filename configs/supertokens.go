package config

import (
	"be-assignment/prisma/db"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword/epmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func InitSupertokens(ctx *gin.Context, client *db.PrismaClient) {
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: os.Getenv("SUPERTOKENS_CONNECTION_URI"),
			APIKey:        os.Getenv("SUPERTOKENS_API_KEY"),
		},
		AppInfo: supertokens.AppInfo{
			AppName:       os.Getenv("SUPERTOKENS_APP_NAME"),
			APIDomain:     os.Getenv("SUPERTOKENS_API_DOMAIN"),
			WebsiteDomain: os.Getenv("SUPERTOKENS_WEBSITE_DOMAIN"),
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(&epmodels.TypeInput{
				Override: &epmodels.OverrideStruct{
					Functions: func(originalImplementation epmodels.RecipeInterface) epmodels.RecipeInterface {
						// create a copy of the originalImplementation func
						originalSignUp := *originalImplementation.SignUp

						// override the sign in up function
						(*originalImplementation.SignUp) = func(email, password, tenantId string, userContext supertokens.UserContext) (epmodels.SignUpResponse, error) {

							// First we call the original implementation of SignUp.
							response, err := originalSignUp(email, password, tenantId, userContext)
							if err != nil {
								return epmodels.SignUpResponse{}, err
							}

							if response.OK != nil {
								// sign up was successful

								// user object contains the ID and email
								user := response.OK.User

								// TODO: Post sign up logic.
								client.User.CreateOne(
									db.User.Email.Set(user.Email),
									db.User.SupertokensUserID.Set(user.ID),
								).Exec(ctx)

								client.Account.CreateOne(
									db.Account.UserID.Set(user.ID),
								).Exec(ctx)

							}
							return response, nil
						}

						// create a copy of the originalImplementation func
						originalSignIn := *originalImplementation.SignIn

						// override the sign in up function
						(*originalImplementation.SignIn) = func(email, password, tenantId string, userContext supertokens.UserContext) (epmodels.SignInResponse, error) {

							// First we call the original implementation of SignIn.
							response, err := originalSignIn(email, password, tenantId, userContext)
							if err != nil {
								return epmodels.SignInResponse{}, err
							}

							if response.OK != nil {
								// sign in was successful

								// user object contains the ID and email
								user := response.OK.User

								// TODO: Post sign in logic.
								fmt.Println(user)

							}
							return response, nil
						}

						return originalImplementation
					},
				},
			}),
			session.Init(nil),
		},
	})

	if err != nil {
		panic(err.Error())
	}
}
