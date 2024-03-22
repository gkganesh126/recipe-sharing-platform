package frontend

const InternalPage = `
<html>
	<head>
		<style>
		body {
			margin:0px ;
			background-image: url("/static/css/bg.jpg");	
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
	</head>
	<body>
			<div class="toptag">
				<img src="/static/js/efi logo.png" style="float:left;" />
				<h1 id="efi"><br>Customer Service Portal</h1>
			</div>
	<br>
			<div class="mid">
				<h1>%s</h1><br>
				<hr>
				<small><b>User:</b> %s</small><br><br>
				<hr>
				<p>What would you like to do?</p><br>
				<a href="/app/">Upload a Picture</a><br><br>
				<a href="/viewimage/">View Pictures</a><br><br>
					<form method="post" action="/logout">
					    <button type="submit">Logout</button>
					</form>
			</div>
			
		<style>
					.mid {
			margin-left: 50px;
			
			margin-top:10px;
			margin-bottom:0px;
			}
						
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
