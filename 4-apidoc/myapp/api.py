"""
@apiDefine ProductNotFoundError

@apiError ProductNotFound The id of the Product was not found.

@apiErrorExample Error-Response:
    HTTP/1.1 404 Not Found
    {
      "error": "ProductNotFound"
    }
"""

"""
@apiDefine admin Admin access only
    Admins permissions only
"""

"""
@api {get} /products_list Request list of all products
@apiName ProductsList
@apiGroup Products

@apiExample {curl} Example usage:
    curl -i http://localhost/products_list

@apiSuccess {String[]} productslist List of all products.

@apiSuccessExample Success-Response:
    HTTP/1.1 200 OK
    {
      [
        "id": "2121",
        "name": "Sony Playstation 4",
        "description": "Home video game console",
        "price": "500$"
      ],
      [
        "id": "2122",
        "name": "Xbox One S",
        "description": "Home video game console",
        "price": "470$"
      ]
    }
"""

"""
@api {get} /products/:id Request information about product
@apiName GetProduct
@apiGroup Products

@apiParam {Number} id Product unique ID.

@apiExample {curl} Example usage:
    curl -i http://localhost/products/2121

@apiSuccess {Number} id          Product id.
@apiSuccess {String} name        Name of the product.
@apiSuccess {String} description Description of the product.
@apiSuccess {String} price       Price of the product.

@apiSuccessExample Success-Response:
    HTTP/1.1 200 OK
    {
      "id": "2121",
      "name": "Sony Playstation 4",
      "description": "Home video game console",
      "price": "500$"
    }

@apiUse ProductNotFoundError
"""

"""
@api {put} /products/:id Modify Product information
@apiName PutProduct
@apiGroup Products

@apiHeader {String} access-key Users unique access-key.
@apiPermission admin

@apiParam {Number} id            Product id.
@apiParam {String} name          Name of the product.
@apiParam {String} description   Description of the product.
@apiParam {String} price         Price of the product.

@apiSuccessExample Success-Response:
    HTTP/1.1 200 OK

@apiUse ProductNotFoundError
 """
