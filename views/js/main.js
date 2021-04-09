$(function () {
	$('#login-form-link').click(function (e) {
		$("#login-form").delay(100).fadeIn(100);
		$("#register-form").fadeOut(100);
		$('#register-form-link').removeClass('active');
		$(this).addClass('active');
		e.preventDefault();
	});
	$('#register-form-link').click(function (e) {
		$("#register-form").delay(100).fadeIn(100);
		$("#login-form").fadeOut(100);
		$('#login-form-link').removeClass('active');
		$(this).addClass('active');
		e.preventDefault();
	});
	$("#usernameRegister").focusout(function () {
		let username = $("#usernameRegister").val()
		checkUserExist(username, function (result) {
			if (result == false) {
				$('#messageUsername').css('color', 'red');
				$('#messageUsername').text('Tài khoản đã tồn tại!');
				$('#register-submit').attr('disabled', 'disabled');
			} else {
				$('#messageUsername').css('color', 'green');
				$('#messageUsername').text('Tài khoản có thể sử dụng!');
				$('#register-submit').removeAttr('disabled');
			}
		});
	});
	$("#emailRegister").focusout(function () {
		let email = $("#emailRegister").val()
		checkEmailExist(email, function (result) {
			if (result == false) {
				$('#messageEmail').css('color', 'red');
				$('#messageEmail').text('Email đã tồn tại!');
				$('#register-submit').attr('disabled', 'disabled');
			} else {
				$('#messageEmail').hide();
				$('#register-submit').removeAttr('disabled');
			}
		})
	});
});
$(document).ready(function () {
	vadidateRegister('register-form');
});
var username = $('#username').val();
var password = $('#password').val();
