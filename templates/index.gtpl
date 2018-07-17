<html>
<h2>Blog</h2>

{{range .}}
<h3>{{.Title}}</h3>
<p>{{.Body}}</p>
<a href="/delete/{{.Pid}}">Delete</a>
{{end}}

<h3>New post</h3>

<form action="/posts/new" method="POST">
<p>Title</p>
<input type="text" name="title" /><br><br>
<textarea name="body"></textarea>
<input type="submit">
</form>

</html>