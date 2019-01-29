$(document).ready(function () {
    $('#editorsTable').DataTable();
    $('#modal-btn-ok-noti').click(function () {
        $('#notification-modal').modal('toggle');
        location.reload()
    });
});
var idPost
var idDelete
var idDeActive
var modalBrowserPost = function (callback) {
    $(".btn-browser-post").on("click", function () {
        idPost = this.id.substring(this.id.indexOf("-") + 1)
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
modalBrowserPost(function (confirm) {
    let comment = $("#comment-active").val()
    if (confirm) {
        $.ajax({
            type: 'post',
            url: 'http://localhost:8000/active-status-post',
            data: {
                'id': idPost,
                'comment': comment
            },
            dataType: 'json',
            success: function (result) {
                if (result.status == 200) {
                    $('#labelNotiModel').text('Duyệt thành công!');
                    $('#notification-modal').modal('show');
                } else {
                    $('#labelNotiModel').text('Duyệt thất bại!');
                    $('#notification-modal').modal('show');
                }
            }
        });
    } else {

    }
});
var modalConfirmDelete = function (callback) {
    $(".btn-delete-post").on("click", function () {
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
            url: 'http://localhost:8000/delete-post',
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
        // $("#result").html("NO CONFIRMADO");
    }
});
var modalDeActivePost = function (callback) {
    $(".btn-not-browser-post").on("click", function () {
        idDeActive = this.id.substring(this.id.indexOf("-") + 1)
        $("#deactive-modal").modal('show');
    });
    $("#modal-btn-ok-deactive").on("click", function () {
        callback(true);
        $("#deactive-modal").modal('hide');
    });

    $("#modal-btn-back").on("click", function () {
        callback(false);
        $("#deactive-modal").modal('hide');
    });
};
modalDeActivePost(function (confirm) {
    let comment = $("#comment-deactive").val()
    if (confirm) {
        $.ajax({
            type: 'post',
            url: 'http://localhost:8000/deactive-status-post',
            data: {
                'id': idDeActive,
                'comment': comment
            },
            dataType: 'json',
            success: function (result) {
                if (result.status == 200) {
                    $('#labelNotiModel').text('Bỏ duyệt thành công!');
                    $('#notification-modal').modal('show');
                } else {
                    $('#labelNotiModel').text('Bỏ duyệt thất bại!');
                    $('#notification-modal').modal('show');
                }
            }

        });
    } else {

    }
});
$('.btn-update-post').click(function () {
    idPost = this.id.substring(this.id.indexOf("-") + 1)
    document.location.href = "/render-update-post/" + idPost;
})

var modalDetailPost = function (callback) {
    $(".btn-detail-post").on("click", function () {

        idPost = this.id.substring(this.id.indexOf("-") + 1)
        $("#detail-post-modal").modal('show');

        callback(true);
    });

};
// editor = CKEDITOR.replace('content-area');
modalDetailPost(function (confirm) {

    if (confirm) {
        $.ajax({
            type: 'get',
            url: 'http://localhost:8000/render-detail-post/' + idPost,
            dataType: 'json',

            success: function (result) {
                if (result.status == 200) {
                    $('#detail-title').val(result.post["title"]);
                    $('#detail-topic').val(result.post["topic"]);
                    $('#detail-description').val(result.post["description"]);
                    // console.log(result.post["content"]);

                    $('#detail-content').html(result.post["content"]);
                    countCurrentComment = $("#detail-post-modal").find('[name = comment-detail]');
                    for (var i = 0; i < countCurrentComment.length; i++) {

                        $("#detail-post-comment").find('[name = comment-detail]').remove();
                    }
                    for (var i = 0; i < result.comment.length; i++) {
                        var detail_comment = "<div style='background:#EEEEEE; border: 1px solid gray'  name='comment-detail' class='comment-detail-" + i + " form-group'><label>Nhận xét: " + (i + 1) + "</label><div class='content-comment'>Nội dung nhận xét: "+result.comment[i]["message"] +"</div><div class='time-created-comment'>Thời gian nhận xét: "+result.comment[i]["created_at"]+"</div><div class='creator-comment'>Người nhận xét: "+result.username[i]+"</div></div>"
                        //var detail_comment = ("<div name='comment-detail' class='comment-detail-" + i + " form-group'><label>Comment " + (i + 1) + "</label><input disabled type='comment' name='comment' id='comment' tabindex='1' class='form-control ' value = " + "'" + result.comment[i]["message"] + " - " + result.comment[i]["created_at"] + " - " + result.username[i] + "'" + " required /></div>");
                        $('#detail-post-comment').append(detail_comment);
                    }
                } else {

                }
            }
        });
    } else {

    }
});


