package graphql

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPartService mocks the PartService
type MockPartService struct {
	mock.Mock
}

func (m *MockPartService) GetParts(ctx context.Context) ([]model.Part, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Part), args.Error(1)
}

func (m *MockPartService) GetPartByID(ctx context.Context, id uuid.UUID) (*model.Part, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Part), args.Error(1)
}

func (m *MockPartService) SearchParts(ctx context.Context, name string) ([]model.Part, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Part), args.Error(1)
}

func (m *MockPartService) CreatePart(ctx context.Context, input model.CreatePartInput) (*model.Part, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Part), args.Error(1)
}

func (m *MockPartService) UpdatePart(ctx context.Context, id uuid.UUID, input model.UpdatePartInput) (*model.Part, error) {
	args := m.Called(ctx, id, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Part), args.Error(1)
}

func (m *MockPartService) DeletePart(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockManufacturerService mocks the ManufacturerService
type MockManufacturerService struct {
	mock.Mock
}

func (m *MockManufacturerService) GetManufacturers(ctx context.Context) ([]model.Manufacturer, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Manufacturer), args.Error(1)
}

func (m *MockManufacturerService) GetManufacturerByID(ctx context.Context, id uuid.UUID) (*model.Manufacturer, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Manufacturer), args.Error(1)
}

func (m *MockManufacturerService) CreateManufacturer(ctx context.Context, input model.CreateManufacturerInput) (*model.Manufacturer, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Manufacturer), args.Error(1)
}

func (m *MockManufacturerService) UpdateManufacturer(ctx context.Context, id uuid.UUID, input model.UpdateManufacturerInput) (*model.Manufacturer, error) {
	args := m.Called(ctx, id, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Manufacturer), args.Error(1)
}

func (m *MockManufacturerService) DeleteManufacturer(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockCategoryService mocks the CategoryService
type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) GetCategories(ctx context.Context) ([]model.Category, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Category), args.Error(1)
}

func (m *MockCategoryService) GetCategoryByID(ctx context.Context, id uuid.UUID) (*model.Category, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *MockCategoryService) CreateCategory(ctx context.Context, input model.CreateCategoryInput) (*model.Category, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *MockCategoryService) UpdateCategory(ctx context.Context, id uuid.UUID, input model.UpdateCategoryInput) (*model.Category, error) {
	args := m.Called(ctx, id, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *MockCategoryService) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockPartsInventoryService mocks the PartsInventoryService
type MockPartsInventoryService struct {
	mock.Mock
}

func (m *MockPartsInventoryService) GetPartsInventory(ctx context.Context) ([]model.PartsInventory, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.PartsInventory), args.Error(1)
}

func (m *MockPartsInventoryService) GetPartsInventoryByID(ctx context.Context, id uuid.UUID) (*model.PartsInventory, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.PartsInventory), args.Error(1)
}

func (m *MockPartsInventoryService) GetPartsInventoryByPartID(ctx context.Context, partID uuid.UUID) ([]model.PartsInventory, error) {
	args := m.Called(ctx, partID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.PartsInventory), args.Error(1)
}

func (m *MockPartsInventoryService) CreatePartsInventory(ctx context.Context, input model.CreatePartsInventoryInput) (*model.PartsInventory, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.PartsInventory), args.Error(1)
}

func (m *MockPartsInventoryService) UpdatePartsInventory(ctx context.Context, id uuid.UUID, input model.UpdatePartsInventoryInput) (*model.PartsInventory, error) {
	args := m.Called(ctx, id, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.PartsInventory), args.Error(1)
}

func (m *MockPartsInventoryService) DeletePartsInventory(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockProductService mocks the ProductService
type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) GetProducts(ctx context.Context) ([]model.Product, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockProductService) GetProductByID(ctx context.Context, id uuid.UUID) (*model.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Product), args.Error(1)
}

func (m *MockProductService) SearchProducts(ctx context.Context, name string) ([]model.Product, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockProductService) CreateProduct(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Product), args.Error(1)
}

func (m *MockProductService) UpdateProduct(ctx context.Context, id uuid.UUID, input model.UpdateProductInput) (*model.Product, error) {
	args := m.Called(ctx, id, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Product), args.Error(1)
}

func (m *MockProductService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockProductService) AddProductPart(ctx context.Context, input model.AddProductPartInput) (*model.ProductPart, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ProductPart), args.Error(1)
}

func (m *MockProductService) UpdateProductPart(ctx context.Context, id uuid.UUID, input model.UpdateProductPartInput) (*model.ProductPart, error) {
	args := m.Called(ctx, id, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ProductPart), args.Error(1)
}

func (m *MockProductService) RemoveProductPart(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockProductInventoryService mocks the ProductInventoryService
type MockProductInventoryService struct {
	mock.Mock
}

func (m *MockProductInventoryService) GetProductInventory(ctx context.Context) ([]model.ProductInventory, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.ProductInventory), args.Error(1)
}

func (m *MockProductInventoryService) GetProductInventoryByID(ctx context.Context, id uuid.UUID) (*model.ProductInventory, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ProductInventory), args.Error(1)
}

func (m *MockProductInventoryService) GetProductInventoryByProductID(ctx context.Context, productID uuid.UUID) ([]model.ProductInventory, error) {
	args := m.Called(ctx, productID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.ProductInventory), args.Error(1)
}

func (m *MockProductInventoryService) CreateProductInventory(ctx context.Context, input model.CreateProductInventoryInput) (*model.ProductInventory, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ProductInventory), args.Error(1)
}

func (m *MockProductInventoryService) UpdateProductInventory(ctx context.Context, id uuid.UUID, input model.UpdateProductInventoryInput) (*model.ProductInventory, error) {
	args := m.Called(ctx, id, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.ProductInventory), args.Error(1)
}

func (m *MockProductInventoryService) DeleteProductInventory(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockProductInventoryService) AddPartToProductInventory(ctx context.Context, productInventoryID uuid.UUID, partsInventoryID uuid.UUID) error {
	args := m.Called(ctx, productInventoryID, partsInventoryID)
	return args.Error(0)
}

func (m *MockProductInventoryService) RemovePartFromProductInventory(ctx context.Context, productInventoryID uuid.UUID, partsInventoryID uuid.UUID) error {
	args := m.Called(ctx, productInventoryID, partsInventoryID)
	return args.Error(0)
}

// ============================================================================
// QUERY TESTS
// ============================================================================

func TestQueryParts(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	testPart := model.Part{
		ID:   testID,
		Name: "Test Part",
	}

	mockPartSvc.On("GetParts", mock.Anything).Return([]model.Part{testPart}, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	parts, err := r.Parts(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, parts)
	assert.Equal(t, 1, len(parts))
	assert.Equal(t, testPart.Name, parts[0].Name)

	mockPartSvc.AssertExpectations(t)
}

func TestQueryPartByID(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	testPart := model.Part{
		ID:   testID,
		Name: "Test Part",
	}

	mockPartSvc.On("GetPartByID", mock.Anything, testID).Return(&testPart, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	part, err := r.Part(context.Background(), testID)
	assert.NoError(t, err)
	assert.NotNil(t, part)
	assert.Equal(t, testPart.Name, part.Name)

	mockPartSvc.AssertExpectations(t)
}

func TestQuerySearchParts(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	testPart := model.Part{
		ID:   testID,
		Name: "Resistor",
	}

	mockPartSvc.On("SearchParts", mock.Anything, "Resistor").Return([]model.Part{testPart}, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	parts, err := r.SearchParts(context.Background(), "Resistor")
	assert.NoError(t, err)
	assert.NotNil(t, parts)
	assert.Equal(t, 1, len(parts))
	assert.Equal(t, testPart.Name, parts[0].Name)

	mockPartSvc.AssertExpectations(t)
}

func TestQueryCategories(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	testCategory := model.Category{
		ID:   testID,
		Name: "Electronics",
	}

	mockCatSvc.On("GetCategories", mock.Anything).Return([]model.Category{testCategory}, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	categories, err := r.Categories(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, categories)
	assert.Equal(t, 1, len(categories))
	assert.Equal(t, testCategory.Name, categories[0].Name)

	mockCatSvc.AssertExpectations(t)
}

func TestQueryCategoryByID(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	testCategory := model.Category{
		ID:   testID,
		Name: "Electronics",
	}

	mockCatSvc.On("GetCategoryByID", mock.Anything, testID).Return(&testCategory, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	category, err := r.Category(context.Background(), testID)
	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, testCategory.Name, category.Name)

	mockCatSvc.AssertExpectations(t)
}

func TestQueryManufacturers(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	testMfg := model.Manufacturer{
		ID:   testID,
		Name: "Samsung",
	}

	mockMfgSvc.On("GetManufacturers", mock.Anything).Return([]model.Manufacturer{testMfg}, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	manufacturers, err := r.Manufacturers(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, manufacturers)
	assert.Equal(t, 1, len(manufacturers))
	assert.Equal(t, testMfg.Name, manufacturers[0].Name)

	mockMfgSvc.AssertExpectations(t)
}

func TestQueryManufacturerByID(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	testMfg := model.Manufacturer{
		ID:   testID,
		Name: "Samsung",
	}

	mockMfgSvc.On("GetManufacturerByID", mock.Anything, testID).Return(&testMfg, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	manufacturer, err := r.Manufacturer(context.Background(), testID)
	assert.NoError(t, err)
	assert.NotNil(t, manufacturer)
	assert.Equal(t, testMfg.Name, manufacturer.Name)

	mockMfgSvc.AssertExpectations(t)
}

func TestQueryProducts(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	testProduct := model.Product{
		ID:   testID,
		Name: "Widget 3000",
	}

	mockProdSvc.On("GetProducts", mock.Anything).Return([]model.Product{testProduct}, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	products, err := r.Products(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 1, len(products))
	assert.Equal(t, testProduct.Name, products[0].Name)

	mockProdSvc.AssertExpectations(t)
}

func TestQueryProductByID(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	testProduct := model.Product{
		ID:   testID,
		Name: "Widget 3000",
	}

	mockProdSvc.On("GetProductByID", mock.Anything, testID).Return(&testProduct, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	product, err := r.Product(context.Background(), testID)
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, testProduct.Name, product.Name)

	mockProdSvc.AssertExpectations(t)
}

func TestQuerySearchProducts(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	testProduct := model.Product{
		ID:   testID,
		Name: "Widget",
	}

	mockProdSvc.On("SearchProducts", mock.Anything, "Widget").Return([]model.Product{testProduct}, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	products, err := r.SearchProducts(context.Background(), "Widget")
	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, 1, len(products))
	assert.Equal(t, testProduct.Name, products[0].Name)

	mockProdSvc.AssertExpectations(t)
}

func TestQueryPartsInventory(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	partID := uuid.New()
	testPartsInv := model.PartsInventory{
		ID:           testID,
		PartID:       partID,
		SerialNumber: "SN123",
		IsAvailable:  true,
	}

	mockPartsInvSvc.On("GetPartsInventory", mock.Anything).Return([]model.PartsInventory{testPartsInv}, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	inventory, err := r.PartsInventory(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, inventory)
	assert.Equal(t, 1, len(inventory))
	assert.Equal(t, testPartsInv.SerialNumber, inventory[0].SerialNumber)

	mockPartsInvSvc.AssertExpectations(t)
}

func TestQueryPartsInventoryItem(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	partID := uuid.New()
	testPartsInv := model.PartsInventory{
		ID:           testID,
		PartID:       partID,
		SerialNumber: "SN123",
		IsAvailable:  true,
	}

	mockPartsInvSvc.On("GetPartsInventoryByID", mock.Anything, testID).Return(&testPartsInv, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	item, err := r.PartsInventoryItem(context.Background(), testID)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, testPartsInv.SerialNumber, item.SerialNumber)

	mockPartsInvSvc.AssertExpectations(t)
}

func TestQueryPartsInventoryByPart(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	partID := uuid.New()
	testPartsInv := model.PartsInventory{
		ID:           testID,
		PartID:       partID,
		SerialNumber: "SN123",
		IsAvailable:  true,
	}

	mockPartsInvSvc.On("GetPartsInventoryByPartID", mock.Anything, partID).Return([]model.PartsInventory{testPartsInv}, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	items, err := r.PartsInventoryByPart(context.Background(), partID)
	assert.NoError(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 1, len(items))
	assert.Equal(t, testPartsInv.SerialNumber, items[0].SerialNumber)

	mockPartsInvSvc.AssertExpectations(t)
}

func TestQueryProductInventory(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	productID := uuid.New()
	testProdInv := model.ProductInventory{
		ID:           testID,
		ProductID:    productID,
		SerialNumber: "PROD123",
		IsAvailable:  true,
	}

	mockProdInvSvc.On("GetProductInventory", mock.Anything).Return([]model.ProductInventory{testProdInv}, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	inventory, err := r.ProductInventory(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, inventory)
	assert.Equal(t, 1, len(inventory))
	assert.Equal(t, testProdInv.SerialNumber, inventory[0].SerialNumber)

	mockProdInvSvc.AssertExpectations(t)
}

func TestQueryProductInventoryItem(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	productID := uuid.New()
	testProdInv := model.ProductInventory{
		ID:           testID,
		ProductID:    productID,
		SerialNumber: "PROD123",
		IsAvailable:  true,
	}

	mockProdInvSvc.On("GetProductInventoryByID", mock.Anything, testID).Return(&testProdInv, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	item, err := r.ProductInventoryItem(context.Background(), testID)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, testProdInv.SerialNumber, item.SerialNumber)

	mockProdInvSvc.AssertExpectations(t)
}

func TestQueryProductInventoryByProduct(t *testing.T) {
	mockPartSvc := new(MockPartService)
	mockMfgSvc := new(MockManufacturerService)
	mockCatSvc := new(MockCategoryService)
	mockPartsInvSvc := new(MockPartsInventoryService)
	mockProdSvc := new(MockProductService)
	mockProdInvSvc := new(MockProductInventoryService)

	testID := uuid.New()
	productID := uuid.New()
	testProdInv := model.ProductInventory{
		ID:           testID,
		ProductID:    productID,
		SerialNumber: "PROD123",
		IsAvailable:  true,
	}

	mockProdInvSvc.On("GetProductInventoryByProductID", mock.Anything, productID).Return([]model.ProductInventory{testProdInv}, nil)

	resolver := NewResolver(mockPartSvc, mockMfgSvc, mockCatSvc, mockPartsInvSvc, mockProdSvc, mockProdInvSvc)
	r := resolver.Query()

	items, err := r.ProductInventoryByProduct(context.Background(), productID)
	assert.NoError(t, err)
	assert.NotNil(t, items)
	assert.Equal(t, 1, len(items))
	assert.Equal(t, testProdInv.SerialNumber, items[0].SerialNumber)

	mockProdInvSvc.AssertExpectations(t)
}
