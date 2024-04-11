package providers

import (
	"encoding/json"
	"errors"

	"github.com/fedemiodo/Crabi-code-challenge/internal/model"
)

var PLDProvider = &ExternalProviderRouting{BaseUrl: "http://44.210.144.170", ContentType: "application/json"}

// returns bool = isUsedBlacklisted on response OK, otherwise error.
func ConsumePLDService(user *model.User) (bool, error) {
	// Adapt to provider contract
	reqUser := &PldServiceRequest{User: user}
	if reqBody, err := json.Marshal(reqUser); err != nil {
		return false, err
	} else {
		// request body ok to send
		if resp, err := PLDProvider.Post("/check-blacklist", reqBody); err != nil {
			// failure
			return false, err
		} else {
			defer resp.Body.Close()
			var res map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
				// decoding error
				return false, err
			} else {
				// Provider error response
				if resp := res["is_in_blacklist"]; resp == nil {
					errorResp, _ := json.Marshal(res)
					return false, errors.New(string(errorResp)) // not pretty but works
				} else {
					// Response ok - should have "is_in_blacklist"
					return res["is_in_blacklist"].(bool), nil
				}
			}
		}
	}
}

type PldServiceRequest struct {
	*model.User
}

func (userReq *PldServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}{
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Email:     userReq.Email,
	})
}
