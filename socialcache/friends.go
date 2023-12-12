package socialcache

import (
	"errors"
	. "github.com/dotabook/faceit-go-steam/protocol/steamlang"
	"github.com/dotabook/faceit-go-steam/steamid"
	"sync"
)

// FriendsList is a thread safe map
// They can be iterated over like so:
//
//	for id, friend := range client.Social.Friends.GetCopy() {
//		log.Println(id, friend.Name)
//	}
type FriendsList struct {
	mutex sync.RWMutex
	byId  map[steamid.SteamId]*Friend
}

// NewFriendsList builds a new friends list
func NewFriendsList() *FriendsList {
	return &FriendsList{byId: make(map[steamid.SteamId]*Friend)}
}

// Add adds a friend to the friend list
func (list *FriendsList) Add(friend Friend) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	_, exists := list.byId[steamid.SteamId(friend.SteamId)]
	if !exists { //make sure this doesnt already exist
		list.byId[steamid.SteamId(friend.SteamId)] = &friend
	}
}

// Remove removes a friend from the friend list
func (list *FriendsList) Remove(id steamid.SteamId) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	delete(list.byId, id)
}

// Returns a copy of the friends map
func (list *FriendsList) GetCopy() map[steamid.SteamId]Friend {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	flist := make(map[steamid.SteamId]Friend)
	for key, friend := range list.byId {
		flist[key] = *friend
	}
	return flist
}

// Returns a copy of the friend of a given SteamId
func (list *FriendsList) ById(id steamid.SteamId) (Friend, error) {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	if val, ok := list.byId[id]; ok {
		return *val, nil
	}
	return Friend{}, errors.New("Friend not found")
}

// Returns the number of friends
func (list *FriendsList) Count() int {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	return len(list.byId)
}

// Setter methods
func (list *FriendsList) SetName(id steamid.SteamId, name string) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if val, ok := list.byId[id]; ok {
		val.Name = name
	}
}

func (list *FriendsList) SetAvatar(id steamid.SteamId, hash string) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if val, ok := list.byId[id]; ok {
		val.Avatar = hash
	}
}

func (list *FriendsList) SetRelationship(id steamid.SteamId, relationship EFriendRelationship) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if val, ok := list.byId[id]; ok {
		val.Relationship = relationship
	}
}

func (list *FriendsList) SetPersonaState(id steamid.SteamId, state EPersonaState) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if val, ok := list.byId[id]; ok {
		val.PersonaState = state
	}
}

func (list *FriendsList) SetPersonaStateFlags(id steamid.SteamId, flags EPersonaStateFlag) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if val, ok := list.byId[id]; ok {
		val.PersonaStateFlags = flags
	}
}

func (list *FriendsList) SetGameAppId(id steamid.SteamId, gameappid uint32) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if val, ok := list.byId[id]; ok {
		val.GameAppId = gameappid
	}
}

func (list *FriendsList) SetGameId(id steamid.SteamId, gameid uint64) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if val, ok := list.byId[id]; ok {
		val.GameId = gameid
	}
}

func (list *FriendsList) SetGameName(id steamid.SteamId, name string) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if val, ok := list.byId[id]; ok {
		val.GameName = name
	}
}

// A Friend
type Friend struct {
	SteamId           steamid.SteamId `json:",string"`
	Name              string
	Avatar            string
	Relationship      EFriendRelationship
	PersonaState      EPersonaState
	PersonaStateFlags EPersonaStateFlag
	GameAppId         uint32
	GameId            uint64 `json:",string"`
	GameName          string
}
