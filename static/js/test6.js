function progressHandlingFunction(evt) {
    if (evt.lengthComputable) {
         
	
    } else {
        alert('error buddy');
    }
}


$(function(){
	
	$('#text').keydown(function (e){
    		if(e.keyCode == 13){
	        
		var fd = new FormData(document.getElementById("id-commentdata"));
		
			$.ajax({
			  url: $("#id-commentdata").attr("action"),
			  type: "POST",
			  data: fd,
			  processData: false,  // tell jQuery not to process the data
			  contentType: false,   // tell jQuery not to set contentType,
			 
				xhr: function(){  // custom xhr
		        	    myXhr = $.ajaxSettings.xhr();
		        	    if(myXhr.upload){ // check if upload property exists
		            		myXhr.upload.addEventListener('progress',progressHandlingFunction, false); // for handling the progress of the upload
		        	    }
		        	    return myXhr;
	    			}			
			}).success(function(data){
				
			var text = document.getElementById("text");
		var existing = document.getElementById("commentArea");
		var username = document.getElementById("user");
		console.log(username);
		existing.innerHTML = existing.innerHTML  + "<h3> " + username.innerHTML + "</h3> " +  text.value + "<br><hr>" ;
		document.getElementById("text").value = "";
		existing.scrollTop = existing.scrollHeight;
				
					}).fail(function(data){
				alert("failure");
				
							});
			
	}
});
	
	
   

});			
