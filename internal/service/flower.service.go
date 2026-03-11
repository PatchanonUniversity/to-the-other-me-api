package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"to-the-other-me/internal/model"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiService struct{}

func (service *GeminiService) GenerateFlower(ctx context.Context, req model.GeminiRequest) (*model.AIResponse, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	modelAI := client.GenerativeModel("gemini-2.5-flash")
	
	prompt := fmt.Sprintf(`
	คุณคือ 'Future Self' (ตัวตนในอนาคต) ที่มีความเมตตา
	จงประมวลผลข้อมูลของ %s
	ที่ชอบ %s
	และอยากทำ %s
	โดยตอนนี้รู้สึก %s
	ซึ่งตอนนี้กำลังไล่ตาม %s
	ชีวิตช่วงนี้คือ %s
	เขาเพิ่งวางความคาดหวังเรื่อง '%s'
	ทิ้งไว้ในตะกร้า
	เหตุผลต้องเชื่อมโยงกับ hobby, dream และ feeling
	จงเลือกดอกไม้ 1 ชนิดที่ห้ามใช้สัญลักษณ์ทางศาสนา
	ต้องเป็นดอกไม้ที่สะท้อน hobby, dream และ feeling โดยตรง
	ตอบกลับเป็น JSON เท่านั้น:
	{
	"name": "ชื่อดอกไม้",
	"meaning": "เหตุผลที่กินใจ โดยลงท้ายว่า 'ไม่ว่าผลลัพธ์จะเป็นยังไง ชั้นภูมิใจในตัวเธอเสมอนะ'"
	}
	`,
		req.UserName,
		req.UserHobby,
		req.UserDream,
		req.UserFeeling,
		req.UserChasing,
		req.UserLife,
		req.UserExpectation,
	)

	res, err := modelAI.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, err
	}

	part := res.Candidates[0].Content.Parts[0]
	text := fmt.Sprintf("%v", part)
	cleanJson := strings.Trim(text, " \n`json")

	var aiResult model.AIResponse
	if err := json.Unmarshal([]byte(cleanJson), &aiResult); err != nil {
		return nil, err
	}

	return &aiResult, nil
}