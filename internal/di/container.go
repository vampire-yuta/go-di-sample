package di

import (
	domainRepos "go-di/internal/domain/repositories"
	"go-di/internal/infrastructure/handlers"
	infraRepos "go-di/internal/infrastructure/repositories"
	"go-di/internal/usecase"
)

// Container は依存性注入コンテナです
type Container struct {
	// Repositories
	UserRepository domainRepos.UserRepository

	// Services
	UserService *usecase.UserService

	// Handlers
	UserHandler *handlers.UserHandler
}

// NewContainer は新しいDIコンテナを作成し、依存関係を注入します
func NewContainer() *Container {
	container := &Container{}

	// リポジトリの初期化
	container.UserRepository = infraRepos.NewUserRepositoryImpl()

	// サービスの初期化（リポジトリを注入）
	container.UserService = usecase.NewUserService(container.UserRepository)

	// ハンドラーの初期化（サービスを注入）
	container.UserHandler = handlers.NewUserHandler(container.UserService)

	return container
}

// GetUserHandler はユーザーハンドラーを取得します
func (c *Container) GetUserHandler() *handlers.UserHandler {
	return c.UserHandler
}

// GetUserService はユーザーサービスを取得します
func (c *Container) GetUserService() *usecase.UserService {
	return c.UserService
}

// GetUserRepository はユーザーリポジトリを取得します
func (c *Container) GetUserRepository() domainRepos.UserRepository {
	return c.UserRepository
}
