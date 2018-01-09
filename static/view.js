function getDefinitions(word, callback) {
	var xmlHttp = new XMLHttpRequest();
	xmlHttp.onreadystatechange = function() { 
		if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
			var resp = JSON.parse(xmlHttp.responseText);
			if (resp.error !== undefined) {
				console.warn("Error for "+word+": "+resp.error)
			}
			else {
				callback(resp);
			}
		}
	}
	xmlHttp.open("GET", "/def/"+word, true);
	xmlHttp.send(null);
}

function Network(startWord) {
	self = {}

	self.expand = function(params) {
		// Must have a node clicked
		if (params.nodes.length === 0) {
			return;
		}
		
		// Get word to lookup
		var word = params.nodes[0];

		// Don't look it up if we already did
		if (self.clicked[word]) {
			return;
		}
		self.clicked[word] = true;

		getDefinitions(word, function(response) {
			var seen = {}
			response.definitions[0].replace(/\b\S+\b/g, function(w) {
				if (!self.exists[w]) {
					self.nodes.add({id: w, label: w});
					self.exists[w] = true;
				}
				if (!seen[w]) {
					self.edges.add({from: word, to: w});
				}
				seen[w] = true;
			});
		});
	}

	self.clicked = {}
	self.exists = {}

	// Set up single node data first
	self.nodes = new vis.DataSet([
		{id: startWord, label: startWord},
	]);
	self.edges = new vis.DataSet([]);

	// provide the data in the vis format
	var data = {
		nodes: self.nodes,
		edges: self.edges
	};
	var options = {};

	// initialize your network!
	var container = document.getElementById("graph");
	var network = new vis.Network(container, data, options);
	network.on("click", self.expand)

	return self
}

startWord = document.getElementById("start-word").innerText;
network = new Network(startWord);
