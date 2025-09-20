package controllers

import (
	"ecosnap/pkg/errs"
	"ecosnap/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"net/http"
	"strings"
)

// Функция для преобразования gRPC-кодов в HTTP-коды
func grpcCodeToHTTP(code codes.Code) int {
	switch code {
	case codes.InvalidArgument:
		return http.StatusBadRequest // 400
	case codes.Unauthenticated:
		return http.StatusUnauthorized // 401
	case codes.PermissionDenied:
		return http.StatusForbidden // 403
	case codes.NotFound:
		return http.StatusNotFound // 404
	case codes.AlreadyExists:
		return http.StatusConflict // 409
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests // 429
	case codes.FailedPrecondition, codes.Aborted:
		return http.StatusPreconditionFailed // 412
	case codes.Internal:
		return http.StatusInternalServerError // 500
	case codes.Unavailable:
		return http.StatusServiceUnavailable // 503
	default:
		return http.StatusBadRequest // 400 по умолчанию
	}
}

func parseGRPCError(err error) (codes.Code, string) {
	errMsg := err.Error()
	codePrefix := "code ="
	descPrefix := "desc ="

	// Ищем позицию начала "code ="
	idxCode := strings.Index(errMsg, codePrefix)
	if idxCode == -1 {
		return codes.Internal, errMsg
	}
	// Ищем позицию начала "desc ="
	idxDesc := strings.Index(errMsg, descPrefix)
	if idxDesc == -1 {
		return codes.Internal, errMsg
	}

	// Извлекаем код ошибки: берём строку между "code =" и "desc ="
	codeStr := strings.TrimSpace(errMsg[idxCode+len(codePrefix) : idxDesc])

	// Извлекаем описание ошибки: берём всё, что идёт после "desc ="
	descStr := strings.TrimSpace(errMsg[idxDesc+len(descPrefix):])

	// Ищем соответствующий код из пакета grpc/codes
	var grpcCode codes.Code = codes.Internal
	for c := codes.OK; c <= codes.Unauthenticated; c++ {
		if strings.EqualFold(c.String(), codeStr) {
			grpcCode = c
			break
		}
	}

	return grpcCode, descStr
}

// Обработка ошибок, которые приводят к статусу 400 (Bad Request)
func handleBadRequestErrors(err error) bool {
	return errors.Is(err, errs.ErrUsernameUniquenessFailed) ||
		errors.Is(err, errs.ErrUsernameIsRequired) ||
		errors.Is(err, errs.ErrPasswordIsRequired) ||
		errors.Is(err, errs.ErrFirstNameIsRequired) ||
		errors.Is(err, errs.ErrLastNameIsRequired) ||
		errors.Is(err, errs.ErrAppLoginIsRequired) ||
		errors.Is(err, errs.ErrEmailIsRequired) ||
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword) ||
		errors.Is(err, errs.ErrCategoryNameUniquenessFailed) ||
		errors.Is(err, errs.ErrOrderStatusNameUniquenessFailed) ||
		errors.Is(err, errs.ErrInvalidCredentials) ||
		errors.Is(err, errs.ErrOrderNotFound) ||
		errors.Is(err, errs.ErrInvalidStoreReviewID) ||
		errors.Is(err, errs.ErrPathParametrized) ||
		errors.Is(err, errs.ErrInvalidProductID) ||
		errors.Is(err, errs.ErrInvalidAddressID) ||
		errors.Is(err, errs.ErrInvalidOrderID) ||
		errors.Is(err, errs.ErrInvalidProductReviewID) ||
		errors.Is(err, errs.ErrInvalidAccountID) ||
		errors.Is(err, errs.ErrInvalidFeaturedProductID) ||
		errors.Is(err, errs.ErrInvalidAddressName) ||
		errors.Is(err, errs.ErrInvalidAccountNumber) ||
		errors.Is(err, errs.ErrAddressNameUniquenessFailed) ||
		errors.Is(err, errs.ErrAccountNumberUniquenessFailed) ||
		errors.Is(err, errs.ErrInvalidMinPrice) ||
		errors.Is(err, errs.ErrInvalidMaxPrice) ||
		errors.Is(err, errs.ErrInvalidPrice) ||
		errors.Is(err, errs.ErrInvalidID) ||
		errors.Is(err, errs.ErrInvalidPaymentID) ||
		errors.Is(err, errs.ErrInvalidRating) ||
		errors.Is(err, errs.ErrInvalidContent) ||
		errors.Is(err, errs.ErrInvalidComment) ||
		errors.Is(err, errs.ErrInvalidField) ||
		errors.Is(err, errs.ErrInvalidCategory) ||
		errors.Is(err, errs.ErrEmailIsEmpty) ||
		errors.Is(err, errs.ErrPasswordIsEmpty) ||
		errors.Is(err, errs.ErrUsernameIsEmpty) ||
		errors.Is(err, errs.ErrInvalidStore) ||
		errors.Is(err, errs.ErrInvalidStoreID) ||
		errors.Is(err, errs.ErrValidationFailed) ||
		errors.Is(err, errs.ErrStoreNameUniquenessFailed) ||
		errors.Is(err, errs.ErrNotEnoughProductInStock) ||
		errors.Is(err, errs.ErrDeleteFailed) ||
		errors.Is(err, errs.ErrInvalidTitle) ||
		errors.Is(err, errs.ErrOrderAlreadyPayed) ||
		errors.Is(err, errs.ErrInvalidDescription) ||
		errors.Is(err, errs.ErrInvalidAmount) ||
		errors.Is(err, errs.ErrInvalidQuantity) ||
		errors.Is(err, errs.ErrInsufficientFunds)
}

// Обработка ошибок, которые приводят к статусу 404 (Not Found)
func handleNotFoundErrors(err error) bool {
	return errors.Is(err, errs.ErrRecordNotFound) ||
		errors.Is(err, errs.ErrCategoryNotFound) ||
		errors.Is(err, errs.ErrOrderStatusNotFound) ||
		errors.Is(err, errs.ErrOrderNotFound) ||
		errors.Is(err, errs.ErrProductReviewNotFound) ||
		errors.Is(err, errs.ErrFileNotFound) ||
		errors.Is(err, errs.ErrProductNotFound) ||
		errors.Is(err, errs.ErrAddressNotFound) ||
		errors.Is(err, errs.ErrFeaturedProductNotFound) ||
		errors.Is(err, errs.ErrPaymentNotFound) ||
		errors.Is(err, errs.ErrAccountNotFound) ||
		errors.Is(err, errs.ErrStoreNotFound) ||
		errors.Is(err, errs.ErrStoreReviewNotFound)
}

// Обработка ошибок, которые приводят к статусу 401 (Unauthorized)
func handleUnauthorizedErrors(err error) bool {
	return errors.Is(err, errs.ErrInvalidToken) ||
		errors.Is(err, errs.ErrUnauthorized) ||
		errors.Is(err, errs.ErrRefreshTokenExpired)
}

// HandleError Основная функция обработки ошибок
func HandleError(c *gin.Context, err error) {
	if handleBadRequestErrors(err) {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	} else if errors.Is(err, errs.ErrPermissionDenied) {
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))
	} else if handleNotFoundErrors(err) {
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	} else if errors.Is(err, errs.ErrFetchingProducts) {
		c.JSON(http.StatusInternalServerError, newErrorResponse(err.Error()))
	} else if errors.Is(err, errs.WarningNoProductsFound) {
		c.JSON(http.StatusOK, gin.H{"message": errs.WarningNoProductsFound.Error()})
	} else if handleUnauthorizedErrors(err) {
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))
	} else {
		logger.Error.Printf("Err: %s", err)
		c.JSON(http.StatusInternalServerError, newErrorResponse(errs.ErrSomethingWentWrong.Error()))
	}
}
