package services

import (
	"errors"

	"github.com/kacperkg/vellum/internal/dto"
	"github.com/kacperkg/vellum/internal/models"
	"github.com/kacperkg/vellum/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repositories.UserRepository
	userSettingsRepository repositories.UserSettingsRepository
}

func NewUserService(userRepository repositories.UserRepository, userSettingsRepository repositories.UserSettingsRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
		userSettingsRepository: userSettingsRepository,
	}
}

func (s *UserService) Register(req dto.RegisterRequest) error {
	// Check if email already exists
	user, err := s.userRepository.FindByEmail(req.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("user with this email already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user = &models.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PasswordHash: string(hashedPassword),
	}

	err = s.userRepository.Create(user)
	if err != nil {
		return err
	}

	// Create default user settings
	defaultSettings := &models.UserSettings{
		UserID: user.ID,
		Currency: models.GBP,
		DateFormat: models.DDMMYYYY,
		NumberFormat: models.CommaDecimal,
		WeekStartsOn: models.Monday,
	}

	err = s.userSettingsRepository.Create(defaultSettings)
	if err != nil {
		return err
	}

	return nil
}