package controllers

import (
	"github.com/gkganesh126/recipe-sharing-platform/common"
	"gopkg.in/mgo.v2"
)

// Struct used for maintaining HTTP Request Context
type Context struct {
	MongoSession *mgo.Session
}

// Close mgo.Session
func (c *Context) Close() {
	c.MongoSession.Close()
}

// Returns mgo.collection for the given name
func (c *Context) RecipeSharingPlatformDbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.RecipeSharingPlatformDatabase).C(name)
}

// Create a new Context object for each HTTP request
func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}
