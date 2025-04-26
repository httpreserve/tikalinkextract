<!--markdownlint-disable-->

<div>
<p align="center">
<img height="300px" width="300px" id="logo" src="https://github.com/httpreserve/httpreserve/raw/main/src/images/httpreserve-logo.png" alt="httpreserve"/>
</p>
</div>

<!--markdownlint-enable-->

# tikalinkextract

Tika client for httpreserve.

## About

Tikalinkextract requires users start the Tika HTTP server, and then it provides
a way for them to automate the batch processing of those files into its text
extraction mechanism. The text is then processed to look for hyperlinks which
are extracted and output to stdout. There are examples you can try below.

More information is available on the OPF website:
[Hyperlinks in your files? How to get them out using tikalinkextract][opf-1]

[opf-1]: https://openpreservation.org/blogs/hyperlinks-in-your-files-how-to-get-them-out-using-tikalinkextract/

## Demo

[![asciicast](https://asciinema.org/a/143271.png)](https://asciinema.org/a/143271)

## Use with Wget

### Extract the links from your files using seeds option

```sh
./tikalinkextract -seeds -file archives-nz-demo/ > transferlinks.txt
```

### Use the seeds to generate a warc file

<!--markdownlint-disable-->

```sh
wget -T 10 --tries=1 --page-requisites --span-hosts --convert-links  --execute robots=off --adjust-extension --no-directories --directory-prefix=output --warc-cdx --warc-file=accession --wait=0.1 --user-agent=httpreserve-wget/0.0.1 -i transferlinks.txt
```

See [explainshell.com][explain-1]

[explain-1]: https://explainshell.com/explain?cmd=wget+-T+10+--tries%3D1+--page-requisites+--span-hosts+--convert-links++--execute+robots%3Doff+--adjust-extension+--no-directories+--directory-prefix%3Doutput+--warc-cdx+--warc-file%3Daccession+--wait%3D0.1+--user-agent%3Dhttpreserve-wget%2F0.0.1+-i+transferlinks.txt

<!--markdownlint-enable-->

## Resources that might be useful

* [REGEX Guru: Detecting URLS in text][regex-1]

[regex-1]: http://www.regexguru.com/2008/11/detecting-urls-in-a-block-of-text/

## License

Tika is licensed as [Apache License 2.0][tika-license].

This tool is licensed [GNU General Public License Version 3](LICENSE).

[tika-license]: http://www.apache.org/licenses/
