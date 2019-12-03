package model

// Code generated by goradd. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/goradd/goradd/pkg/orm/broadcast"
	"github.com/goradd/goradd/pkg/orm/db"
	"github.com/goradd/goradd/pkg/orm/op"
	. "github.com/goradd/goradd/pkg/orm/op"
	"github.com/goradd/goradd/pkg/orm/query"
	"github.com/goradd/goradd/pkg/stringmap"
	"github.com/goradd/goradd/web/examples/gen/goradd/model/node"

	//"./node"
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// personBase is a base structure to be embedded in a "subclass" and provides the ORM access to the database.
// Do not directly access the internal variables, but rather use the accessor functions, since this class maintains internal state
// related to the variables.

type personBase struct {
	id        string
	idIsValid bool
	idIsDirty bool

	firstName        string
	firstNameIsValid bool
	firstNameIsDirty bool

	lastName        string
	lastNameIsValid bool
	lastNameIsDirty bool

	// Reverse reference objects.
	oProjectsAsManager        []*Project          // Objects in the order they were queried
	mProjectsAsManager        map[string]*Project // Objects by PK
	oProjectsAsManagerIsDirty bool
	oLogin                    *Login
	oLoginIsDirty             bool
	oEmployeeInfo             *EmployeeInfo
	oEmployeeInfoIsDirty      bool
	oAddresses                []*Address          // Objects in the order they were queried
	mAddresses                map[string]*Address // Objects by PK
	oAddressesIsDirty         bool

	// Many-Many reference objects.
	oProjectsAsTeamMember []*Project
	mProjectsAsTeamMember map[string]*Project // Objects by PK
	oPersonTypes          []PersonType

	// Custom aliases, if specified
	_aliases map[string]interface{}

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update
	_restored bool
}

const (
	PersonIDDefault        = ""
	PersonFirstNameDefault = ""
	PersonLastNameDefault  = ""
)

const (
	Person_ID               = `ID`
	Person_FirstName        = `FirstName`
	Person_LastName         = `LastName`
	PersonProjectsAsManager = `ProjectsAsManager`

	PersonLogin = `Login`

	PersonEmployeeInfo         = `EmployeeInfo`
	PersonAddresses            = `Addresses`
	PersonProjectAsTeamMember  = `ProjectAsTeamMember`
	PersonProjectsAsTeamMember = `ProjectsAsTeamMember`
	PersonPersonType           = `PersonType`
	PersonPersonTypes          = `PersonTypes`
)

// Initialize or re-initialize a Person database object to default values.
func (o *personBase) Initialize() {

	o.id = ""
	o.idIsValid = false
	o.idIsDirty = false

	o.firstName = ""
	o.firstNameIsValid = false
	o.firstNameIsDirty = false

	o.lastName = ""
	o.lastNameIsValid = false
	o.lastNameIsDirty = false

	o._restored = false
}

func (o *personBase) PrimaryKey() string {
	return o.id
}

// ID returns the loaded value of ID.
func (o *personBase) ID() string {
	return fmt.Sprint(o.id)
}

// IDIsValid returns true if the value was loaded from the database or has been set.
func (o *personBase) IDIsValid() bool {
	return o._restored && o.idIsValid
}

func (o *personBase) FirstName() string {
	if o._restored && !o.firstNameIsValid {
		panic("firstName was not selected in the last query and so is not valid")
	}
	return o.firstName
}

// FirstNameIsValid returns true if the value was loaded from the database or has been set.
func (o *personBase) FirstNameIsValid() bool {
	return o.firstNameIsValid
}

// SetFirstName sets the value of FirstName in the object, to be saved later using the Save() function.
func (o *personBase) SetFirstName(v string) {
	o.firstNameIsValid = true
	if o.firstName != v || !o._restored {
		o.firstName = v
		o.firstNameIsDirty = true
	}

}

func (o *personBase) LastName() string {
	if o._restored && !o.lastNameIsValid {
		panic("lastName was not selected in the last query and so is not valid")
	}
	return o.lastName
}

// LastNameIsValid returns true if the value was loaded from the database or has been set.
func (o *personBase) LastNameIsValid() bool {
	return o.lastNameIsValid
}

// SetLastName sets the value of LastName in the object, to be saved later using the Save() function.
func (o *personBase) SetLastName(v string) {
	o.lastNameIsValid = true
	if o.lastName != v || !o._restored {
		o.lastName = v
		o.lastNameIsDirty = true
	}

}

// GetAlias returns the alias for the given key.
func (o *personBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
		return query.NewAliasValue([]byte{})
	}
}

// ProjectAsTeamMember returns a single Project object, if one was loaded
// otherwise, it will return nil.
func (o *personBase) ProjectAsTeamMember() *Project {
	if o.oProjectsAsTeamMember == nil {
		return nil
	}
	return o.oProjectsAsTeamMember[0]
}

// ProjectsAsTeamMember returns a slice of Project objects if loaded. If not loaded, will return nil.
func (o *personBase) ProjectsAsTeamMember() []*Project {
	if o.oProjectsAsTeamMember == nil {
		return nil
	}
	return o.oProjectsAsTeamMember
}

// PersonTypes returns a slice of PersonType objects if loaded.
func (o *personBase) PersonTypes() []PersonType {
	if o.oPersonTypes == nil {
		return nil
	}
	return o.oPersonTypes
}

// PersonType returns a single PersonType object, if one was loaded
// otherwise, it will return zero.
func (o *personBase) PersonType() PersonType {
	if o.oPersonTypes == nil {
		return 0
	}
	return o.oPersonTypes[0]
}

// ProjectAsManager returns a single Project object by primary key, if one was loaded.
// Otherwise, it will return nil. It will not return Project objects that are not saved.
func (o *personBase) ProjectAsManager(pk string) *Project {
	if o.mProjectsAsManager == nil {
		return nil
	}
	v, _ := o.mProjectsAsManager[pk]
	return v
}

// ProjectsAsManager returns a slice of Project objects if loaded.
func (o *personBase) ProjectsAsManager() []*Project {
	if o.oProjectsAsManager == nil {
		return nil
	}
	return o.oProjectsAsManager
}

// LoadProjectsAsManager loads a new slice of Project objects and returns it.
func (o *personBase) LoadProjectsAsManager(ctx context.Context, conditions ...interface{}) []*Project {
	qb := queryProjects(ctx)
	cond := Equal(node.Project().ManagerID(), o.PrimaryKey())
	if conditions != nil {
		conditions = append(conditions, cond)
		cond = And(conditions...)
	}

	o.oProjectsAsManager = qb.Where(cond).Load(ctx)
	return o.oProjectsAsManager
}

// SetProjectsAsManager associates the given objects with the Person.
// If it has items already associated with it that will not be associated after a save,
// the foreign keys for those will be set to null.
// If you did not use a join to query the items in the first place, used a conditional join,
// or joined with an expansion, be particularly careful, since you may be changing items
// that are not currently attached to this Person.
func (o *personBase) SetProjectsAsManager(objs []*Project) {
	for _, obj := range o.oProjectsAsManager {
		if obj.IsDirty() {
			panic("You cannot overwrite items that have changed but have not been saved.")
		}
	}

	o.oProjectsAsManager = objs
	for _, obj := range o.oProjectsAsManager {
		pk := obj.ID()
		if pk != "" {
			if o.mProjectsAsManager == nil {
				o.mProjectsAsManager = make(map[string]*Project)
			}
			o.mProjectsAsManager[pk] = obj
		}
	}
	o.oProjectsAsManagerIsDirty = true
}

// Login returns the connected Login object, if one was loaded
// otherwise, it will return nil.
func (o *personBase) Login() *Login {
	if o.oLogin == nil {
		return nil
	}
	return o.oLogin
}

// LoadLogin returns the connected Login object, if one was loaded
// otherwise, it will return nil.
func (o *personBase) LoadLogin(ctx context.Context) *Login {
	if o.oLogin == nil {
		o.oLogin = LoadLoginByPersonID(ctx, o.ID())
	}
	return o.oLogin
}

// SetLogin associates the given object with the Person.
// If it has an item already associated with it,
// the foreign key for that item will be set to null.
// If you did not use a join to query the items in the first place, used a conditional join,
// or joined with an expansion, be particularly careful, since you may be changing an item
// that is not currently attached to this Person.
func (o *personBase) SetLogin(obj *Login) {
	if o.oLogin != nil && o.oLogin.IsDirty() {
		panic("The Login has changed. You must save it first before changing to a different one.")
	}
	o.oLogin = obj
	o.oLoginIsDirty = true
}

// EmployeeInfo returns the connected EmployeeInfo object, if one was loaded
// otherwise, it will return nil.
func (o *personBase) EmployeeInfo() *EmployeeInfo {
	if o.oEmployeeInfo == nil {
		return nil
	}
	return o.oEmployeeInfo
}

// LoadEmployeeInfo returns the connected EmployeeInfo object, if one was loaded
// otherwise, it will return nil.
func (o *personBase) LoadEmployeeInfo(ctx context.Context) *EmployeeInfo {
	if o.oEmployeeInfo == nil {
		o.oEmployeeInfo = LoadEmployeeInfoByPersonID(ctx, o.ID())
	}
	return o.oEmployeeInfo
}

// SetEmployeeInfo associates the given object with the Person.
// WARNING! If it has an item already associated with it,
// that item will be DELETED since it cannot be null.
// If you did not use a join to query the items in the first place, used a conditional join,
// or joined with an expansion, be particularly careful, since you may be changing an item
// that is not currently attached to this Person.
func (o *personBase) SetEmployeeInfo(obj *EmployeeInfo) {
	o.oEmployeeInfo = obj
	o.oEmployeeInfoIsDirty = true
}

// Address returns a single Address object by primary key, if one was loaded.
// Otherwise, it will return nil. It will not return Address objects that are not saved.
func (o *personBase) Address(pk string) *Address {
	if o.mAddresses == nil {
		return nil
	}
	v, _ := o.mAddresses[pk]
	return v
}

// Addresses returns a slice of Address objects if loaded.
func (o *personBase) Addresses() []*Address {
	if o.oAddresses == nil {
		return nil
	}
	return o.oAddresses
}

// LoadAddresses loads a new slice of Address objects and returns it.
func (o *personBase) LoadAddresses(ctx context.Context, conditions ...interface{}) []*Address {
	qb := queryAddresses(ctx)
	cond := Equal(node.Address().PersonID(), o.PrimaryKey())
	if conditions != nil {
		conditions = append(conditions, cond)
		cond = And(conditions...)
	}

	o.oAddresses = qb.Where(cond).Load(ctx)
	return o.oAddresses
}

// SetAddresses associates the given objects with the Person.
// WARNING! If it has items already associated with it that will not be associated after a save,
// those items will be DELETED since they cannot be null.
// If you did not use a join to query the items in the first place, used a conditional join,
// or joined with an expansion, be particularly careful, since you may be changing items
// that are not currently attached to this Person.
func (o *personBase) SetAddresses(objs []*Address) {
	for _, obj := range o.oAddresses {
		if obj.IsDirty() {
			panic("You cannot overwrite items that have changed but have not been saved.")
		}
	}

	o.oAddresses = objs
	for _, obj := range o.oAddresses {
		pk := obj.ID()
		if pk != "" {
			if o.mAddresses == nil {
				o.mAddresses = make(map[string]*Address)
			}
			o.mAddresses[pk] = obj
		}
	}
	o.oAddressesIsDirty = true
}

// Load returns a Person from the database.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See Join() and Select() for more info.
func LoadPerson(ctx context.Context, primaryKey string, joinOrSelectNodes ...query.NodeI) *Person {
	return queryPeople(ctx).Where(Equal(node.Person().ID(), primaryKey)).joinOrSelect(joinOrSelectNodes...).Get(ctx)
}

// The PeopleBuilder uses the QueryBuilderI interface from the database to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, Count, or Delete
type PeopleBuilder struct {
	base                query.QueryBuilderI
	hasConditionalJoins bool
}

func newPersonBuilder() *PeopleBuilder {
	b := &PeopleBuilder{
		base: db.GetDatabase("goradd").
			NewBuilder(),
	}
	return b.Join(node.Person())
}

// Load terminates the query builder, performs the query, and returns a slice of Person objects. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice
func (b *PeopleBuilder) Load(ctx context.Context) (personSlice []*Person) {
	results := b.base.Load(ctx)
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Person)
		o.load(item, !b.hasConditionalJoins, o, nil, "")
		personSlice = append(personSlice, o)
	}
	return personSlice
}

// LoadI terminates the query builder, performs the query, and returns a slice of interfaces. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice.
func (b *PeopleBuilder) LoadI(ctx context.Context) (personSlice []interface{}) {
	results := b.base.Load(ctx)
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Person)
		o.load(item, !b.hasConditionalJoins, o, nil, "")
		personSlice = append(personSlice, o)
	}
	return personSlice
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
func (b *PeopleBuilder) Get(ctx context.Context) *Person {
	results := b.Load(ctx)
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

// Expand expands an array type node so that it will produce individual rows instead of an array of items
func (b *PeopleBuilder) Expand(n query.NodeI) *PeopleBuilder {
	b.base.Expand(n)
	return b
}

// Join adds a node to the node tree so that its fields will appear in the query. Optionally add conditions to filter
// what gets included. The conditions will be AND'd with the basic condition matching the primary keys of the join.
func (b *PeopleBuilder) Join(n query.NodeI, conditions ...query.NodeI) *PeopleBuilder {
	var condition query.NodeI
	if len(conditions) > 1 {
		condition = And(conditions)
	} else if len(conditions) == 1 {
		condition = conditions[0]
	}
	b.base.Join(n, condition)
	if condition != nil {
		b.hasConditionalJoins = true
	}
	return b
}

// Where adds a condition to filter what gets selected.
func (b *PeopleBuilder) Where(c query.NodeI) *PeopleBuilder {
	b.base.Condition(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
func (b *PeopleBuilder) OrderBy(nodes ...query.NodeI) *PeopleBuilder {
	b.base.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified
func (b *PeopleBuilder) Limit(maxRowCount int, offset int) *PeopleBuilder {
	b.base.Limit(maxRowCount, offset)
	return b
}

// Select optimizes the query to only return the specified fields. Once you put a Select in your query, you must
// specify all the fields that you will eventually read out. Be careful when selecting fields in joined tables, as joined
// tables will also contain pointers back to the parent table, and so the parent node should have the same field selected
// as the child node if you are querying those fields.
func (b *PeopleBuilder) Select(nodes ...query.NodeI) *PeopleBuilder {
	b.base.Select(nodes...)
	return b
}

// Alias lets you add a node with a custom name. After the query, you can read out the data using GetAlias() on a
// returned object. Alias is useful for adding calculations or subqueries to the query.
func (b *PeopleBuilder) Alias(name string, n query.NodeI) *PeopleBuilder {
	b.base.Alias(name, n)
	return b
}

// Distinct removes duplicates from the results of the query. Adding a Select() may help you get to the data you want, although
// using Distinct with joined tables is often not effective, since we force joined tables to include primary keys in the query, and this
// often ruins the effect of Distinct.
func (b *PeopleBuilder) Distinct() *PeopleBuilder {
	b.base.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions in an Alias() call.
func (b *PeopleBuilder) GroupBy(nodes ...query.NodeI) *PeopleBuilder {
	b.base.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query.
func (b *PeopleBuilder) Having(node query.NodeI) *PeopleBuilder {
	b.base.Having(node)
	return b
}

// Count terminates a query and returns just the number of items selected.
func (b *PeopleBuilder) Count(ctx context.Context, distinct bool, nodes ...query.NodeI) uint {
	return b.base.Count(ctx, distinct, nodes...)
}

// Delete uses the query builder to delete a group of records that match the criteria
func (b *PeopleBuilder) Delete(ctx context.Context) {
	b.base.Delete(ctx)
}

// Subquery uses the query builder to define a subquery within a larger query. You MUST include what
// you are selecting by adding Alias or Select functions on the subquery builder. Generally you would use
// this as a node to an Alias function on the surrounding query builder.
func (b *PeopleBuilder) Subquery() *query.SubqueryNode {
	return b.base.Subquery()
}

// joinOrSelect is a private helper function for the Load* functions
func (b *PeopleBuilder) joinOrSelect(nodes ...query.NodeI) *PeopleBuilder {
	for _, n := range nodes {
		switch n.(type) {
		case query.TableNodeI:
			b.base.Join(n, nil)
		case *query.ColumnNode:
			b.Select(n)
		}
	}
	return b
}

func CountPersonByID(ctx context.Context, id string) uint {
	return queryPeople(ctx).Where(Equal(node.Person().ID(), id)).Count(ctx, false)
}

func CountPersonByFirstName(ctx context.Context, firstName string) uint {
	return queryPeople(ctx).Where(Equal(node.Person().FirstName(), firstName)).Count(ctx, false)
}

func CountPersonByLastName(ctx context.Context, lastName string) uint {
	return queryPeople(ctx).Where(Equal(node.Person().LastName(), lastName)).Count(ctx, false)
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
// If linkParent is true we will have child relationships use a pointer back to the parent object. If false, it will create a separate object.
// Care must be taken in the query, as Select clauses might not be honored if the child object has fields selected which the parent object does not have.
// Also, if any joins are conditional, that might affect which child objects are included, so in this situation, linkParent should be false
func (o *personBase) load(m map[string]interface{}, linkParent bool, objThis *Person, objParent interface{}, parentKey string) {
	if v, ok := m["id"]; ok && v != nil {
		if o.id, ok = v.(string); ok {
			o.idIsValid = true
			o.idIsDirty = false
		} else {
			panic("Wrong type found for id.")
		}
	} else {
		o.idIsValid = false
		o.id = ""
	}

	if v, ok := m["first_name"]; ok && v != nil {
		if o.firstName, ok = v.(string); ok {
			o.firstNameIsValid = true
			o.firstNameIsDirty = false
		} else {
			panic("Wrong type found for first_name.")
		}
	} else {
		o.firstNameIsValid = false
		o.firstName = ""
	}

	if v, ok := m["last_name"]; ok && v != nil {
		if o.lastName, ok = v.(string); ok {
			o.lastNameIsValid = true
			o.lastNameIsDirty = false
		} else {
			panic("Wrong type found for last_name.")
		}
	} else {
		o.lastNameIsValid = false
		o.lastName = ""
	}

	if v, ok := m["ProjectsAsTeamMember"]; ok {
		if oProjectsAsTeamMember, ok2 := v.([]db.ValueMap); ok2 {
			o.oProjectsAsTeamMember = []*Project{}
			o.mProjectsAsTeamMember = map[string]*Project{}

			for _, v2 := range oProjectsAsTeamMember {
				obj := new(Project)
				obj.load(v2, linkParent, obj, objThis, "TeamMembers")
				if linkParent && parentKey == "ProjectsAsTeamMember" && obj.id == objParent.(*Project).id {
					obj = objParent.(*Project)
				}
				o.oProjectsAsTeamMember = append(o.oProjectsAsTeamMember, obj)
				o.mProjectsAsTeamMember[obj.PrimaryKey()] = obj
			}
		} else {
			panic("Wrong type found for oProjectsAsTeamMember object.")
		}
	} else {
		o.oProjectsAsTeamMember = nil
	}

	if v, ok := m["PersonTypes"]; ok {
		if oPersonTypes, ok2 := v.([]uint); ok2 {
			o.oPersonTypes = []PersonType{}
			for _, m := range oPersonTypes {
				o.oPersonTypes = append(o.oPersonTypes, PersonType(m))
			}
		} else {
			panic("Wrong type found for oPersonTypes object.")
		}
	} else {
		o.oPersonTypes = nil
	}

	if v, ok := m["ProjectsAsManager"]; ok {
		switch oProjectsAsManager := v.(type) {
		case []db.ValueMap:
			o.oProjectsAsManager = make([]*Project, 0, len(oProjectsAsManager))
			o.mProjectsAsManager = make(map[string]*Project, len(oProjectsAsManager))
			for _, v2 := range oProjectsAsManager {
				obj := new(Project)
				obj.load(v2, linkParent, obj, objThis, "Manager")
				if linkParent && parentKey == "ProjectsAsManager" && obj.managerID == objParent.(*Project).managerID {
					obj = objParent.(*Project)
				}
				o.oProjectsAsManager = append(o.oProjectsAsManager, obj)
				o.mProjectsAsManager[obj.PrimaryKey()] = obj
				o.oProjectsAsManagerIsDirty = false
			}
		case db.ValueMap: // single expansion
			obj := new(Project)
			obj.load(oProjectsAsManager, linkParent, obj, objThis, "Manager")
			if linkParent && parentKey == "ProjectsAsManager" && obj.managerID == objParent.(*Project).managerID {
				obj = objParent.(*Project)
			}
			o.oProjectsAsManager = []*Project{obj}
			o.oProjectsAsManagerIsDirty = false
		default:
			panic("Wrong type found for oProjectsAsManager object.")
		}
	} else {
		o.oProjectsAsManager = nil
		o.oProjectsAsManagerIsDirty = false
	}

	if v, ok := m["Login"]; ok {
		if oLogin, ok2 := v.(db.ValueMap); ok2 {
			o.oLogin = new(Login)
			o.oLogin.load(oLogin, linkParent, o.oLogin, objThis, "Person")
			o.oLoginIsDirty = false
		} else {
			panic("Wrong type found for oLogin object.")
		}
	} else {
		o.oLogin = nil
		o.oLoginIsDirty = false
	}

	if v, ok := m["EmployeeInfo"]; ok {
		if oEmployeeInfo, ok2 := v.(db.ValueMap); ok2 {
			o.oEmployeeInfo = new(EmployeeInfo)
			o.oEmployeeInfo.load(oEmployeeInfo, linkParent, o.oEmployeeInfo, objThis, "Person")
			o.oEmployeeInfoIsDirty = false
		} else {
			panic("Wrong type found for oEmployeeInfo object.")
		}
	} else {
		o.oEmployeeInfo = nil
		o.oEmployeeInfoIsDirty = false
	}

	if v, ok := m["Addresses"]; ok {
		switch oAddresses := v.(type) {
		case []db.ValueMap:
			o.oAddresses = make([]*Address, 0, len(oAddresses))
			o.mAddresses = make(map[string]*Address, len(oAddresses))
			for _, v2 := range oAddresses {
				obj := new(Address)
				obj.load(v2, linkParent, obj, objThis, "Person")
				if linkParent && parentKey == "Addresses" && obj.personID == objParent.(*Address).personID {
					obj = objParent.(*Address)
				}
				o.oAddresses = append(o.oAddresses, obj)
				o.mAddresses[obj.PrimaryKey()] = obj
				o.oAddressesIsDirty = false
			}
		case db.ValueMap: // single expansion
			obj := new(Address)
			obj.load(oAddresses, linkParent, obj, objThis, "Person")
			if linkParent && parentKey == "Addresses" && obj.personID == objParent.(*Address).personID {
				obj = objParent.(*Address)
			}
			o.oAddresses = []*Address{obj}
			o.oAddressesIsDirty = false
		default:
			panic("Wrong type found for oAddresses object.")
		}
	} else {
		o.oAddresses = nil
		o.oAddressesIsDirty = false
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = map[string]interface{}(v.(db.ValueMap))
	}
	o._restored = true
}

// Save will update or insert the object, depending on the state of the object.
// If it has any auto-generated ids, those will be updated.
func (o *personBase) Save(ctx context.Context) {
	if o._restored {
		if !o.IsDirty() {
			return
		}
		o.Update(ctx)
	} else {
		o.Insert(ctx)
	}
}

// Update will update the values in the database, saving any changed values.
func (o *personBase) Update(ctx context.Context) {
	if !o._restored {
		panic("Cannot update a record that was not originally read from the database.")
	}
	m := o.getModifiedFields()
	if len(m) == 0 {
		return
	}
	d := db.GetDatabase("goradd")
	txid := d.Begin(ctx)
	defer d.Rollback(ctx, txid)
	d.Update(ctx, "person", m, "id", fmt.Sprint(o.id))

	if o.oProjectsAsManagerIsDirty {
		objs := QueryProjects(ctx).
			Where(op.Equal(node.Project().ManagerID(), o.PrimaryKey())).
			Load(ctx)
		// TODO:select only the required fields
		for _, obj := range objs {
			if _, ok := o.mProjectsAsManager[obj.PrimaryKey()]; !ok {
				// The old object is not in the group of new objects
				obj.SetManagerID(nil)
				obj.Save(ctx)
			}
		}
		for _, obj := range o.oProjectsAsManager {
			obj.SetManagerID(o.PrimaryKey())
			obj.Save(ctx)
		}

	}
	if o.oLoginIsDirty {
		obj := QueryLogins(ctx).
			Where(op.Equal(node.Login().PersonID(), o.PrimaryKey())).
			Get(ctx)
		if obj != nil && obj.PrimaryKey() != o.oLogin.PrimaryKey() {
			obj.SetPersonID(nil)
			obj.Save(ctx)
		}
		o.oLogin.SetPersonID(o.PrimaryKey())
		o.oLogin.Save(ctx)
	}
	if o.oEmployeeInfoIsDirty {

		// Since the other side of the relationship cannot be null, the object to be detached must be deleted
		obj := QueryEmployeeInfos(ctx).
			Where(op.Equal(node.EmployeeInfo().PersonID(), o.PrimaryKey())).
			Get(ctx)
		if obj != nil && obj.PrimaryKey() != o.oEmployeeInfo.PrimaryKey() {
			obj.Delete(ctx)
		}
		o.oEmployeeInfo.SetPersonID(o.PrimaryKey())
		o.oEmployeeInfo.Save(ctx)
	}
	if o.oAddressesIsDirty {

		// Since the other side of the relationship cannot be null, the objects to be detached must be deleted
		// We take care to only delete objects that are not being reattached
		objs := QueryAddresses(ctx).
			Where(op.Equal(node.Address().PersonID(), o.PrimaryKey())).
			Load(ctx)
		// TODO: select only the required fields
		for _, obj := range objs {
			if _, ok := o.mAddresses[obj.PrimaryKey()]; !ok {
				// The old object is not in the group of new objects
				obj.Delete(ctx)
			}
		}
		for _, obj := range o.oAddresses {
			obj.SetPersonID(o.PrimaryKey())
			obj.Save(ctx)
		}
	}
	d.Commit(ctx, txid)
	o.resetDirtyStatus()
	broadcast.Update(ctx, "goradd", "person", fmt.Sprintf("%v", o.id), stringmap.SortedKeys(m)...)
}

// Insert forces the object to be inserted into the database. If the object was loaded from the database originally,
// this will create a duplicate in the database.
func (o *personBase) Insert(ctx context.Context) {
	m := o.getModifiedFields()
	if len(m) == 0 {
		return
	}
	d := db.GetDatabase("goradd")
	txid := d.Begin(ctx)
	defer d.Rollback(ctx, txid)

	id := d.Insert(ctx, "person", m)
	o.id = id
	if o.oProjectsAsManagerIsDirty {

		for _, obj := range o.oProjectsAsManager {
			obj.SetManagerID(id)
			obj.Save(ctx)
			if o.mProjectsAsManager == nil {
				o.mProjectsAsManager = make(map[string]*Project)
			}
			o.mProjectsAsManager[obj.PrimaryKey()] = obj
		}
	}
	if o.oLoginIsDirty {

		o.oLogin.SetPersonID(id)
		o.oLogin.Save(ctx)
	}
	if o.oEmployeeInfoIsDirty {

		o.oEmployeeInfo.SetPersonID(id)
		o.oEmployeeInfo.Save(ctx)
	}
	if o.oAddressesIsDirty {

		for _, obj := range o.oAddresses {
			obj.SetPersonID(id)
			obj.Save(ctx)
			if o.mAddresses == nil {
				o.mAddresses = make(map[string]*Address)
			}
			o.mAddresses[obj.PrimaryKey()] = obj
		}
	}
	d.Commit(ctx, txid)
	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd", "person", fmt.Sprint(o.id))
}

func (o *personBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}

	if o.firstNameIsDirty {
		fields["first_name"] = o.firstName
	}

	if o.lastNameIsDirty {
		fields["last_name"] = o.lastName
	}

	return
}

// Delete deletes the associated record from the database.
func (o *personBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := db.GetDatabase("goradd")
	txid := d.Begin(ctx)
	defer d.Rollback(ctx, txid)
	obj := QueryEmployeeInfos(ctx).
		Where(op.Equal(node.EmployeeInfo().PersonID(), o.PrimaryKey())).
		Select(node.EmployeeInfo().PrimaryKeyNode()).
		Get(ctx)
	if obj != nil {
		obj.Delete(ctx)
	}
	o.oEmployeeInfo = nil
	objs := QueryAddresses(ctx).
		Where(op.Equal(node.Address().PersonID(), o.PrimaryKey())).
		Select(node.Address().PrimaryKeyNode()).
		Load(ctx)
	for _, obj := range objs {
		obj.Delete(ctx)
	}
	o.oAddresses = nil

	d.Delete(ctx, "person", "id", o.id)
	d.Commit(ctx, txid)
	broadcast.Delete(ctx, "goradd", "person", fmt.Sprintf("%v", o.id))
}

// deletePerson deletes the associated record from the database.
func deletePerson(ctx context.Context, pk string) {
	if obj := LoadPerson(ctx, pk, node.Person().PrimaryKeyNode()); obj != nil {
		obj.Delete(ctx)
	}
}

func (o *personBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.firstNameIsDirty = false
	o.lastNameIsDirty = false
	o.oProjectsAsManagerIsDirty = false
	o.oLoginIsDirty = false
	o.oEmployeeInfoIsDirty = false
	o.oAddressesIsDirty = false

}

func (o *personBase) IsDirty() bool {
	return o.idIsDirty ||
		o.firstNameIsDirty ||
		o.lastNameIsDirty ||
		o.oProjectsAsManagerIsDirty ||
		o.oLoginIsDirty ||
		o.oEmployeeInfoIsDirty ||
		o.oAddressesIsDirty
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *personBase) Get(key string) interface{} {

	switch key {
	case "ID":
		if !o.idIsValid {
			return nil
		}
		return o.id

	case "FirstName":
		if !o.firstNameIsValid {
			return nil
		}
		return o.firstName

	case "LastName":
		if !o.lastNameIsValid {
			return nil
		}
		return o.lastName

	case "ProjectsAsManager":
		return o.ProjectsAsManager()

	case "Login":
		return o.Login()

	case "EmployeeInfo":
		return o.EmployeeInfo()

	case "Addresses":
		return o.Addresses()

	case "ProjectAsTeamMember":
		return o.ProjectAsTeamMember()
	case "ProjectsAsTeamMember":
		return o.ProjectsAsTeamMember()

	case "PersonType":
		return o.PersonType()
	case "PersonTypes":
		return o.PersonTypes()

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database object over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *personBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err := encoder.Encode(o.id); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.idIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.idIsDirty); err != nil {
		return nil, err
	}

	if err := encoder.Encode(o.firstName); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.firstNameIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.firstNameIsDirty); err != nil {
		return nil, err
	}

	if err := encoder.Encode(o.lastName); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.lastNameIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.lastNameIsDirty); err != nil {
		return nil, err
	}

	if err := encoder.Encode(o.oProjectsAsManager); err != nil {
		return nil, err
	}

	if err := encoder.Encode(o.oLogin); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.oEmployeeInfo); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.oAddresses); err != nil {
		return nil, err
	}

	if err := encoder.Encode(o.oProjectsAsTeamMember); err != nil {
		return nil, err
	}

	if err := encoder.Encode(o.oPersonTypes); err != nil {
		return nil, err
	}

	if o._aliases == nil {
		if err := encoder.Encode(false); err != nil {
			return nil, err
		}
	} else {
		if err := encoder.Encode(true); err != nil {
			return nil, err
		}
		if err := encoder.Encode(o._aliases); err != nil {
			return nil, err
		}
	}

	if err := encoder.Encode(o._restored); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (o *personBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&o.id); err != nil {
		return
	}
	if err = dec.Decode(&o.idIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.firstName); err != nil {
		return
	}
	if err = dec.Decode(&o.firstNameIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.firstNameIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.lastName); err != nil {
		return
	}
	if err = dec.Decode(&o.lastNameIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.lastNameIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.oProjectsAsManager); err != nil {
		return
	}
	if len(o.oProjectsAsManager) > 0 {
		o.mProjectsAsManager = make(map[string]*Project)
		for _, p := range o.oProjectsAsManager {
			o.mProjectsAsManager[p.PrimaryKey()] = p
		}
	}
	if err = dec.Decode(&o.oLogin); err != nil {
		return
	}
	if err = dec.Decode(&o.oEmployeeInfo); err != nil {
		return
	}
	if err = dec.Decode(&o.oAddresses); err != nil {
		return
	}
	if len(o.oAddresses) > 0 {
		o.mAddresses = make(map[string]*Address)
		for _, p := range o.oAddresses {
			o.mAddresses[p.PrimaryKey()] = p
		}
	}

	if err = dec.Decode(&o.oProjectsAsTeamMember); err != nil {
		return
	}
	if len(o.oProjectsAsTeamMember) > 0 {
		o.mProjectsAsTeamMember = make(map[string]*Project)

		for _, p := range o.oProjectsAsTeamMember {
			o.mProjectsAsTeamMember[p.PrimaryKey()] = p
		}
	}
	if err = dec.Decode(&o.oPersonTypes); err != nil {
		return
	}

	var hasAliases bool
	if err = dec.Decode(&hasAliases); err != nil {
		return
	}
	if hasAliases {
		if err = dec.Decode(&o._aliases); err != nil {
			return
		}
	}

	if err = dec.Decode(&o._restored); err != nil {
		return
	}

	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object.
func (o *personBase) MarshalJSON() (data []byte, err error) {
	v := make(map[string]interface{})

	if o.idIsValid {
		v["id"] = o.id
	}

	if o.firstNameIsValid {
		v["firstName"] = o.firstName
	}

	if o.lastNameIsValid {
		v["lastName"] = o.lastName
	}

	if val := o.ProjectsAsManager(); val != nil {
		v["manager"] = val
	}

	if val := o.Login(); val != nil {
		v["person"] = val
	}

	if val := o.EmployeeInfo(); val != nil {
		v["person"] = val
	}

	if val := o.Addresses(); val != nil {
		v["person"] = val
	}

	if val := o.ProjectsAsTeamMember(); val != nil {
		v["projectsAsTeamMember"] = val
	}
	if val := o.PersonTypes(); val != nil {
		v["personTypes"] = val
	}

	return json.Marshal(v)
}
