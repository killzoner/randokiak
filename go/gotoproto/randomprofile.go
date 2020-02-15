package randomprofile

//this file is an extract from https://github.com/Pallinder/go-randomdata/blob/master/fullprofile.go
//with un-nested structure so it can be used to generate related proto file

type ProfileName struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Title string `json:"title"`
}
type ProfileLocation struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	Postcode int    `json:"postcode"`
}

type ProfileLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Md5      string `json:"md5"`
	Sha1     string `json:"sha1"`
	Sha256   string `json:"sha256"`
}

type ProfileId struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type ProfilePicture struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}

// Profile : alias to generate proto
type Profile struct {
	Gender     string          `json:"gender"`
	Name       ProfileName     `json:"name"`
	Location   ProfileLocation `json:"location"`
	Email      string          `json:"email"`
	Login      ProfileLogin    `json:"login"`
	Dob        string          `json:"dob"`
	Registered string          `json:"registered"`
	Phone      string          `json:"phone"`
	Cell       string          `json:"cell"`
	ID         ProfileId       `json:"id"`
	Picture    ProfilePicture  `json:"picture"`
	Nat        string          `json:"nat"`
}
