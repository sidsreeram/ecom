package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductUsecase services.ProductUsecase
}

func NewProductHandler(Productusecase services.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: Productusecase,
	}

}
func (cr *ProductHandler) CreateCategory(c *gin.Context) {
	var category helperstruct.Category
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind ",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	Newcategory, err := cr.ProductUsecase.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Create a New Category",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category Created Successfully",
		Data:       Newcategory,
		Errors:     nil,
	})

}
func (cr *ProductHandler) UpdateCategory(c *gin.Context) {
	var category helperstruct.Category
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updatecategory, err := cr.ProductUsecase.UpdateCategory(category, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed to Update category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category Updated Successfully",
		Data:       updatecategory,
		Errors:     nil,
	})
}
func (cr *ProductHandler) DeleteCategory(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.ProductUsecase.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed To Delete",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category Deleted",
		Data:       nil,
		Errors:     nil,
	})
}
func (cr *ProductHandler) ListCategories(c *gin.Context) {
	category, err := cr.ProductUsecase.ListCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Load categories",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Categories are ",
		Data:       category,
		Errors:     nil,
	})
}
func (cr *ProductHandler) DisplayACategory(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	category, err := cr.ProductUsecase.DisplayACategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't fetch the data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Data fetched successfully",
		Data:       category,
		Errors:     nil,
	})
}

func (cr *ProductHandler) AddProduct(c *gin.Context) {
	var product helperstruct.Product
	err := c.Bind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind ",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	Newproduct, err := cr.ProductUsecase.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Add new Product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "New Product added Successfully",
		Data:       Newproduct,
		Errors:     nil,
	})

}
func (cr *ProductHandler) UpdateProduct(c *gin.Context) {
	var Product helperstruct.Product
	err := c.Bind(&Product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updateProduct, err := cr.ProductUsecase.UpdateProduct(id, Product)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed to Update Porduct",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product Updated Successfully",
		Data:       updateProduct,
		Errors:     nil,
	})
}
func (cr *ProductHandler) DeleteProduct(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.ProductUsecase.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed To Delete",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product Deleted",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *ProductHandler) ListAllProduct(c *gin.Context) {

	var viewProduct helperstruct.QueryParams

	viewProduct.Page, _ = strconv.Atoi(c.Query("page"))
	viewProduct.Limit, _ = strconv.Atoi(c.Query("limit"))
	viewProduct.Query = c.Query("query")
	viewProduct.Filter = c.Query("filter")
	viewProduct.SortBy = c.Query("sort_by")
	viewProduct.SortDesc, _ = strconv.ParseBool(c.Query("sort_desc"))
	fmt.Println(viewProduct)

	products, err := cr.ProductUsecase.ListAllProduct(viewProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product",
		Data:       products,
		Errors:     nil,
	})
}
func (cr *ProductHandler) DisplayAProduct(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	product, err := cr.ProductUsecase.DisplayAProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product",
		Data:       product,
		Errors:     nil,
	})
}

// -----------------------------------------------------ProductItem------------------------------------------------------------
func (cr *ProductHandler) AddProductitem(c *gin.Context) {
	var productItem helperstruct.ProductItem
	err := c.Bind(&productItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind ",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	newProductItem, err := cr.ProductUsecase.AddProductitem(productItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Add new Product-item",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "New Product-item added Successfully",
		Data:       newProductItem,
		Errors:     nil,
	})

}

//----------------------------------------------------update ProductItem-------------------------------------------------------------

func (cr *ProductHandler) UpdateProductitem(c *gin.Context) {
	var Productitem helperstruct.ProductItem
	err := c.Bind(&Productitem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updateProduct, err := cr.ProductUsecase.UpdateProductItem(id, Productitem)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed to Update Porduct",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product Updated Successfully",
		Data:       updateProduct,
		Errors:     nil,
	})
}

// --------------------------------------Delete ProductItem---------------------------------------------------------------
func (cr *ProductHandler) DeleteProductItem(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.ProductUsecase.DeleteProductItem(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Failed To Delete",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Product Deleted",
		Data:       nil,
		Errors:     nil,
	})
}
func (cr *ProductHandler) DisaplyaAllProductItems(c *gin.Context) {
	var viewProductaItem helperstruct.QueryParams

	viewProductaItem.Page, _ = strconv.Atoi(c.Query("page"))
	viewProductaItem.Limit, _ = strconv.Atoi(c.Query("limit"))
	viewProductaItem.Query = c.Query("query")
	viewProductaItem.Filter = c.Query("filter")
	viewProductaItem.SortBy = c.Query("sort_by")
	viewProductaItem.SortDesc, _ = strconv.ParseBool(c.Query("sort_desc"))

	productItems, err := cr.ProductUsecase.DisaplyaAllProductItems(viewProductaItem)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't disaply items",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product items are",
		Data:       productItems,
		Errors:     nil,
	})
}
func (cr *ProductHandler) DisplayAproductitem(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	product, err := cr.ProductUsecase.DisplayAproductitem(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product",
		Data:       product,
		Errors:     nil,
	})
}
func (cr *ProductHandler) UploadImage(c *gin.Context) {

	id := c.Param("id")
	productId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find product id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["images"]

	for _, file := range files {
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "asset/uploads"+file.Filename)
		fmt.Println(file.Filename)
		err := cr.ProductUsecase.UploadImage(file.Filename, productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "cant upload images",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, response.Response{
			StatusCode: 200,
			Message:    " images uploaded succesfully",
			Data:       nil,
			Errors:     nil,
		})
	}
}

func (cr *ProductHandler) UploadImageBinary(c *gin.Context) {
    id := c.Param("id")
    productId, err := strconv.Atoi(id)

    if err != nil {
        c.JSON(http.StatusBadRequest, response.Response{
            StatusCode: 400,
            Message:    "can't find product id",
            Data:       nil,
            Errors:     err.Error(),
        })
        return
    }

    // Multipart form
    form, _ := c.MultipartForm()
    files := form.File["images"]

    for _, file := range files {
        // Open the file
        uploadedFile, err := file.Open()
        if err != nil {
            c.JSON(http.StatusBadRequest, response.Response{
                StatusCode: 400,
                Message:    "can't open uploaded file",
                Data:       nil,
                Errors:     err.Error(),
            })
            return
        }
        defer uploadedFile.Close()

        // Read the file content into a byte slice
        fileBytes, err := io.ReadAll(uploadedFile)
        if err != nil {
            c.JSON(http.StatusBadRequest, response.Response{
                StatusCode: 400,
                Message:    "can't read file content",
                Data:       nil,
                Errors:     err.Error(),
            })
            return
        }

        // Upload the image as bytes
        err = cr.ProductUsecase.UploadImageBinary(fileBytes, file.Filename, productId)
        if err != nil {
            c.JSON(http.StatusBadRequest, response.Response{
                StatusCode: 400,
                Message:    "can't upload image",
                Data:       nil,
                Errors:     err.Error(),
            })
            return
        }

        c.JSON(http.StatusOK, response.Response{
            StatusCode: 200,
            Message:    "image uploaded successfully",
            Data:       nil,
            Errors:     nil,
        })
    }
}
func (cr *ProductHandler) GetProductImages(c *gin.Context){
	id := c.Param("id")
    productId, err := strconv.Atoi(id)

    if err != nil {
        c.JSON(http.StatusBadRequest, response.Response{
            StatusCode: 400,
            Message:    "can't find product id",
            Data:       nil,
            Errors:     err.Error(),
        })
        return
    }
	images, err := cr.ProductUsecase.GetProductImages(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find product images",
			Data: nil,
			Errors: err.Error(),
		
		})
		return
	}
	c.JSON(http.StatusOK,response.Response{
		StatusCode: 200,
		Message: "image for the product item fetched successfully",
		Data: images,
		Errors: nil,
	})
}