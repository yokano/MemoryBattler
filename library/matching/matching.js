var config = {
	maxplayer: 4,
	gamekey: "test"
}

var player = {
	name: "",
	id: "",
	inputName: function() {
		this.name = prompt("名前を入力してください");
	}
}

var view = {
	icons: {
		none: "/ultrarich/img/none.png",
		human: "/ultrarich/img/human.png",
		cpu: "/ultrarich/img/cpu.png"
	},
	create: function() {
		$("#players").append($('<div class="player"><img class="icon" src="' + this.icons.human + '"/><span class="name">' + player.name + '</span></div>'));
		for(var i = 1; i < config.maxplayer; i++) {
			$("#players").append('<div class="player"><img class="icon" src="' + this.icons.none + '"/><span class="name">空席</span></div>');
		}
	}
}

var ajax = {
	status: "",
	join: function() {
		$.ajax({
			url:"/matching",
			type:"GET",
			async: false,
			data:{
				gamekey:"test",
				name:player.name,
				action:"join"
			},
			dataType:"json",
			error:function() {
				console.log("部屋参加時にエラー");
			},
			success:function(data) {
				var channel = new goog.appengine.Channel(data.token);
				var socket = channel.open();
				player.id = data.id;
				console.log(player.id);
				socket.onmessage = function() {
					console.log("on message")
				};
				socket.onerror = function() {
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
				gamekey: "test",
				action: "leave",
				id: player.id
			},
			dataType:"json",
			error:function() {
				console.log("退室時にエラー");
			},
			success: function() {
				socket.close();
				ajax.status = "left";
			}
		});
	},
	get: function() {
		$.ajax({
			url:"/matching",
			type:"GET",
			data:{
				gamekey: "test",
				action: "get",
				id: player.id
			},
			dataType: "json",
			error: function() {
				console.log("ユーザ取得時にエラー");
			},
			success: function(data) {
				console.log("ユーザ取得成功");
				console.log(data);
			}
		});
	}
}

// entry
$(function() {
	player.inputName();
	view.create();
	ajax.join();
	if(ajax.status == "joined") {
		$(window).bind("unload", ajax.leave);
		ajax.get();
	}
});