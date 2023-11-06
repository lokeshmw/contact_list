package contact

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



