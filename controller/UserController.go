package controller

import (
	"gin-go-wp/common"
	"gin-go-wp/dto"
	"gin-go-wp/model"
	"gin-go-wp/response"
	"gin-go-wp/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	//第一种
	// var requestMap = make(map[string]string)
	// json.NewDecoder(ctx.Request.Body).Decode(&requestMap)

	//第二种
	// var requestUser = model.User{}
	// json.NewDecoder(ctx.Request.Body).Decode(&requestUser)

	//第三种
	var requestUser = model.User{}
	// json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&requestUser)
	//获取数据
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	// name := ctx.PostForm("name")
	// telephone := ctx.PostForm("telephone")
	// password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11号")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11号"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码为6位"})
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)
	//判断手机号是否存在
	if isTelephoneExists(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}
	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	db.Create(&newUser)
	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	response.Success(ctx, gin.H{"token": token}, "注册成功")
	// ctx.JSON(200, gin.H{
	// 	"msg": "注册成功",
	// })
}

func Login(ctx *gin.Context) {
	db := common.GetDB()
	var requestUser = model.User{}
	// json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&requestUser)
	//获取参数
	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11号"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码为6位"})
		return
	}
	//判断
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	//判断密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登陆成功",
	})
}

func isTelephoneExists(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return user.ID != 0
	}
	return false
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.TOUserDto(user.(model.User))}})
}

//https://www.cnblogs.com/davygeek/p/7995306.html  golang学习资料
