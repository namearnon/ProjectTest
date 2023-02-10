package service

import (
	"ProjectTest/middlewares/core"
	"ProjectTest/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type BeerService struct {
	ctx core.IContext
}

func BeerSer(ctx core.IContext) *BeerService {
	return &BeerService{ctx: ctx}
}

func (s *BeerService) GetBeer(postData *model.PostData) []map[string]interface{} {
	db, err := s.ctx.DB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	var base64Encoding string
	var mimeType string
	var bytes []byte
	result := []map[string]interface{}{}
	data := []model.Beer{}
	db.Raw("SELECT * FROM beer Where replace(beer_name, ' ', '') LIKE CASE WHEN length('" + strings.Replace(postData.BeerName, " ", "", -1) + "') > 0 THEN '%" + strings.Replace(postData.BeerName, " ", "", -1) + "%' ELSE replace(beer_name, ' ', '') END ").Scan(&data)
	for _, v := range data {
		bytes, err = ioutil.ReadFile("./images/" + strconv.Itoa(v.ID) + ".png")
		if err != nil {
			log.Fatal(err)
		}
		mimeType = http.DetectContentType(bytes)

		// Prepend the appropriate URI scheme header depending

		// on the MIME type
		switch mimeType {
		case "image/jpeg":
			base64Encoding += "data:image/jpeg;base64,"
		case "image/png":
			base64Encoding += "data:image/png;base64,"
		}

		// Append the base64 encoded output
		base64Encoding += model.ToBase64(bytes)
		result = append(result, map[string]interface{}{
			"id":    v.ID,
			"beer":  v.BeerName,
			"type":  v.BeerType,
			"desc":  v.BeerDesc,
			"image": base64Encoding,
		})
		base64Encoding = ""

		// result = append(result, map[string]interface{}{"beer" : })
	}
	return result
}

// data:image/png;base64,i
func (s *BeerService) PostBeer(postData *model.GetBeerData) map[string]interface{} {
	file := postData.BeerImage

	db, err := s.ctx.DB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	if file != nil {
		src, err := file.Open()
		if err != nil {
			log.Println(err)
		}
		check := strings.Split(file.Filename, "/")
		if len(check) > 1 {
			return map[string]interface{}{
				"message": "ชื่อรูปห้ามมีสัญญลักษณ์",
			}
		}
		expension := strings.Split(postData.BeerImage.Filename, ".")[1]
		Beer := model.Beer{
			BeerName:  postData.BeerName,
			BeerType:  postData.BeerType,
			BeerDesc:  postData.BeerDesc,
			BearImage: postData.BeerImage.Filename}
		Log := model.Log{
			LogMethod: "POST",
			LogDesc:   "ซื่อเบียร์ : " + postData.BeerName + "ประเภท : " + postData.BeerType + " รายละเอียด : " + postData.BeerDesc,
		}

		db.Create(&Beer)
		db.Create(&Log)
		dst, err := os.Create("./images/" + strconv.Itoa(Beer.ID) + "." + expension)
		if io.Copy(dst, src); err != nil {
			log.Println(err)
		}
		return map[string]interface{}{
			"message": "success",
		}
	}
	return map[string]interface{}{
		"message": "โปรดใส่รูปภาพด้วย",
	}

}

func (s *BeerService) PutBeer(postData *model.GetBeerData) map[string]interface{} {
	file := postData.BeerImage
	db, err := s.ctx.DB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	if file != nil {
		src, err := file.Open()
		if err != nil {
			log.Println(err)
		}
		check := strings.Split(file.Filename, "/")
		if len(check) > 1 {
			return map[string]interface{}{
				"message": "ชื่อรูปห้ามมีสัญญลักษณ์",
			}
		}

		expension := strings.Split(postData.BeerImage.Filename, ".")[1]

		bytes, err := os.OpenFile("./images/"+postData.ID+"."+expension, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)

		if io.Copy(bytes, src); err != nil {
			log.Println(err)
			return map[string]interface{}{
				"message": 500,
			}
		}
	}
	bl := db.Exec("Update beer set beer_name = CASE WHEN length('"+postData.BeerName+"') > 0 Then '"+postData.BeerName+"' ELSE beer_name END, "+
		"beer_type = CASE WHEN length('"+postData.BeerType+"') > 0 Then '"+postData.BeerType+"' ELSE beer_type END,"+
		"beer_desc = CASE WHEN length('"+postData.BeerDesc+"') > 0 Then '"+postData.BeerDesc+"' ELSE beer_desc END Where id = ?", postData.ID)
	log.Println(bl.Error)
	Log := model.Log{
		LogMethod: "PUT",
		LogDesc:   "ซื่อเบียร์ : " + postData.BeerName + "ประเภท : " + postData.BeerType + " รายละเอียด : " + postData.BeerDesc,
	}
	db.Create(&Log)
	return map[string]interface{}{
		"message": "success",
	}
}

func (s *BeerService) DeleteBeer(postData *model.GetBeerData) map[string]interface{} {
	db, err := s.ctx.DB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	Beer := model.Beer{}
	db.Raw("SELECT bear_image FROM beer Where id = ?", postData.ID).Scan(&Beer)
	e := os.Remove("./images/" + Beer.BearImage)
	if e != nil {
		log.Fatal(e)
	}
	db.Exec("DELETE FROM beer Where id = ?", postData.ID)

	return map[string]interface{}{"message": "success"}
}
