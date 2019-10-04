package readonly

import (
	. "code.cloudfoundry.org/cli/cf/util/testhelpers/matchers"
	"code.cloudfoundry.org/cli/integration/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("packages command", func() {
	var (
		orgName   string
		spaceName string
		appName   string
	)

	BeforeEach(func() {
		orgName = helpers.NewOrgName()
		spaceName = helpers.NewSpaceName()
		appName = helpers.NewAppName()
	})

	Describe("help", func() {
		When("--help flag is set", func() {
			It("appears in cf help -a", func() {
				session := helpers.CF("help", "-a")
				Eventually(session).Should(Exit(0))
				Expect(session).To(HaveCommandInCategoryWithDescription("packages", "APPS", "List packages of an app"))
			})

			It("Displays command usage to output", func() {
				session := helpers.CF("packages", "--help")

				Eventually(session).Should(Say("NAME:"))
				Eventually(session).Should(Say("packages - List packages of an app"))
				Eventually(session).Should(Say("USAGE:"))
				Eventually(session).Should(Say("cf packages APP_NAME"))
				Eventually(session).Should(Say("SEE ALSO:"))
				Eventually(session).Should(Say("app, create-package, droplets, push"))

				Eventually(session).Should(Exit(0))
			})
		})
	})

	When("the app name is not provided", func() {
		It("tells the user that the app name is required, prints help text, and exits 1", func() {
			session := helpers.CF("packages")

			Eventually(session.Err).Should(Say("Incorrect Usage: the required argument `APP_NAME` was not provided"))
			Eventually(session).Should(Say("NAME:"))
			Eventually(session).Should(Exit(1))
		})
	})

	When("the environment is not setup correctly", func() {
		It("fails with the appropriate errors", func() {
			helpers.CheckEnvironmentTargetedCorrectly(true, true, ReadOnlyOrg, "packages", appName)
		})
	})

	When("the environment is set up correctly", func() {
		var userName string

		BeforeEach(func() {
			helpers.SetupCF(orgName, spaceName)
			userName, _ = helpers.GetCredentials()
		})

		AfterEach(func() {
			helpers.QuickDeleteOrg(orgName)
		})

		When("the app does not exist", func() {
			It("displays app not found and exits 1", func() {
				session := helpers.CF("packages", appName)
				userName, _ = helpers.GetCredentials()

				Eventually(session).Should(Say(`Getting packages of app %s in org %s / space %s as %s\.\.\.`, appName, orgName, spaceName, userName))
				Eventually(session.Err).Should(Say(`App '%s' not found\.`, appName))
				Eventually(session).Should(Say("FAILED"))

				Eventually(session).Should(Exit(1))
			})
		})

		When("the app exists", func() {
			Context("with no packages", func() {
				BeforeEach(func() {
					Eventually(helpers.CF("create-app", appName)).Should(Exit(0))
				})

				It("displays empty list", func() {
					session := helpers.CF("packages", appName)
					Eventually(session).Should(Say(`Getting packages of app %s in org %s / space %s as %s\.\.\.`, appName, orgName, spaceName, userName))
					Eventually(session).Should(Say(`No packages found\.`))
					Eventually(session).Should(Exit(0))
				})
			})

			Context("with existing packages", func() {
				BeforeEach(func() {
					helpers.WithHelloWorldApp(func(dir string) {
						Eventually(helpers.CustomCF(helpers.CFEnv{WorkingDirectory: dir}, "v3-push", appName, "--no-start")).Should(Exit(0))
					})
				})

				It("displays packages in the list", func() {
					session := helpers.CF("packages", appName)
					Eventually(session).Should(Say(`Getting packages of app %s in org %s / space %s as %s\.\.\.`, appName, orgName, spaceName, userName))
					Eventually(session).Should(Say(`guid\s+state\s+created`))
					Eventually(session).Should(Say(`.*\s+ready\s+.*`))

					Eventually(session).Should(Exit(0))
				})
			})
		})
	})
})
