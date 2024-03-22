package frontend

const Templ = `
<!DOCTYPE html>
<html>
    <head>
        <script src="http://ajax.googleapis.com/ajax/libs/jquery/1/jquery.js"></script>
		<script type="text/javascript" src="https://code.jquery.com/jquery-2.1.3.min.js"></script>
        <script src="/static/galleria/galleria-1.4.2.min.js"></script>
		<script type="text/javascript">
		
		var observe;
		if (window.attachEvent) {
		    observe = function (element, event, handler) {
		        element.attachEvent('on'+event, handler);
		    };
		}
		else {
		    observe = function (element, event, handler) {
		        element.addEventListener(event, handler, false);
		    };
		}
		function init () {
		    var text = document.getElementById('text');
			var x = document.getElementsByTagName("commentArea");
		console.log(x);
		    function resize () {
		        text.style.height = '28px';
		        text.style.height = text.scrollHeight+'px';
		    }
		    /* 0-timeout to get the already changed text */
		    function delayedResize () {
		        window.setTimeout(resize, 0);
		    }
		    observe(text, 'change',  resize);
		    observe(text, 'cut',     delayedResize);
		    observe(text, 'paste',   delayedResize);
		    observe(text, 'drop',    delayedResize);
		    observe(text, 'keydown', delayedResize);
		
		    text.focus();
		    text.select();
		    resize();
		}
		</script>
		<script src="/static/js/test6.js"></script>
    </head>
    <body onload="init();">
		<div class="toptag">
				<h1 id="efi"><br>Recipe sharing platform</h1>
		</div>
		
		<p style="margin-left:10%; margin-top:20px; height: 6px;">Hello <b id="user">{{.Username}}</b></p>
		<b><p id="currentImage" name="currentImage">
			
		</p></b>
		
		<div class="galleriaContainer">
	        <div class="galleria">
			{{range .Imm}}
			<a href="/static/server/scaledloc/{{.}}" download="{{.}}"><img src="/static/server/scaledloc/{{.}}" /></a>
	         {{end}}   
	            
	        </div>
		</div>
		
		<script>
		Galleria.loadTheme('/static/galleria/themes/classic/galleria.classic.min.js');
            Galleria.run('.galleria');
		
		Galleria.ready(function(options) {

    // 'this' is the gallery instance
    // 'options' is the gallery options



    	this.bind('image', function(e) {
   		console.log('Now viewing ' + e.imageTarget.src);
		var y = e.imageTarget.src ;
		n=y.lastIndexOf("/");
		var name = y.slice(n+1);
		document.getElementById("currentImage").innerHTML = name;
		
		var source=document.getElementById("currentImage").innerHTML;
		
		var dest=document.getElementById("tesla");
		dest.value=source;
		
		function progressHandlingFunction(evt) {
		    if (evt.lengthComputable) {
		    
		    } else {
		        alert('error buddy');
		    }
		}	
			
		var df = new FormData(document.getElementById("id-commentdata"));
		
			$.ajax({
			  url: "/readcmntfromdb/",
			  type: "POST",
			  data: df,
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
				document.getElementById("commentArea").innerHTML = data ;
				document.getElementById("text").value = "";
				
			}).fail(function(data){
				alert("failure");
							});	
    });
});
		</script>
		<div class="rhs">
			<div class="commentArea" id="commentArea">
		
			</div>

			<form action="/writecmnttodb/" method="POST" enctype="multipart/form-data"  id="id-commentdata" name="id-commentdata">
				<input type="hidden" id="tesla" name="tesla" />	
									
				<div class="commentbox" onload="this.style.height='28px'" >
					<textarea  style="width:100%;height:28px;" id="text" name="currentComment"></textarea>
				</div>
				
			</form>
	</div>
	
	<style>
	#currentImage{width:50% ; height: 6px; margin-left : 10% ; margin-top : 26px ;}
	.commentbox{  height:20% ;  margin-left : 2% ; margin-top : 0px ; margin-right:2%; }
	.commentbox > textarea{width:inherit;}
    .galleria{ width: 100%; height: 80%;   margin-top : 0% ;float:left;}
	.galleriaContainer{width: 50%; height: 600px;  margin-left : 10% ; margin-top : 37px ;float:left;}
	.commentArea {   height:92%  ; margin:2% ; background-color: black; opacity:0.5; overflow: auto; color: white;
 }
	.rhs { width:30%; height:480px  ; margin-left : 60% ;margin-top : 37px ;  
	}
	body {
				background-image: url("/static/css/bg.jpg");
			    background-color: #cccccc;
				}
	#efi {
				color : white ;
				}
			.toptag {
				margin-top : 0px ;
				margin-left : 0px ;
				margin-right : 0px ;
				height : 90px ;
				background-color : black ;
				opacity : 0.95;
				}
		</style>
	<style type="text/css">
		textarea {
		    border: 1px solid grey;
			border-spacing: 10px;
		    overflow: hidden;
		    padding: 0;
		    outline: none;
		    background-color: #D0D0D0;
		    resize: none;
			
			max-height : 50px;
			overflow-y:initial;
			width=100%;
		}
		hr{ opacity:0.5; color:}
</style>	
        
	<style>
						
					html, body, div, span, applet, object, iframe,textarea
					h1, h2, h4, h5, h6, p, blockquote, pre,
					a, abbr, acronym, address, big, cite, code,
					del, dfn, em, img, ins, kbd, q, s, samp,
					small, strike, strong, sub, sup, tt, var,
					 u, i, center,
					dl, dt, dd, ol, ul, li,
					fieldset, form, label, legend,
					table, caption, tbody, tfoot, thead, tr, th, td,
					article, aside, canvas, details, embed, 
					figure, figcaption, footer, header, hgroup, 
					menu, nav, output, ruby, section, summary,
					time, mark, audio, video {
						margin: 0;
						padding: 0;
						border: 0;
						-webkit-margin-before: 0 em;
  						-webkit-margin-after: 0 em;
						-webkit-padding-start: 0px;
						vertical-align: baseline;
					}
					h1 {
					  display: block;
					  font-size: 2em;
					  -webkit-margin-before: 0em;
					  -webkit-margin-after: 0em;
					  -webkit-margin-start: 0px;
					  -webkit-margin-end: 0px;
					  font-weight: bold;
					}
					/* HTML5 display-role reset for older browsers */
					article, aside, details, figcaption, figure, 
					footer, header, hgroup, menu, nav, section {
						display: block;
					}
					body {
						
					}
					ol, ul {
						list-style: none;
					}
					blockquote, q {
						quotes: none;
					}
					blockquote:before, blockquote:after,
					q:before, q:after {
						content: '';
						content: none;
					}
					table {
						border-collapse: collapse;
						border-spacing: 0;
					}
					
					inline style {
						height : 28px;
						width : 100%;
						}
		</style>
	
    </body>
</html>
`
