package controller

import (
	"giftCard/model"
	"giftCard/response"
	"giftCard/utils"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

var WhatsAppIdList = []string{
	"6281212121212",
	"6281313131313",
	"6281414141414",
	// Add more numbers here
}

func AllContact(ctx *gin.Context) {
	var contacts []model.Contact
	utils.Db.Find(&contacts)
	response.Success(ctx, contacts, "success")
}

func AddContact(ctx *gin.Context) {
	var contact model.Contact
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		response.Fail(ctx, nil, err.Error())
		return
	}
	utils.Db.Create(&contact)
	response.Success(ctx, nil, "success")
}

func DeleteContact(ctx *gin.Context) {
	id := ctx.Query("contactId")
	var contact model.Contact
	utils.Db.Where("id = ?", id).First(&contact)
	if contact.ID == 0 {
		response.Fail(ctx, nil, "no contact")
		return
	} else {
		utils.Db.Delete(&contact)
		response.Success(ctx, nil, "success")
	}
}

func GetCategories(ctx *gin.Context) {
	var categories []model.Category
	utils.Db.Find(&categories)
	response.Success(ctx, categories, "success")
}

func CreateCategory(ctx *gin.Context) {
	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		response.Fail(ctx, nil, err.Error())
		return
	}
	utils.Db.Create(&category)
	response.Success(ctx, nil, "success")
}

func DeleteCategory(ctx *gin.Context) {
	id := ctx.Query("categoryId")
	var category model.Category
	utils.Db.Where("id = ?", id).First(&category)
	if category.ID == 0 {
		response.Fail(ctx, nil, "no category")
		return
	} else {
		utils.Db.Delete(&category)
		response.Success(ctx, nil, "success")
	}
}

func UpdateCategoryById(ctx *gin.Context) {
	id := ctx.Query("categoryId")
	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		response.Fail(ctx, nil, err.Error())
		return
	} else {
		utils.Db.Model(&category).Where("id = ?", id).Updates(&category)
		response.Success(ctx, nil, "success")
	}
}

func GetCardsByCategory(ctx *gin.Context) {
	categoryId := ctx.Query("categoryId") // c.Request.URL.Query().Get("lastname") 的一种快捷方式

	cards := make([]model.Card, 0)
	utils.Db.Where("category_id = ?", categoryId).Find(&cards)
	if len(cards) == 0 {
		response.Fail(ctx, nil, "no cards")
		return
	}
	response.Success(ctx, cards, "success")
}

func CreateCard(ctx *gin.Context) {
	var card model.Card
	if err := ctx.ShouldBindJSON(&card); err != nil {
		response.Fail(ctx, nil, err.Error())
		return
	}
	utils.Db.Create(&card)
	response.Success(ctx, nil, "success")
}

func DeleteCard(ctx *gin.Context) {
	id := ctx.Query("cardId")
	var card model.Card
	utils.Db.Where("id = ?", id).First(&card)
	if card.ID == 0 {
		response.Fail(ctx, nil, "no card")
		return
	} else {
		utils.Db.Delete(&card)
		response.Success(ctx, nil, "success")
	}
}

func UpdateCardById(ctx *gin.Context) {
	Id := ctx.Query("cardId")
	var card model.Card
	if err := ctx.ShouldBindJSON(&card); err != nil {
		response.Fail(ctx, nil, err.Error())
		return
	} else {
		utils.Db.Model(&card).Where("id = ?", Id).Updates(card)
		response.Success(ctx, nil, "success")
	}
}

func SellCard(ctx *gin.Context) {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	randomIndex := rand.Intn(len(WhatsAppIdList))
	response.Success(ctx, gin.H{
		"whatsappId": WhatsAppIdList[randomIndex],
	}, "success")
}
