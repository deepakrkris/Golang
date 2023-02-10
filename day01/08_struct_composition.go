package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
                     u.Name, u.Location)
}

type Player struct {
	u User
	GameId int
}

func GreetingsForPlayer(p Player) string{
  return p.u.Greetings(); //modify the statement to return the required string
}

func main() {
	p := Player {
		User{Id : 11, Name : "Deepak", Location : "India"},
		223,
	}
	fmt.Println(GreetingsForPlayer(p))
}
