package thirdcloud

import (
	"fmt"
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func init() {
	common.SetTestEnv()
}

var (
	ThirdCloudDeviceId = common.Ed.TestDataThirdCloudDeviceId
	CountryCode        = common.Ed.TestDataCountryCode
	AppSchema          = common.Ed.TestDataAppSchema
	TuyaUsername       = common.Ed.TestDataTuyaUsername
	TuyaProductId      = common.Ed.TestDataTuyaProductId
	ParentDeviceId     = common.Ed.TestDataParentDeviceId
)

func TestPostDevicesBind(t *testing.T) {
	resp, err := PostDevicesBind(ThirdCloudDeviceId, CountryCode, AppSchema, TuyaUsername, TuyaProductId, ParentDeviceId)
	if err != nil {
		t.Errorf("PostDevicesBind req has err:%v,resp:%v \n", err, resp)
	}
	fmt.Printf("resp:%v \n", resp)
	if !resp.Success {
		t.Errorf("resp:%v is not success \n", resp)
	}
}
