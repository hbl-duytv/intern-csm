$(document).ready(function () {
    var creator, title, topic, description, content;
    $('.btn-create-post').click(function () {
        if (confirm("Ban chac chan muon them post nay")) {
            creator = this.id.substring(this.id.indexOf("-") + 1);
            console.log("id", creator);
            title = $('.title').val();
            console.log(title);
            topic = $('.topic').val();
            description = $('.description').val();
            content = $('.content-area').val();
            Post();
        }
        else {

        }
    })
    function Post() {
        console.log(creator, title, topic, description, content);
        $.ajax({
            type: 'post',
            url: 'http://localhost:8000/create-post',
            data: { 'creator': creator, 'title': title, 'topic': topic, 'description': description, 'content': content },
            dataType: 'json',
            success: function () {
                alert("them thanh cong");
            },
            // error: function (response) {
            //     if (response.status == 200) {
            //         alert("them thanh cong error");
            //         ResetVAlue()
            //     } else {
            //         alert("them that bai");
            //     }

            // }


        });
    }
    function ResetVAlue() {
        title = $('.title').val("");
        
        topic = $('.topic').val("");
        description = $('.description').val("");
        content = $('.content-area').val("");
    }

});
