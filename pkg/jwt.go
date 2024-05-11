package pkg

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type Claims struct {
	AntiCsrfToken           *string `json:"antiCsrfToken"`
	Exp                     int64   `json:"exp"`
	Iat                     int64   `json:"iat"`
	Iss                     string  `json:"iss"`
	ParentRefreshTokenHash1 *string `json:"parentRefreshTokenHash1"`
	RefreshTokenHash1       string  `json:"refreshTokenHash1"`
	SessionHandle           string  `json:"sessionHandle"`
	Sub                     string  `json:"sub"`
	TId                     string  `json:"tId"`
}

func DecodeJWT(token string) (*Claims, error) {
	accessToken := token
	segments := strings.Split(accessToken, ".")
	if len(segments) != 3 {
		return nil, nil
	}

	data, err := base64.StdEncoding.DecodeString(segments[1])
	if err != nil {
		return nil, err
	}

	claims := &Claims{}
	if err = json.Unmarshal(data, &claims); err != nil {
		return nil, err
	}

	claimsData := &Claims{
		AntiCsrfToken:           claims.AntiCsrfToken,
		Exp:                     claims.Exp,
		Iat:                     claims.Iat,
		Iss:                     claims.Iss,
		ParentRefreshTokenHash1: claims.ParentRefreshTokenHash1,
		RefreshTokenHash1:       claims.RefreshTokenHash1,
		SessionHandle:           claims.SessionHandle,
		Sub:                     claims.Sub,
		TId:                     claims.TId,
	}

	return claimsData, nil
}
