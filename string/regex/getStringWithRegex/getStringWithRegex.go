package main

import (
	"fmt"
	"regexp"
)

var s1 = `“Er — yes, Harry — about this cupboard. Your aunt and I have been thinking… you’re really
  getting a bit big for it… we think it might be nice if you moved into Dudley’s second bedroom.
   “Why?” said Harry.
   “Don’t ask questions!” snapped his uncle. “Take this stuff upstairs, now.”
   The Dursleys’ house had four bedrooms: one for Uncle Vernon and Aunt Petunia, one for
  visitors (usually Uncle Vernon’s sister, Marge), one where Dudley slept, and one where Dudley
  kept all the toys and things that wouldn’t fit into his first bedroom. It only took Harry one trip
  upstairs to move everything he owned from the cupboard to this room. He sat down on the bed
  and stared around him. Nearly everything in here was broken. The month-old video camera was
  lying on top of a small, working tank Dudley had once driven over the next door neighbor’s dog;
  in the corner was Dudley’s first-ever television set, which he’d put his foot through when his
  favorite program had been canceled; there was a large birdcage, which had once held a parrot
  that Dudley had swapped at school for a real air rifle, which was up on a shelf with the end all
  bent because Dudley had sat on it. Other shelves were full of books. They were the only things in
  the room that looked as though they’d never been touched.
   From downstairs came the sound of Dudley bawling at his mother, “I don’t want him in there… I
  need that room… make him get out…”
   Harry sighed and stretched out on the bed. Yesterday he’d have given anything to be up here.
  Today he’d rather be back in his cupboard with that letter than up here without it.
   Next morning at breakfast, everyone was rather quiet. Dudley was in shock. He’d screamed,
  whacked his father with his Smelting stick, been sick on purpose, kicked his mother, and thrown
  his tortoise through the greenhouse roof, and he still didn’t have his room back. Harry was
  thinking about this time yesterday and bitterly wishing he’d opened the letter in the hall. Uncle
  Vernon and Aunt Petunia kept looking at each other darkly.
   When the mail arrived, Uncle Vernon, who seemed to be trying to be nice to Harry, made Dudley
  go and get it. They heard him banging things with his Smelting stick all the way down the hall.
  Then he shouted, “There’s another one! ‘Mr. H. Potter, The Smallest Bedroom, 4 Privet Drive —
  ’”
   With a strangled cry, Uncle Vernon leapt from his seat and ran down the hall, Harry right behind
  him. Uncle Vernon had to wrestle Dudley to the ground to get the letter from him, which was
  made difficult by the fact that Harry had grabbed Uncle Vernon around the neck from behind.
  After a minute of confused fighting, in which everyone got hit a lot by the Smelting stick, Uncle
  Vernon straightened up, gasping for breath, with Harry’s letter clutched in his hand.8ef49396-f057-4c4e-ab07-9df29be43963`

var s2 = `“If yeh know where to go,” said Hagrid.
   Harry had never been to London before. Although Hagrid seemed to know where he was going,
  he was obviously not used to getting there in an ordinary way. He got stuck in the ticket barrier on the
  Underground, and complained loudly that the seats were too small and the trains too slow.
   “I don’t know how the Muggles manage without magic,” he said as they climbed a broken-down
  escalator that led up to a bustling road lined with shops.
   Hagrid was so huge that he parted the crowd easily; all Harry had to do was keep close behind
  him. They passed book shops and music stores, hamburger restaurants and cinemas, but nowhere
  that looked as if it could sell you a magic wand. This was just an ordinary street full of ordinary
  people. Could there really be piles of wizard gold buried miles beneath them? Were there really
  shops that sold spell books and broomsticks? Might this not all be some huge joke that the
  Dursleys had cooked up? If Harry hadn’t known that the Dursleys had no sense of humor, he
  might have thought so; yet somehow, even though everything Hagrid had told him so far was
  unbelievable, Harry couldn’t help trusting him.
   “This is it,” said Hagrid, coming to a halt, “the Leaky Cauldron. It’s a famous place.”
   It was a tiny, grubby-looking pub. If Hagrid hadn’t pointed it out, Harry wouldn’t have noticed it
  was there. The people hurrying by didn’t glance at it. Their eyes slid from the big book shop on
  one side to the record shop on the other as if they couldn’t see the Leaky Cauldron at all. In fact,
  Harry had the most peculiar feeling that only he and Hagrid could see it. Before he could
  mention this, Hagrid had steered him inside.
   For a famous place, it was very dark and shabby. A few old women were sitting in a corner,
  drinking tiny glasses of sherry. One of them was smoking a long pipe. A little man in a top hat
  was talking to the old bartender, who was quite bald and looked like a toothless walnut. The low
  buzz of chatter stopped when they walked in. Everyone seemed to know Hagrid; they waved and
  smiled at him, and the bartender reached for a glass, saying, “The usual, Hagrid?”
   “Can’t, Tom, I’m on Hogwarts business,” said Hagrid, clapping his great hand on Harry’s
  shoulder and making Harry’s knees buckle.
   “Good Lord,” said the bartender, peering at Harry, “is this — can this be —?”
   The Leaky Cauldron had suddenly gone completely still and silent.
   “Bless my soul,” whispered the old bartender, “Harry Potter… what an honor.”`

func main() {
	var paragraph string
	paragraph = getParagraph(paragraph)
	fmt.Println(paragraph)
}

func getParagraph(paragraph string) string {
	var paragraphString string
	currentParaLength := len(paragraph)
	if currentParaLength > 0 {
		concatLength := 1
		tmpString := getNumberAndStringFromParagraph()
		requireLength := 500 - currentParaLength
		tmplength := len(tmpString)
		if tmplength < requireLength {
			concatLength = tmplength
		} else {
			concatLength = requireLength
		}
		paragraph += string([]byte(tmpString)[:concatLength])
		if len(paragraph) < 500 {
			paragraph = getParagraph(paragraph)
		}
		return paragraph
	} else {
		paragraphString = getNumberAndStringFromParagraph()
		if len(paragraphString) < 500 {
			paragraphString = getParagraph(paragraphString)
		}
	}
	return paragraphString
}

func getNumberAndStringFromParagraph() string {
	str1 := "Hello X42 I'm a Y-32.35 string Z30"
	re := regexp.MustCompile(`[A-Za-z0-9]`)

	fmt.Printf("Pattern: %v\n", re.String()) // Print Pattern

	fmt.Printf("String contains any match: %v\n", re.MatchString(str1)) // True

	submatchall := re.FindAllString(str1, -1)
	var finalString string
	for _, element := range submatchall {
		finalString += element
	}
	return finalString
}

func getNumberFromParagraph() {
	str1 := "Hello X42 I'm a Y-32.35 string Z30"
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	fmt.Printf("Pattern: %v\n", re.String()) // Print Pattern

	fmt.Printf("String contains any match: %v\n", re.MatchString(str1)) // True

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
