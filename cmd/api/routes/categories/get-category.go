package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"

	c_repository "github.com/PedroVMB/GO-CATEGORIES-MCSV/internal/categories/repository"
	c_use_cases "github.com/PedroVMB/GO-CATEGORIES-MCSV/internal/categories/use-cases"
	"github.com/PedroVMB/GO-CATEGORIES-MCSV/pkg/utils"
)

func getCategory(c *gin.Context, repository *c_repository.CategoryRepository) {
	id, err := utils.StringToUint(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	use_case := c_use_cases.NewGetCategoryUseCase(repository)
	categories, err := use_case.Execute(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "result": categories})
}
