package handlers

import (
	"net/http"

	"github.com/erkylima/allGolangRithm/api/v1/algorithms/dto"
	"github.com/erkylima/allGolangRithm/internal/algorithms/commands"
	"github.com/erkylima/allGolangRithm/internal/algorithms/entity"
	"github.com/gin-gonic/gin"
)

func AddTwoNumbers(c *gin.Context) {
	twoNumbersEtty := new(dto.AddTwoNumbers)

	if err := c.ShouldBindJSON(&twoNumbersEtty); err != nil {
		c.Writer.Header().Add("X-Error-Message", err.Error())
		c.JSON(http.StatusBadRequest, "Error in payload. Verify your payload.")
	}
	twoNumbers := &commands.OnAddTwoNumbers{
		L1: &entity.ListNode{
			Val: twoNumbersEtty.Value1,
		},
		L2: &entity.ListNode{
			Val: twoNumbersEtty.Value2,
		},
	}

	result := twoNumbers.Stringfy()

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
