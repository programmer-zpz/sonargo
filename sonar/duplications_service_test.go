package sonargo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/programmer-zpz/sonargo/sonar"
)

var _ = Describe("SonarCLI integration test", func() {
	Describe("Test Show in api/duplications", func() {
		It("Should be ok", func() {
			opt := &DuplicationsShowOption{
				Key:  "",
				Uuid: testComponentId,
			}
			v, resp, err := client.Duplications.Show(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
})
