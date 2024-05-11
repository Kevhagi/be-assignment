package config

import (
	"be-assignment/prisma/db"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword/epmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
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
								createdUser, _ := client.User.CreateOne(
									db.User.Email.Set(user.Email),
									db.User.SupertokensUserID.Set(user.ID),
								).Exec(ctx)

								print(createdUser.ID)

								client.Account.CreateOne(
									db.Account.User.Link(
										db.User.ID.Equals(createdUser.ID),
									),
								).Exec(ctx)

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

// This is a function that wraps the supertokens verification function
// to work the gin
func WrapVerifySession(options *sessmodels.VerifySessionOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		session.VerifySession(options, func(rw http.ResponseWriter, r *http.Request) {
			c.Request = c.Request.WithContext(r.Context())
			c.Next()
		})(c.Writer, c.Request)
		// we call Abort so that the next handler in the chain is not called, unless we call Next explicitly
		c.Abort()
	}
}
