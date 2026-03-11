package handler

import (
	"net/http"
	"to-the-other-me/internal/model"
	"to-the-other-me/internal/service"

	"github.com/gin-gonic/gin"
)

type GeminiHandler struct {
	Service *service.GeminiService
}

func (geminiHandler *GeminiHandler) HandleGemini(ctx *gin.Context) {
	var req model.GeminiRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := geminiHandler.Service.GenerateFlower(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusOK, model.AIResponse{
			Name: "ดอกเดซี่",
			Meaning: "ดอกเดซี่เป็นสัญลักษณ์ของชีวิตที่เรียบง่ายแต่มั่นคง เติบโตได้แม้ในพื้นที่ธรรมดา ไม่ต้องโดดเด่นหรือสมบูรณ์แบบเพื่อจะมีคุณค่า สื่อถึงการเริ่มต้นใหม่ ความบริสุทธิ์ใจ และความหวังเล็ก ๆ ที่คอยย้ำเตือนว่า ต่อให้โลกจะวุ่นวายหรือโหดร้ายเพียงใด เราก็ยังสามารถยืนหยัด เป็นตัวของตัวเอง และค่อย ๆ เบ่งบานในจังหวะของเราได้เสมอ ไม่ว่าผลลัพธ์จะเป็นยังไง ชั้นภูมิใจในตัวเธอเสมอนะ",
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}