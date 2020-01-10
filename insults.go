// Copyright 2020 Vanessa Sochat. All rights reserved.
// Use of this source code is governed by the Polyform Strict license
// that can be found in the LICENSE file and available at
// https://polyformproject.org/licenses/noncommercial/1.0.0

// STOPPED HERE - get insults form
// https://github.com/ncitron/insultAPI and generate API

package main

import (
	"math/rand"
	"syscall/js"
	"time"
)

func getInsult() string {

	insults := [...]string{"Away, you three-inch fool!",
		"Away, you starvelling, you elf-skin, you dried neat’s-tongue, bull’s-pizzle, you stock-fish!",
		"A most notable coward, an infinite and endless liar, an hourly promise breaker, the owner of no one good quality.",
		"Come, come, you froward and unable worms!",
		"Go, prick thy face, and over-red thy fear, Thou lily-liver’d boy.",
		"His wit’s as thick as a Tewkesbury mustard.",
		"I am pigeon-liver’d and lack gall.",
		"I am sick when I do look on thee.",
		"I must tell you friendly in your ear, sell when you can, you are not for all markets.",
		"If thou wilt needs marry, marry a fool; for wise men know well enough what monsters you make of them.",
		"I’ll beat thee, but I would infect my hands.",
		"I scorn you, scurvy companion.",
		"Methink’st thou art a general offence and every man should beat thee.",
		"More of your conversation would infect my brain.",
		"My wife’s a hobby horse!",
		"Peace, ye fat guts!",
		"Poisonous bunch-backed toad!",
		"The rankest compound of villainous smell that ever offended nostril",
		"The tartness of his face sours ripe grapes.",
		"There’s no more faith in thee than in a stewed prune.",
		"Thine forward voice, now, is to speak well of thine friend; thine backward voice is to utter foul speeches and to detract.",
		"That trunk of humours, that bolting-hutch of beastliness, that swollen parcel of dropsies, that huge bombard of sack, that stuffed cloak-bag of guts, that roasted Manningtree ox with pudding in his belly, that reverend vice, that grey Iniquity, that father ruffian, that vanity in years?",
		"Thine face is not worth sunburning.",
		"This woman’s an easy glove, my lord, she goes off and on at pleasure.",
		"Thou art a boil, a plague sore",
		"Was the Duke a flesh-monger, a fool and a coward?",
		"Thou art as fat as butter.",
		"Here is the babe, as loathsome as a toad.",
		"Like the toad; ugly and venomous.",
		"Thou art unfit for any place but hell.",
		"Thou cream faced loon",
		"Thou clay-brained guts, thou knotty-pated fool, thou whoreson obscene greasy tallow-catch!",
		"Thou damned and luxurious mountain goat.",
		"Thou elvish-mark’d, abortive, rooting hog!",
		"Thou leathern-jerkin, crystal-button, knot-pated, agatering, puke-stocking, caddis-garter, smooth-tongue, Spanish pouch!",
		"Thou lump of foul deformity",
		"That poisonous bunch-back’d toad!",
		"Thou sodden-witted lord! Thou hast no more brain than I have in mine elbows.",
		"Thou subtle, perjur’d, false, disloyal man!",
		"Thou whoreson zed , thou unnecessary letter!",
		"Thy sin’s not accidental, but a trade.",
		"Thy tongue outvenoms all the worms of Nile.",
		"Would thou wert clean enough to spit upon.",
		"Would thou wouldst burst!",
		"You poor, base, rascally, cheating lack-linen mate!",
		"You are as a candle, the better burnt out.",
		"You scullion! You rampallian! You fustilarian! I’ll tickle your catastrophe!",
		"You starvelling, you eel-skin, you dried neat’s-tongue, you bull’s-pizzle, you stock-fish–O for breath to utter what is like thee!-you tailor’s-yard, you sheath, you bow-case, you vile standing tuck!",
		"Your brain is as dry as the remainder biscuit after voyage.",
		"Virginity breeds mites, much like a cheese.",
		"Villain, I have done thy mother",
	}

	return insults[rand.Intn(len(insults))]
}

// returnResult back to the browser, in the innerHTML of the result element
func returnResult(output string, divid string) {
	js.Global().Get("document").
		Call("getElementById", divid).
		Set("innerHTML", output)
}

// generateInsult is linked with the JavaScript function of the same name.
// It returns a randomly selected insult.
func generateInsult(this js.Value, val []js.Value) interface{} {
	//fmt.Println("The container binary is:", val[0])

	rand.Seed(time.Now().Unix())
	insult := getInsult()

	// Send result back to browser, key is div id, content is string
	returnResult(insult, "insult")
	return nil
}

// randomInsult is linked with the JavaScript function of the same name.
// It returns a randomly generated insult
func generateRandomInsult(this js.Value, val []js.Value) interface{} {
	//fmt.Println("The container binary is:", val[0])

	insult := randomInsult()

	// Send result back to browser, key is div id, content is string
	returnResult(insult, "insult")
	return nil
}

func randomInsult() string {

	first := [...]string{"artless",
		"bawdy",
		"beslubbering",
		"bootless",
		"churlish",
		"cockered",
		"clouted",
		"craven",
		"currish",
		"dankish",
		"dissembling",
		"droning",
		"errant",
		"fawning",
		"fobbing",
		"froward",
		"frothy",
		"gleeking",
		"goatish",
		"gorbellied",
		"impertinent",
		"infectious",
		"jarring",
		"loggerheaded",
		"lumpish",
		"mammering",
		"mangled",
		"mewling",
		"paunchy",
		"pribbling",
		"puking",
		"puny",
		"qualling",
		"rank",
		"reeky",
		"roguish",
		"ruttish",
		"saucy",
		"spleeny",
		"spongy",
		"surly",
		"tottering",
		"unmuzzled",
		"vain",
		"venomed",
		"villainous",
		"warped",
		"wayward",
		"weedy",
		"yeasty"}

	second := [...]string{
		"base-court",
		"bat-fowling",
		"beef-witted",
		"beetle-headed",
		"boil-brained",
		"clapper-clawed",
		"clay-brained",
		"common-kissing",
		"crook-pated",
		"dismal-dreaming",
		"dizzy-eyed",
		"doghearted",
		"dread-bolted",
		"earth-vexing",
		"elf-skinned",
		"fat-kidneyed",
		"fen-sucked",
		"flap-mouthed",
		"fly-bitten",
		"folly-fallen",
		"fool-born",
		"full-gorged",
		"guts-griping",
		"half-faced",
		"hasty-witted",
		"hedge-born",
		"hell-hated",
		"idle-headed",
		"ill-breeding",
		"ill-nurtured",
		"knotty-pated",
		"milk-livered",
		"motley-minded",
		"onion-eyed",
		"plume-plucked",
		"pottle-deep",
		"pox-marked",
		"reeling-ripe",
		"rough-hewn",
		"rude-growing",
		"rump-fed",
		"shard-borne",
		"sheep-biting",
		"spur-galled",
		"swag-bellied",
		"tardy-gaited",
		"tickle-brained",
		"toad-spotted",
		"unchin-snouted",
		"weather-bitten"}

	nouns := [...]string{
		"apple-john",
		"baggage",
		"barnacle",
		"bladder",
		"boar-pig",
		"bugbear",
		"bum-bailey",
		"canker-blossom",
		"clack-dish",
		"clotpole",
		"coxcomb",
		"codpiece",
		"death-token",
		"dewberry",
		"flap-dragon",
		"flax-wench",
		"flirt-gill",
		"foot-licker",
		"fustilarian",
		"giglet",
		"gudgeon",
		"haggard",
		"harpy",
		"hedge-pig",
		"horn-beast",
		"hugger-mugger",
		"joithead",
		"lewdster",
		"lout",
		"maggot-pie",
		"malt-worm",
		"mammet",
		"measle",
		"minnow",
		"miscreant",
		"moldwarp",
		"mumble-news",
		"nut-hook",
		"pigeon-egg",
		"pignut",
		"puttock",
		"pumpion",
		"ratsbane",
		"scut",
		"skainsmate",
		"strumpet",
		"varlot",
		"vassal",
		"whey-face",
		"wagtail"}

	return "Thou " + first[rand.Intn(len(first))] + " " + second[rand.Intn(len(second))] + " " + nouns[rand.Intn(len(nouns))] + "!"

}
