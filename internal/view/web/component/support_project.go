package component

import (
	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
	lucide "github.com/nodxdev/nodxgo-lucide"
)

func SupportProjectModal() ModalResult {

	thanksSection := nodx.Div(
		nodx.Class("space-y-1"),
		nodx.P(
			nodx.Text("🙏 Thank you for considering supporting PG Back Web!"),
			nodx.Text(" Every contribution, big or small, helps keep the project running and growing! 🚀"),
		),
		nodx.P(
			nodx.Text("PG Back Web"),
			BText(" is 100% free and open-source"),
			nodx.Text(", but maintaining and improving it takes time and resources."),
		),
		PText("You can support the project in two ways:"),
	)

	sponsorSection := nodx.Div(
		nodx.Class("space-y-1"),
		H3Text("1. 💛 Become a Sponsor & Get Featured!"),
		nodx.P(
			nodx.Text("By becoming a sponsor, "),
			BText("your personal/business name, photo/logo and a link will be featured inside every self-hosted instance of PG Back Web and in our GitHub repository"),
			nodx.Text(" giving you visibility while supporting open-source development."),
		),
		PText("Click below to read more, become a sponsor and be part of something great!"),

		nodx.Div(
			nodx.Class("flex justify-center pt-1"),
			nodx.A(
				alpine.XBind("href", "sponsorsLink"),
				nodx.Target("_blank"),
				nodx.Class("btn btn-success btn-lg flex items-center space-x-1 text-lg"),
				SpanText("Become a sponsor"),
				lucide.HeartHandshake(),
			),
		),
	)

	referralSection := nodx.Div(
		nodx.Class("space-y-1"),
		H3Text("2. 🔗 Use PG Back Web's Referral Links & Get Bonuses!"),
		nodx.P(
			nodx.Text("Want to support PG Back Web without spending a dime?"),
			nodx.Text(" Use PG Back Web's referral links to access "),
			BText("amazing deals, free credits, and special bonuses"),
			nodx.Text(" on top services we personally recommend."),
		),
		nodx.P(
			nodx.Text("You get exclusive perks, and a small commission helps support PG Back Web - "),
			BText("win-win!"),
		),

		nodx.Div(
			nodx.Class("pt-2 space-y-4"),
			alpine.Template(
				alpine.XFor("ref in referralLinks"),
				CardBoxSimpleBgBase200(
					nodx.Div(
						nodx.Class("space-y-2"),
						nodx.Div(
							nodx.Class("p-2 bg-white rounded-btn"),
							nodx.Img(
								nodx.Class("max-h-[50px] mx-auto"),
								alpine.XBind("alt", "ref.name"),
								alpine.XBind("src", "ref.logo"),
							),
						),
						H4(alpine.XText("ref.name")),
						nodx.P(alpine.XText("ref.description")),
						nodx.Div(
							nodx.Class("flex justify-end"),
							nodx.A(
								alpine.XBind("href", "ref.link"),
								nodx.Target("_blank"),
								nodx.Class("btn btn-success btn-sm"),
								nodx.SpanEl(alpine.XText("'Get ' + ref.name + ' deal'")),
								lucide.HeartHandshake(),
							),
						),
					),
				),
			),
		),
	)

	mo := Modal(ModalParams{
		Size:  SizeMd,
		Title: "Help support PG Back Web",
		Content: []nodx.Node{
			nodx.Div(
				alpine.XData("alpineSupportProjectData()"),
				alpine.Template(
					alpine.XIf("!isLoading"),
					nodx.Div(
						nodx.Class("space-y-6"),
						thanksSection,
						sponsorSection,
						referralSection,
					),
				),
			),
		},
	})

	return mo
}

func SupportProjectButton(size size) nodx.Node {
	mo := SupportProjectModal()

	return nodx.Group(
		mo.HTML,
		nodx.Button(
			mo.OpenerAttr,
			nodx.ClassMap{
				"btn btn-success": true,
				"btn-sm":          size == SizeSm,
				"btn-lg":          size == SizeLg,
			},
			lucide.HeartHandshake(),
			SpanText("Support the project"),
		),
	)
}

func SupportProjectAnchor(text string) nodx.Node {
	mo := SupportProjectModal()

	return nodx.Group(
		mo.HTML,
		nodx.A(
			mo.OpenerAttr,
			nodx.Class("link"),
			SpanText(text),
		),
	)
}
