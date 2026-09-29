package main

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PrescriptionItem struct {
	DrugID       string  `json:"drug_id"`
	DrugName     string  `json:"drug_name"`
	Dosage       string  `json:"dosage"`
	Frequency    string  `json:"frequency"`
	Duration     string  `json:"duration"`
	Quantity     int     `json:"quantity"`
	UnitPrice    float64 `json:"unit_price"`
	Instructions string  `json:"instructions"`
}

type Prescription struct {
	ID                 string             `json:"id"`
	RxNumber           string             `json:"rx_number"`
	PatientID          string             `json:"patient_id"`
	PatientName        string             `json:"patient_name"`
	PatientPhone       string             `json:"patient_phone"`
	DeliveryAddress    string             `json:"delivery_address"`
	UploadedAt         time.Time          `json:"uploaded_at"`
	ImageURL           string             `json:"image_url"`
	Status             string             `json:"status"` // submitted, under_review, verified_quoted, ready
	AssignedPharmacyID string             `json:"assigned_pharmacy_id"`
	PharmacistNotes    string             `json:"pharmacist_notes"`
	TotalAmount        float64            `json:"total_amount"`
	Items              []PrescriptionItem `json:"items"`
}

var prescriptions = make(map[string]Prescription)

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "pharmrx-go-core"})
	})

	// Upload prescription by patient
	r.POST("/api/prescriptions/upload", func(c *gin.Context) {
		var req Prescription
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.ID = uuid.New().String()
		req.RxNumber = "RX-" + time.Now().Format("20060102") + "-" + req.ID[:5]
		req.UploadedAt = time.Now()
		req.Status = "under_review"
		prescriptions[req.ID] = req
		c.JSON(http.StatusCreated, req)
	})

	// Get prescriptions for a specific pharmacy
	r.GET("/api/pharmacies/:pharmacyId/prescriptions", func(c *gin.Context) {
		pharmacyID := c.Param("pharmacyId")
		var result []Prescription
		for _, p := range prescriptions {
			if p.AssignedPharmacyID == pharmacyID || p.AssignedPharmacyID == "" {
				result = append(result, p)
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// Pharmacist review and approval
	r.PUT("/api/prescriptions/:id/verify", func(c *gin.Context) {
		id := c.Param("id")
		p, exists := prescriptions[id]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Prescription not found"})
			return
		}

		type VerifyPayload struct {
			Status          string             `json:"status"`
			PharmacistNotes string             `json:"pharmacist_notes"`
			Items           []PrescriptionItem `json:"items"`
			TotalAmount     float64            `json:"total_amount"`
		}
		var payload VerifyPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		p.Status = payload.Status
		p.PharmacistNotes = payload.PharmacistNotes
		p.Items = payload.Items
		p.TotalAmount = payload.TotalAmount
		prescriptions[id] = p

		c.JSON(http.StatusOK, p)
	})

	r.Run(":8080")
}
