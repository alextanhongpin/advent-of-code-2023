// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"maps"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solve(example)) // 19114 167409079868000
	fmt.Println(solve(input))   // 434147 136146366355609
}

func solve(s string) (int, int) {
	parts := strings.Split(s, "\n\n")

	var coords []map[string]int
	for _, c := range strings.Split(parts[1], "\n") {
		coords = append(coords, parseCoords(c))
	}

	rules := make(map[string]Rule)
	for _, rule := range strings.Split(parts[0], "\n") {
		r := parseRule(rule)
		rules[r.name] = r
	}

	var result []map[string]int
	for _, c := range coords {
		if eval(rules, c, "in") {
			result = append(result, c)
		}
	}

	var part1 int
	for _, r := range result {
		for _, v := range r {
			part1 += v
		}
	}

	counts := count(rules, map[string][2]int{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}, "in")

	var part2 int
	for _, m := range counts {
		mul := 1
		for _, v := range m {
			// The plus one is because the range is inclusive.
			mul *= v[1] - v[0] + 1
		}
		part2 += mul
	}

	return part1, part2
}

func eval(rules map[string]Rule, coord map[string]int, start string) bool {
	if start == "A" {
		return true
	}
	if start == "R" {
		return false
	}

	rule := rules[start]
	for _, r := range rule.conds {
		s, ok := r.Eval(coord)
		if ok {
			return eval(rules, coord, s)
		}
	}

	return eval(rules, coord, rule.last)
}

type Range struct {
	min int
	max int
}

func (r Range) Intersect(r2 Range) Range {
	return Range{max(r.min, r2.min), min(r.max, r2.max)}
}

func count(rules map[string]Rule, coord map[string][2]int, start string) []map[string][2]int {
	if start == "A" {
		return []map[string][2]int{coord}
	}
	if start == "R" {
		return nil
	}

	last := maps.Clone(coord)
	var success []map[string][2]int

	rule := rules[start]
	for _, r := range rule.conds {
		when, then, elif := r.count(last)
		// Separate into success and failure.
		// If failed, we just append the condition to the last.
		success = append(success, count(rules, when, then)...)
		last = elif
	}

	// Evaluate the last rule.
	return append(success, count(rules, last, rule.last)...)
}

type Cond struct {
	val string
}

func (c Cond) count(m map[string][2]int) (map[string][2]int, string, map[string][2]int) {
	cond, ret, ok := strings.Cut(c.val, ":")
	if !ok {
		panic("invalid")
	}

	k, op, v := cond[:1], cond[1], cond[2:]
	n := toInt(v)

	one := maps.Clone(m)
	two := maps.Clone(m)

	// Just set the low and high value.
	if op == '<' {
		one[k] = [2]int{one[k][0], n - 1}
		two[k] = [2]int{n, two[k][1]}
		return one, ret, two
	}

	one[k] = [2]int{n + 1, one[k][1]}
	two[k] = [2]int{two[k][0], n}
	return one, ret, two
}

func (c Cond) Eval(m map[string]int) (string, bool) {
	cond, ret, ok := strings.Cut(c.val, ":")
	if !ok {
		panic("invalid")
	}

	a, op, b := cond[:1], cond[1], cond[2:]
	n := toInt(b)

	switch op {
	case '<':
		return ret, m[a] < n
	case '>':
		return ret, m[a] > n
	}

	return "", false
}

type Rule struct {
	name  string
	conds []Cond
	last  string
}

func parseRule(s string) Rule {
	i := strings.Index(s, "{")
	wf, rules := s[:i], s[i+1:len(s)-1]

	conds := strings.Split(rules, ",")

	r := Rule{name: wf, last: conds[len(conds)-1]}
	for _, rule := range conds[:len(conds)-1] {
		r.conds = append(r.conds, Cond{rule})
	}

	return r
}

func parseCoords(in string) map[string]int {
	var x, m, a, s int
	fmt.Fscanf(strings.NewReader(in), `{x=%d,m=%d,a=%d,s=%d}`, &x, &m, &a, &s)

	r := make(map[string]int)
	r["x"] = x
	r["m"] = m
	r["a"] = a
	r["s"] = s
	return r
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var example = `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`
var input = `xkv{a>2504:A,x<1530:R,A}
qz{a<1970:A,x>3109:A,a>2070:A,A}
mxz{s>699:R,s>378:R,A}
lhf{m<1064:drs,A}
cnm{a>3392:hml,x<3423:nd,x<3797:vjg,nfn}
tqg{a<247:A,x>3063:A,A}
vsk{m<3711:R,x<3262:A,m<3859:A,A}
ffv{a>534:R,m>2824:ptx,qkp}
sk{x<2768:R,m>3533:A,m>3475:R,R}
nd{s>3057:R,x<3063:lf,m>1124:qg,kkk}
vdz{x<3453:R,rg}
lfm{x>3090:sn,a<3647:zf,jpx}
dpg{x<2909:tx,x>3041:zxr,m<2728:rhk,R}
bfn{m<3889:A,a<2136:A,s>1544:R,A}
hg{m>2435:R,A}
srq{a<2592:R,a<2942:A,R}
nh{a<695:R,a<862:R,s<2863:R,A}
sf{m>867:R,m>545:A,A}
lfv{m<553:A,m<622:A,R}
brl{s>1779:R,x>618:R,s<1558:R,R}
dpz{a>1553:mb,ld}
xfd{x>1604:kmd,lds}
sj{m<2454:R,xns}
mkl{a>937:A,m>2342:A,m<2200:A,A}
cll{a<3546:A,a>3717:R,R}
rl{a<3156:A,R}
vzj{x>1388:R,m<2520:R,m>2595:R,R}
jqd{m<1237:A,s<575:R,s<822:mxs,A}
bcv{m<2674:xn,x>1635:pbj,m>2789:ffv,hh}
jm{x>3465:bdh,kc}
hjk{s>2982:R,R}
xxt{m<3838:lft,sg}
hh{m<2726:fsd,m>2760:R,a>513:A,A}
pqj{m>727:A,x<2610:A,x<2698:vj,R}
vjs{s<2593:A,a<3380:R,A}
mkk{s>1188:ktb,a<1065:gl,m>2676:tfs,pqx}
td{m<2540:R,m>2852:R,m>2695:R,A}
lv{x>2575:R,m<699:A,A}
qlg{x<2593:A,R}
skp{m>543:A,x<2918:mp,s>1383:R,A}
zpc{x<1132:R,m<2752:R,m>2820:A,R}
jxm{s<527:A,m<790:R,R}
nhm{a>546:kh,m<344:qcp,A}
vq{s<2024:gt,ts}
xl{x<3081:R,m<3856:A,x>3675:A,R}
ldp{m<973:phf,a<1111:jgb,df}
slv{m>3643:zv,vxb}
zjc{a>658:R,rz}
xx{a<1631:R,A}
mth{s>256:A,s<123:A,R}
ck{m<1689:R,A}
vb{s>1286:A,s>1148:R,a>2321:R,R}
hzz{s>2923:A,qz}
kfr{s<2749:A,m<1366:A,A}
xr{a<1884:nr,m<3307:kq,zb}
ft{x<2965:R,A}
drs{s>1186:R,x>3784:A,a<3137:A,R}
fng{m<3631:R,s<1005:R,x<2207:A,R}
ktb{s<1718:hps,s<2041:R,s<2216:A,krn}
vf{m<489:R,s<2772:xd,s>3022:R,npx}
pfj{s<1850:nkp,x<2932:nfd,A}
svn{a<2529:dbb,a>2578:lc,a>2558:vsk,A}
xjl{a>1147:A,m<3416:R,m>3684:A,A}
gd{m>1151:R,m<747:mz,s>3664:R,A}
szt{m<741:A,R}
jhz{s<980:A,a>3272:ckv,x>1933:cl,A}
cx{a>1197:A,a<1004:mkl,bx}
qt{s>2057:R,s<1706:R,m<2851:A,ddd}
sg{m<3903:R,x<3774:R,A}
lxr{m>2278:sv,R}
vgv{a>3352:js,x>983:jk,s<2648:pq,thj}
dr{s<2247:kgg,xqc}
hm{s>2064:rgq,s<959:lrn,s>1473:sfv,A}
qnz{x>1926:A,x<1370:R,m>2633:R,A}
nnb{a>3159:R,x<2873:srq,a<2778:A,R}
dxj{a>1283:R,x<2996:A,m>1289:R,A}
xs{x<2067:tl,hj}
xzd{a>2939:dn,x<1759:vk,x<2245:prc,vd}
jtg{s>2244:R,R}
lqh{x>3627:A,a>599:qfx,x<3478:fhq,tv}
mz{a<1184:R,A}
lqv{m<2699:A,a<1655:R,A}
xng{x>399:A,s<1539:bn,m<3597:R,dp}
tkc{x>2151:A,x<2100:A,s<622:A,A}
pbj{m<2744:A,A}
ngg{s<909:R,A}
gvd{m>2546:R,x>3024:A,s<3693:R,R}
vk{a>2480:lh,x<1478:A,rph}
ztf{m>3809:A,s>683:A,A}
vkg{s>2160:A,x>457:R,A}
skc{m<3382:A,R}
frg{s>1178:A,md}
jl{a>1808:R,x<926:R,R}
kq{s<1440:xs,x>2476:jf,knk}
gb{m<230:A,s>2201:A,R}
xxb{x<1200:R,x<2033:xkv,A}
vv{a>1280:ffj,hb}
fpc{a<1017:bh,x>353:hrr,ns}
cpf{x>3065:cj,m>834:tm,vf}
klm{s<458:A,A}
mj{x<2645:A,m>2520:R,x<2760:A,R}
xbz{m<2351:A,s>1323:R,a>801:A,A}
llb{m>3555:R,R}
lxf{x<455:A,R}
vj{a<2492:A,m<383:A,R}
gp{s<2556:zjk,m<419:dbv,R}
ccn{a<1268:thv,vqr}
sn{s<2445:R,s<3147:kcj,glg}
cjx{m<2349:A,m>2479:R,s>538:A,R}
bdh{x<3524:A,a>764:R,A}
tlj{m<1192:pmp,m<1622:kf,a>2944:qlg,R}
tgs{m>3553:dmx,m<3129:R,x<887:lxf,A}
db{a>1012:R,s<1705:A,R}
pj{a<1295:A,A}
hvn{x>121:A,m<3599:R,A}
tfs{m<2774:A,gs}
jn{x>3163:ccn,x<2943:vv,cpf}
lq{x<2782:A,s<3109:R,x>2989:R,R}
hj{s>667:gng,a<3079:pf,m<2650:vl,qjd}
vh{a<690:nzz,skc}
fbq{x<3401:fp,a<385:hfh,a<428:R,A}
bl{x<2519:mr,s>1935:xqn,cb}
ldq{a>3030:cnm,hd}
mqg{s<3488:A,a<2368:A,gvd}
jph{a<403:A,R}
krs{a<3217:dnn,nnc}
hps{s>1533:R,a<1022:R,A}
nz{s>2249:A,ckp}
mhq{a>1371:tz,m>2680:R,pj}
flf{m<3401:R,R}
zk{m>2438:R,m<2379:fz,R}
rvv{m<3843:R,x<2952:R,s<536:R,A}
rbx{x>3829:R,s>3393:A,m>697:A,R}
jgb{x<2114:lqf,s>2913:A,A}
vg{a<319:R,a<405:R,x>3070:A,A}
zlf{x<2457:A,x<2489:R,A}
ds{x>260:A,x<159:A,m<2784:A,A}
px{s<2272:A,m>904:A,a>1418:R,A}
ldj{x>704:R,A}
krj{x>887:R,A}
vp{m>2403:pk,A}
pb{m>729:R,m>335:A,R}
xns{x>3814:R,a<2631:R,m>2934:R,R}
nm{a<2596:lmc,a<2803:qt,a<2905:qc,lbk}
hf{a<652:fkm,mjh}
lds{m<3701:R,m<3876:jb,mn}
mg{m<2739:R,A}
hds{s<2792:ghd,svn}
brc{x>3034:R,R}
hkk{a>979:A,m>2258:R,R}
nbb{m>2423:R,R}
kgg{x<2344:zn,x<3351:vh,m<3538:gr,xxt}
jxh{a>1001:A,m<2137:R,a<375:lq,skf}
nmz{a<1092:A,A}
qfn{s<1082:A,x>1366:R,m>2732:R,R}
jls{m>2690:dqd,x<1106:fpg,m>2623:A,A}
qks{m>1240:R,s>2857:A,m>691:hx,R}
dbb{x<3173:R,x<3499:R,x<3725:R,R}
hb{a>1088:R,a<1044:A,qnq}
tl{m<2679:kdf,cc}
jk{x<1761:vmc,x>2006:psp,a>3258:rq,R}
kmd{a>2412:A,A}
jg{m>3738:vgv,x>1473:jgh,a<3310:xtk,dhr}
mxs{s>724:A,s>650:A,a>2016:R,A}
qmg{m<1042:R,A}
vzc{a>1174:A,R}
cxt{x>2937:zc,m>2802:gkf,a>2348:R,mj}
ts{a<1237:R,s<2852:A,A}
dvx{s>1425:ftg,m>1578:vb,lgn}
jb{x>1233:R,R}
sfv{x>1239:A,s<1729:R,R}
rpf{s>3314:blv,s<2936:ttz,s>3160:xx,lqv}
bf{x>2831:R,s<2534:R,m>2358:R,A}
lb{x>2846:R,A}
fm{m<353:dds,x<3303:A,a<2482:A,jr}
fc{m>1200:dvx,x<2763:pqj,a>2421:skp,mk}
vvr{a<1315:krl,s<922:A,s>1463:R,R}
vx{x<1200:brl,x<1637:xzs,kz}
gxg{s<3604:R,s<3783:A,x>3794:R,R}
thv{s>2646:nmz,btc}
hxj{x>3637:zjc,jm}
vpb{m<2328:sb,x>3163:hxj,a<906:hpn,mtf}
hbv{m<2899:R,x<713:R,s>515:R,R}
nsv{s>831:A,m>3641:R,m>3615:A,R}
sdh{x>2642:A,m>1701:R,a<3462:A,A}
qkq{x>3166:A,m>844:A,s>2950:R,R}
bsm{m>2129:A,R}
vqc{s>2017:dbx,x>3160:rgv,srz}
pr{a<2417:ghf,a>2738:sdq,x<3580:vdz,sj}
bzz{s<2323:mkk,a<1189:bcv,pqs}
hd{a<2258:hzz,a>2670:qks,gdm}
cg{a<2596:vvr,lhf}
fz{m>2314:A,a>1114:A,x>963:R,A}
kxv{x<1929:A,R}
vl{a<3502:A,s<413:mth,x<3274:R,cjx}
dg{m>2538:R,m>2294:A,a>2151:A,R}
dbv{a>777:R,A}
rnm{m<2279:gfg,x>1238:vp,x>432:zk,vzc}
lx{a<3286:A,s>1162:A,m<2544:A,A}
fkm{a<340:tlk,a<471:fbq,rtr}
jzz{m<3464:A,R}
qm{m<3931:A,s<1999:R,R}
lp{x>2718:lr,x<2467:ppz,a<2199:R,dsl}
gdz{m<1186:A,a>71:R,s>2098:R,R}
dc{s<1380:A,s<2762:R,R}
cpt{a<3214:A,s>1400:R,x<173:R,R}
dpc{s<2578:A,R}
kz{s<1798:R,R}
fl{s>1758:A,A}
jnz{s<452:A,s<622:A,R}
dhb{a>1730:hk,x>580:lhl,fpc}
blr{x<3554:A,x>3646:A,a>1220:A,A}
kr{s>3219:R,s<3074:R,R}
fx{m>2195:A,s<562:R,x<1281:R,A}
dl{a>1635:jqd,m<793:nhm,ffn}
nnk{m>781:R,x<3583:R,m>475:R,A}
lgn{a<2362:R,a<3061:R,R}
qvf{s>2815:A,R}
sx{m>405:bm,vrg}
fhq{m>1205:R,R}
phf{m>420:kxv,hjk}
kc{x<3336:R,s<1853:R,R}
ckp{x>1689:A,R}
pk{x<1838:R,m<2492:R,A}
vmc{m>3872:R,m>3827:A,R}
lrn{s>618:R,m<3453:R,m>3553:A,A}
krn{a<1053:A,x<874:A,R}
tx{x<2767:R,m<2729:A,x>2839:R,R}
rhk{s<1926:A,A}
gr{x>3682:xrl,x>3552:R,A}
vd{x<2390:R,a<2338:pzq,zlf}
sv{m<2417:A,m>2503:R,x<2108:R,A}
tk{s<1028:A,rt}
khl{m>809:A,m>386:R,R}
js{s<2489:hfd,s<3162:R,R}
rph{m<1329:A,A}
tlk{x>3223:nnk,sf}
xjr{s>2377:dm,s>2160:pg,kvc}
qcp{a>197:A,m<200:R,x<2621:A,A}
lqf{m>1653:A,A}
fp{x<2917:R,R}
ln{s>1308:A,A}
dt{s>1989:xxf,s>1968:R,A}
gkv{s>1015:fc,x>2705:vbb,a<2584:dl,cmd}
qx{x<1945:tk,s>888:frg,x<2319:jtd,cs}
pqs{a<1565:mhq,a<1747:rpf,s>3108:jls,jl}
kvc{s>2021:bd,dt}
cd{s<2526:jd,s>3302:rnm,zd}
dm{s>2764:tng,m>1176:kb,a>527:gp,qrc}
ql{s<2564:tbf,x<2090:R,jzz}
hgc{m<2261:R,x<2854:R,a>795:ft,lg}
dh{s>1881:A,m<3578:R,x>285:xf,nsv}
tng{m<1292:qkq,s>2941:R,a<377:tqg,nh}
bv{s>498:R,m<3410:tzx,sk}
nj{x<3582:rp,cg}
ghd{m<3751:R,xl}
rgq{s<2895:A,x<1522:R,a>3921:A,A}
cgt{m>3497:sbs,s<2628:R,s>2826:R,A}
bqs{m<3785:R,x>2872:A,x>2772:R,A}
bgh{s<3855:R,A}
tj{a>3425:R,A}
dp{x>266:R,R}
src{a>2635:vqc,a<2363:ksc,s<2166:bs,hds}
phk{m<2309:R,A}
zxr{m<2787:A,a>593:A,x>3103:A,A}
qjk{m<2326:qph,a<300:nbb,hg}
jjr{s<1690:R,A}
dsl{m<3686:R,R}
gf{m>702:A,A}
cn{a>982:R,A}
nnc{s>2522:R,m<3498:R,R}
xn{s<3346:R,x>918:qnz,a>407:bz,A}
dhr{x>554:tt,dh}
knk{s<2786:gn,m>2849:gbm,tlx}
gdm{x<3428:A,R}
dbx{s>3277:mhm,dj}
dnn{a<3186:A,s<2228:R,A}
qnq{m<1274:R,a>1073:A,R}
cv{a<149:A,mg}
gx{x>2639:R,lv}
vdh{x>1520:R,x>532:A,s<2641:A,R}
tz{s>2935:R,x<1274:R,m<2670:R,R}
hlg{m<3754:xng,x<545:gdg,fpn}
qfx{s<2301:R,m<1147:A,a>818:A,R}
hbz{x>3913:R,s<3651:R,a<3166:R,R}
gn{a>3212:qs,s<2067:vx,mvs}
bxc{x<3197:R,x>3560:A,R}
zjk{a>754:R,s<2460:R,A}
gbm{a<2967:xxb,nf}
pf{s<249:R,x>2973:tdl,x>2386:nl,A}
pmp{x>2595:R,s>323:A,m>478:A,A}
xrl{m<3211:R,R}
gq{s<3648:lm,m>923:blr,bgh}
pcp{x>1217:A,m>3432:A,x<700:A,R}
tzx{x<2781:A,A}
lmc{m<2606:A,A}
lbk{s<2122:zpx,R}
ptx{a<261:R,m<2848:R,R}
gdg{s<2157:pll,x>318:jtz,x>170:R,bp}
jvd{a<355:R,m>2337:R,m<2226:R,A}
btc{x>3666:A,a>1109:R,A}
cl{x>2003:A,s>1440:A,m<3540:A,R}
bs{x>3332:rgj,s>1026:kvv,m>3607:zkg,bv}
ld{m<2732:R,m<2787:R,R}
pv{m>3665:A,x>3530:A,R}
rq{x<1898:A,m>3909:R,A}
fg{s<1737:qx,a<1670:ldp,xzd}
qc{s<2026:td,m<2495:dpc,s<2480:A,R}
hz{x>3104:A,x<3091:R,s>3021:A,R}
zg{a<2057:R,x<2147:A,s<2713:A,R}
bx{m>2248:R,a<1117:R,a>1157:A,R}
mb{x<2754:R,x>2970:R,m>2768:A,R}
fb{a<1061:A,A}
rtr{a<591:hgl,x>3132:A,R}
kv{a<3193:R,x>890:A,R}
cb{x>3061:nj,gkv}
rh{a>1126:A,x>701:R,s>1039:R,A}
rs{m>3465:fng,m>3406:R,s<991:tkc,A}
jrs{s>3348:xjl,a<891:A,A}
ddd{x<3632:A,R}
xxf{x<3431:A,x>3715:R,s>2006:R,A}
prc{x<1961:R,a<2400:zg,A}
vql{m>820:A,A}
pq{a>3283:jjr,A}
ksc{x<3116:lp,pv}
dn{s>3219:pb,s<2596:R,m<1274:A,cll}
vqr{m<1322:R,R}
zx{m<2382:hkk,rh}
jkc{x<3561:A,x>3832:dc,x<3677:R,bsm}
ffj{a>1476:A,a<1381:hgz,s<2497:R,A}
nf{m<3057:R,A}
st{a>3782:R,x<2780:A,R}
zhg{a<442:R,s>2291:A,m<777:R,ntm}
mtf{m>2586:dpz,m>2477:pfj,fl}
qrc{m>533:vg,R}
jtd{x>2147:R,m<901:A,a>1562:zr,jnz}
khv{x>858:xfd,hlg}
kkk{s>2469:A,A}
cs{a<1678:jxm,x>2445:R,R}
xqn{a>1577:ldq,s>3254:hf,a<962:xjr,jn}
rp{m>925:ktn,s>733:fm,sx}
tlx{a>2763:phk,kr}
qr{s<541:tj,s>742:R,m<1196:gf,sdh}
xv{x<620:A,s<778:A,m>3490:A,A}
ghx{x>1044:knn,qm}
gkf{m>2982:R,x>2719:A,x>2592:R,A}
xlx{x>3224:R,R}
hk{s<2134:qnh,m>751:kv,fjr}
tkr{x<3564:R,R}
zv{a>3765:cbm,m>3825:ghx,zmg}
xzs{a<2714:A,m<2663:R,s<1777:A,R}
cqf{m>2360:A,m<2170:R,x<3630:R,R}
qkp{s>3185:A,a<322:R,m>2807:A,R}
tv{x>3566:R,R}
ktn{s<1031:lgp,a<1510:cn,hcx}
tm{s>2552:dxj,A}
lc{x>3433:A,R}
kvv{s<1419:A,R}
qjd{a<3493:R,m<3085:A,A}
vbb{a>2256:nnb,lb}
qnh{m>1065:A,cf}
zp{s>2503:R,A}
kb{m>1619:bxc,brc}
in{m<2023:bl,xr}
vrg{x>3360:R,s>364:R,x<3192:R,R}
xhp{m>3517:hvn,a>3235:A,a>3194:cpt,flf}
pll{m>3918:R,R}
ckv{a>3367:R,A}
vqb{m<1154:R,m>1540:A,R}
hl{m>1440:R,a<1338:R,A}
kdf{m<2326:fx,a>2603:A,R}
zpk{m<1233:R,x<2632:A,a<661:R,A}
df{m>1379:ck,x>1765:nq,A}
hcx{x>3377:A,x>3180:A,a<2921:R,A}
kh{a>1214:R,a>973:R,A}
ghf{m<2871:dg,s<3381:R,R}
xf{s<1083:R,A}
npx{x>3013:A,A}
mjh{x<3451:gd,x>3699:kxb,gq}
nkp{x<2919:R,s<642:A,x>3028:A,A}
ffn{s>566:ckd,zpk}
gs{m>2826:R,x<1232:R,x>1796:R,A}
ntm{s<2244:R,x<2798:R,R}
bm{m<644:A,m<769:A,x>3408:R,A}
vr{s>1608:db,xbz}
bz{x<529:R,A}
zc{m>2772:R,A}
qv{a<3949:A,a<3967:R,R}
lm{m>693:R,a>1249:R,R}
nzz{a<430:A,x<2928:R,m>3593:A,A}
jpx{a>3871:qv,s<2638:A,st}
pzq{x>2434:A,s>3006:R,A}
zmg{m<3755:jtg,s>1412:A,A}
gk{m>2092:R,A}
rm{a<287:R,m>817:R,A}
rz{m<2523:R,x<3802:R,R}
bd{x<3348:pgs,a>412:R,a>184:rm,gdz}
jr{a<3459:A,a<3744:A,x<3432:R,R}
bn{m<3567:A,s<565:A,x<157:R,A}
qph{x<1929:A,s>2826:A,R}
mr{x>1335:fg,dhb}
fpg{m<2631:R,m>2670:A,s>3595:A,R}
nl{m<2777:R,a<2491:R,m>3033:R,A}
xqc{x<1847:tgs,s>2971:jrs,cgt}
krl{s>1051:R,a<562:A,s<531:R,R}
jf{a>3016:lfm,x<3356:nt,s<2887:nm,pr}
tbf{a<3320:A,a>3374:A,R}
dj{s>2519:R,x>3124:A,a>3435:R,R}
rkm{m<3789:A,R}
jtz{a>2325:R,x<447:A,x<490:A,R}
dmx{s<3222:A,R}
cbm{a>3853:A,x<1454:A,a>3795:A,khs}
zkg{x<2690:ztf,rvv}
jgh{x<1856:nz,s>1740:ql,x<2058:jhz,rs}
sb{x>3160:jkc,s>2434:jxh,m<2179:fb,hgc}
hfd{s>1479:R,R}
blv{s>3651:A,A}
rt{s>1306:A,A}
glg{m<2517:A,A}
skf{a<731:R,a<826:R,R}
bjm{x<2925:R,x>2972:R,a<1178:R,R}
vjg{a<3193:A,s<2765:A,R}
mn{s>2360:A,m>3934:R,R}
dds{x>3365:A,a<2607:R,a<3223:R,A}
qq{a<1623:R,A}
qg{m>1556:A,A}
rgj{m<3753:A,x<3748:A,ln}
ldc{m<695:A,krj}
gt{s>798:R,R}
hgz{s>2480:A,m<1250:R,R}
pgs{m<1105:R,A}
tt{s>1428:vjs,A}
thj{m<3885:nvr,A}
bh{x>271:A,R}
hml{s>3305:krt,m<797:xlx,x>3203:kfr,R}
tdl{s<491:A,a<2325:R,x>3404:A,A}
knn{x<1795:R,a>3644:R,R}
krt{a<3789:A,a>3924:A,x>3157:A,R}
sbs{x<2764:A,x<3450:R,A}
rg{m<2733:R,A}
jd{x<1436:zx,s<1164:klm,s>1855:lxr,vr}
hbr{s>826:R,A}
sdq{m<2615:cqf,x<3602:A,a>2909:gxg,A}
hgl{a<520:R,x>3189:A,x>2780:A,A}
nvr{m>3788:A,A}
hrr{x>447:px,m>1045:A,A}
nr{m>2861:dr,x>2517:vpb,m<2569:cd,bzz}
zpx{m<2756:R,R}
fjr{m<415:R,s<3163:qvf,lfv}
vn{a<321:A,m>916:A,m<610:R,szt}
cmd{a>3095:qr,s>650:jx,a<2861:gx,tlj}
zd{a>811:cx,x>851:qjk,bjp}
ftg{x<2737:A,R}
zr{a>2461:A,x<2036:A,a<1980:A,A}
zb{x>2298:src,a<3159:khv,a<3456:jg,slv}
kf{s<300:A,R}
lft{s>758:R,m<3680:A,A}
gfg{a>862:A,m>2178:R,a>300:R,gk}
mvs{m>2720:zp,s<2397:R,m>2373:vzj,vdh}
ttz{s>2636:A,a>1671:A,R}
khs{a>3780:A,R}
xtk{x>705:krs,x>307:vkg,xhp}
nfd{m<2529:R,R}
cj{s<2807:vqb,hz}
psp{s>2091:A,s<813:A,R}
cc{a<2910:hbv,nv}
mp{s>1454:A,m<274:A,m>431:R,A}
zxn{s<1411:A,a>1408:A,s<1938:A,R}
kxb{s>3573:qmg,rbx}
ppz{x>2375:A,R}
qs{x>1155:R,x>764:A,ds}
lr{m>3727:A,A}
gng{x<3343:R,a<2792:A,s<1002:tkr,lx}
zn{x<1440:xv,a<1159:A,m<3535:qq,zxn}
nv{x<1036:R,s<804:A,m<2949:R,R}
rgv{x<3541:llb,m<3745:R,m>3877:R,rkm}
lgp{m<1570:R,A}
mk{s<1468:bjm,R}
hx{x<3408:R,x<3748:R,a>2865:A,A}
hfh{x>3664:A,a>366:A,A}
nq{a>1413:A,A}
gl{s<587:R,s<808:R,s<970:zpc,qfn}
lh{m<769:R,A}
bjp{x<540:A,jvd}
cf{a>2587:R,s>758:A,R}
lhl{x<1051:ldc,a<866:vn,vq}
dqd{x>1354:A,x>512:A,m<2802:A,R}
pg{x>3309:lqh,zhg}
fsd{m<2701:A,A}
ns{m>884:hl,m<479:gb,s>2080:A,A}
kcj{x>3594:A,m>2843:A,m>2338:R,R}
mhm{m<3741:A,x<3194:R,m<3832:A,A}
vxb{a<3768:pcp,hm}
ckd{s<845:R,R}
lf{x>2800:A,A}
nfn{m>1010:A,s>3191:hbz,x>3866:A,rl}
xd{a<1222:A,x>3024:R,m<660:R,A}
md{s<1014:R,R}
pqx{a>1369:A,s<449:A,m<2606:hbr,ngg}
hpn{m<2654:jph,a>367:dpg,cv}
bp{a>2680:R,a>2177:A,A}
nt{s<2996:cxt,mqg}
zf{m<2571:bf,m>3027:A,s<3097:R,A}
jx{a>2831:vql,khl}
lg{a>435:A,a>273:R,A}
fpn{a>2344:A,s<2420:bfn,m>3914:ldj,A}
srz{m<3646:mxz,x<2707:A,bqs}

{x=1243,m=275,a=647,s=591}
{x=1563,m=724,a=801,s=2826}
{x=550,m=64,a=200,s=2487}
{x=3030,m=730,a=2108,s=414}
{x=2656,m=1297,a=440,s=213}
{x=415,m=297,a=330,s=799}
{x=1412,m=1705,a=1503,s=2893}
{x=998,m=1610,a=2121,s=810}
{x=507,m=1643,a=615,s=1575}
{x=387,m=936,a=2718,s=663}
{x=47,m=628,a=2908,s=2478}
{x=124,m=248,a=108,s=984}
{x=753,m=102,a=1968,s=265}
{x=3102,m=112,a=164,s=375}
{x=2790,m=832,a=840,s=1286}
{x=74,m=875,a=11,s=20}
{x=523,m=348,a=147,s=17}
{x=2139,m=1534,a=923,s=1168}
{x=3098,m=2409,a=956,s=44}
{x=813,m=1180,a=1119,s=1338}
{x=958,m=452,a=25,s=101}
{x=867,m=58,a=558,s=106}
{x=2475,m=116,a=876,s=1507}
{x=560,m=34,a=420,s=9}
{x=3474,m=489,a=332,s=537}
{x=220,m=2061,a=844,s=621}
{x=367,m=2406,a=3040,s=2115}
{x=468,m=97,a=1344,s=1531}
{x=307,m=2342,a=2535,s=430}
{x=760,m=166,a=3234,s=20}
{x=668,m=154,a=659,s=2947}
{x=801,m=440,a=295,s=876}
{x=183,m=1620,a=1482,s=1311}
{x=1924,m=439,a=2913,s=481}
{x=158,m=126,a=657,s=1985}
{x=1536,m=543,a=556,s=691}
{x=352,m=134,a=1757,s=43}
{x=1252,m=647,a=386,s=2010}
{x=166,m=430,a=1755,s=2037}
{x=1062,m=333,a=576,s=86}
{x=1667,m=171,a=1686,s=370}
{x=461,m=246,a=206,s=1334}
{x=160,m=977,a=1821,s=3078}
{x=1095,m=39,a=1490,s=3115}
{x=202,m=1946,a=448,s=439}
{x=1006,m=296,a=1038,s=8}
{x=33,m=1599,a=701,s=631}
{x=255,m=677,a=925,s=47}
{x=553,m=100,a=681,s=1987}
{x=327,m=140,a=928,s=153}
{x=2873,m=2219,a=749,s=72}
{x=141,m=1498,a=55,s=980}
{x=1565,m=645,a=722,s=1976}
{x=2181,m=1127,a=963,s=33}
{x=348,m=829,a=52,s=612}
{x=3397,m=1970,a=766,s=3682}
{x=2346,m=866,a=1608,s=3249}
{x=1252,m=3102,a=2114,s=1085}
{x=3139,m=1683,a=873,s=1568}
{x=1028,m=2253,a=82,s=801}
{x=2517,m=728,a=1701,s=71}
{x=2270,m=2530,a=436,s=329}
{x=215,m=203,a=265,s=862}
{x=653,m=184,a=3426,s=1560}
{x=1980,m=1749,a=581,s=319}
{x=373,m=2208,a=1135,s=35}
{x=790,m=1980,a=438,s=498}
{x=883,m=1740,a=259,s=915}
{x=200,m=358,a=1321,s=45}
{x=176,m=485,a=2125,s=2}
{x=38,m=914,a=408,s=1235}
{x=93,m=313,a=13,s=280}
{x=2595,m=1080,a=564,s=385}
{x=304,m=107,a=53,s=3167}
{x=1105,m=419,a=861,s=1235}
{x=875,m=1210,a=255,s=575}
{x=3230,m=124,a=16,s=1419}
{x=1630,m=939,a=647,s=160}
{x=116,m=86,a=3272,s=45}
{x=499,m=1248,a=263,s=80}
{x=45,m=328,a=325,s=1706}
{x=3792,m=1921,a=143,s=1945}
{x=1496,m=373,a=288,s=212}
{x=1979,m=55,a=880,s=449}
{x=846,m=387,a=2878,s=395}
{x=2458,m=906,a=3115,s=2373}
{x=617,m=116,a=1149,s=902}
{x=1289,m=1140,a=1016,s=752}
{x=3717,m=1357,a=809,s=3197}
{x=1101,m=1185,a=68,s=1556}
{x=39,m=1553,a=290,s=1956}
{x=1941,m=1953,a=370,s=469}
{x=542,m=393,a=310,s=205}
{x=723,m=36,a=1044,s=2478}
{x=526,m=1439,a=361,s=315}
{x=2004,m=596,a=1530,s=872}
{x=2852,m=984,a=2689,s=234}
{x=597,m=744,a=338,s=1888}
{x=622,m=320,a=804,s=462}
{x=729,m=1919,a=284,s=1584}
{x=42,m=1579,a=940,s=970}
{x=1602,m=1789,a=177,s=815}
{x=1144,m=11,a=217,s=511}
{x=2278,m=882,a=2492,s=88}
{x=455,m=1366,a=1237,s=1550}
{x=343,m=1062,a=788,s=1043}
{x=1276,m=2736,a=2418,s=289}
{x=1782,m=1774,a=455,s=155}
{x=570,m=1015,a=123,s=2307}
{x=157,m=1034,a=65,s=2512}
{x=223,m=3317,a=167,s=695}
{x=860,m=1377,a=1260,s=1202}
{x=929,m=297,a=13,s=2762}
{x=1407,m=315,a=62,s=3098}
{x=2905,m=1109,a=759,s=1762}
{x=99,m=1970,a=410,s=43}
{x=45,m=1553,a=2815,s=137}
{x=210,m=87,a=139,s=45}
{x=100,m=33,a=2596,s=23}
{x=3424,m=3113,a=4,s=2886}
{x=241,m=895,a=1443,s=267}
{x=1380,m=1592,a=3407,s=2345}
{x=295,m=532,a=1318,s=2327}
{x=46,m=1707,a=37,s=185}
{x=823,m=344,a=345,s=776}
{x=90,m=1805,a=3750,s=1161}
{x=857,m=59,a=452,s=1252}
{x=122,m=416,a=414,s=528}
{x=1672,m=1192,a=285,s=1333}
{x=582,m=147,a=3405,s=75}
{x=234,m=243,a=408,s=289}
{x=44,m=72,a=631,s=555}
{x=867,m=359,a=1830,s=1231}
{x=397,m=358,a=24,s=1169}
{x=259,m=951,a=1053,s=1909}
{x=540,m=1597,a=489,s=328}
{x=867,m=971,a=751,s=1768}
{x=121,m=2141,a=616,s=94}
{x=373,m=911,a=1576,s=1026}
{x=2103,m=38,a=209,s=708}
{x=1938,m=1508,a=918,s=2583}
{x=60,m=3003,a=1445,s=156}
{x=1001,m=2132,a=1182,s=600}
{x=482,m=1079,a=2152,s=1008}
{x=724,m=362,a=2459,s=2691}
{x=28,m=2660,a=286,s=1810}
{x=49,m=1106,a=1899,s=861}
{x=212,m=1780,a=1307,s=815}
{x=96,m=2581,a=1524,s=1463}
{x=10,m=212,a=2500,s=239}
{x=1019,m=779,a=920,s=476}
{x=16,m=743,a=135,s=1450}
{x=564,m=495,a=1259,s=93}
{x=897,m=2508,a=2414,s=928}
{x=350,m=142,a=280,s=1806}
{x=150,m=273,a=1174,s=258}
{x=3254,m=777,a=2033,s=22}
{x=520,m=1796,a=642,s=2391}
{x=82,m=983,a=247,s=682}
{x=2800,m=88,a=70,s=20}
{x=399,m=644,a=2699,s=1657}
{x=1242,m=976,a=1361,s=530}
{x=647,m=903,a=2150,s=10}
{x=1840,m=88,a=2003,s=2086}
{x=180,m=1806,a=3537,s=1941}
{x=2608,m=910,a=55,s=1923}
{x=2784,m=1083,a=730,s=2226}
{x=762,m=689,a=26,s=142}
{x=1763,m=1033,a=47,s=2210}
{x=1,m=1530,a=2951,s=806}
{x=187,m=898,a=572,s=1885}
{x=243,m=1263,a=527,s=213}
{x=181,m=414,a=2414,s=1040}
{x=712,m=229,a=1870,s=897}
{x=459,m=371,a=108,s=676}
{x=992,m=1193,a=40,s=225}
{x=1238,m=3,a=577,s=2792}
{x=1587,m=547,a=842,s=2891}
{x=3424,m=98,a=1142,s=453}
{x=258,m=229,a=368,s=1298}
{x=2956,m=55,a=568,s=385}
{x=357,m=3142,a=11,s=1636}
{x=1120,m=708,a=2200,s=1765}
{x=977,m=304,a=1485,s=245}
{x=135,m=1934,a=554,s=3201}
{x=1660,m=350,a=1669,s=577}
{x=1160,m=874,a=408,s=3168}
{x=367,m=57,a=90,s=319}
{x=2948,m=44,a=473,s=1256}
{x=336,m=848,a=1036,s=2203}
{x=321,m=2009,a=2133,s=1266}
{x=209,m=2937,a=710,s=521}
{x=198,m=693,a=479,s=719}
{x=1447,m=1068,a=1496,s=269}
{x=1401,m=960,a=1467,s=856}
{x=2256,m=483,a=1957,s=3515}
{x=52,m=2064,a=665,s=2876}
{x=184,m=300,a=1061,s=1314}
{x=2405,m=590,a=776,s=109}
{x=452,m=165,a=1349,s=1386}`
