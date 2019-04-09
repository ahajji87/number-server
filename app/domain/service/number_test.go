package service_test

import (
	"errors"
	"number-server/app/domain/model"
	"number-server/app/domain/service"
	"number-server/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

func TestNumberServiceSuite(t *testing.T) {
	suite.Run(t, new(NumberServiceTestSuite))
}

type NumberServiceTestSuite struct {
	suite.Suite
	numberRepository *mocks.MockNumberRepository
	underTest        service.NumberService
}

func (suite *NumberServiceTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()
	suite.numberRepository = mocks.NewMockNumberRepository(mockCtrl)
	suite.underTest = service.NewNumberService(suite.numberRepository)
}

func (suite *NumberServiceTestSuite) TestStoreNewNumber() {
	number := &model.Number{
		Value: "123456789",
	}

	suite.numberRepository.EXPECT().Exists(gomock.AssignableToTypeOf(number)).Return(false, nil).Times(1)

	suite.numberRepository.EXPECT().Save(gomock.AssignableToTypeOf(number)).Return(nil).Times(1)

	err := suite.underTest.Store(number)

	suite.NoError(err, "Shouldn't error")
}

func (suite *NumberServiceTestSuite) TestStoreDuplicatedNumber() {
	number := &model.Number{
		Value: "123456789",
	}

	suite.numberRepository.EXPECT().Exists(gomock.AssignableToTypeOf(number)).Return(true, nil).Times(1)

	suite.numberRepository.EXPECT().Save(gomock.AssignableToTypeOf(number)).Return(nil).Times(0)

	err := suite.underTest.Store(number)

	suite.NoError(err, "Shouldn't error")
}

func (suite *NumberServiceTestSuite) TestStoreFail() {
	number := &model.Number{
		Value: "123456789",
	}

	fail := errors.New("something wrong happens")
	suite.numberRepository.EXPECT().Exists(gomock.AssignableToTypeOf(number)).Return(false, nil).Times(1)

	suite.numberRepository.EXPECT().Save(gomock.AssignableToTypeOf(number)).Return(fail).Times(1)

	err := suite.underTest.Store(number)

	suite.Error(err, "Should return error")
}

func (suite *NumberServiceTestSuite) TestIsValidNumberSuccess() {
	n := &model.Number{
		Value: "123456789",
	}

	result := suite.underTest.IsValidNumber(n)

	suite.Equal(true, result)
}

func (suite *NumberServiceTestSuite) TestIsValidNumberIsNotNineDigits() {
	n := &model.Number{
		Value: "1234",
	}

	result := suite.underTest.IsValidNumber(n)

	suite.Equal(false, result)
}

func (suite *NumberServiceTestSuite) TestIsValidNumberIsNotNumeric() {
	n := &model.Number{
		Value: "1234abcdf",
	}

	result := suite.underTest.IsValidNumber(n)

	suite.Equal(false, result)
}
