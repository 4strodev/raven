package injector

import (
	templateProviders "github.com/4strodev/raven/pkg/features/template/providers"
	sharedProviders "github.com/4strodev/raven/pkg/shared/providers"
	"github.com/spf13/afero"
)

type Injector struct {
	TemplateRepository *templateProviders.TemplateRepository
	FileSystem *sharedProviders.FileSystem
}

func NewInjector() Injector {
	fs := sharedProviders.NewFileSystem(afero.NewOsFs())
	return Injector{
		TemplateRepository: templateProviders.NewTemplateRepository(fs),
		FileSystem: fs,
	}
}

