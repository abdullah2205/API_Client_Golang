package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
	"io/ioutil"
	"encoding/json"
)

var token string

func sendLoginRequest(email, password string) (string, error) {
    // URL endpoint login pada server API
    url := "http://127.0.0.1:8000/api/login"

    // Membuat permintaan HTTP dengan kredensial
    data := strings.NewReader(fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password))
    req, err := http.NewRequest("POST", url, data)
    if err != nil {
        return "", err
    }
    req.Header.Set("Content-Type", "application/json")

    // Melakukan permintaan HTTP
    client := &http.Client{}
    resp, err := client.Do(req)
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

func login(c *gin.Context) {
		
    // Ambil kredensial pengguna dari permintaan (dalam format JSON)
    var requestBody struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Kirim kredensial ke server API
    receivedToken, err := sendLoginRequest(requestBody.Email, requestBody.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Gagal masuk"})
        return
    }
    //uraikan respon token dalam bentuk json
    var responseMap map[string]interface{}
    if err := json.Unmarshal([]byte(receivedToken), &responseMap); err != nil {
        fmt.Println("Gagal menguraikan JSON:", err) 	
        return
    }
    //jika tidak ada
    getToken, found := responseMap["token"].(string)
    if !found {
        fmt.Println("Token tidak ditemukan dalam respons JSON")
        return
    }

    token = getToken
    // Jika kredensial valid, kirimkan token akses sebagai respons
    c.JSON(http.StatusOK, gin.H{"pesan": "Berhasil Login go", "token": token})
}

func getUser(c *gin.Context) {
    url := "http://127.0.0.1:8000/api/user" // Ganti port dengan port API Laravel Anda

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Gagal membuat permintaan HTTP:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat permintaan HTTP"})
        return
    }

    req.Header.Set("Authorization", "Bearer "+token)

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

func getBuku(c *gin.Context) {
    url := "http://127.0.0.1:8000/api/buku"

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Gagal membuat permintaan HTTP:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat permintaan HTTP"})
        return
    }

    req.Header.Set("Authorization", "Bearer "+token)

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

func addBuku(c *gin.Context){
    //
}

func getBukuByID(c *gin.Context) {
    // Mendapatkan ID dari parameter URL
    id := c.Param("id_buku")

    url := "http://127.0.0.1:8000/api/buku/" + id

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Gagal membuat permintaan HTTP:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat permintaan HTTP"})
        return
    }

    req.Header.Set("Authorization", "Bearer "+token)

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

func updateBuku(c *gin.Context){
    //
}

func logout(c *gin.Context) {
	url := "http://127.0.0.1:8000/api/logout" // Ganti dengan URL API Laravel Anda

    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        fmt.Println("Gagal membuat permintaan HTTP:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat permintaan HTTP"})
        return
    }

    req.Header.Set("Authorization", "Bearer "+token)

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

func main() {
    router := gin.Default()
    router.POST("/api/login", login)

	router.GET("/api/user", getUser)

	router.GET("/api/buku", getBuku)

    router.POST("/api/buku/add", addBuku)

    router.GET("/api/buku/:id_buku", getBukuByID)

    router.PUT("/api/buku/:id_buku", updateBuku)

	router.POST("/api/logout", logout)

    router.Run(":8090")
}
