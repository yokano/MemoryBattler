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
	join: function() {
		$.ajax({
			url:"/matching",
			type:"GET",
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
					alert("MESSAGE")
				};
				socket.onerror = function() {
					alert("ERROR")	
				};
				$(window).bind("unload", ajax.leave);
			}
		})
	},
	leave: function() {
		$.ajax({
			async: false,
			url:"/matching",
			type:"GET",
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
			}
		});
	}
}

// entry
$(function() {
	player.inputName();
	view.create();
	ajax.join();
});