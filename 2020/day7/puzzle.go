package main

var input = `striped tan bags contain 2 light silver bags, 1 drab black bag, 2 clear tan bags, 2 mirrored tan bags.
dark black bags contain 1 vibrant indigo bag, 5 muted gold bags, 4 bright tomato bags, 3 dull tan bags.
dim silver bags contain 1 vibrant black bag, 3 muted cyan bags, 4 plaid turquoise bags, 4 faded orange bags.
faded maroon bags contain 3 pale aqua bags.
dim crimson bags contain 1 faded orange bag, 2 plaid silver bags.
plaid indigo bags contain 1 dull purple bag, 2 plaid cyan bags, 4 mirrored gold bags, 3 striped magenta bags.
mirrored brown bags contain 4 posh bronze bags, 3 dark brown bags, 5 shiny violet bags.
faded gray bags contain 3 drab beige bags, 5 vibrant olive bags, 3 muted cyan bags.
bright salmon bags contain 2 striped white bags, 5 bright plum bags.
plaid brown bags contain 1 mirrored violet bag, 1 dark turquoise bag, 5 dark gold bags.
vibrant lime bags contain 3 posh yellow bags, 5 dim aqua bags, 1 plaid coral bag.
pale gold bags contain 5 vibrant cyan bags, 2 faded red bags, 2 mirrored purple bags.
dotted orange bags contain 5 faded fuchsia bags, 1 light indigo bag.
mirrored chartreuse bags contain 3 bright black bags.
drab blue bags contain 1 plaid turquoise bag, 3 dim red bags, 5 dim lime bags, 3 dark maroon bags.
dark violet bags contain 5 plaid lavender bags, 1 dark salmon bag, 1 muted yellow bag, 3 drab beige bags.
pale salmon bags contain 3 dark red bags, 4 clear violet bags, 4 dark blue bags, 2 pale olive bags.
wavy fuchsia bags contain 3 striped plum bags, 2 light violet bags, 3 dotted yellow bags.
shiny gray bags contain 4 dull salmon bags, 2 dotted brown bags, 3 bright tomato bags.
dark plum bags contain 3 pale red bags, 3 striped gold bags, 3 posh bronze bags, 5 drab tomato bags.
mirrored green bags contain 5 faded red bags, 2 posh fuchsia bags, 2 dim green bags, 2 posh green bags.
drab brown bags contain 4 light violet bags, 2 mirrored salmon bags.
shiny magenta bags contain 4 clear turquoise bags, 2 posh silver bags.
clear olive bags contain 5 muted gold bags, 1 bright orange bag, 2 clear turquoise bags.
dim blue bags contain 2 pale red bags.
drab indigo bags contain 4 clear fuchsia bags, 3 plaid orange bags.
dull tomato bags contain 2 wavy purple bags, 2 faded brown bags, 4 plaid indigo bags, 3 faded silver bags.
light salmon bags contain 4 clear crimson bags, 3 clear gold bags.
light green bags contain 1 dull olive bag, 2 pale gold bags.
muted coral bags contain 4 pale indigo bags, 1 dark lime bag.
plaid chartreuse bags contain 4 dark gray bags, 5 clear tomato bags, 3 dotted brown bags, 1 mirrored olive bag.
dull plum bags contain 1 wavy gray bag, 4 wavy indigo bags, 4 dark black bags.
plaid crimson bags contain 3 mirrored blue bags, 2 dull turquoise bags, 1 dotted bronze bag, 4 clear coral bags.
light tomato bags contain 3 mirrored chartreuse bags, 4 faded bronze bags, 1 bright tomato bag, 2 muted bronze bags.
muted fuchsia bags contain 4 plaid turquoise bags, 5 light cyan bags.
wavy salmon bags contain 3 dim red bags, 2 muted crimson bags, 1 dotted purple bag.
dim turquoise bags contain 5 dull salmon bags, 4 pale crimson bags.
striped red bags contain 2 clear turquoise bags, 5 muted orange bags, 4 mirrored coral bags, 1 dim aqua bag.
light purple bags contain 1 clear chartreuse bag, 3 dull cyan bags, 5 light fuchsia bags.
pale purple bags contain 3 plaid fuchsia bags, 2 posh cyan bags.
mirrored violet bags contain 1 dark silver bag, 2 faded silver bags, 4 posh gray bags, 2 vibrant crimson bags.
mirrored olive bags contain 1 posh fuchsia bag, 5 dim turquoise bags, 5 dark red bags, 5 shiny magenta bags.
shiny aqua bags contain 3 vibrant olive bags.
bright crimson bags contain 4 muted cyan bags, 4 dim red bags, 1 dotted chartreuse bag.
light gray bags contain 5 plaid coral bags.
faded chartreuse bags contain 4 dark coral bags, 1 muted beige bag, 4 vibrant purple bags.
clear gold bags contain 2 plaid coral bags, 5 light chartreuse bags, 2 posh olive bags, 2 mirrored bronze bags.
faded tomato bags contain 2 light silver bags, 4 striped yellow bags, 1 light indigo bag.
wavy crimson bags contain 4 dotted beige bags, 5 light violet bags.
plaid aqua bags contain 5 dark aqua bags, 3 vibrant gold bags, 5 bright orange bags.
drab yellow bags contain 1 muted bronze bag, 4 muted blue bags.
dim yellow bags contain 3 wavy olive bags, 5 bright blue bags, 3 faded orange bags.
dotted teal bags contain 1 posh tomato bag, 4 drab green bags, 3 light plum bags, 1 light green bag.
plaid gold bags contain 2 faded beige bags, 5 clear teal bags, 2 pale coral bags, 3 dull gold bags.
faded indigo bags contain 1 muted yellow bag, 5 faded brown bags.
dull coral bags contain 1 dotted tan bag, 5 dull olive bags, 4 shiny red bags.
wavy tomato bags contain 3 dull purple bags.
pale plum bags contain 2 mirrored fuchsia bags, 4 muted olive bags, 2 bright purple bags.
bright red bags contain 2 muted turquoise bags, 2 dark crimson bags, 5 dotted purple bags, 1 muted beige bag.
vibrant maroon bags contain 4 bright indigo bags, 3 mirrored turquoise bags, 1 dull plum bag, 1 dark yellow bag.
dull gray bags contain 1 dark chartreuse bag, 1 light orange bag.
plaid silver bags contain 2 pale gold bags, 5 dotted purple bags.
shiny violet bags contain 3 pale crimson bags.
shiny fuchsia bags contain 1 mirrored silver bag, 4 vibrant olive bags.
drab bronze bags contain 5 wavy beige bags, 5 dim chartreuse bags.
dotted crimson bags contain 1 wavy black bag, 2 muted teal bags, 1 plaid cyan bag, 5 bright orange bags.
dim beige bags contain 5 clear orange bags, 1 dim violet bag.
posh red bags contain 4 faded indigo bags, 4 dim coral bags, 1 striped beige bag, 3 dark lime bags.
dim teal bags contain 2 bright red bags, 5 dull salmon bags.
drab magenta bags contain 3 dark crimson bags, 1 muted gold bag, 3 posh silver bags, 2 dull aqua bags.
posh aqua bags contain 5 clear crimson bags, 5 dim turquoise bags, 1 clear coral bag.
mirrored tan bags contain 5 shiny olive bags, 1 light violet bag, 4 muted gold bags.
bright chartreuse bags contain 2 plaid brown bags.
drab purple bags contain 5 wavy tan bags, 5 vibrant fuchsia bags, 3 pale brown bags.
faded coral bags contain 5 light tan bags, 2 clear crimson bags, 1 dark gray bag.
posh indigo bags contain 3 plaid silver bags.
plaid green bags contain 5 vibrant teal bags, 4 mirrored gold bags, 2 striped silver bags.
clear tan bags contain 5 striped magenta bags, 4 faded chartreuse bags, 1 dotted purple bag.
posh white bags contain 3 bright black bags.
dotted maroon bags contain 4 clear teal bags, 1 wavy fuchsia bag, 1 posh turquoise bag.
dark magenta bags contain 1 vibrant violet bag.
shiny indigo bags contain 4 dull lime bags, 1 vibrant turquoise bag, 1 light gray bag.
light fuchsia bags contain 4 faded silver bags, 1 mirrored indigo bag.
light plum bags contain 5 clear crimson bags, 1 posh bronze bag.
dark purple bags contain 4 clear indigo bags, 4 bright tomato bags, 1 faded olive bag.
posh plum bags contain 4 mirrored black bags, 5 plaid silver bags.
wavy bronze bags contain 4 faded magenta bags, 4 dotted chartreuse bags, 3 striped beige bags, 2 striped crimson bags.
clear teal bags contain 2 shiny olive bags.
plaid coral bags contain 4 drab tomato bags, 3 striped gold bags, 1 dotted lime bag.
bright violet bags contain 3 muted olive bags, 5 plaid orange bags.
vibrant white bags contain 5 light brown bags, 2 pale olive bags, 5 vibrant violet bags.
shiny blue bags contain 1 vibrant coral bag.
pale yellow bags contain 2 dark tan bags.
wavy cyan bags contain 2 faded brown bags, 3 dotted tomato bags, 2 drab cyan bags.
dim aqua bags contain 4 dim fuchsia bags, 1 dark red bag.
mirrored indigo bags contain 1 dotted purple bag, 1 dim violet bag, 5 dark silver bags, 2 mirrored blue bags.
faded red bags contain 3 drab gray bags, 2 shiny magenta bags, 4 faded silver bags, 2 dull tan bags.
clear silver bags contain 5 shiny gray bags, 4 light violet bags.
striped yellow bags contain 4 plaid turquoise bags, 2 vibrant green bags, 4 posh bronze bags, 2 dark crimson bags.
light orange bags contain 4 bright blue bags.
drab fuchsia bags contain 3 light lime bags, 1 dull maroon bag.
vibrant indigo bags contain 5 faded silver bags, 2 striped white bags, 5 shiny magenta bags.
pale black bags contain 1 dull aqua bag, 1 striped gold bag.
clear lavender bags contain 4 posh bronze bags.
bright tan bags contain 5 posh bronze bags, 2 faded coral bags, 2 bright tomato bags, 3 posh fuchsia bags.
vibrant blue bags contain 2 drab magenta bags.
drab red bags contain 5 bright purple bags, 1 mirrored fuchsia bag, 1 plaid salmon bag.
faded lime bags contain 3 dark salmon bags, 2 drab gray bags, 1 vibrant cyan bag.
faded teal bags contain 1 faded bronze bag, 3 bright blue bags, 4 dull salmon bags, 4 dull white bags.
dark yellow bags contain 5 posh chartreuse bags.
dull white bags contain 3 light green bags, 3 drab gray bags, 2 mirrored tan bags, 2 bright black bags.
clear yellow bags contain 4 clear beige bags, 1 faded blue bag.
light coral bags contain 5 dotted orange bags, 2 vibrant plum bags, 2 wavy purple bags.
faded orange bags contain 1 light violet bag, 2 faded olive bags, 5 posh silver bags.
bright olive bags contain 1 clear purple bag.
shiny beige bags contain 5 pale silver bags, 3 drab orange bags.
dark aqua bags contain 3 drab gray bags, 5 shiny plum bags, 5 dull purple bags.
pale green bags contain 4 clear olive bags, 5 drab crimson bags, 2 dim silver bags, 1 striped fuchsia bag.
muted violet bags contain 4 faded plum bags, 3 bright beige bags.
shiny brown bags contain 4 dull tomato bags.
light cyan bags contain 4 shiny bronze bags, 5 faded beige bags, 2 shiny plum bags, 4 mirrored purple bags.
plaid yellow bags contain 1 vibrant bronze bag, 5 light black bags, 1 plaid coral bag, 4 drab tan bags.
light chartreuse bags contain 3 drab magenta bags, 5 mirrored coral bags.
shiny orange bags contain 1 drab gray bag, 4 bright tomato bags.
mirrored blue bags contain 1 dark coral bag, 3 vibrant purple bags, 1 mirrored turquoise bag.
clear green bags contain 4 wavy orange bags, 5 muted black bags, 5 striped gray bags.
bright silver bags contain 2 plaid aqua bags, 1 bright blue bag.
clear blue bags contain 5 dotted tan bags.
posh silver bags contain no other bags.
light lavender bags contain 5 dark gray bags, 2 clear tomato bags, 5 muted cyan bags, 4 pale cyan bags.
dull salmon bags contain 4 dark silver bags, 5 mirrored coral bags, 4 bright red bags.
clear purple bags contain 5 wavy yellow bags, 2 striped orange bags.
dark lavender bags contain 3 clear purple bags, 4 vibrant cyan bags.
dotted white bags contain 4 light fuchsia bags.
mirrored cyan bags contain 4 dotted plum bags, 2 posh salmon bags, 3 light orange bags, 3 muted teal bags.
shiny cyan bags contain 1 dotted blue bag, 4 striped crimson bags.
clear tomato bags contain 5 vibrant green bags, 1 drab lavender bag, 1 dim coral bag.
pale olive bags contain 1 faded red bag, 5 dim lime bags.
wavy red bags contain 4 faded orange bags, 1 dull white bag, 1 dotted chartreuse bag.
pale violet bags contain 4 muted turquoise bags, 5 vibrant yellow bags, 2 bright plum bags, 4 posh plum bags.
dull lavender bags contain 4 faded turquoise bags, 5 striped gray bags.
drab gray bags contain 5 faded silver bags, 1 plaid turquoise bag.
vibrant aqua bags contain 2 drab brown bags.
shiny tan bags contain 2 dark red bags.
clear turquoise bags contain no other bags.
pale teal bags contain 4 dotted brown bags, 1 dull lime bag, 5 dotted purple bags.
posh teal bags contain 4 vibrant purple bags.
drab orange bags contain 1 striped magenta bag.
plaid salmon bags contain 3 light chartreuse bags, 3 dull purple bags.
mirrored white bags contain 2 light beige bags.
mirrored beige bags contain 3 dim tan bags, 5 bright salmon bags, 5 pale tan bags.
mirrored maroon bags contain 1 dotted salmon bag, 1 vibrant yellow bag, 4 bright cyan bags, 1 clear blue bag.
dim lime bags contain 5 light violet bags, 4 drab gray bags.
pale bronze bags contain 5 dotted fuchsia bags, 1 dull tan bag, 2 faded fuchsia bags.
vibrant turquoise bags contain 1 dotted tan bag, 2 clear purple bags, 5 dotted chartreuse bags.
vibrant crimson bags contain 1 faded orange bag, 3 light chartreuse bags, 3 dotted purple bags.
dotted green bags contain 2 vibrant lime bags, 1 dotted bronze bag, 3 faded chartreuse bags.
mirrored gray bags contain 4 shiny maroon bags, 2 faded salmon bags.
muted purple bags contain 3 clear magenta bags, 1 light green bag, 3 faded orange bags.
wavy maroon bags contain 4 muted beige bags, 4 posh gray bags.
clear aqua bags contain 3 clear fuchsia bags.
mirrored red bags contain 3 muted indigo bags, 4 pale lime bags, 1 dark silver bag, 3 mirrored tomato bags.
vibrant olive bags contain 2 shiny gold bags, 5 mirrored tan bags, 2 pale olive bags.
dull purple bags contain 5 dotted purple bags, 5 dotted fuchsia bags, 4 dark salmon bags, 3 faded olive bags.
muted maroon bags contain 1 posh green bag, 5 faded bronze bags, 4 striped gold bags.
shiny chartreuse bags contain 3 bright crimson bags, 4 dim purple bags, 2 dark violet bags.
pale crimson bags contain 4 dark silver bags, 3 dark crimson bags.
vibrant fuchsia bags contain 3 shiny gold bags, 1 wavy bronze bag, 4 wavy purple bags, 5 dark gray bags.
pale white bags contain 4 clear purple bags, 5 muted turquoise bags, 4 clear lavender bags.
faded silver bags contain 1 striped white bag.
dim magenta bags contain 5 plaid salmon bags, 3 dark chartreuse bags.
dark beige bags contain 3 bright brown bags, 4 dull turquoise bags, 3 muted black bags.
wavy beige bags contain 5 faded olive bags.
dark turquoise bags contain 4 striped indigo bags, 5 pale blue bags.
pale tan bags contain 3 dim olive bags, 4 posh green bags.
muted beige bags contain 2 pale indigo bags, 1 bright black bag, 3 bright plum bags.
faded fuchsia bags contain 4 faded silver bags.
shiny olive bags contain no other bags.
dark lime bags contain 3 clear olive bags, 2 clear blue bags, 2 dull cyan bags.
dim tan bags contain 4 plaid chartreuse bags.
light bronze bags contain 5 bright aqua bags, 1 plaid magenta bag, 5 dim lime bags.
muted chartreuse bags contain 5 faded cyan bags, 4 posh fuchsia bags, 2 drab tomato bags.
bright blue bags contain 5 vibrant green bags, 5 dull white bags, 1 bright orange bag.
light beige bags contain 3 dim magenta bags, 5 plaid tomato bags, 5 wavy black bags.
dull maroon bags contain 1 dotted gold bag.
bright lavender bags contain 3 drab blue bags, 2 dotted yellow bags.
drab teal bags contain 5 shiny crimson bags.
faded tan bags contain 2 muted cyan bags.
wavy lime bags contain 4 mirrored silver bags, 3 dark purple bags.
dull turquoise bags contain 4 posh green bags, 4 shiny gold bags, 5 clear blue bags.
clear salmon bags contain 4 striped blue bags, 2 faded chartreuse bags, 1 muted tan bag, 4 dark coral bags.
wavy magenta bags contain 3 dim gold bags, 5 shiny white bags, 4 striped gold bags.
striped crimson bags contain 3 pale aqua bags, 3 pale indigo bags, 4 pale tan bags.
light white bags contain 4 clear fuchsia bags, 1 dark silver bag, 5 shiny olive bags.
bright magenta bags contain 4 pale beige bags.
posh tan bags contain 5 plaid salmon bags.
pale tomato bags contain 1 posh salmon bag.
bright purple bags contain 4 mirrored coral bags.
dark orange bags contain 5 drab olive bags, 4 faded yellow bags.
light olive bags contain 2 striped brown bags, 1 wavy tan bag.
clear red bags contain 4 light chartreuse bags.
wavy coral bags contain 3 muted purple bags, 1 bright tomato bag.
dull teal bags contain 3 dull tan bags, 5 faded silver bags, 3 striped orange bags.
dark crimson bags contain no other bags.
muted yellow bags contain 1 dull teal bag, 3 drab olive bags.
vibrant orange bags contain 2 striped beige bags.
faded gold bags contain 4 wavy magenta bags, 4 mirrored teal bags.
clear indigo bags contain 1 dark gray bag, 2 shiny magenta bags, 1 vibrant olive bag.
plaid olive bags contain 5 striped lavender bags, 5 faded orange bags.
dotted salmon bags contain 1 pale tan bag, 5 light magenta bags.
dull silver bags contain 4 mirrored tan bags, 3 light tan bags, 2 dark black bags.
plaid fuchsia bags contain 2 mirrored bronze bags.
dotted beige bags contain 2 faded orange bags, 3 clear orange bags.
dark fuchsia bags contain 2 clear turquoise bags, 1 vibrant indigo bag, 5 wavy blue bags, 3 dotted brown bags.
muted gray bags contain 4 faded magenta bags, 5 mirrored salmon bags, 5 vibrant purple bags.
muted aqua bags contain 1 vibrant turquoise bag.
wavy indigo bags contain 1 dotted salmon bag, 4 light violet bags.
plaid black bags contain 3 faded teal bags, 5 wavy teal bags.
muted plum bags contain 4 vibrant teal bags, 4 pale aqua bags, 2 light bronze bags, 4 pale silver bags.
striped salmon bags contain 3 shiny turquoise bags, 5 drab tan bags, 3 clear silver bags.
mirrored black bags contain 1 dotted chartreuse bag, 4 dull olive bags, 2 dim coral bags.
striped black bags contain 1 vibrant black bag, 3 faded coral bags, 3 wavy gray bags.
faded beige bags contain 3 dim olive bags, 5 vibrant bronze bags, 4 mirrored gold bags.
striped fuchsia bags contain 5 dotted yellow bags.
light indigo bags contain 2 pale gold bags, 3 light black bags.
mirrored silver bags contain 4 dark lime bags, 3 clear orange bags, 3 clear lime bags.
vibrant black bags contain 3 posh silver bags, 2 dull aqua bags, 2 dark black bags.
vibrant tomato bags contain 5 striped beige bags, 3 posh coral bags, 4 posh brown bags.
clear chartreuse bags contain 2 dull bronze bags, 2 drab turquoise bags.
dotted coral bags contain 1 dull brown bag, 5 clear teal bags, 3 mirrored olive bags.
drab beige bags contain 5 dotted olive bags, 5 dotted chartreuse bags, 4 dark coral bags, 3 wavy yellow bags.
wavy plum bags contain 1 dotted blue bag, 5 vibrant brown bags.
posh brown bags contain 2 dark plum bags, 3 shiny white bags, 2 striped plum bags, 3 dim lavender bags.
pale lime bags contain 5 posh orange bags, 2 posh magenta bags.
plaid tomato bags contain 4 shiny gray bags, 2 faded olive bags.
shiny purple bags contain 5 mirrored tan bags, 3 faded red bags, 1 clear coral bag, 4 posh purple bags.
vibrant red bags contain 3 wavy tomato bags, 5 mirrored magenta bags.
dim lavender bags contain 4 dim red bags.
dull beige bags contain 2 shiny gold bags.
vibrant cyan bags contain 4 shiny magenta bags.
striped chartreuse bags contain 2 dark plum bags.
pale magenta bags contain 2 drab tomato bags, 2 striped chartreuse bags, 4 clear crimson bags, 5 plaid green bags.
wavy teal bags contain 1 muted gold bag.
plaid tan bags contain 2 plaid turquoise bags, 5 pale aqua bags.
striped lime bags contain 1 light yellow bag, 1 drab coral bag, 1 posh blue bag.
light turquoise bags contain 2 clear brown bags, 1 pale fuchsia bag, 1 vibrant teal bag, 3 dark salmon bags.
dim indigo bags contain 4 dark red bags, 3 drab violet bags, 5 bright white bags.
faded crimson bags contain 4 dotted brown bags.
vibrant gold bags contain 2 pale turquoise bags, 2 vibrant tan bags.
drab turquoise bags contain 3 wavy teal bags.
dark silver bags contain 3 clear turquoise bags, 2 posh green bags, 3 dim coral bags, 1 clear coral bag.
mirrored bronze bags contain 3 light black bags, 3 dim olive bags, 5 posh purple bags, 2 dim lime bags.
dim gold bags contain 2 dotted lavender bags, 1 dotted olive bag.
dull aqua bags contain no other bags.
shiny red bags contain 5 vibrant magenta bags, 2 dull silver bags.
dull orange bags contain 2 drab orange bags.
striped violet bags contain 4 mirrored tan bags, 5 clear crimson bags, 3 dotted lavender bags, 5 light plum bags.
faded blue bags contain 4 striped blue bags, 1 light silver bag, 5 drab black bags.
dotted aqua bags contain 5 plaid fuchsia bags, 1 bright plum bag, 4 shiny red bags.
shiny lavender bags contain 5 wavy indigo bags, 5 wavy beige bags.
faded salmon bags contain 4 shiny red bags.
muted orange bags contain 4 dim silver bags, 3 clear turquoise bags, 4 pale red bags, 2 pale olive bags.
drab silver bags contain 1 vibrant yellow bag, 3 pale tomato bags, 1 dull silver bag.
dim white bags contain 5 drab orange bags, 3 dark fuchsia bags, 4 muted coral bags, 2 shiny aqua bags.
muted lavender bags contain 3 faded lime bags, 5 light plum bags.
pale blue bags contain 4 dotted chartreuse bags.
clear bronze bags contain 4 dull magenta bags, 3 mirrored fuchsia bags, 2 shiny lime bags, 4 plaid violet bags.
clear magenta bags contain 4 wavy blue bags, 1 dotted lavender bag, 1 shiny olive bag.
dotted gold bags contain 3 shiny coral bags, 5 bright tan bags.
vibrant violet bags contain 4 bright olive bags.
posh fuchsia bags contain 1 striped crimson bag, 5 mirrored gold bags, 3 muted beige bags.
faded white bags contain 5 plaid purple bags, 3 posh lime bags, 3 clear blue bags.
dim plum bags contain 1 posh purple bag, 2 pale indigo bags, 1 dark lavender bag.
dull lime bags contain 2 striped magenta bags, 5 drab tomato bags.
plaid orange bags contain 5 faded tan bags, 1 posh aqua bag, 2 faded red bags, 2 clear red bags.
muted white bags contain 2 striped white bags.
bright gold bags contain 3 wavy beige bags, 2 faded coral bags, 5 drab chartreuse bags, 1 vibrant brown bag.
pale maroon bags contain 2 plaid salmon bags, 1 muted turquoise bag.
posh cyan bags contain 1 dotted brown bag, 3 shiny magenta bags, 2 pale brown bags.
striped tomato bags contain 5 vibrant lime bags, 1 dark black bag.
dotted magenta bags contain 2 wavy turquoise bags, 2 striped teal bags, 5 dotted yellow bags, 2 drab brown bags.
wavy lavender bags contain 5 light turquoise bags.
clear plum bags contain 4 faded tomato bags, 3 clear yellow bags, 5 drab silver bags, 5 dim yellow bags.
muted silver bags contain 5 pale maroon bags, 2 posh silver bags, 2 dotted yellow bags, 1 muted indigo bag.
muted indigo bags contain 5 dull beige bags, 4 striped orange bags, 2 wavy fuchsia bags, 3 muted magenta bags.
light teal bags contain 1 clear magenta bag, 4 faded olive bags, 5 dotted purple bags, 4 clear coral bags.
posh purple bags contain 2 bright tomato bags, 5 dark gray bags, 5 clear coral bags, 5 clear turquoise bags.
muted lime bags contain 1 mirrored bronze bag.
pale coral bags contain 1 vibrant lime bag, 4 vibrant black bags, 2 vibrant magenta bags.
mirrored orange bags contain 1 dim orange bag, 4 faded purple bags, 5 dull bronze bags, 5 dull green bags.
vibrant bronze bags contain 1 muted gold bag.
plaid white bags contain 4 dull yellow bags.
dark white bags contain 2 drab magenta bags, 5 dull beige bags.
dull olive bags contain 2 muted teal bags, 4 posh purple bags, 2 muted lime bags.
dull fuchsia bags contain 1 wavy olive bag, 4 bright aqua bags, 5 vibrant purple bags, 3 shiny purple bags.
dim purple bags contain 5 light brown bags.
mirrored purple bags contain 5 clear coral bags.
dim bronze bags contain 3 dull teal bags, 3 posh blue bags, 1 dull purple bag, 4 striped chartreuse bags.
light brown bags contain 1 plaid olive bag, 1 vibrant tan bag, 2 dotted orange bags.
faded turquoise bags contain 2 dotted fuchsia bags, 5 muted orange bags.
striped gold bags contain 3 drab olive bags, 1 vibrant green bag, 3 shiny gold bags.
light black bags contain 3 shiny magenta bags, 4 posh green bags.
dull indigo bags contain 2 muted bronze bags, 5 plaid magenta bags.
light red bags contain 4 dim teal bags, 3 muted yellow bags, 1 dotted gray bag.
dark coral bags contain 3 dotted chartreuse bags, 3 bright orange bags, 4 shiny magenta bags.
dark cyan bags contain 3 striped magenta bags.
plaid bronze bags contain 3 shiny blue bags.
muted cyan bags contain 1 posh green bag.
wavy orange bags contain 1 posh green bag, 1 vibrant cyan bag, 1 dark maroon bag, 1 dull teal bag.
shiny salmon bags contain 1 dotted yellow bag, 1 striped olive bag, 4 dull teal bags, 3 drab magenta bags.
pale turquoise bags contain 4 bright crimson bags, 5 dotted fuchsia bags, 1 faded turquoise bag, 5 muted turquoise bags.
dotted silver bags contain 5 dark orange bags, 4 dotted olive bags, 4 pale bronze bags, 3 dull yellow bags.
dotted red bags contain 4 dull white bags, 2 shiny plum bags, 2 pale gold bags.
striped blue bags contain 4 shiny orange bags, 1 striped orange bag.
dotted tan bags contain 3 dark brown bags, 3 wavy yellow bags, 5 dotted fuchsia bags.
shiny green bags contain 2 dark olive bags, 2 clear indigo bags, 2 striped salmon bags.
wavy tan bags contain 4 mirrored purple bags, 3 mirrored blue bags, 3 clear turquoise bags, 4 vibrant purple bags.
dark brown bags contain 2 shiny olive bags, 1 dull aqua bag.
plaid cyan bags contain 3 dull aqua bags, 1 muted teal bag.
dark tomato bags contain 2 faded lime bags, 2 dark cyan bags, 4 vibrant gray bags.
dark gold bags contain 3 pale teal bags, 3 posh chartreuse bags, 3 pale tan bags, 2 drab magenta bags.
striped indigo bags contain 4 striped black bags, 3 vibrant black bags, 2 mirrored turquoise bags, 3 dark coral bags.
faded lavender bags contain 1 muted turquoise bag, 3 dull bronze bags, 3 bright red bags, 2 dull turquoise bags.
wavy black bags contain 3 dotted purple bags.
faded plum bags contain 4 plaid lavender bags, 5 dark lavender bags, 5 dark black bags, 3 striped crimson bags.
posh olive bags contain 2 striped beige bags, 1 plaid crimson bag.
dim green bags contain 5 faded brown bags.
mirrored magenta bags contain 5 clear yellow bags.
pale red bags contain 2 dotted purple bags, 5 dotted tan bags, 2 dark brown bags.
bright black bags contain 1 dark maroon bag.
dark gray bags contain no other bags.
mirrored plum bags contain 5 bright orange bags, 3 muted purple bags.
bright turquoise bags contain 5 vibrant yellow bags.
posh beige bags contain 4 dim green bags, 3 posh gold bags.
bright maroon bags contain 2 muted blue bags, 5 shiny plum bags.
bright brown bags contain 2 muted black bags.
shiny yellow bags contain 5 muted olive bags, 5 wavy tomato bags, 1 mirrored green bag.
dim tomato bags contain 5 drab chartreuse bags.
posh turquoise bags contain 4 drab green bags, 2 dim gold bags.
posh blue bags contain 3 shiny bronze bags, 3 dull white bags, 3 clear beige bags, 4 dull cyan bags.
wavy green bags contain 5 pale bronze bags, 3 dull maroon bags, 4 dark purple bags.
posh violet bags contain 4 plaid fuchsia bags, 1 clear salmon bag.
bright beige bags contain 1 faded maroon bag, 2 dim brown bags, 1 dull chartreuse bag.
dim violet bags contain 5 vibrant bronze bags.
mirrored gold bags contain 5 light violet bags, 2 mirrored tan bags.
mirrored lavender bags contain 5 dotted white bags, 3 wavy chartreuse bags, 2 shiny lavender bags.
pale silver bags contain 4 mirrored turquoise bags, 3 faded coral bags, 5 wavy yellow bags, 5 shiny gold bags.
shiny lime bags contain 1 plaid indigo bag.
dim cyan bags contain 4 dim teal bags.
wavy violet bags contain 2 plaid beige bags.
striped brown bags contain 4 plaid silver bags, 4 striped violet bags, 2 muted gold bags, 1 plaid yellow bag.
pale chartreuse bags contain 3 dark coral bags, 2 faded olive bags.
mirrored fuchsia bags contain 1 plaid lavender bag, 5 dark coral bags, 5 light green bags, 1 dotted bronze bag.
dotted cyan bags contain 5 posh gold bags, 5 vibrant red bags.
shiny gold bags contain 2 dark brown bags, 2 mirrored coral bags, 1 faded olive bag, 1 posh bronze bag.
striped coral bags contain 3 clear red bags, 2 mirrored crimson bags.
striped beige bags contain 5 plaid salmon bags, 4 dim lime bags, 3 dark brown bags.
plaid lavender bags contain 2 dotted olive bags, 4 faded olive bags, 4 shiny magenta bags.
muted crimson bags contain 3 shiny purple bags.
mirrored crimson bags contain 2 muted teal bags, 3 dotted purple bags, 5 pale teal bags.
posh salmon bags contain 5 vibrant turquoise bags, 3 muted black bags, 3 pale indigo bags, 2 shiny plum bags.
vibrant coral bags contain 2 drab magenta bags, 2 faded red bags, 4 plaid magenta bags, 4 muted cyan bags.
shiny crimson bags contain 1 dotted beige bag, 2 mirrored turquoise bags, 5 clear gray bags, 1 dim lavender bag.
posh bronze bags contain 5 clear coral bags, 1 dotted lime bag, 5 faded olive bags, 3 dim red bags.
faded purple bags contain 1 plaid white bag, 4 posh lime bags, 5 pale turquoise bags.
posh magenta bags contain 2 dull turquoise bags, 2 dark black bags, 1 bright red bag.
faded brown bags contain 5 mirrored bronze bags, 1 mirrored coral bag, 5 muted white bags, 5 dim orange bags.
faded violet bags contain 4 pale white bags, 1 bright tomato bag, 4 posh gray bags.
mirrored teal bags contain 1 pale crimson bag, 4 striped lime bags, 5 plaid purple bags.
drab crimson bags contain 5 mirrored lime bags.
posh maroon bags contain 5 faded olive bags, 4 bright black bags, 5 muted tomato bags.
muted magenta bags contain 5 shiny olive bags, 4 posh purple bags.
muted olive bags contain 3 dull plum bags, 4 clear turquoise bags, 4 clear purple bags.
dark tan bags contain 4 vibrant teal bags, 3 dotted chartreuse bags, 5 clear crimson bags.
dotted violet bags contain 4 dull cyan bags, 2 wavy blue bags, 1 posh yellow bag.
mirrored turquoise bags contain 4 pale crimson bags, 2 faded lime bags.
muted teal bags contain 5 vibrant indigo bags, 2 bright tomato bags, 5 drab magenta bags, 2 dim red bags.
plaid plum bags contain 1 mirrored bronze bag, 5 dim orange bags, 2 dim beige bags.
posh gray bags contain 1 faded coral bag, 2 dark crimson bags.
dotted tomato bags contain 1 plaid silver bag, 4 clear gold bags, 1 faded violet bag.
shiny teal bags contain 2 muted plum bags, 3 mirrored olive bags.
posh orange bags contain 1 faded aqua bag, 1 wavy indigo bag, 2 muted magenta bags, 2 faded beige bags.
light tan bags contain 1 faded lime bag.
pale cyan bags contain 4 drab black bags, 2 wavy black bags.
dark olive bags contain 4 dark gray bags, 5 drab olive bags, 2 light coral bags.
drab olive bags contain 1 dotted fuchsia bag, 4 dull tan bags, 3 plaid turquoise bags.
vibrant brown bags contain 3 striped bronze bags, 1 dull yellow bag, 5 wavy yellow bags, 4 drab cyan bags.
wavy gray bags contain 1 faded lime bag, 3 shiny purple bags, 3 vibrant cyan bags, 2 dotted fuchsia bags.
pale gray bags contain 2 striped orange bags, 1 plaid green bag, 2 faded orange bags, 3 mirrored blue bags.
posh crimson bags contain 3 clear orange bags, 5 faded indigo bags.
vibrant teal bags contain 5 dark black bags, 4 dim silver bags.
faded bronze bags contain 2 posh bronze bags, 4 bright tomato bags.
plaid violet bags contain 3 plaid indigo bags, 5 dull black bags, 1 striped indigo bag, 2 drab blue bags.
muted black bags contain 1 pale indigo bag.
bright green bags contain 5 vibrant orange bags, 3 muted orange bags, 1 pale violet bag, 1 muted olive bag.
wavy chartreuse bags contain 1 dark fuchsia bag, 1 vibrant teal bag, 1 wavy purple bag.
dotted gray bags contain 3 plaid chartreuse bags, 2 faded violet bags.
dull violet bags contain 4 wavy gold bags.
bright tomato bags contain no other bags.
plaid purple bags contain 2 drab lavender bags, 3 dim red bags, 5 dark chartreuse bags, 5 dotted brown bags.
vibrant chartreuse bags contain 5 striped orange bags, 4 faded blue bags, 1 plaid salmon bag.
light silver bags contain 5 posh purple bags, 5 light tan bags, 4 pale fuchsia bags.
dim salmon bags contain 4 shiny lime bags, 3 light tan bags, 5 dim orange bags, 5 mirrored coral bags.
plaid gray bags contain 2 bright coral bags, 2 plaid tan bags, 3 dark crimson bags, 4 dotted lavender bags.
dull gold bags contain 2 mirrored indigo bags.
light blue bags contain 2 striped gray bags, 1 wavy tan bag, 4 clear crimson bags.
wavy brown bags contain 3 light coral bags, 4 dotted maroon bags, 5 dim silver bags, 2 plaid black bags.
drab black bags contain 4 dull yellow bags.
shiny plum bags contain 5 drab olive bags, 4 dark tan bags.
dim red bags contain 2 light violet bags, 4 clear coral bags.
vibrant salmon bags contain 5 striped violet bags, 2 pale silver bags, 5 clear orange bags.
dull crimson bags contain 1 vibrant blue bag, 4 posh red bags, 5 muted bronze bags, 4 dim olive bags.
vibrant purple bags contain 3 muted turquoise bags, 5 striped gold bags.
vibrant plum bags contain 1 dull teal bag, 5 posh purple bags.
shiny silver bags contain 1 clear plum bag.
plaid beige bags contain 1 shiny magenta bag, 5 posh gray bags.
dull tan bags contain 1 posh purple bag.
mirrored coral bags contain 3 pale aqua bags.
dim brown bags contain 2 dark turquoise bags, 4 dark orange bags, 3 posh olive bags, 2 light chartreuse bags.
muted gold bags contain no other bags.
drab lavender bags contain 5 shiny purple bags, 5 dotted bronze bags.
muted red bags contain 3 drab lavender bags, 1 dim tomato bag, 3 light olive bags, 2 dull aqua bags.
dotted lime bags contain 4 drab magenta bags, 1 dark black bag, 2 posh purple bags.
plaid turquoise bags contain no other bags.
clear orange bags contain 5 shiny magenta bags, 4 light indigo bags, 1 muted yellow bag.
dim gray bags contain 4 posh fuchsia bags, 5 clear gray bags, 2 clear beige bags.
muted tan bags contain 5 vibrant turquoise bags, 2 mirrored gold bags.
shiny white bags contain 5 dim indigo bags, 4 shiny blue bags, 5 clear lavender bags, 3 dim teal bags.
pale lavender bags contain 1 faded green bag, 3 light plum bags, 1 dim bronze bag.
bright plum bags contain 2 mirrored turquoise bags, 2 shiny orange bags, 5 bright tomato bags, 1 bright aqua bag.
vibrant tan bags contain 5 dark lavender bags.
wavy olive bags contain 4 dull salmon bags, 5 mirrored purple bags.
vibrant yellow bags contain 2 striped white bags, 3 pale gold bags, 4 muted gold bags, 2 dull cyan bags.
clear maroon bags contain 4 faded turquoise bags, 3 bright blue bags, 4 dark brown bags, 2 shiny indigo bags.
muted tomato bags contain 4 dim aqua bags, 5 light black bags.
striped lavender bags contain 4 striped gold bags, 4 dark turquoise bags.
dull chartreuse bags contain 1 dotted purple bag, 2 wavy indigo bags, 2 posh blue bags.
dark blue bags contain 2 drab orange bags.
clear coral bags contain 1 muted gold bag, 5 drab magenta bags, 5 bright tomato bags.
clear black bags contain 1 faded coral bag.
dotted blue bags contain 3 shiny gold bags, 3 mirrored purple bags, 5 pale maroon bags.
muted turquoise bags contain 3 posh green bags.
dotted indigo bags contain 1 mirrored tan bag.
dotted olive bags contain 2 vibrant indigo bags, 4 dull purple bags, 4 dull salmon bags.
dull yellow bags contain 1 dim red bag, 3 pale fuchsia bags, 5 shiny gold bags, 4 dull aqua bags.
wavy white bags contain 2 pale salmon bags, 2 striped beige bags, 1 pale aqua bag.
clear beige bags contain 3 dull yellow bags, 2 shiny red bags.
light gold bags contain 2 striped magenta bags, 2 pale salmon bags.
dotted brown bags contain 1 drab blue bag.
posh gold bags contain 1 wavy purple bag, 1 clear magenta bag, 3 clear turquoise bags.
bright yellow bags contain 2 dull maroon bags, 2 muted olive bags, 5 wavy teal bags.
striped plum bags contain 2 striped magenta bags.
striped turquoise bags contain 4 pale magenta bags, 3 muted green bags, 4 pale blue bags.
dull cyan bags contain 4 plaid cyan bags, 1 vibrant cyan bag, 5 shiny violet bags, 3 drab lavender bags.
vibrant silver bags contain 3 vibrant gray bags, 3 light tomato bags.
light lime bags contain 5 plaid crimson bags.
posh yellow bags contain 1 mirrored blue bag, 1 bright orange bag, 2 drab lavender bags.
dull blue bags contain 4 pale chartreuse bags, 2 light maroon bags.
faded black bags contain 4 mirrored brown bags.
bright bronze bags contain 5 dim silver bags.
drab salmon bags contain 1 plaid purple bag, 4 vibrant crimson bags.
drab maroon bags contain 5 shiny gray bags, 4 dotted tan bags.
plaid red bags contain 5 dotted indigo bags.
posh chartreuse bags contain 4 posh silver bags, 4 pale white bags, 3 light chartreuse bags, 1 light black bag.
wavy yellow bags contain 5 plaid turquoise bags, 1 dark gray bag.
vibrant beige bags contain 4 pale gray bags, 4 light blue bags.
light yellow bags contain 3 dim fuchsia bags, 2 clear violet bags, 4 clear teal bags.
dim black bags contain 1 wavy beige bag, 1 dark lavender bag.
shiny bronze bags contain 1 faded lime bag, 5 vibrant green bags.
striped bronze bags contain 1 vibrant cyan bag, 1 faded tomato bag, 5 faded silver bags, 5 dim violet bags.
dim olive bags contain 5 clear fuchsia bags, 4 dark maroon bags.
muted blue bags contain 1 plaid salmon bag, 3 posh purple bags, 5 drab orange bags.
pale fuchsia bags contain 3 dull teal bags.
drab tan bags contain 4 dull fuchsia bags, 2 dark violet bags.
striped white bags contain no other bags.
bright gray bags contain 4 dark bronze bags, 2 vibrant lavender bags, 5 faded cyan bags.
pale aqua bags contain 1 dim red bag.
posh tomato bags contain 3 bright orange bags, 1 dull bronze bag.
mirrored yellow bags contain 5 posh brown bags, 2 plaid brown bags, 2 muted gold bags.
drab cyan bags contain 2 dull silver bags, 1 light silver bag, 5 vibrant gray bags, 5 mirrored crimson bags.
wavy gold bags contain 2 muted plum bags.
plaid teal bags contain 4 shiny black bags, 4 bright olive bags, 3 mirrored maroon bags.
posh black bags contain 4 faded orange bags.
striped silver bags contain 5 muted cyan bags.
bright orange bags contain 1 dull tan bag, 5 dull aqua bags, 5 plaid turquoise bags, 1 bright tomato bag.
clear lime bags contain 1 pale violet bag.
vibrant gray bags contain 5 bright black bags, 1 vibrant purple bag, 5 mirrored coral bags, 5 plaid magenta bags.
mirrored salmon bags contain 2 faded coral bags, 4 striped silver bags.
dim coral bags contain 4 striped white bags, 1 light violet bag, 1 dark brown bag.
dark indigo bags contain 5 vibrant cyan bags.
dotted bronze bags contain 1 drab olive bag, 4 bright orange bags, 1 striped yellow bag.
dotted plum bags contain 3 wavy beige bags, 2 striped magenta bags.
mirrored aqua bags contain 2 wavy purple bags, 5 dull brown bags, 1 clear maroon bag, 4 posh orange bags.
dark green bags contain 2 pale yellow bags, 4 vibrant cyan bags, 4 dim turquoise bags.
dotted chartreuse bags contain 2 drab magenta bags, 1 clear coral bag, 5 striped white bags.
pale brown bags contain 5 faded olive bags, 4 light turquoise bags, 2 bright blue bags.
drab gold bags contain 4 wavy chartreuse bags, 2 plaid aqua bags, 4 wavy orange bags.
plaid blue bags contain 4 faded fuchsia bags, 3 pale bronze bags, 5 faded violet bags.
bright lime bags contain 2 muted turquoise bags.
clear cyan bags contain 1 muted crimson bag, 5 dotted gold bags, 5 dotted crimson bags.
bright indigo bags contain 4 muted turquoise bags, 2 shiny black bags.
posh lavender bags contain 1 faded tan bag, 4 dark coral bags, 2 muted plum bags.
striped maroon bags contain 4 dull magenta bags, 5 pale tan bags, 2 dark teal bags, 3 dotted gray bags.
light aqua bags contain 1 faded lime bag.
shiny coral bags contain 2 dim orange bags, 4 striped beige bags.
vibrant green bags contain 3 shiny magenta bags, 4 drab magenta bags, 4 posh silver bags, 2 shiny orange bags.
plaid maroon bags contain 3 dark coral bags, 5 dim green bags, 4 dim lime bags.
dull red bags contain 4 shiny fuchsia bags.
light crimson bags contain 2 dull tomato bags.
pale beige bags contain 4 muted gold bags.
dotted lavender bags contain 1 muted gold bag, 1 dim violet bag.
dim orange bags contain 2 faded chartreuse bags.
clear white bags contain 3 posh teal bags.
muted brown bags contain 5 light cyan bags, 4 drab purple bags, 2 light black bags, 1 dark indigo bag.
striped aqua bags contain 2 faded aqua bags.
dotted turquoise bags contain 2 shiny white bags, 1 striped lavender bag.
dim maroon bags contain 3 drab tan bags, 4 muted tomato bags, 2 striped lavender bags, 3 plaid silver bags.
drab tomato bags contain 1 wavy gray bag, 2 vibrant teal bags.
striped green bags contain 4 vibrant tomato bags, 3 dull brown bags, 4 light tomato bags.
dull magenta bags contain 2 plaid chartreuse bags, 2 light indigo bags.
vibrant lavender bags contain 1 dull purple bag, 2 shiny crimson bags.
faded olive bags contain 1 dark gray bag.
bright aqua bags contain 2 pale crimson bags.
wavy aqua bags contain 3 light gray bags, 1 posh red bag.
shiny turquoise bags contain 4 dotted black bags, 4 drab gray bags, 1 muted magenta bag.
muted green bags contain 5 striped white bags, 3 vibrant turquoise bags, 1 dull turquoise bag.
faded yellow bags contain 2 dim red bags.
clear fuchsia bags contain 1 dark salmon bag, 2 dim lime bags, 5 drab magenta bags, 2 faded orange bags.
bright coral bags contain 4 wavy gray bags, 1 dull tan bag.
muted bronze bags contain 3 drab turquoise bags, 5 clear aqua bags, 4 dull aqua bags, 2 dim violet bags.
clear violet bags contain 1 plaid salmon bag, 5 dark lavender bags.
dark teal bags contain 2 clear teal bags, 3 dull chartreuse bags, 3 wavy purple bags.
posh green bags contain 5 dark crimson bags, 2 dotted chartreuse bags.
faded magenta bags contain 1 striped silver bag, 5 light beige bags, 1 posh olive bag, 5 vibrant coral bags.
clear crimson bags contain 3 drab gray bags.
pale indigo bags contain 1 faded red bag.
dull brown bags contain 1 clear purple bag.
clear brown bags contain 2 mirrored coral bags, 5 pale white bags, 1 dotted bronze bag, 5 wavy orange bags.
drab aqua bags contain 1 vibrant chartreuse bag, 4 drab tomato bags, 2 shiny gold bags.
shiny black bags contain 2 dull lime bags.
dotted purple bags contain 3 faded orange bags, 1 dotted chartreuse bag, 4 pale indigo bags.
posh lime bags contain 5 dotted tan bags, 1 shiny magenta bag.
drab plum bags contain 3 plaid silver bags, 2 dark aqua bags, 2 bright crimson bags, 1 posh purple bag.
striped gray bags contain 1 posh purple bag.
light magenta bags contain 5 dark brown bags.
striped magenta bags contain 5 faded orange bags, 2 faded red bags, 2 posh silver bags, 3 dark gray bags.
dark bronze bags contain 1 faded coral bag, 1 shiny plum bag, 5 shiny white bags.
bright teal bags contain 4 dark tan bags.
mirrored tomato bags contain 1 muted tomato bag, 2 dull yellow bags, 2 posh white bags.
bright white bags contain 5 faded silver bags, 5 dark chartreuse bags.
wavy turquoise bags contain 3 drab fuchsia bags, 3 dull violet bags, 4 dark tomato bags.
posh coral bags contain 5 shiny gold bags, 3 vibrant olive bags, 3 dotted brown bags, 2 dim chartreuse bags.
dotted fuchsia bags contain 5 dim coral bags, 5 drab gray bags, 4 striped white bags, 4 bright orange bags.
vibrant magenta bags contain 3 plaid salmon bags, 1 dotted lime bag, 1 clear crimson bag, 2 dotted tan bags.
faded cyan bags contain 3 faded violet bags, 2 light green bags.
dark red bags contain 1 dim violet bag, 4 vibrant olive bags, 5 drab olive bags, 5 faded yellow bags.
faded aqua bags contain 5 bright plum bags, 3 light black bags, 1 muted yellow bag, 3 mirrored bronze bags.
wavy blue bags contain 5 wavy teal bags, 5 shiny magenta bags, 5 drab olive bags, 5 dark coral bags.
striped orange bags contain 4 striped white bags, 3 dim red bags, 3 clear coral bags, 4 shiny gold bags.
pale orange bags contain 3 dim green bags, 3 dark olive bags.
drab violet bags contain 3 striped gold bags, 5 plaid magenta bags.
dark salmon bags contain 5 clear coral bags, 2 muted gold bags, 5 bright tomato bags.
mirrored lime bags contain 2 plaid maroon bags, 1 shiny olive bag, 2 vibrant coral bags.
dull black bags contain 3 vibrant brown bags, 4 clear black bags.
clear gray bags contain 5 drab cyan bags, 3 dotted salmon bags.
light maroon bags contain 5 plaid fuchsia bags, 4 muted purple bags.
drab white bags contain 5 shiny gray bags, 5 posh orange bags.
dim chartreuse bags contain 3 wavy blue bags, 1 bright aqua bag.
dotted black bags contain 5 pale bronze bags.
striped purple bags contain 3 vibrant black bags.
shiny tomato bags contain 2 faded blue bags.
drab coral bags contain 2 pale beige bags.
striped cyan bags contain 2 dark lime bags, 5 light plum bags.
dotted yellow bags contain 4 dotted bronze bags.
dark maroon bags contain 1 dark silver bag, 1 dark black bag, 2 striped silver bags, 5 striped white bags.
drab chartreuse bags contain 4 muted maroon bags.
plaid lime bags contain 2 dim teal bags, 2 posh chartreuse bags.
dim fuchsia bags contain 4 mirrored turquoise bags, 4 light gray bags, 5 mirrored black bags, 3 faded yellow bags.
bright cyan bags contain 4 plaid crimson bags, 3 dim plum bags, 4 mirrored olive bags, 4 muted plum bags.
striped olive bags contain 3 shiny purple bags.
dull bronze bags contain 2 drab tomato bags, 3 dull teal bags, 4 dim silver bags, 4 faded turquoise bags.
wavy purple bags contain 4 wavy black bags, 4 clear purple bags, 2 dark fuchsia bags.
wavy silver bags contain 2 plaid tomato bags.
plaid magenta bags contain 3 faded red bags, 2 vibrant black bags.
dark chartreuse bags contain 1 dim red bag, 4 vibrant olive bags, 4 dark black bags.
striped teal bags contain 1 mirrored fuchsia bag, 1 wavy black bag, 5 dim yellow bags, 1 dotted orange bag.
dull green bags contain 3 posh tomato bags, 4 bright plum bags.
drab lime bags contain 2 shiny blue bags, 5 light bronze bags, 4 dim magenta bags, 3 mirrored crimson bags.
faded green bags contain 1 posh fuchsia bag, 5 dim plum bags.
light violet bags contain 2 dark crimson bags, 3 dark gray bags, 5 muted gold bags.
bright fuchsia bags contain 4 dull tomato bags, 4 wavy fuchsia bags, 4 clear salmon bags.
drab green bags contain 4 vibrant black bags, 5 pale salmon bags, 4 pale olive bags, 5 clear crimson bags.
muted salmon bags contain 2 clear magenta bags, 3 wavy crimson bags.
shiny maroon bags contain 2 dim crimson bags.`
