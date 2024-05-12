package response

import "github.com/Marvit-Solutions/csw-golang/library/struct/model"

type MentorHomeList struct {
	ID            int     `json:"id"`
	UUID          string  `json:"uuid"`
	Name          string  `json:"name"`
	TeachingField string  `json:"teaching_field"`
	Description   string  `json:"description"`
	Motto         string  `json:"motto"`
	MediaID       int     `json:"media_id"`
	Rating        float64 `json:"rating"`
}

type MentorHome struct {
	UUID          string               `json:"uuid"`
	Name          string               `json:"name"`
	TeachingField string               `json:"teaching_field"`
	Description   string               `json:"description"`
	Motto         string               `json:"motto"`
	Media         *model.MultiResImage `json:"media"`
	Rating        float64              `json:"rating"`
}

type MentorDetailHome struct {
	UUID           string   `json:"uuid"`
	Name           string   `json:"name"`
	TeachingField  string   `json:"teaching_field"`
	Description    string   `json:"description"`
	ProfilePicture string   `json:"profile_picture"`
	Uniques        []string `json:"uniques"`
}

type PlanHome struct {
	UUID       string  `json:"uuid"`
	ModuleName string  `json:"module_name"`
	Name       string  `json:"name"`
	Picture    string  `json:"picture"`
	Price      float64 `json:"price"`
	Group      bool    `json:"group"`
	Exercise   int     `json:"exercise"`
	Access     int     `json:"access"`
	Module     bool    `json:"module"`
	TryOut     int     `json:"try_out"`
	Zoom       bool    `json:"zoom"`
}

type TestimonialHome struct {
	UUID           string  `json:"uuid"`
	Name           string  `json:"name"`
	Class          string  `json:"class"`
	ProfilePicture string  `json:"profile_picture"`
	Comment        string  `json:"comment"`
	Rating         float64 `json:"rating"`
}

type MentorStats struct {
	Rating      float64
	Testimonial int
}
