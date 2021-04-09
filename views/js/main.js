$(function() {
    $('#login-form-link').click(function(e) {
		$("#login-form").delay(100).fadeIn(100);
 		$("#register-form").fadeOut(100);
		$('#register-form-link').removeClass('active');
		$(this).addClass('active');
		e.preventDefault();
	});
	$('#register-form-link').click(function(e) {
		$("#register-form").delay(100).fadeIn(100);
 		$("#login-form").fadeOut(100);
		$('#login-form-link').removeClass('active');
		$(this).addClass('active');
		e.preventDefault();
	});
	$("#usernameRegister").focusout(function() {
		$.ajax({
			type:'post',
			url:'http://localhost:8000/check-user-exist',
			data:{'username':$("#usernameRegister").val()},
			dataType:'json',
			success: function(result){
				if(result.check == false){
					$('#messageUsername').css('color', 'red');
				  $('#messageUsername').text('Tài khoản đã tồn tại!');
				  $('#register-submit').attr('disabled','disabled');
				} else {
					$('#messageUsername').css('color', 'green');
					$('#messageUsername').text('Tài khoản có thể sử dụng!');
				  $('#register-submit').removeAttr('disabled');
				}
			}
		});
	});
	$("#emailRegister").focusout(function() {
		$.ajax({
			type:'post',
			url:'http://localhost:8000/check-email-exist',
			data:{'email':$("#emailRegister").val()},
			dataType:'json',
			success: function(result){
				if(result.check == false){
					$('#messageEmail').css('color', 'red');
				  $('#messageEmail').text('Email đã tồn tại!');
				  $('#register-submit').attr('disabled','disabled');
				} else {
					$('#messageEmail').hide();
				  $('#register-submit').removeAttr('disabled');
				}
			}
		});
	});
});
$(document).ready(function () {
  // check confirm password
	$("#register-form").validate({
		rules: {
			passwordRegister: { 
					required: true,
						 minlength: 6,
						 maxlength: 10,
				} , 
				confirmPassword: { 
						  equalTo: "#passwordRegister"
				}
		},
		messages:{
		passwordRegister: { 
						required:"The password is required"
					}
		}
		});
});
