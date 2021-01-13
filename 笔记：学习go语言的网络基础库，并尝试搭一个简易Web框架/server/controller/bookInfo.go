package controller

import (
	"encoding/json"
	"log"
	"server/jun"
)

type bookinfo struct {
	Id        string
	Name      string
	Price     string
	Info      string
	Publisher string
}

var bookInfoList []bookinfo

func AddBookInfo(c *jun.Context) {
	req := c.Request

	id := req.PostFormValue("id")
	name := req.PostFormValue("name")
	price := req.PostFormValue("price")
	info := req.PostFormValue("info")
	publisher := req.PostFormValue("publisher")

	book := bookinfo{
		Id:        id,
		Name:      name,
		Price:     price,
		Info:      info,
		Publisher: publisher,
	}

	bookInfoList = append(bookInfoList, book)

	c.Writer.Write(getBookInfoListBytes())
}

func DeleteBookInfo(c *jun.Context) {
	id := c.Request.PostFormValue("id")
	if id == "" {
		c.Writer.Write([]byte("id can't be null"))
		return
	}
	for i, b := range bookInfoList {
		if b.Id == id {
			removeFromBookInfoList(i)
		}
	}
	c.Writer.Write(getBookInfoListBytes())
}

func removeFromBookInfoList(index int) {
	bookInfoList[index] = bookInfoList[len(bookInfoList)-1]
	bookInfoList = bookInfoList[:len(bookInfoList)-1]
}

func UpdateBookInfo(c *jun.Context) {
	req := c.Request
	w := c.Writer

	id := req.PostFormValue("id")
	if id == "" {
		w.Write([]byte("id can't be null"))
		return
	}
	name := req.PostFormValue("name")
	price := req.PostFormValue("price")
	info := req.PostFormValue("info")
	publisher := req.PostFormValue("publisher")

	for i, b := range bookInfoList {
		if b.Id == id {
			bookInfoList[i].Name = name
			bookInfoList[i].Price = price
			bookInfoList[i].Info = info
			bookInfoList[i].Publisher = publisher
		}
	}
	w.Write(getBookInfoListBytes())
}

func GetBookInfo(c *jun.Context) {
	req := c.Request
	w := c.Writer

	id, ok := req.URL.Query()["id"]

	if !ok || id[0] == "" {
		w.Write([]byte("id can't be null"))
		return
	}
	var result bookinfo
	for _, b := range bookInfoList {
		if b.Id == id[0] {
			result = b
		}
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(bytes)
}

func getBookInfoListBytes() []byte {
	//show top ten bookInfo
	var result []bookinfo
	for i, _ := range bookInfoList {
		if i >= 10 {
			break
		}
		result = append(result, bookInfoList[i])
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
