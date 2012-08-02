$(function() {

	var config = {};
	var players = [];
		
	// 設定画面の決定ボタン
	$("#fix_config").bind("tap", function() {
		config.maxplayer = parseInt($("#maxplayer option:selected").attr("value"), 10);
		config.joker = parseInt($(".joker:checked").attr("value"), 10);
		$.mobile.changePage("#wait", {"transition":"flip"});
	});
	
	// 参加者募集画面
	$("#wait").bind("pagebeforeshow", function() {
		var players = $("#players");
		players.empty();
		players.append($('<div class="player"><img class="icon" src="/ultrarich/img/human.png"/><span class="name">あなた</span></div>'));
		for(var i = 1; i < config.maxplayer; i++) {
			players.append($('<div class="player"><img class="icon" src="/ultrarich/img/none.png"/><span class="name">空席</span></div>'));
		}
		
		$(".icon[src!='/ultrarich/img/human.png']").bind("tap", function() {
			var imgsrc = $(this).attr("src")
			if(imgsrc == "/ultrarich/img/cpu.png") {
				imgsrc = "/ultrarich/img/none.png";
			} else {
				imgsrc = "/ultrarich/img/cpu.png";
			}
			$(this).attr("src", imgsrc);
		});
	}).bind("pageinit", function() {
		$("#start").bind("tap", function() {
			$.mobile.changePage("#game")
		});
	});
});
