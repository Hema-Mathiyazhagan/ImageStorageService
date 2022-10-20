package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"errors"
	"fmt"
	"net/http"
	"gopkg.in/resty.v0"
	"github.com/Hema-Mathiyazhagan/image-storage-service/lib"
)

func RegisterErrors() {
	resty.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {

		if (r.StatusCode() > 399) {
			return errors.New("Status code error: " + fmt.Sprint(r.StatusCode()))
		}

		return nil
	})
}

var _ = Describe("Backend", func() {
	BeforeEach(func(){
		RegisterErrors()
	})
	AfterEach(func(){
	})
	Context("create album", func(){
		It("Cannot create the album", func(){
			statusCode,err := lib.Add_Album("abc")
			Expect(err).To(Not(BeNil()), "Error")
			Expect(statusCode).To(Equal(http.StatusInternalServerError), "Status code") 
		})
		It("Cannot create the album", func(){
			statusCode,err := lib.Add_Album(" ")
			Expect(err).To(Not(BeNil()), "Error")
			Expect(statusCode).To(Equal(http.StatusInternalServerError), "Status code") 
		})
		It("Cannot create the album", func(){
			statusCode,err := lib.Add_Album("abcdef")
			Expect(err).To(Not(BeNil()), "Error")
			Expect(statusCode).To(Equal(http.StatusInternalServerError), "Status code") 
		})
		It("Cannot create the album", func(){
			statusCode,err := lib.Add_Album("abc123")
			Expect(err).To(Not(BeNil()), "Error")
			Expect(statusCode).To(Equal(http.StatusInternalServerError), "Status code") 
		})
	})

	Context("Delete Album", func(){
		It("Cannot delete the album", func(){
			statusCode,err := lib.Delete_Album("***")
			Expect(err).To(Not(BeNil()), "Error")
			Expect(statusCode).To(Equal(http.StatusBadRequest), "Status code") 
		})
		It("Cannot delete the album", func(){
			statusCode,err := lib.Delete_Album("000124")
			Expect(err).To(Not(BeNil()), "Error")
			Expect(statusCode).To(Equal(http.StatusBadRequest), "Status code") 
		})
		It("Cannot delete the album", func(){
			statusCode,err := lib.Delete_Album(" ")
			Expect(err).To(Not(BeNil()), "Error")
			Expect(statusCode).To(Equal(http.StatusBadRequest), "Status code") 
		})
		It("Cannot delete the album", func(){
			statusCode,err := lib.Delete_Album("abcdef")
			Expect(err).To(Not(BeNil()), "Error")
			Expect(statusCode).To(Equal(http.StatusBadRequest), "Status code") 
		})
	})

	Context("Delete Image from album", func(){
		It("Cannot delete the image", func(){
			statusCode,err := lib.Delete_image("abcdef","png.png")
			Expect(err).To(Not(BeNil()), "Error")
			Expect(statusCode).To(Equal(http.StatusBadRequest), "Status code") 
		})
		It("Cannot delete the image", func(){
			statusCode,err := lib.Delete_image("abcdef","123.png")
			Expect(err).To(Not(BeNil()), "Error")
			Expect(statusCode).To(Equal(http.StatusBadRequest), "Status code") 
		})
	})

	Context("Retrieve Image from album", func(){
		It("Cannot retrieve the image", func(){
			imglist,err := lib.List_images("abcdef")
			Expect(err).To(Not(BeNil()), "Error")
			Expect(imglist).To(HaveLen(0),"Image list length") 
		})
		It("Cannot get the album", func(){
			imglist,err := lib.List_images(" ")
			Expect(err).To(Not(BeNil()), "Error")
			Expect(imglist).To(HaveLen(0),"Image list length") 
		})
		It("Cannot find the album", func(){
			imglist,err := lib.List_images("123***")
			Expect(err).To(Not(BeNil()), "Error")
			Expect(imglist).To(HaveLen(0),"Image list length") 
		})
	})
})

