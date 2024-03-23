function progressHandlingFunction(evt) {
    if (evt.lengthComputable) {
        var percentComplete = evt.loaded / evt.total;
		
		 $("#id-progress-bar").show();
	    $("#id-upload-progress").attr("style","width:50" + parseInt(percentComplete * 100) + "%");
		var temp = parseInt(percentComplete * 100);
		document.getElementById("id-percentile").innerHTML = temp + "% Completed" ;
  	    
    } else {
        alert('error buddy');
    }
}
	
$(function(){
	alert("on load");
	var bar = $('.bar');
	var percent = $('.percent');
	var status = $('#status');
	$('#id-file-d').change(function(){
			
			var fd = new FormData(document.getElementById("id-filedata"));	
			
			var filename = $("#id-file-d").val();
			
        // Use a regular expression to trim everything before final dot
        var extension = filename.replace(/^.*\./, '');

        // Iff there is no dot anywhere in filename, we would have extension == filename,
        // so we account for this possibility now
        if (extension == filename) {
            extension = '';
        } else {
            // if there is an extension, we convert to lower case
            // (N.B. this conversion will not effect the value of the extension
            // on the file upload.)
            extension = extension.toLowerCase();
        }

        switch (extension) {
            case 'jpg':
            case 'jpeg':
            case 'png':
                //alert("it's got an extension which suggests it's a PNG or JPG image (but N.B. that's only its name, so let's be sure that we, say, check the mime-type server-side!)");

            // uncomment the next line to allow the form to submitted in this case:
         break;

            default:
                // Cancel the form submission
				alert('invalid filetype');
				document.location.reload(true);
                submitEvent.preventDefault();
				
        }
	console.log(fd);

			
			$.ajax({
			  url: $("#id-filedata").attr("action"),
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
				alert('successful');
				document.location.reload(true);
					}).fail(function(data){
				alert("No data available!");
				document.location.reload(true);
							});
			
			
			
			
	
	
	})
	
	
   

});