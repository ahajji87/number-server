package usecase_test

import (
	"errors"
	"fmt"
	"number-server/app/domain/model"
	"number-server/app/usecase"
	"number-server/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func TestNumberUseCaseSuite(t *testing.T) {
	suite.Run(t, new(NumberUseCaseTestSuite))
}

type NumberUseCaseTestSuite struct {
	suite.Suite
	numberService *mocks.MockNumberService
	pipeline      chan *model.Number
	underTest     usecase.NumberUseCase
}

func (suite *NumberUseCaseTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()
	suite.pipeline = make(chan *model.Number, 1)
	suite.numberService = mocks.NewMockNumberService(mockCtrl)
	suite.underTest = usecase.NewNumberUseCase(suite.numberService, suite.pipeline)
}

func (suite *NumberUseCaseTestSuite) TestReadNumberSucces() {

	number := &model.Number{
		Value: "123456789",
	}
	suite.numberService.EXPECT().IsValidNumber(gomock.AssignableToTypeOf(number)).Return(true)

	err := suite.underTest.ReadNumber(number)
	result := <-suite.pipeline

	suite.NoError(err, "Shouldn't error")
	suite.NotNil(result, "should not be null")
	suite.Equal(number, result)
}

func (suite *NumberUseCaseTestSuite) TestReadNumberInvalidNumber() {
	number := &model.Number{
		Value: "12345678",
	}
	suite.numberService.EXPECT().IsValidNumber(gomock.AssignableToTypeOf(number)).Return(false)

	err := suite.underTest.ReadNumber(number)

	suite.Error(err, "Should return error")
	suite.Equal(err.Error(), "Invalid number")
}

func (suite *NumberUseCaseTestSuite) TestStore() {
	number := &model.Number{
		Value: "123123123",
	}

	suite.numberService.EXPECT().Store(gomock.AssignableToTypeOf(number)).Return(nil)

	err := suite.underTest.Store(number)

	suite.NoError(err, "Shouldn't error")
}

func (suite *NumberUseCaseTestSuite) TestStoreFail() {
	number := &model.Number{
		Value: "123123123",
	}

	fail := errors.New("something wrong happens")
	suite.numberService.EXPECT().Store(gomock.AssignableToTypeOf(number)).Return(fail).Times(1)

	err := suite.underTest.Store(number)

	suite.Error(err, "Should return error")
}

func (suite *NumberUseCaseTestSuite) TestGetReport() {
	n := &model.Report{
		Duplicats: 1,
		Adds:      1,
		Total:     10,
	}
	expected := fmt.Sprintf("%d new numbers, %d duplicated entries, %d total numbers", n.Adds, n.Duplicats, n.Total)
	suite.numberService.EXPECT().GetCounters().Return(n)

	suite.numberService.EXPECT().ResetCounters().Times(1)

	result := suite.underTest.GetReport()

	suite.Equal(expected, result)
}
