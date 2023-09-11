package controllers

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "io/ioutil"
    "strings"
)

func GetBuku(c *gin.Context) {
    url := "http://127.0.0.1:8000/api/buku"

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

func AddBuku(c *gin.Context) {
    var requestBody struct {
        Judul    string `json:"judul"`
        Tahun string `json:"tahun"`
    }

    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	url := "http://127.0.0.1:8000/api/buku/store"

	data := strings.NewReader(fmt.Sprintf(`{"judul": "%s", "tahun": "%s"}`, requestBody.Judul, requestBody.Tahun))

    req, err := http.NewRequest("POST", url, data)
    if err != nil {
        fmt.Println("Gagal membuat permintaan HTTP:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat permintaan HTTP"})
        return
    }
    req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("Gagal melakukan permintaan HTTP:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal melakukan permintaan HTTP"})
        return
    }
    defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
        fmt.Println("Gagal menambahkan data, status kode:", resp.StatusCode)
        c.JSON(http.StatusBadGateway, gin.H{"error": "Gagal menambahkan data"})
        return
    }

	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Gagal membaca respons:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca respons"})
        return
    }

    data_add := string(body)
    c.JSON(http.StatusCreated, gin.H{"data": data_add})
}

func GetBukuByID(c *gin.Context) {
    id := c.Param("id_buku")

    url := "http://127.0.0.1:8000/api/buku/" + id

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

func UpdateBuku(c *gin.Context) {
    id := c.Param("id_buku")

    var requestBody struct {
        Judul   string `json:"judul"`
        Tahun   string `json:"tahun"`
    }

    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	url := "http://127.0.0.1:8000/api/buku/" + id

    data := strings.NewReader(fmt.Sprintf(`{"judul": "%s", "tahun": "%s"}`, requestBody.Judul, requestBody.Tahun))

    req, err := http.NewRequest("PUT", url, data)
    if err != nil {
        fmt.Println("Gagal membuat permintaan HTTP:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat permintaan HTTP"})
        return
    }
    req.Header.Set("Content-Type", "application/json")

    req.Header.Set("Authorization", "Bearer "+token)

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("Gagal melakukan permintaan HTTP:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal melakukan permintaan HTTP"})
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Println("Gagal memperbarui data, status kode:", resp.StatusCode)
        c.JSON(http.StatusBadGateway, gin.H{"error": "Gagal memperbarui data"})
        return
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Gagal membaca respons:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca respons"})
        return
    }

    data_update := string(body)
    c.JSON(http.StatusOK, gin.H{"data": data_update})
}


func DeleteBuku(c *gin.Context){
    //
}