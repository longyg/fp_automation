package main

import (
	"fmt"
	"./nidd"
)

func main() {
	loginUrl := "http://esnidd053.emea.nsn-net.net/nidd/webapi/HomePage/UserLogin"
	//addProductUrl := "http://esnidd053.emea.nsn-net.net/nidd/webapi/Product/ProductAdd"
	//addPackageUrl := "http://esnidd053.emea.nsn-net.net/nidd/webapi/package/Import"
	importAlarmUrl := "http://esnidd053.emea.nsn-net.net/nidd/webapi/Alarm/Import"

	err := nidd.Login(loginUrl, "", "")
	if nil != err {
		fmt.Println("Login to NIDD failed:", err.Error())
		return
	}

	fmt.Println("Login to NIDD successful")

	//err = nidd.AddProduct(addProductUrl, "IMS", "CSCF1",
	//	"CSCF1 test", "NO",
	//	"NO", "CSCF TEST PRODUCT")
	//if nil != err {
	//	fmt.Println("Add product fail:", err)
	//}
	//
	//err = nidd.CreateNonKpiPackage(addPackageUrl, "IMS", "ALB", "alb_test", "18.0vnf", "PM", "Initial")
	//if nil != err {
	//	fmt.Println("Create package fail:", err)
	//	return
	//}
	//fmt.Println("Package is created succefully")

	//package id: 23713

	err = nidd.ImportAlarms(importAlarmUrl, "IMS", "CSCF", "23713", "true", "D:\\com.nsn.cscf.man.xml")
	if nil != err {
		fmt.Println("Import alarm fail:", err)
		return
	}
	fmt.Println("Alarm is imported succefully")
}








