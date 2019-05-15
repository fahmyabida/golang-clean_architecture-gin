package utils

import "github.com/gin-gonic/gin"

func ResponseList(res interface{}, length int, result gin.H, info ...interface{}) gin.H{
	if info != nil {
		result = gin.H{
			"data": res,
			"info": info,
			"count": length,
		}
	} else {
		result = gin.H{
			"data": res,
			"count": length,
		}
	}
	return result
}

func ResponseListWithOneInfo(res interface{}, length int, result gin.H, info interface{}) gin.H{
	if info != nil {
		result = gin.H{
			"data": res,
			"info": info,
			"count": length,
		}
	} else {
		result = gin.H{
			"data": res,
			"count": length,
		}
	}
	return result
}

func ResponseObject(res interface{}, result gin.H) gin.H{
	result = gin.H{
		"data": res,
	}
	return result
}