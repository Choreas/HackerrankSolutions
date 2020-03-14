class MorganAndaString {

	static charlist = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

	MorganAndaString() {

	}

	getMinLex(a, b) {

		var sa = a.split('');
		var sb = b.split('');
		var combined = sa.length + sb.length;
		var empty = false
		var res = new String();
		var la = sa.length;
		var lb = sb.length;

		var lRes = 1;
		//return combined;
		for (var ct = 0; ct <= combined - 1 && empty == false; ct++) {
			//console.log(sb);

			var _, src = this.chooseLex(sa, sb, la, lb, 0)
			if (src == "a") {
				res = res.padEnd(lRes, sa.shift());
				if (la - 1 == 0) {
					empty = true;
				}
				la--;
			} else {
				res = res.padEnd(lRes, sb.shift());
				if (lb - 1 == 0) {
					empty = true;
				}
				lb--;
			}
			lRes++;

		}
		if (la > 0) {
			for (var ii = 0; ii < la; ii++) {
				res = res.padEnd(lRes, sa.shift());
				lRes++;
			}
		} else if (lb > 0) {
			for (var ii = 0; ii < lb; ii++) {
				res = res.padEnd(lRes, sb.shift());
				lRes++;
			}
		}
		return res
	}

	chooseLex(a, b, la, lb, idx) {

		var ca = a[idx];
		var cb = b[idx];
		var ia = MorganAndaString.charlist.indexOf(ca);
		var ib = MorganAndaString.charlist.indexOf(cb);
		var char;
		var src;

		if (ib == ia) {
			if (la > idx + 1 && lb > idx + 1) {
				_, src = this.chooseLex(a, b, la, lb, idx + 1);
				if (src == "a") {
					char = ca;
				} else {
					char = cb;
				}
			} else if (la < lb) {
				src = "b";
				char = cb;
			} else {
				src = "a";
				char = ca;
			}
		} else if (ia > ib) {
			src = "b";
			char = cb;
		} else {
			src = "a";
			char = ca;
		}

		return char, src;
	}


}

module.exports = MorganAndaString;