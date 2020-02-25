$(document).ready(function () {


	var currentUrl = window.location.href
	console.log(currentUrl)

	//console.log(currentUrl.replace("signup", ""))
	//console.log(String(currentUrl) - "/signup")

	$("#signupbutton").click(function () {
		console.log($("#userName").val())
		console.log($("#password").val())
		console.log($("#email").val())

		window.userName = $("#userName").val();
		window.password = $("#password").val();
		window.email = $("#email").val();

		$.ajax({
			data: {
				"username": window.userName,
				"password": window.password,
				"email": window.email,
				"urlforactivation": currentUrl.replace("signup", "confirm"),
			},
			//dataType: "json",
			type: "POST",

			url: currentUrl + "submit",


			success: function (data) {
				$("#signupresults").empty()
				console.log("Data sent: ", data)
				$("#signupresults").append(data);
				console.log(data)


			},
			error: function (req, status, err) {
				//console.log(req.responseText)
				console.log(req)

				console.log('Something went wrong', status, err);
				console.log(err)

			}
		});
//		setTimeout(function () {
//				window.location = currentUrl.replace("signup", "");
//
//		}, 4000);

	});
	$("#homepageredirect").click(function (e) {
		e.preventDefault();
		//console.log(currentUrl.replace("login", ""))
		window.location = currentUrl.replace("signup", "");
	});
	$("#logout").click(function (e) {
		e.preventDefault();
		//console.log(currentUrl.replace("login", ""))
		window.location = currentUrl.replace("signup", "logout");
	});


});