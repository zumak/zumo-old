import $ from "/static/js/minilib.js";

class ReconnetSocket {
	static protocal() {
		return document.location.protocol.includes("https") ? "wss:" : "ws:"
	}
	constructor(url){
		// connect -> open -> message -> close -> disconnect
		// connect -> open -> message -> close -> (reconnect) -> open -> message -> close -> disconnect
		this._url = url;
		this._listeners = {
			"open": [],
			"close": [],
			"connect": [],
			"disconnect": [],
			"message": [],
			"error": [],
		};
		this._tryConnect();
	}
	on(event, listener){
		this._listeners[event].push(listener)
	}
	send(data, options) {
		try{
			this._socket.send(data, options);
		} catch (e){
			this._emit("error", e)
		}
	}

	_emit(event, data) {
		this._listeners[event].forEach(function(l){
			l(data);
		})
	}
	_tryConnect() {
		console.debug("[ReconnetSocket:_tryConnect]")
		this._socket = new WebSocket(this._url);

		this._socket.on("open", this._emit.bind(this, "open"));
		this._socket.on("close", this._reciveClose.bind(this));
		this._socket.on("message", this._emit.bind(this, "message"));
		this._socket.on("error", this._emit.bind(this, "error"));
	}

	async _reciveClose(){
		await $.timeout(500); // wait 500ms
		this._tryConnect();
	}
}

export default ReconnetSocket;
