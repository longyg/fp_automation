package nidd

import (
	"../httputil"
)

func ImportAlarms(postUrl string, domain string, product string, packageId string, newAlarm string, importFile string) error {
	values := make(map[string]string)
	values["domain"] = domain
	values["product"] = product
	values["package"] = packageId
	values["newAlarm"] = newAlarm

	return httputil.UploadFile(postUrl, values, "importFile", importFile)
}