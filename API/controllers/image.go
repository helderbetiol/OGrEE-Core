package controllers

import (
	"fmt"
	"net/http"
	"p3/models"
	u "p3/utils"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const GetImagePath = "/api/images/"

// Transforms image inside data from an ID to a URL
func imageIDToUrl(entity int, object map[string]any) map[string]any {
	if entity == u.TAG {
		if imageID, hasImage := object["image"].(primitive.ObjectID); hasImage && imageID != primitive.NilObjectID {
			object["image"] = GetImagePath + imageID.Hex()
		} else if imageID, hasImage := object["image"].(string); hasImage && imageID != "" {
			object["image"] = GetImagePath + imageID
		} else {
			object["image"] = ""
		}
	}

	return object
}

// swagger:operation GET /api/images/{id} Objects GetImage
// Gets an image by its id.
// Returns a HTTP response with the binary of the image.
// The format of the image will correspond to the content type of the response.
// ---
// security:
// - bearer: []
// produces:
// - image/*
// - application/json
// parameters:
// - name: id
//   in: path
//   description: ID of image
//   required: true
//   type: string
// responses:
//     '200':
//         description: Found. A response body will be returned with the image binary.
//     '404':
//         description: Not Found. An error message will be returned.

func GetImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******************************************************")
	fmt.Println("FUNCTION CALL: 	 GetImage ")
	fmt.Println("******************************************************")
	DispRequestMetaData(r)

	image, err := models.GetImage(mux.Vars(r)["id"])
	if err != nil {
		u.ErrLog("Error while getting image",
			"GET GetImage", err.Message, r)
		u.RespondWithError(w, err)
	} else if r.Method == "OPTIONS" {
		w.Header().Add("Content-Type", image.MIMEType)
		w.Header().Add("Allow", "GET, OPTIONS")
	} else {
		w.Header().Add("Content-Type", image.MIMEType)
		w.Write(image.Data)
	}
}
