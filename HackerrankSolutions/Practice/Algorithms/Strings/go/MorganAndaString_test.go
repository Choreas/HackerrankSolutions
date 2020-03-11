package strings

import "testing"

func Test_morganAndString(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Case 1", args{"ACA", "BCF"}, "ABCACF"},
		{"Case 2", args{"JACK", "DANIEL"}, "DAJACKNIEL"},
		{"Case 3", args{"ABACABA", "ABACABA"}, "AABABACABACABA"},
		{"Case 4", args{longA, longB}, longWant},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := morganAndString(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("morganAndString() = %v, want %v", got, tt.want)
			}
		})
	}
}

const longA string = "YKBTHKJXKPXPZRVQXMFNDSMZXVJNBKZOGSCHQIZWRZNWRBQKYUXLEEEGKKKQAZGUTNWREXASKAEAREVNHRZMHMROYWSRJKPNQLSGKTDISVWXILBPAVAZFFTZYQETTLAIBFSJTBXZRGALUHPAFCMWJKVTIXOTOLZKSQSGGZDGKRZJLSHCWBZVBMKVEGLIMTRCGRKVTHXAJKTDYAKUNBYERJYVVXTQPXRNAODMFYJWZULIIYSFJPSXPQGYBOSKTJMPRGFWXNKEVSYDEQSREGVSTDXWHYBREKRADCGOBWGILBVDQOQAECOPOOVTGVVZEUNAYRPPEQITNLPYIIQCQBQQIMNXZHTHERRURXNGCMZCDFUKOVMMAIWIFKBOZRBZYWAVPAYORRQVTTEIBEKWJRXQBNLUCIWQGEFQHFCFCLQDGJNUIKYPIVOAWVRVZZVLJQLRGDDWIZMBAVKMCFXFLBCYCKVSSCYPVYYBILEOTTANEGQEGPHHNCLUFINTGIDIYVJFSAHDZFUQDOPUOOBBHLNWWZUJNYNCUMCKCMYSOLLXGUJKQDKWFXWMNZCMFCSWJXKYUIZAQCHDQOEZBWLJKUWMYLNVZAGFFQZJCMTXLNFQPASBKXKBESKXECRGYMZNDICDVICQPKBTCZTXTWKTQABQTMWJOTJAZDSYMHHFLGYAAWXTUKCRPNQMHWCDENBDWGIDCPLYLOXTBWDSBTRUJSREGNXTDRJINMUJLGQHNWMAQFEBIQEOSRGATZLGTCWKGFXNVLDJYEZSXUDDNWMCSVCCIDJIRBUJTAIGDXCPJZICRTHGPBFEECHVDCQEMYCXMIOQTQWQWZDCQCZVIGZPIPSYRWZYZLSQJQKXRQQFEZULLKTNIAPZSFULWUDENPIUGSSDRCAAVXVVIFHVZMPRTVHGFCLOVDDQHQYERDUDFMPPBPKAVIKGKQAFBYTTSAEEDWETSYOXKDLQDUMTXRKKWUHAGDNQSKFCXMYMOPCJYMGJ"
const longB string = "QTIYJTGDQGXXIUBJOKQTRTCJPKEVLGXFSLMORYPCXGQHTJEUNTOAAJMNQHCOORXYCLOMMLPBVPYMRKMHUQGMEQULDUCFWGNHLGBIKXDDSXLSDVNIIMVWUMFVGCELJBYBMOVCJQQQUAJHONUFWZLMQGQCZBWJAMZOPBOTMLRPQUEGFQEEKBADRWPEQKXXTIVKPLDYNWPMXLSYNIDOCRDLFGEDAYTCEDRXXAUGMHRCXSATBMPWYCXWULHDPYJELWJXIHOJZUGLKYYBCEMYXNBYAURETPMIYAOYXEJRPNXNSCOCSCLDHWNNIYEBRUEOTMSWXRCVJPBOPPPOHRJMXLCXBOKFJVOIOWHJPVXAFOYWOAXCKCFUNFKTQOUFBKKOEWYWZCQOQKRBNDLVFIEVZPLQBSEBMQOREQJCKSMEXZJALKQCIYKFTIYEDAXAXJLCCLQVCQLVZPCYOKXVCJGJWHUYPNAPITGFBQTPBDONKUFFCCWGKAGAWQYNAEECYXCYRJSGFMTAVTWHSFRBUJOWHUTHSLRPZNVNVGEQMHRYIJVAJPPISHAOWYEDZILXSDTDANELCCHNVXJLVAFPHWLKPEWOBMDJSNQEBPDVUPVULLJQCYPVTUSDNWURVEPRQDEDBMRQBSVFBXEMUDMELQJULAMYHLKUTQTRNMKFTQMAYEXYFVYEDGKHQFUWJTXZRZHWTGBBLNTJIKQCPHUYNXUVUPBOWTJXDFYUKEWASYAVCWTIHTSOFEXZTMPINBHIYGRCLIDCDBNGWKFNVMIIHRVSPTOLXTDEWRHDDLYQEWBZHWNZNGTIRHKHXXYYNBIRTISUGIIVRRASLOAPOWRIQJASRRPYNJOZFYXMBOOPPRVBYZUVXNZAWDAFWOOGPQBMUOKBSECZQUYPWUPWIYEUSDFZYMUFCBSEEPILTYHUYVMXUTSSCVWMKOYSSUJRFJFXKPQKJJAZCXEOECLFCARVOQLLVZJLSPBBCQFSRVVYRWUJYBGZ"
const longWant string = "QTIYJTGDQGXXIUBJOKQTRTCJPKEVLGXFSLMORYKBTHKJXKPXPYPCXGQHTJEUNTOAAJMNQHCOORXYCLOMMLPBVPYMRKMHUQGMEQULDUCFWGNHLGBIKXDDSXLSDVNIIMVWUMFVGCELJBYBMOVCJQQQUAJHONUFWZLMQGQCZBWJAMZOPBOTMLRPQUEGFQEEKBADRWPEQKXXTIVKPLDYNWPMXLSYNIDOCRDLFGEDAYTCEDRXXAUGMHRCXSATBMPWYCXWULHDPYJELWJXIHOJZRVQXMFNDSMZUGLKYYBCEMYXNBYAURETPMIYAOYXEJRPNXNSCOCSCLDHWNNIYEBRUEOTMSWXRCVJPBOPPPOHRJMXLCXBOKFJVOIOWHJPVXAFOYWOAXCKCFUNFKTQOUFBKKOEWYWZCQOQKRBNDLVFIEVZPLQBSEBMQOREQJCKSMEXZJALKQCIYKFTIYEDAXAXJLCCLQVCQLVZPCYOKXVCJGJWHUYPNAPITGFBQTPBDONKUFFCCWGKAGAWQYNAEECYXCYRJSGFMTAVTWHSFRBUJOWHUTHSLRPZNVNVGEQMHRYIJVAJPPISHAOWYEDZILXSDTDANELCCHNVXJLVAFPHWLKPEWOBMDJSNQEBPDVUPVULLJQCYPVTUSDNWURVEPRQDEDBMRQBSVFBXEMUDMELQJULAMYHLKUTQTRNMKFTQMAYEXYFVYEDGKHQFUWJTXZRZHWTGBBLNTJIKQCPHUYNXUVUPBOWTJXDFYUKEWASYAVCWTIHTSOFEXZTMPINBHIYGRCLIDCDBNGWKFNVMIIHRVSPTOLXTDEWRHDDLYQEWBZHWNZNGTIRHKHXXYYNBIRTISUGIIVRRASLOAPOWRIQJASRRPYNJOZFYXMBOOPPRVBYZUVXNZAWDAFWOOGPQBMUOKBSECZQUYPWUPWIYEUSDFZXVJNBKZOGSCHQIZWRZNWRBQKYUXLEEEGKKKQAZGUTNWREXASKAEAREVNHRZMHMROYWSRJKPNQLSGKTDISVWXILBPAVAZFFTZYMUFCBSEEPILTYHUYVMXUTSSCVWMKOYSSUJRFJFXKPQKJJAZCXEOECLFCARVOQLLVZJLSPBBCQFSRVVYRWUJYBGZYQETTLAIBFSJTBXZRGALUHPAFCMWJKVTIXOTOLZKSQSGGZDGKRZJLSHCWBZVBMKVEGLIMTRCGRKVTHXAJKTDYAKUNBYERJYVVXTQPXRNAODMFYJWZULIIYSFJPSXPQGYBOSKTJMPRGFWXNKEVSYDEQSREGVSTDXWHYBREKRADCGOBWGILBVDQOQAECOPOOVTGVVZEUNAYRPPEQITNLPYIIQCQBQQIMNXZHTHERRURXNGCMZCDFUKOVMMAIWIFKBOZRBZYWAVPAYORRQVTTEIBEKWJRXQBNLUCIWQGEFQHFCFCLQDGJNUIKYPIVOAWVRVZZVLJQLRGDDWIZMBAVKMCFXFLBCYCKVSSCYPVYYBILEOTTANEGQEGPHHNCLUFINTGIDIYVJFSAHDZFUQDOPUOOBBHLNWWZUJNYNCUMCKCMYSOLLXGUJKQDKWFXWMNZCMFCSWJXKYUIZAQCHDQOEZBWLJKUWMYLNVZAGFFQZJCMTXLNFQPASBKXKBESKXECRGYMZNDICDVICQPKBTCZTXTWKTQABQTMWJOTJAZDSYMHHFLGYAAWXTUKCRPNQMHWCDENBDWGIDCPLYLOXTBWDSBTRUJSREGNXTDRJINMUJLGQHNWMAQFEBIQEOSRGATZLGTCWKGFXNVLDJYEZSXUDDNWMCSVCCIDJIRBUJTAIGDXCPJZICRTHGPBFEECHVDCQEMYCXMIOQTQWQWZDCQCZVIGZPIPSYRWZYZLSQJQKXRQQFEZULLKTNIAPZSFULWUDENPIUGSSDRCAAVXVVIFHVZMPRTVHGFCLOVDDQHQYERDUDFMPPBPKAVIKGKQAFBYTTSAEEDWETSYOXKDLQDUMTXRKKWUHAGDNQSKFCXMYMOPCJYMGJZ"
