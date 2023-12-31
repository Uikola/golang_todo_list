package app

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
	"todolist/internal"
	"todolist/internal/handler"
	"todolist/internal/model"
	"todolist/internal/repository"
	"todolist/internal/service"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("SSL_MODE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("failed to connect to db due to error %s", err)
	}
	err = db.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		logrus.Fatalf("failed to migrate due to error %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(internal.Server)
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error ocured while running http server: %s", err.Error())
		}
	}()

	logrus.Println("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("TodoApp shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error while shutting down: %s", err.Error())
	}
}
