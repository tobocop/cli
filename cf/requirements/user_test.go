package requirements_test

import (
	"errors"

	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/requirements"

	fakeapi "github.com/cloudfoundry/cli/cf/api/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserRequirement", func() {
	var (
		userRepo        *fakeapi.FakeUserRepository
		userRequirement requirements.UserRequirement
	)

	BeforeEach(func() {
		userRepo = &fakeapi.FakeUserRepository{}
	})

	Describe("Execute", func() {
		Context("when wantGuid is true", func() {
			BeforeEach(func() {
				userRequirement = requirements.NewUserRequirement("the-username", userRepo, true)
			})

			It("tries to find the user in CC", func() {
				userRequirement.Execute()
				Expect(userRepo.FindByUsernameCallCount()).To(Equal(1))
				Expect(userRepo.FindByUsernameArgsForCall(0)).To(Equal("the-username"))
			})

			Context("when the call to find the user succeeds", func() {
				var user models.UserFields
				BeforeEach(func() {
					user = models.UserFields{Username: "the-username", Guid: "the-guid"}
					userRepo.FindByUsernameReturns(user, nil)
				})

				It("stores the user that was found", func() {
					userRequirement.Execute()
					Expect(userRequirement.GetUser()).To(Equal(user))
				})

				It("does not error", func() {
					err := userRequirement.Execute()
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when the call to find the user fails", func() {
				userError := errors.New("some error")
				BeforeEach(func() {
					userRepo.FindByUsernameReturns(models.UserFields{}, userError)
				})

				It("errors", func() {
					err := userRequirement.Execute()

					Expect(err).To(HaveOccurred())
					Expect(err).To(Equal(userError))
				})
			})
		})

		Context("when wantGuid is false", func() {
			BeforeEach(func() {
				userRequirement = requirements.NewUserRequirement("the-username", userRepo, false)
			})

			It("does not try to find the user in CC", func() {
				userRequirement.Execute()
				Expect(userRepo.FindByUsernameCallCount()).To(Equal(0))
			})

			It("stores a user with just Username set", func() {
				userRequirement.Execute()
				expectedUser := models.UserFields{Username: "the-username"}
				Expect(userRequirement.GetUser()).To(Equal(expectedUser))
			})
		})
	})
})
