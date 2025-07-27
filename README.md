# ofx2

**_ofx2_** is a bash script for extracting data from _OFX_ files.

I wrote it to generate _CSV_ files from the transactions in the _OFX_ files
that I download for my bank accounts and credit-cards.

The _**ofx2xml**_ utility remarshals _OFX_ files,
ensuring they are at least version-2 _OFX_ files,
which are valid XML and thereby easier to parse.
The utility also interprets `<INTU.BID>` tags as `<FI><FID>` tag pairs,
as my banks tend to include the former, but not the latter, in their _OFX_ files.



### ofx2xml

**_ofx2_** depends on the utility _**ofx2xml**_,
the source of which was once included here,
but which now has its own repository: <https://github.com/heckman/ofx2xml>.

 _**ofx2xml**_ is used to transmute the provided OFX file to valid XML.

Note that _**ofx2xml**_ is licensed under [GPL2](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html) and not [MIT](https://opensource.org/license/MIT) as is **_ofx2_**.



## Usage

To perform its primary function, run `ofx2 csv -w <ofx-file>`.
This invocation will create a _CSV_ file of the transactions found in the _OFX_ file,
the name of which will be generated from
the financial instutution,
the account number,
and the date range of the transactions.
I've inlcuded a helper-script **_ofx2csv_** which calls **_ofx2_** in this way.

The script includes other options,
and a few other tangentally-related operations.
For complete usage information,
run `ofx2 --help`.



## Installation

I wrote this for _MacOS_,
but it should work on _posix_-like systems.

Both _**ofx2**_ and **_ofx2csv_** are bash scripts
suitable for putting somewhere on your path.
(In fact, ***ofx2csv*** will not work unless ***ofx2*** is on the path.)

***ofx2*** requires the _**ofx2xml**_ utility found here: <https://github.com/heckman/ofx2xml>.
It can be downloaded and installed with ***go*** >1.17:

```shell
go install github.com/heckman/ofx2xml@latest
```

In addition to [_**ofx2xml**_](https://github.com/heckman/ofx2xml), ***ofx2*** requires _**jq**_ and **_xq_**.
Optionally, to generate hashes of transactions, _**mlr**_ is also required.
(It's not needed if the hashing option is never used.)
_**jq**_, **_xq_**, and _**mlr**_ are available via [_homebrew_](https://brew.sh).

## Configuration

Some configuration is required to customize the script
for use with OFX files from specific financial institutions.

In the _ofx-txns_ script, immediately following the usage information,
you'll find the definition of a varable `jq_fidir`.
It contains a json object
that maps `INTU.BID` values
to the names of financial institutions.
Values for `INTU.BID` for institutions in North America
can be found from [intuit.com's list](https://ofx-prod-filist.intuit.com/qm2400/data/fidir.txt).
Add the ones you use to this variable.

## Motivation

In order to track my finances,
I used to download my transactions as CSV files
and then import them into a spreadsheet.

The problem with this approach is that there is
no consistency between the format of the CSV files
from my various financial institutions.

The _ofx2_ script attempts to address this problem
by using the standardized _OFX_ format to generate
a consistent format for the CSV files.
