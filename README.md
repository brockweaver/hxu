# hxu - Utility for Helix

This is an open source (MIT licensed) utility for more easily consuming resources from Q2's Helix product.  It currently focuses on batch files -- both parsing what you can download from Helix and generating what you need to upload to Helix.

Planned future features:
- sftp sync capability to automatically pull down expected files to a local repository
- event webhooks to handling consuming events off Azure Service Bus and calling a REST-friendly webhook you implement

## Commands
`hxu file parse` - processes stdin stream as one of the fixed-width files emitted by Helix. Will emit a jsonlines-formatted stream to stdout. Auto detects file content based on file header line and parses appropriately.  Options exist to override this default behavior. Pass the `--help` flag to see all of them.

`hxu file generate` - processes stdin stream and emits a valid file for uploading to Helix.  Several options exist, run `hxu help file generate` for details.

### Options
running `hxu help file parse` or `hxu file parse --help`:
```
Converts native fixed-width Helix files to different formats

Usage:
  hxu file parse [inputFile] [outputFile] [flags]

Flags:
      --format string   format [jsonlines,tsv,csv,verifyonly] (default "jsonlines")
      --help            display help content
      --noheader        suppress emitting the header line
      --preserve        preserve exact original content sans insignificant whitespace; do not apply conversions
      --strict          error if unexpected bytes appear in content lines
      --verifyonly      only verifies the file; does not emit content. synonym for --format=verifyonly
      --windows         use windows-style line endings (\r\n) instead of unix-style (\n)
```

Note the type of the file (account balance, event notification, etc) is autodetected by inspecting the Header line present in every file emitted by Helix.

The `--preserve` flag is to disable the auto-conversion of fields from string to their appropriate types and format.  For example, in the account balance file the `accountBalance` field as it appears in the file is `000000000001234`. By default, this is emitted by `hxu` as `12.34`, which is its semantic meaning.  To disable this conversion, specify the `--preserve` flag and it will emit from `hxu` as the original value, `000000000001234`. 

The `file` command does some implicit validation to ensure the file is valid:
- Ensures all bytes for all fields on every line exist (properly ignores extra fields)

If you want to perform this validation without converting and emitting file content, pass the `--format=verifyonly` flag
If you want additional validation, pass the `--strict` flag.

If no arguments are specified (i.e. missing `inputfile` and `outputfile`), reads from `stdin` and writes to `stdout`.

`hxu help` - displays helpful information for using non-default behavior of hxu. e.g. arguments and flags for each individual command.

### Parse Example
With the [example ach transaction file](https://github.com/q2baas/SFTPFileExamples/blob/prod/202003011500_ACHTRANSACTION.TXT) downloaded from [Helix's API Site](https://docs.helix.q2.com/docs), run `hxu` with the following command:

> `hxu file parse achtransaction.txt`

It will emit a `jsonlines` formatted file like this:
```json
{"recordType":"H","fileName":"202003161500_ACHTRANSACTION.TXT","recordCount":1,"fileCreatedDate":"2020-03-16T15:00:11992-05:00","fileEffectiveDate":"2020-03-15T23:59:59.000-05:00"}
{"customerId":1062592,"customerTag":"e5f82d07-ae0f-4fc4-bc92-825d010eae23","accountId":1062595 "accountTag":"f44c6019-5f3b-4b7f-80c2-dc73b04e7f3c","accountName":"My Savings Account","transactionId":90741777 "transactionTypeCode":"UNKDEP","traceNumber":"071108830000005","standardEntryClassCode":"PPD","companyName":"ACH File Name","companyDiscretionaryData":"123456890","companyEntryDescription":"PAID","receivingCompanyName":"ACH SAMPLE" "identificationNumber":"001002005861076"}
{"customerId":1062410,"customerTag":"e5f2faf9-a5a5-bdc8-7791-151a44bb33b9","accountId":1062414,"accountTag":"a2148bb3-b3b5-c110-84a4-06715eae41fc","accountName":"My Savings Account","transactionId":90741779,"transactionTypeCode":"CPDEP","traceNumber":"071108830000006","standardEntryClassCode":"PPD","companyName":"ACH File Name","companyDiscretionaryData":"Discretionary Data","companyEntryDescription":"PAID","receivingCompanyName":"ACH SAMPLE","identificationNumber":"000001205139352"}
```

Making this a bit easier to read:
> Header:
```json
{
    "recordType": "H",
    "fileName": "202003161500_ACHTRANSACTION.TXT",
    "recordCount": 1,
    "fileCreatedDate": "2020-03-16T15:00:11992-05:00",
    "fileEffectiveDate": "2020-03-15T23:59:59.000-05:00"
}
```

> Content Line 1:
```json
{
    "customerId": 1062592,
    "customerTag": "e5f82d07-ae0f-4fc4-bc92-825d010eae23",
    "accountId": 1062595 "accountTag": "f44c6019-5f3b-4b7f-80c2-dc73b04e7f3c",
    "accountName": "My Savings Account",
    "transactionId": 90741777 "transactionTypeCode": "UNKDEP",
    "traceNumber": "071108830000005",
    "standardEntryClassCode": "PPD",
    "companyName": "ACH File Name",
    "companyDiscretionaryData": "123456890",
    "companyEntryDescription": "PAID",
    "receivingCompanyName": "ACH SAMPLE",
    "identificationNumber": "001002005861076"
}
```

> Content Line 2:
```json
{
    "customerId": 1062410,
    "customerTag": "e5f2faf9-a5a5-bdc8-7791-151a44bb33b9",
    "accountId": 1062414,
    "accountTag": "a2148bb3-b3b5-c110-84a4-06715eae41fc",
    "accountName": "My Savings Account",
    "transactionId": 90741779,
    "transactionTypeCode": "CPDEP",
    "traceNumber": "071108830000006",
    "standardEntryClassCode": "PPD",
    "companyName": "ACH File Name",
    "companyDiscretionaryData": "Discretionary Data",
    "companyEntryDescription": "PAID",
    "receivingCompanyName": "ACH SAMPLE",
    "identificationNumber": "000001205139352"
}
```

Note it did not need to be told the file is an ACH transaction file; it auto-detected the file and parsed appropriately. This is true for all files Helix emits.  Also, you can output it in different formats. Suppose you wanted a comma-delimited file:

> `hxu file parse achtransaction.txt --format=csv`

It will emit a `csv` formatted file like this:
```
CustomerId,CustomerTag,AccountId,AccountTag,AccountName,TransactionId,TransactionTag,TransactionTypeCode,TraceNumber,StandardEntryClassCode,CompanyName,CompanyDiscretionaryData,CompanyEntryDescription,ReceivingCompanyName,IdentificationNumber
1062592,e5f82d07-ae0f-4fc4-bc92-825d010eae23,1062595,f44c6019-5f3b-4b7f-80c2-dc73b04e7f3c,My Savings Account,90741777,,UNKDEP,071108830000005,PPD,ACH File Name,123456890,PAID,ACH SAMPLE,001002005861076
1062410,e5f2faf9-a5a5-bdc8-7791-151a44bb33b9,1062414,a2148bb3-b3b5-c110-84a4-06715eae41fc,My Savings Account,90741779,,CPDEP,071108830000006,PPD,ACH File Name,Discretionary Data,PAID,ACH SAMPLE,000001205139352
```

### Caveats
jsonlines-formatted files are, by definition, a single json object per line.  json does not enforce any kind of ordering of properties on an object when serializing, so json properties may be in different order from one line to the next.  Fields with no significant characters (e.g. empty string or all whitespace) will be omitted to conserve file size.

The intent of this conversion utility is not necessarily to have smaller sized files; it is to create files that are more easily consumed. However, jsonlines-formatted files are on average about 1/2 the original file size.  CSV and TSV files are roughly about 1/6 the original file size.

jsonlines-formatted files retain the original header information as its first line, and as such, has a different schema than all remaining lines in the file.  CSV- and TSV-formatted files do not retain original header information; it is lost in the conversion. The first line of CSV- and TSV-formatted files are the individual field names.  The header line can be suppressed for any output format using the `--noheader` flag.

This is an open source project done in my free time. I'd love to hear any feedback you may have, good or bad.