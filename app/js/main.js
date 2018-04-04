//import $ from "/static/js/minilib.js";
//$._registerGlobal(); // load lib first, for custom element


$.get("zumo-channel-list").on("change-channel", function(evt){
	//var channelId = evt.details.ID;
	var channelId = evt.detail.ID;
	$.get("zumo-input-box").attr("channel-id", channelId);
	$.get(".channel-header h2").innerHTML = evt.detail.Name;
	$.get("zumo-messages").load(channelId);
});

$.get("button.create-channel").on("click", async function() {
	$.get("zumo-dialog.create-channel").show();
});
$.get("zumo-dialog.create-channel").on("ok", async function(evt){
	var name = $.get(this, "input[type=text]").value.trim();
	try {
		var channel = await $.request("POST", "/api/v1/channels", {
			body: {
				Name: name
			}
		});
		console.log("create channel", name)
		$.get("zumo-alert-box").info(`${name} channel created!`);
		this.hide();
		this.clear();
	} catch(e) {
		// TODO show error message
		console.warn("error on create channel", name)
		$.get("zumo-alert-box").error("error on create channel!");
		this.hide();
	}
}).on("cancel", function(){
	$.get(this, "input[type=text]").value = "";
	this.hide();
	$.get("zumo-alert-box").info("create channel canceled");
});
$.get("zumo-dialog.create-bot").on("ok", async function(evt){
	var name = $.get(this, "input[name=name]").value.trim();
	var driver = $.get(this, "input[name=driver]").value.trim();

	if (name == "") {
		$.get("zumo-alert-box").error("name is blank");
		return
	}

	if (driver == "") {
		$.get("zumo-alert-box").error("driver is blank");
		return
	}

	try {
		await $.request("POST", "/api/v1/bots", {
			body: {
				Name: name,
				Driver: driver,
			}
		});
		$.get("zumo-alert-box").info("bot created");
		this.hide()
	} catch(e) {
		$.get("zumo-alert-box").error("failed", e);
		this.hide()
	}

}).on("cancel", function(){
	$.get("zumo-alert-box").info("create bot canceled");
	this.hide();
});

$.get("button.join-channel").on("click", async function() {
	// TODO load channel
	var res = await $.request("GET", "/api/v1/channels")
	$.get("zumo-dialog.join-channel select").clear();
	res.json.map(function(e) {
		return $.create("option", {
			$text: `${e.Name}`,
			"value": e.ID
		});
	}).forEach(function(e){
		$.get("zumo-dialog.join-channel select").appendChild(e);
	}, this);

	$.get("zumo-dialog.join-channel").show();
});
$.get("zumo-dialog.join-channel").on("ok", async function() {
	try {
		var channelID = $.get(this, "select").value
		await $.request("PUT", `/api/v1/channels/${channelID}/join`, {})
		this.hide();
		$.get("zumo-alert-box").info("join channel");
	} catch(e) {
		this.hide();
		$.get("zumo-alert-box").error("join channel failed");
	}
}).on("cancel", function(){
	this.hide();
});

$.get("button.logout").on("click", async function() {
	try {
		await $.request("GET", "/api/v1/channels", {
			$auth:{
				user: "__",
				password: "__",
			}
		});
	} catch(e) {
		//$.get("zumo-alert-box").error("Error on Logout");
	}
	document.location = "/"
});

$.get("zumo-menu.application").on("menu", function(e) {
	switch(e.detail.name) {
		case "create-bot":
			$.get("zumo-dialog.create-bot").show();
			this.hide();
			break;
	}
});

window.onload = function() {
	$.get("zumo-channel-list").load();
	$.get("zumo-socket").connect();
	$.get("zumo-socket").on("message", function(evt) {
		$.get("zumo-messages").appendText(evt.detail);
		$.get("zumo-messages").scrollToEnd();
	});
}
