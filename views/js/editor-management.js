$(document).ready(function () {
    $('#editorsTable').DataTable();
    $('#modal-btn-ok-noti').click(function () {
        $('#notification-modal').modal('toggle');
        location.reload()
    });
    vadidateRegister('register-form');
    $(function () {
        $("#birthdayRegister").datepicker({
            dateFormat: 'dd/mm/yy',
            changeMonth: true,
            changeYear: true,
            yearRange: '-100y:c+nn',
            maxDate: '-1d'
        });
    });

});

var idActive
var idDeactive
var idDelete
var modalConfirmActive = function (callback) {
    $(".btn-confirm-active").on("click", function () {
        idActive = this.id.substring(this.id.indexOf("-") + 1)
        $("#active-modal").modal('show');
    });
    $("#modal-btn-ok-active").on("click", function () {
        callback(true);
        $("#active-modal").modal('hide');
    });

    $("#modal-btn-back-active").on("click", function () {
        callback(false);
        $("#active-modal").modal('hide');
    });
};
modalConfirmActive(function (confirm) {
    if (confirm) {
        $.ajax({
            type: 'post',
            url: 'http://localhost:8000/active-editor',
            data: { 'id': idActive },
            dataType: 'json',
            success: function (result) {
                if (result.status == 200) {
                    $('#labelNotiModel').text('Kích hoạt thành công!');
                    $('#notification-modal').modal('show');
                } else {
                    $('#labelNotiModel').text('Kích hoạt thất bại!');
                    $('#notification-modal').modal('show');
                }
            }
        });
    } else {

    }
});
var modalConfirmDeactive = function (callback) {
    $(".btn-confirm-deactive").on("click", function () {
        idDeactive = this.id.substring(this.id.indexOf("-") + 1)
        $("#deactive-modal").modal('show');
    });
    $("#modal-btn-ok-deactive").on("click", function () {
        callback(true);
        $("#deactive-modal").modal('hide');
    });

    $("#modal-btn-back-deactive").on("click", function () {
        callback(false);
        $("#deactive-modal").modal('hide');
    });
};
modalConfirmDeactive(function (confirm) {
    if (confirm) {
        $.ajax({
            type: 'post',
            url: 'http://localhost:8000/deactive-editor',
            data: { 'id': idDeactive },
            dataType: 'json',
            success: function (result) {
                if (result.status == 200) {
                    $('#labelNotiModel').text('Bỏ kích hoạt thành công!');
                    $('#notification-modal').modal('show');
                } else {
                    $('#labelNotiModel').text('Bỏ kích hoạt thất bại!');
                    $('#notification-modal').modal('show');
                }
            }
        });
    } else {

    }
});
var modalConfirmDelete = function (callback) {
    $(".btn-confirm-delete").on("click", function () {
        idDelete = this.id.substring(this.id.indexOf("-") + 1)
        $("#delete-modal").modal('show');
    });
    $("#modal-btn-ok-delete").on("click", function () {
        callback(true);
        $("#delete-modal").modal('hide');
    });

    $("#modal-btn-back-delete").on("click", function () {
        callback(false);
        $("#delete-modal").modal('hide');
    });
};
modalConfirmDelete(function (confirm) {
    if (confirm) {
        $.ajax({
            type: 'post',
            url: 'http://localhost:8000/delete-user',
            data: { 'id': idDelete },
            dataType: 'json',
            success: function (result) {
                if (result.status == 200) {
                    $("#labelNotiModel").text('Xóa thành công!');
                    $("#notification-modal").modal('show');
                } else {
                    $("#labelNotiModel").text('Xóa thất bại!');
                    $("#notification-modal").modal('show');
                }
            }
        });
    } else {
    }
});
var modalCreateEditor = function (callback) {
    $("#btn-add-editor").on("click", function () {
        $("#create-editor-modal").modal('show');
    });
    $("#modal-btn-ok-create").on("click", function () {
        callback(true);
        $("#create-editor-modal").modal('hide');
    });
    $("#modal-btn-back-create").on("click", function () {
        callback(false);
        $("#create-editor-modal").modal('hide');
    });
};
modalCreateEditor(function (confirm) {
    var usernameRegister = $("#usernameRegister").val()
    var passwordRegister = $("#passwordRegister").val()
    var emailRegister = $("#emailRegister").val()
    var confirmPasswordRegister = $("#confirmPasswordRegister").val()
    var nameRegister = $("#nameRegister").val()
    var genderRegister = $("#genderRegister").val()
    var birthdayRegister = $("#birthdayRegister").val()
    var phoneNumberRegister = $("#phoneNumberRegister").val()
    var statusRegister = $("#statusRegister").val()
    if (passwordRegister !== confirmPasswordRegister) {
        $("#messageConfirmPass").text('Hai mật khẩu khác nhau!');
    } else if (confirm) {
        $.ajax({
            type: 'post',
            url: 'http://localhost:8000/create-user',
            data: {
                'username': usernameRegister,
                'password': passwordRegister,
                'email': emailRegister,
                'name': nameRegister,
                'gender': genderRegister,
                'birthday': birthdayRegister,
                'phone_number': phoneNumberRegister,
                'status': statusRegister
            },
            dataType: 'json',
            success: function (result) {
                if (result.status == 201) {
                    $("#labelNotiModel").text('Tạo biên tập viên thành công!');
                    $("#notification-modal").modal('show');
                } else {
                    $("#labelNotiModel").text('Tạo biên tập viên thất bại!');
                    $("#notification-modal").modal('show');
                }
            }
        });
    } else {

    }
});
$("#confirmPasswordRegister").focusout(function () {
    var passwordRegister = $("#passwordRegister").val()
    var confirmPasswordRegister = $("#confirmPasswordRegister").val()
    if (passwordRegister !== confirmPasswordRegister) {
        $('#messageConfirmPass').css('color', 'red');
        $("#messageConfirmPass").text('Hai mật khẩu khác nhau!');
        $('#modal-btn-ok-create').attr('disabled', 'disabled');
    } else {
        $('#messageConfirmPass').css('color', 'green');
        $("#messageConfirmPass").text('');
        $('#modal-btn-ok-create').removeAttr('disabled');
    }
});
$("#usernameRegister").focusout(function () {
    let username = $("#usernameRegister").val()
    checkUserExist(username, function (result) {
        if (result == false) {
            $('#messageUsername').css('color', 'red');
            $('#messageUsername').text('Tài khoản đã tồn tại!');
            $('#modal-btn-ok-create').attr('disabled', 'disabled');
        } else {
            $('#messageUsername').css('color', 'green');
            $('#messageUsername').text('Tài khoản có thể sử dụng!');
            $('#modal-btn-ok-create').removeAttr('disabled');
        }
    });
});
$("#emailRegister").focusout(function () {
    let email = $("#emailRegister").val()
    checkEmailExist(email, function (result) {
        if (result == false) {
            $('#messageEmail').css('color', 'red');
            $('#messageEmail').text('Email đã tồn tại!');
            $('#modal-btn-ok-create').attr('disabled', 'disabled');
        } else {
            $('#messageEmail').hide();
            $('#modal-btn-ok-create').removeAttr('disabled');
        }
    })
});