# Body1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Secret ID. | [default to null]
**Name** | **string** | Secret name. | [default to null]
**Target** | **string** | Secret target name. | [optional] [default to null]
**TargetPrivilegeAccount** | **string** | Enable/disable the current secret as the target privilege account. | [optional] [default to null]
**PwdChgUsePolicyDefault** | **string** | Enable/Disable using policy default for Password Changer    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**HeartbeatUsePolicyDefault** | **string** | Enable/Disable using policy default for Password Heartbeat    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**Folder** | **int32** | ID of parent folder. | [default to null]
**Template** | **string** | Secret template name. | [default to null]
**AssociatedSecret** | **int32** | Associated secret setting. | [optional] [default to null]
**SshConnectWith** | **string** | Connect to the target server with current secret or using associate secret.    self:Connect and authenticate to server using current secret.    associate-secret:Connect and authenticate to server using associate secret. | [optional] [default to null]
**SshAutoPassword** | **string** | Automatically send password to server when privileged command (e.g., sudo) in SSH terminal detected.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**ZtnaTagsMatchLogic** | **string** | ZTNA tag matching logic.    or:Match ZTNA tags using a logical OR operator.    and:Match ZTNA tags using a logical AND operator. | [optional] [default to null]
**RsaSignAlgo** | **string** | Signing algorithm of RSA-based publicKey authentication for ssh launcher.    ssh-rsa:RSA publicKey authentication using RSA SHA-1 signing algorithm.    rsa-sha2-256:RSA publicKey authentication using RSA SHA-256 signing algorithm.    rsa-sha2-512:RSA publicKey authentication using RSA SHA-512 signing algorithm. | [optional] [default to null]
**KeyboardLayout** | **string** | Keyboard layout (Applies to web-rdp only).    ar-101:Arabic (101).    ar-102:Arabic (102).    ar-102-azerty:Arabic (102) AZERTY.    can-mul:Canadian Multilingual Standard.    cz:Czech.    cz-qwerty:Czech (QWERTY).    cz-pr:Czech Programmers.    da:Danish.    nl:Dutch.    de:German.    de-ch:German, Switzerland.    de-ibm:German (IBM).    en-uk:English, United Kingdom.    en-uk-ext:English, United Kingdom Extended.    en-us:English, United States.    en-us-dvorak:English, United States-Dvorak.    es:Spanish.    es-var:Spanish Variation.    fi:Finish.    fi-sami:Finnish with Sami.    fr:French.    fr-ca:French, Canada.    fr-ch:French, Switzerland.    fr-be:French, Belgian.    hr:Croatian.    hu:Hungarian.    hu-101:Hungarian 101-Key.    it:Italian.    it-142:Italian (142).    ja:Japanese.    ko:Korean.    lt:Lithuanian.    lt-ibm:Lithuanian IBM.    lt-std:Lithuanian Standard.    lav-std:Latvian (Standard).    lav-leg:Latvian (Legacy).    mk:Macedonian (FYROM).    mk-std:Macedonia (FYROM) - Standard.    no:Norwegian.    no-sami:Norwegian with Sami.    pol-214:Polish (214).    pol-pr:Polish (Programmers).    pt:Portuguese.    pt-br:Portuguese (Brazilian ABNT).    pt-br-abnt2:Portuguese (Brazilian ABNT2).    ru:Russian.    ru-mne:Russian - Mnemonic.    ru-t:Russian (Typewriter).    sl:Slovenian.    sv:Swedish.    sv-sami:Swedish with Sami.    tuk:Turkmen.    tur-f:Turkish F.    tur-q:Turkish Q.    zh-sym-sg-us:Chinese (Simplified, Singapore) - US keyboard.    zh-sym-us:Chinese (Simplified) - US Keyboard.    zh-tr-hk:Chinese (Traditional, Hong Kong S.A.R.).    zh-tr-mo:Chinese (Traditional Macao S.A.R.) - US Keyboard.    zh-tr-us:Chinese (Traditional) - US keyboard. | [optional] [default to null]
**Description** | **string** | Description. | [optional] [default to null]
**PasswordChanger** | **string** | Setting of password changer.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**PasswordHeartbeat** | **string** | Setting of password heartbeat.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**Checkout** | **string** | Setting of checkout.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**Recording** | **string** | Setting of screen recording.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**NeedApproval** | **string** | Setting of approval for launching secret.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**NeedApprovalJob** | **string** | Setting of approval for activating secret job.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**Proxy** | **string** | Setting of proxy for secret.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**SshFilter** | **string** | Setting of if secret need ssh filter profile.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**AvScan** | **string** | Setting of if secret need antivirus profile.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**DlpStatus** | **string** | Setting of if seceret need dlp profile.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**RdpSecurityLevel** | **string** | RDP Security Level    rdp:Standard RDP encryption (web-rdp only).    tls:RDP over TLS.    nla:Network Level Authentication (CredSSP).    best-effort:Best effort (use NLA if server supports, RDP over TLS otherwise). | [optional] [default to null]
**RdpRestrictedAdminMode** | **string** | Credential-less logon over NLA.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**BlockRdpClipboard** | **string** | Setting of whether block clipboard copy &amp; paste of rdp auth proxy.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**PwdChgSchFrequency** | **string** | Password Change Frequency. Repeats Monthly/Weekly/Daily    month:Repeat Monthly    week:Repeat Weekly.    day:Repeat Daily. | [optional] [default to null]
**PwdChgSchInterval** | **int32** | Password Change Interval. Ex. Every x (month/week/day(s), based on frequency) | [optional] [default to null]
**PwdChgSchMonthType** | **string** | Password change schedule Monthly Type.    first:Repeat in the first week of each month.    second:Repeat in the second week of each month.    third:Repeat in the third week of each month.    last:Repeat in the last week of each month.    last-day:Repeat on the last day of each month.    specific:Repeat on a specific day of each month. | [optional] [default to null]
**PwdChgSchMonthDay** | **string** | Password change schedule month days. Repeat Monthly on specified date(s). (May be many)    1:1    2:2    3:3    4:4    5:5    6:6    7:7    8:8    9:9    10:10    11:11    12:12    13:13    14:14    15:15    16:16    17:17    18:18    19:19    20:20    21:21    22:22    23:23    24:24    25:25    26:26    27:27    28:28 | [optional] [default to null]
**PwdChgSchWeekDay** | **string** | Password change schedule week days. (May be many)    sunday:Sunday.    monday:Monday.    tuesday:Tuesday.    wednesday:Wednesday.    thursday:Thursday.    friday:Friday.    saturday:Saturday. | [optional] [default to null]
**PwdChgSchStartDate** | [**time.Time**](time.Time.md) | Password change schedule start date | [optional] [default to null]
**HeartbeatInterval** | **int32** | Heartbeat interval(mins). | [optional] [default to null]
**HeartbeatStartTime** | [**time.Time**](time.Time.md) | Password heartbeat start time | [optional] [default to null]
**CheckoutDuration** | **int32** | Setting of checkout duration(minutes). | [optional] [default to null]
**CheckinPwdChange** | **string** | Setting to change password on checkin    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**CheckoutDurationRenew** | **string** | Setting to renew checkout duration    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**CheckoutRenewTimes** | **int32** | Setting of the max amount of times you can renew checkout. | [optional] [default to null]
**TunnelEncryption** | **string** | Setting of using encrypted tunnel.    enable:Enable setting.    disable:Disable setting. | [optional] [default to null]
**ApprovalProfile** | **string** | Setting of approval-profile for secret. | [optional] [default to null]
**SshFilterProfile** | **string** | SSH filter profile. | [optional] [default to null]
**DlpSensor** | **string** | DLP sensor name. | [optional] [default to null]
**AvProfile** | **string** | Name of Antivirus profile. | [optional] [default to null]
**InheritPermission** | **string** | Enable/disable inherit permission from folder.    enable:Enable setting.    disable:Disable setting. | [default to null]
**SshServiceStatus** | **string** | SSH service status.    up:Service up on server.    down:Service down on server. | [optional] [default to null]
**SshServicePort** | **int32** | SSH service port (unset to use default). | [optional] [default to null]
**RdpServiceStatus** | **string** | RDP service status.    up:Service up on server.    down:Service down on server. | [optional] [default to null]
**RdpServicePort** | **int32** | RDP service port (unset to use default). | [optional] [default to null]
**VncServiceStatus** | **string** | VNC service status.    up:Service up on server.    down:Service down on server. | [optional] [default to null]
**VncServicePort** | **int32** | VNC service port (unset to use default). | [optional] [default to null]
**LdapsServiceStatus** | **string** | LDAPS service status.    up:Service up on server.    down:Service down on server. | [optional] [default to null]
**LdapsServicePort** | **int32** | LDAPS service port (unset to use default). | [optional] [default to null]
**SambaServiceStatus** | **string** | Samba service status.    up:Service up on server.    down:Service down on server. | [optional] [default to null]
**SambaServicePort** | **int32** | Samba service port (unset to use default). | [optional] [default to null]
**DisplayNum** | **int32** | Display number for vnc (unset to use default) | [optional] [default to null]
**AddressBlacklist** | [**[]CmdbsecretdatabaseidAddressblacklist**](cmdbsecretdatabaseid_addressblacklist.md) | RDP target blacklist addresses. | [optional] [default to null]
**AddressWhitelist** | [**[]CmdbsecretdatabaseidAddresswhitelist**](cmdbsecretdatabaseid_addresswhitelist.md) | RDP target whitelist addresses. | [optional] [default to null]
**ZtnaEmsTag** | [**[]CmdbsecretdatabaseidZtnaemstag**](cmdbsecretdatabaseid_ztnaemstag.md) | ZTNA ems tag that controls launching | [optional] [default to null]
**UserPermission** | [**[]CmdbsecretdatabaseidUserpermission**](cmdbsecretdatabaseid_userpermission.md) | Configure user permission. | [optional] [default to null]
**GroupPermission** | [**[]CmdbsecretdatabaseidGrouppermission**](cmdbsecretdatabaseid_grouppermission.md) | Configure group permission. | [optional] [default to null]
**Field** | [**[]CmdbsecretdatabaseField**](cmdbsecretdatabase_field.md) | Configure secret fields. Current secret credential information | [default to null]
**TotpSetting** | [**[]CmdbsecretdatabaseidTotpsetting**](cmdbsecretdatabaseid_totpsetting.md) | Third party time based one time password setting. For Fortitoken totp activation support, please use internal/secret-ftk api to update this field. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

