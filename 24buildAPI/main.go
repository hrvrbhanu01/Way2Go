package main

// model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB
var courses []Course

//middleware or helper - file
func(c *Course) IsEmpty() bool{
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}
