$(document).ready(function() {
    $('#editorsTable').DataTable();
    $('#modal-btn-ok-noti').click(function(){
      $('#notification-modal').modal('toggle');
      location.reload()
    });
});
var idPost
var modalBrowserPost = function(callback){
  $(".btn-browser-post").on("click", function(){
    idPost = this.id.substring(this.id.indexOf("-")+1)
    $("#active-modal").modal('show');
  });
  $("#modal-btn-ok-active").on("click", function(){
    callback(true);
    $("#active-modal").modal('hide');
  });
  
  $("#modal-btn-back-active").on("click", function(){
    callback(false);
    $("#active-modal").modal('hide');
  });
}; 
modalBrowserPost(function(confirm){
  if(confirm){
      $.ajax({
    type:'post',
    url:'http://localhost:8000/update-status-post',
    data:{'id':idPost},
    dataType:'json',
    success: function(result){
      if(result.status == 200){
        $('#labelNotiModel').text('Duyệt thành công!');
        $('#notification-modal').modal('show');
      } else {
        $('#labelNotiModel').text('Duyệt thất bại!');
        $('#notification-modal').modal('show');
      }
    }
  });
  }else{
      
  }
});