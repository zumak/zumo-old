import $ from "/static/js/minilib.js";
$._registerGlobal(); // load lib first, for custom element

var messageBox = $.get("zumo-message-box");
//messageBox.info("test");

$.get("zumo-channel-list").on("change-channel", function(evt){
	//var channelId = evt.details.ID;
	var channelId = evt.detail.ID;
	$.get("zumo-input-box").attr("channel-id", channelId);
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
		$.get("zumo-message-box").info(`${name} channel created!`);
		this.hide();
		this.clear();
	} catch(e) {
		// TODO show error message
		console.warn("error on create channel", name)
		$.get("zumo-message-box").error("error on create channel!");
		this.hide();
	}
}).on("cancel", function(){
	$.get(this, "input[type=text]").value = "";
	this.hide();
	$.get("zumo-message-box").info("create channel canceled");
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
		$.get("zumo-message-box").info("join channel");
	} catch(e) {
		this.hide();
		$.get("zumo-message-box").error("join channel failed");
	}
}).on("cancel", function(){
	this.hide();
});

window.onload = function() {
	$.get("zumo-channel-list").load();
	$.get("zumo-socket").connect();
}
