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
  if (confirm) {
    $.ajax({
      type: 'post',
      url: 'http://localhost:8000/active-status-post',
      data: { 'id': idPost },
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
  if (confirm) {
    $.ajax({
      type: 'post',
      url: 'http://localhost:8000/deactive-status-post',
      data: { 'id': idDeActive },
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
});

var modalDetailPost = function (callback) {
  $(".btn-detail-post").on("click", function () {
    idPost = this.id.substring(this.id.indexOf("-") + 1)
    $("#detail-post-modal").modal('show');
    callback(true);
  });

};
editor = CKEDITOR.replace('content-area');
modalDetailPost(function (confirm) {
  console.log("id" + idPost);
  if (confirm) {
    $.ajax({
      type: 'get',
      url: 'http://localhost:8000/render-detail-post/' + idPost,
      dataType: 'json',

      success: function (result) {
        if (result.status == 200) {
          
          // console.log("data:", result.post);
          // console.log("title", result.post["title"]);
          $('#detail-title').val(result.post["title"]);
          $('#detail-topic').val(result.post["topic"]);
          $('#detail-description').val(result.post["description"]);
          // console.log(result.post["content"]);
          editor.setData(result.post["content"]);
          
        } else {

        }
      }
    });
  } else {

  }
})



