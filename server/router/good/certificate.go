package good

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CertificateRouter struct{}

func (s *CertificateRouter) InitCertificateRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	certificateRouter := Router.Group("certificate").Use(middleware.OperationRecord())
	certificateRouterWithoutRecord := Router.Group("certificate")
	certificateRouterWithoutAuth := PublicRouter.Group("certificate")
	{
		certificateRouter.POST("createCertificate", certificateApi.CreateCertificate)
		certificateRouter.DELETE("deleteCertificate", certificateApi.DeleteCertificate)
		certificateRouter.DELETE("deleteCertificateByIds", certificateApi.DeleteCertificateByIds)
		certificateRouter.PUT("updateCertificate", certificateApi.UpdateCertificate)
	}
	{
		certificateRouterWithoutRecord.GET("findCertificate", certificateApi.FindCertificate)
		certificateRouterWithoutRecord.GET("getCertificateList", certificateApi.GetCertificateList)
	}
	{
		certificateRouterWithoutAuth.GET("getCertificatePublic", certificateApi.GetCertificatePublic)
		certificateRouterWithoutAuth.POST("count", certificateApi.Equcount)
	}
}
