package component

import (
	"strings"

	nodx "github.com/nodxdev/nodxgo"
	alpine "github.com/nodxdev/nodxgo-alpine"
)

func SupportProjectSponsors() nodx.Node {
	sponsorCard := func(tier string) nodx.Node {
		tier = strings.ToLower(tier)
		isGold := tier == "gold"
		isSilver := tier == "silver"
		isBronze := tier == "bronze"

		return nodx.A(
			nodx.Class("group tooltip mr-2 mb-2"),
			alpine.XBind("data-tip", "sponsor.name"),
			alpine.XBind("href", "sponsor.link"),
			nodx.Target("_blank"),

			nodx.Div(
				nodx.ClassMap{
					"rounded-full overflow-hidden": true,
					"border-4 border-transparent":  true,
					"hover:border-primary":         true,
				},
				nodx.Img(
					nodx.ClassMap{
						"size-[150px]": isGold,
						"size-[80px]":  isSilver,
						"size-[50px]":  isBronze,
						"object-cover": true,
					},
					alpine.XBind("src", "sponsor.logo"),
					alpine.XBind("alt", "sponsor.name"),
				),
			),
		)
	}

	return CardBoxSimple(
		H2Text("PG Back Web Sponsors"),
		nodx.P(
			nodx.Class("mt-1"),
			nodx.Text("A big thank you to the following sponsors for supporting PG Back Web!"),
			nodx.Text(" Your contributions help keep the project running and growing! 🚀"),
		),
		nodx.P(
			nodx.Class("mt-1"),
			nodx.Text("You can become a sponsor or contribute to the project"),
			nodx.Text(" even without sending money by using our referral links. "),
			SupportProjectAnchor(`Learn more here.`),
		),

		nodx.Div(
			alpine.XData("alpineSupportProjectData()"),
			nodx.Class("mt-4 space-y-6"),

			nodx.Div(
				H1Text("Gold Sponsors 🏆"),
				nodx.Div(
					nodx.Class("pt-2"),
					alpine.Template(
						alpine.XFor("sponsor in goldSponsors"),
						sponsorCard("gold"),
					),
				),
			),

			nodx.Div(
				H3Text("Silver Sponsors 🥈"),
				nodx.Div(
					nodx.Class("pt-2"),
					alpine.Template(
						alpine.XFor("sponsor in silverSponsors"),
						sponsorCard("silver"),
					),
				),
			),

			nodx.Div(
				H4Text("Bronze Sponsors 🥉"),
				nodx.Div(
					nodx.Class("pt-2"),
					alpine.Template(
						alpine.XFor("sponsor in bronzeSponsors"),
						sponsorCard("bronze"),
					),
				),
			),
		),
	)
}
