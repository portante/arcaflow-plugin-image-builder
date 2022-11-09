package images

import (
	"fmt"
	"github.com/arcalot/arcaflow-plugin-image-builder/internal/dto"
	"github.com/arcalot/arcaflow-plugin-image-builder/mocks/ce_client"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.arcalot.io/log"
	"testing"
)

func TestBuildImage(t *testing.T) {
	logger := log.NewLogger(log.LevelInfo, log.NewNOOPLogger())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cec := mocks.NewMockContainerEngineClient(ctrl)
	cec.EXPECT().
		Build("use", "the", []string{"forks"}).
		Return(nil).
		Times(1)
	assert.Nil(t, BuildImage(true, true, cec, "use", "the", "forks", logger))
}

func TestPushImage(t *testing.T) {
	logger := log.NewLogger(log.LevelInfo, log.NewNOOPLogger())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cec := mocks.NewMockContainerEngineClient(ctrl)
	rg1 := dto.Registry{
		Url:       "reg1.io",
		Username:  "user1",
		Password:  "secret1",
		Namespace: "allyourbases",
	}
	image_name := "usethe"
	image_tag := "forks"

	destination := fmt.Sprintf("%s/%s/%s:%s", rg1.Url, rg1.Namespace, image_name, image_tag)
	cec.EXPECT().
		Tag(fmt.Sprintf("%s:%s", image_name, image_tag), destination).
		Return(nil).
		Times(1)
	cec.EXPECT().
		Push(destination, rg1.Username, rg1.Password, rg1.Url).
		Return(nil).
		Times(1)
	assert.Nil(t, PushImage(true, true, true, cec, image_name, image_tag,
		rg1.Username, rg1.Password, rg1.Url, rg1.Namespace, logger))
}
