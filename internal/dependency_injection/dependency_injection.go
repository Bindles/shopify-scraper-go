package dependency_injection

import (
	"github.com/bindles/shopify-scraper-go/internal/resources"
	"github.com/bindles/shopify-scraper-go/internal/services"
	"net/url"
)

type RequiredObjects struct {
	ProductsCSVWriterService services.ProductsCSVWriterService
}

func ConstructRequiredObjects(shopifyStoreUrl url.URL) RequiredObjects {
	productCSVConversionService := services.NewProductCSVConversionService()
	shopifyResource := resources.NewShopifyResource(shopifyStoreUrl)
	productsRetrievalService := services.NewProductsRetrievalService(shopifyResource)
	productsCSVWriterService := services.NewProductsCSVWriterService(productCSVConversionService, productsRetrievalService)
	return RequiredObjects{
		ProductsCSVWriterService: productsCSVWriterService,
	}
}
