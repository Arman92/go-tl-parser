package tdlib

// AuthenticationCodeType Provides information about the method by which an authentication code is delivered to the user 
type AuthenticationCodeType interface {
MessageType() string
}

// AuthorizationState Represents the current authorization state of the client 
type AuthorizationState interface {
MessageType() string
}

// InputFile Points to a file 
type InputFile interface {
MessageType() string
}

// MaskPoint Part of the face, relative to which a mask should be placed 
type MaskPoint interface {
MessageType() string
}

// LinkState Represents the relationship between user A and user B. For incoming_link, user A is the current user; for outgoing_link, user B is the current user 
type LinkState interface {
MessageType() string
}

// UserType Represents the type of the user. The following types are possible: regular users, deleted users and bots 
type UserType interface {
MessageType() string
}

// ChatMemberStatus Provides information about the status of a member in a chat 
type ChatMemberStatus interface {
MessageType() string
}

// SupergroupMembersFilter Specifies the kind of chat members to return in getSupergroupMembers 
type SupergroupMembersFilter interface {
MessageType() string
}

// SecretChatState Describes the current secret chat state 
type SecretChatState interface {
MessageType() string
}

// MessageForwardInfo Contains information about the initial sender of a forwarded message 
type MessageForwardInfo interface {
MessageType() string
}

// MessageSendingState Contains information about the sending state of the message 
type MessageSendingState interface {
MessageType() string
}

// NotificationSettingsScope Describes the types of chats for which notification settings are applied 
type NotificationSettingsScope interface {
MessageType() string
}

// ChatType Describes the type of a chat 
type ChatType interface {
MessageType() string
}

// KeyboardButtonType Describes a keyboard button type 
type KeyboardButtonType interface {
MessageType() string
}

// InlineKeyboardButtonType Describes the type of an inline keyboard button 
type InlineKeyboardButtonType interface {
MessageType() string
}

// ReplyMarkup Contains a description of a custom keyboard and actions that can be done with it to quickly reply to bots 
type ReplyMarkup interface {
MessageType() string
}

// RichText Describes a text object inside an instant-view web page 
type RichText interface {
MessageType() string
}

// PageBlock Describes a block of an instant view web page 
type PageBlock interface {
MessageType() string
}

// InputCredentials Contains information about the payment method chosen by the user 
type InputCredentials interface {
MessageType() string
}

// MessageContent Contains the content of a message 
type MessageContent interface {
MessageType() string
}

// TextEntityType Represents a part of the text which must be formatted differently 
type TextEntityType interface {
MessageType() string
}

// InputMessageContent The content of a message to send 
type InputMessageContent interface {
MessageType() string
}

// SearchMessagesFilter Represents a filter for message search results 
type SearchMessagesFilter interface {
MessageType() string
}

// ChatAction Describes the different types of activity in a chat 
type ChatAction interface {
MessageType() string
}

// UserStatus Describes the last time the user was online 
type UserStatus interface {
MessageType() string
}

// CallDiscardReason Describes the reason why a call was discarded 
type CallDiscardReason interface {
MessageType() string
}

// CallState Describes the current call state 
type CallState interface {
MessageType() string
}

// InputInlineQueryResult Represents a single result of an inline query; for bots only 
type InputInlineQueryResult interface {
MessageType() string
}

// InlineQueryResult Represents a single result of an inline query 
type InlineQueryResult interface {
MessageType() string
}

// CallbackQueryPayload Represents a payload of a callback query 
type CallbackQueryPayload interface {
MessageType() string
}

// ChatEventAction Represents a chat event 
type ChatEventAction interface {
MessageType() string
}

// DeviceToken Represents a data needed to subscribe for push notifications. To use specific push notification service, you must specify the correct application platform and upload valid server authentication data at https://my.telegram.org 
type DeviceToken interface {
MessageType() string
}

// CheckChatUsernameResult Represents result of checking whether a username can be set for a chat 
type CheckChatUsernameResult interface {
MessageType() string
}

// OptionValue Represents the value of an option 
type OptionValue interface {
MessageType() string
}

// UserPrivacySettingRule Represents a single rule for managing privacy settings 
type UserPrivacySettingRule interface {
MessageType() string
}

// UserPrivacySetting Describes available user privacy settings 
type UserPrivacySetting interface {
MessageType() string
}

// ChatReportReason Describes the reason why a chat is reported 
type ChatReportReason interface {
MessageType() string
}

// FileType Represents the type of a file 
type FileType interface {
MessageType() string
}

// NetworkType Represents the type of a network 
type NetworkType interface {
MessageType() string
}

// NetworkStatisticsEntry Contains statistics about network usage 
type NetworkStatisticsEntry interface {
MessageType() string
}

// ConnectionState Describes the current state of the connection to Telegram servers 
type ConnectionState interface {
MessageType() string
}

// TopChatCategory Represents the categories of chats for which a list of frequently used chats can be retrieved 
type TopChatCategory interface {
MessageType() string
}

// TMeURLType Describes the type of a URL linking to an internal Telegram entity 
type TMeURLType interface {
MessageType() string
}

// TextParseMode Describes the way the text should be parsed for TextEntities 
type TextParseMode interface {
MessageType() string
}

// Proxy Contains information about a proxy server 
type Proxy interface {
MessageType() string
}

// Update Contains notifications about data changes 
type Update interface {
MessageType() string
}

type tdCommon struct {
Type string `json:"@type"`
Extra string `json:"extra"`
}

// Error An object of this type can be returned on every function call, in case of an error 
type Error struct {
tdCommon
Code int32 `json:"code"` // Error code; subject to future changes. If the error code is 406, the error message must not be processed in any way and must not be displayed to the user
Message string `json:"message"` // Error message; subject to future changes
}

// MessageType return the string telegram-type of Error 
func (error *Error) MessageType() string {
 return "error" }

// Ok An object of this type is returned on a successful function call for certain functions 
type Ok struct {
tdCommon

}

// MessageType return the string telegram-type of Ok 
func (ok *Ok) MessageType() string {
 return "ok" }

// TdlibParameters Contains parameters for TDLib initialization 
type TdlibParameters struct {
tdCommon
UseTestDc bool `json:"use_test_dc"` // If set to true, the Telegram test environment will be used instead of the production environment
UseSecretChats bool `json:"use_secret_chats"` // If set to true, support for secret chats will be enabled
APIHash string `json:"api_hash"` // Application identifier hash for Telegram API access, which can be obtained at https://my.telegram.org
APIID int32 `json:"api_id"` // Application identifier for Telegram API access, which can be obtained at https://my.telegram.org
DatabaseDirectory string `json:"database_directory"` // The path to the directory for the persistent database; if empty, the current working directory will be used
UseChatInfoDatabase bool `json:"use_chat_info_database"` // If set to true, the library will maintain a cache of users, basic groups, supergroups, channels and secret chats. Implies use_file_database
UseMessageDatabase bool `json:"use_message_database"` // If set to true, the library will maintain a cache of chats and messages. Implies use_chat_info_database
SystemVersion string `json:"system_version"` // Version of the operating system the application is being run on; must be non-empty
ApplicationVersion string `json:"application_version"` // Application version; must be non-empty
FilesDirectory string `json:"files_directory"` // The path to the directory for storing files; if empty, database_directory will be used
UseFileDatabase bool `json:"use_file_database"` // If set to true, information about downloaded and uploaded files will be saved between application restarts
SystemLanguageCode string `json:"system_language_code"` // IETF language tag of the user's operating system language; must be non-empty
DeviceModel string `json:"device_model"` // Model of the device the application is being run on; must be non-empty
EnableStorageOptimizer bool `json:"enable_storage_optimizer"` // If set to true, old files will automatically be deleted
IgnoreFileNames bool `json:"ignore_file_names"` // If set to true, original file names will be ignored. Otherwise, downloaded files will be saved under names as close as possible to the original name
}

// MessageType return the string telegram-type of TdlibParameters 
func (tdlibParameters *TdlibParameters) MessageType() string {
 return "tdlibParameters" }

// AuthenticationCodeTypeTelegramMessage An authentication code is delivered via a private Telegram message, which can be viewed in another client  
type AuthenticationCodeTypeTelegramMessage struct {
tdCommon
Length int32 `json:"length"` // Length of the code
}

// MessageType return the string telegram-type of AuthenticationCodeTypeTelegramMessage 
func (authenticationCodeTypeTelegramMessage *AuthenticationCodeTypeTelegramMessage) MessageType() string {
 return "authenticationCodeTypeTelegramMessage" }

// AuthenticationCodeTypeSms An authentication code is delivered via an SMS message to the specified phone number  
type AuthenticationCodeTypeSms struct {
tdCommon
Length int32 `json:"length"` // Length of the code
}

// MessageType return the string telegram-type of AuthenticationCodeTypeSms 
func (authenticationCodeTypeSms *AuthenticationCodeTypeSms) MessageType() string {
 return "authenticationCodeTypeSms" }

// AuthenticationCodeTypeCall An authentication code is delivered via a phone call to the specified phone number  
type AuthenticationCodeTypeCall struct {
tdCommon
Length int32 `json:"length"` // Length of the code
}

// MessageType return the string telegram-type of AuthenticationCodeTypeCall 
func (authenticationCodeTypeCall *AuthenticationCodeTypeCall) MessageType() string {
 return "authenticationCodeTypeCall" }

// AuthenticationCodeTypeFlashCall An authentication code is delivered by an immediately cancelled call to the specified phone number. The number from which the call was made is the code  
type AuthenticationCodeTypeFlashCall struct {
tdCommon
Pattern string `json:"pattern"` // Pattern of the phone number from which the call will be made
}

// MessageType return the string telegram-type of AuthenticationCodeTypeFlashCall 
func (authenticationCodeTypeFlashCall *AuthenticationCodeTypeFlashCall) MessageType() string {
 return "authenticationCodeTypeFlashCall" }

// AuthenticationCodeInfo Information about the authentication code that was sent  
type AuthenticationCodeInfo struct {
tdCommon
Type AuthenticationCodeType `json:"type"` // Describes the way the code was sent to the user 
NextType AuthenticationCodeType `json:"next_type"` // Describes the way the next code will be sent to the user; may be null 
Timeout int32 `json:"timeout"` // Timeout before the code should be re-sent, in seconds
PhoneNumber string `json:"phone_number"` // A phone number that is being authenticated 
}

// MessageType return the string telegram-type of AuthenticationCodeInfo 
func (authenticationCodeInfo *AuthenticationCodeInfo) MessageType() string {
 return "authenticationCodeInfo" }

// AuthorizationStateWaitTdlibParameters TDLib needs TdlibParameters for initialization 
type AuthorizationStateWaitTdlibParameters struct {
tdCommon

}

// MessageType return the string telegram-type of AuthorizationStateWaitTdlibParameters 
func (authorizationStateWaitTdlibParameters *AuthorizationStateWaitTdlibParameters) MessageType() string {
 return "authorizationStateWaitTdlibParameters" }

// AuthorizationStateWaitEncryptionKey TDLib needs an encryption key to decrypt the local database  
type AuthorizationStateWaitEncryptionKey struct {
tdCommon
IsEncrypted bool `json:"is_encrypted"` // True, if the database is currently encrypted
}

// MessageType return the string telegram-type of AuthorizationStateWaitEncryptionKey 
func (authorizationStateWaitEncryptionKey *AuthorizationStateWaitEncryptionKey) MessageType() string {
 return "authorizationStateWaitEncryptionKey" }

// AuthorizationStateWaitPhoneNumber TDLib needs the user's phone number to authorize 
type AuthorizationStateWaitPhoneNumber struct {
tdCommon

}

// MessageType return the string telegram-type of AuthorizationStateWaitPhoneNumber 
func (authorizationStateWaitPhoneNumber *AuthorizationStateWaitPhoneNumber) MessageType() string {
 return "authorizationStateWaitPhoneNumber" }

// AuthorizationStateWaitCode TDLib needs the user's authentication code to finalize authorization  
type AuthorizationStateWaitCode struct {
tdCommon
IsRegistered bool `json:"is_registered"` // True, if the user is already registered 
CodeInfo AuthenticationCodeInfo `json:"code_info"` // Information about the authorization code that was sent
}

// MessageType return the string telegram-type of AuthorizationStateWaitCode 
func (authorizationStateWaitCode *AuthorizationStateWaitCode) MessageType() string {
 return "authorizationStateWaitCode" }

// AuthorizationStateWaitPassword The user has been authorized, but needs to enter a password to start using the application  
type AuthorizationStateWaitPassword struct {
tdCommon
PasswordHint string `json:"password_hint"` // Hint for the password; can be empty 
HasRecoveryEmailAddress bool `json:"has_recovery_email_address"` // True if a recovery email address has been set up
RecoveryEmailAddressPattern string `json:"recovery_email_address_pattern"` // Pattern of the email address to which the recovery email was sent; empty until a recovery email has been sent
}

// MessageType return the string telegram-type of AuthorizationStateWaitPassword 
func (authorizationStateWaitPassword *AuthorizationStateWaitPassword) MessageType() string {
 return "authorizationStateWaitPassword" }

// AuthorizationStateReady The user has been successfully authorized. TDLib is now ready to answer queries 
type AuthorizationStateReady struct {
tdCommon

}

// MessageType return the string telegram-type of AuthorizationStateReady 
func (authorizationStateReady *AuthorizationStateReady) MessageType() string {
 return "authorizationStateReady" }

// AuthorizationStateLoggingOut The user is currently logging out 
type AuthorizationStateLoggingOut struct {
tdCommon

}

// MessageType return the string telegram-type of AuthorizationStateLoggingOut 
func (authorizationStateLoggingOut *AuthorizationStateLoggingOut) MessageType() string {
 return "authorizationStateLoggingOut" }

// AuthorizationStateClosing TDLib is closing, all subsequent queries will be answered with the error 500. Note that closing TDLib can take a while. All resources will be freed only after authorizationStateClosed has been received 
type AuthorizationStateClosing struct {
tdCommon

}

// MessageType return the string telegram-type of AuthorizationStateClosing 
func (authorizationStateClosing *AuthorizationStateClosing) MessageType() string {
 return "authorizationStateClosing" }

// AuthorizationStateClosed TDLib client is in its final state. All databases are closed and all resources are released. No other updates will be received after this. All queries will be responded to 
type AuthorizationStateClosed struct {
tdCommon

}

// MessageType return the string telegram-type of AuthorizationStateClosed 
func (authorizationStateClosed *AuthorizationStateClosed) MessageType() string {
 return "authorizationStateClosed" }

// PasswordState Represents the current state of 2-step verification  
type PasswordState struct {
tdCommon
HasPassword bool `json:"has_password"` // True if a 2-step verification password has been set up 
PasswordHint string `json:"password_hint"` // Hint for the password; can be empty 
HasRecoveryEmailAddress bool `json:"has_recovery_email_address"` // True if a recovery email has been set up 
UnconfirmedRecoveryEmailAddressPattern string `json:"unconfirmed_recovery_email_address_pattern"` // Pattern of the email address to which a confirmation email was sent
}

// MessageType return the string telegram-type of PasswordState 
func (passwordState *PasswordState) MessageType() string {
 return "passwordState" }

// PasswordRecoveryInfo Contains information available to the user after requesting password recovery  
type PasswordRecoveryInfo struct {
tdCommon
RecoveryEmailAddressPattern string `json:"recovery_email_address_pattern"` // Pattern of the email address to which a recovery email was sent
}

// MessageType return the string telegram-type of PasswordRecoveryInfo 
func (passwordRecoveryInfo *PasswordRecoveryInfo) MessageType() string {
 return "passwordRecoveryInfo" }

// RecoveryEmailAddress Contains information about the current recovery email address  
type RecoveryEmailAddress struct {
tdCommon
RecoveryEmailAddress string `json:"recovery_email_address"` // Recovery email address
}

// MessageType return the string telegram-type of RecoveryEmailAddress 
func (recoveryEmailAddress *RecoveryEmailAddress) MessageType() string {
 return "recoveryEmailAddress" }

// TemporaryPasswordState Returns information about the availability of a temporary password, which can be used for payments  
type TemporaryPasswordState struct {
tdCommon
HasPassword bool `json:"has_password"` // True, if a temporary password is available 
ValidFor int32 `json:"valid_for"` // Time left before the temporary password expires, in seconds
}

// MessageType return the string telegram-type of TemporaryPasswordState 
func (temporaryPasswordState *TemporaryPasswordState) MessageType() string {
 return "temporaryPasswordState" }

// LocalFile Represents a local file 
type LocalFile struct {
tdCommon
CanBeDownloaded bool `json:"can_be_downloaded"` // True, if it is possible to try to download or generate the file
CanBeDeleted bool `json:"can_be_deleted"` // True, if the file can be deleted
IsDownloadingActive bool `json:"is_downloading_active"` // True, if the file is currently being downloaded (or a local copy is being generated by some other means)
IsDownloadingCompleted bool `json:"is_downloading_completed"` // True, if the local copy is fully available
DownloadedPrefixSize int32 `json:"downloaded_prefix_size"` // If is_downloading_completed is false, then only some prefix of the file is ready to be read. downloaded_prefix_size is the size of that prefix
DownloadedSize int32 `json:"downloaded_size"` // Total downloaded file bytes. Should be used only for calculating download progress. The actual file size may be bigger, and some parts of it may contain garbage
Path string `json:"path"` // Local path to the locally available file part; may be empty
}

// MessageType return the string telegram-type of LocalFile 
func (localFile *LocalFile) MessageType() string {
 return "localFile" }

// RemoteFile Represents a remote file 
type RemoteFile struct {
tdCommon
ID string `json:"id"` // Remote file identifier, may be empty. Can be used across application restarts or even from other devices for the current user. If the ID starts with "http://" or "https://", it represents the HTTP URL of the file. TDLib is currently unable to download files if only their URL is known.
IsUploadingActive bool `json:"is_uploading_active"` // True, if the file is currently being uploaded (or a remote copy is being generated by some other means)
IsUploadingCompleted bool `json:"is_uploading_completed"` // True, if a remote copy is fully available
UploadedSize int32 `json:"uploaded_size"` // Size of the remote available part of the file; 0 if unknown
}

// MessageType return the string telegram-type of RemoteFile 
func (remoteFile *RemoteFile) MessageType() string {
 return "remoteFile" }

// File Represents a file 
type File struct {
tdCommon
ID int32 `json:"id"` // Unique file identifier
Size int32 `json:"size"` // File size; 0 if unknown
ExpectedSize int32 `json:"expected_size"` // Expected file size in case the exact file size is unknown, but an approximate size is known. Can be used to show download/upload progress
Local LocalFile `json:"local"` // Information about the local copy of the file
Remote RemoteFile `json:"remote"` // Information about the remote copy of the file
}

// MessageType return the string telegram-type of File 
func (file *File) MessageType() string {
 return "file" }

// InputFileID A file defined by its unique ID  
type InputFileID struct {
tdCommon
ID int32 `json:"id"` // Unique file identifier
}

// MessageType return the string telegram-type of InputFileID 
func (inputFileID *InputFileID) MessageType() string {
 return "inputFileId" }

// InputFileRemote A file defined by its remote ID  
type InputFileRemote struct {
tdCommon
ID string `json:"id"` // Remote file identifier
}

// MessageType return the string telegram-type of InputFileRemote 
func (inputFileRemote *InputFileRemote) MessageType() string {
 return "inputFileRemote" }

// InputFileLocal A file defined by a local path  
type InputFileLocal struct {
tdCommon
Path string `json:"path"` // Local path to the file
}

// MessageType return the string telegram-type of InputFileLocal 
func (inputFileLocal *InputFileLocal) MessageType() string {
 return "inputFileLocal" }

// InputFileGenerated A file generated by the client  
type InputFileGenerated struct {
tdCommon
OriginalPath string `json:"original_path"` // Local path to a file from which the file is generated, may be empty if there is no such file 
Conversion string `json:"conversion"` // String specifying the conversion applied to the original file; should be persistent across application restarts 
ExpectedSize int32 `json:"expected_size"` // Expected size of the generated file; 0 if unknown
}

// MessageType return the string telegram-type of InputFileGenerated 
func (inputFileGenerated *InputFileGenerated) MessageType() string {
 return "inputFileGenerated" }

// PhotoSize Photo description  
type PhotoSize struct {
tdCommon
Type string `json:"type"` // Thumbnail type (see https://core.telegram.org/constructor/photoSize) 
Photo File `json:"photo"` // Information about the photo file 
Width int32 `json:"width"` // Photo width 
Height int32 `json:"height"` // Photo height
}

// MessageType return the string telegram-type of PhotoSize 
func (photoSize *PhotoSize) MessageType() string {
 return "photoSize" }

// MaskPointForehead A mask should be placed relatively to the forehead 
type MaskPointForehead struct {
tdCommon

}

// MessageType return the string telegram-type of MaskPointForehead 
func (maskPointForehead *MaskPointForehead) MessageType() string {
 return "maskPointForehead" }

// MaskPointEyes A mask should be placed relatively to the eyes 
type MaskPointEyes struct {
tdCommon

}

// MessageType return the string telegram-type of MaskPointEyes 
func (maskPointEyes *MaskPointEyes) MessageType() string {
 return "maskPointEyes" }

// MaskPointMouth A mask should be placed relatively to the mouth 
type MaskPointMouth struct {
tdCommon

}

// MessageType return the string telegram-type of MaskPointMouth 
func (maskPointMouth *MaskPointMouth) MessageType() string {
 return "maskPointMouth" }

// MaskPointChin A mask should be placed relatively to the chin 
type MaskPointChin struct {
tdCommon

}

// MessageType return the string telegram-type of MaskPointChin 
func (maskPointChin *MaskPointChin) MessageType() string {
 return "maskPointChin" }

// MaskPosition Position on a photo where a mask should be placed  
type MaskPosition struct {
tdCommon
Point MaskPoint `json:"point"` // Part of the face, relative to which the mask should be placed
XShift float64 `json:"x_shift"` // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. (For example, -1.0 will place the mask just to the left of the default mask position)
YShift float64 `json:"y_shift"` // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. (For example, 1.0 will place the mask just below the default mask position)
Scale float64 `json:"scale"` // Mask scaling coefficient. (For example, 2.0 means a doubled size)
}

// MessageType return the string telegram-type of MaskPosition 
func (maskPosition *MaskPosition) MessageType() string {
 return "maskPosition" }

// TextEntity Represents a part of the text that needs to be formatted in some unusual way  
type TextEntity struct {
tdCommon
Length int32 `json:"length"` // Length of the entity, in UTF-16 code points 
Type TextEntityType `json:"type"` // Type of the entity
Offset int32 `json:"offset"` // Offset of the entity in UTF-16 code points 
}

// MessageType return the string telegram-type of TextEntity 
func (textEntity *TextEntity) MessageType() string {
 return "textEntity" }

// TextEntities Contains a list of text entities  
type TextEntities struct {
tdCommon
Entities []TextEntity `json:"entities"` // List of text entities
}

// MessageType return the string telegram-type of TextEntities 
func (textEntities *TextEntities) MessageType() string {
 return "textEntities" }

// FormattedText A text with some entities  
type FormattedText struct {
tdCommon
Text string `json:"text"` // The text 
Entities []TextEntity `json:"entities"` // Entities contained in the text
}

// MessageType return the string telegram-type of FormattedText 
func (formattedText *FormattedText) MessageType() string {
 return "formattedText" }

// Animation Describes an animation file. The animation must be encoded in GIF or MPEG4 format  
type Animation struct {
tdCommon
Height int32 `json:"height"` // Height of the animation
FileName string `json:"file_name"` // Original name of the file; as defined by the sender 
MimeType string `json:"mime_type"` // MIME type of the file, usually "image/gif" or "video/mp4" 
Thumbnail PhotoSize `json:"thumbnail"` // Animation thumbnail; may be null 
Animation File `json:"animation"` // File containing the animation
Duration int32 `json:"duration"` // Duration of the animation, in seconds; as defined by the sender 
Width int32 `json:"width"` // Width of the animation 
}

// MessageType return the string telegram-type of Animation 
func (animation *Animation) MessageType() string {
 return "animation" }

// Audio Describes an audio file. Audio is usually in MP3 format  
type Audio struct {
tdCommon
AlbumCoverThumbnail PhotoSize `json:"album_cover_thumbnail"` // The thumbnail of the album cover; as defined by the sender. The full size thumbnail should be extracted from the downloaded file; may be null 
Audio File `json:"audio"` // File containing the audio
Duration int32 `json:"duration"` // Duration of the audio, in seconds; as defined by the sender 
Title string `json:"title"` // Title of the audio; as defined by the sender 
Performer string `json:"performer"` // Performer of the audio; as defined by the sender
FileName string `json:"file_name"` // Original name of the file; as defined by the sender 
MimeType string `json:"mime_type"` // The MIME type of the file; as defined by the sender 
}

// MessageType return the string telegram-type of Audio 
func (audio *Audio) MessageType() string {
 return "audio" }

// Document Describes a document of any type  
type Document struct {
tdCommon
MimeType string `json:"mime_type"` // MIME type of the file; as defined by the sender
Thumbnail PhotoSize `json:"thumbnail"` // Document thumbnail; as defined by the sender; may be null 
Document File `json:"document"` // File containing the document
FileName string `json:"file_name"` // Original name of the file; as defined by the sender 
}

// MessageType return the string telegram-type of Document 
func (document *Document) MessageType() string {
 return "document" }

// Photo Describes a photo  
type Photo struct {
tdCommon
Sizes []PhotoSize `json:"sizes"` // Available variants of the photo, in different sizes
ID int64 `json:"id"` // Photo identifier; 0 for deleted photos 
HasStickers bool `json:"has_stickers"` // True, if stickers were added to the photo 
}

// MessageType return the string telegram-type of Photo 
func (photo *Photo) MessageType() string {
 return "photo" }

// Sticker Describes a sticker  
type Sticker struct {
tdCommon
SetID int64 `json:"set_id"` // The identifier of the sticker set to which the sticker belongs; 0 if none 
Width int32 `json:"width"` // Sticker width; as defined by the sender 
Height int32 `json:"height"` // Sticker height; as defined by the sender
Emoji string `json:"emoji"` // Emoji corresponding to the sticker 
IsMask bool `json:"is_mask"` // True, if the sticker is a mask 
MaskPosition MaskPosition `json:"mask_position"` // Position where the mask should be placed; may be null 
Thumbnail PhotoSize `json:"thumbnail"` // Sticker thumbnail in WEBP or JPEG format; may be null 
Sticker File `json:"sticker"` // File containing the sticker
}

// MessageType return the string telegram-type of Sticker 
func (sticker *Sticker) MessageType() string {
 return "sticker" }

// Video Describes a video file  
type Video struct {
tdCommon
MimeType string `json:"mime_type"` // MIME type of the file; as defined by the sender 
HasStickers bool `json:"has_stickers"` // True, if stickers were added to the photo
SupportsStreaming bool `json:"supports_streaming"` // True, if the video should be tried to be streamed 
Thumbnail PhotoSize `json:"thumbnail"` // Video thumbnail; as defined by the sender; may be null 
Video File `json:"video"` // File containing the video
Duration int32 `json:"duration"` // Duration of the video, in seconds; as defined by the sender 
Width int32 `json:"width"` // Video width; as defined by the sender 
Height int32 `json:"height"` // Video height; as defined by the sender
FileName string `json:"file_name"` // Original name of the file; as defined by the sender 
}

// MessageType return the string telegram-type of Video 
func (video *Video) MessageType() string {
 return "video" }

// VideoNote Describes a video note. The video must be equal in width and height, cropped to a circle, and stored in MPEG4 format  
type VideoNote struct {
tdCommon
Duration int32 `json:"duration"` // Duration of the video, in seconds; as defined by the sender 
Length int32 `json:"length"` // Video width and height; as defined by the sender 
Thumbnail PhotoSize `json:"thumbnail"` // Video thumbnail; as defined by the sender; may be null 
Video File `json:"video"` // File containing the video
}

// MessageType return the string telegram-type of VideoNote 
func (videoNote *VideoNote) MessageType() string {
 return "videoNote" }

// VoiceNote Describes a voice note. The voice note must be encoded with the Opus codec, and stored inside an OGG container. Voice notes can have only a single audio channel  
type VoiceNote struct {
tdCommon
Duration int32 `json:"duration"` // Duration of the voice note, in seconds; as defined by the sender
Waveform []byte `json:"waveform"` // A waveform representation of the voice note in 5-bit format 
MimeType string `json:"mime_type"` // MIME type of the file; as defined by the sender 
Voice File `json:"voice"` // File containing the voice note
}

// MessageType return the string telegram-type of VoiceNote 
func (voiceNote *VoiceNote) MessageType() string {
 return "voiceNote" }

// Contact Describes a user contact  
type Contact struct {
tdCommon
FirstName string `json:"first_name"` // First name of the user; 1-255 characters in length 
LastName string `json:"last_name"` // Last name of the user 
UserID int32 `json:"user_id"` // Identifier of the user, if known; otherwise 0
PhoneNumber string `json:"phone_number"` // Phone number of the user 
}

// MessageType return the string telegram-type of Contact 
func (contact *Contact) MessageType() string {
 return "contact" }

// Location Describes a location on planet Earth  
type Location struct {
tdCommon
Longitude float64 `json:"longitude"` // Longitude of the location, in degrees; as defined by the sender
Latitude float64 `json:"latitude"` // Latitude of the location in degrees; as defined by the sender 
}

// MessageType return the string telegram-type of Location 
func (location *Location) MessageType() string {
 return "location" }

// Venue Describes a venue  
type Venue struct {
tdCommon
Provider string `json:"provider"` // Provider of the venue database; as defined by the sender. Currently only "foursquare" needs to be supported
ID string `json:"id"` // Identifier of the venue in the provider database; as defined by the sender
Location Location `json:"location"` // Venue location; as defined by the sender 
Title string `json:"title"` // Venue name; as defined by the sender 
Address string `json:"address"` // Venue address; as defined by the sender 
}

// MessageType return the string telegram-type of Venue 
func (venue *Venue) MessageType() string {
 return "venue" }

// Game Describes a game  
type Game struct {
tdCommon
Photo Photo `json:"photo"` // Game photo 
Animation Animation `json:"animation"` // Game animation; may be null
ID int64 `json:"id"` // Game ID 
ShortName string `json:"short_name"` // Game short name. To share a game use the URL https://t.me/{bot_username}?game={game_short_name} 
Title string `json:"title"` // Game title 
Text FormattedText `json:"text"` // Game text, usually containing scoreboards for a game
Description string `json:"description"` // 
}

// MessageType return the string telegram-type of Game 
func (game *Game) MessageType() string {
 return "game" }

// ProfilePhoto Describes a user profile photo  
type ProfilePhoto struct {
tdCommon
ID int64 `json:"id"` // Photo identifier; 0 for an empty photo. Can be used to find a photo in a list of userProfilePhotos
Small File `json:"small"` // A small (160x160) user profile photo 
Big File `json:"big"` // A big (640x640) user profile photo
}

// MessageType return the string telegram-type of ProfilePhoto 
func (profilePhoto *ProfilePhoto) MessageType() string {
 return "profilePhoto" }

// ChatPhoto Describes the photo of a chat  
type ChatPhoto struct {
tdCommon
Small File `json:"small"` // A small (160x160) chat photo 
Big File `json:"big"` // A big (640x640) chat photo
}

// MessageType return the string telegram-type of ChatPhoto 
func (chatPhoto *ChatPhoto) MessageType() string {
 return "chatPhoto" }

// LinkStateNone The phone number of user A is not known to user B 
type LinkStateNone struct {
tdCommon

}

// MessageType return the string telegram-type of LinkStateNone 
func (linkStateNone *LinkStateNone) MessageType() string {
 return "linkStateNone" }

// LinkStateKnowsPhoneNumber The phone number of user A is known but that number has not been saved to the contacts list of user B 
type LinkStateKnowsPhoneNumber struct {
tdCommon

}

// MessageType return the string telegram-type of LinkStateKnowsPhoneNumber 
func (linkStateKnowsPhoneNumber *LinkStateKnowsPhoneNumber) MessageType() string {
 return "linkStateKnowsPhoneNumber" }

// LinkStateIsContact The phone number of user A has been saved to the contacts list of user B 
type LinkStateIsContact struct {
tdCommon

}

// MessageType return the string telegram-type of LinkStateIsContact 
func (linkStateIsContact *LinkStateIsContact) MessageType() string {
 return "linkStateIsContact" }

// UserTypeRegular A regular user 
type UserTypeRegular struct {
tdCommon

}

// MessageType return the string telegram-type of UserTypeRegular 
func (userTypeRegular *UserTypeRegular) MessageType() string {
 return "userTypeRegular" }

// UserTypeDeleted A deleted user or deleted bot. No information on the user besides the user_id is available. It is not possible to perform any active actions on this type of user 
type UserTypeDeleted struct {
tdCommon

}

// MessageType return the string telegram-type of UserTypeDeleted 
func (userTypeDeleted *UserTypeDeleted) MessageType() string {
 return "userTypeDeleted" }

// UserTypeBot A bot (see https://core.telegram.org/bots)  
type UserTypeBot struct {
tdCommon
CanJoinGroups bool `json:"can_join_groups"` // True, if the bot can be invited to basic group and supergroup chats
CanReadAllGroupMessages bool `json:"can_read_all_group_messages"` // True, if the bot can read all messages in basic group or supergroup chats and not just those addressed to the bot. In private and channel chats a bot can always read all messages
IsInline bool `json:"is_inline"` // True, if the bot supports inline queries 
InlineQueryPlaceholder string `json:"inline_query_placeholder"` // Placeholder for inline queries (displayed on the client input field) 
NeedLocation bool `json:"need_location"` // True, if the location of the user should be sent with every inline query to this bot
}

// MessageType return the string telegram-type of UserTypeBot 
func (userTypeBot *UserTypeBot) MessageType() string {
 return "userTypeBot" }

// UserTypeUnknown No information on the user besides the user_id is available, yet this user has not been deleted. This object is extremely rare and must be handled like a deleted user. It is not possible to perform any actions on users of this type 
type UserTypeUnknown struct {
tdCommon

}

// MessageType return the string telegram-type of UserTypeUnknown 
func (userTypeUnknown *UserTypeUnknown) MessageType() string {
 return "userTypeUnknown" }

// BotCommand Represents commands supported by a bot  
type BotCommand struct {
tdCommon
Command string `json:"command"` // Text of the bot command 
Description string `json:"description"` // 
}

// MessageType return the string telegram-type of BotCommand 
func (botCommand *BotCommand) MessageType() string {
 return "botCommand" }

// BotInfo Provides information about a bot and its supported commands  
type BotInfo struct {
tdCommon
Description string `json:"description"` // 
Commands []BotCommand `json:"commands"` // A list of commands supported by the bot
}

// MessageType return the string telegram-type of BotInfo 
func (botInfo *BotInfo) MessageType() string {
 return "botInfo" }

// User Represents a user  
type User struct {
tdCommon
Status UserStatus `json:"status"` // Current online status of the user 
ProfilePhoto ProfilePhoto `json:"profile_photo"` // Profile photo of the user; may be null
RestrictionReason string `json:"restriction_reason"` // If non-empty, it contains the reason why access to this user must be restricted. The format of the string is "{type}: {description}".
HaveAccess bool `json:"have_access"` // If false, the user is inaccessible, and the only information known about the user is inside this class. It can't be passed to any method except GetUser 
LastName string `json:"last_name"` // Last name of the user 
Username string `json:"username"` // Username of the user
IsVerified bool `json:"is_verified"` // True, if the user is verified 
ID int32 `json:"id"` // User identifier 
IncomingLink LinkState `json:"incoming_link"` // Relationship from the other user to the current user 
Type UserType `json:"type"` // Type of the user 
LanguageCode string `json:"language_code"` // IETF language tag of the user's language; only available to bots
FirstName string `json:"first_name"` // First name of the user 
PhoneNumber string `json:"phone_number"` // Phone number of the user 
OutgoingLink LinkState `json:"outgoing_link"` // Relationship from the current user to the other user 
}

// MessageType return the string telegram-type of User 
func (user *User) MessageType() string {
 return "user" }

// UserFullInfo Contains full information about a user (except the full list of profile photos)  
type UserFullInfo struct {
tdCommon
IsBlocked bool `json:"is_blocked"` // True, if the user is blacklisted by the current user 
CanBeCalled bool `json:"can_be_called"` // True, if the user can be called 
HasPrivateCalls bool `json:"has_private_calls"` // True, if the user can't be called due to their privacy settings
Bio string `json:"bio"` // A short user bio 
ShareText string `json:"share_text"` // For bots, the text that is included with the link when users share the bot 
GroupInCommonCount int32 `json:"group_in_common_count"` // Number of group chats where both the other user and the current user are a member; 0 for the current user 
BotInfo BotInfo `json:"bot_info"` // If the user is a bot, information about the bot; may be null
}

// MessageType return the string telegram-type of UserFullInfo 
func (userFullInfo *UserFullInfo) MessageType() string {
 return "userFullInfo" }

// UserProfilePhotos Contains part of the list of user photos  
type UserProfilePhotos struct {
tdCommon
Photos []Photo `json:"photos"` // A list of photos
TotalCount int32 `json:"total_count"` // Total number of user profile photos 
}

// MessageType return the string telegram-type of UserProfilePhotos 
func (userProfilePhotos *UserProfilePhotos) MessageType() string {
 return "userProfilePhotos" }

// Users Represents a list of users  
type Users struct {
tdCommon
TotalCount int32 `json:"total_count"` // Approximate total count of users found 
UserIDs []int32 `json:"user_ids"` // A list of user identifiers
}

// MessageType return the string telegram-type of Users 
func (users *Users) MessageType() string {
 return "users" }

// ChatMemberStatusCreator The user is the creator of a chat and has all the administrator privileges  
type ChatMemberStatusCreator struct {
tdCommon
IsMember bool `json:"is_member"` // True, if the user is a member of the chat
}

// MessageType return the string telegram-type of ChatMemberStatusCreator 
func (chatMemberStatusCreator *ChatMemberStatusCreator) MessageType() string {
 return "chatMemberStatusCreator" }

// ChatMemberStatusAdministrator The user is a member of a chat and has some additional privileges. In basic groups, administrators can edit and delete messages sent by others, add new members, and ban unprivileged members. In supergroups and channels, there are more detailed options for administrator privileges 
type ChatMemberStatusAdministrator struct {
tdCommon
CanBeEdited bool `json:"can_be_edited"` // True, if the current user can edit the administrator privileges for the called user
CanChangeInfo bool `json:"can_change_info"` // True, if the administrator can change the chat title, photo, and other settings
CanPostMessages bool `json:"can_post_messages"` // True, if the administrator can create channel posts; applicable to channels only
CanEditMessages bool `json:"can_edit_messages"` // True, if the administrator can edit messages of other users and pin messages; applicable to channels only
CanDeleteMessages bool `json:"can_delete_messages"` // True, if the administrator can delete messages of other users
CanInviteUsers bool `json:"can_invite_users"` // True, if the administrator can invite new users to the chat
CanPinMessages bool `json:"can_pin_messages"` // True, if the administrator can pin messages; applicable to supergroups only
CanRestrictMembers bool `json:"can_restrict_members"` // True, if the administrator can restrict, ban, or unban chat members
CanPromoteMembers bool `json:"can_promote_members"` // True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that were directly or indirectly promoted by him
}

// MessageType return the string telegram-type of ChatMemberStatusAdministrator 
func (chatMemberStatusAdministrator *ChatMemberStatusAdministrator) MessageType() string {
 return "chatMemberStatusAdministrator" }

// ChatMemberStatusMember The user is a member of a chat, without any additional privileges or restrictions 
type ChatMemberStatusMember struct {
tdCommon

}

// MessageType return the string telegram-type of ChatMemberStatusMember 
func (chatMemberStatusMember *ChatMemberStatusMember) MessageType() string {
 return "chatMemberStatusMember" }

// ChatMemberStatusRestricted The user is under certain restrictions in the chat. Not supported in basic groups and channels 
type ChatMemberStatusRestricted struct {
tdCommon
CanSendMediaMessages bool `json:"can_send_media_messages"` // True, if the user can send audio files, documents, photos, videos, video notes, and voice notes. Implies can_send_messages permissions
CanSendOtherMessages bool `json:"can_send_other_messages"` // True, if the user can send animations, games, and stickers and use inline bots. Implies can_send_media_messages permissions
CanAddWebPagePreviews bool `json:"can_add_web_page_previews"` // True, if the user may add a web page preview to his messages. Implies can_send_messages permissions
IsMember bool `json:"is_member"` // True, if the user is a member of the chat
RestrictedUntilDate int32 `json:"restricted_until_date"` // Point in time (Unix timestamp) when restrictions will be lifted from the user; 0 if never. If the user is restricted for more than 366 days or for less than 30 seconds from the current time, the user is considered to be restricted forever
CanSendMessages bool `json:"can_send_messages"` // True, if the user can send text messages, contacts, locations, and venues
}

// MessageType return the string telegram-type of ChatMemberStatusRestricted 
func (chatMemberStatusRestricted *ChatMemberStatusRestricted) MessageType() string {
 return "chatMemberStatusRestricted" }

// ChatMemberStatusLeft The user is not a chat member 
type ChatMemberStatusLeft struct {
tdCommon

}

// MessageType return the string telegram-type of ChatMemberStatusLeft 
func (chatMemberStatusLeft *ChatMemberStatusLeft) MessageType() string {
 return "chatMemberStatusLeft" }

// ChatMemberStatusBanned The user was banned (and hence is not a member of the chat). Implies the user can't return to the chat or view messages 
type ChatMemberStatusBanned struct {
tdCommon
BannedUntilDate int32 `json:"banned_until_date"` // Point in time (Unix timestamp) when the user will be unbanned; 0 if never. If the user is banned for more than 366 days or for less than 30 seconds from the current time, the user is considered to be banned forever
}

// MessageType return the string telegram-type of ChatMemberStatusBanned 
func (chatMemberStatusBanned *ChatMemberStatusBanned) MessageType() string {
 return "chatMemberStatusBanned" }

// ChatMember A user with information about joining/leaving a chat  
type ChatMember struct {
tdCommon
UserID int32 `json:"user_id"` // User identifier of the chat member 
InviterUserID int32 `json:"inviter_user_id"` // Identifier of a user that invited/promoted/banned this member in the chat; 0 if unknown
JoinedChatDate int32 `json:"joined_chat_date"` // Point in time (Unix timestamp) when the user joined a chat 
Status ChatMemberStatus `json:"status"` // Status of the member in the chat 
BotInfo BotInfo `json:"bot_info"` // If the user is a bot, information about the bot; may be null. Can be null even for a bot if the bot is not a chat member
}

// MessageType return the string telegram-type of ChatMember 
func (chatMember *ChatMember) MessageType() string {
 return "chatMember" }

// ChatMembers Contains a list of chat members  
type ChatMembers struct {
tdCommon
TotalCount int32 `json:"total_count"` // Approximate total count of chat members found 
Members []ChatMember `json:"members"` // A list of chat members
}

// MessageType return the string telegram-type of ChatMembers 
func (chatMembers *ChatMembers) MessageType() string {
 return "chatMembers" }

// SupergroupMembersFilterRecent Returns recently active users in reverse chronological order 
type SupergroupMembersFilterRecent struct {
tdCommon

}

// MessageType return the string telegram-type of SupergroupMembersFilterRecent 
func (supergroupMembersFilterRecent *SupergroupMembersFilterRecent) MessageType() string {
 return "supergroupMembersFilterRecent" }

// SupergroupMembersFilterAdministrators Returns the creator and administrators 
type SupergroupMembersFilterAdministrators struct {
tdCommon

}

// MessageType return the string telegram-type of SupergroupMembersFilterAdministrators 
func (supergroupMembersFilterAdministrators *SupergroupMembersFilterAdministrators) MessageType() string {
 return "supergroupMembersFilterAdministrators" }

// SupergroupMembersFilterSearch Used to search for supergroup or channel members via a (string) query  
type SupergroupMembersFilterSearch struct {
tdCommon
Query string `json:"query"` // Query to search for
}

// MessageType return the string telegram-type of SupergroupMembersFilterSearch 
func (supergroupMembersFilterSearch *SupergroupMembersFilterSearch) MessageType() string {
 return "supergroupMembersFilterSearch" }

// SupergroupMembersFilterRestricted Returns restricted supergroup members; can be used only by administrators  
type SupergroupMembersFilterRestricted struct {
tdCommon
Query string `json:"query"` // Query to search for
}

// MessageType return the string telegram-type of SupergroupMembersFilterRestricted 
func (supergroupMembersFilterRestricted *SupergroupMembersFilterRestricted) MessageType() string {
 return "supergroupMembersFilterRestricted" }

// SupergroupMembersFilterBanned Returns users banned from the supergroup or channel; can be used only by administrators  
type SupergroupMembersFilterBanned struct {
tdCommon
Query string `json:"query"` // Query to search for
}

// MessageType return the string telegram-type of SupergroupMembersFilterBanned 
func (supergroupMembersFilterBanned *SupergroupMembersFilterBanned) MessageType() string {
 return "supergroupMembersFilterBanned" }

// SupergroupMembersFilterBots Returns bot members of the supergroup or channel 
type SupergroupMembersFilterBots struct {
tdCommon

}

// MessageType return the string telegram-type of SupergroupMembersFilterBots 
func (supergroupMembersFilterBots *SupergroupMembersFilterBots) MessageType() string {
 return "supergroupMembersFilterBots" }

// BasicGroup Represents a basic group of 0-200 users (must be upgraded to a supergroup to accommodate more than 200 users) 
type BasicGroup struct {
tdCommon
UpgradedToSupergroupID int32 `json:"upgraded_to_supergroup_id"` // Identifier of the supergroup to which this group was upgraded; 0 if none
ID int32 `json:"id"` // Group identifier
MemberCount int32 `json:"member_count"` // Number of members in the group
Status ChatMemberStatus `json:"status"` // Status of the current user in the group
EveryoneIsAdministrator bool `json:"everyone_is_administrator"` // True, if all members have been granted administrator rights in the group
IsActive bool `json:"is_active"` // True, if the group is active
}

// MessageType return the string telegram-type of BasicGroup 
func (basicGroup *BasicGroup) MessageType() string {
 return "basicGroup" }

// BasicGroupFullInfo Contains full information about a basic group  
type BasicGroupFullInfo struct {
tdCommon
CreatorUserID int32 `json:"creator_user_id"` // User identifier of the creator of the group; 0 if unknown 
Members []ChatMember `json:"members"` // Group members 
InviteLink string `json:"invite_link"` // Invite link for this group; available only for the group creator and only after it has been generated at least once
}

// MessageType return the string telegram-type of BasicGroupFullInfo 
func (basicGroupFullInfo *BasicGroupFullInfo) MessageType() string {
 return "basicGroupFullInfo" }

// Supergroup Represents a supergroup or channel with zero or more members (subscribers in the case of channels). From the point of view of the system, a channel is a special kind of a supergroup: only administrators can post and see the list of members, and posts from all administrators use the name and photo of the channel instead of individual names and profile photos. Unlike supergroups, channels can have an unlimited number of subscribers 
type Supergroup struct {
tdCommon
MemberCount int32 `json:"member_count"` // Member count; 0 if unknown. Currently it is guaranteed to be known only if the supergroup or channel was found through SearchPublicChats
AnyoneCanInvite bool `json:"anyone_can_invite"` // True, if any member of the supergroup can invite other members. This field has no meaning for channels
SignMessages bool `json:"sign_messages"` // True, if messages sent to the channel should contain information about the sender. This field is only applicable to channels
IsChannel bool `json:"is_channel"` // True, if the supergroup is a channel
ID int32 `json:"id"` // Supergroup or channel identifier
Date int32 `json:"date"` // Point in time (Unix timestamp) when the current user joined, or the point in time when the supergroup or channel was created, in case the user is not a member
Status ChatMemberStatus `json:"status"` // Status of the current user in the supergroup or channel
IsVerified bool `json:"is_verified"` // True, if the supergroup or channel is verified
RestrictionReason string `json:"restriction_reason"` // If non-empty, contains the reason why access to this supergroup or channel must be restricted. Format of the string is "{type}: {description}".
Username string `json:"username"` // Username of the supergroup or channel; empty for private supergroups or channels
}

// MessageType return the string telegram-type of Supergroup 
func (supergroup *Supergroup) MessageType() string {
 return "supergroup" }

// SupergroupFullInfo Contains full information about a supergroup or channel 
type SupergroupFullInfo struct {
tdCommon
AdministratorCount int32 `json:"administrator_count"` // Number of privileged users in the supergroup or channel; 0 if unknown
BannedCount int32 `json:"banned_count"` // Number of users banned from chat; 0 if unknown
Description string `json:"description"` // 
CanGetMembers bool `json:"can_get_members"` // True, if members of the chat can be retrieved
InviteLink string `json:"invite_link"` // Invite link for this chat
CanSetUsername bool `json:"can_set_username"` // True, if the chat can be made public
CanSetStickerSet bool `json:"can_set_sticker_set"` // True, if the supergroup sticker set can be changed
StickerSetID int64 `json:"sticker_set_id"` // Identifier of the supergroup sticker set; 0 if none
MemberCount int32 `json:"member_count"` // Number of members in the supergroup or channel; 0 if unknown
RestrictedCount int32 `json:"restricted_count"` // Number of restricted users in the supergroup; 0 if unknown
IsAllHistoryAvailable bool `json:"is_all_history_available"` // True, if new chat members will have access to old messages. In public supergroups and both public and private channels, old messages are always available, so this option affects only private supergroups. The value of this field is only available for chat administrators
PinnedMessageID int64 `json:"pinned_message_id"` // Identifier of the pinned message in the chat; 0 if none
UpgradedFromBasicGroupID int32 `json:"upgraded_from_basic_group_id"` // Identifier of the basic group from which supergroup was upgraded; 0 if none
UpgradedFromMaxMessageID int64 `json:"upgraded_from_max_message_id"` // Identifier of the last message in the basic group from which supergroup was upgraded; 0 if none
}

// MessageType return the string telegram-type of SupergroupFullInfo 
func (supergroupFullInfo *SupergroupFullInfo) MessageType() string {
 return "supergroupFullInfo" }

// SecretChatStatePending The secret chat is not yet created; waiting for the other user to get online 
type SecretChatStatePending struct {
tdCommon

}

// MessageType return the string telegram-type of SecretChatStatePending 
func (secretChatStatePending *SecretChatStatePending) MessageType() string {
 return "secretChatStatePending" }

// SecretChatStateReady The secret chat is ready to use 
type SecretChatStateReady struct {
tdCommon

}

// MessageType return the string telegram-type of SecretChatStateReady 
func (secretChatStateReady *SecretChatStateReady) MessageType() string {
 return "secretChatStateReady" }

// SecretChatStateClosed The secret chat is closed 
type SecretChatStateClosed struct {
tdCommon

}

// MessageType return the string telegram-type of SecretChatStateClosed 
func (secretChatStateClosed *SecretChatStateClosed) MessageType() string {
 return "secretChatStateClosed" }

// SecretChat Represents a secret chat 
type SecretChat struct {
tdCommon
UserID int32 `json:"user_id"` // Identifier of the chat partner
State SecretChatState `json:"state"` // State of the secret chat
IsOutbound bool `json:"is_outbound"` // True, if the chat was created by the current user; otherwise false
TTL int32 `json:"ttl"` // Current message Time To Live setting (self-destruct timer) for the chat, in seconds
KeyHash []byte `json:"key_hash"` // Hash of the currently used key for comparison with the hash of the chat partner's key. This is a string of 36 bytes, which must be used to make a 12x12 square image with a color depth of 4. The first 16 bytes should be used to make a central 8x8 square, while the remaining 20 bytes should be used to construct a 2-pixel-wide border around that square.
Layer int32 `json:"layer"` // Secret chat layer; determines features supported by the other client. Video notes are supported if the layer >= 66
ID int32 `json:"id"` // Secret chat identifier
}

// MessageType return the string telegram-type of SecretChat 
func (secretChat *SecretChat) MessageType() string {
 return "secretChat" }

// MessageForwardedFromUser The message was originally written by a known user  
type MessageForwardedFromUser struct {
tdCommon
SenderUserID int32 `json:"sender_user_id"` // Identifier of the user that originally sent this message 
Date int32 `json:"date"` // Point in time (Unix timestamp) when the message was originally sent
ForwardedFromChatID int64 `json:"forwarded_from_chat_id"` // For messages forwarded to the chat with the current user (saved messages), the identifier of the chat from which the message was forwarded; 0 if unknown
ForwardedFromMessageID int64 `json:"forwarded_from_message_id"` // For messages forwarded to the chat with the current user (saved messages) the identifier of the original message from which the new message was forwarded; 0 if unknown
}

// MessageType return the string telegram-type of MessageForwardedFromUser 
func (messageForwardedFromUser *MessageForwardedFromUser) MessageType() string {
 return "messageForwardedFromUser" }

// MessageForwardedPost The message was originally a post in a channel  
type MessageForwardedPost struct {
tdCommon
ChatID int64 `json:"chat_id"` // Identifier of the chat from which the message was forwarded 
AuthorSignature string `json:"author_signature"` // Post author signature
Date int32 `json:"date"` // Point in time (Unix timestamp) when the message was originally sent 
MessageID int64 `json:"message_id"` // Message identifier of the original message from which the new message was forwarded; 0 if unknown
ForwardedFromChatID int64 `json:"forwarded_from_chat_id"` // For messages forwarded to the chat with the current user (saved messages), the identifier of the chat from which the message was forwarded; 0 if unknown
ForwardedFromMessageID int64 `json:"forwarded_from_message_id"` // For messages forwarded to the chat with the current user (saved messages), the identifier of the original message from which the new message was forwarded; 0 if unknown
}

// MessageType return the string telegram-type of MessageForwardedPost 
func (messageForwardedPost *MessageForwardedPost) MessageType() string {
 return "messageForwardedPost" }

// MessageSendingStatePending The message is being sent now, but has not yet been delivered to the server 
type MessageSendingStatePending struct {
tdCommon

}

// MessageType return the string telegram-type of MessageSendingStatePending 
func (messageSendingStatePending *MessageSendingStatePending) MessageType() string {
 return "messageSendingStatePending" }

// MessageSendingStateFailed The message failed to be sent 
type MessageSendingStateFailed struct {
tdCommon

}

// MessageType return the string telegram-type of MessageSendingStateFailed 
func (messageSendingStateFailed *MessageSendingStateFailed) MessageType() string {
 return "messageSendingStateFailed" }

// Message Describes a message 
type Message struct {
tdCommon
IsOutgoing bool `json:"is_outgoing"` // True, if the message is outgoing
CanBeForwarded bool `json:"can_be_forwarded"` // True, if the message can be forwarded
SenderUserID int32 `json:"sender_user_id"` // Identifier of the user who sent the message; 0 if unknown. It is unknown for channel posts
ChatID int64 `json:"chat_id"` // Chat identifier
ForwardInfo MessageForwardInfo `json:"forward_info"` // Information about the initial message sender; may be null
ReplyToMessageID int64 `json:"reply_to_message_id"` // If non-zero, the identifier of the message this message is replying to; can be the identifier of a deleted message
MediaAlbumID int64 `json:"media_album_id"` // Unique identifier of an album this message belongs to. Only photos and videos can be grouped together in albums
ID int64 `json:"id"` // Unique message identifier
ContainsUnreadMention bool `json:"contains_unread_mention"` // True, if the message contains an unread mention for the current user
CanBeDeletedForAllUsers bool `json:"can_be_deleted_for_all_users"` // True, if the message can be deleted for all users
Date int32 `json:"date"` // Point in time (Unix timestamp) when the message was sent
EditDate int32 `json:"edit_date"` // Point in time (Unix timestamp) when the message was last edited
TTLExpiresIn float64 `json:"ttl_expires_in"` // Time left before the message expires, in seconds
ViaBotUserID int32 `json:"via_bot_user_id"` // If non-zero, the user identifier of the bot through which this message was sent
AuthorSignature string `json:"author_signature"` // For channel posts, optional author signature
SendingState MessageSendingState `json:"sending_state"` // Information about the sending state of the message; may be null
CanBeDeletedOnlyForSelf bool `json:"can_be_deleted_only_for_self"` // True, if the message can be deleted only for the current user while other users will continue to see it
TTL int32 `json:"ttl"` // For self-destructing messages, the message's TTL (Time To Live), in seconds; 0 if none. TDLib will send updateDeleteMessages or updateMessageContent once the TTL expires
Views int32 `json:"views"` // Number of times this message was viewed
Content MessageContent `json:"content"` // Content of the message
ReplyMarkup ReplyMarkup `json:"reply_markup"` // Reply markup for the message; may be null
CanBeEdited bool `json:"can_be_edited"` // True, if the message can be edited
IsChannelPost bool `json:"is_channel_post"` // True, if the message is a channel post. All messages to channels are channel posts, all other messages are not channel posts
}

// MessageType return the string telegram-type of Message 
func (message *Message) MessageType() string {
 return "message" }

// Messages Contains a list of messages  
type Messages struct {
tdCommon
TotalCount int32 `json:"total_count"` // Approximate total count of messages found 
Messages []Message `json:"messages"` // List of messages; messages may be null
}

// MessageType return the string telegram-type of Messages 
func (messages *Messages) MessageType() string {
 return "messages" }

// FoundMessages Contains a list of messages found by a search  
type FoundMessages struct {
tdCommon
NextFromSearchID int64 `json:"next_from_search_id"` // Value to pass as from_search_id to get more results
Messages []Message `json:"messages"` // List of messages 
}

// MessageType return the string telegram-type of FoundMessages 
func (foundMessages *FoundMessages) MessageType() string {
 return "foundMessages" }

// NotificationSettingsScopeChat Notification settings applied to a particular chat  
type NotificationSettingsScopeChat struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier
}

// MessageType return the string telegram-type of NotificationSettingsScopeChat 
func (notificationSettingsScopeChat *NotificationSettingsScopeChat) MessageType() string {
 return "notificationSettingsScopeChat" }

// NotificationSettingsScopePrivateChats Notification settings applied to all private chats 
type NotificationSettingsScopePrivateChats struct {
tdCommon

}

// MessageType return the string telegram-type of NotificationSettingsScopePrivateChats 
func (notificationSettingsScopePrivateChats *NotificationSettingsScopePrivateChats) MessageType() string {
 return "notificationSettingsScopePrivateChats" }

// NotificationSettingsScopeBasicGroupChats Notification settings applied to all basic groups and channels. (Supergroups have no common settings) 
type NotificationSettingsScopeBasicGroupChats struct {
tdCommon

}

// MessageType return the string telegram-type of NotificationSettingsScopeBasicGroupChats 
func (notificationSettingsScopeBasicGroupChats *NotificationSettingsScopeBasicGroupChats) MessageType() string {
 return "notificationSettingsScopeBasicGroupChats" }

// NotificationSettingsScopeAllChats Notification settings applied to all chats 
type NotificationSettingsScopeAllChats struct {
tdCommon

}

// MessageType return the string telegram-type of NotificationSettingsScopeAllChats 
func (notificationSettingsScopeAllChats *NotificationSettingsScopeAllChats) MessageType() string {
 return "notificationSettingsScopeAllChats" }

// NotificationSettings Contains information about notification settings for a chat or several chats  
type NotificationSettings struct {
tdCommon
Sound string `json:"sound"` // An audio file name for notification sounds; only applies to iOS applications 
ShowPreview bool `json:"show_preview"` // True, if message content should be displayed in notifications
MuteFor int32 `json:"mute_for"` // Time left before notifications will be unmuted, in seconds 
}

// MessageType return the string telegram-type of NotificationSettings 
func (notificationSettings *NotificationSettings) MessageType() string {
 return "notificationSettings" }

// DraftMessage Contains information about a message draft  
type DraftMessage struct {
tdCommon
ReplyToMessageID int64 `json:"reply_to_message_id"` // Identifier of the message to reply to; 0 if none 
InputMessageText InputMessageContent `json:"input_message_text"` // Content of the message draft; this should always be of type inputMessageText
}

// MessageType return the string telegram-type of DraftMessage 
func (draftMessage *DraftMessage) MessageType() string {
 return "draftMessage" }

// ChatTypePrivate An ordinary chat with a user  
type ChatTypePrivate struct {
tdCommon
UserID int32 `json:"user_id"` // User identifier
}

// MessageType return the string telegram-type of ChatTypePrivate 
func (chatTypePrivate *ChatTypePrivate) MessageType() string {
 return "chatTypePrivate" }

// ChatTypeBasicGroup A basic group (i.e., a chat with 0-200 other users)  
type ChatTypeBasicGroup struct {
tdCommon
BasicGroupID int32 `json:"basic_group_id"` // Basic group identifier
}

// MessageType return the string telegram-type of ChatTypeBasicGroup 
func (chatTypeBasicGroup *ChatTypeBasicGroup) MessageType() string {
 return "chatTypeBasicGroup" }

// ChatTypeSupergroup A supergroup (i.e. a chat with up to GetOption("supergroup_max_size") other users), or channel (with unlimited members)  
type ChatTypeSupergroup struct {
tdCommon
IsChannel bool `json:"is_channel"` // True, if the supergroup is a channel
SupergroupID int32 `json:"supergroup_id"` // Supergroup or channel identifier 
}

// MessageType return the string telegram-type of ChatTypeSupergroup 
func (chatTypeSupergroup *ChatTypeSupergroup) MessageType() string {
 return "chatTypeSupergroup" }

// ChatTypeSecret A secret chat with a user  
type ChatTypeSecret struct {
tdCommon
SecretChatID int32 `json:"secret_chat_id"` // Secret chat identifier 
UserID int32 `json:"user_id"` // User identifier of the secret chat peer
}

// MessageType return the string telegram-type of ChatTypeSecret 
func (chatTypeSecret *ChatTypeSecret) MessageType() string {
 return "chatTypeSecret" }

// Chat A chat. (Can be a private chat, basic group, supergroup, or secret chat) 
type Chat struct {
tdCommon
UnreadCount int32 `json:"unread_count"` // Number of unread messages in the chat
LastReadInboxMessageID int64 `json:"last_read_inbox_message_id"` // Identifier of the last read incoming message
LastReadOutboxMessageID int64 `json:"last_read_outbox_message_id"` // Identifier of the last read outgoing message
CanBeReported bool `json:"can_be_reported"` // True, if the chat can be reported to Telegram moderators through reportChat
Title string `json:"title"` // Chat title
Photo ChatPhoto `json:"photo"` // Chat photo; may be null
Order int64 `json:"order"` // Descending parameter by which chats are sorted in the main chat list. If the order number of two chats is the same, they must be sorted in descending order by ID. If 0, the position of the chat in the list is undetermined
UnreadMentionCount int32 `json:"unread_mention_count"` // Number of unread messages with a mention/reply in the chat
DraftMessage DraftMessage `json:"draft_message"` // A draft of a message in the chat; may be null
ClientData string `json:"client_data"` // Contains client-specific data associated with the chat. (For example, the chat position or local chat notification settings can be stored here.) Persistent if a message database is used
ID int64 `json:"id"` // Chat unique identifier
LastMessage Message `json:"last_message"` // Last message in the chat; may be null
IsPinned bool `json:"is_pinned"` // True, if the chat is pinned
NotificationSettings NotificationSettings `json:"notification_settings"` // Notification settings for this chat
ReplyMarkupMessageID int64 `json:"reply_markup_message_id"` // Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
Type ChatType `json:"type"` // Type of the chat
}

// MessageType return the string telegram-type of Chat 
func (chat *Chat) MessageType() string {
 return "chat" }

// Chats Represents a list of chats  
type Chats struct {
tdCommon
ChatIDs []int64 `json:"chat_ids"` // List of chat identifiers
}

// MessageType return the string telegram-type of Chats 
func (chats *Chats) MessageType() string {
 return "chats" }

// ChatInviteLink Contains a chat invite link  
type ChatInviteLink struct {
tdCommon
InviteLink string `json:"invite_link"` // Chat invite link
}

// MessageType return the string telegram-type of ChatInviteLink 
func (chatInviteLink *ChatInviteLink) MessageType() string {
 return "chatInviteLink" }

// ChatInviteLinkInfo Contains information about a chat invite link 
type ChatInviteLinkInfo struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier of the invite link; 0 if the user is not a member of this chat
Type ChatType `json:"type"` // Contains information about the type of the chat
Title string `json:"title"` // Title of the chat
Photo ChatPhoto `json:"photo"` // Chat photo; may be null
MemberCount int32 `json:"member_count"` // Number of members
MemberUserIDs []int32 `json:"member_user_ids"` // User identifiers of some chat members that may be known to the current user
IsPublic bool `json:"is_public"` // True, if the chat is a public supergroup or channel with a username
}

// MessageType return the string telegram-type of ChatInviteLinkInfo 
func (chatInviteLinkInfo *ChatInviteLinkInfo) MessageType() string {
 return "chatInviteLinkInfo" }

// KeyboardButtonTypeText A simple button, with text that should be sent when the button is pressed 
type KeyboardButtonTypeText struct {
tdCommon

}

// MessageType return the string telegram-type of KeyboardButtonTypeText 
func (keyboardButtonTypeText *KeyboardButtonTypeText) MessageType() string {
 return "keyboardButtonTypeText" }

// KeyboardButtonTypeRequestPhoneNumber A button that sends the user's phone number when pressed; available only in private chats 
type KeyboardButtonTypeRequestPhoneNumber struct {
tdCommon

}

// MessageType return the string telegram-type of KeyboardButtonTypeRequestPhoneNumber 
func (keyboardButtonTypeRequestPhoneNumber *KeyboardButtonTypeRequestPhoneNumber) MessageType() string {
 return "keyboardButtonTypeRequestPhoneNumber" }

// KeyboardButtonTypeRequestLocation A button that sends the user's location when pressed; available only in private chats 
type KeyboardButtonTypeRequestLocation struct {
tdCommon

}

// MessageType return the string telegram-type of KeyboardButtonTypeRequestLocation 
func (keyboardButtonTypeRequestLocation *KeyboardButtonTypeRequestLocation) MessageType() string {
 return "keyboardButtonTypeRequestLocation" }

// KeyboardButton Represents a single button in a bot keyboard  
type KeyboardButton struct {
tdCommon
Text string `json:"text"` // Text of the button 
Type KeyboardButtonType `json:"type"` // Type of the button
}

// MessageType return the string telegram-type of KeyboardButton 
func (keyboardButton *KeyboardButton) MessageType() string {
 return "keyboardButton" }

// InlineKeyboardButtonTypeURL A button that opens a specified URL  
type InlineKeyboardButtonTypeURL struct {
tdCommon
URL string `json:"url"` // URL to open
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeURL 
func (inlineKeyboardButtonTypeURL *InlineKeyboardButtonTypeURL) MessageType() string {
 return "inlineKeyboardButtonTypeUrl" }

// InlineKeyboardButtonTypeCallback A button that sends a special callback query to a bot  
type InlineKeyboardButtonTypeCallback struct {
tdCommon
Data []byte `json:"data"` // Data to be sent to the bot via a callback query
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeCallback 
func (inlineKeyboardButtonTypeCallback *InlineKeyboardButtonTypeCallback) MessageType() string {
 return "inlineKeyboardButtonTypeCallback" }

// InlineKeyboardButtonTypeCallbackGame A button with a game that sends a special callback query to a bot. This button must be in the first column and row of the keyboard and can be attached only to a message with content of the type messageGame 
type InlineKeyboardButtonTypeCallbackGame struct {
tdCommon

}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeCallbackGame 
func (inlineKeyboardButtonTypeCallbackGame *InlineKeyboardButtonTypeCallbackGame) MessageType() string {
 return "inlineKeyboardButtonTypeCallbackGame" }

// InlineKeyboardButtonTypeSwitchInline A button that forces an inline query to the bot to be inserted in the input field  
type InlineKeyboardButtonTypeSwitchInline struct {
tdCommon
Query string `json:"query"` // Inline query to be sent to the bot 
InCurrentChat bool `json:"in_current_chat"` // True, if the inline query should be sent from the current chat
}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeSwitchInline 
func (inlineKeyboardButtonTypeSwitchInline *InlineKeyboardButtonTypeSwitchInline) MessageType() string {
 return "inlineKeyboardButtonTypeSwitchInline" }

// InlineKeyboardButtonTypeBuy A button to buy something. This button must be in the first column and row of the keyboard and can be attached only to a message with content of the type messageInvoice 
type InlineKeyboardButtonTypeBuy struct {
tdCommon

}

// MessageType return the string telegram-type of InlineKeyboardButtonTypeBuy 
func (inlineKeyboardButtonTypeBuy *InlineKeyboardButtonTypeBuy) MessageType() string {
 return "inlineKeyboardButtonTypeBuy" }

// InlineKeyboardButton Represents a single button in an inline keyboard  
type InlineKeyboardButton struct {
tdCommon
Text string `json:"text"` // Text of the button 
Type InlineKeyboardButtonType `json:"type"` // Type of the button
}

// MessageType return the string telegram-type of InlineKeyboardButton 
func (inlineKeyboardButton *InlineKeyboardButton) MessageType() string {
 return "inlineKeyboardButton" }

// ReplyMarkupRemoveKeyboard Instructs clients to remove the keyboard once this message has been received. This kind of keyboard can't be received in an incoming message; instead, UpdateChatReplyMarkup with message_id == 0 will be sent 
type ReplyMarkupRemoveKeyboard struct {
tdCommon
IsPersonal bool `json:"is_personal"` // True, if the keyboard is removed only for the mentioned users or the target user of a reply
}

// MessageType return the string telegram-type of ReplyMarkupRemoveKeyboard 
func (replyMarkupRemoveKeyboard *ReplyMarkupRemoveKeyboard) MessageType() string {
 return "replyMarkupRemoveKeyboard" }

// ReplyMarkupForceReply Instructs clients to force a reply to this message 
type ReplyMarkupForceReply struct {
tdCommon
IsPersonal bool `json:"is_personal"` // True, if a forced reply must automatically be shown to the current user. For outgoing messages, specify true to show the forced reply only for the mentioned users and for the target user of a reply
}

// MessageType return the string telegram-type of ReplyMarkupForceReply 
func (replyMarkupForceReply *ReplyMarkupForceReply) MessageType() string {
 return "replyMarkupForceReply" }

// ReplyMarkupShowKeyboard Contains a custom keyboard layout to quickly reply to bots 
type ReplyMarkupShowKeyboard struct {
tdCommon
Rows [][]KeyboardButton `json:"rows"` // A list of rows of bot keyboard buttons
ResizeKeyboard bool `json:"resize_keyboard"` // True, if the client needs to resize the keyboard vertically
OneTime bool `json:"one_time"` // True, if the client needs to hide the keyboard after use
IsPersonal bool `json:"is_personal"` // True, if the keyboard must automatically be shown to the current user. For outgoing messages, specify true to show the keyboard only for the mentioned users and for the target user of a reply
}

// MessageType return the string telegram-type of ReplyMarkupShowKeyboard 
func (replyMarkupShowKeyboard *ReplyMarkupShowKeyboard) MessageType() string {
 return "replyMarkupShowKeyboard" }

// ReplyMarkupInlineKeyboard Contains an inline keyboard layout 
type ReplyMarkupInlineKeyboard struct {
tdCommon
Rows [][]InlineKeyboardButton `json:"rows"` // A list of rows of inline keyboard buttons
}

// MessageType return the string telegram-type of ReplyMarkupInlineKeyboard 
func (replyMarkupInlineKeyboard *ReplyMarkupInlineKeyboard) MessageType() string {
 return "replyMarkupInlineKeyboard" }

// RichTextPlain A plain text  
type RichTextPlain struct {
tdCommon
Text string `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextPlain 
func (richTextPlain *RichTextPlain) MessageType() string {
 return "richTextPlain" }

// RichTextBold A bold rich text  
type RichTextBold struct {
tdCommon
Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextBold 
func (richTextBold *RichTextBold) MessageType() string {
 return "richTextBold" }

// RichTextItalic An italicized rich text  
type RichTextItalic struct {
tdCommon
Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextItalic 
func (richTextItalic *RichTextItalic) MessageType() string {
 return "richTextItalic" }

// RichTextUnderline An underlined rich text  
type RichTextUnderline struct {
tdCommon
Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextUnderline 
func (richTextUnderline *RichTextUnderline) MessageType() string {
 return "richTextUnderline" }

// RichTextStrikethrough A strike-through rich text  
type RichTextStrikethrough struct {
tdCommon
Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextStrikethrough 
func (richTextStrikethrough *RichTextStrikethrough) MessageType() string {
 return "richTextStrikethrough" }

// RichTextFixed A fixed-width rich text  
type RichTextFixed struct {
tdCommon
Text RichText `json:"text"` // Text
}

// MessageType return the string telegram-type of RichTextFixed 
func (richTextFixed *RichTextFixed) MessageType() string {
 return "richTextFixed" }

// RichTextURL A rich text URL link  
type RichTextURL struct {
tdCommon
Text RichText `json:"text"` // Text 
URL string `json:"url"` // URL
}

// MessageType return the string telegram-type of RichTextURL 
func (richTextURL *RichTextURL) MessageType() string {
 return "richTextUrl" }

// RichTextEmailAddress A rich text email link  
type RichTextEmailAddress struct {
tdCommon
Text RichText `json:"text"` // Text 
EmailAddress string `json:"email_address"` // Email address
}

// MessageType return the string telegram-type of RichTextEmailAddress 
func (richTextEmailAddress *RichTextEmailAddress) MessageType() string {
 return "richTextEmailAddress" }

// RichTexts A concatenation of rich texts  
type RichTexts struct {
tdCommon
Texts []RichText `json:"texts"` // Texts
}

// MessageType return the string telegram-type of RichTexts 
func (richTexts *RichTexts) MessageType() string {
 return "richTexts" }

// PageBlockTitle The title of a page  
type PageBlockTitle struct {
tdCommon
Title RichText `json:"title"` // Title
}

// MessageType return the string telegram-type of PageBlockTitle 
func (pageBlockTitle *PageBlockTitle) MessageType() string {
 return "pageBlockTitle" }

// PageBlockSubtitle The subtitle of a page  
type PageBlockSubtitle struct {
tdCommon
Subtitle RichText `json:"subtitle"` // Subtitle
}

// MessageType return the string telegram-type of PageBlockSubtitle 
func (pageBlockSubtitle *PageBlockSubtitle) MessageType() string {
 return "pageBlockSubtitle" }

// PageBlockAuthorDate The author and publishing date of a page  
type PageBlockAuthorDate struct {
tdCommon
Author RichText `json:"author"` // Author 
PublishDate int32 `json:"publish_date"` // Point in time (Unix timestamp) when the article was published; 0 if unknown
}

// MessageType return the string telegram-type of PageBlockAuthorDate 
func (pageBlockAuthorDate *PageBlockAuthorDate) MessageType() string {
 return "pageBlockAuthorDate" }

// PageBlockHeader A header  
type PageBlockHeader struct {
tdCommon
Header RichText `json:"header"` // Header
}

// MessageType return the string telegram-type of PageBlockHeader 
func (pageBlockHeader *PageBlockHeader) MessageType() string {
 return "pageBlockHeader" }

// PageBlockSubheader A subheader  
type PageBlockSubheader struct {
tdCommon
Subheader RichText `json:"subheader"` // Subheader
}

// MessageType return the string telegram-type of PageBlockSubheader 
func (pageBlockSubheader *PageBlockSubheader) MessageType() string {
 return "pageBlockSubheader" }

// PageBlockParagraph A text paragraph  
type PageBlockParagraph struct {
tdCommon
Text RichText `json:"text"` // Paragraph text
}

// MessageType return the string telegram-type of PageBlockParagraph 
func (pageBlockParagraph *PageBlockParagraph) MessageType() string {
 return "pageBlockParagraph" }

// PageBlockPreformatted A preformatted text paragraph  
type PageBlockPreformatted struct {
tdCommon
Text RichText `json:"text"` // Paragraph text 
Language string `json:"language"` // Programming language for which the text should be formatted
}

// MessageType return the string telegram-type of PageBlockPreformatted 
func (pageBlockPreformatted *PageBlockPreformatted) MessageType() string {
 return "pageBlockPreformatted" }

// PageBlockFooter The footer of a page  
type PageBlockFooter struct {
tdCommon
Footer RichText `json:"footer"` // Footer
}

// MessageType return the string telegram-type of PageBlockFooter 
func (pageBlockFooter *PageBlockFooter) MessageType() string {
 return "pageBlockFooter" }

// PageBlockDivider An empty block separating a page 
type PageBlockDivider struct {
tdCommon

}

// MessageType return the string telegram-type of PageBlockDivider 
func (pageBlockDivider *PageBlockDivider) MessageType() string {
 return "pageBlockDivider" }

// PageBlockAnchor An invisible anchor on a page, which can be used in a URL to open the page from the specified anchor  
type PageBlockAnchor struct {
tdCommon
Name string `json:"name"` // Name of the anchor
}

// MessageType return the string telegram-type of PageBlockAnchor 
func (pageBlockAnchor *PageBlockAnchor) MessageType() string {
 return "pageBlockAnchor" }

// PageBlockList A list of texts  
type PageBlockList struct {
tdCommon
Items []RichText `json:"items"` // Texts 
IsOrdered bool `json:"is_ordered"` // True, if the items should be marked with numbers
}

// MessageType return the string telegram-type of PageBlockList 
func (pageBlockList *PageBlockList) MessageType() string {
 return "pageBlockList" }

// PageBlockBlockQuote A block quote  
type PageBlockBlockQuote struct {
tdCommon
Text RichText `json:"text"` // Quote text 
Caption RichText `json:"caption"` // Quote caption
}

// MessageType return the string telegram-type of PageBlockBlockQuote 
func (pageBlockBlockQuote *PageBlockBlockQuote) MessageType() string {
 return "pageBlockBlockQuote" }

// PageBlockPullQuote A pull quote  
type PageBlockPullQuote struct {
tdCommon
Text RichText `json:"text"` // Quote text 
Caption RichText `json:"caption"` // Quote caption
}

// MessageType return the string telegram-type of PageBlockPullQuote 
func (pageBlockPullQuote *PageBlockPullQuote) MessageType() string {
 return "pageBlockPullQuote" }

// PageBlockAnimation An animation  
type PageBlockAnimation struct {
tdCommon
Caption RichText `json:"caption"` // Animation caption 
NeedAutoplay bool `json:"need_autoplay"` // True, if the animation should be played automatically
Animation Animation `json:"animation"` // Animation file; may be null 
}

// MessageType return the string telegram-type of PageBlockAnimation 
func (pageBlockAnimation *PageBlockAnimation) MessageType() string {
 return "pageBlockAnimation" }

// PageBlockAudio An audio file  
type PageBlockAudio struct {
tdCommon
Audio Audio `json:"audio"` // Audio file; may be null 
Caption RichText `json:"caption"` // Audio file caption
}

// MessageType return the string telegram-type of PageBlockAudio 
func (pageBlockAudio *PageBlockAudio) MessageType() string {
 return "pageBlockAudio" }

// PageBlockPhoto A photo  
type PageBlockPhoto struct {
tdCommon
Photo Photo `json:"photo"` // Photo file; may be null 
Caption RichText `json:"caption"` // Photo caption
}

// MessageType return the string telegram-type of PageBlockPhoto 
func (pageBlockPhoto *PageBlockPhoto) MessageType() string {
 return "pageBlockPhoto" }

// PageBlockVideo A video  
type PageBlockVideo struct {
tdCommon
Video Video `json:"video"` // Video file; may be null 
Caption RichText `json:"caption"` // Video caption 
NeedAutoplay bool `json:"need_autoplay"` // True, if the video should be played automatically 
IsLooped bool `json:"is_looped"` // True, if the video should be looped
}

// MessageType return the string telegram-type of PageBlockVideo 
func (pageBlockVideo *PageBlockVideo) MessageType() string {
 return "pageBlockVideo" }

// PageBlockCover A page cover  
type PageBlockCover struct {
tdCommon
Cover PageBlock `json:"cover"` // Cover
}

// MessageType return the string telegram-type of PageBlockCover 
func (pageBlockCover *PageBlockCover) MessageType() string {
 return "pageBlockCover" }

// PageBlockEmbedded An embedded web page  
type PageBlockEmbedded struct {
tdCommon
Width int32 `json:"width"` // Block width 
Height int32 `json:"height"` // Block height 
Caption RichText `json:"caption"` // Block caption 
IsFullWidth bool `json:"is_full_width"` // True, if the block should be full width 
AllowScrolling bool `json:"allow_scrolling"` // True, if scrolling should be allowed
URL string `json:"url"` // Web page URL, if available 
HTML string `json:"html"` // HTML-markup of the embedded page 
PosterPhoto Photo `json:"poster_photo"` // Poster photo, if available; may be null 
}

// MessageType return the string telegram-type of PageBlockEmbedded 
func (pageBlockEmbedded *PageBlockEmbedded) MessageType() string {
 return "pageBlockEmbedded" }

// PageBlockEmbeddedPost An embedded post  
type PageBlockEmbeddedPost struct {
tdCommon
URL string `json:"url"` // Web page URL 
Author string `json:"author"` // Post author 
AuthorPhoto Photo `json:"author_photo"` // Post author photo 
Date int32 `json:"date"` // Point in time (Unix timestamp) when the post was created; 0 if unknown 
PageBlocks []PageBlock `json:"page_blocks"` // Post content 
Caption RichText `json:"caption"` // Post caption
}

// MessageType return the string telegram-type of PageBlockEmbeddedPost 
func (pageBlockEmbeddedPost *PageBlockEmbeddedPost) MessageType() string {
 return "pageBlockEmbeddedPost" }

// PageBlockCollage A collage  
type PageBlockCollage struct {
tdCommon
PageBlocks []PageBlock `json:"page_blocks"` // Collage item contents 
Caption RichText `json:"caption"` // Block caption
}

// MessageType return the string telegram-type of PageBlockCollage 
func (pageBlockCollage *PageBlockCollage) MessageType() string {
 return "pageBlockCollage" }

// PageBlockSlideshow A slideshow  
type PageBlockSlideshow struct {
tdCommon
PageBlocks []PageBlock `json:"page_blocks"` // Slideshow item contents 
Caption RichText `json:"caption"` // Block caption
}

// MessageType return the string telegram-type of PageBlockSlideshow 
func (pageBlockSlideshow *PageBlockSlideshow) MessageType() string {
 return "pageBlockSlideshow" }

// PageBlockChatLink A link to a chat  
type PageBlockChatLink struct {
tdCommon
Title string `json:"title"` // Chat title 
Photo ChatPhoto `json:"photo"` // Chat photo; may be null 
Username string `json:"username"` // Chat username, by which all other information about the chat should be resolved
}

// MessageType return the string telegram-type of PageBlockChatLink 
func (pageBlockChatLink *PageBlockChatLink) MessageType() string {
 return "pageBlockChatLink" }

// WebPageInstantView Describes an instant view page for a web page  
type WebPageInstantView struct {
tdCommon
PageBlocks []PageBlock `json:"page_blocks"` // Content of the web page 
IsFull bool `json:"is_full"` // True, if the instant view contains the full page. A network request might be needed to get the full web page instant view
}

// MessageType return the string telegram-type of WebPageInstantView 
func (webPageInstantView *WebPageInstantView) MessageType() string {
 return "webPageInstantView" }

// WebPage Describes a web page preview  
type WebPage struct {
tdCommon
DisplayURL string `json:"display_url"` // URL to display
Type string `json:"type"` // Type of the web page. Can be: article, photo, audio, video, document, profile, app, or something else
Duration int32 `json:"duration"` // Duration of the content, in seconds
Video Video `json:"video"` // Preview of the content as a video, if available; may be null
VideoNote VideoNote `json:"video_note"` // Preview of the content as a video note, if available; may be null
Description string `json:"description"` // 
EmbedWidth int32 `json:"embed_width"` // Width of the embedded preview
Author string `json:"author"` // Author of the content
Sticker Sticker `json:"sticker"` // Preview of the content as a sticker for small WEBP files, if available; may be null
SiteName string `json:"site_name"` // Short name of the site (e.g., Google Docs, App Store) 
Photo Photo `json:"photo"` // Image representing the content; may be null
EmbedType string `json:"embed_type"` // MIME type of the embedded preview, (e.g., text/html or video/mp4)
Document Document `json:"document"` // Preview of the content as a document, if available (currently only available for small PDF files and ZIP archives); may be null
VoiceNote VoiceNote `json:"voice_note"` // Preview of the content as a voice note, if available; may be null
HasInstantView bool `json:"has_instant_view"` // True, if the web page has an instant view
URL string `json:"url"` // Original URL of the link 
Title string `json:"title"` // Title of the content 
EmbedURL string `json:"embed_url"` // URL to show in the embedded preview
EmbedHeight int32 `json:"embed_height"` // Height of the embedded preview
Animation Animation `json:"animation"` // Preview of the content as an animation, if available; may be null
Audio Audio `json:"audio"` // Preview of the content as an audio file, if available; may be null
}

// MessageType return the string telegram-type of WebPage 
func (webPage *WebPage) MessageType() string {
 return "webPage" }

// LabeledPricePart Portion of the price of a product (e.g., "delivery cost", "tax amount")  
type LabeledPricePart struct {
tdCommon
Label string `json:"label"` // Label for this portion of the product price 
Amount int64 `json:"amount"` // Currency amount in minimal quantity of the currency
}

// MessageType return the string telegram-type of LabeledPricePart 
func (labeledPricePart *LabeledPricePart) MessageType() string {
 return "labeledPricePart" }

// Invoice Product invoice  
type Invoice struct {
tdCommon
Currency string `json:"currency"` // ISO 4217 currency code 
NeedName bool `json:"need_name"` // True, if the user's name is needed for payment 
NeedPhoneNumber bool `json:"need_phone_number"` // True, if the user's phone number is needed for payment 
IsFlexible bool `json:"is_flexible"` // True, if the total price depends on the shipping method
PriceParts []LabeledPricePart `json:"price_parts"` // A list of objects used to calculate the total price of the product 
IsTest bool `json:"is_test"` // True, if the payment is a test payment
NeedEmailAddress bool `json:"need_email_address"` // True, if the user's email address is needed for payment
NeedShippingAddress bool `json:"need_shipping_address"` // True, if the user's shipping address is needed for payment 
SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider"` // True, if the user's phone number will be sent to the provider
SendEmailAddressToProvider bool `json:"send_email_address_to_provider"` // True, if the user's email address will be sent to the provider 
}

// MessageType return the string telegram-type of Invoice 
func (invoice *Invoice) MessageType() string {
 return "invoice" }

// ShippingAddress Describes a shipping address  
type ShippingAddress struct {
tdCommon
CountryCode string `json:"country_code"` // Two-letter ISO 3166-1 alpha-2 country code 
State string `json:"state"` // State, if applicable 
City string `json:"city"` // City 
StreetLine1 string `json:"street_line1"` // First line of the address 
StreetLine2 string `json:"street_line2"` // Second line of the address 
PostalCode string `json:"postal_code"` // Address postal code
}

// MessageType return the string telegram-type of ShippingAddress 
func (shippingAddress *ShippingAddress) MessageType() string {
 return "shippingAddress" }

// OrderInfo Order information  
type OrderInfo struct {
tdCommon
EmailAddress string `json:"email_address"` // Email address of the user 
ShippingAddress ShippingAddress `json:"shipping_address"` // Shipping address for this order; may be null
Name string `json:"name"` // Name of the user 
PhoneNumber string `json:"phone_number"` // Phone number of the user 
}

// MessageType return the string telegram-type of OrderInfo 
func (orderInfo *OrderInfo) MessageType() string {
 return "orderInfo" }

// ShippingOption One shipping option  
type ShippingOption struct {
tdCommon
ID string `json:"id"` // Shipping option identifier 
Title string `json:"title"` // Option title 
PriceParts []LabeledPricePart `json:"price_parts"` // A list of objects used to calculate the total shipping costs
}

// MessageType return the string telegram-type of ShippingOption 
func (shippingOption *ShippingOption) MessageType() string {
 return "shippingOption" }

// SavedCredentials Contains information about saved card credentials  
type SavedCredentials struct {
tdCommon
ID string `json:"id"` // Unique identifier of the saved credentials 
Title string `json:"title"` // Title of the saved credentials
}

// MessageType return the string telegram-type of SavedCredentials 
func (savedCredentials *SavedCredentials) MessageType() string {
 return "savedCredentials" }

// InputCredentialsSaved Applies if a user chooses some previously saved payment credentials. To use their previously saved credentials, the user must have a valid temporary password  
type InputCredentialsSaved struct {
tdCommon
SavedCredentialsID string `json:"saved_credentials_id"` // Identifier of the saved credentials
}

// MessageType return the string telegram-type of InputCredentialsSaved 
func (inputCredentialsSaved *InputCredentialsSaved) MessageType() string {
 return "inputCredentialsSaved" }

// InputCredentialsNew Applies if a user enters new credentials on a payment provider website  
type InputCredentialsNew struct {
tdCommon
Data string `json:"data"` // Contains JSON-encoded data with a credential identifier from the payment provider 
AllowSave bool `json:"allow_save"` // True, if the credential identifier can be saved on the server side
}

// MessageType return the string telegram-type of InputCredentialsNew 
func (inputCredentialsNew *InputCredentialsNew) MessageType() string {
 return "inputCredentialsNew" }

// InputCredentialsAndroidPay Applies if a user enters new credentials using Android Pay  
type InputCredentialsAndroidPay struct {
tdCommon
Data string `json:"data"` // JSON-encoded data with the credential identifier
}

// MessageType return the string telegram-type of InputCredentialsAndroidPay 
func (inputCredentialsAndroidPay *InputCredentialsAndroidPay) MessageType() string {
 return "inputCredentialsAndroidPay" }

// InputCredentialsApplePay Applies if a user enters new credentials using Apple Pay  
type InputCredentialsApplePay struct {
tdCommon
Data string `json:"data"` // JSON-encoded data with the credential identifier
}

// MessageType return the string telegram-type of InputCredentialsApplePay 
func (inputCredentialsApplePay *InputCredentialsApplePay) MessageType() string {
 return "inputCredentialsApplePay" }

// PaymentsProviderStripe Stripe payment provider  
type PaymentsProviderStripe struct {
tdCommon
NeedPostalCode bool `json:"need_postal_code"` // True, if the user ZIP/postal code must be provided 
NeedCardholderName bool `json:"need_cardholder_name"` // True, if the cardholder name must be provided
PublishableKey string `json:"publishable_key"` // Stripe API publishable key 
NeedCountry bool `json:"need_country"` // True, if the user country must be provided 
}

// MessageType return the string telegram-type of PaymentsProviderStripe 
func (paymentsProviderStripe *PaymentsProviderStripe) MessageType() string {
 return "paymentsProviderStripe" }

// PaymentForm Contains information about an invoice payment form  
type PaymentForm struct {
tdCommon
Invoice Invoice `json:"invoice"` // Full information of the invoice 
URL string `json:"url"` // Payment form URL 
PaymentsProvider PaymentsProviderStripe `json:"payments_provider"` // Contains information about the payment provider, if available, to support it natively without the need for opening the URL; may be null
SavedOrderInfo OrderInfo `json:"saved_order_info"` // Saved server-side order information; may be null 
SavedCredentials SavedCredentials `json:"saved_credentials"` // Contains information about saved card credentials; may be null 
CanSaveCredentials bool `json:"can_save_credentials"` // True, if the user can choose to save credentials 
NeedPassword bool `json:"need_password"` // True, if the user will be able to save credentials protected by a password they set up
}

// MessageType return the string telegram-type of PaymentForm 
func (paymentForm *PaymentForm) MessageType() string {
 return "paymentForm" }

// ValidatedOrderInfo Contains a temporary identifier of validated order information, which is stored for one hour. Also contains the available shipping options  
type ValidatedOrderInfo struct {
tdCommon
OrderInfoID string `json:"order_info_id"` // Temporary identifier of the order information 
ShippingOptions []ShippingOption `json:"shipping_options"` // Available shipping options
}

// MessageType return the string telegram-type of ValidatedOrderInfo 
func (validatedOrderInfo *ValidatedOrderInfo) MessageType() string {
 return "validatedOrderInfo" }

// PaymentResult Contains the result of a payment request  
type PaymentResult struct {
tdCommon
Success bool `json:"success"` // True, if the payment request was successful; otherwise the verification_url will be not empty 
VerificationURL string `json:"verification_url"` // URL for additional payment credentials verification
}

// MessageType return the string telegram-type of PaymentResult 
func (paymentResult *PaymentResult) MessageType() string {
 return "paymentResult" }

// PaymentReceipt Contains information about a successful payment  
type PaymentReceipt struct {
tdCommon
Date int32 `json:"date"` // Point in time (Unix timestamp) when the payment was made 
PaymentsProviderUserID int32 `json:"payments_provider_user_id"` // User identifier of the payment provider bot 
Invoice Invoice `json:"invoice"` // Contains information about the invoice
OrderInfo OrderInfo `json:"order_info"` // Contains order information; may be null 
ShippingOption ShippingOption `json:"shipping_option"` // Chosen shipping option; may be null 
CredentialsTitle string `json:"credentials_title"` // Title of the saved credentials
}

// MessageType return the string telegram-type of PaymentReceipt 
func (paymentReceipt *PaymentReceipt) MessageType() string {
 return "paymentReceipt" }

// MessageText A text message  
type MessageText struct {
tdCommon
Text FormattedText `json:"text"` // Text of the message 
WebPage WebPage `json:"web_page"` // A preview of the web page that's mentioned in the text; may be null
}

// MessageType return the string telegram-type of MessageText 
func (messageText *MessageText) MessageType() string {
 return "messageText" }

// MessageAnimation An animation message (GIF-style).  
type MessageAnimation struct {
tdCommon
Animation Animation `json:"animation"` // Message content 
Caption FormattedText `json:"caption"` // Animation caption 
IsSecret bool `json:"is_secret"` // True, if the animation thumbnail must be blurred and the animation must be shown only while tapped
}

// MessageType return the string telegram-type of MessageAnimation 
func (messageAnimation *MessageAnimation) MessageType() string {
 return "messageAnimation" }

// MessageAudio An audio message  
type MessageAudio struct {
tdCommon
Audio Audio `json:"audio"` // Message content 
Caption FormattedText `json:"caption"` // Audio caption
}

// MessageType return the string telegram-type of MessageAudio 
func (messageAudio *MessageAudio) MessageType() string {
 return "messageAudio" }

// MessageDocument A document message (general file)  
type MessageDocument struct {
tdCommon
Document Document `json:"document"` // Message content 
Caption FormattedText `json:"caption"` // Document caption
}

// MessageType return the string telegram-type of MessageDocument 
func (messageDocument *MessageDocument) MessageType() string {
 return "messageDocument" }

// MessagePhoto A photo message  
type MessagePhoto struct {
tdCommon
Caption FormattedText `json:"caption"` // Photo caption 
IsSecret bool `json:"is_secret"` // True, if the photo must be blurred and must be shown only while tapped
Photo Photo `json:"photo"` // Message content 
}

// MessageType return the string telegram-type of MessagePhoto 
func (messagePhoto *MessagePhoto) MessageType() string {
 return "messagePhoto" }

// MessageExpiredPhoto An expired photo message (self-destructed after TTL has elapsed) 
type MessageExpiredPhoto struct {
tdCommon

}

// MessageType return the string telegram-type of MessageExpiredPhoto 
func (messageExpiredPhoto *MessageExpiredPhoto) MessageType() string {
 return "messageExpiredPhoto" }

// MessageSticker A sticker message  
type MessageSticker struct {
tdCommon
Sticker Sticker `json:"sticker"` // Message content
}

// MessageType return the string telegram-type of MessageSticker 
func (messageSticker *MessageSticker) MessageType() string {
 return "messageSticker" }

// MessageVideo A video message  
type MessageVideo struct {
tdCommon
Video Video `json:"video"` // Message content 
Caption FormattedText `json:"caption"` // Video caption 
IsSecret bool `json:"is_secret"` // True, if the video thumbnail must be blurred and the video must be shown only while tapped
}

// MessageType return the string telegram-type of MessageVideo 
func (messageVideo *MessageVideo) MessageType() string {
 return "messageVideo" }

// MessageExpiredVideo An expired video message (self-destructed after TTL has elapsed) 
type MessageExpiredVideo struct {
tdCommon

}

// MessageType return the string telegram-type of MessageExpiredVideo 
func (messageExpiredVideo *MessageExpiredVideo) MessageType() string {
 return "messageExpiredVideo" }

// MessageVideoNote A video note message  
type MessageVideoNote struct {
tdCommon
VideoNote VideoNote `json:"video_note"` // Message content 
IsViewed bool `json:"is_viewed"` // True, if at least one of the recipients has viewed the video note 
IsSecret bool `json:"is_secret"` // True, if the video note thumbnail must be blurred and the video note must be shown only while tapped
}

// MessageType return the string telegram-type of MessageVideoNote 
func (messageVideoNote *MessageVideoNote) MessageType() string {
 return "messageVideoNote" }

// MessageVoiceNote A voice note message  
type MessageVoiceNote struct {
tdCommon
VoiceNote VoiceNote `json:"voice_note"` // Message content 
Caption FormattedText `json:"caption"` // Voice note caption 
IsListened bool `json:"is_listened"` // True, if at least one of the recipients has listened to the voice note
}

// MessageType return the string telegram-type of MessageVoiceNote 
func (messageVoiceNote *MessageVoiceNote) MessageType() string {
 return "messageVoiceNote" }

// MessageLocation A message with a location  
type MessageLocation struct {
tdCommon
Location Location `json:"location"` // Message content 
LivePeriod int32 `json:"live_period"` // Time relative to the message sent date until which the location can be updated, in seconds
ExpiresIn int32 `json:"expires_in"` // Left time for which the location can be updated, in seconds. updateMessageContent is not sent when this field changes
}

// MessageType return the string telegram-type of MessageLocation 
func (messageLocation *MessageLocation) MessageType() string {
 return "messageLocation" }

// MessageVenue A message with information about a venue  
type MessageVenue struct {
tdCommon
Venue Venue `json:"venue"` // Message content
}

// MessageType return the string telegram-type of MessageVenue 
func (messageVenue *MessageVenue) MessageType() string {
 return "messageVenue" }

// MessageContact A message with a user contact  
type MessageContact struct {
tdCommon
Contact Contact `json:"contact"` // Message content
}

// MessageType return the string telegram-type of MessageContact 
func (messageContact *MessageContact) MessageType() string {
 return "messageContact" }

// MessageGame A message with a game  
type MessageGame struct {
tdCommon
Game Game `json:"game"` // Game
}

// MessageType return the string telegram-type of MessageGame 
func (messageGame *MessageGame) MessageType() string {
 return "messageGame" }

// MessageInvoice A message with an invoice from a bot  
type MessageInvoice struct {
tdCommon
Description string `json:"description"` // 
StartParameter string `json:"start_parameter"` // Unique invoice bot start_parameter. To share an invoice use the URL https://t.me/{bot_username}?start={start_parameter} 
IsTest bool `json:"is_test"` // True, if the invoice is a test invoice
ReceiptMessageID int64 `json:"receipt_message_id"` // The identifier of the message with the receipt, after the product has been purchased
Title string `json:"title"` // Product title 
Photo Photo `json:"photo"` // Product photo; may be null 
Currency string `json:"currency"` // Currency for the product price 
TotalAmount int64 `json:"total_amount"` // Product total price in the minimal quantity of the currency
NeedShippingAddress bool `json:"need_shipping_address"` // True, if the shipping address should be specified 
}

// MessageType return the string telegram-type of MessageInvoice 
func (messageInvoice *MessageInvoice) MessageType() string {
 return "messageInvoice" }

// MessageCall A message with information about an ended call  
type MessageCall struct {
tdCommon
Duration int32 `json:"duration"` // Call duration, in seconds
DiscardReason CallDiscardReason `json:"discard_reason"` // Reason why the call was discarded 
}

// MessageType return the string telegram-type of MessageCall 
func (messageCall *MessageCall) MessageType() string {
 return "messageCall" }

// MessageBasicGroupChatCreate A newly created basic group  
type MessageBasicGroupChatCreate struct {
tdCommon
Title string `json:"title"` // Title of the basic group 
MemberUserIDs []int32 `json:"member_user_ids"` // User identifiers of members in the basic group
}

// MessageType return the string telegram-type of MessageBasicGroupChatCreate 
func (messageBasicGroupChatCreate *MessageBasicGroupChatCreate) MessageType() string {
 return "messageBasicGroupChatCreate" }

// MessageSupergroupChatCreate A newly created supergroup or channel  
type MessageSupergroupChatCreate struct {
tdCommon
Title string `json:"title"` // Title of the supergroup or channel
}

// MessageType return the string telegram-type of MessageSupergroupChatCreate 
func (messageSupergroupChatCreate *MessageSupergroupChatCreate) MessageType() string {
 return "messageSupergroupChatCreate" }

// MessageChatChangeTitle An updated chat title  
type MessageChatChangeTitle struct {
tdCommon
Title string `json:"title"` // New chat title
}

// MessageType return the string telegram-type of MessageChatChangeTitle 
func (messageChatChangeTitle *MessageChatChangeTitle) MessageType() string {
 return "messageChatChangeTitle" }

// MessageChatChangePhoto An updated chat photo  
type MessageChatChangePhoto struct {
tdCommon
Photo Photo `json:"photo"` // New chat photo
}

// MessageType return the string telegram-type of MessageChatChangePhoto 
func (messageChatChangePhoto *MessageChatChangePhoto) MessageType() string {
 return "messageChatChangePhoto" }

// MessageChatDeletePhoto A deleted chat photo 
type MessageChatDeletePhoto struct {
tdCommon

}

// MessageType return the string telegram-type of MessageChatDeletePhoto 
func (messageChatDeletePhoto *MessageChatDeletePhoto) MessageType() string {
 return "messageChatDeletePhoto" }

// MessageChatAddMembers New chat members were added  
type MessageChatAddMembers struct {
tdCommon
MemberUserIDs []int32 `json:"member_user_ids"` // User identifiers of the new members
}

// MessageType return the string telegram-type of MessageChatAddMembers 
func (messageChatAddMembers *MessageChatAddMembers) MessageType() string {
 return "messageChatAddMembers" }

// MessageChatJoinByLink A new member joined the chat by invite link 
type MessageChatJoinByLink struct {
tdCommon

}

// MessageType return the string telegram-type of MessageChatJoinByLink 
func (messageChatJoinByLink *MessageChatJoinByLink) MessageType() string {
 return "messageChatJoinByLink" }

// MessageChatDeleteMember A chat member was deleted  
type MessageChatDeleteMember struct {
tdCommon
UserID int32 `json:"user_id"` // User identifier of the deleted chat member
}

// MessageType return the string telegram-type of MessageChatDeleteMember 
func (messageChatDeleteMember *MessageChatDeleteMember) MessageType() string {
 return "messageChatDeleteMember" }

// MessageChatUpgradeTo A basic group was upgraded to a supergroup and was deactivated as the result  
type MessageChatUpgradeTo struct {
tdCommon
SupergroupID int32 `json:"supergroup_id"` // Identifier of the supergroup to which the basic group was upgraded
}

// MessageType return the string telegram-type of MessageChatUpgradeTo 
func (messageChatUpgradeTo *MessageChatUpgradeTo) MessageType() string {
 return "messageChatUpgradeTo" }

// MessageChatUpgradeFrom A supergroup has been created from a basic group  
type MessageChatUpgradeFrom struct {
tdCommon
BasicGroupID int32 `json:"basic_group_id"` // The identifier of the original basic group
Title string `json:"title"` // Title of the newly created supergroup 
}

// MessageType return the string telegram-type of MessageChatUpgradeFrom 
func (messageChatUpgradeFrom *MessageChatUpgradeFrom) MessageType() string {
 return "messageChatUpgradeFrom" }

// MessagePinMessage A message has been pinned  
type MessagePinMessage struct {
tdCommon
MessageID int64 `json:"message_id"` // Identifier of the pinned message, can be an identifier of a deleted message
}

// MessageType return the string telegram-type of MessagePinMessage 
func (messagePinMessage *MessagePinMessage) MessageType() string {
 return "messagePinMessage" }

// MessageScreenshotTaken A screenshot of a message in the chat has been taken 
type MessageScreenshotTaken struct {
tdCommon

}

// MessageType return the string telegram-type of MessageScreenshotTaken 
func (messageScreenshotTaken *MessageScreenshotTaken) MessageType() string {
 return "messageScreenshotTaken" }

// MessageChatSetTTL The TTL (Time To Live) setting messages in a secret chat has been changed  
type MessageChatSetTTL struct {
tdCommon
TTL int32 `json:"ttl"` // New TTL
}

// MessageType return the string telegram-type of MessageChatSetTTL 
func (messageChatSetTTL *MessageChatSetTTL) MessageType() string {
 return "messageChatSetTtl" }

// MessageCustomServiceAction A non-standard action has happened in the chat  
type MessageCustomServiceAction struct {
tdCommon
Text string `json:"text"` // Message text to be shown in the chat
}

// MessageType return the string telegram-type of MessageCustomServiceAction 
func (messageCustomServiceAction *MessageCustomServiceAction) MessageType() string {
 return "messageCustomServiceAction" }

// MessageGameScore A new high score was achieved in a game  
type MessageGameScore struct {
tdCommon
GameMessageID int64 `json:"game_message_id"` // Identifier of the message with the game, can be an identifier of a deleted message 
GameID int64 `json:"game_id"` // Identifier of the game, may be different from the games presented in the message with the game 
Score int32 `json:"score"` // New score
}

// MessageType return the string telegram-type of MessageGameScore 
func (messageGameScore *MessageGameScore) MessageType() string {
 return "messageGameScore" }

// MessagePaymentSuccessful A payment has been completed  
type MessagePaymentSuccessful struct {
tdCommon
InvoiceMessageID int64 `json:"invoice_message_id"` // Identifier of the message with the corresponding invoice; can be an identifier of a deleted message 
Currency string `json:"currency"` // Currency for the price of the product 
TotalAmount int64 `json:"total_amount"` // Total price for the product, in the minimal quantity of the currency
}

// MessageType return the string telegram-type of MessagePaymentSuccessful 
func (messagePaymentSuccessful *MessagePaymentSuccessful) MessageType() string {
 return "messagePaymentSuccessful" }

// MessagePaymentSuccessfulBot A payment has been completed; for bots only  
type MessagePaymentSuccessfulBot struct {
tdCommon
TotalAmount int64 `json:"total_amount"` // Total price for the product, in the minimal quantity of the currency 
InvoicePayload []byte `json:"invoice_payload"` // Invoice payload 
ShippingOptionID string `json:"shipping_option_id"` // Identifier of the shipping option chosen by the user, may be empty if not applicable 
OrderInfo OrderInfo `json:"order_info"` // Information about the order; may be null
TelegramPaymentChargeID string `json:"telegram_payment_charge_id"` // Telegram payment identifier 
ProviderPaymentChargeID string `json:"provider_payment_charge_id"` // Provider payment identifier
InvoiceMessageID int64 `json:"invoice_message_id"` // Identifier of the message with the corresponding invoice; can be an identifier of a deleted message 
Currency string `json:"currency"` // Currency for price of the product
}

// MessageType return the string telegram-type of MessagePaymentSuccessfulBot 
func (messagePaymentSuccessfulBot *MessagePaymentSuccessfulBot) MessageType() string {
 return "messagePaymentSuccessfulBot" }

// MessageContactRegistered A contact has registered with Telegram 
type MessageContactRegistered struct {
tdCommon

}

// MessageType return the string telegram-type of MessageContactRegistered 
func (messageContactRegistered *MessageContactRegistered) MessageType() string {
 return "messageContactRegistered" }

// MessageWebsiteConnected The current user has connected a website by logging in using Telegram Login Widget on it  
type MessageWebsiteConnected struct {
tdCommon
DomainName string `json:"domain_name"` // Domain name of the connected website
}

// MessageType return the string telegram-type of MessageWebsiteConnected 
func (messageWebsiteConnected *MessageWebsiteConnected) MessageType() string {
 return "messageWebsiteConnected" }

// MessageUnsupported Message content that is not supported by the client 
type MessageUnsupported struct {
tdCommon

}

// MessageType return the string telegram-type of MessageUnsupported 
func (messageUnsupported *MessageUnsupported) MessageType() string {
 return "messageUnsupported" }

// TextEntityTypeMention A mention of a user by their username 
type TextEntityTypeMention struct {
tdCommon

}

// MessageType return the string telegram-type of TextEntityTypeMention 
func (textEntityTypeMention *TextEntityTypeMention) MessageType() string {
 return "textEntityTypeMention" }

// TextEntityTypeHashtag A hashtag text, beginning with "#" 
type TextEntityTypeHashtag struct {
tdCommon

}

// MessageType return the string telegram-type of TextEntityTypeHashtag 
func (textEntityTypeHashtag *TextEntityTypeHashtag) MessageType() string {
 return "textEntityTypeHashtag" }

// TextEntityTypeCashtag A cashtag text, beginning with "$" and consisting of capital english letters (i.e. "$USD") 
type TextEntityTypeCashtag struct {
tdCommon

}

// MessageType return the string telegram-type of TextEntityTypeCashtag 
func (textEntityTypeCashtag *TextEntityTypeCashtag) MessageType() string {
 return "textEntityTypeCashtag" }

// TextEntityTypeBotCommand A bot command, beginning with "/". This shouldn't be highlighted if there are no bots in the chat 
type TextEntityTypeBotCommand struct {
tdCommon

}

// MessageType return the string telegram-type of TextEntityTypeBotCommand 
func (textEntityTypeBotCommand *TextEntityTypeBotCommand) MessageType() string {
 return "textEntityTypeBotCommand" }

// TextEntityTypeURL An HTTP URL 
type TextEntityTypeURL struct {
tdCommon

}

// MessageType return the string telegram-type of TextEntityTypeURL 
func (textEntityTypeURL *TextEntityTypeURL) MessageType() string {
 return "textEntityTypeUrl" }

// TextEntityTypeEmailAddress An email address 
type TextEntityTypeEmailAddress struct {
tdCommon

}

// MessageType return the string telegram-type of TextEntityTypeEmailAddress 
func (textEntityTypeEmailAddress *TextEntityTypeEmailAddress) MessageType() string {
 return "textEntityTypeEmailAddress" }

// TextEntityTypeBold A bold text 
type TextEntityTypeBold struct {
tdCommon

}

// MessageType return the string telegram-type of TextEntityTypeBold 
func (textEntityTypeBold *TextEntityTypeBold) MessageType() string {
 return "textEntityTypeBold" }

// TextEntityTypeItalic An italic text 
type TextEntityTypeItalic struct {
tdCommon

}

// MessageType return the string telegram-type of TextEntityTypeItalic 
func (textEntityTypeItalic *TextEntityTypeItalic) MessageType() string {
 return "textEntityTypeItalic" }

// TextEntityTypeCode Text that must be formatted as if inside a code HTML tag 
type TextEntityTypeCode struct {
tdCommon

}

// MessageType return the string telegram-type of TextEntityTypeCode 
func (textEntityTypeCode *TextEntityTypeCode) MessageType() string {
 return "textEntityTypeCode" }

// TextEntityTypePre Text that must be formatted as if inside a pre HTML tag 
type TextEntityTypePre struct {
tdCommon

}

// MessageType return the string telegram-type of TextEntityTypePre 
func (textEntityTypePre *TextEntityTypePre) MessageType() string {
 return "textEntityTypePre" }

// TextEntityTypePreCode Text that must be formatted as if inside pre, and code HTML tags  
type TextEntityTypePreCode struct {
tdCommon
Language string `json:"language"` // Programming language of the code; as defined by the sender
}

// MessageType return the string telegram-type of TextEntityTypePreCode 
func (textEntityTypePreCode *TextEntityTypePreCode) MessageType() string {
 return "textEntityTypePreCode" }

// TextEntityTypeTextURL A text description shown instead of a raw URL  
type TextEntityTypeTextURL struct {
tdCommon
URL string `json:"url"` // URL to be opened when the link is clicked
}

// MessageType return the string telegram-type of TextEntityTypeTextURL 
func (textEntityTypeTextURL *TextEntityTypeTextURL) MessageType() string {
 return "textEntityTypeTextUrl" }

// TextEntityTypeMentionName A text shows instead of a raw mention of the user (e.g., when the user has no username)  
type TextEntityTypeMentionName struct {
tdCommon
UserID int32 `json:"user_id"` // Identifier of the mentioned user
}

// MessageType return the string telegram-type of TextEntityTypeMentionName 
func (textEntityTypeMentionName *TextEntityTypeMentionName) MessageType() string {
 return "textEntityTypeMentionName" }

// TextEntityTypePhoneNumber A phone number 
type TextEntityTypePhoneNumber struct {
tdCommon

}

// MessageType return the string telegram-type of TextEntityTypePhoneNumber 
func (textEntityTypePhoneNumber *TextEntityTypePhoneNumber) MessageType() string {
 return "textEntityTypePhoneNumber" }

// InputThumbnail A thumbnail to be sent along with a file; should be in JPEG or WEBP format for stickers, and less than 200 kB in size  
type InputThumbnail struct {
tdCommon
Height int32 `json:"height"` // Thumbnail height, usually shouldn't exceed 90. Use 0 if unknown
Thumbnail InputFile `json:"thumbnail"` // Thumbnail file to send. Sending thumbnails by file_id is currently not supported
Width int32 `json:"width"` // Thumbnail width, usually shouldn't exceed 90. Use 0 if unknown 
}

// MessageType return the string telegram-type of InputThumbnail 
func (inputThumbnail *InputThumbnail) MessageType() string {
 return "inputThumbnail" }

// InputMessageText A text message  
type InputMessageText struct {
tdCommon
Text FormattedText `json:"text"` // Formatted text to be sent. Only Bold, Italic, Code, Pre, PreCode and TextUrl entities are allowed to be specified manually
DisableWebPagePreview bool `json:"disable_web_page_preview"` // True, if rich web page previews for URLs in the message text should be disabled 
ClearDraft bool `json:"clear_draft"` // True, if a chat message draft should be deleted
}

// MessageType return the string telegram-type of InputMessageText 
func (inputMessageText *InputMessageText) MessageType() string {
 return "inputMessageText" }

// InputMessageAnimation An animation message (GIF-style).  
type InputMessageAnimation struct {
tdCommon
Caption FormattedText `json:"caption"` // Animation caption; 0-200 characters
Animation InputFile `json:"animation"` // Animation file to be sent 
Thumbnail InputThumbnail `json:"thumbnail"` // Animation thumbnail, if available 
Duration int32 `json:"duration"` // Duration of the animation, in seconds 
Width int32 `json:"width"` // Width of the animation; may be replaced by the server 
Height int32 `json:"height"` // Height of the animation; may be replaced by the server 
}

// MessageType return the string telegram-type of InputMessageAnimation 
func (inputMessageAnimation *InputMessageAnimation) MessageType() string {
 return "inputMessageAnimation" }

// InputMessageAudio An audio message  
type InputMessageAudio struct {
tdCommon
Caption FormattedText `json:"caption"` // Audio caption; 0-200 characters
Audio InputFile `json:"audio"` // Audio file to be sent 
AlbumCoverThumbnail InputThumbnail `json:"album_cover_thumbnail"` // Thumbnail of the cover for the album, if available 
Duration int32 `json:"duration"` // Duration of the audio, in seconds; may be replaced by the server 
Title string `json:"title"` // Title of the audio; 0-64 characters; may be replaced by the server
Performer string `json:"performer"` // Performer of the audio; 0-64 characters, may be replaced by the server 
}

// MessageType return the string telegram-type of InputMessageAudio 
func (inputMessageAudio *InputMessageAudio) MessageType() string {
 return "inputMessageAudio" }

// InputMessageDocument A document message (general file)  
type InputMessageDocument struct {
tdCommon
Caption FormattedText `json:"caption"` // Document caption; 0-200 characters
Document InputFile `json:"document"` // Document to be sent 
Thumbnail InputThumbnail `json:"thumbnail"` // Document thumbnail, if available 
}

// MessageType return the string telegram-type of InputMessageDocument 
func (inputMessageDocument *InputMessageDocument) MessageType() string {
 return "inputMessageDocument" }

// InputMessagePhoto A photo message  
type InputMessagePhoto struct {
tdCommon
TTL int32 `json:"ttl"` // Photo TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
Photo InputFile `json:"photo"` // Photo to send 
Thumbnail InputThumbnail `json:"thumbnail"` // Photo thumbnail to be sent, this is sent to the other party in secret chats only 
AddedStickerFileIDs []int32 `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the photo, if applicable 
Width int32 `json:"width"` // Photo width 
Height int32 `json:"height"` // Photo height 
Caption FormattedText `json:"caption"` // Photo caption; 0-200 characters
}

// MessageType return the string telegram-type of InputMessagePhoto 
func (inputMessagePhoto *InputMessagePhoto) MessageType() string {
 return "inputMessagePhoto" }

// InputMessageSticker A sticker message  
type InputMessageSticker struct {
tdCommon
Sticker InputFile `json:"sticker"` // Sticker to be sent 
Thumbnail InputThumbnail `json:"thumbnail"` // Sticker thumbnail, if available 
Width int32 `json:"width"` // Sticker width 
Height int32 `json:"height"` // Sticker height
}

// MessageType return the string telegram-type of InputMessageSticker 
func (inputMessageSticker *InputMessageSticker) MessageType() string {
 return "inputMessageSticker" }

// InputMessageVideo A video message  
type InputMessageVideo struct {
tdCommon
Video InputFile `json:"video"` // Video to be sent 
Thumbnail InputThumbnail `json:"thumbnail"` // Video thumbnail, if available 
AddedStickerFileIDs []int32 `json:"added_sticker_file_ids"` // File identifiers of the stickers added to the video, if applicable
Duration int32 `json:"duration"` // Duration of the video, in seconds 
Width int32 `json:"width"` // Video width 
SupportsStreaming bool `json:"supports_streaming"` // True, if the video should be tried to be streamed
Height int32 `json:"height"` // Video height 
Caption FormattedText `json:"caption"` // Video caption; 0-200 characters 
TTL int32 `json:"ttl"` // Video TTL (Time To Live), in seconds (0-60). A non-zero TTL can be specified only in private chats
}

// MessageType return the string telegram-type of InputMessageVideo 
func (inputMessageVideo *InputMessageVideo) MessageType() string {
 return "inputMessageVideo" }

// InputMessageVideoNote A video note message  
type InputMessageVideoNote struct {
tdCommon
VideoNote InputFile `json:"video_note"` // Video note to be sent 
Thumbnail InputThumbnail `json:"thumbnail"` // Video thumbnail, if available 
Duration int32 `json:"duration"` // Duration of the video, in seconds 
Length int32 `json:"length"` // Video width and height; must be positive and not greater than 640
}

// MessageType return the string telegram-type of InputMessageVideoNote 
func (inputMessageVideoNote *InputMessageVideoNote) MessageType() string {
 return "inputMessageVideoNote" }

// InputMessageVoiceNote A voice note message  
type InputMessageVoiceNote struct {
tdCommon
Waveform []byte `json:"waveform"` // Waveform representation of the voice note, in 5-bit format 
Caption FormattedText `json:"caption"` // Voice note caption; 0-200 characters
VoiceNote InputFile `json:"voice_note"` // Voice note to be sent 
Duration int32 `json:"duration"` // Duration of the voice note, in seconds 
}

// MessageType return the string telegram-type of InputMessageVoiceNote 
func (inputMessageVoiceNote *InputMessageVoiceNote) MessageType() string {
 return "inputMessageVoiceNote" }

// InputMessageLocation A message with a location  
type InputMessageLocation struct {
tdCommon
Location Location `json:"location"` // Location to be sent 
LivePeriod int32 `json:"live_period"` // Period for which the location can be updated, in seconds; should bebetween 60 and 86400 for a live location and 0 otherwise
}

// MessageType return the string telegram-type of InputMessageLocation 
func (inputMessageLocation *InputMessageLocation) MessageType() string {
 return "inputMessageLocation" }

// InputMessageVenue A message with information about a venue  
type InputMessageVenue struct {
tdCommon
Venue Venue `json:"venue"` // Venue to send
}

// MessageType return the string telegram-type of InputMessageVenue 
func (inputMessageVenue *InputMessageVenue) MessageType() string {
 return "inputMessageVenue" }

// InputMessageContact A message containing a user contact  
type InputMessageContact struct {
tdCommon
Contact Contact `json:"contact"` // Contact to send
}

// MessageType return the string telegram-type of InputMessageContact 
func (inputMessageContact *InputMessageContact) MessageType() string {
 return "inputMessageContact" }

// InputMessageGame A message with a game; not supported for channels or secret chats  
type InputMessageGame struct {
tdCommon
BotUserID int32 `json:"bot_user_id"` // User identifier of the bot that owns the game 
GameShortName string `json:"game_short_name"` // Short name of the game
}

// MessageType return the string telegram-type of InputMessageGame 
func (inputMessageGame *InputMessageGame) MessageType() string {
 return "inputMessageGame" }

// InputMessageInvoice A message with an invoice; can be used only by bots and only in private chats  
type InputMessageInvoice struct {
tdCommon
PhotoSize int32 `json:"photo_size"` // Product photo size 
PhotoWidth int32 `json:"photo_width"` // Product photo width 
PhotoHeight int32 `json:"photo_height"` // Product photo height
Invoice Invoice `json:"invoice"` // Invoice 
Title string `json:"title"` // Product title; 1-32 characters 
Payload []byte `json:"payload"` // The invoice payload 
ProviderToken string `json:"provider_token"` // Payment provider token 
ProviderData string `json:"provider_data"` // JSON-encoded data about the invoice, which will be shared with the payment provider 
StartParameter string `json:"start_parameter"` // Unique invoice bot start_parameter for the generation of this invoice
Description string `json:"description"` // 
PhotoURL string `json:"photo_url"` // Product photo URL; optional 
}

// MessageType return the string telegram-type of InputMessageInvoice 
func (inputMessageInvoice *InputMessageInvoice) MessageType() string {
 return "inputMessageInvoice" }

// InputMessageForwarded A forwarded message  
type InputMessageForwarded struct {
tdCommon
MessageID int64 `json:"message_id"` // Identifier of the message to forward 
InGameShare bool `json:"in_game_share"` // True, if a game message should be shared within a launched game; applies only to game messages
FromChatID int64 `json:"from_chat_id"` // Identifier for the chat this forwarded message came from 
}

// MessageType return the string telegram-type of InputMessageForwarded 
func (inputMessageForwarded *InputMessageForwarded) MessageType() string {
 return "inputMessageForwarded" }

// SearchMessagesFilterEmpty Returns all found messages, no filter is applied 
type SearchMessagesFilterEmpty struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterEmpty 
func (searchMessagesFilterEmpty *SearchMessagesFilterEmpty) MessageType() string {
 return "searchMessagesFilterEmpty" }

// SearchMessagesFilterAnimation Returns only animation messages 
type SearchMessagesFilterAnimation struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterAnimation 
func (searchMessagesFilterAnimation *SearchMessagesFilterAnimation) MessageType() string {
 return "searchMessagesFilterAnimation" }

// SearchMessagesFilterAudio Returns only audio messages 
type SearchMessagesFilterAudio struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterAudio 
func (searchMessagesFilterAudio *SearchMessagesFilterAudio) MessageType() string {
 return "searchMessagesFilterAudio" }

// SearchMessagesFilterDocument Returns only document messages 
type SearchMessagesFilterDocument struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterDocument 
func (searchMessagesFilterDocument *SearchMessagesFilterDocument) MessageType() string {
 return "searchMessagesFilterDocument" }

// SearchMessagesFilterPhoto Returns only photo messages 
type SearchMessagesFilterPhoto struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterPhoto 
func (searchMessagesFilterPhoto *SearchMessagesFilterPhoto) MessageType() string {
 return "searchMessagesFilterPhoto" }

// SearchMessagesFilterVideo Returns only video messages 
type SearchMessagesFilterVideo struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterVideo 
func (searchMessagesFilterVideo *SearchMessagesFilterVideo) MessageType() string {
 return "searchMessagesFilterVideo" }

// SearchMessagesFilterVoiceNote Returns only voice note messages 
type SearchMessagesFilterVoiceNote struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterVoiceNote 
func (searchMessagesFilterVoiceNote *SearchMessagesFilterVoiceNote) MessageType() string {
 return "searchMessagesFilterVoiceNote" }

// SearchMessagesFilterPhotoAndVideo Returns only photo and video messages 
type SearchMessagesFilterPhotoAndVideo struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterPhotoAndVideo 
func (searchMessagesFilterPhotoAndVideo *SearchMessagesFilterPhotoAndVideo) MessageType() string {
 return "searchMessagesFilterPhotoAndVideo" }

// SearchMessagesFilterURL Returns only messages containing URLs 
type SearchMessagesFilterURL struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterURL 
func (searchMessagesFilterURL *SearchMessagesFilterURL) MessageType() string {
 return "searchMessagesFilterUrl" }

// SearchMessagesFilterChatPhoto Returns only messages containing chat photos 
type SearchMessagesFilterChatPhoto struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterChatPhoto 
func (searchMessagesFilterChatPhoto *SearchMessagesFilterChatPhoto) MessageType() string {
 return "searchMessagesFilterChatPhoto" }

// SearchMessagesFilterCall Returns only call messages 
type SearchMessagesFilterCall struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterCall 
func (searchMessagesFilterCall *SearchMessagesFilterCall) MessageType() string {
 return "searchMessagesFilterCall" }

// SearchMessagesFilterMissedCall Returns only incoming call messages with missed/declined discard reasons 
type SearchMessagesFilterMissedCall struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterMissedCall 
func (searchMessagesFilterMissedCall *SearchMessagesFilterMissedCall) MessageType() string {
 return "searchMessagesFilterMissedCall" }

// SearchMessagesFilterVideoNote Returns only video note messages 
type SearchMessagesFilterVideoNote struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterVideoNote 
func (searchMessagesFilterVideoNote *SearchMessagesFilterVideoNote) MessageType() string {
 return "searchMessagesFilterVideoNote" }

// SearchMessagesFilterVoiceAndVideoNote Returns only voice and video note messages 
type SearchMessagesFilterVoiceAndVideoNote struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterVoiceAndVideoNote 
func (searchMessagesFilterVoiceAndVideoNote *SearchMessagesFilterVoiceAndVideoNote) MessageType() string {
 return "searchMessagesFilterVoiceAndVideoNote" }

// SearchMessagesFilterMention Returns only messages with mentions of the current user, or messages that are replies to their messages 
type SearchMessagesFilterMention struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterMention 
func (searchMessagesFilterMention *SearchMessagesFilterMention) MessageType() string {
 return "searchMessagesFilterMention" }

// SearchMessagesFilterUnreadMention Returns only messages with unread mentions of the current user or messages that are replies to their messages. When using this filter the results can't be additionally filtered by a query or by the sending user 
type SearchMessagesFilterUnreadMention struct {
tdCommon

}

// MessageType return the string telegram-type of SearchMessagesFilterUnreadMention 
func (searchMessagesFilterUnreadMention *SearchMessagesFilterUnreadMention) MessageType() string {
 return "searchMessagesFilterUnreadMention" }

// ChatActionTyping The user is typing a message 
type ChatActionTyping struct {
tdCommon

}

// MessageType return the string telegram-type of ChatActionTyping 
func (chatActionTyping *ChatActionTyping) MessageType() string {
 return "chatActionTyping" }

// ChatActionRecordingVideo The user is recording a video 
type ChatActionRecordingVideo struct {
tdCommon

}

// MessageType return the string telegram-type of ChatActionRecordingVideo 
func (chatActionRecordingVideo *ChatActionRecordingVideo) MessageType() string {
 return "chatActionRecordingVideo" }

// ChatActionUploadingVideo The user is uploading a video  
type ChatActionUploadingVideo struct {
tdCommon
Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingVideo 
func (chatActionUploadingVideo *ChatActionUploadingVideo) MessageType() string {
 return "chatActionUploadingVideo" }

// ChatActionRecordingVoiceNote The user is recording a voice note 
type ChatActionRecordingVoiceNote struct {
tdCommon

}

// MessageType return the string telegram-type of ChatActionRecordingVoiceNote 
func (chatActionRecordingVoiceNote *ChatActionRecordingVoiceNote) MessageType() string {
 return "chatActionRecordingVoiceNote" }

// ChatActionUploadingVoiceNote The user is uploading a voice note  
type ChatActionUploadingVoiceNote struct {
tdCommon
Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingVoiceNote 
func (chatActionUploadingVoiceNote *ChatActionUploadingVoiceNote) MessageType() string {
 return "chatActionUploadingVoiceNote" }

// ChatActionUploadingPhoto The user is uploading a photo  
type ChatActionUploadingPhoto struct {
tdCommon
Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingPhoto 
func (chatActionUploadingPhoto *ChatActionUploadingPhoto) MessageType() string {
 return "chatActionUploadingPhoto" }

// ChatActionUploadingDocument The user is uploading a document  
type ChatActionUploadingDocument struct {
tdCommon
Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingDocument 
func (chatActionUploadingDocument *ChatActionUploadingDocument) MessageType() string {
 return "chatActionUploadingDocument" }

// ChatActionChoosingLocation The user is picking a location or venue to send 
type ChatActionChoosingLocation struct {
tdCommon

}

// MessageType return the string telegram-type of ChatActionChoosingLocation 
func (chatActionChoosingLocation *ChatActionChoosingLocation) MessageType() string {
 return "chatActionChoosingLocation" }

// ChatActionChoosingContact The user is picking a contact to send 
type ChatActionChoosingContact struct {
tdCommon

}

// MessageType return the string telegram-type of ChatActionChoosingContact 
func (chatActionChoosingContact *ChatActionChoosingContact) MessageType() string {
 return "chatActionChoosingContact" }

// ChatActionStartPlayingGame The user has started to play a game 
type ChatActionStartPlayingGame struct {
tdCommon

}

// MessageType return the string telegram-type of ChatActionStartPlayingGame 
func (chatActionStartPlayingGame *ChatActionStartPlayingGame) MessageType() string {
 return "chatActionStartPlayingGame" }

// ChatActionRecordingVideoNote The user is recording a video note 
type ChatActionRecordingVideoNote struct {
tdCommon

}

// MessageType return the string telegram-type of ChatActionRecordingVideoNote 
func (chatActionRecordingVideoNote *ChatActionRecordingVideoNote) MessageType() string {
 return "chatActionRecordingVideoNote" }

// ChatActionUploadingVideoNote The user is uploading a video note  
type ChatActionUploadingVideoNote struct {
tdCommon
Progress int32 `json:"progress"` // Upload progress, as a percentage
}

// MessageType return the string telegram-type of ChatActionUploadingVideoNote 
func (chatActionUploadingVideoNote *ChatActionUploadingVideoNote) MessageType() string {
 return "chatActionUploadingVideoNote" }

// ChatActionCancel The user has cancelled the previous action 
type ChatActionCancel struct {
tdCommon

}

// MessageType return the string telegram-type of ChatActionCancel 
func (chatActionCancel *ChatActionCancel) MessageType() string {
 return "chatActionCancel" }

// UserStatusEmpty The user status was never changed 
type UserStatusEmpty struct {
tdCommon

}

// MessageType return the string telegram-type of UserStatusEmpty 
func (userStatusEmpty *UserStatusEmpty) MessageType() string {
 return "userStatusEmpty" }

// UserStatusOnline The user is online  
type UserStatusOnline struct {
tdCommon
Expires int32 `json:"expires"` // Point in time (Unix timestamp) when the user's online status will expire
}

// MessageType return the string telegram-type of UserStatusOnline 
func (userStatusOnline *UserStatusOnline) MessageType() string {
 return "userStatusOnline" }

// UserStatusOffline The user is offline  
type UserStatusOffline struct {
tdCommon
WasOnline int32 `json:"was_online"` // Point in time (Unix timestamp) when the user was last online
}

// MessageType return the string telegram-type of UserStatusOffline 
func (userStatusOffline *UserStatusOffline) MessageType() string {
 return "userStatusOffline" }

// UserStatusRecently The user was online recently 
type UserStatusRecently struct {
tdCommon

}

// MessageType return the string telegram-type of UserStatusRecently 
func (userStatusRecently *UserStatusRecently) MessageType() string {
 return "userStatusRecently" }

// UserStatusLastWeek The user is offline, but was online last week 
type UserStatusLastWeek struct {
tdCommon

}

// MessageType return the string telegram-type of UserStatusLastWeek 
func (userStatusLastWeek *UserStatusLastWeek) MessageType() string {
 return "userStatusLastWeek" }

// UserStatusLastMonth The user is offline, but was online last month 
type UserStatusLastMonth struct {
tdCommon

}

// MessageType return the string telegram-type of UserStatusLastMonth 
func (userStatusLastMonth *UserStatusLastMonth) MessageType() string {
 return "userStatusLastMonth" }

// Stickers Represents a list of stickers  
type Stickers struct {
tdCommon
Stickers []Sticker `json:"stickers"` // List of stickers
}

// MessageType return the string telegram-type of Stickers 
func (stickers *Stickers) MessageType() string {
 return "stickers" }

// StickerEmojis Represents a list of all emoji corresponding to a sticker in a sticker set. The list is only for informational purposes, because a sticker is always sent with a fixed emoji from the corresponding Sticker object  
type StickerEmojis struct {
tdCommon
Emojis []string `json:"emojis"` // List of emojis
}

// MessageType return the string telegram-type of StickerEmojis 
func (stickerEmojis *StickerEmojis) MessageType() string {
 return "stickerEmojis" }

// StickerSet Represents a sticker set  
type StickerSet struct {
tdCommon
IsViewed bool `json:"is_viewed"` // True for already viewed trending sticker sets 
Emojis []StickerEmojis `json:"emojis"` // A list of emoji corresponding to the stickers in the same order
ID int64 `json:"id"` // Identifier of the sticker set 
Title string `json:"title"` // Title of the sticker set 
Name string `json:"name"` // Name of the sticker set 
IsOfficial bool `json:"is_official"` // True, if the sticker set is official 
IsMasks bool `json:"is_masks"` // True, if the stickers in the set are masks
IsInstalled bool `json:"is_installed"` // True, if the sticker set has been installed by the current user
IsArchived bool `json:"is_archived"` // True, if the sticker set has been archived. A sticker set can't be installed and archived simultaneously 
Stickers []Sticker `json:"stickers"` // List of stickers in this set 
}

// MessageType return the string telegram-type of StickerSet 
func (stickerSet *StickerSet) MessageType() string {
 return "stickerSet" }

// StickerSetInfo Represents short information about a sticker set  
type StickerSetInfo struct {
tdCommon
ID int64 `json:"id"` // Identifier of the sticker set 
Title string `json:"title"` // Title of the sticker set 
Name string `json:"name"` // Name of the sticker set 
IsArchived bool `json:"is_archived"` // True, if the sticker set has been archived. A sticker set can't be installed and archived simultaneously 
IsOfficial bool `json:"is_official"` // True, if the sticker set is official 
IsMasks bool `json:"is_masks"` // True, if the stickers in the set are masks
IsViewed bool `json:"is_viewed"` // True for already viewed trending sticker sets 
Covers []Sticker `json:"covers"` // Contains up to the first 5 stickers from the set, depending on the context. If the client needs more stickers the full set should be requested
IsInstalled bool `json:"is_installed"` // True, if the sticker set has been installed by current user
Size int32 `json:"size"` // Total number of stickers in the set 
}

// MessageType return the string telegram-type of StickerSetInfo 
func (stickerSetInfo *StickerSetInfo) MessageType() string {
 return "stickerSetInfo" }

// StickerSets Represents a list of sticker sets  
type StickerSets struct {
tdCommon
TotalCount int32 `json:"total_count"` // Approximate total number of sticker sets found 
Sets []StickerSetInfo `json:"sets"` // List of sticker sets
}

// MessageType return the string telegram-type of StickerSets 
func (stickerSets *StickerSets) MessageType() string {
 return "stickerSets" }

// CallDiscardReasonEmpty The call wasn't discarded, or the reason is unknown 
type CallDiscardReasonEmpty struct {
tdCommon

}

// MessageType return the string telegram-type of CallDiscardReasonEmpty 
func (callDiscardReasonEmpty *CallDiscardReasonEmpty) MessageType() string {
 return "callDiscardReasonEmpty" }

// CallDiscardReasonMissed The call was ended before the conversation started. It was cancelled by the caller or missed by the other party 
type CallDiscardReasonMissed struct {
tdCommon

}

// MessageType return the string telegram-type of CallDiscardReasonMissed 
func (callDiscardReasonMissed *CallDiscardReasonMissed) MessageType() string {
 return "callDiscardReasonMissed" }

// CallDiscardReasonDeclined The call was ended before the conversation started. It was declined by the other party 
type CallDiscardReasonDeclined struct {
tdCommon

}

// MessageType return the string telegram-type of CallDiscardReasonDeclined 
func (callDiscardReasonDeclined *CallDiscardReasonDeclined) MessageType() string {
 return "callDiscardReasonDeclined" }

// CallDiscardReasonDisconnected The call was ended during the conversation because the users were disconnected 
type CallDiscardReasonDisconnected struct {
tdCommon

}

// MessageType return the string telegram-type of CallDiscardReasonDisconnected 
func (callDiscardReasonDisconnected *CallDiscardReasonDisconnected) MessageType() string {
 return "callDiscardReasonDisconnected" }

// CallDiscardReasonHungUp The call was ended because one of the parties hung up 
type CallDiscardReasonHungUp struct {
tdCommon

}

// MessageType return the string telegram-type of CallDiscardReasonHungUp 
func (callDiscardReasonHungUp *CallDiscardReasonHungUp) MessageType() string {
 return "callDiscardReasonHungUp" }

// CallProtocol Specifies the supported call protocols  
type CallProtocol struct {
tdCommon
UDPP2p bool `json:"udp_p2p"` // True, if UDP peer-to-peer connections are supported 
UDPReflector bool `json:"udp_reflector"` // True, if connection through UDP reflectors is supported 
MinLayer int32 `json:"min_layer"` // Minimum supported API layer; use 65 
MaxLayer int32 `json:"max_layer"` // Maximum supported API layer; use 65
}

// MessageType return the string telegram-type of CallProtocol 
func (callProtocol *CallProtocol) MessageType() string {
 return "callProtocol" }

// CallConnection Describes the address of UDP reflectors  
type CallConnection struct {
tdCommon
Port int32 `json:"port"` // Reflector port number 
PeerTag []byte `json:"peer_tag"` // Connection peer tag
ID int64 `json:"id"` // Reflector identifier 
IP string `json:"ip"` // IPv4 reflector address 
IPv6 string `json:"ipv6"` // IPv6 reflector address 
}

// MessageType return the string telegram-type of CallConnection 
func (callConnection *CallConnection) MessageType() string {
 return "callConnection" }

// CallID Contains the call identifier  
type CallID struct {
tdCommon
ID int32 `json:"id"` // Call identifier
}

// MessageType return the string telegram-type of CallID 
func (callID *CallID) MessageType() string {
 return "callId" }

// CallStatePending The call is pending, waiting to be accepted by a user  
type CallStatePending struct {
tdCommon
IsReceived bool `json:"is_received"` // True, if the call has already been received by the other party
IsCreated bool `json:"is_created"` // True, if the call has already been created by the server 
}

// MessageType return the string telegram-type of CallStatePending 
func (callStatePending *CallStatePending) MessageType() string {
 return "callStatePending" }

// CallStateExchangingKeys The call has been answered and encryption keys are being exchanged 
type CallStateExchangingKeys struct {
tdCommon

}

// MessageType return the string telegram-type of CallStateExchangingKeys 
func (callStateExchangingKeys *CallStateExchangingKeys) MessageType() string {
 return "callStateExchangingKeys" }

// CallStateReady The call is ready to use  
type CallStateReady struct {
tdCommon
Protocol CallProtocol `json:"protocol"` // Call protocols supported by the peer 
Connections []CallConnection `json:"connections"` // Available UDP reflectors 
Config string `json:"config"` // A JSON-encoded call config 
EncryptionKey []byte `json:"encryption_key"` // Call encryption key 
Emojis []string `json:"emojis"` // Encryption key emojis fingerprint
}

// MessageType return the string telegram-type of CallStateReady 
func (callStateReady *CallStateReady) MessageType() string {
 return "callStateReady" }

// CallStateHangingUp The call is hanging up after discardCall has been called 
type CallStateHangingUp struct {
tdCommon

}

// MessageType return the string telegram-type of CallStateHangingUp 
func (callStateHangingUp *CallStateHangingUp) MessageType() string {
 return "callStateHangingUp" }

// CallStateDiscarded The call has ended successfully  
type CallStateDiscarded struct {
tdCommon
Reason CallDiscardReason `json:"reason"` // The reason, why the call has ended 
NeedRating bool `json:"need_rating"` // True, if the call rating should be sent to the server 
NeedDebugInformation bool `json:"need_debug_information"` // True, if the call debug information should be sent to the server
}

// MessageType return the string telegram-type of CallStateDiscarded 
func (callStateDiscarded *CallStateDiscarded) MessageType() string {
 return "callStateDiscarded" }

// CallStateError The call has ended with an error  
type CallStateError struct {
tdCommon
Error Error `json:"error"` // Error. An error with the code 4005000 will be returned if an outgoing call is missed because of an expired timeout
}

// MessageType return the string telegram-type of CallStateError 
func (callStateError *CallStateError) MessageType() string {
 return "callStateError" }

// Call Describes a call  
type Call struct {
tdCommon
ID int32 `json:"id"` // Call identifier, not persistent 
UserID int32 `json:"user_id"` // Peer user identifier 
IsOutgoing bool `json:"is_outgoing"` // True, if the call is outgoing 
State CallState `json:"state"` // Call state
}

// MessageType return the string telegram-type of Call 
func (call *Call) MessageType() string {
 return "call" }

// Animations Represents a list of animations  
type Animations struct {
tdCommon
Animations []Animation `json:"animations"` // List of animations
}

// MessageType return the string telegram-type of Animations 
func (animations *Animations) MessageType() string {
 return "animations" }

// ImportedContacts Represents the result of an ImportContacts request  
type ImportedContacts struct {
tdCommon
UserIDs []int32 `json:"user_ids"` // User identifiers of the imported contacts in the same order as they were specified in the request; 0 if the contact is not yet a registered user
ImporterCount []int32 `json:"importer_count"` // The number of users that imported the corresponding contact; 0 for already registered users or if unavailable
}

// MessageType return the string telegram-type of ImportedContacts 
func (importedContacts *ImportedContacts) MessageType() string {
 return "importedContacts" }

// InputInlineQueryResultAnimatedGif Represents a link to an animated GIF  
type InputInlineQueryResultAnimatedGif struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
Title string `json:"title"` // Title of the query result 
ThumbnailURL string `json:"thumbnail_url"` // URL of the static result thumbnail (JPEG or GIF), if it exists
GifDuration int32 `json:"gif_duration"` // Duration of the GIF, in seconds 
GifWidth int32 `json:"gif_width"` // Width of the GIF 
GifURL string `json:"gif_url"` // The URL of the GIF-file (file size must not exceed 1MB) 
GifHeight int32 `json:"gif_height"` // Height of the GIF
ReplyMarkup ReplyMarkup `json:"reply_markup"` // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageAnimation, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultAnimatedGif 
func (inputInlineQueryResultAnimatedGif *InputInlineQueryResultAnimatedGif) MessageType() string {
 return "inputInlineQueryResultAnimatedGif" }

// InputInlineQueryResultAnimatedMpeg4 Represents a link to an animated (i.e. without sound) H.264/MPEG-4 AVC video  
type InputInlineQueryResultAnimatedMpeg4 struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
Mpeg4URL string `json:"mpeg4_url"` // The URL of the MPEG4-file (file size must not exceed 1MB) 
Mpeg4Duration int32 `json:"mpeg4_duration"` // Duration of the video, in seconds 
Mpeg4Width int32 `json:"mpeg4_width"` // Width of the video 
Title string `json:"title"` // Title of the result 
ThumbnailURL string `json:"thumbnail_url"` // URL of the static result thumbnail (JPEG or GIF), if it exists
Mpeg4Height int32 `json:"mpeg4_height"` // Height of the video
ReplyMarkup ReplyMarkup `json:"reply_markup"` // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageAnimation, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultAnimatedMpeg4 
func (inputInlineQueryResultAnimatedMpeg4 *InputInlineQueryResultAnimatedMpeg4) MessageType() string {
 return "inputInlineQueryResultAnimatedMpeg4" }

// InputInlineQueryResultArticle Represents a link to an article or web page  
type InputInlineQueryResultArticle struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
HideURL bool `json:"hide_url"` // True, if the URL must be not shown 
Title string `json:"title"` // Title of the result
Description string `json:"description"` // 
ThumbnailHeight int32 `json:"thumbnail_height"` // Thumbnail height, if known
InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageLocation, InputMessageVenue or InputMessageContact
URL string `json:"url"` // URL of the result, if it exists 
ThumbnailURL string `json:"thumbnail_url"` // URL of the result thumbnail, if it exists 
ThumbnailWidth int32 `json:"thumbnail_width"` // Thumbnail width, if known 
ReplyMarkup ReplyMarkup `json:"reply_markup"` // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
}

// MessageType return the string telegram-type of InputInlineQueryResultArticle 
func (inputInlineQueryResultArticle *InputInlineQueryResultArticle) MessageType() string {
 return "inputInlineQueryResultArticle" }

// InputInlineQueryResultAudio Represents a link to an MP3 audio file  
type InputInlineQueryResultAudio struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
Title string `json:"title"` // Title of the audio file 
Performer string `json:"performer"` // Performer of the audio file
AudioURL string `json:"audio_url"` // The URL of the audio file 
AudioDuration int32 `json:"audio_duration"` // Audio file duration, in seconds
ReplyMarkup ReplyMarkup `json:"reply_markup"` // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageAudio, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultAudio 
func (inputInlineQueryResultAudio *InputInlineQueryResultAudio) MessageType() string {
 return "inputInlineQueryResultAudio" }

// InputInlineQueryResultContact Represents a user contact  
type InputInlineQueryResultContact struct {
tdCommon
ThumbnailWidth int32 `json:"thumbnail_width"` // Thumbnail width, if known 
ThumbnailHeight int32 `json:"thumbnail_height"` // Thumbnail height, if known
ReplyMarkup ReplyMarkup `json:"reply_markup"` // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageLocation, InputMessageVenue or InputMessageContact
ID string `json:"id"` // Unique identifier of the query result 
Contact Contact `json:"contact"` // User contact 
ThumbnailURL string `json:"thumbnail_url"` // URL of the result thumbnail, if it exists 
}

// MessageType return the string telegram-type of InputInlineQueryResultContact 
func (inputInlineQueryResultContact *InputInlineQueryResultContact) MessageType() string {
 return "inputInlineQueryResultContact" }

// InputInlineQueryResultDocument Represents a link to a file  
type InputInlineQueryResultDocument struct {
tdCommon
InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageDocument, InputMessageLocation, InputMessageVenue or InputMessageContact
ID string `json:"id"` // Unique identifier of the query result 
Title string `json:"title"` // Title of the resulting file 
MimeType string `json:"mime_type"` // MIME type of the file content; only "application/pdf" and "application/zip" are currently allowed
ReplyMarkup ReplyMarkup `json:"reply_markup"` // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
ThumbnailHeight int32 `json:"thumbnail_height"` // Height of the thumbnail
Description string `json:"description"` // 
DocumentURL string `json:"document_url"` // URL of the file 
ThumbnailURL string `json:"thumbnail_url"` // The URL of the file thumbnail, if it exists 
ThumbnailWidth int32 `json:"thumbnail_width"` // Width of the thumbnail 
}

// MessageType return the string telegram-type of InputInlineQueryResultDocument 
func (inputInlineQueryResultDocument *InputInlineQueryResultDocument) MessageType() string {
 return "inputInlineQueryResultDocument" }

// InputInlineQueryResultGame Represents a game  
type InputInlineQueryResultGame struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
GameShortName string `json:"game_short_name"` // Short name of the game 
ReplyMarkup ReplyMarkup `json:"reply_markup"` // Message reply markup. Must be of type replyMarkupInlineKeyboard or null
}

// MessageType return the string telegram-type of InputInlineQueryResultGame 
func (inputInlineQueryResultGame *InputInlineQueryResultGame) MessageType() string {
 return "inputInlineQueryResultGame" }

// InputInlineQueryResultLocation Represents a point on the map  
type InputInlineQueryResultLocation struct {
tdCommon
LivePeriod int32 `json:"live_period"` // Amount of time relative to the message sent time until the location can be updated, in seconds 
ThumbnailURL string `json:"thumbnail_url"` // URL of the result thumbnail, if it exists 
ReplyMarkup ReplyMarkup `json:"reply_markup"` // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageLocation, InputMessageVenue or InputMessageContact
ID string `json:"id"` // Unique identifier of the query result 
Location Location `json:"location"` // Location result 
Title string `json:"title"` // Title of the result 
ThumbnailWidth int32 `json:"thumbnail_width"` // Thumbnail width, if known 
ThumbnailHeight int32 `json:"thumbnail_height"` // Thumbnail height, if known
}

// MessageType return the string telegram-type of InputInlineQueryResultLocation 
func (inputInlineQueryResultLocation *InputInlineQueryResultLocation) MessageType() string {
 return "inputInlineQueryResultLocation" }

// InputInlineQueryResultPhoto Represents link to a JPEG image  
type InputInlineQueryResultPhoto struct {
tdCommon
Title string `json:"title"` // Title of the result, if known 
Description string `json:"description"` // 
PhotoWidth int32 `json:"photo_width"` // Width of the photo 
ReplyMarkup ReplyMarkup `json:"reply_markup"` // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessagePhoto, InputMessageLocation, InputMessageVenue or InputMessageContact
ID string `json:"id"` // Unique identifier of the query result 
PhotoURL string `json:"photo_url"` // The URL of the JPEG photo (photo size must not exceed 5MB) 
PhotoHeight int32 `json:"photo_height"` // Height of the photo
ThumbnailURL string `json:"thumbnail_url"` // URL of the photo thumbnail, if it exists
}

// MessageType return the string telegram-type of InputInlineQueryResultPhoto 
func (inputInlineQueryResultPhoto *InputInlineQueryResultPhoto) MessageType() string {
 return "inputInlineQueryResultPhoto" }

// InputInlineQueryResultSticker Represents a link to a WEBP sticker  
type InputInlineQueryResultSticker struct {
tdCommon
ReplyMarkup ReplyMarkup `json:"reply_markup"` // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, inputMessageSticker, InputMessageLocation, InputMessageVenue or InputMessageContact
ID string `json:"id"` // Unique identifier of the query result 
ThumbnailURL string `json:"thumbnail_url"` // URL of the sticker thumbnail, if it exists
StickerURL string `json:"sticker_url"` // The URL of the WEBP sticker (sticker file size must not exceed 5MB) 
StickerWidth int32 `json:"sticker_width"` // Width of the sticker 
StickerHeight int32 `json:"sticker_height"` // Height of the sticker
}

// MessageType return the string telegram-type of InputInlineQueryResultSticker 
func (inputInlineQueryResultSticker *InputInlineQueryResultSticker) MessageType() string {
 return "inputInlineQueryResultSticker" }

// InputInlineQueryResultVenue Represents information about a venue  
type InputInlineQueryResultVenue struct {
tdCommon
ThumbnailHeight int32 `json:"thumbnail_height"` // Thumbnail height, if known
ReplyMarkup ReplyMarkup `json:"reply_markup"` // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageLocation, InputMessageVenue or InputMessageContact
ID string `json:"id"` // Unique identifier of the query result 
Venue Venue `json:"venue"` // Venue result 
ThumbnailURL string `json:"thumbnail_url"` // URL of the result thumbnail, if it exists 
ThumbnailWidth int32 `json:"thumbnail_width"` // Thumbnail width, if known 
}

// MessageType return the string telegram-type of InputInlineQueryResultVenue 
func (inputInlineQueryResultVenue *InputInlineQueryResultVenue) MessageType() string {
 return "inputInlineQueryResultVenue" }

// InputInlineQueryResultVideo Represents a link to a page containing an embedded video player or a video file  
type InputInlineQueryResultVideo struct {
tdCommon
InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageVideo, InputMessageLocation, InputMessageVenue or InputMessageContact
Title string `json:"title"` // Title of the result 
ThumbnailURL string `json:"thumbnail_url"` // The URL of the video thumbnail (JPEG), if it exists 
MimeType string `json:"mime_type"` // MIME type of the content of the video URL, only "text/html" or "video/mp4" are currently supported
VideoWidth int32 `json:"video_width"` // Width of the video 
VideoHeight int32 `json:"video_height"` // Height of the video 
ReplyMarkup ReplyMarkup `json:"reply_markup"` // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
ID string `json:"id"` // Unique identifier of the query result 
Description string `json:"description"` // 
VideoURL string `json:"video_url"` // URL of the embedded video player or video file 
VideoDuration int32 `json:"video_duration"` // Video duration, in seconds
}

// MessageType return the string telegram-type of InputInlineQueryResultVideo 
func (inputInlineQueryResultVideo *InputInlineQueryResultVideo) MessageType() string {
 return "inputInlineQueryResultVideo" }

// InputInlineQueryResultVoiceNote Represents a link to an opus-encoded audio file within an OGG container, single channel audio  
type InputInlineQueryResultVoiceNote struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
Title string `json:"title"` // Title of the voice note
VoiceNoteURL string `json:"voice_note_url"` // The URL of the voice note file 
VoiceNoteDuration int32 `json:"voice_note_duration"` // Duration of the voice note, in seconds
ReplyMarkup ReplyMarkup `json:"reply_markup"` // The message reply markup. Must be of type replyMarkupInlineKeyboard or null
InputMessageContent InputMessageContent `json:"input_message_content"` // The content of the message to be sent. Must be one of the following types: InputMessageText, InputMessageVoiceNote, InputMessageLocation, InputMessageVenue or InputMessageContact
}

// MessageType return the string telegram-type of InputInlineQueryResultVoiceNote 
func (inputInlineQueryResultVoiceNote *InputInlineQueryResultVoiceNote) MessageType() string {
 return "inputInlineQueryResultVoiceNote" }

// InlineQueryResultArticle Represents a link to an article or web page  
type InlineQueryResultArticle struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
URL string `json:"url"` // URL of the result, if it exists 
HideURL bool `json:"hide_url"` // True, if the URL must be not shown 
Title string `json:"title"` // Title of the result
Description string `json:"description"` // 
Thumbnail PhotoSize `json:"thumbnail"` // Result thumbnail; may be null
}

// MessageType return the string telegram-type of InlineQueryResultArticle 
func (inlineQueryResultArticle *InlineQueryResultArticle) MessageType() string {
 return "inlineQueryResultArticle" }

// InlineQueryResultContact Represents a user contact  
type InlineQueryResultContact struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
Contact Contact `json:"contact"` // A user contact 
Thumbnail PhotoSize `json:"thumbnail"` // Result thumbnail; may be null
}

// MessageType return the string telegram-type of InlineQueryResultContact 
func (inlineQueryResultContact *InlineQueryResultContact) MessageType() string {
 return "inlineQueryResultContact" }

// InlineQueryResultLocation Represents a point on the map  
type InlineQueryResultLocation struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
Location Location `json:"location"` // Location result 
Title string `json:"title"` // Title of the result 
Thumbnail PhotoSize `json:"thumbnail"` // Result thumbnail; may be null
}

// MessageType return the string telegram-type of InlineQueryResultLocation 
func (inlineQueryResultLocation *InlineQueryResultLocation) MessageType() string {
 return "inlineQueryResultLocation" }

// InlineQueryResultVenue Represents information about a venue  
type InlineQueryResultVenue struct {
tdCommon
Thumbnail PhotoSize `json:"thumbnail"` // Result thumbnail; may be null
ID string `json:"id"` // Unique identifier of the query result 
Venue Venue `json:"venue"` // Venue result 
}

// MessageType return the string telegram-type of InlineQueryResultVenue 
func (inlineQueryResultVenue *InlineQueryResultVenue) MessageType() string {
 return "inlineQueryResultVenue" }

// InlineQueryResultGame Represents information about a game  
type InlineQueryResultGame struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
Game Game `json:"game"` // Game result
}

// MessageType return the string telegram-type of InlineQueryResultGame 
func (inlineQueryResultGame *InlineQueryResultGame) MessageType() string {
 return "inlineQueryResultGame" }

// InlineQueryResultAnimation Represents an animation file  
type InlineQueryResultAnimation struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
Animation Animation `json:"animation"` // Animation file 
Title string `json:"title"` // Animation title
}

// MessageType return the string telegram-type of InlineQueryResultAnimation 
func (inlineQueryResultAnimation *InlineQueryResultAnimation) MessageType() string {
 return "inlineQueryResultAnimation" }

// InlineQueryResultAudio Represents an audio file  
type InlineQueryResultAudio struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
Audio Audio `json:"audio"` // Audio file
}

// MessageType return the string telegram-type of InlineQueryResultAudio 
func (inlineQueryResultAudio *InlineQueryResultAudio) MessageType() string {
 return "inlineQueryResultAudio" }

// InlineQueryResultDocument Represents a document  
type InlineQueryResultDocument struct {
tdCommon
Document Document `json:"document"` // Document 
Title string `json:"title"` // Document title 
Description string `json:"description"` // 
ID string `json:"id"` // Unique identifier of the query result 
}

// MessageType return the string telegram-type of InlineQueryResultDocument 
func (inlineQueryResultDocument *InlineQueryResultDocument) MessageType() string {
 return "inlineQueryResultDocument" }

// InlineQueryResultPhoto Represents a photo  
type InlineQueryResultPhoto struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
Photo Photo `json:"photo"` // Photo 
Title string `json:"title"` // Title of the result, if known 
Description string `json:"description"` // 
}

// MessageType return the string telegram-type of InlineQueryResultPhoto 
func (inlineQueryResultPhoto *InlineQueryResultPhoto) MessageType() string {
 return "inlineQueryResultPhoto" }

// InlineQueryResultSticker Represents a sticker  
type InlineQueryResultSticker struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
Sticker Sticker `json:"sticker"` // Sticker
}

// MessageType return the string telegram-type of InlineQueryResultSticker 
func (inlineQueryResultSticker *InlineQueryResultSticker) MessageType() string {
 return "inlineQueryResultSticker" }

// InlineQueryResultVideo Represents a video  
type InlineQueryResultVideo struct {
tdCommon
Description string `json:"description"` // 
ID string `json:"id"` // Unique identifier of the query result 
Video Video `json:"video"` // Video 
Title string `json:"title"` // Title of the video 
}

// MessageType return the string telegram-type of InlineQueryResultVideo 
func (inlineQueryResultVideo *InlineQueryResultVideo) MessageType() string {
 return "inlineQueryResultVideo" }

// InlineQueryResultVoiceNote Represents a voice note  
type InlineQueryResultVoiceNote struct {
tdCommon
ID string `json:"id"` // Unique identifier of the query result 
VoiceNote VoiceNote `json:"voice_note"` // Voice note 
Title string `json:"title"` // Title of the voice note
}

// MessageType return the string telegram-type of InlineQueryResultVoiceNote 
func (inlineQueryResultVoiceNote *InlineQueryResultVoiceNote) MessageType() string {
 return "inlineQueryResultVoiceNote" }

// InlineQueryResults Represents the results of the inline query. Use sendInlineQueryResultMessage to send the result of the query  
type InlineQueryResults struct {
tdCommon
InlineQueryID int64 `json:"inline_query_id"` // Unique identifier of the inline query 
NextOffset string `json:"next_offset"` // The offset for the next request. If empty, there are no more results 
Results []InlineQueryResult `json:"results"` // Results of the query
SwitchPmText string `json:"switch_pm_text"` // If non-empty, this text should be shown on the button, which opens a private chat with the bot and sends the bot a start message with the switch_pm_parameter 
SwitchPmParameter string `json:"switch_pm_parameter"` // Parameter for the bot start message
}

// MessageType return the string telegram-type of InlineQueryResults 
func (inlineQueryResults *InlineQueryResults) MessageType() string {
 return "inlineQueryResults" }

// CallbackQueryPayloadData The payload from a general callback button  
type CallbackQueryPayloadData struct {
tdCommon
Data []byte `json:"data"` // Data that was attached to the callback button
}

// MessageType return the string telegram-type of CallbackQueryPayloadData 
func (callbackQueryPayloadData *CallbackQueryPayloadData) MessageType() string {
 return "callbackQueryPayloadData" }

// CallbackQueryPayloadGame The payload from a game callback button  
type CallbackQueryPayloadGame struct {
tdCommon
GameShortName string `json:"game_short_name"` // A short name of the game that was attached to the callback button
}

// MessageType return the string telegram-type of CallbackQueryPayloadGame 
func (callbackQueryPayloadGame *CallbackQueryPayloadGame) MessageType() string {
 return "callbackQueryPayloadGame" }

// CallbackQueryAnswer Contains a bot's answer to a callback query  
type CallbackQueryAnswer struct {
tdCommon
Text string `json:"text"` // Text of the answer 
ShowAlert bool `json:"show_alert"` // True, if an alert should be shown to the user instead of a toast notification 
URL string `json:"url"` // URL to be opened
}

// MessageType return the string telegram-type of CallbackQueryAnswer 
func (callbackQueryAnswer *CallbackQueryAnswer) MessageType() string {
 return "callbackQueryAnswer" }

// CustomRequestResult Contains the result of a custom request  
type CustomRequestResult struct {
tdCommon
Result string `json:"result"` // A JSON-serialized result
}

// MessageType return the string telegram-type of CustomRequestResult 
func (customRequestResult *CustomRequestResult) MessageType() string {
 return "customRequestResult" }

// GameHighScore Contains one row of the game high score table  
type GameHighScore struct {
tdCommon
Score int32 `json:"score"` // User score
Position int32 `json:"position"` // Position in the high score table 
UserID int32 `json:"user_id"` // User identifier 
}

// MessageType return the string telegram-type of GameHighScore 
func (gameHighScore *GameHighScore) MessageType() string {
 return "gameHighScore" }

// GameHighScores Contains a list of game high scores  
type GameHighScores struct {
tdCommon
Scores []GameHighScore `json:"scores"` // A list of game high scores
}

// MessageType return the string telegram-type of GameHighScores 
func (gameHighScores *GameHighScores) MessageType() string {
 return "gameHighScores" }

// ChatEventMessageEdited A message was edited  
type ChatEventMessageEdited struct {
tdCommon
OldMessage Message `json:"old_message"` // The original message before the edit 
NewMessage Message `json:"new_message"` // The message after it was edited
}

// MessageType return the string telegram-type of ChatEventMessageEdited 
func (chatEventMessageEdited *ChatEventMessageEdited) MessageType() string {
 return "chatEventMessageEdited" }

// ChatEventMessageDeleted A message was deleted  
type ChatEventMessageDeleted struct {
tdCommon
Message Message `json:"message"` // Deleted message
}

// MessageType return the string telegram-type of ChatEventMessageDeleted 
func (chatEventMessageDeleted *ChatEventMessageDeleted) MessageType() string {
 return "chatEventMessageDeleted" }

// ChatEventMessagePinned A message was pinned  
type ChatEventMessagePinned struct {
tdCommon
Message Message `json:"message"` // Pinned message
}

// MessageType return the string telegram-type of ChatEventMessagePinned 
func (chatEventMessagePinned *ChatEventMessagePinned) MessageType() string {
 return "chatEventMessagePinned" }

// ChatEventMessageUnpinned A message was unpinned 
type ChatEventMessageUnpinned struct {
tdCommon

}

// MessageType return the string telegram-type of ChatEventMessageUnpinned 
func (chatEventMessageUnpinned *ChatEventMessageUnpinned) MessageType() string {
 return "chatEventMessageUnpinned" }

// ChatEventMemberJoined A new member joined the chat 
type ChatEventMemberJoined struct {
tdCommon

}

// MessageType return the string telegram-type of ChatEventMemberJoined 
func (chatEventMemberJoined *ChatEventMemberJoined) MessageType() string {
 return "chatEventMemberJoined" }

// ChatEventMemberLeft A member left the chat 
type ChatEventMemberLeft struct {
tdCommon

}

// MessageType return the string telegram-type of ChatEventMemberLeft 
func (chatEventMemberLeft *ChatEventMemberLeft) MessageType() string {
 return "chatEventMemberLeft" }

// ChatEventMemberInvited A new chat member was invited  
type ChatEventMemberInvited struct {
tdCommon
UserID int32 `json:"user_id"` // New member user identifier 
Status ChatMemberStatus `json:"status"` // New member status
}

// MessageType return the string telegram-type of ChatEventMemberInvited 
func (chatEventMemberInvited *ChatEventMemberInvited) MessageType() string {
 return "chatEventMemberInvited" }

// ChatEventMemberPromoted A chat member has gained/lost administrator status, or the list of their administrator privileges has changed  
type ChatEventMemberPromoted struct {
tdCommon
NewStatus ChatMemberStatus `json:"new_status"` // New status of the chat member
UserID int32 `json:"user_id"` // Chat member user identifier 
OldStatus ChatMemberStatus `json:"old_status"` // Previous status of the chat member 
}

// MessageType return the string telegram-type of ChatEventMemberPromoted 
func (chatEventMemberPromoted *ChatEventMemberPromoted) MessageType() string {
 return "chatEventMemberPromoted" }

// ChatEventMemberRestricted A chat member was restricted/unrestricted or banned/unbanned, or the list of their restrictions has changed  
type ChatEventMemberRestricted struct {
tdCommon
UserID int32 `json:"user_id"` // Chat member user identifier 
OldStatus ChatMemberStatus `json:"old_status"` // Previous status of the chat member 
NewStatus ChatMemberStatus `json:"new_status"` // New status of the chat member
}

// MessageType return the string telegram-type of ChatEventMemberRestricted 
func (chatEventMemberRestricted *ChatEventMemberRestricted) MessageType() string {
 return "chatEventMemberRestricted" }

// ChatEventTitleChanged The chat title was changed  
type ChatEventTitleChanged struct {
tdCommon
OldTitle string `json:"old_title"` // Previous chat title 
NewTitle string `json:"new_title"` // New chat title
}

// MessageType return the string telegram-type of ChatEventTitleChanged 
func (chatEventTitleChanged *ChatEventTitleChanged) MessageType() string {
 return "chatEventTitleChanged" }

// ChatEventDescriptionChanged The chat description was changed  
type ChatEventDescriptionChanged struct {
tdCommon
OldDescription string `json:"old_description"` // Previous chat description 
NewDescription string `json:"new_description"` // New chat description
}

// MessageType return the string telegram-type of ChatEventDescriptionChanged 
func (chatEventDescriptionChanged *ChatEventDescriptionChanged) MessageType() string {
 return "chatEventDescriptionChanged" }

// ChatEventUsernameChanged The chat username was changed  
type ChatEventUsernameChanged struct {
tdCommon
OldUsername string `json:"old_username"` // Previous chat username 
NewUsername string `json:"new_username"` // New chat username
}

// MessageType return the string telegram-type of ChatEventUsernameChanged 
func (chatEventUsernameChanged *ChatEventUsernameChanged) MessageType() string {
 return "chatEventUsernameChanged" }

// ChatEventPhotoChanged The chat photo was changed  
type ChatEventPhotoChanged struct {
tdCommon
OldPhoto ChatPhoto `json:"old_photo"` // Previous chat photo value; may be null 
NewPhoto ChatPhoto `json:"new_photo"` // New chat photo value; may be null
}

// MessageType return the string telegram-type of ChatEventPhotoChanged 
func (chatEventPhotoChanged *ChatEventPhotoChanged) MessageType() string {
 return "chatEventPhotoChanged" }

// ChatEventInvitesToggled The anyone_can_invite setting of a supergroup chat was toggled  
type ChatEventInvitesToggled struct {
tdCommon
AnyoneCanInvite bool `json:"anyone_can_invite"` // New value of anyone_can_invite
}

// MessageType return the string telegram-type of ChatEventInvitesToggled 
func (chatEventInvitesToggled *ChatEventInvitesToggled) MessageType() string {
 return "chatEventInvitesToggled" }

// ChatEventSignMessagesToggled The sign_messages setting of a channel was toggled  
type ChatEventSignMessagesToggled struct {
tdCommon
SignMessages bool `json:"sign_messages"` // New value of sign_messages
}

// MessageType return the string telegram-type of ChatEventSignMessagesToggled 
func (chatEventSignMessagesToggled *ChatEventSignMessagesToggled) MessageType() string {
 return "chatEventSignMessagesToggled" }

// ChatEventStickerSetChanged The supergroup sticker set was changed  
type ChatEventStickerSetChanged struct {
tdCommon
OldStickerSetID int64 `json:"old_sticker_set_id"` // Previous identifier of the chat sticker set; 0 if none 
NewStickerSetID int64 `json:"new_sticker_set_id"` // New identifier of the chat sticker set; 0 if none
}

// MessageType return the string telegram-type of ChatEventStickerSetChanged 
func (chatEventStickerSetChanged *ChatEventStickerSetChanged) MessageType() string {
 return "chatEventStickerSetChanged" }

// ChatEventIsAllHistoryAvailableToggled The is_all_history_available setting of a supergroup was toggled  
type ChatEventIsAllHistoryAvailableToggled struct {
tdCommon
IsAllHistoryAvailable bool `json:"is_all_history_available"` // New value of is_all_history_available
}

// MessageType return the string telegram-type of ChatEventIsAllHistoryAvailableToggled 
func (chatEventIsAllHistoryAvailableToggled *ChatEventIsAllHistoryAvailableToggled) MessageType() string {
 return "chatEventIsAllHistoryAvailableToggled" }

// ChatEvent Represents a chat event  
type ChatEvent struct {
tdCommon
UserID int32 `json:"user_id"` // Identifier of the user who performed the action that triggered the event 
Action ChatEventAction `json:"action"` // Action performed by the user
ID int64 `json:"id"` // Chat event identifier 
Date int32 `json:"date"` // Point in time (Unix timestamp) when the event happened 
}

// MessageType return the string telegram-type of ChatEvent 
func (chatEvent *ChatEvent) MessageType() string {
 return "chatEvent" }

// ChatEvents Contains a list of chat events  
type ChatEvents struct {
tdCommon
Events []ChatEvent `json:"events"` // List of events
}

// MessageType return the string telegram-type of ChatEvents 
func (chatEvents *ChatEvents) MessageType() string {
 return "chatEvents" }

// ChatEventLogFilters Represents a set of filters used to obtain a chat event log 
type ChatEventLogFilters struct {
tdCommon
MessageEdits bool `json:"message_edits"` // True, if message edits should be returned
MemberLeaves bool `json:"member_leaves"` // True, if members leaving events should be returned
MemberPromotions bool `json:"member_promotions"` // True, if member promotion/demotion events should be returned
InfoChanges bool `json:"info_changes"` // True, if changes in chat information should be returned
SettingChanges bool `json:"setting_changes"` // True, if changes in chat settings should be returned
MessageDeletions bool `json:"message_deletions"` // True, if message deletions should be returned
MessagePins bool `json:"message_pins"` // True, if pin/unpin events should be returned
MemberJoins bool `json:"member_joins"` // True, if members joining events should be returned
MemberInvites bool `json:"member_invites"` // True, if invited member events should be returned
MemberRestrictions bool `json:"member_restrictions"` // True, if member restricted/unrestricted/banned/unbanned events should be returned
}

// MessageType return the string telegram-type of ChatEventLogFilters 
func (chatEventLogFilters *ChatEventLogFilters) MessageType() string {
 return "chatEventLogFilters" }

// DeviceTokenGoogleCloudMessaging A token for Google Cloud Messaging  
type DeviceTokenGoogleCloudMessaging struct {
tdCommon
Token string `json:"token"` // Device registration token, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenGoogleCloudMessaging 
func (deviceTokenGoogleCloudMessaging *DeviceTokenGoogleCloudMessaging) MessageType() string {
 return "deviceTokenGoogleCloudMessaging" }

// DeviceTokenApplePush A token for Apple Push Notification service  
type DeviceTokenApplePush struct {
tdCommon
DeviceToken string `json:"device_token"` // Device token, may be empty to de-register a device 
IsAppSandbox bool `json:"is_app_sandbox"` // True, if App Sandbox is enabled
}

// MessageType return the string telegram-type of DeviceTokenApplePush 
func (deviceTokenApplePush *DeviceTokenApplePush) MessageType() string {
 return "deviceTokenApplePush" }

// DeviceTokenApplePushVoIP A token for Apple Push Notification service VoIP notifications  
type DeviceTokenApplePushVoIP struct {
tdCommon
DeviceToken string `json:"device_token"` // Device token, may be empty to de-register a device 
IsAppSandbox bool `json:"is_app_sandbox"` // True, if App Sandbox is enabled
}

// MessageType return the string telegram-type of DeviceTokenApplePushVoIP 
func (deviceTokenApplePushVoIP *DeviceTokenApplePushVoIP) MessageType() string {
 return "deviceTokenApplePushVoIP" }

// DeviceTokenWindowsPush A token for Windows Push Notification Services  
type DeviceTokenWindowsPush struct {
tdCommon
AccessToken string `json:"access_token"` // The access token that will be used to send notifications, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenWindowsPush 
func (deviceTokenWindowsPush *DeviceTokenWindowsPush) MessageType() string {
 return "deviceTokenWindowsPush" }

// DeviceTokenMicrosoftPush A token for Microsoft Push Notification Service  
type DeviceTokenMicrosoftPush struct {
tdCommon
ChannelURI string `json:"channel_uri"` // Push notification channel URI, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenMicrosoftPush 
func (deviceTokenMicrosoftPush *DeviceTokenMicrosoftPush) MessageType() string {
 return "deviceTokenMicrosoftPush" }

// DeviceTokenMicrosoftPushVoIP A token for Microsoft Push Notification Service VoIP channel  
type DeviceTokenMicrosoftPushVoIP struct {
tdCommon
ChannelURI string `json:"channel_uri"` // Push notification channel URI, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenMicrosoftPushVoIP 
func (deviceTokenMicrosoftPushVoIP *DeviceTokenMicrosoftPushVoIP) MessageType() string {
 return "deviceTokenMicrosoftPushVoIP" }

// DeviceTokenWebPush A token for web Push API  
type DeviceTokenWebPush struct {
tdCommon
Endpoint string `json:"endpoint"` // Absolute URL exposed by the push service where the application server can send push messages, may be empty to de-register a device
P256dhBase64url string `json:"p256dh_base64url"` // Base64url-encoded P-256 elliptic curve Diffie-Hellman public key 
AuthBase64url string `json:"auth_base64url"` // Base64url-encoded authentication secret
}

// MessageType return the string telegram-type of DeviceTokenWebPush 
func (deviceTokenWebPush *DeviceTokenWebPush) MessageType() string {
 return "deviceTokenWebPush" }

// DeviceTokenSimplePush A token for Simple Push API for Firefox OS  
type DeviceTokenSimplePush struct {
tdCommon
Endpoint string `json:"endpoint"` // Absolute URL exposed by the push service where the application server can send push messages, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenSimplePush 
func (deviceTokenSimplePush *DeviceTokenSimplePush) MessageType() string {
 return "deviceTokenSimplePush" }

// DeviceTokenUbuntuPush A token for Ubuntu Push Client service  
type DeviceTokenUbuntuPush struct {
tdCommon
Token string `json:"token"` // Token, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenUbuntuPush 
func (deviceTokenUbuntuPush *DeviceTokenUbuntuPush) MessageType() string {
 return "deviceTokenUbuntuPush" }

// DeviceTokenBlackBerryPush A token for BlackBerry Push Service  
type DeviceTokenBlackBerryPush struct {
tdCommon
Token string `json:"token"` // Token, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenBlackBerryPush 
func (deviceTokenBlackBerryPush *DeviceTokenBlackBerryPush) MessageType() string {
 return "deviceTokenBlackBerryPush" }

// DeviceTokenTizenPush A token for Tizen Push Service  
type DeviceTokenTizenPush struct {
tdCommon
RegID string `json:"reg_id"` // Push service registration identifier, may be empty to de-register a device
}

// MessageType return the string telegram-type of DeviceTokenTizenPush 
func (deviceTokenTizenPush *DeviceTokenTizenPush) MessageType() string {
 return "deviceTokenTizenPush" }

// Wallpaper Contains information about a wallpaper  
type Wallpaper struct {
tdCommon
ID int32 `json:"id"` // Unique persistent wallpaper identifier 
Sizes []PhotoSize `json:"sizes"` // Available variants of the wallpaper in different sizes. These photos can only be downloaded; they can't be sent in a message 
Color int32 `json:"color"` // Main color of the wallpaper in RGB24 format; should be treated as background color if no photos are specified
}

// MessageType return the string telegram-type of Wallpaper 
func (wallpaper *Wallpaper) MessageType() string {
 return "wallpaper" }

// Wallpapers Contains a list of wallpapers  
type Wallpapers struct {
tdCommon
Wallpapers []Wallpaper `json:"wallpapers"` // A list of wallpapers
}

// MessageType return the string telegram-type of Wallpapers 
func (wallpapers *Wallpapers) MessageType() string {
 return "wallpapers" }

// Hashtags Contains a list of hashtags  
type Hashtags struct {
tdCommon
Hashtags []string `json:"hashtags"` // A list of hashtags
}

// MessageType return the string telegram-type of Hashtags 
func (hashtags *Hashtags) MessageType() string {
 return "hashtags" }

// CheckChatUsernameResultOk The username can be set 
type CheckChatUsernameResultOk struct {
tdCommon

}

// MessageType return the string telegram-type of CheckChatUsernameResultOk 
func (checkChatUsernameResultOk *CheckChatUsernameResultOk) MessageType() string {
 return "checkChatUsernameResultOk" }

// CheckChatUsernameResultUsernameInvalid The username is invalid 
type CheckChatUsernameResultUsernameInvalid struct {
tdCommon

}

// MessageType return the string telegram-type of CheckChatUsernameResultUsernameInvalid 
func (checkChatUsernameResultUsernameInvalid *CheckChatUsernameResultUsernameInvalid) MessageType() string {
 return "checkChatUsernameResultUsernameInvalid" }

// CheckChatUsernameResultUsernameOccupied The username is occupied 
type CheckChatUsernameResultUsernameOccupied struct {
tdCommon

}

// MessageType return the string telegram-type of CheckChatUsernameResultUsernameOccupied 
func (checkChatUsernameResultUsernameOccupied *CheckChatUsernameResultUsernameOccupied) MessageType() string {
 return "checkChatUsernameResultUsernameOccupied" }

// CheckChatUsernameResultPublicChatsTooMuch The user has too much public chats, one of them should be made private first 
type CheckChatUsernameResultPublicChatsTooMuch struct {
tdCommon

}

// MessageType return the string telegram-type of CheckChatUsernameResultPublicChatsTooMuch 
func (checkChatUsernameResultPublicChatsTooMuch *CheckChatUsernameResultPublicChatsTooMuch) MessageType() string {
 return "checkChatUsernameResultPublicChatsTooMuch" }

// CheckChatUsernameResultPublicGroupsUnavailable The user can't be a member of a public supergroup 
type CheckChatUsernameResultPublicGroupsUnavailable struct {
tdCommon

}

// MessageType return the string telegram-type of CheckChatUsernameResultPublicGroupsUnavailable 
func (checkChatUsernameResultPublicGroupsUnavailable *CheckChatUsernameResultPublicGroupsUnavailable) MessageType() string {
 return "checkChatUsernameResultPublicGroupsUnavailable" }

// OptionValueBoolean Boolean option  
type OptionValueBoolean struct {
tdCommon
Value bool `json:"value"` // The value of the option
}

// MessageType return the string telegram-type of OptionValueBoolean 
func (optionValueBoolean *OptionValueBoolean) MessageType() string {
 return "optionValueBoolean" }

// OptionValueEmpty An unknown option or an option which has a default value 
type OptionValueEmpty struct {
tdCommon

}

// MessageType return the string telegram-type of OptionValueEmpty 
func (optionValueEmpty *OptionValueEmpty) MessageType() string {
 return "optionValueEmpty" }

// OptionValueInteger An integer option  
type OptionValueInteger struct {
tdCommon
Value int32 `json:"value"` // The value of the option
}

// MessageType return the string telegram-type of OptionValueInteger 
func (optionValueInteger *OptionValueInteger) MessageType() string {
 return "optionValueInteger" }

// OptionValueString A string option  
type OptionValueString struct {
tdCommon
Value string `json:"value"` // The value of the option
}

// MessageType return the string telegram-type of OptionValueString 
func (optionValueString *OptionValueString) MessageType() string {
 return "optionValueString" }

// UserPrivacySettingRuleAllowAll A rule to allow all users to do something 
type UserPrivacySettingRuleAllowAll struct {
tdCommon

}

// MessageType return the string telegram-type of UserPrivacySettingRuleAllowAll 
func (userPrivacySettingRuleAllowAll *UserPrivacySettingRuleAllowAll) MessageType() string {
 return "userPrivacySettingRuleAllowAll" }

// UserPrivacySettingRuleAllowContacts A rule to allow all of a user's contacts to do something 
type UserPrivacySettingRuleAllowContacts struct {
tdCommon

}

// MessageType return the string telegram-type of UserPrivacySettingRuleAllowContacts 
func (userPrivacySettingRuleAllowContacts *UserPrivacySettingRuleAllowContacts) MessageType() string {
 return "userPrivacySettingRuleAllowContacts" }

// UserPrivacySettingRuleAllowUsers A rule to allow certain specified users to do something  
type UserPrivacySettingRuleAllowUsers struct {
tdCommon
UserIDs []int32 `json:"user_ids"` // The user identifiers
}

// MessageType return the string telegram-type of UserPrivacySettingRuleAllowUsers 
func (userPrivacySettingRuleAllowUsers *UserPrivacySettingRuleAllowUsers) MessageType() string {
 return "userPrivacySettingRuleAllowUsers" }

// UserPrivacySettingRuleRestrictAll A rule to restrict all users from doing something 
type UserPrivacySettingRuleRestrictAll struct {
tdCommon

}

// MessageType return the string telegram-type of UserPrivacySettingRuleRestrictAll 
func (userPrivacySettingRuleRestrictAll *UserPrivacySettingRuleRestrictAll) MessageType() string {
 return "userPrivacySettingRuleRestrictAll" }

// UserPrivacySettingRuleRestrictContacts A rule to restrict all contacts of a user from doing something 
type UserPrivacySettingRuleRestrictContacts struct {
tdCommon

}

// MessageType return the string telegram-type of UserPrivacySettingRuleRestrictContacts 
func (userPrivacySettingRuleRestrictContacts *UserPrivacySettingRuleRestrictContacts) MessageType() string {
 return "userPrivacySettingRuleRestrictContacts" }

// UserPrivacySettingRuleRestrictUsers A rule to restrict all specified users from doing something  
type UserPrivacySettingRuleRestrictUsers struct {
tdCommon
UserIDs []int32 `json:"user_ids"` // The user identifiers
}

// MessageType return the string telegram-type of UserPrivacySettingRuleRestrictUsers 
func (userPrivacySettingRuleRestrictUsers *UserPrivacySettingRuleRestrictUsers) MessageType() string {
 return "userPrivacySettingRuleRestrictUsers" }

// UserPrivacySettingRules A list of privacy rules. Rules are matched in the specified order. The first matched rule defines the privacy setting for a given user. If no rule matches, the action is not allowed  
type UserPrivacySettingRules struct {
tdCommon
Rules []UserPrivacySettingRule `json:"rules"` // A list of rules
}

// MessageType return the string telegram-type of UserPrivacySettingRules 
func (userPrivacySettingRules *UserPrivacySettingRules) MessageType() string {
 return "userPrivacySettingRules" }

// UserPrivacySettingShowStatus A privacy setting for managing whether the user's online status is visible 
type UserPrivacySettingShowStatus struct {
tdCommon

}

// MessageType return the string telegram-type of UserPrivacySettingShowStatus 
func (userPrivacySettingShowStatus *UserPrivacySettingShowStatus) MessageType() string {
 return "userPrivacySettingShowStatus" }

// UserPrivacySettingAllowChatInvites A privacy setting for managing whether the user can be invited to chats 
type UserPrivacySettingAllowChatInvites struct {
tdCommon

}

// MessageType return the string telegram-type of UserPrivacySettingAllowChatInvites 
func (userPrivacySettingAllowChatInvites *UserPrivacySettingAllowChatInvites) MessageType() string {
 return "userPrivacySettingAllowChatInvites" }

// UserPrivacySettingAllowCalls A privacy setting for managing whether the user can be called 
type UserPrivacySettingAllowCalls struct {
tdCommon

}

// MessageType return the string telegram-type of UserPrivacySettingAllowCalls 
func (userPrivacySettingAllowCalls *UserPrivacySettingAllowCalls) MessageType() string {
 return "userPrivacySettingAllowCalls" }

// AccountTTL Contains information about the period of inactivity after which the current user's account will automatically be deleted  
type AccountTTL struct {
tdCommon
Days int32 `json:"days"` // Number of days of inactivity before the account will be flagged for deletion; should range from 30-366 days
}

// MessageType return the string telegram-type of AccountTTL 
func (accountTTL *AccountTTL) MessageType() string {
 return "accountTtl" }

// Session Contains information about one session in a Telegram application used by the current user  
type Session struct {
tdCommon
IsCurrent bool `json:"is_current"` // True, if this session is the current session
DeviceModel string `json:"device_model"` // Model of the device the application has been run or is running on, as provided by the application 
IP string `json:"ip"` // IP address from which the session was created, in human-readable format
ID int64 `json:"id"` // Session identifier 
APIID int32 `json:"api_id"` // Telegram API identifier, as provided by the application 
Platform string `json:"platform"` // Operating system the application has been run or is running on, as provided by the application
Region string `json:"region"` // Region code from which the session was created, based on the IP address
LogInDate int32 `json:"log_in_date"` // Point in time (Unix timestamp) when the user has logged in
LastActiveDate int32 `json:"last_active_date"` // Point in time (Unix timestamp) when the session was last used 
Country string `json:"country"` // A two-letter country code for the country from which the session was created, based on the IP address 
ApplicationName string `json:"application_name"` // Name of the application, as provided by the application
ApplicationVersion string `json:"application_version"` // The version of the application, as provided by the application 
IsOfficialApplication bool `json:"is_official_application"` // True, if the application is an official application or uses the api_id of an official application
SystemVersion string `json:"system_version"` // Version of the operating system the application has been run or is running on, as provided by the application 
}

// MessageType return the string telegram-type of Session 
func (session *Session) MessageType() string {
 return "session" }

// Sessions Contains a list of sessions  
type Sessions struct {
tdCommon
Sessions []Session `json:"sessions"` // List of sessions
}

// MessageType return the string telegram-type of Sessions 
func (sessions *Sessions) MessageType() string {
 return "sessions" }

// ConnectedWebsite Contains information about one website the current user is logged in with Telegram 
type ConnectedWebsite struct {
tdCommon
IP string `json:"ip"` // IP address from which the user was logged in, in human-readable format
ID int64 `json:"id"` // Website identifier
DomainName string `json:"domain_name"` // The domain name of the website
BotUserID int32 `json:"bot_user_id"` // User identifier of a bot linked with the website
Browser string `json:"browser"` // The version of a browser used to log in
Platform string `json:"platform"` // Operating system the browser is running on
LogInDate int32 `json:"log_in_date"` // Point in time (Unix timestamp) when the user was logged in
LastActiveDate int32 `json:"last_active_date"` // Point in time (Unix timestamp) when obtained authorization was last used
Location string `json:"location"` // Human-readable description of a country and a region, from which the user was logged in, based on the IP address
}

// MessageType return the string telegram-type of ConnectedWebsite 
func (connectedWebsite *ConnectedWebsite) MessageType() string {
 return "connectedWebsite" }

// ConnectedWebsites Contains a list of websites the current user is logged in with Telegram  
type ConnectedWebsites struct {
tdCommon
Websites []ConnectedWebsite `json:"websites"` // List of connected websites
}

// MessageType return the string telegram-type of ConnectedWebsites 
func (connectedWebsites *ConnectedWebsites) MessageType() string {
 return "connectedWebsites" }

// ChatReportSpamState Contains information about the availability of the "Report spam" action for a chat  
type ChatReportSpamState struct {
tdCommon
CanReportSpam bool `json:"can_report_spam"` // True, if a prompt with the "Report spam" action should be shown to the user
}

// MessageType return the string telegram-type of ChatReportSpamState 
func (chatReportSpamState *ChatReportSpamState) MessageType() string {
 return "chatReportSpamState" }

// ChatReportReasonSpam The chat contains spam messages 
type ChatReportReasonSpam struct {
tdCommon

}

// MessageType return the string telegram-type of ChatReportReasonSpam 
func (chatReportReasonSpam *ChatReportReasonSpam) MessageType() string {
 return "chatReportReasonSpam" }

// ChatReportReasonViolence The chat promotes violence 
type ChatReportReasonViolence struct {
tdCommon

}

// MessageType return the string telegram-type of ChatReportReasonViolence 
func (chatReportReasonViolence *ChatReportReasonViolence) MessageType() string {
 return "chatReportReasonViolence" }

// ChatReportReasonPornography The chat contains pornographic messages 
type ChatReportReasonPornography struct {
tdCommon

}

// MessageType return the string telegram-type of ChatReportReasonPornography 
func (chatReportReasonPornography *ChatReportReasonPornography) MessageType() string {
 return "chatReportReasonPornography" }

// ChatReportReasonCustom A custom reason provided by the user  
type ChatReportReasonCustom struct {
tdCommon
Text string `json:"text"` // Report text
}

// MessageType return the string telegram-type of ChatReportReasonCustom 
func (chatReportReasonCustom *ChatReportReasonCustom) MessageType() string {
 return "chatReportReasonCustom" }

// PublicMessageLink Contains a public HTTPS link to a message in a public supergroup or channel  
type PublicMessageLink struct {
tdCommon
HTML string `json:"html"` // HTML-code for embedding the message
Link string `json:"link"` // Message link 
}

// MessageType return the string telegram-type of PublicMessageLink 
func (publicMessageLink *PublicMessageLink) MessageType() string {
 return "publicMessageLink" }

// FileTypeNone The data is not a file 
type FileTypeNone struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeNone 
func (fileTypeNone *FileTypeNone) MessageType() string {
 return "fileTypeNone" }

// FileTypeAnimation The file is an animation 
type FileTypeAnimation struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeAnimation 
func (fileTypeAnimation *FileTypeAnimation) MessageType() string {
 return "fileTypeAnimation" }

// FileTypeAudio The file is an audio file 
type FileTypeAudio struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeAudio 
func (fileTypeAudio *FileTypeAudio) MessageType() string {
 return "fileTypeAudio" }

// FileTypeDocument The file is a document 
type FileTypeDocument struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeDocument 
func (fileTypeDocument *FileTypeDocument) MessageType() string {
 return "fileTypeDocument" }

// FileTypePhoto The file is a photo 
type FileTypePhoto struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypePhoto 
func (fileTypePhoto *FileTypePhoto) MessageType() string {
 return "fileTypePhoto" }

// FileTypeProfilePhoto The file is a profile photo 
type FileTypeProfilePhoto struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeProfilePhoto 
func (fileTypeProfilePhoto *FileTypeProfilePhoto) MessageType() string {
 return "fileTypeProfilePhoto" }

// FileTypeSecret The file was sent to a secret chat (the file type is not known to the server) 
type FileTypeSecret struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeSecret 
func (fileTypeSecret *FileTypeSecret) MessageType() string {
 return "fileTypeSecret" }

// FileTypeSticker The file is a sticker 
type FileTypeSticker struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeSticker 
func (fileTypeSticker *FileTypeSticker) MessageType() string {
 return "fileTypeSticker" }

// FileTypeThumbnail The file is a thumbnail of another file 
type FileTypeThumbnail struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeThumbnail 
func (fileTypeThumbnail *FileTypeThumbnail) MessageType() string {
 return "fileTypeThumbnail" }

// FileTypeUnknown The file type is not yet known 
type FileTypeUnknown struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeUnknown 
func (fileTypeUnknown *FileTypeUnknown) MessageType() string {
 return "fileTypeUnknown" }

// FileTypeVideo The file is a video 
type FileTypeVideo struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeVideo 
func (fileTypeVideo *FileTypeVideo) MessageType() string {
 return "fileTypeVideo" }

// FileTypeVideoNote The file is a video note 
type FileTypeVideoNote struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeVideoNote 
func (fileTypeVideoNote *FileTypeVideoNote) MessageType() string {
 return "fileTypeVideoNote" }

// FileTypeVoiceNote The file is a voice note 
type FileTypeVoiceNote struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeVoiceNote 
func (fileTypeVoiceNote *FileTypeVoiceNote) MessageType() string {
 return "fileTypeVoiceNote" }

// FileTypeWallpaper The file is a wallpaper 
type FileTypeWallpaper struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeWallpaper 
func (fileTypeWallpaper *FileTypeWallpaper) MessageType() string {
 return "fileTypeWallpaper" }

// FileTypeSecretThumbnail The file is a thumbnail of a file from a secret chat 
type FileTypeSecretThumbnail struct {
tdCommon

}

// MessageType return the string telegram-type of FileTypeSecretThumbnail 
func (fileTypeSecretThumbnail *FileTypeSecretThumbnail) MessageType() string {
 return "fileTypeSecretThumbnail" }

// StorageStatisticsByFileType Contains the storage usage statistics for a specific file type  
type StorageStatisticsByFileType struct {
tdCommon
FileType FileType `json:"file_type"` // File type 
Size int64 `json:"size"` // Total size of the files 
Count int32 `json:"count"` // Total number of files
}

// MessageType return the string telegram-type of StorageStatisticsByFileType 
func (storageStatisticsByFileType *StorageStatisticsByFileType) MessageType() string {
 return "storageStatisticsByFileType" }

// StorageStatisticsByChat Contains the storage usage statistics for a specific chat  
type StorageStatisticsByChat struct {
tdCommon
Size int64 `json:"size"` // Total size of the files in the chat 
Count int32 `json:"count"` // Total number of files in the chat 
ByFileType []StorageStatisticsByFileType `json:"by_file_type"` // Statistics split by file types
ChatID int64 `json:"chat_id"` // Chat identifier; 0 if none 
}

// MessageType return the string telegram-type of StorageStatisticsByChat 
func (storageStatisticsByChat *StorageStatisticsByChat) MessageType() string {
 return "storageStatisticsByChat" }

// StorageStatistics Contains the exact storage usage statistics split by chats and file type  
type StorageStatistics struct {
tdCommon
Size int64 `json:"size"` // Total size of files 
Count int32 `json:"count"` // Total number of files 
ByChat []StorageStatisticsByChat `json:"by_chat"` // Statistics split by chats
}

// MessageType return the string telegram-type of StorageStatistics 
func (storageStatistics *StorageStatistics) MessageType() string {
 return "storageStatistics" }

// StorageStatisticsFast Contains approximate storage usage statistics, excluding files of unknown file type  
type StorageStatisticsFast struct {
tdCommon
DatabaseSize int64 `json:"database_size"` // Size of the database
FilesSize int64 `json:"files_size"` // Approximate total size of files 
FileCount int32 `json:"file_count"` // Approximate number of files 
}

// MessageType return the string telegram-type of StorageStatisticsFast 
func (storageStatisticsFast *StorageStatisticsFast) MessageType() string {
 return "storageStatisticsFast" }

// NetworkTypeNone The network is not available 
type NetworkTypeNone struct {
tdCommon

}

// MessageType return the string telegram-type of NetworkTypeNone 
func (networkTypeNone *NetworkTypeNone) MessageType() string {
 return "networkTypeNone" }

// NetworkTypeMobile A mobile network 
type NetworkTypeMobile struct {
tdCommon

}

// MessageType return the string telegram-type of NetworkTypeMobile 
func (networkTypeMobile *NetworkTypeMobile) MessageType() string {
 return "networkTypeMobile" }

// NetworkTypeMobileRoaming A mobile roaming network 
type NetworkTypeMobileRoaming struct {
tdCommon

}

// MessageType return the string telegram-type of NetworkTypeMobileRoaming 
func (networkTypeMobileRoaming *NetworkTypeMobileRoaming) MessageType() string {
 return "networkTypeMobileRoaming" }

// NetworkTypeWiFi A Wi-Fi network 
type NetworkTypeWiFi struct {
tdCommon

}

// MessageType return the string telegram-type of NetworkTypeWiFi 
func (networkTypeWiFi *NetworkTypeWiFi) MessageType() string {
 return "networkTypeWiFi" }

// NetworkTypeOther A different network type (e.g., Ethernet network) 
type NetworkTypeOther struct {
tdCommon

}

// MessageType return the string telegram-type of NetworkTypeOther 
func (networkTypeOther *NetworkTypeOther) MessageType() string {
 return "networkTypeOther" }

// NetworkStatisticsEntryFile Contains information about the total amount of data that was used to send and receive files  
type NetworkStatisticsEntryFile struct {
tdCommon
SentBytes int64 `json:"sent_bytes"` // Total number of bytes sent 
ReceivedBytes int64 `json:"received_bytes"` // Total number of bytes received
FileType FileType `json:"file_type"` // Type of the file the data is part of 
NetworkType NetworkType `json:"network_type"` // Type of the network the data was sent through. Call setNetworkType to maintain the actual network type
}

// MessageType return the string telegram-type of NetworkStatisticsEntryFile 
func (networkStatisticsEntryFile *NetworkStatisticsEntryFile) MessageType() string {
 return "networkStatisticsEntryFile" }

// NetworkStatisticsEntryCall Contains information about the total amount of data that was used for calls  
type NetworkStatisticsEntryCall struct {
tdCommon
SentBytes int64 `json:"sent_bytes"` // Total number of bytes sent 
ReceivedBytes int64 `json:"received_bytes"` // Total number of bytes received 
Duration float64 `json:"duration"` // Total call duration, in seconds
NetworkType NetworkType `json:"network_type"` // Type of the network the data was sent through. Call setNetworkType to maintain the actual network type
}

// MessageType return the string telegram-type of NetworkStatisticsEntryCall 
func (networkStatisticsEntryCall *NetworkStatisticsEntryCall) MessageType() string {
 return "networkStatisticsEntryCall" }

// NetworkStatistics A full list of available network statistic entries  
type NetworkStatistics struct {
tdCommon
SinceDate int32 `json:"since_date"` // Point in time (Unix timestamp) when the app began collecting statistics 
Entries []NetworkStatisticsEntry `json:"entries"` // Network statistics entries
}

// MessageType return the string telegram-type of NetworkStatistics 
func (networkStatistics *NetworkStatistics) MessageType() string {
 return "networkStatistics" }

// ConnectionStateWaitingForNetwork Currently waiting for the network to become available. Use SetNetworkType to change the available network type 
type ConnectionStateWaitingForNetwork struct {
tdCommon

}

// MessageType return the string telegram-type of ConnectionStateWaitingForNetwork 
func (connectionStateWaitingForNetwork *ConnectionStateWaitingForNetwork) MessageType() string {
 return "connectionStateWaitingForNetwork" }

// ConnectionStateConnectingToProxy Currently establishing a connection with a proxy server 
type ConnectionStateConnectingToProxy struct {
tdCommon

}

// MessageType return the string telegram-type of ConnectionStateConnectingToProxy 
func (connectionStateConnectingToProxy *ConnectionStateConnectingToProxy) MessageType() string {
 return "connectionStateConnectingToProxy" }

// ConnectionStateConnecting Currently establishing a connection to the Telegram servers 
type ConnectionStateConnecting struct {
tdCommon

}

// MessageType return the string telegram-type of ConnectionStateConnecting 
func (connectionStateConnecting *ConnectionStateConnecting) MessageType() string {
 return "connectionStateConnecting" }

// ConnectionStateUpdating Downloading data received while the client was offline 
type ConnectionStateUpdating struct {
tdCommon

}

// MessageType return the string telegram-type of ConnectionStateUpdating 
func (connectionStateUpdating *ConnectionStateUpdating) MessageType() string {
 return "connectionStateUpdating" }

// ConnectionStateReady There is a working connection to the Telegram servers 
type ConnectionStateReady struct {
tdCommon

}

// MessageType return the string telegram-type of ConnectionStateReady 
func (connectionStateReady *ConnectionStateReady) MessageType() string {
 return "connectionStateReady" }

// TopChatCategoryUsers A category containing frequently used private chats with non-bot users 
type TopChatCategoryUsers struct {
tdCommon

}

// MessageType return the string telegram-type of TopChatCategoryUsers 
func (topChatCategoryUsers *TopChatCategoryUsers) MessageType() string {
 return "topChatCategoryUsers" }

// TopChatCategoryBots A category containing frequently used private chats with bot users 
type TopChatCategoryBots struct {
tdCommon

}

// MessageType return the string telegram-type of TopChatCategoryBots 
func (topChatCategoryBots *TopChatCategoryBots) MessageType() string {
 return "topChatCategoryBots" }

// TopChatCategoryGroups A category containing frequently used basic groups and supergroups 
type TopChatCategoryGroups struct {
tdCommon

}

// MessageType return the string telegram-type of TopChatCategoryGroups 
func (topChatCategoryGroups *TopChatCategoryGroups) MessageType() string {
 return "topChatCategoryGroups" }

// TopChatCategoryChannels A category containing frequently used channels 
type TopChatCategoryChannels struct {
tdCommon

}

// MessageType return the string telegram-type of TopChatCategoryChannels 
func (topChatCategoryChannels *TopChatCategoryChannels) MessageType() string {
 return "topChatCategoryChannels" }

// TopChatCategoryInlineBots A category containing frequently used chats with inline bots sorted by their usage in inline mode 
type TopChatCategoryInlineBots struct {
tdCommon

}

// MessageType return the string telegram-type of TopChatCategoryInlineBots 
func (topChatCategoryInlineBots *TopChatCategoryInlineBots) MessageType() string {
 return "topChatCategoryInlineBots" }

// TopChatCategoryCalls A category containing frequently used chats used for calls 
type TopChatCategoryCalls struct {
tdCommon

}

// MessageType return the string telegram-type of TopChatCategoryCalls 
func (topChatCategoryCalls *TopChatCategoryCalls) MessageType() string {
 return "topChatCategoryCalls" }

// TMeURLTypeUser A URL linking to a user  
type TMeURLTypeUser struct {
tdCommon
UserID int32 `json:"user_id"` // Identifier of the user
}

// MessageType return the string telegram-type of TMeURLTypeUser 
func (tMeURLTypeUser *TMeURLTypeUser) MessageType() string {
 return "tMeUrlTypeUser" }

// TMeURLTypeSupergroup A URL linking to a public supergroup or channel  
type TMeURLTypeSupergroup struct {
tdCommon
SupergroupID int64 `json:"supergroup_id"` // Identifier of the supergroup or channel
}

// MessageType return the string telegram-type of TMeURLTypeSupergroup 
func (tMeURLTypeSupergroup *TMeURLTypeSupergroup) MessageType() string {
 return "tMeUrlTypeSupergroup" }

// TMeURLTypeChatInvite A chat invite link  
type TMeURLTypeChatInvite struct {
tdCommon
Info ChatInviteLinkInfo `json:"info"` // Chat invite link info
}

// MessageType return the string telegram-type of TMeURLTypeChatInvite 
func (tMeURLTypeChatInvite *TMeURLTypeChatInvite) MessageType() string {
 return "tMeUrlTypeChatInvite" }

// TMeURLTypeStickerSet A URL linking to a sticker set  
type TMeURLTypeStickerSet struct {
tdCommon
StickerSetID int64 `json:"sticker_set_id"` // Identifier of the sticker set
}

// MessageType return the string telegram-type of TMeURLTypeStickerSet 
func (tMeURLTypeStickerSet *TMeURLTypeStickerSet) MessageType() string {
 return "tMeUrlTypeStickerSet" }

// TMeURL Represents a URL linking to an internal Telegram entity  
type TMeURL struct {
tdCommon
URL string `json:"url"` // URL 
Type TMeURLType `json:"type"` // Type of the URL
}

// MessageType return the string telegram-type of TMeURL 
func (tMeURL *TMeURL) MessageType() string {
 return "tMeUrl" }

// TMeURLs Contains a list of t.me URLs  
type TMeURLs struct {
tdCommon
URLs []TMeURL `json:"urls"` // List of URLs
}

// MessageType return the string telegram-type of TMeURLs 
func (tMeURLs *TMeURLs) MessageType() string {
 return "tMeUrls" }

// Count Contains a counter  
type Count struct {
tdCommon
Count int32 `json:"count"` // Count
}

// MessageType return the string telegram-type of Count 
func (count *Count) MessageType() string {
 return "count" }

// Text Contains some text  
type Text struct {
tdCommon
Text string `json:"text"` // Text
}

// MessageType return the string telegram-type of Text 
func (text *Text) MessageType() string {
 return "text" }

// TextParseModeMarkdown The text should be parsed in markdown-style 
type TextParseModeMarkdown struct {
tdCommon

}

// MessageType return the string telegram-type of TextParseModeMarkdown 
func (textParseModeMarkdown *TextParseModeMarkdown) MessageType() string {
 return "textParseModeMarkdown" }

// TextParseModeHTML The text should be parsed in HTML-style 
type TextParseModeHTML struct {
tdCommon

}

// MessageType return the string telegram-type of TextParseModeHTML 
func (textParseModeHTML *TextParseModeHTML) MessageType() string {
 return "textParseModeHTML" }

// ProxyEmpty An empty proxy server 
type ProxyEmpty struct {
tdCommon

}

// MessageType return the string telegram-type of ProxyEmpty 
func (proxyEmpty *ProxyEmpty) MessageType() string {
 return "proxyEmpty" }

// ProxySocks5 A SOCKS5 proxy server  
type ProxySocks5 struct {
tdCommon
Username string `json:"username"` // Username for logging in 
Password string `json:"password"` // Password for logging in
Server string `json:"server"` // Proxy server IP address 
Port int32 `json:"port"` // Proxy server port 
}

// MessageType return the string telegram-type of ProxySocks5 
func (proxySocks5 *ProxySocks5) MessageType() string {
 return "proxySocks5" }

// InputSticker Describes a sticker that should be added to a sticker set  
type InputSticker struct {
tdCommon
PngSticker InputFile `json:"png_sticker"` // PNG image with the sticker; must be up to 512 kB in size and fit in a 512x512 square 
Emojis string `json:"emojis"` // Emoji corresponding to the sticker 
MaskPosition MaskPosition `json:"mask_position"` // For masks, position where the mask should be placed; may be null
}

// MessageType return the string telegram-type of InputSticker 
func (inputSticker *InputSticker) MessageType() string {
 return "inputSticker" }

// UpdateAuthorizationState The user authorization state has changed  
type UpdateAuthorizationState struct {
tdCommon
AuthorizationState AuthorizationState `json:"authorization_state"` // New authorization state
}

// MessageType return the string telegram-type of UpdateAuthorizationState 
func (updateAuthorizationState *UpdateAuthorizationState) MessageType() string {
 return "updateAuthorizationState" }

// UpdateNewMessage A new message was received; can also be an outgoing message  
type UpdateNewMessage struct {
tdCommon
Message Message `json:"message"` // The new message 
DisableNotification bool `json:"disable_notification"` // True, if this message must not generate a notification 
ContainsMention bool `json:"contains_mention"` // True, if the message contains a mention of the current user
}

// MessageType return the string telegram-type of UpdateNewMessage 
func (updateNewMessage *UpdateNewMessage) MessageType() string {
 return "updateNewMessage" }

// UpdateMessageSendAcknowledged A request to send a message has reached the Telegram server. This doesn't mean that the message will be sent successfully or even that the send message request will be processed. This update will be sent only if the option "use_quick_ack" is set to true. This update may be sent multiple times for the same message 
type UpdateMessageSendAcknowledged struct {
tdCommon
MessageID int64 `json:"message_id"` // A temporary message identifier
ChatID int64 `json:"chat_id"` // The chat identifier of the sent message 
}

// MessageType return the string telegram-type of UpdateMessageSendAcknowledged 
func (updateMessageSendAcknowledged *UpdateMessageSendAcknowledged) MessageType() string {
 return "updateMessageSendAcknowledged" }

// UpdateMessageSendSucceeded A message has been successfully sent  
type UpdateMessageSendSucceeded struct {
tdCommon
Message Message `json:"message"` // Information about the sent message. Usually only the message identifier, date, and content are changed, but almost all other fields can also change 
OldMessageID int64 `json:"old_message_id"` // The previous temporary message identifier
}

// MessageType return the string telegram-type of UpdateMessageSendSucceeded 
func (updateMessageSendSucceeded *UpdateMessageSendSucceeded) MessageType() string {
 return "updateMessageSendSucceeded" }

// UpdateMessageSendFailed A message failed to send. Be aware that some messages being sent can be irrecoverably deleted, in which case updateDeleteMessages will be received instead of this update 
type UpdateMessageSendFailed struct {
tdCommon
ErrorCode int32 `json:"error_code"` // An error code 
ErrorMessage string `json:"error_message"` // Error message
Message Message `json:"message"` // Contains information about the message that failed to send 
OldMessageID int64 `json:"old_message_id"` // The previous temporary message identifier 
}

// MessageType return the string telegram-type of UpdateMessageSendFailed 
func (updateMessageSendFailed *UpdateMessageSendFailed) MessageType() string {
 return "updateMessageSendFailed" }

// UpdateMessageContent The message content has changed  
type UpdateMessageContent struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier 
MessageID int64 `json:"message_id"` // Message identifier 
NewContent MessageContent `json:"new_content"` // New message content
}

// MessageType return the string telegram-type of UpdateMessageContent 
func (updateMessageContent *UpdateMessageContent) MessageType() string {
 return "updateMessageContent" }

// UpdateMessageEdited A message was edited. Changes in the message content will come in a separate updateMessageContent  
type UpdateMessageEdited struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier 
MessageID int64 `json:"message_id"` // Message identifier 
EditDate int32 `json:"edit_date"` // Point in time (Unix timestamp) when the message was edited 
ReplyMarkup ReplyMarkup `json:"reply_markup"` // New message reply markup; may be null
}

// MessageType return the string telegram-type of UpdateMessageEdited 
func (updateMessageEdited *UpdateMessageEdited) MessageType() string {
 return "updateMessageEdited" }

// UpdateMessageViews The view count of the message has changed  
type UpdateMessageViews struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier 
MessageID int64 `json:"message_id"` // Message identifier 
Views int32 `json:"views"` // New value of the view count
}

// MessageType return the string telegram-type of UpdateMessageViews 
func (updateMessageViews *UpdateMessageViews) MessageType() string {
 return "updateMessageViews" }

// UpdateMessageContentOpened The message content was opened. Updates voice note messages to "listened", video note messages to "viewed" and starts the TTL timer for self-destructing messages  
type UpdateMessageContentOpened struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier 
MessageID int64 `json:"message_id"` // Message identifier
}

// MessageType return the string telegram-type of UpdateMessageContentOpened 
func (updateMessageContentOpened *UpdateMessageContentOpened) MessageType() string {
 return "updateMessageContentOpened" }

// UpdateMessageMentionRead A message with an unread mention was read  
type UpdateMessageMentionRead struct {
tdCommon
UnreadMentionCount int32 `json:"unread_mention_count"` // The new number of unread mention messages left in the chat
ChatID int64 `json:"chat_id"` // Chat identifier 
MessageID int64 `json:"message_id"` // Message identifier 
}

// MessageType return the string telegram-type of UpdateMessageMentionRead 
func (updateMessageMentionRead *UpdateMessageMentionRead) MessageType() string {
 return "updateMessageMentionRead" }

// UpdateNewChat A new chat has been loaded/created. This update is guaranteed to come before the chat identifier is returned to the client. The chat field changes will be reported through separate updates  
type UpdateNewChat struct {
tdCommon
Chat Chat `json:"chat"` // The chat
}

// MessageType return the string telegram-type of UpdateNewChat 
func (updateNewChat *UpdateNewChat) MessageType() string {
 return "updateNewChat" }

// UpdateChatTitle The title of a chat was changed  
type UpdateChatTitle struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier 
Title string `json:"title"` // The new chat title
}

// MessageType return the string telegram-type of UpdateChatTitle 
func (updateChatTitle *UpdateChatTitle) MessageType() string {
 return "updateChatTitle" }

// UpdateChatPhoto A chat photo was changed  
type UpdateChatPhoto struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier 
Photo ChatPhoto `json:"photo"` // The new chat photo; may be null
}

// MessageType return the string telegram-type of UpdateChatPhoto 
func (updateChatPhoto *UpdateChatPhoto) MessageType() string {
 return "updateChatPhoto" }

// UpdateChatLastMessage The last message of a chat was changed. If last_message is null then the last message in the chat became unknown. Some new unknown messages might be added to the chat in this case  
type UpdateChatLastMessage struct {
tdCommon
Order int64 `json:"order"` // New value of the chat order
ChatID int64 `json:"chat_id"` // Chat identifier 
LastMessage Message `json:"last_message"` // The new last message in the chat; may be null 
}

// MessageType return the string telegram-type of UpdateChatLastMessage 
func (updateChatLastMessage *UpdateChatLastMessage) MessageType() string {
 return "updateChatLastMessage" }

// UpdateChatOrder The order of the chat in the chats list has changed. Instead of this update updateChatLastMessage, updateChatIsPinned or updateChatDraftMessage might be sent  
type UpdateChatOrder struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier 
Order int64 `json:"order"` // New value of the order
}

// MessageType return the string telegram-type of UpdateChatOrder 
func (updateChatOrder *UpdateChatOrder) MessageType() string {
 return "updateChatOrder" }

// UpdateChatIsPinned A chat was pinned or unpinned  
type UpdateChatIsPinned struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier 
IsPinned bool `json:"is_pinned"` // New value of is_pinned 
Order int64 `json:"order"` // New value of the chat order
}

// MessageType return the string telegram-type of UpdateChatIsPinned 
func (updateChatIsPinned *UpdateChatIsPinned) MessageType() string {
 return "updateChatIsPinned" }

// UpdateChatReadInbox Incoming messages were read or number of unread messages has been changed  
type UpdateChatReadInbox struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier 
LastReadInboxMessageID int64 `json:"last_read_inbox_message_id"` // Identifier of the last read incoming message 
UnreadCount int32 `json:"unread_count"` // The number of unread messages left in the chat
}

// MessageType return the string telegram-type of UpdateChatReadInbox 
func (updateChatReadInbox *UpdateChatReadInbox) MessageType() string {
 return "updateChatReadInbox" }

// UpdateChatReadOutbox Outgoing messages were read  
type UpdateChatReadOutbox struct {
tdCommon
LastReadOutboxMessageID int64 `json:"last_read_outbox_message_id"` // Identifier of last read outgoing message
ChatID int64 `json:"chat_id"` // Chat identifier 
}

// MessageType return the string telegram-type of UpdateChatReadOutbox 
func (updateChatReadOutbox *UpdateChatReadOutbox) MessageType() string {
 return "updateChatReadOutbox" }

// UpdateChatUnreadMentionCount The chat unread_mention_count has changed  
type UpdateChatUnreadMentionCount struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier 
UnreadMentionCount int32 `json:"unread_mention_count"` // The number of unread mention messages left in the chat
}

// MessageType return the string telegram-type of UpdateChatUnreadMentionCount 
func (updateChatUnreadMentionCount *UpdateChatUnreadMentionCount) MessageType() string {
 return "updateChatUnreadMentionCount" }

// UpdateNotificationSettings Notification settings for some chats were updated  
type UpdateNotificationSettings struct {
tdCommon
Scope NotificationSettingsScope `json:"scope"` // Types of chats for which notification settings were updated 
NotificationSettings NotificationSettings `json:"notification_settings"` // The new notification settings
}

// MessageType return the string telegram-type of UpdateNotificationSettings 
func (updateNotificationSettings *UpdateNotificationSettings) MessageType() string {
 return "updateNotificationSettings" }

// UpdateChatReplyMarkup The default chat reply markup was changed. Can occur because new messages with reply markup were received or because an old reply markup was hidden by the user 
type UpdateChatReplyMarkup struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier 
ReplyMarkupMessageID int64 `json:"reply_markup_message_id"` // Identifier of the message from which reply markup needs to be used; 0 if there is no default custom reply markup in the chat
}

// MessageType return the string telegram-type of UpdateChatReplyMarkup 
func (updateChatReplyMarkup *UpdateChatReplyMarkup) MessageType() string {
 return "updateChatReplyMarkup" }

// UpdateChatDraftMessage A draft has changed. Be aware that the update may come in the currently opened chat but with old content of the draft. If the user has changed the content of the draft, this update shouldn't be applied  
type UpdateChatDraftMessage struct {
tdCommon
Order int64 `json:"order"` // New value of the chat order
ChatID int64 `json:"chat_id"` // Chat identifier 
DraftMessage DraftMessage `json:"draft_message"` // The new draft message; may be null 
}

// MessageType return the string telegram-type of UpdateChatDraftMessage 
func (updateChatDraftMessage *UpdateChatDraftMessage) MessageType() string {
 return "updateChatDraftMessage" }

// UpdateDeleteMessages Some messages were deleted  
type UpdateDeleteMessages struct {
tdCommon
ChatID int64 `json:"chat_id"` // Chat identifier 
MessageIDs []int64 `json:"message_ids"` // Identifiers of the deleted messages
IsPermanent bool `json:"is_permanent"` // True, if the messages are permanently deleted by a user (as opposed to just becoming unaccessible)
FromCache bool `json:"from_cache"` // True, if the messages are deleted only from the cache and can possibly be retrieved again in the future
}

// MessageType return the string telegram-type of UpdateDeleteMessages 
func (updateDeleteMessages *UpdateDeleteMessages) MessageType() string {
 return "updateDeleteMessages" }

// UpdateUserChatAction User activity in the chat has changed  
type UpdateUserChatAction struct {
tdCommon
Action ChatAction `json:"action"` // The action description
ChatID int64 `json:"chat_id"` // Chat identifier 
UserID int32 `json:"user_id"` // Identifier of a user performing an action 
}

// MessageType return the string telegram-type of UpdateUserChatAction 
func (updateUserChatAction *UpdateUserChatAction) MessageType() string {
 return "updateUserChatAction" }

// UpdateUserStatus The user went online or offline  
type UpdateUserStatus struct {
tdCommon
UserID int32 `json:"user_id"` // User identifier 
Status UserStatus `json:"status"` // New status of the user
}

// MessageType return the string telegram-type of UpdateUserStatus 
func (updateUserStatus *UpdateUserStatus) MessageType() string {
 return "updateUserStatus" }

// UpdateUser Some data of a user has changed. This update is guaranteed to come before the user identifier is returned to the client  
type UpdateUser struct {
tdCommon
User User `json:"user"` // New data about the user
}

// MessageType return the string telegram-type of UpdateUser 
func (updateUser *UpdateUser) MessageType() string {
 return "updateUser" }

// UpdateBasicGroup Some data of a basic group has changed. This update is guaranteed to come before the basic group identifier is returned to the client  
type UpdateBasicGroup struct {
tdCommon
BasicGroup BasicGroup `json:"basic_group"` // New data about the group
}

// MessageType return the string telegram-type of UpdateBasicGroup 
func (updateBasicGroup *UpdateBasicGroup) MessageType() string {
 return "updateBasicGroup" }

// UpdateSupergroup Some data of a supergroup or a channel has changed. This update is guaranteed to come before the supergroup identifier is returned to the client  
type UpdateSupergroup struct {
tdCommon
Supergroup Supergroup `json:"supergroup"` // New data about the supergroup
}

// MessageType return the string telegram-type of UpdateSupergroup 
func (updateSupergroup *UpdateSupergroup) MessageType() string {
 return "updateSupergroup" }

// UpdateSecretChat Some data of a secret chat has changed. This update is guaranteed to come before the secret chat identifier is returned to the client  
type UpdateSecretChat struct {
tdCommon
SecretChat SecretChat `json:"secret_chat"` // New data about the secret chat
}

// MessageType return the string telegram-type of UpdateSecretChat 
func (updateSecretChat *UpdateSecretChat) MessageType() string {
 return "updateSecretChat" }

// UpdateUserFullInfo Some data from userFullInfo has been changed  
type UpdateUserFullInfo struct {
tdCommon
UserID int32 `json:"user_id"` // User identifier 
UserFullInfo UserFullInfo `json:"user_full_info"` // New full information about the user
}

// MessageType return the string telegram-type of UpdateUserFullInfo 
func (updateUserFullInfo *UpdateUserFullInfo) MessageType() string {
 return "updateUserFullInfo" }

// UpdateBasicGroupFullInfo Some data from basicGroupFullInfo has been changed  
type UpdateBasicGroupFullInfo struct {
tdCommon
BasicGroupID int32 `json:"basic_group_id"` // Identifier of a basic group 
BasicGroupFullInfo BasicGroupFullInfo `json:"basic_group_full_info"` // New full information about the group
}

// MessageType return the string telegram-type of UpdateBasicGroupFullInfo 
func (updateBasicGroupFullInfo *UpdateBasicGroupFullInfo) MessageType() string {
 return "updateBasicGroupFullInfo" }

// UpdateSupergroupFullInfo Some data from supergroupFullInfo has been changed  
type UpdateSupergroupFullInfo struct {
tdCommon
SupergroupID int32 `json:"supergroup_id"` // Identifier of the supergroup or channel 
SupergroupFullInfo SupergroupFullInfo `json:"supergroup_full_info"` // New full information about the supergroup
}

// MessageType return the string telegram-type of UpdateSupergroupFullInfo 
func (updateSupergroupFullInfo *UpdateSupergroupFullInfo) MessageType() string {
 return "updateSupergroupFullInfo" }

// UpdateServiceNotification Service notification from the server. Upon receiving this the client must show a popup with the content of the notification  
type UpdateServiceNotification struct {
tdCommon
Type string `json:"type"` // Notification type 
Content MessageContent `json:"content"` // Notification content
}

// MessageType return the string telegram-type of UpdateServiceNotification 
func (updateServiceNotification *UpdateServiceNotification) MessageType() string {
 return "updateServiceNotification" }

// UpdateFile Information about a file was updated  
type UpdateFile struct {
tdCommon
File File `json:"file"` // New data about the file
}

// MessageType return the string telegram-type of UpdateFile 
func (updateFile *UpdateFile) MessageType() string {
 return "updateFile" }

// UpdateFileGenerationStart The file generation process needs to be started by the client 
type UpdateFileGenerationStart struct {
tdCommon
GenerationID int64 `json:"generation_id"` // Unique identifier for the generation process
OriginalPath string `json:"original_path"` // The path to a file from which a new file is generated, may be empty
DestinationPath string `json:"destination_path"` // The path to a file that should be created and where the new file should be generated
Conversion string `json:"conversion"` // String specifying the conversion applied to the original file. If conversion is "#url#" than original_path contains a HTTP/HTTPS URL of a file, which should be downloaded by the client
}

// MessageType return the string telegram-type of UpdateFileGenerationStart 
func (updateFileGenerationStart *UpdateFileGenerationStart) MessageType() string {
 return "updateFileGenerationStart" }

// UpdateFileGenerationStop File generation is no longer needed  
type UpdateFileGenerationStop struct {
tdCommon
GenerationID int64 `json:"generation_id"` // Unique identifier for the generation process
}

// MessageType return the string telegram-type of UpdateFileGenerationStop 
func (updateFileGenerationStop *UpdateFileGenerationStop) MessageType() string {
 return "updateFileGenerationStop" }

// UpdateCall New call was created or information about a call was updated  
type UpdateCall struct {
tdCommon
Call Call `json:"call"` // New data about a call
}

// MessageType return the string telegram-type of UpdateCall 
func (updateCall *UpdateCall) MessageType() string {
 return "updateCall" }

// UpdateUserPrivacySettingRules Some privacy setting rules have been changed  
type UpdateUserPrivacySettingRules struct {
tdCommon
Setting UserPrivacySetting `json:"setting"` // The privacy setting 
Rules UserPrivacySettingRules `json:"rules"` // New privacy rules
}

// MessageType return the string telegram-type of UpdateUserPrivacySettingRules 
func (updateUserPrivacySettingRules *UpdateUserPrivacySettingRules) MessageType() string {
 return "updateUserPrivacySettingRules" }

// UpdateUnreadMessageCount Number of unread messages has changed. This update is sent only if a message database is used  
type UpdateUnreadMessageCount struct {
tdCommon
UnreadUnmutedCount int32 `json:"unread_unmuted_count"` // Total number of unread messages in unmuted chats
UnreadCount int32 `json:"unread_count"` // Total number of unread messages 
}

// MessageType return the string telegram-type of UpdateUnreadMessageCount 
func (updateUnreadMessageCount *UpdateUnreadMessageCount) MessageType() string {
 return "updateUnreadMessageCount" }

// UpdateOption An option changed its value  
type UpdateOption struct {
tdCommon
Value OptionValue `json:"value"` // The new option value
Name string `json:"name"` // The option name 
}

// MessageType return the string telegram-type of UpdateOption 
func (updateOption *UpdateOption) MessageType() string {
 return "updateOption" }

// UpdateInstalledStickerSets The list of installed sticker sets was updated  
type UpdateInstalledStickerSets struct {
tdCommon
IsMasks bool `json:"is_masks"` // True, if the list of installed mask sticker sets was updated 
StickerSetIDs []int64 `json:"sticker_set_ids"` // The new list of installed ordinary sticker sets
}

// MessageType return the string telegram-type of UpdateInstalledStickerSets 
func (updateInstalledStickerSets *UpdateInstalledStickerSets) MessageType() string {
 return "updateInstalledStickerSets" }

// UpdateTrendingStickerSets The list of trending sticker sets was updated or some of them were viewed  
type UpdateTrendingStickerSets struct {
tdCommon
StickerSets StickerSets `json:"sticker_sets"` // The new list of trending sticker sets
}

// MessageType return the string telegram-type of UpdateTrendingStickerSets 
func (updateTrendingStickerSets *UpdateTrendingStickerSets) MessageType() string {
 return "updateTrendingStickerSets" }

// UpdateRecentStickers The list of recently used stickers was updated  
type UpdateRecentStickers struct {
tdCommon
IsAttached bool `json:"is_attached"` // True, if the list of stickers attached to photo or video files was updated, otherwise the list of sent stickers is updated 
StickerIDs []int32 `json:"sticker_ids"` // The new list of file identifiers of recently used stickers
}

// MessageType return the string telegram-type of UpdateRecentStickers 
func (updateRecentStickers *UpdateRecentStickers) MessageType() string {
 return "updateRecentStickers" }

// UpdateFavoriteStickers The list of favorite stickers was updated  
type UpdateFavoriteStickers struct {
tdCommon
StickerIDs []int32 `json:"sticker_ids"` // The new list of file identifiers of favorite stickers
}

// MessageType return the string telegram-type of UpdateFavoriteStickers 
func (updateFavoriteStickers *UpdateFavoriteStickers) MessageType() string {
 return "updateFavoriteStickers" }

// UpdateSavedAnimations The list of saved animations was updated  
type UpdateSavedAnimations struct {
tdCommon
AnimationIDs []int32 `json:"animation_ids"` // The new list of file identifiers of saved animations
}

// MessageType return the string telegram-type of UpdateSavedAnimations 
func (updateSavedAnimations *UpdateSavedAnimations) MessageType() string {
 return "updateSavedAnimations" }

// UpdateConnectionState The connection state has changed  
type UpdateConnectionState struct {
tdCommon
State ConnectionState `json:"state"` // The new connection state
}

// MessageType return the string telegram-type of UpdateConnectionState 
func (updateConnectionState *UpdateConnectionState) MessageType() string {
 return "updateConnectionState" }

// UpdateNewInlineQuery A new incoming inline query; for bots only  
type UpdateNewInlineQuery struct {
tdCommon
SenderUserID int32 `json:"sender_user_id"` // Identifier of the user who sent the query 
UserLocation Location `json:"user_location"` // User location, provided by the client; may be null 
Query string `json:"query"` // Text of the query 
Offset string `json:"offset"` // Offset of the first entry to return
ID int64 `json:"id"` // Unique query identifier 
}

// MessageType return the string telegram-type of UpdateNewInlineQuery 
func (updateNewInlineQuery *UpdateNewInlineQuery) MessageType() string {
 return "updateNewInlineQuery" }

// UpdateNewChosenInlineResult The user has chosen a result of an inline query; for bots only  
type UpdateNewChosenInlineResult struct {
tdCommon
Query string `json:"query"` // Text of the query 
ResultID string `json:"result_id"` // Identifier of the chosen result 
InlineMessageID string `json:"inline_message_id"` // Identifier of the sent inline message, if known
SenderUserID int32 `json:"sender_user_id"` // Identifier of the user who sent the query 
UserLocation Location `json:"user_location"` // User location, provided by the client; may be null 
}

// MessageType return the string telegram-type of UpdateNewChosenInlineResult 
func (updateNewChosenInlineResult *UpdateNewChosenInlineResult) MessageType() string {
 return "updateNewChosenInlineResult" }

// UpdateNewCallbackQuery A new incoming callback query; for bots only  
type UpdateNewCallbackQuery struct {
tdCommon
MessageID int64 `json:"message_id"` // Identifier of the message, from which the query originated 
ChatInstance int64 `json:"chat_instance"` // Identifier that uniquely corresponds to the chat to which the message was sent 
Payload CallbackQueryPayload `json:"payload"` // Query payload
ID int64 `json:"id"` // Unique query identifier 
SenderUserID int32 `json:"sender_user_id"` // Identifier of the user who sent the query 
ChatID int64 `json:"chat_id"` // Identifier of the chat, in which the query was sent
}

// MessageType return the string telegram-type of UpdateNewCallbackQuery 
func (updateNewCallbackQuery *UpdateNewCallbackQuery) MessageType() string {
 return "updateNewCallbackQuery" }

// UpdateNewInlineCallbackQuery A new incoming callback query from a message sent via a bot; for bots only  
type UpdateNewInlineCallbackQuery struct {
tdCommon
Payload CallbackQueryPayload `json:"payload"` // Query payload
ID int64 `json:"id"` // Unique query identifier 
SenderUserID int32 `json:"sender_user_id"` // Identifier of the user who sent the query 
InlineMessageID string `json:"inline_message_id"` // Identifier of the inline message, from which the query originated
ChatInstance int64 `json:"chat_instance"` // An identifier uniquely corresponding to the chat a message was sent to 
}

// MessageType return the string telegram-type of UpdateNewInlineCallbackQuery 
func (updateNewInlineCallbackQuery *UpdateNewInlineCallbackQuery) MessageType() string {
 return "updateNewInlineCallbackQuery" }

// UpdateNewShippingQuery A new incoming shipping query; for bots only. Only for invoices with flexible price  
type UpdateNewShippingQuery struct {
tdCommon
ID int64 `json:"id"` // Unique query identifier 
SenderUserID int32 `json:"sender_user_id"` // Identifier of the user who sent the query 
InvoicePayload string `json:"invoice_payload"` // Invoice payload 
ShippingAddress ShippingAddress `json:"shipping_address"` // User shipping address
}

// MessageType return the string telegram-type of UpdateNewShippingQuery 
func (updateNewShippingQuery *UpdateNewShippingQuery) MessageType() string {
 return "updateNewShippingQuery" }

// UpdateNewPreCheckoutQuery A new incoming pre-checkout query; for bots only. Contains full information about a checkout  
type UpdateNewPreCheckoutQuery struct {
tdCommon
ShippingOptionID string `json:"shipping_option_id"` // Identifier of a shipping option chosen by the user; may be empty if not applicable 
OrderInfo OrderInfo `json:"order_info"` // Information about the order; may be null
ID int64 `json:"id"` // Unique query identifier 
SenderUserID int32 `json:"sender_user_id"` // Identifier of the user who sent the query 
Currency string `json:"currency"` // Currency for the product price 
TotalAmount int64 `json:"total_amount"` // Total price for the product, in the minimal quantity of the currency
InvoicePayload []byte `json:"invoice_payload"` // Invoice payload 
}

// MessageType return the string telegram-type of UpdateNewPreCheckoutQuery 
func (updateNewPreCheckoutQuery *UpdateNewPreCheckoutQuery) MessageType() string {
 return "updateNewPreCheckoutQuery" }

// UpdateNewCustomEvent A new incoming event; for bots only  
type UpdateNewCustomEvent struct {
tdCommon
Event string `json:"event"` // A JSON-serialized event
}

// MessageType return the string telegram-type of UpdateNewCustomEvent 
func (updateNewCustomEvent *UpdateNewCustomEvent) MessageType() string {
 return "updateNewCustomEvent" }

// UpdateNewCustomQuery A new incoming query; for bots only  
type UpdateNewCustomQuery struct {
tdCommon
Data string `json:"data"` // JSON-serialized query data 
Timeout int32 `json:"timeout"` // Query timeout
ID int64 `json:"id"` // The query identifier 
}

// MessageType return the string telegram-type of UpdateNewCustomQuery 
func (updateNewCustomQuery *UpdateNewCustomQuery) MessageType() string {
 return "updateNewCustomQuery" }

// TestInt A simple object containing a number; for testing only  
type TestInt struct {
tdCommon
Value int32 `json:"value"` // Number
}

// MessageType return the string telegram-type of TestInt 
func (testInt *TestInt) MessageType() string {
 return "testInt" }

// TestString A simple object containing a string; for testing only  
type TestString struct {
tdCommon
Value string `json:"value"` // String
}

// MessageType return the string telegram-type of TestString 
func (testString *TestString) MessageType() string {
 return "testString" }

// TestBytes A simple object containing a sequence of bytes; for testing only  
type TestBytes struct {
tdCommon
Value []byte `json:"value"` // Bytes
}

// MessageType return the string telegram-type of TestBytes 
func (testBytes *TestBytes) MessageType() string {
 return "testBytes" }

// TestVectorInt A simple object containing a vector of numbers; for testing only  
type TestVectorInt struct {
tdCommon
Value []int32 `json:"value"` // Vector of numbers
}

// MessageType return the string telegram-type of TestVectorInt 
func (testVectorInt *TestVectorInt) MessageType() string {
 return "testVectorInt" }

// TestVectorIntObject A simple object containing a vector of objects that hold a number; for testing only  
type TestVectorIntObject struct {
tdCommon
Value []TestInt `json:"value"` // Vector of objects
}

// MessageType return the string telegram-type of TestVectorIntObject 
func (testVectorIntObject *TestVectorIntObject) MessageType() string {
 return "testVectorIntObject" }

// TestVectorString A simple object containing a vector of strings; for testing only  
type TestVectorString struct {
tdCommon
Value []string `json:"value"` // Vector of strings
}

// MessageType return the string telegram-type of TestVectorString 
func (testVectorString *TestVectorString) MessageType() string {
 return "testVectorString" }

// TestVectorStringObject A simple object containing a vector of objects that hold a string; for testing only  
type TestVectorStringObject struct {
tdCommon
Value []TestString `json:"value"` // Vector of objects
}

// MessageType return the string telegram-type of TestVectorStringObject 
func (testVectorStringObject *TestVectorStringObject) MessageType() string {
 return "testVectorStringObject" }

fication Service notification from the server. Upon receiving this the client must show a popup with the content of the notification
type UpdateServiceNotification struct {
	tdCommon
	Type    string         `json:"type"`    // Notification type
	Content MessageContent `json:"content"` // Notification content
}

// MessageType return the string telegram-type of UpdateServiceNotification
func (updateServiceNotification *UpdateServiceNotification) MessageType() string {
	return "updateServiceNotification"
}

// UpdateFile Information about a file was updated
type UpdateFile struct {
	tdCommon
	File File `json:"file"` // New data about the file
}

// MessageType return the string telegram-type of UpdateFile
func (updateFile *UpdateFile) MessageType() string {
	return "updateFile"
}

// UpdateFileGenerationStart The file generation process needs to be started by the client
type UpdateFileGenerationStart struct {
	tdCommon
	Conversion      string `json:"conversion"`       // String specifying the conversion applied to the original file. If conversion is "#url#" than original_path contains a HTTP/HTTPS URL of a file, which should be downloaded by the client
	GenerationID    int64  `json:"generation_id"`    // Unique identifier for the generation process
	OriginalPath    string `json:"original_path"`    // The path to a file from which a new file is generated, may be empty
	DestinationPath string `json:"destination_path"` // The path to a file that should be created and where the new file should be generated
}

// MessageType return the string telegram-type of UpdateFileGenerationStart
func (updateFileGenerationStart *UpdateFileGenerationStart) MessageType() string {
	return "updateFileGenerationStart"
}

// UpdateFileGenerationStop File generation is no longer needed
type UpdateFileGenerationStop struct {
	tdCommon
	GenerationID int64 `json:"generation_id"` // Unique identifier for the generation process
}

// MessageType return the string telegram-type of UpdateFileGenerationStop
func (updateFileGenerationStop *UpdateFileGenerationStop) MessageType() string {
	return "updateFileGenerationStop"
}

// UpdateCall New call was created or information about a call was updated
type UpdateCall struct {
	tdCommon
	Call Call `json:"call"` // New data about a call
}

// MessageType return the string telegram-type of UpdateCall
func (updateCall *UpdateCall) MessageType() string {
	return "updateCall"
}

// UpdateUserPrivacySettingRules Some privacy setting rules have been changed
type UpdateUserPrivacySettingRules struct {
	tdCommon
	Setting UserPrivacySetting      `json:"setting"` // The privacy setting
	Rules   UserPrivacySettingRules `json:"rules"`   // New privacy rules
}

// MessageType return the string telegram-type of UpdateUserPrivacySettingRules
func (updateUserPrivacySettingRules *UpdateUserPrivacySettingRules) MessageType() string {
	return "updateUserPrivacySettingRules"
}

// UpdateUnreadMessageCount Number of unread messages has changed. This update is sent only if a message database is used
type UpdateUnreadMessageCount struct {
	tdCommon
	UnreadCount        int32 `json:"unread_count"`         // Total number of unread messages
	UnreadUnmutedCount int32 `json:"unread_unmuted_count"` // Total number of unread messages in unmuted chats
}

// MessageType return the string telegram-type of UpdateUnreadMessageCount
func (updateUnreadMessageCount *UpdateUnreadMessageCount) MessageType() string {
	return "updateUnreadMessageCount"
}

// UpdateOption An option changed its value
type UpdateOption struct {
	tdCommon
	Name  string      `json:"name"`  // The option name
	Value OptionValue `json:"value"` // The new option value
}

// MessageType return the string telegram-type of UpdateOption
func (updateOption *UpdateOption) MessageType() string {
	return "updateOption"
}

// UpdateInstalledStickerSets The list of installed sticker sets was updated
type UpdateInstalledStickerSets struct {
	tdCommon
	IsMasks       bool    `json:"is_masks"`        // True, if the list of installed mask sticker sets was updated
	StickerSetIDs []int64 `json:"sticker_set_ids"` // The new list of installed ordinary sticker sets
}

// MessageType return the string telegram-type of UpdateInstalledStickerSets
func (updateInstalledStickerSets *UpdateInstalledStickerSets) MessageType() string {
	return "updateInstalledStickerSets"
}

// UpdateTrendingStickerSets The list of trending sticker sets was updated or some of them were viewed
type UpdateTrendingStickerSets struct {
	tdCommon
	StickerSets StickerSets `json:"sticker_sets"` // The new list of trending sticker sets
}

// MessageType return the string telegram-type of UpdateTrendingStickerSets
func (updateTrendingStickerSets *UpdateTrendingStickerSets) MessageType() string {
	return "updateTrendingStickerSets"
}

// UpdateRecentStickers The list of recently used stickers was updated
type UpdateRecentStickers struct {
	tdCommon
	IsAttached bool    `json:"is_attached"` // True, if the list of stickers attached to photo or video files was updated, otherwise the list of sent stickers is updated
	StickerIDs []int32 `json:"sticker_ids"` // The new list of file identifiers of recently used stickers
}

// MessageType return the string telegram-type of UpdateRecentStickers
func (updateRecentStickers *UpdateRecentStickers) MessageType() string {
	return "updateRecentStickers"
}

// UpdateFavoriteStickers The list of favorite stickers was updated
type UpdateFavoriteStickers struct {
	tdCommon
	StickerIDs []int32 `json:"sticker_ids"` // The new list of file identifiers of favorite stickers
}

// MessageType return the string telegram-type of UpdateFavoriteStickers
func (updateFavoriteStickers *UpdateFavoriteStickers) MessageType() string {
	return "updateFavoriteStickers"
}

// UpdateSavedAnimations The list of saved animations was updated
type UpdateSavedAnimations struct {
	tdCommon
	AnimationIDs []int32 `json:"animation_ids"` // The new list of file identifiers of saved animations
}

// MessageType return the string telegram-type of UpdateSavedAnimations
func (updateSavedAnimations *UpdateSavedAnimations) MessageType() string {
	return "updateSavedAnimations"
}

// UpdateConnectionState The connection state has changed
type UpdateConnectionState struct {
	tdCommon
	State ConnectionState `json:"state"` // The new connection state
}

// MessageType return the string telegram-type of UpdateConnectionState
func (updateConnectionState *UpdateConnectionState) MessageType() string {
	return "updateConnectionState"
}

// UpdateNewInlineQuery A new incoming inline query; for bots only
type UpdateNewInlineQuery struct {
	tdCommon
	ID           int64    `json:"id"`             // Unique query identifier
	SenderUserID int32    `json:"sender_user_id"` // Identifier of the user who sent the query
	UserLocation Location `json:"user_location"`  // User location, provided by the client; may be null
	Query        string   `json:"query"`          // Text of the query
	Offset       string   `json:"offset"`         // Offset of the first entry to return
}

// MessageType return the string telegram-type of UpdateNewInlineQuery
func (updateNewInlineQuery *UpdateNewInlineQuery) MessageType() string {
	return "updateNewInlineQuery"
}

// UpdateNewChosenInlineResult The user has chosen a result of an inline query; for bots only
type UpdateNewChosenInlineResult struct {
	tdCommon
	SenderUserID    int32    `json:"sender_user_id"`    // Identifier of the user who sent the query
	UserLocation    Location `json:"user_location"`     // User location, provided by the client; may be null
	Query           string   `json:"query"`             // Text of the query
	ResultID        string   `json:"result_id"`         // Identifier of the chosen result
	InlineMessageID string   `json:"inline_message_id"` // Identifier of the sent inline message, if known
}

// MessageType return the string telegram-type of UpdateNewChosenInlineResult
func (updateNewChosenInlineResult *UpdateNewChosenInlineResult) MessageType() string {
	return "updateNewChosenInlineResult"
}

// UpdateNewCallbackQuery A new incoming callback query; for bots only
type UpdateNewCallbackQuery struct {
	tdCommon
	Payload      CallbackQueryPayload `json:"payload"`        // Query payload
	ID           int64                `json:"id"`             // Unique query identifier
	SenderUserID int32                `json:"sender_user_id"` // Identifier of the user who sent the query
	ChatID       int64                `json:"chat_id"`        // Identifier of the chat, in which the query was sent
	MessageID    int64                `json:"message_id"`     // Identifier of the message, from which the query originated
	ChatInstance int64                `json:"chat_instance"`  // Identifier that uniquely corresponds to the chat to which the message was sent
}

// MessageType return the string telegram-type of UpdateNewCallbackQuery
func (updateNewCallbackQuery *UpdateNewCallbackQuery) MessageType() string {
	return "updateNewCallbackQuery"
}

// UpdateNewInlineCallbackQuery A new incoming callback query from a message sent via a bot; for bots only
type UpdateNewInlineCallbackQuery struct {
	tdCommon
	ChatInstance    int64                `json:"chat_instance"`     // An identifier uniquely corresponding to the chat a message was sent to
	Payload         CallbackQueryPayload `json:"payload"`           // Query payload
	ID              int64                `json:"id"`                // Unique query identifier
	SenderUserID    int32                `json:"sender_user_id"`    // Identifier of the user who sent the query
	InlineMessageID string               `json:"inline_message_id"` // Identifier of the inline message, from which the query originated
}

// MessageType return the string telegram-type of UpdateNewInlineCallbackQuery
func (updateNewInlineCallbackQuery *UpdateNewInlineCallbackQuery) MessageType() string {
	return "updateNewInlineCallbackQuery"
}

// UpdateNewShippingQuery A new incoming shipping query; for bots only. Only for invoices with flexible price
type UpdateNewShippingQuery struct {
	tdCommon
	InvoicePayload  string          `json:"invoice_payload"`  // Invoice payload
	ShippingAddress ShippingAddress `json:"shipping_address"` // User shipping address
	ID              int64           `json:"id"`               // Unique query identifier
	SenderUserID    int32           `json:"sender_user_id"`   // Identifier of the user who sent the query
}

// MessageType return the string telegram-type of UpdateNewShippingQuery
func (updateNewShippingQuery *UpdateNewShippingQuery) MessageType() string {
	return "updateNewShippingQuery"
}

// UpdateNewPreCheckoutQuery A new incoming pre-checkout query; for bots only. Contains full information about a checkout
type UpdateNewPreCheckoutQuery struct {
	tdCommon
	Currency         string    `json:"currency"`           // Currency for the product price
	TotalAmount      int64     `json:"total_amount"`       // Total price for the product, in the minimal quantity of the currency
	InvoicePayload   []byte    `json:"invoice_payload"`    // Invoice payload
	ShippingOptionID string    `json:"shipping_option_id"` // Identifier of a shipping option chosen by the user; may be empty if not applicable
	OrderInfo        OrderInfo `json:"order_info"`         // Information about the order; may be null
	ID               int64     `json:"id"`                 // Unique query identifier
	SenderUserID     int32     `json:"sender_user_id"`     // Identifier of the user who sent the query
}

// MessageType return the string telegram-type of UpdateNewPreCheckoutQuery
func (updateNewPreCheckoutQuery *UpdateNewPreCheckoutQuery) MessageType() string {
	return "updateNewPreCheckoutQuery"
}

// UpdateNewCustomEvent A new incoming event; for bots only
type UpdateNewCustomEvent struct {
	tdCommon
	Event string `json:"event"` // A JSON-serialized event
}

// MessageType return the string telegram-type of UpdateNewCustomEvent
func (updateNewCustomEvent *UpdateNewCustomEvent) MessageType() string {
	return "updateNewCustomEvent"
}

// UpdateNewCustomQuery A new incoming query; for bots only
type UpdateNewCustomQuery struct {
	tdCommon
	Data    string `json:"data"`    // JSON-serialized query data
	Timeout int32  `json:"timeout"` // Query timeout
	ID      int64  `json:"id"`      // The query identifier
}

// MessageType return the string telegram-type of UpdateNewCustomQuery
func (updateNewCustomQuery *UpdateNewCustomQuery) MessageType() string {
	return "updateNewCustomQuery"
}

// TestInt A simple object containing a number; for testing only
type TestInt struct {
	tdCommon
	Value int32 `json:"value"` // Number
}

// MessageType return the string telegram-type of TestInt
func (testInt *TestInt) MessageType() string {
	return "testInt"
}

// TestString A simple object containing a string; for testing only
type TestString struct {
	tdCommon
	Value string `json:"value"` // String
}

// MessageType return the string telegram-type of TestString
func (testString *TestString) MessageType() string {
	return "testString"
}

// TestBytes A simple object containing a sequence of bytes; for testing only
type TestBytes struct {
	tdCommon
	Value []byte `json:"value"` // Bytes
}

// MessageType return the string telegram-type of TestBytes
func (testBytes *TestBytes) MessageType() string {
	return "testBytes"
}

// TestVectorInt A simple object containing a vector of numbers; for testing only
type TestVectorInt struct {
	tdCommon
	Value []int32 `json:"value"` // Vector of numbers
}

// MessageType return the string telegram-type of TestVectorInt
func (testVectorInt *TestVectorInt) MessageType() string {
	return "testVectorInt"
}

// TestVectorIntObject A simple object containing a vector of objects that hold a number; for testing only
type TestVectorIntObject struct {
	tdCommon
	Value []TestInt `json:"value"` // Vector of objects
}

// MessageType return the string telegram-type of TestVectorIntObject
func (testVectorIntObject *TestVectorIntObject) MessageType() string {
	return "testVectorIntObject"
}

// TestVectorString A simple object containing a vector of strings; for testing only
type TestVectorString struct {
	tdCommon
	Value []string `json:"value"` // Vector of strings
}

// MessageType return the string telegram-type of TestVectorString
func (testVectorString *TestVectorString) MessageType() string {
	return "testVectorString"
}

// TestVectorStringObject A simple object containing a vector of objects that hold a string; for testing only
type TestVectorStringObject struct {
	tdCommon
	Value []TestString `json:"value"` // Vector of objects
}

// MessageType return the string telegram-type of TestVectorStringObject
func (testVectorStringObject *TestVectorStringObject) MessageType() string {
	return "testVectorStringObject"
}
