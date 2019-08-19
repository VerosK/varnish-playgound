vcl 4.0;

import directors;

backend working {
  .host = "working:8983";
}
backend broken {
  .host = "broken:8983";
}
sub vcl_init {
    new cluster = directors.round_robin();

    cluster.add_backend(working);
    cluster.add_backend(broken);
}

sub vcl_recv {
    set req.backend_hint = cluster.backend();
    if (req.url ~ "brokenbroken") {
	return (synth(503));
    }
    if (req.url == "/favicon.ico") {
	return (synth(404));
    }
    return (pass); // disable caching
}


sub vcl_deliver {
   if (resp.status == 503) {
	return (restart);
   }
   if (resp.status == 500) {
	return (restart);
   } 
   set resp.http.X-Restarts = req.restarts;
}
