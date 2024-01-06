# wayback-keyword-search 

[IMPORTANT NOTE: WAYBACK IS NOW RATE LIMITING. I HAD TO UPDATE THE TOOL TO GET IT WORKING; HOWEVER IT IS SLOWER NOW AS IT CAN'T USE PARALLELISM ANYMORE. THE PYTHON VERSION IS UPDATED; THE GO VERSION IS NOT (YET)]

This tools downloads each page from the Wayback Machine for a specific input domain and saves each page as a local .txt file, so that you can later search for keyword matches within the saved files.

downloading is done with the "download" file; and searching with the "search" file.

You can download pages saved in specific years (i.e.: 2020), or years and months (i.e.: 202001), or years and months and days (i.e: 20200101), just specifying the date format in the prompt. If you want to download everything in the 2000's or 19**'s regardless the saved date, just type "2" (for the pages saved past 2000) or "1" (for the pages saved in the XXth century) in the prompt, and Wayback will save each page saved matching that criteria. So, if you want to save a website that has been saved across 1999 and 2000, you will need to run the tool twice.

There is a Python3 version and a Go version.

--------------------------

[*] Python usage:

> python3 download.py > specify your domain like: nytimes.com (no quotes!)

When the download is completed, a directory named as the domain will be saved in the local path.

So you can search for keyword matches within each file in the local dir using the "search.py" file:

> python3 search.py > specify your keyword (no quotes!).

--------------------------

[*] Go usage: [NOT WORKING NOW DUE TO ARCHIVE BLOCKING TOO MANY PARALLEL REQUESTS]

> go run download.go

and then:

> go run search.go

The best way to use the Go version is by running the compiled executables:

go build search.go

go build download.go

Notice that the Go version also features a download_channels.go version (thanks to Stephen Paulger for such improvement) which is a bit more efficient. Consider testing both!

