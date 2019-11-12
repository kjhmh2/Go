var is_gameing = 0;
var is_hinting = 0;
var is_clicked = 0;
var time = 120;

// 移动的设置
function move(event) {
	var get = document.getElementById("part16");
	var blank_top = get.offsetTop;
	var blank_left = get.offsetLeft;
	var this_top = this.offsetTop;
	var this_left = this.offsetLeft;
	if ((Math.abs(blank_top - this_top) == 100 && blank_left == this_left) || (Math.abs(blank_left - this_left) == 100 && blank_top == this_top)) {
		var str = get.className;
		get.className = this.className;
		this.className = str;
		check();
	}
}

// 判断是否已经成功了
function check() {
	if (is_gameing == 0) return;
	for (var i = 1; i <= 16; ++ i) {
		var temp = document.getElementById("part" + i);
		if (temp.className != "common" + " position" + i) return;
	}
	// 获胜
	is_gameing = 0;
	$("#result").text("You win!");
	clearInterval(clear);
	is_clicked = 0;
}

// 随机摆放图片
function random_pos() {
	is_gameing = 1;
	time = 120;
	if (is_clicked == 1) clearInterval(clear);
	clear = setInterval(pass, 10);
	is_clicked = 1;
	$("#result").text("Gaming");
	for (var i = 1; i <= 16; ++ i) {
		document.getElementById("part" + i).className = "common" + " position" + i;
	}
	store = [];
	for (var j = 0; j <= 14; ++ j) {
		store[j] = j + 1;
	}
	while(true) {
		store.sort(cmp);
		if (isVaild()) break;
	}
	for (var k = 1; k <= 15; ++ k) {
		document.getElementById("part" + k).className = "common" + " position" + store[k - 1];
	}
}

// 将store数组随机打乱排列
function cmp() {
	return 0.5 - Math.random();
}

// 从逻辑上判断是否能够拼图成功
function isVaild() {
	var count = 0;
    for (var i = 0; i <= 15; i ++) {
        for (var j = i + 1; j <= 15; j ++) {
            if (store[j] < store[i]) {
                count ++;
            }
        }
    }
    // 判断错序系数的奇偶性
    return count % 2 === 0;
}

// 给每个方块加上一个标记
function sign() {
	if (is_hinting == 0) {
		for (var i = 1; i <= 16; ++ i) {
			$("<div class = 'number' id = " + i + ">" + i + "</div>").appendTo($("#part" + i));
		}
		is_hinting = 1;
	}
	else {
		for (var i = 1; i <= 16; ++ i) {
			$("#" + i).remove();
		}
		is_hinting = 0;
	}
}

function pass() {
	// 失败
	if (time <= 0) {
		is_gameing = 0;
		$("#result").text("You Lose!");
		$("#color").css("opacity", "0");
		clearInterval(clear);
		is_clicked = 0;
		return;
	}
	// 颜色的渐变
	$("#color").css("width", time * 2.5 + "px");
	var sub = 120 - time;
	var r,g,b;
	if (time >= 60) {
		r = 80 + parseInt(2.28 * sub);
		g = 217;
	}
	else {
		r = 217;
		g = 217 - parseInt(2.42 * (sub - 60));
	}
	var b = 55;
	$("#color").css("backgroundColor", "rgb(" + r + "," + g + "," + b + ")");
	time = time - 0.01;
}

$(document).ready(function() {
	$("#result").text("Please click the restart button to start a new game!");
	for (var i = 1; i <= 16; ++ i) {
		$("#part" + i).on("click", move);
	}
	$("#start").on("click", random_pos);
	$("#hint").on("click", sign);
});