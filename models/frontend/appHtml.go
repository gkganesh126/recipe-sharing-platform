package frontend

const AppHtml = `
<html >
	<head>
		 <title>Recipe Sharing Platform</title>
		<script type="text/javascript" src="https://code.jquery.com/jquery-2.1.3.min.js"></script>
		<script type="text/javascript" src="/static/js/test5.js"></script>
		<link href="/static/css/vendor/bootstrap.min.css" rel="stylesheet">
		<style type="text/css">
					.progress {
					  height: 12px;
					  background: #ebedef;
					  border-radius: 32px;
					  box-shadow: none;
					  }
					.progress-bar {
					  line-height: 12px;
					  background: #1abc9c;
					  box-shadow: none;
					  }
					.progress-bar-success {
					  background-color: #2ecc71;
					  }
					.progress-bar-warning {
					  background-color: #f1c40f;
					  }
					.progress-bar-danger {
					  background-color: #e74c3c;
					  }
					.progress-bar-info {
					  background-color: #3498db;
					  }
					body {
						background-image : url("static/css/bg.jpg") ;
						background-repeat: repeat;
						margin-top :195 px ;
						text-align : center ;
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
						text-align : left ;
						}
		</style>
	</head>
	<body>
		<div id="main">    
			   <div id="content" >
					<div class="content_item" align='middle'><br><br>
						<h3>Hello <b>%s</b>, </h3><br><br>
					  	<h1>Welcome To Recipe Sharing Platform/h1> <br><br>
			      		<p>Browse for the image to upload :</p>	 <br> 
				  					<form action="/upload/" method="POST"  enctype="multipart/form-data"  id="id-filedata" name="id-filedata">
											
											<input type="file" name="id-file-d" id="id-file-d" >
											
									</form>
							 <!--close button_small-->
				  	</div><!--close content_item-->
				</div>
		</div><!--close main-->
		<br>
			<p>
				<div class="progress" id="id-progress-bar" align="center"  style="display:none;">
			        <div class="progress-bar" id="id-upload-progress" align="center" color="black" style="">
								
					</div>
			    </div>
				<div id="id-percentile" align="center"></div>
			</p>
		<style>						
					html, body, div, span, applet, object, iframe,
					h1, h2, h3, h4, h5, h6, p, blockquote, pre,
					a, abbr, acronym, address, big, cite, code,
					del, dfn, em, img, ins, kbd, q, s, samp,
					small, strike, strong, sub, sup, tt, var,
					b, u, i, center,
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
						font-size: 100 %;
						font: inherit;
						vertical-align: baseline;
					}
					/* HTML5 display-role reset for older browsers */
					article, aside, details, figcaption, figure, 
					footer, header, hgroup, menu, nav, section {
						display: block;
					}
					body {
						line-height: 1;
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
		</style>
	</body>
</html>
`
