package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"p3/models"
	u "p3/utils"
	"strings"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	TENANT = iota
	SITE
	BLDG
	ROOM
	RACK
	DEVICE
	SUBDEV
	SUBDEV1
	AC
	PWRPNL
	WALL
	CABINET
	AISLE
	TILE
	GROUP
	CORIDOR
	ROOMTMPL
	OBJTMPL
)

func getObjID(x string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(x)
	if err != nil {
		return objID, err
	}
	return objID, nil
}

func parseDataForNonStdResult(ent string, eNum int, data map[string]interface{}) map[string][]map[string]interface{} {

	ans := map[string][]map[string]interface{}{}
	add := []map[string]interface{}{}

	firstIndex := u.EntityToString(eNum + 1)
	firstArr := data[firstIndex+"s"].([]map[string]interface{})

	ans[firstIndex+"s"] = firstArr

	for i := range firstArr {
		nxt := u.EntityToString(eNum + 2)
		add = append(add, firstArr[i][nxt+"s"].([]map[string]interface{})...)
	}

	ans[u.EntityToString(eNum+2)+"s"] = add
	newAdd := []map[string]interface{}{}
	for i := range add {
		nxt := u.EntityToString(eNum + 3)
		newAdd = append(newAdd, add[i][nxt+"s"].([]map[string]interface{})...)
	}

	ans[u.EntityToString(eNum+3)+"s"] = newAdd

	newAdd2 := []map[string]interface{}{}
	for i := range newAdd {
		nxt := u.EntityToString(eNum + 4)
		newAdd2 = append(newAdd2, newAdd[i][nxt+"s"].([]map[string]interface{})...)
	}

	ans[u.EntityToString(eNum+4)+"s"] = newAdd2
	newAdd3 := []map[string]interface{}{}

	for i := range newAdd2 {
		nxt := u.EntityToString(eNum + 5)
		newAdd3 = append(newAdd3, newAdd2[i][nxt+"s"].([]map[string]interface{})...)
	}
	ans[u.EntityToString(eNum+5)+"s"] = newAdd3

	newAdd4 := []map[string]interface{}{}

	for i := range newAdd3 {
		nxt := u.EntityToString(eNum + 6)
		newAdd4 = append(newAdd4, newAdd3[i][nxt+"s"].([]map[string]interface{})...)
	}

	ans[u.EntityToString(eNum+6)+"s"] = newAdd4

	newAdd5 := []map[string]interface{}{}

	for i := range newAdd4 {
		nxt := u.EntityToString(eNum + 7)
		newAdd5 = append(newAdd5, newAdd4[i][nxt+"s"].([]map[string]interface{})...)
	}

	ans[u.EntityToString(eNum+7)+"s"] = newAdd5

	//add := []map[string]interface{}{}

	//Get All first entities
	/*for i := eNum + 1; i < SUBDEV1; i++ {
		add = append(add, firstArr[i])
	}*/
	return ans
}

//This function is useful for debugging
//purposes. It displays any JSON
func viewJson(r *http.Request) {
	var updateData map[string]interface{}
	json.NewDecoder(r.Body).Decode(&updateData)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")

	if err := enc.Encode(updateData); err != nil {
		log.Fatal(err)
	}
}

// swagger:operation POST /api/{obj} objects CreateObject
// Creates an object in the system.
// ---
// produces:
// - application/json
// parameters:
// - name: objs
//   in: query
//   description: 'Indicates the Object. Only values of "tenants", "sites",
//   "buildings", "rooms", "racks", "devices", "acs", "panels",
//   "walls","aisles", "tiles", "cabinets", "groups", "corridors",
//   "room-templates", "obj-templates",
//   "racksensors","devicesensors" are acceptable'
//   required: true
//   type: string
//   default: "sites"
// - name: Name
//   in: query
//   description: Name of object
//   required: true
//   type: string
//   default: "Object A"
// - name: Category
//   in: query
//   description: Category of Object (ex. Consumer Electronics, Medical)
//   required: true
//   type: string
//   default: "Research"
// - name: Domain
//   description: 'Domain of Object'
//   required: true
//   type: string
//   default: 999
// - name: ParentID
//   description: 'All objects are linked to a
//   parent with the exception of Tenant since it has no parent'
//   required: true
//   type: int
//   default: 999
// - name: Description
//   in: query
//   description: Description of Object
//   required: false
//   type: string[]
//   default: ["Some abandoned object in Grenoble"]
// - name: Attributes
//   in: query
//   description: 'Any other object attributes can be added.
//   They are required depending on the obj type.'
//   required: true
//   type: json
// responses:
//     '201':
//         description: Created
//     '400':
//         description: Bad request

var CreateEntity = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******************************************************")
	fmt.Println("FUNCTION CALL: 	 CreateEntity ")
	fmt.Println("******************************************************")
	var e string
	var resp map[string]interface{}
	entity := map[string]interface{}{}
	err := json.NewDecoder(r.Body).Decode(&entity)

	//strip the '/api' in URL
	entStr, e1 := mux.Vars(r)["entity"]
	if e1 == false {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Error while parsing path params"))
		u.ErrLog("Error while parsing path params", "CREATE "+entStr, "", r)
		return
	}

	entStr = entStr[:len(entStr)-1] // and the trailing 's'
	entUpper := strings.ToUpper(entStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		u.ErrLog("Error while decoding request body", "CREATE "+entStr, "", r)
		return
	}

	switch entStr {
	case "ac", "panel", "wall", "cabinet", "tile", //NESTED Objects
		"aisle", "group", "corridor", "racksensor", "devicesensor":
		entity["id"] = primitive.NewObjectID().Hex()

		i := u.EntityStrToInt(entStr)
		println("ENT: ", entStr)
		println("ENUM VAL: ", i)
		resp, e = models.CreateNestedEntity(i, entStr, entity)
	default:
		//If creating templates, format them
		if idx := strings.Index(entStr, "-"); idx != -1 {
			//entStr[idx] = '_'
			entStr = entStr[:idx] + "_" + entStr[idx+1:]
		}
		i := u.EntityStrToInt(entStr)

		println("ENT: ", entStr)
		println("ENUM VAL: ", i)

		//Prevents Mongo from creating a new unidentified collection
		if i < 0 {
			w.WriteHeader(http.StatusNotFound)
			u.Respond(w, u.Message(false, "Invalid object: Please provide a valid object"))
			u.ErrLog("Cannot create invalid object", "CREATE "+entStr, "", r)
			return
		}

		resp, e = models.CreateEntity(i, entity)

	}

	switch e {
	case "validate", "duplicate":
		w.WriteHeader(http.StatusBadRequest)
		u.ErrLog("Error while creating "+entStr, "CREATE "+entUpper, e, r)
	case "":
		w.WriteHeader(http.StatusCreated)
	default:
		if strings.Split(e, " ")[1] == "duplicate" {
			w.WriteHeader(http.StatusBadRequest)
			u.ErrLog("Error: Duplicate "+entStr+" is forbidden",
				"CREATE "+entUpper, e, r)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			u.ErrLog("Error while creating "+entStr, "CREATE "+entUpper, e, r)
		}
	}

	u.Respond(w, resp)
}

// swagger:operation GET /api/{objs}/{id} objects GetObject
// Gets an Object from the system.
// The ID must be provided in the URL parameter
// The name can be used instead of ID if the obj is tenant
// ---
// produces:
// - application/json
// parameters:
// - name: objs
//   in: query
//   description: 'Indicates the location. Only values of "tenants", "sites",
//   "buildings", "rooms", "racks", "devices", "room-templates",
//   "obj-templates" are acceptable
//   For rooms, walls, acs, panels, aisles, tiles, cabinets, groups,
//   corridors, racksensors, and devicesensors
//   can be obtained by appending /subobj type and subobj id'
//   required: true
//   type: string
//   default: "sites"
// - name: ID
//   in: path
//   description: ID of desired object or Name of Tenant
//   required: true
//   type: int
//   default: 999
// responses:
//     '200':
//         description: Found
//     '400':
//         description: Bad request
//     '404':
//         description: Not Found
var GetEntity = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******************************************************")
	fmt.Println("FUNCTION CALL: 	 GetEntity ")
	fmt.Println("******************************************************")
	var data map[string]interface{}
	var id, e1 string
	var x primitive.ObjectID
	var e bool
	var e2 error

	resp := u.Message(true, "success")

	//Get entity type and strip trailing 's'
	s, _ := mux.Vars(r)["entity"]
	s = s[:len(s)-1]

	//GET By ID
	if id, e = mux.Vars(r)["id"]; e == true {

		if nestID, err := mux.Vars(r)["nest"]; err == true { //If nested
			//Get subentity type and strip trailing 's'
			s, _ = mux.Vars(r)["subent"]
			s = s[:len(s)-1]
			ID, _ := getObjID(id)

			data, e1 = models.GetNestedEntity(ID, s, nestID)
		} else { // Not Nested

			x, e2 = getObjID(id)
			if e2 != nil {
				u.Respond(w, u.Message(false, "Error while converting ID to ObjectID"))
				u.ErrLog("Error while converting ID to ObjectID", "GET ENTITY", "", r)
				return
			}

			data, e1 = models.GetEntity(x, s)

		}

	} else if id, e = mux.Vars(r)["name"]; e == true { //GET By String
		//If templates, format them
		if idx := strings.Index(s, "-"); idx != -1 { //GET By Slug
			s = s[:idx] + "_" + s[idx+1:]
			data, e1 = models.GetEntityBySlug(id, s)
		} else {
			data, e1 = models.GetEntityByName(id, s) //GET By Name
		}
	}

	if e == false {
		u.Respond(w, u.Message(false, "Error while parsing path parameters"))
		u.ErrLog("Error while parsing path parameters", "GET ENTITY", "", r)
		return
	}

	if data == nil {
		resp = u.Message(false, "Error while getting "+s+": "+e1)
		u.ErrLog("Error while getting "+s, "GET "+strings.ToUpper(s), "", r)

		switch e1 {
		case "record not found":
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusNotFound) //For now
		}

	} else {
		resp = u.Message(true, "success")
	}

	resp["data"] = data
	u.Respond(w, resp)
}

// swagger:operation GET /api/{objs} objects GetAllObjects
// Gets all present objects for specified category in the system.
// Returns JSON body with all specified objects of type and their IDs
// ---
// produces:
// - application/json
// parameters:
// - name: objs
//   in: query
//   description: 'Indicates the object. Only values of "tenants", "sites",
//   "buildings", "rooms", "racks", "devices", "room-templates",
//   "obj-templates" are acceptable
//   For rooms, walls, acs, panels, aisles, tiles, cabinets, groups,
//   corridors, racksensors, and devicesensors
//   can be obtained by appending /subobj type and subobj id'
//   required: true
//   type: string
//   default: "sites"
// responses:
//     '200':
//         description: Found
//     '404':
//         description: Nothing Found
var GetAllEntities = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******************************************************")
	fmt.Println("FUNCTION CALL: 	 GetAllEntities ")
	fmt.Println("******************************************************")
	var data []map[string]interface{}
	var e, entStr string
	arr := strings.Split(r.URL.Path, "/")
	if len(arr) < 4 { //Main hierarchy objects

		//entStr = arr[2][:len(arr[2])-1]
		entStr, _ = mux.Vars(r)["entity"]
		entStr = entStr[:len(entStr)-1]
		println("ENTSTR: ", entStr)

		//If templates, format them
		if idx := strings.Index(entStr, "-"); idx != -1 {
			entStr = entStr[:idx] + "_" + entStr[idx+1:]
		}

		data, e = models.GetAllEntities(entStr)

	} else { //Nested objects

		entStr, e2 := mux.Vars(r)["subent"]
		id, e1 := mux.Vars(r)["id"]
		if e1 == false || e2 == false {
			u.Respond(w, u.Message(false, "Error while parsing path parameters"))
			u.ErrLog("Error while parsing path parameters", "GET ENTITY", "", r)
			return
		}

		//strip trailing 's'
		entStr = entStr[:len(entStr)-1]

		ID, _ := getObjID(id)
		data, e = models.GetAllNestedEntities(ID, entStr)
	}

	entUpper := strings.ToUpper(entStr) // and the trailing 's'
	resp := u.Message(true, "success")

	if len(data) == 0 {
		resp = u.Message(false, "Error while getting "+entStr+": "+e)
		u.ErrLog("Error while getting "+entStr+"s", "GET ALL "+entUpper, e, r)

		switch e {
		case "":
			resp = u.Message(false,
				"Error while getting "+entStr+"s: No Records Found")
			w.WriteHeader(http.StatusNotFound)
		default:
		}

	} else {
		resp = u.Message(true, "success")
	}

	resp["data"] = map[string]interface{}{"objects": data}

	u.Respond(w, resp)
}

// swagger:operation DELETE /api/{objs}/{id} objects DeleteObject
// Deletes an Object in the system.
// ---
// produces:
// - application/json
// parameters:
// - name: objs
//   in: query
//   description: 'Indicates the object. Only values of "tenants", "sites",
//   "buildings", "rooms", "racks", "devices", "room-templates",
//   "obj-templates" are acceptable
//   For rooms, walls, acs, panels, aisles, tiles, cabinets, groups,
//   corridors, racksensors, and devicesensors
//   can be obtained by appending /subobj type and subobj id'
//   required: true
//   type: string
//   default: "sites"
// - name: ID
//   in: path
//   description: ID of desired object
//   required: true
//   type: int
//   default: 999
// responses:
//     '204':
//        description: Successful
//     '404':
//        description: Not found
var DeleteEntity = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******************************************************")
	fmt.Println("FUNCTION CALL: 	 DeleteEntity ")
	fmt.Println("******************************************************")
	var v map[string]interface{}
	id, e := mux.Vars(r)["id"]
	nest, e1 := mux.Vars(r)["nest"]
	name, e2 := mux.Vars(r)["name"]

	switch {
	case e2 == true: // DELETE SLUG
		//Get entity from URL and strip trailing 's'
		entity, _ := mux.Vars(r)["entity"]
		entity = entity[:len(entity)-1]

		//If templates, format them
		if idx := strings.Index(entity, "-"); idx != -1 {
			entity = entity[:idx] + "_" + entity[idx+1:]
		}
		v, _ = models.DeleteEntityBySlug(entity, name)

	case e1 == true && e == true: // DELETE NESTED

		//Get entity from URL and strip trailing 's'
		entity, _ := mux.Vars(r)["subent"]
		entity = entity[:len(entity)-1]

		objID, _ := getObjID(id)
		v, _ = models.DeleteNestedEntity(entity, objID, nest)

	case e == true: // DELETE NORMAL
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			u.Respond(w, u.Message(false, "Error while converting ID to ObjectID"))
			u.ErrLog("Error while converting ID to ObjectID", "DELETE ENTITY", "", r)
			return
		}

		//Get entity from URL and strip trailing 's'
		entity, _ := mux.Vars(r)["entity"]
		entity = entity[:len(entity)-1]

		if entity == "device" {
			v, _ = models.DeleteDeviceF(objID)
		} else {
			v, _ = models.DeleteEntity(entity, objID)
		}

	default:
		u.Respond(w, u.Message(false, "Error while parsing path parameters"))
		u.ErrLog("Error while parsing path parameters", "DELETE ENTITY", "", r)
		return
	}

	if v["status"] == false {
		w.WriteHeader(http.StatusNotFound)
		u.ErrLog("Error while deleting entity", "DELETE ENTITY", "Not Found", r)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}

	u.Respond(w, v)
}

// swagger:operation PUT /api/{objs}/{id} objects UpdateObject
// Changes Object data in the system.
// If no new or any information is provided
// an OK will still be returned
// ---
// produces:
// - application/json
// parameters:
// - name: objs
//   in: query
//   description: 'Indicates the object. Only values of "tenants", "sites",
//   "buildings", "rooms", "racks", "devices", "room-templates",
//   "obj-templates" are acceptable.
//   For rooms, walls, acs, panels, aisles, tiles, cabinets, groups,
//   corridors, racksensors, and devicesensors
//   can be obtained by appending /subobj type and subobj id'
//   required: true
//   type: string
//   default: "sites"
// - name: ID
//   in: path
//   description: ID of desired Object
//   required: true
//   type: int
//   default: 999
// - name: Name
//   in: query
//   description: Name of Object
//   required: false
//   type: string
//   default: "INFINITI"
// - name: Category
//   in: query
//   description: Category of Object (ex. Consumer Electronics, Medical)
//   required: false
//   type: string
//   default: "Auto"
// - name: Description
//   in: query
//   description: Description of Object
//   required: false
//   type: string[]
//   default: "High End Worldwide automotive company"
// - name: Domain
//   description: 'Domain of the Object'
//   required: false
//   type: string
//   default: "High End Auto"
// - name: Attributes
//   in: query
//   description: Any other object attributes can be updated
//   required: false
//   type: json
// responses:
//     '200':
//         description: Updated
//     '400':
//         description: Bad request
//     '404':
//         description: Not Found

var UpdateEntity = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******************************************************")
	fmt.Println("FUNCTION CALL: 	 UpdateEntity ")
	fmt.Println("******************************************************")
	var v map[string]interface{}
	var e3 string
	var entity string

	updateData := map[string]interface{}{}
	id, e := mux.Vars(r)["id"]
	nest, e1 := mux.Vars(r)["nest"]
	name, e2 := mux.Vars(r)["name"]

	//viewJson(r)

	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		u.ErrLog("Error while decoding request body", "UPDATE ENTITY", "", r)
	}

	switch {
	case e2 == true: // UPDATE SLUG
		println("updating slug")
		//Get entity from URL and strip trailing 's'
		entity, _ = mux.Vars(r)["entity"]
		entity = entity[:len(entity)-1]

		//If templates, format them
		if idx := strings.Index(entity, "-"); idx != -1 {
			entity = entity[:idx] + "_" + entity[idx+1:]
		}

		v, e3 = models.UpdateEntityBySlug(entity, name, &updateData)
	case e1 == true && e == true: // UPDATE NESTED
		println("updating nested")
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			u.Respond(w, u.Message(false, "Error while converting ID to ObjectID"))
			u.ErrLog("Error while converting ID to ObjectID", "DELETE ENTITY", "", r)
			return
		}

		//Get entity from URL
		idx := strings.SplitAfter(r.URL.Path, "/")[4]
		entity = idx[:len(idx)-2]

		v, e3 = models.UpdateNestedEntity(entity, objID, nest, updateData)

	case e == true: // UPDATE NORMAL
		println("updating Normale")
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			u.Respond(w, u.Message(false, "Error while converting ID to ObjectID"))
			u.ErrLog("Error while converting ID to ObjectID", "UPDATE ENTITY", "", r)
			return
		}

		//Get entity from URL and strip trailing 's'
		entity, _ = mux.Vars(r)["entity"]
		entity = entity[:len(entity)-1]
		println("OBJID:", objID.Hex())
		println("Entity;", entity)

		v, e3 = models.UpdateEntity(entity, objID, &updateData)

	default:
		println("ERRYO")
		w.WriteHeader(http.StatusBadRequest)
		u.Respond(w, u.Message(false, "Error while extracting from path parameters"))
		u.ErrLog("Error while extracting from path parameters", "UPDATE ENTITY", "", r)
		return
	}

	switch e3 {
	case "validate":
		w.WriteHeader(http.StatusBadRequest)
		u.ErrLog("Error while updating "+entity, "UPDATE "+strings.ToUpper(entity), e3, r)
	case "internal":
		w.WriteHeader(http.StatusInternalServerError)
		u.ErrLog("Error while updating "+entity, "UPDATE "+strings.ToUpper(entity), e3, r)
	case "mongo: no documents in result", "parent not found":
		w.WriteHeader(http.StatusNotFound)
		u.ErrLog("Error while updating "+entity, "UPDATE "+strings.ToUpper(entity), e3, r)
	default:
	}

	u.Respond(w, v)
}

// swagger:operation GET /api/{objs}? objects GetObject
// Gets an Object using any attribute (with the exception of description)
// via query in the system
// The attributes are in the form {attr}=xyz&{attr1}=abc
// And any combination can be used given that at least 1 is provided.
// ---
// produces:
// - application/json
// parameters:
// - name: objs
//   in: query
//   description: 'Indicates the object. Only values of "tenants", "sites",
//   "buildings", "rooms", "racks", "devices", "room-templates",
//   "obj-templates", "walls","acs","panels", "aisles", "tiles",
//   "cabinets", "groups", "corridors", "racksensors", and "devicesensors"
//   are acceptable'
//   required: true
//   type: string
//   default: "sites"
// - name: Name
//   in: query
//   description: Name of tenant
//   required: false
//   type: string
//   default: "INFINITI"
// - name: Category
//   in: query
//   description: Category of Tenant (ex. Consumer Electronics, Medical)
//   required: false
//   type: string
//   default: "Auto"
// - name: Domain
//   description: 'Domain of the Tenant'
//   required: false
//   type: string
//   default: "High End Auto"
// - name: Attributes
//   in: query
//   description: Any other object attributes can be queried
//   required: false
//   type: json
// responses:
//     '204':
//        description: Found
//     '404':
//        description: Not found
var GetEntityByQuery = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******************************************************")
	fmt.Println("FUNCTION CALL: 	 GetEntityByQuery ")
	fmt.Println("******************************************************")
	var data []map[string]interface{}
	var resp map[string]interface{}
	var bsonMap bson.M
	var e, entStr string

	if strings.Count(r.URL.Path, "/") == 3 { //We are querying NestedObjs
		query := u.ParamsParse(r.URL, u.EntityStrToInt(entStr))
		js, _ := json.Marshal(query)
		json.Unmarshal(js, &bsonMap)

		arr := strings.Split(r.URL.Path, "/")
		entStr = arr[2][:len(arr[2])-1]
		sub := arr[3][:len(arr[3])-1] //Not sure why this doesn't include rest of string

		data, e = models.GetNestedEntityByQuery(entStr, sub, bsonMap)
	} else {

		entStr = r.URL.Path[5 : len(r.URL.Path)-1]

		//If templates, format them
		if idx := strings.Index(entStr, "-"); idx != -1 {
			entStr = entStr[:idx] + "_" + entStr[idx+1:]
		}

		query := u.ParamsParse(r.URL, u.EntityStrToInt(entStr))
		js, _ := json.Marshal(query)
		json.Unmarshal(js, &bsonMap)

		data, e = models.GetEntityByQuery(entStr, bsonMap)

	}

	if len(data) == 0 {
		resp = u.Message(false, "Error: "+e)
		u.ErrLog("Error while getting "+entStr, "GET ENTITYQUERY", e, r)

		switch e {
		case "record not found":
			w.WriteHeader(http.StatusNotFound)
		case "":
			resp = u.Message(false, "Error: No Records Found")
			w.WriteHeader(http.StatusNotFound)
		default:
			resp = u.Message(false, "Error: No Records Found")
			w.WriteHeader(http.StatusNotFound)
		}

	} else {
		resp = u.Message(true, "success")
	}

	resp["data"] = map[string]interface{}{"objects": data}

	u.Respond(w, resp)
}

// swagger:operation GET /api/{obj}/{id}/{subent} objects GetFromObject
// Obtain all objects 2 levels lower in the system.
// For Example: /api/tenants/{id}/buildings
// Will return all buildings of a tenant
// Returns JSON body with all subobjects under the Object
// ---
// produces:
// - application/json
// parameters:
// - name: obj
//   in: query
//   description: 'Indicates the object. Only values of "tenants", "sites",
//   "buildings", "rooms" are acceptable'
//   required: true
//   type: string
//   default: "tenants"
// - name: ID
//   in: query
//   description: ID of object
//   required: true
//   type: int
//   default: 999
// - name: subent
//   in: query
//   description: Objects which 2 are levels lower in the hierarchy.
//   required: true
//   type: string
//   default: buildings
// responses:
//     '200':
//         description: Found
//     '404':
//         description: Nothing Found
var GetEntitiesOfAncestor = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******************************************************")
	fmt.Println("FUNCTION CALL: 	 GetEntitiesOfAncestor ")
	fmt.Println("******************************************************")
	var id string
	var e bool
	//Extract string between /api and /{id}
	idx := strings.Index(r.URL.Path[5:], "/") + 4
	entStr := r.URL.Path[5:idx]

	//s, _ := getObjID(id)
	enum := u.EntityStrToInt(entStr)
	//childBase := u.EntityToString(enum + 1)

	resp := u.Message(true, "success")

	if enum == TENANT {
		id, e = mux.Vars(r)["tenant_name"]
	} else {
		id, e = mux.Vars(r)["id"]
	}

	if e == false {
		u.Respond(w, u.Message(false, "Error while parsing path parameters"))
		u.ErrLog("Error while parsing path parameters", "GET CHILDRENOFPARENT", "", r)
		return
	}

	data, e1 := models.GetEntitiesOfAncestor(id, enum, entStr)
	if data == nil {
		resp = u.Message(false, "Error while getting "+entStr+"s: "+e1)
		u.ErrLog("Error while getting children of "+entStr,
			"GET CHILDRENOFPARENT", e1, r)

		switch e1 {
		case "record not found":
			w.WriteHeader(http.StatusNotFound)
		default:
		}

	} else {
		resp = u.Message(true, "success")
	}

	resp["data"] = map[string]interface{}{"objects": data}

	u.Respond(w, resp)
}

// swagger:operation GET /api/{objs}/{id}/all objects GetFromObject
// Obtain all objects related to specified object in the system.
// Returns JSON body with all subobjects under the Object
// ---
// produces:
// - application/json
// parameters:
// - name: objs
//   in: query
//   description: 'Indicates the object. Only values of "tenants", "sites",
//   "buildings", "rooms", "racks", "devices" are acceptable'
//   required: true
//   type: string
//   default: "sites"
// - name: ID
//   in: query
//   description: ID of object
//   required: true
//   type: int
//   default: 999
// responses:
//     '200':
//         description: Found
//     '404':
//         description: Nothing Found
var GetEntityHierarchy = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******************************************************")
	fmt.Println("FUNCTION CALL: 	 GetEntityHierarchy ")
	fmt.Println("******************************************************")
	//Extract string between /api and /{id}
	idx := strings.Index(r.URL.Path[5:], "/") + 4
	entity := r.URL.Path[5:idx]
	resp := u.Message(true, "success")
	var limit int
	var data map[string]interface{}
	var e1 string

	id, e := mux.Vars(r)["id"]
	if e == false {

		if entity != "tenant" {
			u.Respond(w, u.Message(false, "Error while parsing path parameters"))
			u.ErrLog("Error while parsing path parameters", "GET ENTITYHIERARCHY", "", r)
			return
		}
		id, e = mux.Vars(r)["tenant_name"]

		if e == false {
			u.Respond(w, u.Message(false, "Error while parsing tenant name"))
			u.ErrLog("Error while parsing path parameters", "GET ENTITYHIERARCHY", "", r)
			return
		}
	}

	if entity == "tenant" {

		_, e := models.GetEntityByName(entity, id)
		if e != "" {
			resp = u.Message(false, "Error while getting :"+entity+","+e)
			u.ErrLog("Error while getting "+entity, "GET "+entity, e, r)
		}

	}

	//Check if the request is a ranged hierarchy
	lastSlashIdx := strings.LastIndex(r.URL.Path, "/")
	indicator := r.URL.Path[lastSlashIdx+1:]
	switch indicator {
	case "all":
		//set to SUBDEV1
		limit = SUBDEV1
	case "nonstd":
		//special case
	default:
		//set to int equivalent
		//This strips the trailing s
		limit = u.EntityStrToInt(indicator[:len(indicator)-1])
	}
	println("Indicator: ", indicator)
	println("The limit is: ", limit)

	oID, _ := getObjID(id)

	entNum := u.EntityStrToInt(entity)

	println("Entity: ", entity, " & OID: ", oID.Hex())
	if entity == "device" {
		println("RETREIVE")
		data, e1 = models.RetrieveDeviceHierarch(oID)
	} else {
		data, e1 = models.GetEntityHierarchy(entity, oID, entNum, limit+1)
	}

	if data == nil {
		resp = u.Message(false, "Error while getting :"+entity+","+e1)
		u.ErrLog("Error while getting "+entity, "GET "+entity, e1, r)

		switch e1 {
		case "record not found":
			w.WriteHeader(http.StatusNotFound)
		default:
		}

	} else {
		resp = u.Message(true, "success")
	}

	resp["data"] = data
	u.Respond(w, resp)
}

// swagger:operation GET /api/tenants/{name}/all objects GetFromObject
// Obtain all objects related to Tenant in the system.
// Returns JSON body with all subobjects under the Tenant
// ---
// produces:
// - application/json
// parameters:
// - name: name
//   in: query
//   description: Name of Tenant
//   required: true
//   type: int
//   default: 999
// responses:
//     '200':
//         description: Found
//     '404':
//         description: Nothing Found
var GetTenantHierarchy = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******************************************************")
	fmt.Println("FUNCTION CALL: 	 GetTenantHierarchy ")
	fmt.Println("******************************************************")
	entity := "tenant"
	resp := u.Message(true, "success")
	var limit int

	id, e := mux.Vars(r)["tenant_name"]
	if e == false {
		u.Respond(w, u.Message(false, "Error while parsing tenant name"))
		u.ErrLog("Error while parsing path parameters", "GET TENANTHIERARCHY", "", r)
		return
	}

	//Check if the request is a ranged hierarchy
	lastSlashIdx := strings.LastIndex(r.URL.Path, "/")
	indicator := r.URL.Path[lastSlashIdx+1:]
	switch indicator {
	case "all":
		//set to SUBDEV1
		limit = SUBDEV1
	case "nonstd":
		//special case
	default:
		//set to int equivalent
		//This strips the trailing s
		limit = u.EntityStrToInt(indicator[:len(indicator)-1])
	}
	println("Indicator: ", indicator)
	println("The limit is: ", limit)

	data, e1 := models.GetTenantHierarchy(entity, id, TENANT, limit+1)

	if data == nil {
		resp = u.Message(false, "Error while getting :"+entity+","+e1)
		u.ErrLog("Error while getting "+entity, "GET "+entity, e1, r)

		switch e1 {
		case "record not found":
			w.WriteHeader(http.StatusNotFound)
		default:
		}

	} else {
		resp = u.Message(true, "success")
	}

	resp["data"] = data
	u.Respond(w, resp)
}

// swagger:operation GET /api/{objs}/{id}/* objects GetFromObect
// A category of objects of a Parent Object can be retrieved from the system.
// ---
// produces:
// - application/json
// parameters:
// - name: objs
//   in: query
//   description: 'Indicates the object. Only values of "tenants", "sites",
//   "buildings", "rooms", "racks", "devices" are acceptable'
//   required: true
//   type: string
//   default: "sites"
// - name: ID
//   in: path
//   description: ID of desired object
//   required: true
//   type: string
//   default: "INFINITI"
// - name: '*'
//   in: path
//   description: Hierarchal path to desired object(s). For rooms it can additionally have "acs" or "panels" or "walls"
//   required: true
//   type: string
//   default: "/buildings/BuildingB/RoomA"
// responses:
//     '200':
//         description: Found
//     '404':
//         description: Not Found
var GetEntitiesUsingNamesOfParents = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******************************************************")
	fmt.Println("FUNCTION CALL: 	 GetEntitiesUsingNamesOfParents ")
	fmt.Println("******************************************************")
	//Extract string between /api and /{id}
	idx := strings.Index(r.URL.Path[5:], "/") + 4
	entity := r.URL.Path[5:idx]
	resp := u.Message(true, "success")

	id, e := mux.Vars(r)["id"]
	tname, e1 := mux.Vars(r)["tenant_name"]
	if e == false && e1 == false {
		u.Respond(w, u.Message(false, "Error while parsing path parameters"))
		u.ErrLog("Error while parsing path parameters", "GET ENTITIESUSINGANCESTORNAMES", "", r)
		return
	}

	arr := (strings.Split(r.URL.Path, "/")[4:])
	ancestry := make([]map[string]string, 0)

	for i, k := range arr {
		if i%2 == 0 { //The keys (entities) are at the even indexes
			if i+1 >= len(arr) {
				ancestry = append(ancestry,
					map[string]string{k[:len(k)-1]: "all"})
			} else {
				ancestry = append(ancestry,
					map[string]string{k[:len(k)-1]: arr[i+1]})
			}
		}
	}

	oID, _ := getObjID(id)

	if len(arr)%2 != 0 { //This means we are getting entities
		var data []map[string]interface{}
		var e3 string
		if e1 == true {
			println("we are getting entities here")
			data, e3 = models.GetEntitiesUsingTenantAsAncestor(entity, tname, ancestry)

		} else {
			data, e3 = models.GetEntitiesUsingAncestorNames(entity, oID, ancestry)
		}

		if data == nil || len(data) == 0 {
			resp = u.Message(false, "Error while getting :"+entity+","+e3)
			u.ErrLog("Error while getting "+entity, "GET "+entity, e3, r)

			switch e3 {
			case "record not found":
				w.WriteHeader(http.StatusNotFound)
			default:
				w.WriteHeader(http.StatusNotFound)
			}

		} else {
			resp = u.Message(true, "success")
		}

		resp["data"] = map[string]interface{}{"objects": data}
		u.Respond(w, resp)
	} else { //We are only retrieving an entity
		var data map[string]interface{}
		var e3 string
		if e1 == true {
			data, e3 = models.GetEntityUsingTenantAsAncestor(entity, tname, ancestry)
		} else {
			data, e3 = models.GetEntityUsingAncestorNames(entity, oID, ancestry)
		}

		//data, e := models.GetEntityUsingAncestorNames(entity, oID, ancestry)
		if data == nil || len(data) == 0 {
			resp = u.Message(false, "Error while getting :"+entity+","+e3)
			u.ErrLog("Error while getting "+entity, "GET "+entity, e3, r)

			switch e3 {
			case "record not found":
				w.WriteHeader(http.StatusNotFound)
			default:
				w.WriteHeader(http.StatusNotFound)
			}

		} else {
			resp = u.Message(true, "success")
		}

		resp["data"] = data
		u.Respond(w, resp)
	}

}

var GetEntityHierarchyNonStd = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******************************************************")
	fmt.Println("FUNCTION CALL: 	 GetEntityHierarchyNonStd ")
	fmt.Println("******************************************************")
	var e, e1 bool
	var err string
	//Extract string between /api and /{id}
	idx := strings.Index(r.URL.Path[5:], "/") + 4
	entity := r.URL.Path[5:idx]

	id, e := mux.Vars(r)["id"]
	resp := u.Message(true, "success")
	data := map[string]interface{}{}
	//result := map[string][]map[string]interface{}{}

	if e == false {
		if id, e1 = mux.Vars(r)["tenant_name"]; e1 == false {
			u.Respond(w, u.Message(false, "Error while parsing Tpath parameters"))
			u.ErrLog("Error while parsing path parameters", "GETHIERARCHYNONSTD", "", r)
			return
		}
	}

	entNum := u.EntityStrToInt(entity)

	if entity == "tenant" {
		println("Getting TENANT HEIRARCHY")
		println("With ID: ", id)
		data, err = models.GetTenantHierarchy(entity, id, entNum, SUBDEV1)
		if err != "" {
			println("We have ERR")
		}
	} else {
		oID, _ := getObjID(id)
		data, err = models.GetEntityHierarchy(entity, oID, entNum, SUBDEV1)
	}

	if data == nil {
		resp = u.Message(false, "Error while getting NonStandard Hierarchy: "+err)
		u.ErrLog("Error while getting NonStdHierarchy", "GETNONSTDHIERARCHY", err, r)

		switch err {
		case "record not found":
			w.WriteHeader(http.StatusNotFound)
		default:
		}

	} else {
		resp = u.Message(true, "success")
		result := parseDataForNonStdResult(entity, entNum, data)
		resp["data"] = result
		//u.Respond(w, resp)
	}

	//resp["data"] = data
	/*resp["data"] = sites
	resp["buildings"] = bldgs
	resp["rooms"] = rooms
	resp["racks"] = racks
	resp["devices"] = devices*/
	u.Respond(w, resp)
}
