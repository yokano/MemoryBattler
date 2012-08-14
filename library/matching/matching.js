var player = {
	name: "",
	id: "",
	inputName: function() {
		this.name = prompt("名前を入力してください");
	},
	host:false
}

var view = {
	icons: {
		none: "/ultrarich/img/none.png",
		human: "/ultrarich/img/human.png",
		cpu: "/ultrarich/img/cpu.png"
	},
	update: function(players) {
		$("#players").empty();
		for(var i = 0; i < config.maxplayer; i++) {
			$("#players").append('<div class="player"><img class="icon" src="' + this.icons.none + '"/><span class="name">空席</span></div>');
		}
		for(i = 0; i < players.length; i++) {
			$(".icon:eq(" + i + ")").attr("src", this.icons.human);
			$(".name:eq(" + i + ")").html(players[i].name);
		}
	},
	full: function() {
		$(":jqmData(role=content)").empty().html("満室のため入室できませんでした。");
	},
	hostOnly: function() {
		$(".host_only").show();
	}
}

var ajax = {
	status: "",
	socket: null,
	join: function() {
		$.ajax({
			url:"/matching",
			type:"GET",
			async: false,
			data:{
				gamekey:config.gamekey,
				name:player.name,
				action:"join"
			},
			dataType:"json",
			error:function() {
				console.log("部屋参加時にエラー");
			},
			success:function(data) {
				if(data == null) {
					ajax.status = "full";
					return;
				}
				var channel = new goog.appengine.Channel(data.token);
				ajax.socket = channel.open();
				player.id = data.id;
				console.log(player.id);
				ajax.socket.onmessage = function() {
					console.log("get message");
					view.update(ajax.get());
				};
				ajax.socket.onerror = function() {
					console.log("socket error")
				};
				ajax.status = "joined"
			}
		})
	},
	leave: function() {
		$.ajax({
			url:"/matching",
			type:"GET",
			async: false,
			data:{
				gamekey: config.gamekey,
				action: "leave",
				id: player.id
			},
			dataType:"json",
			error:function() {
				console.log("退室時にエラー");
			},
			success: function() {
				ajax.socket.close();
				ajax.status = "left";
				ajax.message({id:player.id, content:"update"})
				console.log("leave");
			}
		});
	},
	get: function() {
		var result = null;
		$.ajax({
			url:"/matching",
			type:"GET",
			async: false,
			data:{
				gamekey: config.gamekey,
				action: "get",
				id: player.id
			},
			dataType: "json",
			error: function() {
				console.log("ユーザ取得時にエラー");
			},
			success: function(players) {
				console.log("ユーザ取得成功");
				console.log(players);
				result = players;
			}
		});
		return result;
	},
	message: function(message) {
		$.ajax({
			url: "/message",
			type: "POST",
			data: {
				gamekey: config.gamekey,
				"message": JSON.stringify(message)
			},
			dataType: "json",
			error: function() {
				console.log("メッセージ送信時にエラー");
			},
			success: function() {
				console.log("send message");
			}
		});
	}
}

// entry
$(function() {
	player.inputName();
	ajax.join();
	if(ajax.status == "full") {
		view.full();
	} else if(ajax.status == "joined") {
		$(window).bind("beforeunload", ajax.leave);
		var players = ajax.get();
		if(players.length == 1) {
			player.host = true;
			view.hostOnly();
		}
		view.update(players);
		ajax.message({"id":player.id, "content":"update"});
	}
});
