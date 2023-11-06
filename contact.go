package main

import (
	"fmt"
)

type contact struct {
	contact_name string
	number       int
	email_id     string
}

type contact_list interface {
	add_conatct()
	fav_contacts()
	search_contact()
	list_contacts()
	delete_contact()
	recently_added_contact()
}

type All_contacts struct {
	contacts           []contact
	favourite_contacts []contact
}

func (a *All_contacts) add_contact(x contact) {
	a.contacts = append(a.contacts, x)
}
func (a *All_contacts) list_contacts(){
		for i, contact := range a.contacts {
			fmt.Printf("%d. Name: %s, Phone: %d\n", i+1, contact.contact_name, contact.number)
		}
		
	}

func (a *All_contacts) fav_contact(y contact) {
	a.favourite_contacts = append(a.favourite_contacts, y)
}

func (a *All_contacts)search_contact(z string) {
	found := false
	for _, i := range a.contacts {
		if i.contact_name == z {
			found = true
			break
		}
	}

	if found {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
func (a *All_contacts)delete_contact(w string){
	for i, j:=range a.contacts{
		if j.contact_name == w{
			a.contacts = append(a.contacts[:i], a.contacts[i+1:]...)
		}

	} 

}

func(a *All_contacts)recently_added_contact()contact{
	b:= len(a.contacts)
	return a.contacts[b-1]
}


func main() {
	c := All_contacts{}
	lok := contact{
		contact_name: "lokesh",
		number:       8709782727,
		email_id:     "loki@gamil.com",
	}
	gub := contact{
		contact_name: "h",
		number:       8727,
		email_id:     "gub@gamil.com",
	}
	c.add_contact(lok)
	c.add_contact(gub)
	fmt.Println(c.contacts)
	c.fav_contact(gub)
	fmt.Println(c.favourite_contacts)
	c.search_contact("h")
	c.list_contacts()
	c.delete_contact("h")
	c.list_contacts()
	fmt.Println(c.recently_added_contact())

}
