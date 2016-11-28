package routes

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/sessionauth"
	"github.com/olegkravchuk/life_of_student/models"
	"github.com/olegkravchuk/life_of_student/services"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type PostsComment struct {
	sync.Mutex
	rooms []*Room
}

type Room struct {
	sync.Mutex
	name    string
	clients []*Client
}

type Client struct {
	Name       string
	in         <-chan *models.Message
	out        chan<- *models.Message
	done       <-chan bool
	err        <-chan error
	disconnect chan<- int
}

//Get the room for the given name
func (c *PostsComment) getRoom(name string) *Room {
	c.Lock()

	for _, room := range c.rooms {
		if room.name == name {
			return room
		}
	}

	r := &Room{sync.Mutex{}, name, make([]*Client, 0)}
	c.rooms = append(c.rooms, r)

	defer c.Unlock()

	return r
}

// Add a client to a room
func (r *Room) appendClient(client *Client) {
	r.Lock()
	r.clients = append(r.clients, client)
	r.Unlock()
}

// Remove a client from a room
func (r *Room) removeClient(client *Client) {
	r.Lock()

	for index, c := range r.clients {
		if c == client {
			r.clients = append(r.clients[:index], r.clients[(index+1):]...)
		}
	}
	defer r.Unlock()
}

// Message all the other clients in the same room
func (r *Room) messageOtherClients(client *Client, msg *models.Message) {
	r.Lock()
	for _, c := range r.clients {
		if c != client {
			c.out <- msg
		}
	}
	defer r.Unlock()
}

func NewPostsComment() *PostsComment {
	return &PostsComment{sync.Mutex{}, make([]*Room, 0)}
}

func CreateCommentHandler(postCom *PostsComment, req *http.Request, w http.ResponseWriter, params martini.Params, receiver <-chan *models.Message, sender chan<- *models.Message, done <-chan bool, disconnect chan<- int, errorChannel <-chan error, user sessionauth.User) {
	client := &Client{user.(*models.MyUser).Username, receiver, sender, done, errorChannel, disconnect}
	r := postCom.getRoom(params["id"])
	r.appendClient(client)

	id, _ := params["id"]
	for {
		select {
		case <-client.err:
		//disconnect <- 1001
		case msg := <-client.in:
			postIdInt, _ := strconv.Atoi(id)
			servicePost := services.PostService{}
			post, err := servicePost.GetPost(bson.M{"id": postIdInt})
			if err != nil {
				http.NotFound(w, req)
			}
			serviceComment := services.CommentService{}
			serviceComment.CreateComment(models.Comment{Post: *post, Author: *user.(*models.MyUser), Comment: msg.Comment, CreateDate: time.Now()})
			msg.Author = user.(*models.MyUser).Username
			msg.Date = time.Now().Format("2006-01-02")
			sender <- msg
			r.messageOtherClients(client, msg)
		case <-client.done:
			r.removeClient(client)
			return
		}
	}
}
