package profile

import (
	"github.com/PharbersDeveloper/max-go/phmodel/company"
	"github.com/alfredyang1986/blackmirror/bmconfighandle"
	"github.com/alfredyang1986/blackmirror/bmmodel"
	"github.com/alfredyang1986/blackmirror/bmmodel/request"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PhProfile struct {
	Id  string        `json:"id"`
	Id_ bson.ObjectId `bson:"_id"`

	Username string            `json:"username" bson:"username"`
	Password string            `json:"password" bson:"password"`
	Company  company.PhCompany `json:"Company" jsonapi:"relationships"`
}

/*------------------------------------------------
 * bm object interface
 *------------------------------------------------*/

func (bd *PhProfile) ResetIdWithId_() {
	bmmodel.ResetIdWithId_(bd)
}

func (bd *PhProfile) ResetId_WithID() {
	bmmodel.ResetId_WithID(bd)
}

/*------------------------------------------------
 * bmobject interface
 *------------------------------------------------*/

func (bd *PhProfile) QueryObjectId() bson.ObjectId {
	return bd.Id_
}

func (bd *PhProfile) QueryId() string {
	return bd.Id
}

func (bd *PhProfile) SetObjectId(id_ bson.ObjectId) {
	bd.Id_ = id_
}

func (bd *PhProfile) SetId(id string) {
	bd.Id = id
}

/*------------------------------------------------
 * relationships interface
 *------------------------------------------------*/
func (bd PhProfile) SetConnect(tag string, v interface{}) interface{} {
	switch tag {
	case "Company":
		bd.Company = v.(company.PhCompany)
	}
	return bd
}

func (bd PhProfile) QueryConnect(tag string) interface{} {
	switch tag {
	case "Company":
		return bd.Company
	}
	return bd
}

/*------------------------------------------------
 * mongo interface
 *------------------------------------------------*/

func (bd *PhProfile) InsertBMObject() error {
	return bmmodel.InsertBMObject(bd)
}

func (bd *PhProfile) FindOne(req request.Request) error {
	return bmmodel.FindOne(req, bd)
}

func (bd *PhProfile) UpdateBMObject(req request.Request) error {
	return bmmodel.UpdateOne(req, bd)
}

/*------------------------------------------------
 * profile interface
 *------------------------------------------------*/

func (bd PhProfile) IsUserRegisted() bool {
	var bmMongoConfig bmconfig.BMMongoConfig
	bmMongoConfig.GenerateConfig()
	session, err := mgo.Dial(bmMongoConfig.Host + ":" + bmMongoConfig.Port)
	if err != nil {
		return true
	}
	defer session.Close()

	c := session.DB("test").C("PhProfile")
	n, err := c.Find(bson.M{"username": bd.Username}).Count()
	if err != nil {
		panic(err)
	}

	return n > 0
}

func (bd PhProfile) Valid() bool {
	return bd.Username != ""
}
