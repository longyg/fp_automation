package nidd

import (
	"net/url"
	"../httputil"
)

func CreateNonKpiPackage(postUrl string, domain string, product string, packageName string, packageVersion string, packageType string,
	packageStatus string) error {

	return createNonKpiPackage0(postUrl, domain, product, packageName,
		packageVersion, packageType, "", packageStatus,
		"", "", "", "",
		"", "",
		"", "")
}

func createNonKpiPackage0(postUrl string, domain string, product string, packageName string, packageVersion string, packageType string,
	packageSubType string, packageStatus string, description string, moCTree string, adaptationId string, parentPackage string, relatedContainer string,
	relatedContainerVersion string, counterAbbrUnique string, counterPIIDunique string) error {

	values := make(url.Values)
	values.Set("domain", domain)
	values.Set("product", product)
	values.Set("packageName", packageName)
	values.Set("packageVersion", packageVersion)
	values.Set("packageType", packageType)
	values.Set("packageSubType", packageSubType)

	values.Set("packageStatus", packageStatus)
	values.Set("description", description)
	values.Set("moCTree", moCTree)
	values.Set("adaptationId", adaptationId)
	values.Set("parentPackage", parentPackage)
	values.Set("relatedContainer", relatedContainer)

	values.Set("relatedContainerVersion", relatedContainerVersion)
	values.Set("counterAbbrUnique", counterAbbrUnique)
	values.Set("counterPIIDunique", counterPIIDunique)

	return httputil.PostRequest(postUrl, values)
}
