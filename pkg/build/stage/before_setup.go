package stage

import (
	"github.com/flant/dapp/pkg/build/builder"
	"github.com/flant/dapp/pkg/config"
	"github.com/flant/dapp/pkg/util"
)

func GenerateBeforeSetupStage(dimgConfig config.DimgInterface, extra *builder.Extra) Interface {
	b := getBuilder(dimgConfig, extra)
	if b != nil && !b.IsBeforeSetupEmpty() {
		return newBeforeSetupStage(b)
	}

	return nil
}

func newBeforeSetupStage(builder builder.Builder) *BeforeSetupStage {
	s := &BeforeSetupStage{}
	s.UserStage = newUserStage(builder)
	return s
}

type BeforeSetupStage struct {
	*UserStage
}

func (s *BeforeSetupStage) Name() StageName {
	return BeforeSetup
}

func (s *BeforeSetupStage) GetContext(_ Conveyor) (string, error) {
	stageDependenciesChecksum, err := s.GetStageDependenciesChecksum(BeforeSetup)
	if err != nil {
		return "", err
	}

	return util.Sha256Hash(s.builder.BeforeSetupChecksum(), stageDependenciesChecksum), nil
}

func (s *BeforeSetupStage) PrepareImage(prevBuiltImage, image Image) error {
	if err := s.BaseStage.PrepareImage(prevBuiltImage, image); err != nil {
		return err
	}

	if err := s.builder.BeforeSetup(image.BuilderContainer()); err != nil {
		return err
	}

	return nil
}
