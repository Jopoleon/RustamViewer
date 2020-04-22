
var currentUrl = window.location.href
$(document).ready(function () {
	

	$("#gotosignup").click(function (e) {
		e.preventDefault();
		window.location = currentUrl + "signup";
	});
	$("#gotologin").click(function (e) {
		e.preventDefault();
		window.location = currentUrl + "login";
	});
	$("#logout").click(function (e) {
		e.preventDefault();
		//console.log(currentUrl.replace("login", ""))
		window.location = currentUrl + "logout"
	});
	console.log(currentUrl)
	$("#addtracklocationButton").click(function () {
		console.log($("#locationtrackInput").val())

		window.locationText = $("#locationtrackInput").val();
		//window.usernameText = $("#usernameInput").val();
		console.log(currentUrl + "watchlocation")
		console.log(window.locationText)

		$.ajax({
			data: {
				"location": window.locationText,
				//"username": window.usernameText
			},
			//dataType: "json",
			type: "POST",

			url: currentUrl + "watchlocation",

			success: function (data) {
				console.log(data)
				$("#ajaxResults").empty();
				$('#ajaxResults').append(data)

			},
			error: function (req, status, err) {
				$("#ajaxResults").empty();
				console.log(req.responseText)
				$('#ajaxResults').append(req.responseText);
				console.log('Something went wrong', status, err);
				console.log(err)

			}
		});
	});
	
	$("#deletetrcklocationButton").click(function () {
		console.log($("#locationInput").val())

		window.locationText = $("#locationtrackInput").val();
		//window.usernameText = $("#usernameInput").val();


		$.ajax({
			data: {
				"location": window.locationText,
				//"username": window.usernameText
			},
			//dataType: "json",
			type: "POST",

			url: currentUrl + "deletewatchlocation",

			success: function (data) {
				console.log(data)
				$("#ajaxResults").empty();
				$('#ajaxResults').append(data)

			},
			error: function (req, status, err) {
				$("#ajaxResults").empty();
				console.log(req.responseText)
				$('#ajaxResults').append(req.responseText);
				console.log('Something went wrong', status, err);
				console.log(err)

			}
		});
	});



	$("#locationButton").click(function () {
		console.log($("#locationInput").val())

		window.locationText = $("#locationInput").val();



		$.ajax({
			data: {
				"value": window.locationText
			},
			dataType: "json",
			type: "POST",

			url: currentUrl + "scrapelocation",


			success: function (data) {
				$("#ajaxResults").empty();
				console.log(data)

				console.log(data[1].Title)
				console.log(data.Cost)
				for (var i = 0; i < data.length; i++) {
					var imagetag = "<img src=" + data[i].ImageUrl + " style='width:100px;height:100px;'>";

					$('#ajaxResults').append(data[i].Title + " " + data[i].Cost + "<br>" + imagetag + "<br><br>");


				}

			},
			error: function (req, status, err) {
				$("#ajaxResults").empty();
				console.log(req.responseText)
				$('#ajaxResults').append(req.responseText);
				console.log('Something went wrong', status, err);
				console.log(err)

			}
		});
	});
});