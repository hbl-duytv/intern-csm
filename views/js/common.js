function checkUserExist(username,callback){
    $.ajax({
        type:'post',
        url:'http://localhost:8000/check-user-exist',
        data:{'username':username},
        dataType:'json',
        success: function(result){
            return callback(result.check)
        }
    });
}
function checkEmailExist(email,callback){
    $.ajax({
        type:'post',
        url:'http://localhost:8000/check-email-exist',
        data:{'email':email},
        dataType:'json',
        success: function(result){
            return callback(result.check)
        }
    });
}
function vadidateRegister(id){
    $("#"+ id).validate({
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
}
function getTotalPage(callback){
    $.ajax({
        type:'get',
        url:'http://localhost:8000/total-page',
        data:{},
        dataType:'json',
        success: function(result){
            return callback(result.totalPage)
        }
    });
}