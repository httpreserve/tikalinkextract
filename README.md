<div>
<p align="center">
<img height="300px" width="300px" id="logo" src="https://github.com/httpreserve/httpreserve/raw/main/src/images/httpreserve-logo.png" alt="httpreserve"/>
</p>
</div>

# tikalinkextract 

Tika client for httpreserve

## Demo

[![asciicast](https://asciinema.org/a/143271.png)](https://asciinema.org/a/143271)

## Use with Wget

**Extract the links from your files using seeds option**

    ./tikalinkextract -seeds -file archives-nz-demo/ > transferlinks.txt

**Use the seeds to generate a warc file**

    wget --page-requisites --span-hosts --convert-links  --execute robots=off --adjust-extension --no-directories --directory-prefix=output --warc-cdx --warc-file=accession.warc --wait=0.1 --user-agent=httpreserve-wget/0.0.1 -i transferlinks.txt 


### Known Issues

* HTTP links that are formatted in such a way to be split across lines, thus include a newline \n character. 

### Resources that might help 

* [REGEX Guru: Detecting URLS in text](http://www.regexguru.com/2008/11/detecting-urls-in-a-block-of-text/)

## License

Tika is licensed as follows: http://www.apache.org/licenses/

This tool is licensed GNU General Public License Version 3. [Full Text](LICENSE)
