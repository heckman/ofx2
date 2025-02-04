# ofx2

**_ofx2_** is a bash script for extracting data from _OFX_ files.

I wrote it to generate _CSV_ files from the transactions in the _OFX_ files
that I download for my bank accounts and credit-cards.

It depends on the utility _**ofx2xml**_,
the source for which can be found in the `src/ofx2xml` directory.

The _**ofx2xml**_ utility remarshals _OFX_ files,
ensuring they are at least version-2 _OFX_ files,
which are vavid XML and therefore easier to parse.
The utility also interprets `<INTU.BID>` tags as `<FI><FID>` tag pairs,
as my banks tend to include the former, but not the latter, in their _OFX_ files.

## Usage

To perform its primary function, run `ofx2 csv -w <ofx-file>`.
This invocation will create a _CSV_ file of the transactions found in the _OFX_ file,
the name of which will be generated from
the financial instutution,
the account number,
and the date range of the transactions.
I've inlcuded a helper-script `ofx2csv` which calls **_ofx2_** in this way.

The script includes other options,
and a few other tangentally-related operations.
For complete usage information,
run `ofx2 --help`,
or examine [the source code](./ofx-txns),
where the usage appears near the top,
and which includes a description
of the variables available for custom-formatting the output.

## Installation

I wrote this for _MacOS_,
but it should work on _posix_-like systems.

The script _**ofx2**_
and its helper script _**ofx2csv**_
are both in the root of this repository.

The _**ofx2xml**_ utility is written in _go_.
I needs to be compiled.

The scripts require _**jq**_ and **_xq_** to be installed.
To generate hashes of transactions, _**mlr**_ is also required.
(It's not needed if hashing is never used.)
All three utilities are avaiable via _homebrew_.

The script `dev-bin/install`
builds _**ofx2xml**_, and copies the compiled binary,
along with _**ofx2**_ ann _**ofx2csv**_,
to `$HOME/bin`.
(At least on my machine. YMMV. Don't use it blindly.)

## Configuration

Some configuration is required to customize the script
for use with OFX files from specific financial institutions.

In the _ofx-txns_ script, immediately following the usage information,
you'll find the definition of a varable `jq_fidir`.
It contains a json object
that maps `INTU.BID` values
to the names of financial institutions.
Add the ones you use to this variable.
Values for `INTU.BID` for institutions in North America
can be found from [intuit.com's list](https://ofx-prod-filist.intuit.com/qm2400/data/fidir.txt).

## Development

This project makes use of the _direnv_ utility
to add local directories to the environment's PATH.
It is usefull for putting _ofx2xml_ on the path
before it has been installed,
but you can get along without it.

The `dev-bin` directory contains
scripts to perform common tasks.
The _direnv_ utility will also add it to the PATH.

## Motivation

To track my finances,
I used to download my transactions as CSV files
and then import them into a spreadsheet program.

The problem with this approach is that there is
no consistency between the format of the CSV files
from the various financial institutions.

The _ofx-txns_ script attempts to address this problem.

Most financial institutions
will provide transaction history in OFX format,
although in my experience it tends to be
an older non-XML version of OFX.
Furthermore, in my experience,
they don't include a `<FI>` tag to identify the financial institution,
but instead use a (not-to-spec) `<INTU.BID>` tag.
