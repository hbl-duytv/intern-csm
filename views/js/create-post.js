$(document).ready(function () {
    $('#modal-btn-ok-noti').click(function () {

        $('#notification-modal').modal('toggle');
        location.reload()
    });
});
CKEDITOR.replace('content-area');
var creator, title, topic, description, content;

var modalCreatePost = function (callback) {
    $(".btn-create-post").on("click", function () {
        creator = this.id.substring(this.id.indexOf("-") + 1);
        console.log("id", creator);
        title = $('.title').val();
        console.log(title);
        topic = $('.topic').val();
        description = $('.description').val();
        content = CKEDITOR.instances['content-area'].getData();
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
modalCreatePost(function (confirm) {
    if (confirm) {
        $.ajax({
            type: 'post',
            url: 'http://localhost:8000/create-post',
            data: { 'creator': creator, 'title': title, 'topic': topic, 'description': description, 'content': content },
            dataType: 'json',
            success: function (result) {
                if (result.status == 201) {
                    $('#labelNotiModel').text('Tạo thành công!');
                    $('#notification-modal').modal('show');
                    ResetValue();
                }

            },
            error: function (result) {
                if (result.status == 400) {
                    $('#labelNotiModel').text('Tạo thất bại!');
                    $('#notification-modal').modal('show');
                }
            }
        });
    } else {

    }
});
function ResetValue() {
    $('.title').val("");

    $('.topic').val("");
    $('.description').val("");
    CKEDITOR.instances['content-area'].setData("");
    
}


