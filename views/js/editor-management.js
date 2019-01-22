$(document).ready( function () {
    $('#editorsTable').DataTable();
    $('#modal-btn-ok-noti').click(function(){
      $('#notification-modal').modal('toggle');
      location.reload()
    });
});
var idActive
var idDeactive
var idDelete
var modalConfirmActive = function(callback){
    $(".btn-confirm-active").on("click", function(){
      idActive = this.id.substring(this.id.indexOf("-")+1)
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
modalConfirmActive(function(confirm){
    if(confirm){
        $.ajax({
			type:'post',
			url:'http://localhost:8000/active-editor',
			data:{'id':idActive},
			dataType:'json',
			success: function(result){
				if(result.status == 200){
				  $('#labelNotiModel').text('Kích hoạt thành công!');
				  $('#notification-modal').modal('show');
				} else {
          $('#labelNotiModel').text('Kích hoạt thất bại!');
				  $('#notification-modal').modal('show');
				}
			}
		});
    }else{
        
    }
});
var modalConfirmDeactive = function(callback){
  $(".btn-confirm-deactive").on("click", function(){
    idDeactive = this.id.substring(this.id.indexOf("-")+1)
    $("#deactive-modal").modal('show');
  });
  $("#modal-btn-ok-deactive").on("click", function(){
    callback(true);
    $("#deactive-modal").modal('hide');
  });
  
  $("#modal-btn-back").on("click", function(){
    callback(false);
    $("#deactive-modal").modal('hide');
  });
}; 
modalConfirmDeactive(function(confirm){
  if(confirm){
      $.ajax({
    type:'post',
    url:'http://localhost:8000/deactive-editor',
    data:{'id':idDeactive},
    dataType:'json',
    success: function(result){
      if(result.status == 200){
        $('#labelNotiModel').text('Bỏ kích hoạt thành công!');
        $('#notification-modal').modal('show');
      } else {
        $('#labelNotiModel').text('Bỏ kích hoạt thất bại!');
        $('#notification-modal').modal('show');
      }
    }
  });
  }else{

  }
});
var modalConfirmDelete = function(callback){
    $(".btn-confirm-delete").on("click", function(){
        idDelete = this.id.substring(this.id.indexOf("-")+1)
        $("#delete-modal").modal('show');
      });
    $("#modal-btn-ok").on("click", function(){
      callback(true);
      $("#delete-modal").modal('hide');
    });
    
    $("#modal-btn-back").on("click", function(){
      callback(false);
      $("#delete-modal").modal('hide');
    });
  }; 
modalConfirmDelete(function(confirm){
    if(confirm){
        $.ajax({
			type:'post',
			url:'http://localhost:8000/delete-user',
			data:{'id':idDelete},
			dataType:'json',
			success: function(result){
				if(result.status == 200){
          $("#labelNotiModel").text('Xóa thành công!');
          $("#notification-modal").modal('show');
				} else {
					$("#labelNotiModel").text('Xóa thất bại!');
          $("#notification-modal").modal('show');
				}
			}
		});
    }else{
        // $("#result").html("NO CONFIRMADO");
    }
});