package providers

import (
	"github.com/4strodev/raven/pkg/features/template/domain"
	"github.com/4strodev/raven/pkg/lib/monads/option"
	sharedProviders "github.com/4strodev/raven/pkg/shared/providers"
)

type TemplateRepository struct {
	fileSystem *sharedProviders.FileSystem
}

func NewTemplateRepository(fs *sharedProviders.FileSystem) *TemplateRepository {
	return &TemplateRepository{
		fileSystem: fs,
	}
}

func (r TemplateRepository) Delete() error {
	return nil
}

func (r TemplateRepository) Save() error {
	return nil
}

func (r TemplateRepository) Find() option.Option[[]domain.Template] {
	return option.None[[]domain.Template]()
}
