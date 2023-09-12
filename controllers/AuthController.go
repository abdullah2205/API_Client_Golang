package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
)

var token string

//register
func sendRegisterRequest(email, password, name string) (string, error) {
    url := "http://127.0.0.1:8000/api/register"

    // Membuat permintaan HTTP untuk registrasi
    data := strings.NewReader(fmt.Sprintf(`{"email": "%s", "password": "%s", "name": "%s"}`, email, password, name))
    req, err := http.NewRequest("POST", url, data)
    if err != nil {
        return "", err
    }
    req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // Membaca token akses dari respons jika registrasi berhasil
    if resp.StatusCode == http.StatusOK {
        bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return "", err
        }
        token := string(bodyBytes)
        return token, nil
    } else {
        return "", fmt.Errorf("Gagal registrasi, status kode: %d", resp.StatusCode)
    }
}

func Register(c *gin.Context) {
    var requestBody struct {
		Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	
    receivedToken, err := sendRegisterRequest(requestBody.Email, requestBody.Password, requestBody.Name)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal registrasi"})
        return
    }

	var responseMap map[string]interface{}
    if err := json.Unmarshal([]byte(receivedToken), &responseMap); err != nil { //uraikan respon token dalam bentuk json dari receivedToken
        fmt.Println("Gagal menguraikan JSON:", err) 	
        return
    }

    valueToken, found := responseMap["token"].(string)
    if !found {
        fmt.Println("Token tidak ditemukan dalam respons JSON")
        return
    }

    token = valueToken //isi nilai token
	
    c.JSON(http.StatusOK, gin.H{"pesan": "Registrasi berhasil", "token": receivedToken})
}
//end register

func sendLoginRequest(email, password string) (string, error) {
    url := "http://127.0.0.1:8000/api/login"

    // Membuat permintaan HTTP dengan kredensial
    data := strings.NewReader(fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password))
    req, err := http.NewRequest("POST", url, data)
    if err != nil {
        return "", err
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // Membaca token akses dari respons
    if resp.StatusCode == http.StatusOK {
        bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return "", err
        }
        token := string(bodyBytes)
        return token, nil
    } else {
        return "", fmt.Errorf("Gagal masuk, status kode: %d", resp.StatusCode)
    }
}

func Login(c *gin.Context) {
    var requestBody struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	
    receivedToken, err := sendLoginRequest(requestBody.Email, requestBody.Password) //pemanggilan fungsi sendLoginRequest
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Gagal masuk"})
        return
    }
    
    var responseMap map[string]interface{}
    if err := json.Unmarshal([]byte(receivedToken), &responseMap); err != nil { //uraikan respon token dalam bentuk json dari receivedToken
        fmt.Println("Gagal menguraikan JSON:", err) 	
        return
    }

    valueToken, found := responseMap["token"].(string)
    if !found {
        fmt.Println("Token tidak ditemukan dalam respons JSON")
        return
    }

    token = valueToken //isi nilai token
	
    c.JSON(http.StatusOK, gin.H{"pesan": "Berhasil Login go", "token": token})
}

func Logout(c *gin.Context) {
	url := "http://127.0.0.1:8000/api/logout" // Ganti dengan URL API Laravel Anda

    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        fmt.Println("Gagal membuat permintaan HTTP:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat permintaan HTTP"})
        return
    }

    req.Header.Set("Authorization", "Bearer "+ GetToken())

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("Gagal melakukan permintaan HTTP:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal melakukan permintaan HTTP"})
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("Gagal mengambil data, status kode:", resp.StatusCode)
        c.JSON(http.StatusBadGateway, gin.H{"error": "Gagal mengambil data"})
        return
    }

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Gagal membaca respons:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca respons"})
		return
	}

    token = ""

	c.JSON(http.StatusOK, gin.H{"data": string(body)})
}

func GetToken() string {
    return token
}

func GetUser(c *gin.Context) {
    url := "http://127.0.0.1:8000/api/user" // Ganti port dengan port API Laravel Anda

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Gagal membuat permintaan HTTP:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat permintaan HTTP"})
        return
    }

    req.Header.Set("Authorization", "Bearer "+ GetToken())

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("Gagal melakukan permintaan HTTP:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal melakukan permintaan HTTP"})
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("Gagal mengambil data, status kode:", resp.StatusCode)
        c.JSON(http.StatusBadGateway, gin.H{"error": "Gagal mengambil data"})
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Gagal membaca respons:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca respons"})
        return
    }

    data := string(body)
    c.JSON(http.StatusOK, gin.H{"data": data})
}
