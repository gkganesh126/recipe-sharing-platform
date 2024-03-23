package data

import (
	"fmt"

	"github.com/gkganesh126/recipe-sharing-platform/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	C *mgo.Collection
}

func (r *UserRepository) Create(user *models.User) error {

	isAlreadyRegistered, alreadyRegisteredUserID := r.IsRegistered(user.Username)
	if isAlreadyRegistered {
		user.UserID = bson.ObjectIdHex(alreadyRegisteredUserID)

		//return nil
	} else {
		objID := bson.NewObjectId()
		user.UserID = objID
	}
	_, err := r.C.UpsertId(user.UserID, user)
	//err := r.C.Insert(&user)
	return err
}

func (r *UserRepository) GetAll() []models.User {
	var users []models.User
	iter := r.C.Find(nil).Iter()
	result := models.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (r *UserRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *UserRepository) SetDelete(id string) error {
	err := r.C.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"delete": true}})
	return err
}

func (r *UserRepository) Update(user *models.User) error {
	return r.C.Update(bson.M{"_id": user.UserID}, bson.M{"$set": bson.M{"username": user.Username}})
}

func (r *UserRepository) GetUserID(username string) (models.OnlyUserIDBson, error) {
	var userID models.OnlyUserIDBson
	err := r.C.Find(bson.M{"username": username}).Select(bson.M{"_id": 1}).One(&userID)
	if err != nil {
		return models.OnlyUserIDBson{}, err
	}
	return userID, nil
}

func (r *UserRepository) IsRegistered(mobNum string) (bool, string) {
	userID, err := r.GetUserID(mobNum)
	if err != nil {
		return false, ""
	}
	if userID.UserID.Hex() != "" {
		return true, userID.UserID.Hex()
	}
	return false, ""
}

func (r *UserRepository) GetSession(userID string) (string, error) {
	var session models.OnlySession
	err := r.C.Find(bson.M{"_id": bson.ObjectIdHex(userID)}).Select(bson.M{"session": 1}).One(&session)
	if err != nil {
		return "", err
	}
	return session.Session, nil
}

func (r *UserRepository) UpdateSession(userID string, session string) error {
	fmt.Println("userID: ", userID, " session: ", session)
	return r.C.Update(bson.M{"_id": bson.ObjectIdHex(userID)}, bson.M{"$set": bson.M{"session": session}})
}

func (r *UserRepository) GetUserName(userID string) (string, error) {
	var onlyUsername models.OnlyUsername
	err := r.C.Find(bson.M{"_id": bson.ObjectIdHex(userID)}).Select(bson.M{"username": 1}).One(&onlyUsername)
	if err != nil {
		return "", err
	}
	return onlyUsername.Username, nil
}
