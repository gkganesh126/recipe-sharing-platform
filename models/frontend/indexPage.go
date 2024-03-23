package frontend

const IndexPage = `
<html >
	<head>
	<title>Recipe Sharing Platform</title>
	<script>
		function pwdcheck() {
			var pwd = document.getElementsByTagName("input");
			var form=document.getElementsByTagName("form");
			if((pwd[0].value == pwd[1].value) || (pwd[0].value==pwd[2].value)) {
					alert('Username and password should not be same');
					form.setAttribute("action","/register/");
			} else if(pwd[1].value.length<6 || pwd[2].value.length<6) {
					alert('password length should not be less than 6');
					form.setAttribute("action","/register/");
			}else if(pwd[1].value != pwd[2].value) {
					alert('Passwords are not same');
					form.setAttribute("action","/register/");
			}else {
					alert('Valid Password');
					form.setAttribute("action","/newlogin/");
			}
		}
		</script>
		<style>
		body {
			margin:0px ;
			background-image: url("static/css/bg.jpg");
			background-repeat: repeat;
		    background-color: #cccccc;
			text-align: center;
			}
		#efi {
			color : white ;
			}
		.toptag {
			margin-top : 0px ;
			margin-left : 0px ;
			margin-right : 0px ;
			height : 90px ;
			font-size: 30px;
			background-color : black ;
			opacity : 0.95;
			text-align:left;
			}
		</style>
	</head>
	<body>
			<div class="toptag">
						<h1 id="efi"><br>Recipe sharing platform</h1>
			</div>
	
		   <div id="content" >
		        <div class="content_item" align='middle'><br><br>
				    <h1>Welcome To Recipe sharing platform</h1> 
						<br><br>
					
						<form method="post" action="/internal">
						    <label for="name">User name</label>
						    <input type="text" id="name" name="name">
						    <label for="password">Password</label>
						    <input type="password" id="password" name="password">
						    <button type="submit">Login</button>
						</form><br>
				</div>
			</div>
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
