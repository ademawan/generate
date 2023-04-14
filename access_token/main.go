package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"reflect"
	"time"
	domain "wec-auth/domain/auth_basic"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

var (
	accessTokenWithEncryptAESCBC = "eyJ0eXAiOiJKV1QiLCJraWQiOiJ3VTNpZklJYUxPVUFSZVJCL0ZHNmVNMVAxUU09IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI3YzVjOTZiZS0zMGFjLTQ5NzctYjkyYy02MjY0NjMyYjIwNjkiLCJjdHMiOiJPQVVUSDJfR1JBTlRfU0VUIiwiYXV0aF9sZXZlbCI6MCwiYXVkaXRUcmFja2luZ0lkIjoiMTAzNzI3MDMtZjJkMC00NjYxLThjNDgtMzUyYjRlNmU5MWI0LTgxNSIsImlzcyI6Imh0dHBzOi8vY2lhbWFtcHJlcGRhcHAuY2lhbS50ZWxrb21zZWwuY29tOjEwMDAzL29wZW5hbS9vYXV0aDIvdHNlbC93ZWMvd2ViIiwidG9rZW5OYW1lIjoiYWNjZXNzX3Rva2VuIiwidG9rZW5fdHlwZSI6IkJlYXJlciIsImF1dGhHcmFudElkIjoiOGlYZTFMZ1h0em0yRmdmMldHV2lJWkVKeDZJLjVJWEx0Vk91cGhuTjR1UXN0RXJtRm9qdlh3NCIsIm5vbmNlIjoidHJ1ZSIsImF1ZCI6ImIzOTM2MTg0MzZlNTExZWM4ZDNkMDI0MmFjMTMwMDAzIiwibmJmIjoxNjgxNDUzODU1LCJncmFudF90eXBlIjoiYXV0aG9yaXphdGlvbl9jb2RlIiwic2NvcGUiOlsib3BlbmlkIiwicHJvZmlsZSJdLCJhdXRoX3RpbWUiOjE2ODE0NTM4NTQsInJlYWxtIjoiL3RzZWwvd2VjL3dlYiIsImV4cCI6MTY4MTQ1NDc1NSwiaWF0IjoxNjgxNDUzODU1LCJleHBpcmVzX2luIjo5MDAsImp0aSI6IjhpWGUxTGdYdHptMkZnZjJXR1dpSVpFSng2SS5EYVRSZGdDOEpiQ19QUUxBQmxDWklPQ3plSG8ifQ.bD0f9kzNQBpzqobqHWHXFsgrM-doBKIgl0vpwdWWekVrYcKYXD6dVUqHBSi5XqgVGypwhCJH6JrDXozIsNuSnzLEXAP8KfE8uYaPvZDnjZCHX7wpIEpyvAJm8cxnGao1oMK0BOdQ6gzZnvoTxQggN8hk5wgvGcmET6a74IsTqVbsk1mVwKZvjJ0enSWD40TA9xQf5RzX2gKsBUhcccNH6mHFxK0BJLwoeiy42woAtchGvQ6wLO07gtYTZxPUC23VApt8xNJ4rcS7wJnJ-mb_EEIc1wpTRL6y5VKO9Nr0FV6m7b0iHtFV2yXCO6WH5w2AIJOWzfE21jOgt1bjQj2QbA|24eb20adf69dd404a4024064ef0c349793542dfc9480e0cdd85c0e2828ce2032fafa324263bb6f4d06fd66b1457329fbf234bac597e7a840228593a390b609dd658f42fff2950a0c1b622f89a2294255ae791e12c79706e5a395c29287c795d285d3d8df64df9dc5d27af902fedceb3898de6b5f1b72368cbbb8beb26faaa3bf26e8e546cf0134c2c5d61079acdcd3485f46022797950ead2aa85d1ee5aaaf3de4ad69c352669cc3bee72294ff77c85a5f0c15e8ca58ff7ea830b365ed95ca15d8f86c71e47a763b78fe698cadc4b136"

	accessToken = "eyJ0eXAiOiJKV1QiLCJraWQiOiJ3VTNpZklJYUxPVUFSZVJCL0ZHNmVNMVAxUU09IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI3YzVjOTZiZS0zMGFjLTQ5NzctYjkyYy02MjY0NjMyYjIwNjkiLCJjdHMiOiJPQVVUSDJfR1JBTlRfU0VUIiwiYXV0aF9sZXZlbCI6MCwiYXVkaXRUcmFja2luZ0lkIjoiODM3ZGRmNmUtYWQwZC00MGZkLTk2NGUtZmZiMWUwM2UzMzJmLTQ2NzUiLCJpc3MiOiJodHRwczovL2NpYW1hbXByZXBkYXBwLmNpYW0udGVsa29tc2VsLmNvbToxMDAwMy9vcGVuYW0vb2F1dGgyL3RzZWwvd2VjL3dlYiIsInRva2VuTmFtZSI6ImFjY2Vzc190b2tlbiIsInRva2VuX3R5cGUiOiJCZWFyZXIiLCJhdXRoR3JhbnRJZCI6IjhpWGUxTGdYdHptMkZnZjJXR1dpSVpFSng2SS5JYk1NR3R5RTVVWkFGWWh0OXo4UURRc1F4R1EiLCJub25jZSI6InRydWUiLCJhdWQiOiJiMzkzNjE4NDM2ZTUxMWVjOGQzZDAyNDJhYzEzMDAwMyIsIm5iZiI6MTY4MTQ2MjMzOCwiZ3JhbnRfdHlwZSI6ImF1dGhvcml6YXRpb25fY29kZSIsInNjb3BlIjpbIm9wZW5pZCIsInByb2ZpbGUiXSwiYXV0aF90aW1lIjoxNjgxNDYyMzM3LCJyZWFsbSI6Ii90c2VsL3dlYy93ZWIiLCJleHAiOjE2ODE0NjMyMzgsImlhdCI6MTY4MTQ2MjMzOCwiZXhwaXJlc19pbiI6OTAwLCJqdGkiOiI4aVhlMUxnWHR6bTJGZ2YyV0dXaUlaRUp4NkkubG9LOWk5cGhJbk9wZWRmSXlqTWJZRTFIb2M0In0.bBimpoUs0DzMmD6HPWxTF2XEmxpgH7nspCTm_8PcBVoUaPYWkwhWAinJYRw55jdr2JVle0FxhPbL8Ccw5ylyWMbN2q1wwZrSAT_L30taGjZKGuIWw2n6lYveEb7NkJ6N9poakm9PECJOLGgSOrfdlUWEg-hC1QKZgXguQJPwe8uGP7KoPbtl0O6Q-K6G88bG1DV12pQyUlH-FRqa-z7qIzEntD18SlCnoCkMRzGbDwAsA_JcAIgpKARDkjTsa5a2Ch-35tk4kijjIjh7jTm-dM5lq4O_OaUJrV9Y7DQ3m7p49z_BzIu8gZ5gJOaGDF5WQnCMg4vc46KDtviDaYshsQ"

	encryptedAesCBC = "24eb20adf69dd404a4024064ef0c349793542dfc9480e0cdd85c0e2828ce2032fafa324263bb6f4d06fd66b1457329fbf234bac597e7a840228593a390b609dd658f42fff2950a0c1b622f89a2294255ae791e12c79706e5a395c29287c795d285d3d8df64df9dc5d27af902fedceb3898de6b5f1b72368cbbb8beb26faaa3bf26e8e546cf0134c2c5d61079acdcd3485f46022797950ead2aa85d1ee5aaaf3de4ad69c352669cc3bee72294ff77c85a5f0c15e8ca58ff7ea830b365ed95ca15d8f86c71e47a763b78fe698cadc4b136"
)

func main() {
	fmt.Println(GenerateKey())

	uuid, _ := ValidateJwt("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJURVNUSU5HIjoiVEVTVElORyBBSkEiLCJhdWQiOiJiMzkzNjE4NDM2ZTUxMWVjOGQzZDAyNDJhYzEzMDAwMyIsImF1ZGl0VHJhY2tpbmdJZCI6ImI0ZGFhMjBlLTljMjYtNDUyYy1hN2MyLWZmNDYwZTJiNWRlNy01Mzc1IiwiYXV0aEdyYW50SWQiOiI4aVhlMUxnWHR6bTJGZ2YyV0dXaUlaRUp4NkkuUGx4aGNTMUdIdDZlWGVMa1ZRdW5pZVZSQUcwIiwiYXV0aF9sZXZlbCI6MCwiYXV0aF90aW1lIjoxNjgxNDYwODEzLCJjdHMiOiJPQVVUSDJfR1JBTlRfU0VUIiwiZXhwIjoxNjgxNDYxNzE0LCJleHBpcmVzX2luIjo5MDAsImdyYW50X3R5cGUiOiJhdXRob3JpemF0aW9uX2NvZGUiLCJpYXQiOjE2ODE0NjA4MTQsImlzcyI6Imh0dHBzOi8vY2lhbWFtcHJlcGRhcHAuY2lhbS50ZWxrb21zZWwuY29tOjEwMDAzL29wZW5hbS9vYXV0aDIvdHNlbC93ZWMvd2ViIiwianRpIjoiOGlYZTFMZ1h0em0yRmdmMldHV2lJWkVKeDZJLnV4Nk1NbEJaRVBiVXpGb1JVeHFTNUlUd0M0YyIsIm5iZiI6MTY4MTQ2MDgxNCwibm9uY2UiOiJ0cnVlIiwicmVhbG0iOiIvdHNlbC93ZWMvd2ViIiwic2NvcGUiOlsib3BlbmlkIiwicHJvZmlsZSJdLCJzdWIiOiI3YzVjOTZiZS0zMGFjLTQ5NzctYjkyYy02MjY0NjMyYjIwNjkiLCJ0b2tlbk5hbWUiOiJhY2Nlc3NfdG9rZW4iLCJ0b2tlbl90eXBlIjoiQmVhcmVyIn0.l3y07VVxDPDBWHbdiOuiTWzgVUM4uvH-vDyUqvYczxU")
	fmt.Println("uuuiiiiddddd=====:", uuid)
	claims2, _ := GetClaims("eyJ0eXAiOiJKV1QiLCJjdHkiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.ZXlKMGVYQWlPaUpLVjFRaUxDSmxibU1pT2lKQk1USTRRMEpETFVoVE1qVTJJaXdpWVd4bklqb2lVbE5CTVY4MUluMC5EeEQ2dENoaS1NaEpaNjNMWFh2UXQ0OXFQNUIxYWVCckk4WUxqMGxyWXBSYjFJVDQ1Vjd4RmpkZVJkR2N6Y1lnTGVqU0lDczBxNGdvVlVDM3lLT3BUby1GdjV3Z2FEVlp4dHItbzBUVmx5TmhBZDlRV2UybWx1SGFtNFlCTnFmT002UjB1MU1TNVJKTXZhbnRFejVVNjVlTXJTMUhfV290Vjg3QUlyMWVUZG5scFZkcmdxbjBZMmU5dnplcTVYaEF3LWMyQ2dnRWlZdUlUVzhfTVptZU1RTVpDRllJUk1UZGZrTDE2dHpsVjRIV2NKYXpiY3VoZTl0N0VINVJTREpaUTlWN00yMmJRdUVaOHdyaWJ3U0FRZDFyOTRLUXYzbEJhajRXRGNYaXl4QjlvRkxoY2x0bkdSMTJnS3REeWRSS1FteU00ZXpmaDhxTklvblN5TklRUkEua01uTURfa29MUXdUSTBiMU5OZG1pQS5FSG1NT0RBTkNHOGVySXpDT2xtMFJJekpYZ0lRdHNDU2x5UmhFRkJ2aUFGR3JtYVhGc01VdDhhQThnb0lheGgtTWc3bENTOFQtZmZWd0wtOVQ3enJlTTJtQmJwYXV6VzFfM2VRcThQdnBtb0pkeHFfay1SSmlBX3kzVEtEVWFRdGh1R1QyZ0x3Mm1PY2V1Y2k4MEtyeWhpb1hIU3RnQ2hKZUtBazlib29hTHllZDRycU5tcHhmODJ3TV9vT2ktLWdrSlJwNk1HNGk5dDhLelh2RzZBNk5uR2xYTWJQek5kbUxzekduOGFhSzFFc3RqRkRQaGJqaVVNbUVEeS1HVmJLV3NFcjBkVHE3UkpLdkUwdjNvMFBzSko3OWlFekJqbHVNTkw0VmljTXJ0S2ZCdlZJY2Zvd1pQMXA5X3lsWHFzS1h1QmFpTlFTZUxpcUw3SWE2WWRJWHl5NF9rRDViNV9OcjNuUWo5blhQaE1DUnhnV21zZ0dnQ09hVHRWOWpxRWZtanl0eGc3Wi1SX29rYTFRZElNV0ZEVjdlLWVuVXViUnBxbnYxSjRHcWUtdkhYaXZMQVBmTzFKcVFxcm1pT2pzdWVOaVVkWl9Dc2ZKb21fOUdDb1NVQnVBTUpFZm9hT2pISnZOV2RwQVlRQXRmM09oM3dxX1ZLM1Y0UjI4TXpTdnU0NzBOeGhmejRkWllwSElFMzRvQ1JYUV9BOFBzSGs0a0VKWnI5MFhnOTh2TmtPUmg4LUZHbklneWFGZEdnMmFpLXh6bTgzTVRjX3JRLWM1ODJYTmlvVExCTFpkM0xTdE1PLUw5amJJQ24xYWgyQmVQMncxRjJaMFU1djFpUDhxM0w1S2hwazBCNU5NN1A4c2l5M1RMcGF4QnE0Z2pERk1RQ2ZLcGZuRzVOX3NsOHRuQ0RTaDVLSXpXZWFIYm5rRmxwSTJERHl0MTZuV3A2cS1rSG1xZ3cud2o5VWFRZWZ6QWZtNzNTblQyWjA2QQ.TUZYLg5nQojlmZIMmQNWfNRyxWAbWUpSjdaL-IsKMyE")
	fmt.Println(claims2)

	file2, _ := json.MarshalIndent(claims2, "", " ")

	_ = ioutil.WriteFile("datatest.json", file2, 0644)
	claims, err := getClaimsJWTCIAM(accessToken)
	if err == nil {
		for key, val := range claims {
			fmt.Println(key, ":", val, ":", reflect.TypeOf(val))
		}
	}
	fmt.Println("===================")
	fmt.Println(GenerateToken(claims))
	fmt.Println("===================")

	acc, err := RegenerateAccessToken(accessToken)
	if err == nil {
		fmt.Println(acc)
	} else {
		fmt.Println(err.Error())
	}
}

var mySigningKey = []byte("wectelkomselcom")
var js = `{
	"kty": "RSA",
	"kid": "wU3ifIIaLOUAReRB/FG6eM1P1QM=",
	"use": "sig",
	"x5t": "5eOfy1Nn2MMIKVRRkq0OgFAw348",
	"x5c": [
		"MIIDdzCCAl+gAwIBAgIES3eb+zANBgkqhkiG9w0BAQsFADBsMRAwDgYDVQQGEwdVbmtub3duMRAwDgYDVQQIEwdVbmtub3duMRAwDgYDVQQHEwdVbmtub3duMRAwDgYDVQQKEwdVbmtub3duMRAwDgYDVQQLEwdVbmtub3duMRAwDgYDVQQDEwdVbmtub3duMB4XDTE2MDUyNDEzNDEzN1oXDTI2MDUyMjEzNDEzN1owbDEQMA4GA1UEBhMHVW5rbm93bjEQMA4GA1UECBMHVW5rbm93bjEQMA4GA1UEBxMHVW5rbm93bjEQMA4GA1UEChMHVW5rbm93bjEQMA4GA1UECxMHVW5rbm93bjEQMA4GA1UEAxMHVW5rbm93bjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANdIhkOZeSHagT9ZecG+QQwWaUsi7OMv1JvpBr/7HtAZEZMDGWrxg/zao6vMd/nyjSOOZ1OxOwjgIfII5+iwl37oOexEH4tIDoCoToVXC5iqiBFz5qnmoLzJ3bF1iMupPFjz8Ac0pDeTwyygVyhv19QcFbzhPdu+p68epSatwoDW5ohIoaLzbf+oOaQsYkmqyJNrmht091XuoVCazNFt+UJqqzTPay95Wj4F7Qrs+LCSTd6xp0Kv9uWG1GsFvS9TE1W6isVosjeVm16FlIPLaNQ4aEJ18w8piDIRWuOTUy4cbXR/Qg6a11l1gWls6PJiBXrOciOACVuGUoNTzztlCUkCAwEAAaMhMB8wHQYDVR0OBBYEFMm4/1hF4WEPYS5gMXRmmH0gs6XjMA0GCSqGSIb3DQEBCwUAA4IBAQDVH/Md9lCQWxbSbie5lPdPLB72F4831glHlaqms7kzAM6IhRjXmd0QTYq3Ey1J88KSDf8A0HUZefhudnFaHmtxFv0SF5VdMUY14bJ9UsxJ5f4oP4CVh57fHK0w+EaKGGIw6TQEkL5L/+5QZZAywKgPz67A3o+uk45aKpF3GaNWjGRWEPqcGkyQ0sIC2o7FUTV+MV1KHDRuBgreRCEpqMoY5XGXe/IJc1EJLFDnsjIOQU1rrUzfM+WP/DigEQTPpkKWHJpouP+LLrGRj2ziYVbBDveP8KtHvLFsnexA/TidjOOxChKSLT9LYFyQqsvUyCagBb4aLs009kbW6inN8zA6"
	],
	"n": "10iGQ5l5IdqBP1l5wb5BDBZpSyLs4y_Um-kGv_se0BkRkwMZavGD_Nqjq8x3-fKNI45nU7E7COAh8gjn6LCXfug57EQfi0gOgKhOhVcLmKqIEXPmqeagvMndsXWIy6k8WPPwBzSkN5PDLKBXKG_X1BwVvOE9276nrx6lJq3CgNbmiEihovNt_6g5pCxiSarIk2uaG3T3Ve6hUJrM0W35QmqrNM9rL3laPgXtCuz4sJJN3rGnQq_25YbUawW9L1MTVbqKxWiyN5WbXoWUg8to1DhoQnXzDymIMhFa45NTLhxtdH9CDprXWXWBaWzo8mIFes5yI4AJW4ZSg1PPO2UJSQ",
	"e": "AQAB",
	"alg": "RS256"
}`
var log = logrus.New()

func GenerateJWTAuthIdMsisdn(authId, msisdn, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "wec telkomsel"
	claims["authId"] = authId
	claims["msisdn"] = msisdn
	claims["email"] = email
	if os.Getenv("EXPIRY_JWT") == "true" {
		claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	}

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func GenerateJWTCustom(userData *domain.UserData, userMerchant *domain.UserMerchant) (string, error) {
	var mySigningKey = []byte(os.Getenv("JWT_SIGN_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["aud"] = userData.Uuid
	claims["auditTrackingId"] = userData.FullName
	claims["authGrantId"] = userData.Email
	claims["auth_level"] = userData.LinkPicture
	claims["auth_time"] = userMerchant.MerchantId
	claims["cts"] = userMerchant.TeamId
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	// if os.Getenv("EXPIRY_JWT") == "true" {
	// 	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	// }

	//key := GenerateKey()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func getClaimsJWTCIAM(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return GenerateKey(), nil
	})
	fmt.Println(claims)
	// ... error handling
	if err != nil {
		return nil, err
	}
	return claims, nil

}
func ValidateJwt(tokenString string) (string, bool) {

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return GenerateKey(), nil
	})
	fmt.Println(claims)
	// ... error handling
	if err != nil {
		StringLog("error", err.Error())
	}

	if claims["sub"] != nil {
		return claims["sub"].(string), token.Valid
	}
	if claims["msisdn"] != nil {
		return claims["msisdn"].(string), true
	}

	fmt.Println("token jwt", claims["msisdn"])
	return "", false
}

func GetClaims(tokenString string) (interface{}, bool) {

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return GenerateKey(), nil
	})
	fmt.Println(claims)
	// ... error handling
	if err != nil {
		StringLog("error", err.Error())
	}

	if claims["sub"] != nil {
		return claims, token.Valid
	}
	if claims["msisdn"] != nil {
		return claims, true
	}

	fmt.Println("token jwt", claims["msisdn"])
	return "", false
}

func GenerateKey() interface{} {
	jwk := map[string]string{}
	json.Unmarshal([]byte(js), &jwk)

	if jwk["kty"] != "RSA" {
		StringLog("error", "invalid key type:"+jwk["kty"])
		return nil
	}

	// decode the base64 bytes for n
	nb, err := base64.RawURLEncoding.DecodeString(jwk["n"])
	if err != nil {
		StringLog("error", err.Error())
		return nil
	}

	e := 0
	// The default exponent is usually 65537, so just compare the
	// base64 for [1,0,1] or [0,1,0,1]
	if jwk["e"] == "AQAB" || jwk["e"] == "AAEAAQ" {
		e = 65537
	} else {
		// need to decode "e" as a big-endian int
		StringLog("error", "need to deocde e: "+jwk["e"])
		return nil
	}

	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nb),
		E: e,
	}

	der, err := x509.MarshalPKIXPublicKey(pk)
	if err != nil {
		StringLog("error", err.Error())
		return nil
	}

	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: der,
	}

	var out bytes.Buffer
	pem.Encode(&out, block)

	block, _ = pem.Decode([]byte(out.String()))
	if block == nil {
		StringLog("error", "failed to parse PEM block containing the public key")
		return nil
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		StringLog("error", "failed to parse DER encoded public key: "+err.Error())
		return nil
	}

	return pub
}

func ValidateJwt2(tokenString string) (string, string, float64, bool) {
	jwk := map[string]string{}
	json.Unmarshal([]byte(js), &jwk)

	if jwk["kty"] != "RSA" {
		fmt.Println("error", "invalid key type:"+jwk["kty"])
	}

	// decode the base64 bytes for n
	nb, err := base64.RawURLEncoding.DecodeString(jwk["n"])
	if err != nil {
		fmt.Println("error", err.Error())
	}

	e := 0
	// The default exponent is usually 65537, so just compare the
	// base64 for [1,0,1] or [0,1,0,1]
	if jwk["e"] == "AQAB" || jwk["e"] == "AAEAAQ" {
		e = 65537
	} else {
		// need to decode "e" as a big-endian int
		fmt.Println("error", "need to deocde e: "+jwk["e"])
	}

	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nb),
		E: e,
	}

	der, err := x509.MarshalPKIXPublicKey(pk)
	if err != nil {
		fmt.Println("error", err.Error())
	}

	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: der,
	}

	var out bytes.Buffer
	pem.Encode(&out, block)

	block, _ = pem.Decode([]byte(out.String()))
	if block == nil {
		fmt.Println("error", "failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("error", "failed to parse DER encoded public key: "+err.Error())
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return pub, nil
	})

	// ... error handling
	if err != nil {
		fmt.Println("error", err.Error())
	}

	if claims["sub"] != nil {
		return claims["sub"].(string), claims["iss"].(string), claims["exp"].(float64), token.Valid
	}

	return "", "", 0, false
}
func StringLog(level string, message string) {
	log.Out = os.Stdout
	log_dir := ""

	if os.Getenv("ENV") == "production" || os.Getenv("ENV") == "preproduction" {
		log_dir = "/app/"
	}

	file, err := os.OpenFile(log_dir+"logs/service-auth.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr (string log)")
	}

	if level == "info" {
		log.Info(message)
	} else if level == "warn" {
		log.Warn(message)
	} else if level == "error" {
		log.Error(message)
	}
}

func GenerateToken(claims jwt.MapClaims) (string, error) {
	// getting variable frsom conf environtment
	// secretKey, _ := web.AppConfig.String("SECRET_KEY")
	// making claimes
	// claims := jwt.MapClaims{}
	// claims["authorized"] = true
	// claims["id"] = id
	// claims["merchant_id"] = merchantId
	// claims["role_id"] = roleId
	// claims["is_merchant"] = isMerchant
	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	x := GenerateKey()
	checkTypeOf(x)

	tc, _ := json.Marshal(x.(*rsa.PublicKey))
	return token.SignedString(tc)
}
func checkTypeOf(x interface{}) {
	var reflectValue = reflect.ValueOf(x)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		f, _ := reflectType.FieldByName("Email")
		fmt.Println(f.Tag)
		fmt.Println(f.Tag.Lookup("test"))
		fmt.Println("nama      :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai     :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func RegenerateAccessToken(accessTokenString string) (string, error) {
	claims, err := getClaimsJWTCIAM(accessTokenString)
	if err != nil {
		return "", err
	}
	claims["TESTING"] = "TESTING AJA"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	x := GenerateKey()
	// checkTypeOf(x)

	tc, _ := json.MarshalIndent(x.(*rsa.PublicKey), "", " ")
	return token.SignedString(tc)

}
