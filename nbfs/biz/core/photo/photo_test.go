/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package photo

import (
	"testing"
	//"fmt"
	//"github.com/disintegration/imaging"
	//"bytes"
	"github.com/airwide-code/airwide.datacenter/baselib/mysql_client"
	"github.com/airwide-code/airwide.datacenter/nbfs/biz/dal/dao"
	//"image"
)

func init()  {
	mysqlConfig := mysql_client.MySQLConfig{
		Name:   "immain",
		DSN:    "root:@/nebulaim?charset=utf8",
		Active: 5,
		Idle:   2,
	}
	mysql_client.InstallMysqlClientManager([]mysql_client.MySQLConfig{mysqlConfig})
	dao.InstallMysqlDAOManager(mysql_client.GetMysqlClientManager())
}

func TestResize(t *testing.T) {
	//id := int64(-8540733062663239681)
	//filePartsDO := dao.GetFilePartsDAO(dao.DB_MASTER).SelectFileParts(1, id)
	////fileDatas := []byte{}
	////for _, p := range filePartsDOList {
	////	fileDatas = append(fileDatas, p.Bytes...)
	////}
	//
	//// bufio.Reader{}
	//img, err := imaging.Decode(bytes.NewReader(fileDatas))
	//if err != nil {
	//	fmt.Printf("Decode error: {%v}\n", err)
	//	return
	//}
	//
	//imgSz := makeResizeInfo(img)
	//for i, sz := range sizeList {
	//	var dst *image.NRGBA
	//	if imgSz.isWidth {
	//		dst = imaging.Resize(img, sz, 0, imaging.Lanczos)
	//	} else {
	//		dst = imaging.Resize(img, 0, sz, imaging.Lanczos)
	//	}
	//
	//	err := imaging.Save(dst, fmt.Sprintf("/tmp/telegramd/%d.jpg", i))
	//	if err != nil {
	//		fmt.Printf("Save error: {%v}\n", err)
	//	}
	//}
}
