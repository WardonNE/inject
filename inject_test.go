package inject

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	assert.Equal(t, &container{constructor: "Init"}, NewContainer())
}

func TestSetConstruct(t *testing.T) {
	container := NewContainer()
	container.SetConstructor("Construct")
	assert.Equal(t, "Construct", container.constructor)
}

type TestBase struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type iTestLogger interface {
	Output() string
}

type testLogger struct {
}

func (logger *testLogger) Output() string {
	return "logger"
}

type iTestLogger2 interface {
	Output2() string
}

type testLogger2 struct {
}

func (logger *testLogger2) Output2() string {
	return "logger2"
}

type testUser struct {
	Attributes map[string]any `inject:""`
	ID         int
	Name       string
	Tag        []string     `inject:""`
	Company    *testCompany `inject:""`
	Group      *testGroup   `inject:"group"`
	Logger     iTestLogger  `inject:"logger"`
	Logger2    iTestLogger2 `inject:""`
	TestBase   `inject:""`
}

func (t *testUser) Init() {
	t.ID = 1
	t.Name = "testUser"
	t.Tag = append(t.Tag, "Tag1")
	t.Attributes["gender"] = "female"
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}

type testGroup struct {
	ID int
}

type testCompany struct {
	ID       int
	Name     string
	Country  *testCountry `inject:"country"`
	TestBase `inject:""`
}

func (t *testCompany) Init() {
	t.ID = 2
	t.Name = "testCompany"
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}

type testCountry struct {
	ID       int
	Name     string
	TestBase `inject:""`
}

func (t *testCountry) Init() {
	t.ID = 3
	t.Name = "testCountry"
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}

func TestStoreWithDefaultInit(t *testing.T) {
	container := NewContainer()
	testUser := new(testUser)
	err := container.Provide("logger", new(testLogger))
	err = container.Provide("logger2", new(testLogger2))
	err = container.Provide("country", new(testCountry))
	err = container.Provide("user", testUser)
	testUser.Attributes["gender"] = "female"
	assert.Equal(t, nil, err)
	assert.Equal(t, "female", testUser.Attributes["gender"])
	assert.Equal(t, "testUser", testUser.Name)
	assert.Equal(t, 1, testUser.ID)
	assert.Equal(t, 2, testUser.Company.ID)
	assert.Equal(t, "testCompany", testUser.Company.Name)
	assert.Equal(t, 3, testUser.Company.Country.ID)
	assert.Equal(t, "testCountry", testUser.Company.Country.Name)
}

func TestStoreWithCustomInit(t *testing.T) {
	container := NewContainer()
	testUser := new(testUser)
	err := container.Provide("logger", new(testLogger))
	err = container.Provide("logger2", new(testLogger2))
	err = container.Provide("country", new(testCountry))
	err = container.Provide("user", testUser)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, testUser.ID)
	assert.Equal(t, 2, testUser.Company.ID)
	assert.Equal(t, 3, testUser.Company.Country.ID)
}
