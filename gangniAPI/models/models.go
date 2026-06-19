package models

type BlogAuthor struct {
	Avatar string `json:"avatar"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
}
type BlogCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
type BlogPostListItem struct {
	Author                 BlogAuthor   `json:"author"`
	BlogPicturePathLinux   string       `json:"blogPicturePathLinux"`
	BlogPicturePathWindows string       `json:"blogPicturePathWindows"`
	Category               BlogCategory `json:"category"`
	Excerpt                string       `json:"excerpt"`
	Id                     int          `json:"id"`
	PublishedAt            string       `json:"publishedAt"`
	Slug                   string       `json:"slug"`
	Title                  string       `json:"title"`
}
type BlogPostResponse struct {
	Data BlogPostDetail `json:"data"`
}
type BlogPostsResponse struct {
	Data []BlogPostListItem `json:"data"`
	Meta PaginationMeta     `json:"meta"`
}

type BlogPostDetail struct {
	Author                 BlogAuthor   `json:"author"`
	BlogPicturePathLinux   string       `json:"blogPicturePathLinux"`
	BlogPicturePathWindows string       `json:"blogPicturePathWindows"`
	Category               BlogCategory `json:"category"`
	Content                string       `json:"content"`
	Id                     int          `json:"id"`
	PublishedAt            string       `json:"publishedAt"`
	Seo                    BlogSEO      `json:"seo"`
	Slug                   string       `json:"slug"`
	Tags                   []BlogTag    `json:"tags"`
	Title                  string       `json:"title"`
}

type BlogSEO struct {
	MetaDescription string `json:"metaDescription"`
	MetaTitle       string `json:"metaTitle"`
}

type BlogTag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Comment struct {
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ParentId  int    `json:"parentId"`
	Replies   string `json:"replies"`
}

type CommentCreateRequest struct {
	Content  string `json:"content"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	ParentId int    `json:"parentId"`
}

type CommentResponse struct {
	Comments []Comment `json:"comments"`
	Total    int       `json:"total"`
}

type ContactMail struct {
	Content string `json:"content"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
}

type Loader struct {
	AutoWeight             int    `json:"autoWeight"`
	BatteryType            string `json:"batteryType"`
	BrakeType              string `json:"brakeType"`
	ChargingTime           string `json:"chargingTime"`
	EngineType             string `json:"engineType"`
	ForkLength             int    `json:"forkLength"`
	FrontWheels            string `json:"frontWheels"`
	Height                 int    `json:"height"`
	HydraulicLiftingEngine string `json:"hydraulicLiftingEngine"`
	Id                     int    `json:"id"`
	Length                 int    `json:"length"`
	LiftHeight             int    `json:"liftHeight"`
	LiftingAngle           int    `json:"liftingAngle"`
	LiftingCylinder        string `json:"liftingCylinder"`
	LongmenFrameMaterial   int    `json:"longmenFrameMaterial"`
	MaxLiftWeight          int    `json:"maxLiftWeight"`
	Name                   string `json:"name"`
	PicturePathLinux       string `json:"picturePathLinux"`
	PicturePathWindows     string `json:"picturePathWindows"`
	Price                  int    `json:"price"`
	RearWheels             string `json:"rearWheels"`
	SteeringMode           string `json:"steeringMode"`
	TurningRadius          int    `json:"turningRadius"`
	Voltage                int    `json:"voltage"`
	WheelAxis              string `json:"wheelAxis"`
	Width                  int    `json:"width"`
	WorkingHours           string `json:"workingHours"`
}
type ManualLoader struct {
	BrakeType          string  `json:"brakeType"`
	Control            string  `json:"control"`
	DriveGear          string  `json:"driveGear"`
	ForkLength         int     `json:"forkLength"`
	ForkWidth          int     `json:"forkWidth"`
	Id                 int     `json:"id"`
	Length             int     `json:"length"`
	LiftingSpeed       int     `json:"liftingSpeed"`
	MaxLiftWeight      int     `json:"maxLiftWeight"`
	MaxSpeed           float64 `json:"maxSpeed"`
	Name               string  `json:"name"`
	PicturePathLinux   string  `json:"picturePathLinux"`
	PicturePathWindows string  `json:"picturePathWindows"`
	Price              int     `json:"price"`
}

type PaginationMeta struct {
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
}
