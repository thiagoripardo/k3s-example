package model

type BaconResponse struct {
	InstanceID string `json:"instance_id"`
	Bacon      Bacon  `json:"bacon"`
}

type Bacon struct {
	BaconName   string `json:"bacon_name"`
	Description string `json:"description"`
}

var Bacons = []Bacon{
	{
		BaconName:   "Sizzle McSmoky",
		Description: "Sizzle McSmoky is the master of the grill, infusing every dish with the perfect balance of smoky flavor and crispy goodness.",
	},
	{
		BaconName:   "Crispy Crunchster",
		Description: "Crispy Crunchster is the crunchiest bacon in town, with a satisfying snap in every bite that will leave you craving for more.",
	},
	{
		BaconName:   "Maple Glaze Maven",
		Description: "Maple Glaze Maven is known for its irresistible sweetness, perfectly complementing the savory essence of bacon with a hint of maple perfection.",
	},
	{
		BaconName:   "Smoky Swirlz",
		Description: "Smoky Swirlz brings a unique twist to traditional bacon with its mesmerizing swirl pattern, making it as visually appealing as it is delicious.",
	},
	{
		BaconName:   "Baconator",
		Description: "The legendary Baconator is the great villain in the land of bacon and is the only one who was not generated by AI.",
	},
	{
		BaconName:   "Crispy Curls",
		Description: "Crispy Curls boasts perfectly curled edges that offer an ideal balance of crispy texture and mouthwatering tenderness in every bite.",
	},
	{
		BaconName:   "Golden Glider",
		Description: "Golden Glider is the epitome of perfection, with its golden-brown hue and delectable aroma that beckons you from afar.",
	},
	{
		BaconName:   "Savory Stripsmith",
		Description: "Savory Stripsmith crafts bacon with meticulous care, ensuring each strip is bursting with savory flavor and irresistible juiciness.",
	},
	{
		BaconName:   "Peppered Perfectionist",
		Description: "Peppered Perfectionist takes bacon to the next level with its expertly seasoned pepper coating, adding a delightful kick to every bite.",
	},
	{
		BaconName:   "Candied Crispster",
		Description: "Candied Crispster is a sweet sensation, with its caramelized exterior and heavenly crunch that will satisfy even the most discerning sweet tooth.",
	},
	{
		BaconName:   "Smokey Spiralizer",
		Description: "Smokey Spiralizer is a culinary work of art, featuring mesmerizing spiral patterns that enhance both the visual appeal and flavor profile of this bacon masterpiece.",
	},
	{
		BaconName:   "Crispy Crown",
		Description: "Crispy Crown reigns supreme in the kingdom of bacon, with its perfectly crisped edges and rich, succulent flavor that captivates all who taste it.",
	},
	{
		BaconName:   "Maple Majesty",
		Description: "Maple Majesty exudes elegance and refinement, with its delicate maple infusion adding a touch of sophistication to the classic bacon experience.",
	},
	{
		BaconName:   "Sizzling Serenade",
		Description: "Sizzling Serenade tantalizes the senses with its harmonious symphony of sizzling sounds and mouthwatering aroma, promising a culinary experience like no other.",
	},
	{
		BaconName:   "Hickory Hero",
		Description: "Hickory Hero is a beacon of bold flavor, with its robust hickory smoke essence elevating bacon to legendary status among meat enthusiasts.",
	},
	{
		BaconName:   "Crispy Comet",
		Description: "Crispy Comet streaks across the culinary sky with its blazing trails of crispy perfection, leaving a lasting impression on all who dare to taste its heavenly flavor.",
	},
	{
		BaconName:   "Savory Streaker",
		Description: "Savory Streaker takes bacon to new heights with its streaks of savory goodness, ensuring each bite is packed with irresistible flavor and juiciness.",
	},
	{
		BaconName:   "Peppery Picasso",
		Description: "Peppery Picasso is a masterpiece of flavor, with its expertly applied pepper coating creating a culinary canvas that delights the senses and ignites the taste buds.",
	},
	{
		BaconName:   "Caramelized Crusader",
		Description: "Caramelized Crusader is a sweet and savory sensation, with its caramelized exterior and tender interior combining to create a culinary crusade of epic proportions.",
	},
	{
		BaconName:   "Smoldering Sizzler",
		Description: "Smoldering Sizzler is a force to be reckoned with, with its intense smoky aroma and sizzling hot allure drawing you in for a fiery bacon experience.",
	},
	{
		BaconName:   "Crispy Cascade",
		Description: "Crispy Cascade is a waterfall of flavor, with its cascading layers of crispy perfection promising a culinary journey that flows with every savory bite.",
	},
	{
		BaconName:   "Maple Marauder",
		Description: "Maple Marauder strikes fear into the hearts of bland breakfasts, with its bold maple infusion and savory bacon goodness conquering taste buds with each delicious bite.",
	},
	{
		BaconName:   "Sizzling Symphony",
		Description: "Sizzling Symphony is a culinary masterpiece, with its harmonious blend of sizzling sounds and savory aroma creating a symphony of flavor that delights the senses.",
	},
	{
		BaconName:   "Hickory Haze",
		Description: "Hickory Haze envelops your senses in a cloud of smoky perfection, with its rich hickory essence infusing every bite with an intoxicating aroma and flavor.",
	},
	{
		BaconName:   "Crispy Crescendo",
		Description: "Crispy Crescendo builds to a crescendo of flavor, with each bite crescendoing into a symphony of crispy goodness that leaves you wanting more.",
	},
	{
		BaconName:   "Savory Sentinel",
		Description: "Savory Sentinel stands guard over your taste buds, with its bold flavor and succulent juiciness ensuring that every bacon experience is one to remember.",
	},
	{
		BaconName:   "Peppered Prodigy",
		Description: "Peppered Prodigy is a culinary prodigy, with its expertly seasoned pepper coating adding a burst of flavor that elevates bacon to new heights of deliciousness.",
	},
	{
		BaconName:   "Caramel Conqueror",
		Description: "Caramel Conqueror conquers taste buds with its irresistible combination of sweet caramelization and savory bacon goodness, leaving a trail of satisfied palates in its wake.",
	},
	{
		BaconName:   "Smoky Sensation",
		Description: "Smoky Sensation is a feast for the senses, with its intense smoky aroma and flavor enveloping you in a cloud of bacon bliss with every tantalizing bite.",
	},
	{
		BaconName:   "Crispy Connoisseur",
		Description: "Crispy Connoisseur is a connoisseur of crispiness, with its perfectly cooked bacon offering a symphony of crunch that satisfies even the most discerning palate.",
	},
	{
		BaconName:   "Maple Mastermind",
		Description: "Maple Mastermind is a culinary genius, with its delicate maple infusion adding a touch of sweetness that enhances the savory perfection of bacon with every delectable bite.",
	},
	{
		BaconName:   "Sizzling Savant",
		Description: "Sizzling Savant is a master of the grill, with its sizzling hot allure and savory aroma enticing you to indulge in a bacon experience like no other.",
	},
	{
		BaconName:   "Hickory Heirloom",
		Description: "Hickory Heirloom is a timeless classic, with its rich hickory essence and savory flavor profile reminiscent of bacon's proud heritage and enduring popularity.",
	},
	{
		BaconName:   "Crispy Charmer",
		Description: "Crispy Charmer captivates the senses with its irresistible charm and crispy perfection, leaving a lasting impression that keeps you coming back for more.",
	},
	{
		BaconName:   "Peppery Pioneer",
		Description: "Peppery Pioneer pioneers a new frontier of flavor, with its bold pepper coating adding a spicy kick that revolutionizes the bacon experience with every adventurous bite.",
	},
	{
		BaconName:   "Caramelized Creator",
		Description: "Caramelized Creator is the architect of bacon perfection, with its expertly caramelized exterior and tender interior crafting a culinary masterpiece that delights the senses.",
	},
	{
		BaconName:   "Smoky Sovereign",
		Description: "Smoky Sovereign rules over the realm of bacon with an iron fist of flavor, commanding attention with its intense smoky aroma and savory goodness that reigns supreme.",
	},
	{
		BaconName:   "Crispy Crusader",
		Description: "Crispy Crusader is a bacon warrior on a quest for perfection, with its crispy exterior and succulent interior leading the charge against bland breakfasts everywhere.",
	},
	{
		BaconName:   "Maple Maverick",
		Description: "Maple Maverick blazes a trail of maple-infused glory, with its bold flavor and irresistible sweetness setting it apart as a true maverick in the world of bacon.",
	},
	{
		BaconName:   "Sizzling Sensation",
		Description: "Sizzling Sensation sets your taste buds ablaze with its fiery sizzle and tantalizing aroma, promising a bacon experience that's as hot as it is delicious.",
	},
	{
		BaconName:   "Hickory Highlander",
		Description: "Hickory Highlander is a bacon champion of epic proportions, with its robust hickory flavor and savory goodness transporting you to the misty hills of flavor-filled satisfaction.",
	},
	{
		BaconName:   "Crispy Captain",
		Description: "Crispy Captain steers your taste buds on a course for culinary bliss, with its perfectly crisped edges and savory flavor profile guiding you to bacon nirvana.",
	},
	{
		BaconName:   "Peppery Picasso",
		Description: "Peppery Picasso paints a masterpiece of flavor, with its expertly seasoned pepper coating adding depth and complexity to the savory perfection of bacon with every delicious bite.",
	},
	{
		BaconName:   "Caramelized Connoisseur",
		Description: "Caramelized Connoisseur is a connoisseur of caramelized goodness, with its expertly crafted exterior and tender interior offering a culinary experience that's truly one of a kind.",
	},
	{
		BaconName:   "Smoky Showstopper",
		Description: "Smoky Showstopper steals the spotlight with its intense smoky flavor and savory goodness, leaving all who taste it in awe of its culinary brilliance.",
	},
	{
		BaconName:   "Crispy Creator",
		Description: "Crispy Creator is the architect of bacon perfection, with its expertly cooked strips offering a symphony of crunch and flavor that's simply irresistible.",
	},
	{
		BaconName:   "Maple Maestro",
		Description: "Maple Maestro conducts a sweet symphony of flavor, with its delicate maple infusion adding a touch of sweetness that harmonizes perfectly with the savory essence of bacon.",
	},
	{
		BaconName:   "Sizzling Samurai",
		Description: "Sizzling Samurai wields a mighty blade of flavor, with its sizzling hot allure and savory aroma cutting through the mundane to deliver a bacon experience that's truly legendary.",
	},
	{
		BaconName:   "Hickory Houdini",
		Description: "Hickory Houdini magically transforms ordinary bacon into a savory masterpiece, with its rich hickory flavor and smoky essence captivating your taste buds with every delicious bite.",
	},
}
