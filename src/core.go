package twitter

import _ "net/http"
import _ "fmt"

var bearer string;

func SetBearerToken(token string) error {
  bearer = "Bearer "+token
}
func GetBearer() string {
  return bearer
}
