$(document).ready(function () {
    $('#modal-btn-ok-noti').click(function () {
        $('#notification-modal').modal('toggle');
        location.reload()
    });
})
CKEDITOR.replace('content-area');
var idPost, title, topic, description, content;
var modalConfirmActive = function (callback) {
    $(".btn-update-post").on("click", function () {
        idPost = this.id.substring(this.id.indexOf("-") + 1)
        title = $('#title').val()
        topic = $('#topic').val()
        description = $('#description').val();
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
modalConfirmActive(function (confirm) {
    if (confirm) {
        $.ajax({
            type: 'post',
            url: 'http://localhost:8000/update-content-post',
            data: { 
                'id': idPost, 
                'title': title, 
                'topic': topic, 
                'description': description, 
                'content': content },
            dataType: 'json',
            success: function (result) {
                if (result.status == 200) {
                    $('#labelNotiModel').text('Sửa thành công!');
                    $('#notification-modal').modal('show');
                    ResetValue();
                }

            },
            error: function (result) {
                if (result.status == 400) {
                    $('#labelNotiModel').text('Sửa thất bại!');
                    $('#notification-modal').modal('show');
                }
            },

        });
    } 
});
