# Send NTC an email to report SMS scammers

Automating this [very tedious process](https://region1.ntc.gov.ph/wp-content/uploads/2018/01/COMPLAINTS-ON-TEXT-MESSAGE-TEXT-SPAMTEXT-SCAM-ILLEGALOBSCENE-TEXT-MESSAGE-THREATS-AND-OTHER-RELATED-ITEMS.pdf) of the National Telecommunications Commission (Republic of the Philippines 🇵🇭) to report text spams, text message threats, etc.

This tool allows you to:
* answer the NTC form through a more user friendly interface (Jotforms)
* privately configure an email sending application that sends your complaint to the [NTC](https://ntc.gov.ph/).

## Configuration

1. [Generate an application-specific password for your Gmail](https://devanswers.co/outlook-and-gmail-problem-application-specific-password-required/).
2. Copy `config.sample.yaml` into `config.yaml` and configure your account credentials.
3. Save your government IDs as `govt-id-one.png` and `govt-id-two.png` in this folder.
4. Configure to which email this will send to (e.g. Region 1 if you're in region 1, or to the general consumer email of NTC).


## Usage

Once you've configured it once, you just need to do the following steps to send a report.

1. [Fill up an NTC form](https://form.jotform.com/221512749411046) and download it into this folder.
2. Rename it as `ntc-form.pdf`.
3. Run `go run main.go` - which sends the email from your configured email.
4. 🤔 *Wait for a reply? I don't know what happens next.*

## Limitations

* I use a free version of jotforms, so there's a 100 monthly submissions limit to generate the automatically filled up PDF.
  * There's an option for non-profit orgs to get a 50% discount though, but I'm not exploring that right now.

## License

[MIT](LICENSE.txt).

## References
* [Jotforms](https://www.jotform.com/) for friendly PDF forms.
* [National Telecommunications Commission](https://ntc.gov.ph) of the Republic of the Philippines.

### Personal notes

I coded this from 5:30PM up to 6:45PM of June 1, 2022.
