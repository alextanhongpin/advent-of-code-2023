// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var digits = regexp.MustCompile(`\d+`)

func main() {
	fmt.Println(solve(example)) // 2, 114
	fmt.Println(solve(input))   // 993, 1834108701
}

func solve(input string) (head, tail int) {
	nums := parse(input)
	for _, num := range nums {
		first, last := eval(num)
		head += first
		tail += last
	}
	return
}

func eval(nums []int) (int, int) {
	delta := func(ns []int) []int {
		res := make([]int, len(ns)-1)
		for i := range ns {
			if i == 0 {
				continue
			}
			res[i-1] = ns[i] - ns[i-1]
		}
		return res
	}

	last := nums[len(nums)-1]

	heads := []int{nums[0]}
	for {
		nums = delta(nums)
		if nums[0] == 0 && nums[len(nums)-1] == 0 {
			break
		}
		heads = append(heads, nums[0])
		last += nums[len(nums)-1]
	}
	for len(heads) != 1 {
		var last, secondLast int
		last, heads = heads[len(heads)-1], heads[:len(heads)-1]
		secondLast, heads = heads[len(heads)-1], heads[:len(heads)-1]
		heads = append(heads, secondLast-last)
	}

	return heads[0], last
}

func parse(s string) [][]int {
	lines := strings.Split(s, "\n")
	res := make([][]int, len(lines))

	for i, line := range lines {
		res[i] = toDigits(strings.Fields(line))
	}
	return res
}

func toDigits(ns []string) []int {
	res := make([]int, len(ns))
	for i, s := range ns {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res[i] = n
	}
	return res
}

var example = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

var input = `10 15 15 10 0 -15 -35 -60 -90 -125 -165 -210 -260 -315 -375 -440 -510 -585 -665 -750 -840
4 0 4 37 137 373 875 1885 3829 7404 13667 24105 40656 65641 101556 150661 214290 291792 378998 466093 534755
5 14 35 70 124 218 428 968 2337 5554 12525 26654 53986 105566 202537 387238 747109 1465286 2925459 5924110 12089144
3 0 -7 -18 -33 -52 -75 -102 -133 -168 -207 -250 -297 -348 -403 -462 -525 -592 -663 -738 -817
8 11 11 15 49 169 480 1180 2661 5723 11997 24738 50264 100559 198135 385661 746225 1447698 2841038 5675992 11563810
4 5 8 11 23 75 228 590 1366 2981 6341 13334 27729 56736 113748 223498 432804 835074 1628856 3258428 6744587
15 37 66 105 171 312 630 1308 2652 5183 9850 18489 34736 65730 125136 238304 450789 842025 1546714 2787505 4923849
11 14 19 32 70 163 356 711 1309 2252 3665 5698 8528 12361 17434 24017 32415 42970 56063 72116 91594
10 11 18 43 101 208 380 633 980 1429 2016 2969 5221 11764 31003 84858 227915 591377 1480393 3581719 8391418
10 11 11 7 10 57 227 672 1682 3818 8182 16968 34581 69862 140375 280385 555214 1086293 2093697 3966621 7377616
4 22 60 128 252 497 1007 2084 4355 9126 19101 39767 81946 166353 331575 647845 1240504 2329276 4292420 7768040 13807018
10 15 27 55 108 195 325 507 750 1063 1455 1935 2512 3195 3993 4915 5970 7167 8515 10023 11700
9 26 50 97 200 419 852 1655 3092 5646 10224 18487 33369 60032 108084 197314 372211 736262 1527102 3270289 7075854
6 23 50 90 150 237 359 545 910 1816 4238 10563 26263 63234 146161 324219 692090 1427309 2856482 5572820 10644640
9 18 29 46 84 194 508 1324 3261 7518 16270 33232 64426 119206 211645 362482 601985 974330 1544451 2408810 3712200
-1 13 34 62 108 203 410 844 1714 3420 6766 13388 26550 52541 103021 198813 375797 693667 1248238 2187526 3730644
17 30 50 88 178 400 912 2001 4172 8303 15903 29519 53347 94111 162283 273726 451851 730388 1156880 1797018 2739944
8 22 38 56 76 98 122 148 176 206 238 272 308 346 386 428 472 518 566 616 668
13 25 37 54 92 189 425 961 2122 4579 9743 20597 43428 91443 192424 405152 854784 1808438 3833845 8124774 17152503
14 27 47 91 194 428 934 1968 3966 7638 14107 25116 43334 72801 119562 192551 304798 475045 729871 1106441 1656010
21 44 77 131 237 454 889 1743 3407 6661 13097 26036 52529 107701 224035 470730 992825 2087570 4348253 8925689 17987933
16 41 78 124 175 226 271 303 314 295 236 126 -47 -296 -635 -1079 -1644 -2347 -3206 -4240 -5469
5 0 -3 -6 -1 52 254 801 2021 4413 8687 15796 26930 43422 66570 97738 140394 209562 364240 802888 2118158
16 19 36 81 168 312 526 820 1219 1830 3018 5838 13104 32043 80740 203171 502618 1212337 2837948 6431459 14096328
11 13 31 95 264 641 1400 2845 5530 10478 19546 35992 65309 116400 203177 346676 577789 940723 1497305 2332261 3559606
7 20 44 97 211 436 844 1533 2631 4300 6740 10193 14947 21340 29764 40669 54567 72036 93724 120353 152723
19 31 49 81 150 304 628 1255 2376 4260 7314 12240 20381 34391 59415 105024 188217 337877 601151 1052315 1804784
-7 -4 4 17 35 58 86 119 157 200 248 301 359 422 490 563 641 724 812 905 1003
20 32 57 104 183 305 482 727 1054 1478 2015 2682 3497 4479 5648 7025 8632 10492 12629 15068 17835
10 35 85 173 313 536 923 1657 3103 5938 11380 21639 40913 77737 150571 300782 621788 1322253 2858755 6206037 13391957
2 -6 -17 -33 -48 -28 118 576 1674 3938 8157 15457 27384 45996 73964 114682 172386 252282 360683 505155 694672
13 16 36 97 235 495 932 1623 2710 4529 7960 15307 32375 73111 169497 391824 888927 1965978 4231656 8871265 18144503
1 2 12 46 141 377 906 1987 4028 7640 13712 23519 38876 62349 97528 149356 224491 331654 481884 688580 967159
2 12 40 98 196 341 540 807 1174 1706 2520 3808 5864 9115 14156 21789 33066 49336 72296 104046 147148
24 34 53 94 176 329 607 1122 2125 4196 8676 18587 40453 87698 186772 388113 786082 1556382 3028915 5834240 11204258
6 20 34 49 78 157 370 909 2213 5264 12163 27165 58419 120737 239805 458349 844880 1505764 2601496 4368201 7145540
7 20 37 68 132 254 458 756 1133 1528 1811 1756 1010 -942 -4816 -11572 -22461 -39076 -63407 -97900 -145520
0 3 16 39 67 85 66 -21 -184 -354 -307 403 2397 6084 10507 10689 -7696 -78115 -270042 -716585 -1657937
15 22 32 56 126 304 695 1478 2977 5800 11074 20794 38280 68694 119505 200700 324419 503538 748532 1061716 1427682
8 16 42 102 226 477 985 1996 3936 7490 13696 24054 40650 66295 104679 160540 239848 350004 500054 700918 965634
5 14 31 74 171 372 784 1642 3436 7119 14419 28274 53425 97292 171547 295570 504857 871791 1555795 2920419 5797389
3 0 5 30 101 270 637 1387 2851 5626 10863 20997 41522 85030 179836 387412 836015 1782980 3726077 7596336 15082447
22 39 55 63 64 79 161 407 970 2071 4011 7183 12084 19327 29653 43943 63230 88711 121759 163935 217000
12 19 26 33 40 47 54 61 68 75 82 89 96 103 110 117 124 131 138 145 152
16 17 24 62 168 388 770 1352 2143 3087 3981 4284 2701 -3650 -20783 -60530 -144966 -313379 -632842 -1213742 -2231970
6 19 40 71 106 133 140 138 241 881 3296 10542 29520 75030 177952 401808 876039 1861846 3880099 7955496 16075577
5 9 32 90 217 492 1094 2396 5106 10471 20594 38993 71686 129372 231799 416388 755129 1387844 2588587 4900010 9406491
6 4 -3 -18 -44 -76 -84 12 374 1284 3187 6738 12852 22756 38042 60720 93270 138692 200553 283030 390948
9 14 33 81 191 437 985 2188 4749 9997 20371 40308 77910 148066 278223 518944 963181 1781592 3287525 6056487 11143989
-1 14 46 105 216 439 910 1920 4062 8507 17542 35650 71680 143102 284037 559773 1091909 2100196 3969650 7354671 13335772
5 6 24 76 185 386 742 1377 2539 4717 8859 16785 31976 61066 116590 221869 419365 784436 1447182 2626014 4677711
4 14 28 59 131 286 611 1292 2698 5502 10878 20901 39465 74400 142156 277692 554537 1124231 2292005 4659184 9386660
7 8 21 57 134 290 611 1281 2659 5391 10588 20181 37775 70792 135611 269056 552323 1157758 2435411 5065749 10318224
12 28 48 67 87 139 324 888 2366 5861 13564 29680 62035 124860 243670 463897 866136 1592674 2892544 5195809 9232223
27 39 49 66 116 250 551 1141 2190 3943 6835 11906 22051 45334 103035 247968 604045 1447828 3369512 7578272 16461240
1 -3 -5 13 92 317 864 2098 4759 10287 21353 42681 82266 153115 275662 481034 815373 1345449 2165831 3407917 5251160
-5 3 22 55 103 166 256 436 903 2135 5122 11699 24995 50006 94292 168788 288707 474499 752814 1157399 1729839
19 46 82 119 143 139 112 143 520 2015 6419 17499 42610 95298 199402 395464 750777 1375268 2446788 4251459 7247729
-1 17 60 141 287 553 1040 1933 3586 6701 12686 24338 47089 91209 175683 335218 633501 1189353 2229380 4193603 7943119
15 29 43 57 71 85 99 113 127 141 155 169 183 197 211 225 239 253 267 281 295
0 -2 -2 0 -2 -25 -107 -329 -849 -1926 -3864 -6706 -9308 -7046 13258 83718 275366 734706 1747734 3847880 7992006
19 28 28 12 -27 -94 -192 -327 -526 -876 -1564 -2802 -4297 -3491 8959 59204 210313 599728 1505072 3454428 7404005
14 34 71 136 242 401 616 866 1082 1112 673 -712 -3794 -9697 -20038 -37070 -63850 -104434 -164101 -249608 -369478
24 31 40 56 100 219 490 1014 1894 3189 4841 6606 8140 9729 14997 38881 128413 412577 1214457 3290064 8312645
5 15 28 44 63 85 110 138 169 203 240 280 323 369 418 470 525 583 644 708 775
23 46 87 162 293 505 825 1288 1963 3039 5076 9669 21068 49865 120881 289116 670403 1497672 3220033 6670904 13343955
-3 -4 -4 5 47 164 411 842 1495 2407 3730 6090 11473 25226 60406 147029 351391 815776 1838954 4037649 8667625
13 15 12 8 14 47 129 286 547 943 1506 2268 3260 4511 6047 7890 10057 12559 15400 18576 22074
8 31 64 107 160 223 296 379 472 575 688 811 944 1087 1240 1403 1576 1759 1952 2155 2368
17 36 56 68 65 56 88 287 943 2681 6774 15654 33647 67877 129126 232171 394709 633386 954620 1336800 1698997
17 35 77 159 302 542 958 1728 3223 6149 11747 22061 40284 71192 121676 201382 323469 505495 770441 1147883 1675322
17 29 50 85 145 263 539 1230 2918 6825 15413 33533 70604 144655 289604 567942 1092111 2059397 3807196 6898157 12248075
22 27 39 84 206 470 965 1807 3142 5149 8043 12078 17550 24800 34217 46241 61366 80143 103183 131160 164814
5 0 -7 -15 -11 45 232 675 1547 3068 5505 9179 14487 21949 32292 46585 66441 94304 133841 190461 271985
1 7 15 23 24 3 -69 -249 -656 -1551 -3513 -7794 -16965 -35949 -73364 -142474 -259223 -430017 -610075 -585655 327168
12 22 27 22 16 43 166 482 1153 2518 5395 11781 26339 59395 132811 291403 625279 1314027 2714684 5536213 11180667
-5 -3 2 10 21 35 52 72 95 121 150 182 217 255 296 340 387 437 490 546 605
3 19 48 90 142 213 358 741 1741 4119 9268 19572 38904 73297 131826 227743 379911 614587 967608 1487038 2236338
25 48 96 183 320 509 751 1086 1701 3175 6979 16416 38274 85578 181970 368425 713233 1326450 2380354 4137845 6991212
4 4 4 5 16 63 218 662 1792 4373 9716 19823 37378 65402 106406 161136 227814 304665 402447 581727 1046676
4 -3 -18 -38 -47 0 213 845 2449 6169 14241 30804 63148 123556 231930 419427 733370 1243741 2051608 3299886 5186883
19 30 45 70 108 152 183 191 241 602 1941 5562 13666 29669 58813 109736 198447 357410 654309 1227666 2349912
12 35 84 167 286 446 671 1019 1586 2485 3776 5303 6360 5056 -2825 -25773 -79744 -193193 -414748 -824569 -1550726
9 13 31 84 209 461 914 1663 2837 4658 7632 13053 24184 48847 104894 231467 511605 1116399 2388649 4996382 10205711
5 27 57 85 110 152 276 645 1620 3938 9047 19788 41820 86518 176570 356208 709049 1388191 2666171 5017981 9260988
-6 -15 -24 -29 -26 -11 20 71 146 249 384 555 766 1021 1324 1679 2090 2561 3096 3699 4374
10 23 46 87 155 276 532 1128 2500 5511 11856 24934 51702 106534 219127 450485 924714 1890903 3840374 7722313 15328273
12 17 25 43 78 137 227 355 528 753 1037 1387 1810 2313 2903 3587 4372 5265 6273 7403 8662
3 12 27 60 138 312 679 1416 2830 5445 10178 18703 34167 62507 114724 210601 384509 694130 1233141 2149150 3668456
12 24 42 66 96 132 174 222 276 336 402 474 552 636 726 822 924 1032 1146 1266 1392
-1 10 50 144 341 724 1423 2635 4651 7887 12916 20516 31829 48968 77021 129810 244857 519650 1195369 2843653 6772541
-2 4 15 33 59 97 163 295 565 1106 2191 4441 9294 19927 42865 90495 184568 360432 671079 1188961 2001749
12 39 85 160 280 480 844 1557 2984 5781 11043 20494 36724 63478 106002 171451 269364 412211 616017 901068 1292704
13 26 37 51 84 171 387 904 2122 4936 11250 24946 53691 112255 228462 453555 879683 1668478 3096355 5624317 10002766
9 8 19 57 149 344 723 1409 2577 4464 7379 11713 17949 26672 38579 54489 75353 102264 136467 179369 232549
7 17 27 37 47 57 67 77 87 97 107 117 127 137 147 157 167 177 187 197 207
9 33 82 178 361 701 1327 2483 4619 8523 15498 27586 47839 80635 132035 210175 325685 492125 726426 1049322 1485757
8 17 28 41 56 73 92 113 136 161 188 217 248 281 316 353 392 433 476 521 568
1 13 41 103 233 479 907 1620 2801 4789 8197 14081 24169 41159 69095 113830 183585 289613 446977 675451 1000553
13 22 47 89 150 244 421 812 1713 3746 8165 17415 36102 72592 141527 267626 491229 876142 1520451 2571093 4243102
1 -2 -13 -26 -23 28 189 597 1553 3670 8095 16820 33097 61972 110953 190827 316641 508862 794731 1209826 1799849
18 36 76 160 333 670 1294 2419 4447 8181 15280 29209 57198 114258 231354 471803 963456 1960118 3955188 7885322 15484219
11 7 12 50 163 418 920 1842 3492 6452 11846 21829 40467 75389 141137 266385 509704 993032 1969151 3949503 7930550
1 21 63 142 280 506 856 1373 2107 3115 4461 6216 8458 11272 14750 18991 24101 30193 37387 45810 55596
9 18 36 63 99 144 198 261 333 414 504 603 711 828 954 1089 1233 1386 1548 1719 1899
11 33 72 131 220 375 694 1398 2930 6114 12411 24332 46101 84706 151535 264869 453597 762631 1260634 2050833 3285874
5 17 39 67 98 129 149 130 49 23 729 4438 17258 53617 144707 353665 801829 1712653 3483011 6795925 12794522
8 29 72 150 285 513 896 1560 2787 5198 10073 19863 38958 74784 139311 251063 437730 739491 1213166 1937324 3018483
15 20 34 57 95 180 406 983 2304 5013 10059 18740 32822 55035 90738 152574 272039 526168 1096315 2394224 5328817
4 6 8 10 12 14 16 18 20 22 24 26 28 30 32 34 36 38 40 42 44
16 30 54 99 186 350 648 1184 2182 4169 8376 17528 37276 78627 161854 322519 620420 1152480 2070834 3607641 6108454
-6 4 41 117 244 440 737 1188 1868 2862 4231 5945 7770 9094 8675 4292 -7722 -33088 -80355 -161867 -294952
18 35 54 80 122 188 274 356 422 631 1773 6347 20793 59727 153448 360518 787854 1621495 3171980 5939038 10700968
-3 0 9 33 97 251 572 1169 2225 4151 8002 16439 35742 79716 176825 382628 800814 1618431 3165604 6019862 11195332
4 0 7 46 145 348 729 1406 2558 4465 7617 12973 22495 40135 73515 136611 253832 465974 838627 1473720 2525005
10 16 48 135 331 736 1523 2980 5588 10168 18138 31917 55490 95121 160255 265025 432002 702892 1167508 2035271 3796881
12 22 45 91 170 292 467 705 1016 1410 1897 2487 3190 4016 4975 6077 7332 8750 10341 12115 14082
13 37 88 189 374 698 1256 2211 3831 6535 10948 17965 28824 45188 69236 103763 152289 219177 309760 430477 589018
7 17 47 117 250 482 894 1672 3215 6352 12823 26378 55242 117423 251606 538455 1140399 2372867 4827031 9574107 18494978
3 1 15 65 184 429 900 1785 3458 6656 12742 24025 44085 78138 133903 224663 379068 668010 1269397 2608855 5639543
16 37 81 160 285 459 669 874 987 861 317 -699 -1668 -491 8487 38015 113858 284205 634001 1306038 2531103
-4 2 20 52 98 160 247 392 713 1585 4054 10739 27675 67933 158575 353853 760018 1581404 3204672 6349758 12334232
-1 4 29 83 175 310 482 666 811 836 631 65 -997 -2666 -4988 -7888 -11097 -14060 -15823 -14897 -9097
11 3 5 38 144 400 931 1927 3676 6643 11671 20474 36752 68495 132356 261344 517477 1011375 1930965 3581378 6437564
7 13 11 -5 -44 -122 -267 -529 -1008 -1926 -3786 -7682 -15849 -32571 -65598 -128260 -242507 -443149 -783619 -1343635 -2239194
7 17 28 38 46 64 146 445 1322 3551 8691 19730 42147 85586 166391 311313 562769 986109 1679430 2786566 4513980
-1 -3 -3 18 98 297 705 1469 2868 5500 10719 21596 44905 94995 200967 419430 856405 1703875 3299289 6220314 11432620
10 21 31 39 42 50 111 340 960 2404 5607 12748 28896 65283 145283 314631 657982 1324599 2566783 4795629 8659822
22 52 93 136 166 172 169 228 518 1387 3547 8487 19351 42783 92883 199886 429291 923343 1988338 4271887 9109773
6 19 50 128 296 621 1214 2264 4104 7347 13153 23711 43040 78227 141225 251327 438410 747003 1241172 2010130 3174368
27 40 61 109 218 446 888 1688 3042 5195 8474 13480 21699 36999 68771 137861 286941 597594 1217155 2399271 4563232
1 1 9 45 149 390 884 1848 3733 7511 15257 31299 64465 132441 270169 545965 1092447 2166045 4261848 8335320 16228737
8 14 28 60 122 222 362 554 879 1628 3577 8460 19725 43727 91744 183878 357618 687726 1332134 2633940 5337380
15 33 76 157 289 487 770 1163 1699 2421 3384 4657 6325 8491 11278 14831 19319 24937 31908 40485 50953
12 15 19 24 30 37 45 54 64 75 87 100 114 129 145 162 180 199 219 240 262
8 14 18 20 20 18 14 8 0 -10 -22 -36 -52 -70 -90 -112 -136 -162 -190 -220 -252
16 35 76 160 321 614 1137 2079 3820 7134 13575 26154 50448 96364 181061 334404 610682 1116019 2073512 3974170 7912661
10 16 33 81 202 468 989 1921 3474 5920 9601 14937 22434 32692 46413 64409 87610 117072 153985 199681 255642
0 5 22 65 166 381 796 1545 2861 5190 9418 17316 32438 61971 120515 237571 471770 936751 1846290 3590031 6856250
15 31 70 143 264 464 816 1468 2672 4784 8190 13077 18908 23394 20795 -117 -56226 -160890 -285161 -241228 616733
25 46 87 158 270 435 666 977 1383 1900 2545 3336 4292 5433 6780 8355 10181 12282 14683 17410 20490
7 12 14 17 33 86 216 483 971 1792 3090 5045 7877 11850 17276 24519 33999 46196 61654 80985 104873
17 34 59 87 109 113 90 45 13 80 409 1271 3081 6439 12176 21405 35577 56542 86615 128647 186101
6 2 4 18 55 135 295 615 1295 2853 6573 15415 35706 80060 172122 353890 696532 1315772 2393058 4203832 7154281
5 27 61 107 173 295 576 1251 2785 6011 12315 23875 43961 77303 130534 212715 335949 516091 773561 1134267 1630645
-4 6 30 77 174 394 898 1992 4202 8378 15864 28840 51109 89988 160826 297555 576687 1170497 2461997 5287506 11436011
19 29 42 67 128 278 622 1356 2837 5707 11102 20985 38650 69452 121826 208666 349143 571049 913762 1431935 2200020
13 20 31 47 78 149 296 559 995 1751 3259 6665 14736 33800 77916 177663 397977 872742 1868847 3899763 7919106
8 9 17 49 146 403 1025 2422 5364 11231 22423 43076 80439 147757 270571 500538 945172 1829973 3623933 7284551 14720468
2 10 36 88 184 367 738 1521 3174 6560 13192 25566 47596 85165 146806 244527 394794 619686 948236 1417972 2076672
7 9 19 54 156 407 944 1974 3789 6781 11457 18454 28554 42699 62006 87782 121539 165009 220159 289206 374632
19 25 44 104 246 534 1086 2133 4121 7891 15016 28463 53909 102302 194657 370687 703880 1327497 2479627 4582648 8386169
10 27 56 98 169 326 706 1589 3512 7490 15452 31086 61419 119649 230010 435803 812182 1485861 2664624 4680394 8050669
24 36 54 86 158 344 824 1980 4538 9762 19715 37647 68701 121438 211330 368607 655040 1197912 2255256 4335292 8405940
19 36 65 121 227 415 741 1330 2466 4739 9257 17927 33806 61522 107767 181870 296469 468318 719289 1077661 1579829
18 29 53 112 244 509 995 1824 3158 5205 8225 12536 18520 26629 37391 51416 69402 92141 120525 155552 198332
3 3 -3 -10 -2 66 312 1013 2762 6755 15327 32960 68145 136698 267400 511131 954957 1742843 3104713 5395338 9143848
0 -6 -17 -33 -56 -99 -202 -457 -1044 -2280 -4683 -9053 -16572 -28925 -48444 -78277 -122584 -186762 -277701 -404073 -576656
11 29 66 135 257 461 784 1271 1975 2957 4286 6039 8301 11165 14732 19111 24419 30781 38330 47207 57561
-3 9 35 75 129 197 279 375 485 609 747 899 1065 1245 1439 1647 1869 2105 2355 2619 2897
9 27 65 147 323 691 1425 2814 5321 9667 16928 28607 46636 73351 111824 167814 256493 420851 775942 1607580 3584263
4 8 10 19 61 186 484 1120 2398 4864 9458 17725 32095 56242 95532 157570 252856 395560 604426 903815 1324897
-6 -8 -12 -22 -42 -76 -128 -202 -302 -432 -596 -798 -1042 -1332 -1672 -2066 -2518 -3032 -3612 -4262 -4986
7 20 40 67 101 142 190 245 307 376 452 535 625 722 826 937 1055 1180 1312 1451 1597
5 4 2 -2 0 49 265 924 2575 6206 13468 26966 50626 90147 153547 251812 399657 616408 927014 1363198 1964756
4 14 31 64 151 385 951 2174 4578 8956 16451 28648 47677 76327 118171 177702 260480 373290 524311 723296 981763
17 42 80 131 195 272 362 465 581 710 852 1007 1175 1356 1550 1757 1977 2210 2456 2715 2987
15 17 30 58 109 201 374 732 1555 3552 8381 19645 44683 97606 204169 409205 787455 1458683 2607936 4511660 7570073
17 39 68 104 153 234 398 770 1625 3509 7416 15032 29057 53616 94770 161138 264641 421379 652652 986136 1457225
6 25 54 93 142 201 270 349 438 537 646 765 894 1033 1182 1341 1510 1689 1878 2077 2286
8 11 12 9 0 -15 -29 -17 102 582 2154 6742 19061 50041 123858 291715 657584 1424081 2970676 5981683 11648018
0 16 51 118 249 516 1071 2216 4514 8952 17167 31746 56611 97500 162555 263028 414116 635936 954651 1403758 2025549
24 50 85 122 148 140 61 -144 -552 -1266 -2419 -4178 -6748 -10376 -15355 -22028 -30792 -42102 -56475 -74494 -96812
-3 -13 -25 -36 -43 -43 -33 -10 29 87 167 272 405 569 767 1002 1277 1595 1959 2372 2837
17 15 20 49 125 274 519 871 1317 1805 2226 2393 2017 680 -2195 -7377 -15863 -28917 -48112 -75375 -113035
13 30 70 143 259 428 660 965 1353 1834 2418 3115 3935 4888 5984 7233 8645 10230 11998 13959 16123
30 59 102 159 230 315 414 527 654 795 950 1119 1302 1499 1710 1935 2174 2427 2694 2975 3270
17 15 8 8 34 120 347 921 2330 5624 12873 27869 57149 111427 207534 370976 639231 1065917 1725974 2722014 4192004
22 33 45 73 138 264 479 817 1317 2033 3121 5170 10109 23266 57493 140718 328856 726721 1520445 3025943 5759180
5 -1 -2 9 51 179 535 1441 3551 8089 17235 34811 67612 128095 239789 447893 839405 1580430 2984454 5637286 10627953
7 9 15 30 66 148 324 676 1332 2498 4585 8635 17523 38967 92468 224368 538997 1260591 2852235 6237536 13202672
7 4 11 45 130 305 649 1341 2787 5864 12354 25671 52022 102187 194146 356809 635095 1096524 1839285 3001361 4769650
24 34 57 101 165 242 329 448 692 1325 2985 7064 16369 36203 76045 152053 290664 533620 944809 1619375 2695621
5 25 69 163 360 747 1447 2616 4435 7097 10789 15669 21838 29307 37959 47506 57441 66985 75029 80071 80148
8 27 56 101 185 370 787 1679 3481 6991 13738 26751 52118 102058 200799 395477 775694 1507483 2890452 5451089 10090927
23 36 62 116 223 433 851 1687 3331 6458 12168 22166 38987 66271 109093 174353 271231 411712 611186 889128 1269863
14 38 80 153 274 463 757 1267 2325 4795 10662 24076 53144 113021 231473 459544 893163 1718134 3303969 6396991 12510417
11 7 12 46 138 326 660 1206 2048 3286 5036 7476 11107 17734 33479 76942 199441 531941 1388040 3479538 8349120
7 1 5 39 140 373 842 1701 3165 5521 9139 14483 22122 32741 47152 66305 91299 123393 164017 214783 277496
2 8 20 40 63 87 140 331 938 2549 6267 13972 28597 54316 96455 160817 251954 369718 503174 620656 654387
16 33 59 95 160 306 633 1304 2560 4735 8271 13733 21824 33400 49485 71286 100208 137869 186115 247035 322976
22 31 42 55 70 87 106 127 150 175 202 231 262 295 330 367 406 447 490 535 582
24 44 71 101 142 227 434 918 1968 4127 8473 17291 35635 74789 159543 342727 732900 1544882 3188479 6418965 12584504
-3 7 35 102 237 469 819 1311 2036 3318 6046 12251 26022 54870 111664 217278 404103 720593 1237029 2052700 3304715
16 29 56 100 167 279 509 1058 2413 5660 13089 29350 63674 134217 276721 561971 1130944 2264777 4521540 8997922 17818875
9 22 35 40 40 62 171 497 1303 3144 7190 15795 33367 67520 130411 240269 423928 723949 1221549 2105304 3858295
19 32 65 126 222 376 666 1293 2694 5737 12071 24758 49389 95985 182110 337779 612933 1088478 1892149 3220766 5370800
9 16 30 46 60 74 100 163 308 622 1304 2883 6839 17182 44060 111284 270875 629468 1393778 2945484 5959974
21 35 66 138 292 605 1232 2488 5006 10047 20124 40281 80723 162163 326458 657183 1318227 2623971 5163054 10010380 19080462
12 18 47 119 263 528 1003 1849 3355 6042 10851 19460 34787 61785 108811 190335 332858 588110 1061583 1971129 3760859`
