package infrared

import (
	"encoding/json"
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PostMatchingRemotesTokenReq struct {
	InfraredID string
	CategoryId int    `json:"category_id"`
	PreToken   string `json:"pre_token"`
	Code       string `json:"code"`
}

func (t *PostMatchingRemotesTokenReq) Method() string {
	return common.RequestPost
}

func (t *PostMatchingRemotesTokenReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/matching-remotes/token", t.InfraredID)
}

func (t *PostMatchingRemotesTokenReq) Body() []byte {
	reqBody, _ := json.Marshal(t)
	return reqBody
}

// PostmatchingRemotesToken sets infrared learning mode
func PostmatchingRemotesToken(infraredID string, categoryId int, preToken, code string) (*PostmatchingRemotesTokenResponse, error) {
	a := &PostMatchingRemotesTokenReq{
		InfraredID: infraredID,
		CategoryId: categoryId,
		PreToken:   preToken,
		Code:       code,
	}
	resp := &PostmatchingRemotesTokenResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type PostmatchingRemotesTokenResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		Token      string `json:"token"`
		ExpireTime int    `json:"expire_time"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
