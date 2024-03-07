package routes

import (
	"context"
	"go-auth-template/controllers"
	"go-auth-template/utils"
	"net/http"
)

func AddBaseRoutes(ctx context.Context, mux *http.ServeMux, logger *utils.Logger) {
	mux.Handle("/healthz", controllers.HandleHealthZ(logger))
	mux.Handle("/", controllers.HandleHome(ctx, logger))
}

//func AddRoutes(
//    mux *http.ServeMux,
//    logger *utils.Logger,
//    userController *controllers.UserController,
//) {
//    mux.Handle("/login", userController.Login(logger))
//    mux.Handle("/refreshToken", userController.RefreshToken(logger))
//    mux.Handle("/logout", userController.Logout(logger))
//    mux.Handle("/docs", middleware.VerifyToken(logger, userController.GetDocs(logger)))
//}
