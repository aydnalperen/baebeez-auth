package controller

import (
	"baebeez-auth/models"
	"baebeez-auth/utils"
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Register(ctx *gin.Context) {
	var input models.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.UserAuth

	user.Mail = input.Mail
	user.Uid = uuid.NewString()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	user.Password = string(hashedPassword) // user creation is done

	_, err = user.SaveUserAuth()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	verifCode := new(models.VerifCode)

	verifCode.Uid = user.Uid
	verifCode.VerifCode = models.CreateVerifCode(6)

	verifCode.SaveVerifCode()

	SendMail(ctx, user.Mail, verifCode.VerifCode)
	ctx.JSON(http.StatusOK, gin.H{"message": "registered!", "mail": user.Mail, "password": user.Password})
}
func SaveProfile(ctx *gin.Context) {
	var input models.ProfileInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	user.Bio = input.Bio
	user.Name = input.Name
	user.Major = input.Major
	user.Photo = input.Photo
	user.Year = input.Year
	user.Uid = uuid.NewString()

	_, err := user.SaveUser()

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "validated!"})

}
func Login(ctx *gin.Context) {
	var input models.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
	}
	token, err := models.LoginCheck(input.Mail, input.Password, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func Logout(ctx *gin.Context) {

}
func CurrentUser(ctx *gin.Context) {
	user_id, err := utils.ExtractTokenUID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := models.GetUserByUid(user_id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func SendMail(ctx *gin.Context, receiver string, verifcode string) {
	data := map[string]string{
		"email":    os.Getenv("SENDER_EMAIL"),
		"password": os.Getenv("EMAIL_PASSWORD"),
	}

	jsonValue, _ := json.Marshal(data)

	resp, err := http.Post(os.Getenv("GET_TOKEN_ADDR"), "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Error!"})
		return
	}

	var res map[string]string

	json.NewDecoder(resp.Body).Decode(&res)

	values := map[string]string{
		"senderProfileId":      os.Getenv("SENDER_PROFILE_ID"),
		"receiverEmailAddress": receiver,
		"subject":              "Baebeez Verification",
		"content":              "Your Verification Code : " + verifcode,
	}
	jsonValue, _ = json.Marshal(values)

	bearer := "Bearer " + res["tokenValue"]

	url := os.Getenv("TRANSACTION_ADDR") + res["accountId"] + "/transactional-email"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", bearer)

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	second_res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer second_res.Body.Close()

	if second_res.StatusCode != http.StatusCreated {
		panic(second_res.Status)
	}
}
