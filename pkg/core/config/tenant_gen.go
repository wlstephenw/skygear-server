package config

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *AuthConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zbai uint32
	zbai, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zbai > 0 {
		zbai--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "AUTH_RECORD_KEYS":
			var zcmr uint32
			zcmr, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.AuthRecordKeys) >= int(zcmr) {
				z.AuthRecordKeys = (z.AuthRecordKeys)[:zcmr]
			} else {
				z.AuthRecordKeys = make([][]string, zcmr)
			}
			for zxvk := range z.AuthRecordKeys {
				var zajw uint32
				zajw, err = dc.ReadArrayHeader()
				if err != nil {
					return
				}
				if cap(z.AuthRecordKeys[zxvk]) >= int(zajw) {
					z.AuthRecordKeys[zxvk] = (z.AuthRecordKeys[zxvk])[:zajw]
				} else {
					z.AuthRecordKeys[zxvk] = make([]string, zajw)
				}
				for zbzg := range z.AuthRecordKeys[zxvk] {
					z.AuthRecordKeys[zxvk][zbzg], err = dc.ReadString()
					if err != nil {
						return
					}
				}
			}
		case "CUSTOM_TOKEN_SECRET":
			z.CustomTokenSecret, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *AuthConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "AUTH_RECORD_KEYS"
	err = en.Append(0x82, 0xb0, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4b, 0x45, 0x59, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.AuthRecordKeys)))
	if err != nil {
		return
	}
	for zxvk := range z.AuthRecordKeys {
		err = en.WriteArrayHeader(uint32(len(z.AuthRecordKeys[zxvk])))
		if err != nil {
			return
		}
		for zbzg := range z.AuthRecordKeys[zxvk] {
			err = en.WriteString(z.AuthRecordKeys[zxvk][zbzg])
			if err != nil {
				return
			}
		}
	}
	// write "CUSTOM_TOKEN_SECRET"
	err = en.Append(0xb3, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.CustomTokenSecret)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AuthConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "AUTH_RECORD_KEYS"
	o = append(o, 0x82, 0xb0, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4b, 0x45, 0x59, 0x53)
	o = msgp.AppendArrayHeader(o, uint32(len(z.AuthRecordKeys)))
	for zxvk := range z.AuthRecordKeys {
		o = msgp.AppendArrayHeader(o, uint32(len(z.AuthRecordKeys[zxvk])))
		for zbzg := range z.AuthRecordKeys[zxvk] {
			o = msgp.AppendString(o, z.AuthRecordKeys[zxvk][zbzg])
		}
	}
	// string "CUSTOM_TOKEN_SECRET"
	o = append(o, 0xb3, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	o = msgp.AppendString(o, z.CustomTokenSecret)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AuthConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zwht uint32
	zwht, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zwht > 0 {
		zwht--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "AUTH_RECORD_KEYS":
			var zhct uint32
			zhct, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.AuthRecordKeys) >= int(zhct) {
				z.AuthRecordKeys = (z.AuthRecordKeys)[:zhct]
			} else {
				z.AuthRecordKeys = make([][]string, zhct)
			}
			for zxvk := range z.AuthRecordKeys {
				var zcua uint32
				zcua, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(z.AuthRecordKeys[zxvk]) >= int(zcua) {
					z.AuthRecordKeys[zxvk] = (z.AuthRecordKeys[zxvk])[:zcua]
				} else {
					z.AuthRecordKeys[zxvk] = make([]string, zcua)
				}
				for zbzg := range z.AuthRecordKeys[zxvk] {
					z.AuthRecordKeys[zxvk][zbzg], bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				}
			}
		case "CUSTOM_TOKEN_SECRET":
			z.CustomTokenSecret, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AuthConfiguration) Msgsize() (s int) {
	s = 1 + 17 + msgp.ArrayHeaderSize
	for zxvk := range z.AuthRecordKeys {
		s += msgp.ArrayHeaderSize
		for zbzg := range z.AuthRecordKeys[zxvk] {
			s += msgp.StringPrefixSize + len(z.AuthRecordKeys[zxvk][zbzg])
		}
	}
	s += 20 + msgp.StringPrefixSize + len(z.CustomTokenSecret)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ForgotPasswordConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxhx uint32
	zxhx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxhx > 0 {
		zxhx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "APP_NAME":
			z.AppName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "URL_PREFIX":
			z.URLPrefix, err = dc.ReadString()
			if err != nil {
				return
			}
		case "SECURE_MATCH":
			z.SecureMatch, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "SENDER_NAME":
			z.SenderName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "SENDER":
			z.Sender, err = dc.ReadString()
			if err != nil {
				return
			}
		case "SUBJECT":
			z.Subject, err = dc.ReadString()
			if err != nil {
				return
			}
		case "REPLY_TO_NAME":
			z.ReplyToName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "REPLY_TO":
			z.ReplyTo, err = dc.ReadString()
			if err != nil {
				return
			}
		case "RESET_URL_LIFE_TIME":
			z.ResetURLLifeTime, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "SUCCESS_REDIRECT":
			z.SuccessRedirect, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ERROR_REDIRECT":
			z.ErrorRedirect, err = dc.ReadString()
			if err != nil {
				return
			}
		case "EMAIL_TEXT_URL":
			z.EmailTextURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "EMAIL_HTML_URL":
			z.EmailHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "RESET_HTML_URL":
			z.ResetHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "RESET_SUCCESS_HTML_URL":
			z.ResetSuccessHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "RESET_ERROR_HTML_URL":
			z.ResetErrorHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ForgotPasswordConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 16
	// write "APP_NAME"
	err = en.Append(0xde, 0x0, 0x10, 0xa8, 0x41, 0x50, 0x50, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteString(z.AppName)
	if err != nil {
		return
	}
	// write "URL_PREFIX"
	err = en.Append(0xaa, 0x55, 0x52, 0x4c, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58)
	if err != nil {
		return err
	}
	err = en.WriteString(z.URLPrefix)
	if err != nil {
		return
	}
	// write "SECURE_MATCH"
	err = en.Append(0xac, 0x53, 0x45, 0x43, 0x55, 0x52, 0x45, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.SecureMatch)
	if err != nil {
		return
	}
	// write "SENDER_NAME"
	err = en.Append(0xab, 0x53, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SenderName)
	if err != nil {
		return
	}
	// write "SENDER"
	err = en.Append(0xa6, 0x53, 0x45, 0x4e, 0x44, 0x45, 0x52)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Sender)
	if err != nil {
		return
	}
	// write "SUBJECT"
	err = en.Append(0xa7, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Subject)
	if err != nil {
		return
	}
	// write "REPLY_TO_NAME"
	err = en.Append(0xad, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ReplyToName)
	if err != nil {
		return
	}
	// write "REPLY_TO"
	err = en.Append(0xa8, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x54, 0x4f)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ReplyTo)
	if err != nil {
		return
	}
	// write "RESET_URL_LIFE_TIME"
	err = en.Append(0xb3, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x4c, 0x49, 0x46, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.ResetURLLifeTime)
	if err != nil {
		return
	}
	// write "SUCCESS_REDIRECT"
	err = en.Append(0xb0, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SuccessRedirect)
	if err != nil {
		return
	}
	// write "ERROR_REDIRECT"
	err = en.Append(0xae, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ErrorRedirect)
	if err != nil {
		return
	}
	// write "EMAIL_TEXT_URL"
	err = en.Append(0xae, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.EmailTextURL)
	if err != nil {
		return
	}
	// write "EMAIL_HTML_URL"
	err = en.Append(0xae, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.EmailHTMLURL)
	if err != nil {
		return
	}
	// write "RESET_HTML_URL"
	err = en.Append(0xae, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ResetHTMLURL)
	if err != nil {
		return
	}
	// write "RESET_SUCCESS_HTML_URL"
	err = en.Append(0xb6, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ResetSuccessHTMLURL)
	if err != nil {
		return
	}
	// write "RESET_ERROR_HTML_URL"
	err = en.Append(0xb4, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ResetErrorHTMLURL)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ForgotPasswordConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 16
	// string "APP_NAME"
	o = append(o, 0xde, 0x0, 0x10, 0xa8, 0x41, 0x50, 0x50, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	o = msgp.AppendString(o, z.AppName)
	// string "URL_PREFIX"
	o = append(o, 0xaa, 0x55, 0x52, 0x4c, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58)
	o = msgp.AppendString(o, z.URLPrefix)
	// string "SECURE_MATCH"
	o = append(o, 0xac, 0x53, 0x45, 0x43, 0x55, 0x52, 0x45, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48)
	o = msgp.AppendBool(o, z.SecureMatch)
	// string "SENDER_NAME"
	o = append(o, 0xab, 0x53, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	o = msgp.AppendString(o, z.SenderName)
	// string "SENDER"
	o = append(o, 0xa6, 0x53, 0x45, 0x4e, 0x44, 0x45, 0x52)
	o = msgp.AppendString(o, z.Sender)
	// string "SUBJECT"
	o = append(o, 0xa7, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54)
	o = msgp.AppendString(o, z.Subject)
	// string "REPLY_TO_NAME"
	o = append(o, 0xad, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	o = msgp.AppendString(o, z.ReplyToName)
	// string "REPLY_TO"
	o = append(o, 0xa8, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x54, 0x4f)
	o = msgp.AppendString(o, z.ReplyTo)
	// string "RESET_URL_LIFE_TIME"
	o = append(o, 0xb3, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x55, 0x52, 0x4c, 0x5f, 0x4c, 0x49, 0x46, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45)
	o = msgp.AppendInt(o, z.ResetURLLifeTime)
	// string "SUCCESS_REDIRECT"
	o = append(o, 0xb0, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54)
	o = msgp.AppendString(o, z.SuccessRedirect)
	// string "ERROR_REDIRECT"
	o = append(o, 0xae, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54)
	o = msgp.AppendString(o, z.ErrorRedirect)
	// string "EMAIL_TEXT_URL"
	o = append(o, 0xae, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.EmailTextURL)
	// string "EMAIL_HTML_URL"
	o = append(o, 0xae, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.EmailHTMLURL)
	// string "RESET_HTML_URL"
	o = append(o, 0xae, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.ResetHTMLURL)
	// string "RESET_SUCCESS_HTML_URL"
	o = append(o, 0xb6, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.ResetSuccessHTMLURL)
	// string "RESET_ERROR_HTML_URL"
	o = append(o, 0xb4, 0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.ResetErrorHTMLURL)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ForgotPasswordConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zlqf uint32
	zlqf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zlqf > 0 {
		zlqf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "APP_NAME":
			z.AppName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "URL_PREFIX":
			z.URLPrefix, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SECURE_MATCH":
			z.SecureMatch, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "SENDER_NAME":
			z.SenderName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SENDER":
			z.Sender, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SUBJECT":
			z.Subject, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "REPLY_TO_NAME":
			z.ReplyToName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "REPLY_TO":
			z.ReplyTo, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RESET_URL_LIFE_TIME":
			z.ResetURLLifeTime, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "SUCCESS_REDIRECT":
			z.SuccessRedirect, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ERROR_REDIRECT":
			z.ErrorRedirect, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "EMAIL_TEXT_URL":
			z.EmailTextURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "EMAIL_HTML_URL":
			z.EmailHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RESET_HTML_URL":
			z.ResetHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RESET_SUCCESS_HTML_URL":
			z.ResetSuccessHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RESET_ERROR_HTML_URL":
			z.ResetErrorHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ForgotPasswordConfiguration) Msgsize() (s int) {
	s = 3 + 9 + msgp.StringPrefixSize + len(z.AppName) + 11 + msgp.StringPrefixSize + len(z.URLPrefix) + 13 + msgp.BoolSize + 12 + msgp.StringPrefixSize + len(z.SenderName) + 7 + msgp.StringPrefixSize + len(z.Sender) + 8 + msgp.StringPrefixSize + len(z.Subject) + 14 + msgp.StringPrefixSize + len(z.ReplyToName) + 9 + msgp.StringPrefixSize + len(z.ReplyTo) + 20 + msgp.IntSize + 17 + msgp.StringPrefixSize + len(z.SuccessRedirect) + 15 + msgp.StringPrefixSize + len(z.ErrorRedirect) + 15 + msgp.StringPrefixSize + len(z.EmailTextURL) + 15 + msgp.StringPrefixSize + len(z.EmailHTMLURL) + 15 + msgp.StringPrefixSize + len(z.ResetHTMLURL) + 23 + msgp.StringPrefixSize + len(z.ResetSuccessHTMLURL) + 21 + msgp.StringPrefixSize + len(z.ResetErrorHTMLURL)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *NexmoConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zdaf uint32
	zdaf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zdaf > 0 {
		zdaf--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "API_KEY":
			z.APIKey, err = dc.ReadString()
			if err != nil {
				return
			}
		case "API_SECRET":
			z.APISecret, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FROM":
			z.From, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z NexmoConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "API_KEY"
	err = en.Append(0x83, 0xa7, 0x41, 0x50, 0x49, 0x5f, 0x4b, 0x45, 0x59)
	if err != nil {
		return err
	}
	err = en.WriteString(z.APIKey)
	if err != nil {
		return
	}
	// write "API_SECRET"
	err = en.Append(0xaa, 0x41, 0x50, 0x49, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.APISecret)
	if err != nil {
		return
	}
	// write "FROM"
	err = en.Append(0xa4, 0x46, 0x52, 0x4f, 0x4d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.From)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z NexmoConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "API_KEY"
	o = append(o, 0x83, 0xa7, 0x41, 0x50, 0x49, 0x5f, 0x4b, 0x45, 0x59)
	o = msgp.AppendString(o, z.APIKey)
	// string "API_SECRET"
	o = append(o, 0xaa, 0x41, 0x50, 0x49, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	o = msgp.AppendString(o, z.APISecret)
	// string "FROM"
	o = append(o, 0xa4, 0x46, 0x52, 0x4f, 0x4d)
	o = msgp.AppendString(o, z.From)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NexmoConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zpks uint32
	zpks, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zpks > 0 {
		zpks--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "API_KEY":
			z.APIKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "API_SECRET":
			z.APISecret, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FROM":
			z.From, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z NexmoConfiguration) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.APIKey) + 11 + msgp.StringPrefixSize + len(z.APISecret) + 5 + msgp.StringPrefixSize + len(z.From)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SMTPConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zjfb uint32
	zjfb, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zjfb > 0 {
		zjfb--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "HOST":
			z.Host, err = dc.ReadString()
			if err != nil {
				return
			}
		case "PORT":
			z.Port, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "MODE":
			z.Mode, err = dc.ReadString()
			if err != nil {
				return
			}
		case "LOGIN":
			z.Login, err = dc.ReadString()
			if err != nil {
				return
			}
		case "PASSWORD":
			z.Password, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SMTPConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "HOST"
	err = en.Append(0x85, 0xa4, 0x48, 0x4f, 0x53, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Host)
	if err != nil {
		return
	}
	// write "PORT"
	err = en.Append(0xa4, 0x50, 0x4f, 0x52, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Port)
	if err != nil {
		return
	}
	// write "MODE"
	err = en.Append(0xa4, 0x4d, 0x4f, 0x44, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Mode)
	if err != nil {
		return
	}
	// write "LOGIN"
	err = en.Append(0xa5, 0x4c, 0x4f, 0x47, 0x49, 0x4e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Login)
	if err != nil {
		return
	}
	// write "PASSWORD"
	err = en.Append(0xa8, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Password)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SMTPConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "HOST"
	o = append(o, 0x85, 0xa4, 0x48, 0x4f, 0x53, 0x54)
	o = msgp.AppendString(o, z.Host)
	// string "PORT"
	o = append(o, 0xa4, 0x50, 0x4f, 0x52, 0x54)
	o = msgp.AppendInt(o, z.Port)
	// string "MODE"
	o = append(o, 0xa4, 0x4d, 0x4f, 0x44, 0x45)
	o = msgp.AppendString(o, z.Mode)
	// string "LOGIN"
	o = append(o, 0xa5, 0x4c, 0x4f, 0x47, 0x49, 0x4e)
	o = msgp.AppendString(o, z.Login)
	// string "PASSWORD"
	o = append(o, 0xa8, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44)
	o = msgp.AppendString(o, z.Password)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SMTPConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcxo uint32
	zcxo, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcxo > 0 {
		zcxo--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "HOST":
			z.Host, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "PORT":
			z.Port, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "MODE":
			z.Mode, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "LOGIN":
			z.Login, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "PASSWORD":
			z.Password, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SMTPConfiguration) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Host) + 5 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Mode) + 6 + msgp.StringPrefixSize + len(z.Login) + 9 + msgp.StringPrefixSize + len(z.Password)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SSOConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zeff uint32
	zeff, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zeff > 0 {
		zeff--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "NAME":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "CLIENT_ID":
			z.ClientID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "CLIENT_SECRET":
			z.ClientSecret, err = dc.ReadString()
			if err != nil {
				return
			}
		case "SCOPE":
			z.Scope, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SSOConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "NAME"
	err = en.Append(0x84, 0xa4, 0x4e, 0x41, 0x4d, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "CLIENT_ID"
	err = en.Append(0xa9, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ClientID)
	if err != nil {
		return
	}
	// write "CLIENT_SECRET"
	err = en.Append(0xad, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ClientSecret)
	if err != nil {
		return
	}
	// write "SCOPE"
	err = en.Append(0xa5, 0x53, 0x43, 0x4f, 0x50, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Scope)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SSOConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "NAME"
	o = append(o, 0x84, 0xa4, 0x4e, 0x41, 0x4d, 0x45)
	o = msgp.AppendString(o, z.Name)
	// string "CLIENT_ID"
	o = append(o, 0xa9, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x44)
	o = msgp.AppendString(o, z.ClientID)
	// string "CLIENT_SECRET"
	o = append(o, 0xad, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	o = msgp.AppendString(o, z.ClientSecret)
	// string "SCOPE"
	o = append(o, 0xa5, 0x53, 0x43, 0x4f, 0x50, 0x45)
	o = msgp.AppendString(o, z.Scope)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SSOConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrsw uint32
	zrsw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrsw > 0 {
		zrsw--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "NAME":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "CLIENT_ID":
			z.ClientID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "CLIENT_SECRET":
			z.ClientSecret, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SCOPE":
			z.Scope, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SSOConfiguration) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 10 + msgp.StringPrefixSize + len(z.ClientID) + 14 + msgp.StringPrefixSize + len(z.ClientSecret) + 6 + msgp.StringPrefixSize + len(z.Scope)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SSOSetting) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zobc uint32
	zobc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zobc > 0 {
		zobc--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "URL_PREFIX":
			z.URLPrefix, err = dc.ReadString()
			if err != nil {
				return
			}
		case "JS_SDK_CDN_URL":
			z.JSSDKCDNURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "STATE_JWT_SECRET":
			z.StateJWTSecret, err = dc.ReadString()
			if err != nil {
				return
			}
		case "AUTO_LINK_PROVIDER_KEYS":
			var zsnv uint32
			zsnv, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.AutoLinkProviderKeys) >= int(zsnv) {
				z.AutoLinkProviderKeys = (z.AutoLinkProviderKeys)[:zsnv]
			} else {
				z.AutoLinkProviderKeys = make([]string, zsnv)
			}
			for zxpk := range z.AutoLinkProviderKeys {
				z.AutoLinkProviderKeys[zxpk], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "ALLOWED_CALLBACK_URLS":
			var zkgt uint32
			zkgt, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.AllowedCallbackURLs) >= int(zkgt) {
				z.AllowedCallbackURLs = (z.AllowedCallbackURLs)[:zkgt]
			} else {
				z.AllowedCallbackURLs = make([]string, zkgt)
			}
			for zdnj := range z.AllowedCallbackURLs {
				z.AllowedCallbackURLs[zdnj], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SSOSetting) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "URL_PREFIX"
	err = en.Append(0x85, 0xaa, 0x55, 0x52, 0x4c, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58)
	if err != nil {
		return err
	}
	err = en.WriteString(z.URLPrefix)
	if err != nil {
		return
	}
	// write "JS_SDK_CDN_URL"
	err = en.Append(0xae, 0x4a, 0x53, 0x5f, 0x53, 0x44, 0x4b, 0x5f, 0x43, 0x44, 0x4e, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.JSSDKCDNURL)
	if err != nil {
		return
	}
	// write "STATE_JWT_SECRET"
	err = en.Append(0xb0, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4a, 0x57, 0x54, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.StateJWTSecret)
	if err != nil {
		return
	}
	// write "AUTO_LINK_PROVIDER_KEYS"
	err = en.Append(0xb7, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x4b, 0x45, 0x59, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.AutoLinkProviderKeys)))
	if err != nil {
		return
	}
	for zxpk := range z.AutoLinkProviderKeys {
		err = en.WriteString(z.AutoLinkProviderKeys[zxpk])
		if err != nil {
			return
		}
	}
	// write "ALLOWED_CALLBACK_URLS"
	err = en.Append(0xb5, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x55, 0x52, 0x4c, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.AllowedCallbackURLs)))
	if err != nil {
		return
	}
	for zdnj := range z.AllowedCallbackURLs {
		err = en.WriteString(z.AllowedCallbackURLs[zdnj])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SSOSetting) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "URL_PREFIX"
	o = append(o, 0x85, 0xaa, 0x55, 0x52, 0x4c, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58)
	o = msgp.AppendString(o, z.URLPrefix)
	// string "JS_SDK_CDN_URL"
	o = append(o, 0xae, 0x4a, 0x53, 0x5f, 0x53, 0x44, 0x4b, 0x5f, 0x43, 0x44, 0x4e, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.JSSDKCDNURL)
	// string "STATE_JWT_SECRET"
	o = append(o, 0xb0, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4a, 0x57, 0x54, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	o = msgp.AppendString(o, z.StateJWTSecret)
	// string "AUTO_LINK_PROVIDER_KEYS"
	o = append(o, 0xb7, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x4b, 0x45, 0x59, 0x53)
	o = msgp.AppendArrayHeader(o, uint32(len(z.AutoLinkProviderKeys)))
	for zxpk := range z.AutoLinkProviderKeys {
		o = msgp.AppendString(o, z.AutoLinkProviderKeys[zxpk])
	}
	// string "ALLOWED_CALLBACK_URLS"
	o = append(o, 0xb5, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x55, 0x52, 0x4c, 0x53)
	o = msgp.AppendArrayHeader(o, uint32(len(z.AllowedCallbackURLs)))
	for zdnj := range z.AllowedCallbackURLs {
		o = msgp.AppendString(o, z.AllowedCallbackURLs[zdnj])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SSOSetting) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zema uint32
	zema, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zema > 0 {
		zema--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "URL_PREFIX":
			z.URLPrefix, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "JS_SDK_CDN_URL":
			z.JSSDKCDNURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "STATE_JWT_SECRET":
			z.StateJWTSecret, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "AUTO_LINK_PROVIDER_KEYS":
			var zpez uint32
			zpez, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.AutoLinkProviderKeys) >= int(zpez) {
				z.AutoLinkProviderKeys = (z.AutoLinkProviderKeys)[:zpez]
			} else {
				z.AutoLinkProviderKeys = make([]string, zpez)
			}
			for zxpk := range z.AutoLinkProviderKeys {
				z.AutoLinkProviderKeys[zxpk], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "ALLOWED_CALLBACK_URLS":
			var zqke uint32
			zqke, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.AllowedCallbackURLs) >= int(zqke) {
				z.AllowedCallbackURLs = (z.AllowedCallbackURLs)[:zqke]
			} else {
				z.AllowedCallbackURLs = make([]string, zqke)
			}
			for zdnj := range z.AllowedCallbackURLs {
				z.AllowedCallbackURLs[zdnj], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SSOSetting) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.URLPrefix) + 15 + msgp.StringPrefixSize + len(z.JSSDKCDNURL) + 17 + msgp.StringPrefixSize + len(z.StateJWTSecret) + 24 + msgp.ArrayHeaderSize
	for zxpk := range z.AutoLinkProviderKeys {
		s += msgp.StringPrefixSize + len(z.AutoLinkProviderKeys[zxpk])
	}
	s += 22 + msgp.ArrayHeaderSize
	for zdnj := range z.AllowedCallbackURLs {
		s += msgp.StringPrefixSize + len(z.AllowedCallbackURLs[zdnj])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TenantConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zywj uint32
	zywj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zywj > 0 {
		zywj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "DATABASE_URL":
			z.DBConnectionStr, err = dc.ReadString()
			if err != nil {
				return
			}
		case "API_KEY":
			z.APIKey, err = dc.ReadString()
			if err != nil {
				return
			}
		case "MASTER_KEY":
			z.MasterKey, err = dc.ReadString()
			if err != nil {
				return
			}
		case "APP_NAME":
			z.AppName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "CORS_HOST":
			z.CORSHost, err = dc.ReadString()
			if err != nil {
				return
			}
		case "AUTH":
			err = z.Auth.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "TOKEN_STORE":
			var zjpj uint32
			zjpj, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zjpj > 0 {
				zjpj--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "SECRET":
					z.TokenStore.Secret, err = dc.ReadString()
					if err != nil {
						return
					}
				case "EXPIRY":
					z.TokenStore.Expiry, err = dc.ReadInt64()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "USER_PROFILE":
			var zzpf uint32
			zzpf, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zzpf > 0 {
				zzpf--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "IMPLEMENTATION":
					z.UserProfile.ImplName, err = dc.ReadString()
					if err != nil {
						return
					}
				case "IMPL_STORE_URL":
					z.UserProfile.ImplStoreURL, err = dc.ReadString()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "USER_AUDIT":
			err = z.UserAudit.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "SMTP":
			err = z.SMTP.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "TWILIO":
			var zrfe uint32
			zrfe, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zrfe > 0 {
				zrfe--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "ACCOUNT_SID":
					z.Twilio.AccountSID, err = dc.ReadString()
					if err != nil {
						return
					}
				case "AUTH_TOKEN":
					z.Twilio.AuthToken, err = dc.ReadString()
					if err != nil {
						return
					}
				case "FROM":
					z.Twilio.From, err = dc.ReadString()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "NEXMO":
			var zgmo uint32
			zgmo, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zgmo > 0 {
				zgmo--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "API_KEY":
					z.Nexmo.APIKey, err = dc.ReadString()
					if err != nil {
						return
					}
				case "API_SECRET":
					z.Nexmo.APISecret, err = dc.ReadString()
					if err != nil {
						return
					}
				case "FROM":
					z.Nexmo.From, err = dc.ReadString()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		case "FORGOT_PASSWORD":
			err = z.ForgotPassword.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "WELCOME_EMAIL":
			err = z.WelcomeEmail.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "SSO_SETTING":
			err = z.SSOSetting.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "SSO_PROVIDERS":
			var ztaf uint32
			ztaf, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.SSOProviders) >= int(ztaf) {
				z.SSOProviders = (z.SSOProviders)[:ztaf]
			} else {
				z.SSOProviders = make([]string, ztaf)
			}
			for zqyh := range z.SSOProviders {
				z.SSOProviders[zqyh], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "SSO_CONFIGS":
			var zeth uint32
			zeth, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.SSOConfigs) >= int(zeth) {
				z.SSOConfigs = (z.SSOConfigs)[:zeth]
			} else {
				z.SSOConfigs = make([]SSOConfiguration, zeth)
			}
			for zyzr := range z.SSOConfigs {
				err = z.SSOConfigs[zyzr].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "USER_VERIFY":
			err = z.UserVerify.DecodeMsg(dc)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TenantConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 18
	// write "DATABASE_URL"
	err = en.Append(0xde, 0x0, 0x12, 0xac, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.DBConnectionStr)
	if err != nil {
		return
	}
	// write "API_KEY"
	err = en.Append(0xa7, 0x41, 0x50, 0x49, 0x5f, 0x4b, 0x45, 0x59)
	if err != nil {
		return err
	}
	err = en.WriteString(z.APIKey)
	if err != nil {
		return
	}
	// write "MASTER_KEY"
	err = en.Append(0xaa, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x4b, 0x45, 0x59)
	if err != nil {
		return err
	}
	err = en.WriteString(z.MasterKey)
	if err != nil {
		return
	}
	// write "APP_NAME"
	err = en.Append(0xa8, 0x41, 0x50, 0x50, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteString(z.AppName)
	if err != nil {
		return
	}
	// write "CORS_HOST"
	err = en.Append(0xa9, 0x43, 0x4f, 0x52, 0x53, 0x5f, 0x48, 0x4f, 0x53, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.CORSHost)
	if err != nil {
		return
	}
	// write "AUTH"
	err = en.Append(0xa4, 0x41, 0x55, 0x54, 0x48)
	if err != nil {
		return err
	}
	err = z.Auth.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "TOKEN_STORE"
	// map header, size 2
	// write "SECRET"
	err = en.Append(0xab, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x82, 0xa6, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.TokenStore.Secret)
	if err != nil {
		return
	}
	// write "EXPIRY"
	err = en.Append(0xa6, 0x45, 0x58, 0x50, 0x49, 0x52, 0x59)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.TokenStore.Expiry)
	if err != nil {
		return
	}
	// write "USER_PROFILE"
	// map header, size 2
	// write "IMPLEMENTATION"
	err = en.Append(0xac, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x82, 0xae, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.UserProfile.ImplName)
	if err != nil {
		return
	}
	// write "IMPL_STORE_URL"
	err = en.Append(0xae, 0x49, 0x4d, 0x50, 0x4c, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.UserProfile.ImplStoreURL)
	if err != nil {
		return
	}
	// write "USER_AUDIT"
	err = en.Append(0xaa, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x54)
	if err != nil {
		return err
	}
	err = z.UserAudit.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "SMTP"
	err = en.Append(0xa4, 0x53, 0x4d, 0x54, 0x50)
	if err != nil {
		return err
	}
	err = z.SMTP.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "TWILIO"
	// map header, size 3
	// write "ACCOUNT_SID"
	err = en.Append(0xa6, 0x54, 0x57, 0x49, 0x4c, 0x49, 0x4f, 0x83, 0xab, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Twilio.AccountSID)
	if err != nil {
		return
	}
	// write "AUTH_TOKEN"
	err = en.Append(0xaa, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Twilio.AuthToken)
	if err != nil {
		return
	}
	// write "FROM"
	err = en.Append(0xa4, 0x46, 0x52, 0x4f, 0x4d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Twilio.From)
	if err != nil {
		return
	}
	// write "NEXMO"
	// map header, size 3
	// write "API_KEY"
	err = en.Append(0xa5, 0x4e, 0x45, 0x58, 0x4d, 0x4f, 0x83, 0xa7, 0x41, 0x50, 0x49, 0x5f, 0x4b, 0x45, 0x59)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Nexmo.APIKey)
	if err != nil {
		return
	}
	// write "API_SECRET"
	err = en.Append(0xaa, 0x41, 0x50, 0x49, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Nexmo.APISecret)
	if err != nil {
		return
	}
	// write "FROM"
	err = en.Append(0xa4, 0x46, 0x52, 0x4f, 0x4d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Nexmo.From)
	if err != nil {
		return
	}
	// write "FORGOT_PASSWORD"
	err = en.Append(0xaf, 0x46, 0x4f, 0x52, 0x47, 0x4f, 0x54, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44)
	if err != nil {
		return err
	}
	err = z.ForgotPassword.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "WELCOME_EMAIL"
	err = en.Append(0xad, 0x57, 0x45, 0x4c, 0x43, 0x4f, 0x4d, 0x45, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c)
	if err != nil {
		return err
	}
	err = z.WelcomeEmail.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "SSO_SETTING"
	err = en.Append(0xab, 0x53, 0x53, 0x4f, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47)
	if err != nil {
		return err
	}
	err = z.SSOSetting.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "SSO_PROVIDERS"
	err = en.Append(0xad, 0x53, 0x53, 0x4f, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.SSOProviders)))
	if err != nil {
		return
	}
	for zqyh := range z.SSOProviders {
		err = en.WriteString(z.SSOProviders[zqyh])
		if err != nil {
			return
		}
	}
	// write "SSO_CONFIGS"
	err = en.Append(0xab, 0x53, 0x53, 0x4f, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.SSOConfigs)))
	if err != nil {
		return
	}
	for zyzr := range z.SSOConfigs {
		err = z.SSOConfigs[zyzr].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "USER_VERIFY"
	err = en.Append(0xab, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59)
	if err != nil {
		return err
	}
	err = z.UserVerify.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TenantConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 18
	// string "DATABASE_URL"
	o = append(o, 0xde, 0x0, 0x12, 0xac, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.DBConnectionStr)
	// string "API_KEY"
	o = append(o, 0xa7, 0x41, 0x50, 0x49, 0x5f, 0x4b, 0x45, 0x59)
	o = msgp.AppendString(o, z.APIKey)
	// string "MASTER_KEY"
	o = append(o, 0xaa, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x4b, 0x45, 0x59)
	o = msgp.AppendString(o, z.MasterKey)
	// string "APP_NAME"
	o = append(o, 0xa8, 0x41, 0x50, 0x50, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	o = msgp.AppendString(o, z.AppName)
	// string "CORS_HOST"
	o = append(o, 0xa9, 0x43, 0x4f, 0x52, 0x53, 0x5f, 0x48, 0x4f, 0x53, 0x54)
	o = msgp.AppendString(o, z.CORSHost)
	// string "AUTH"
	o = append(o, 0xa4, 0x41, 0x55, 0x54, 0x48)
	o, err = z.Auth.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "TOKEN_STORE"
	// map header, size 2
	// string "SECRET"
	o = append(o, 0xab, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x82, 0xa6, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	o = msgp.AppendString(o, z.TokenStore.Secret)
	// string "EXPIRY"
	o = append(o, 0xa6, 0x45, 0x58, 0x50, 0x49, 0x52, 0x59)
	o = msgp.AppendInt64(o, z.TokenStore.Expiry)
	// string "USER_PROFILE"
	// map header, size 2
	// string "IMPLEMENTATION"
	o = append(o, 0xac, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x82, 0xae, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e)
	o = msgp.AppendString(o, z.UserProfile.ImplName)
	// string "IMPL_STORE_URL"
	o = append(o, 0xae, 0x49, 0x4d, 0x50, 0x4c, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.UserProfile.ImplStoreURL)
	// string "USER_AUDIT"
	o = append(o, 0xaa, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x54)
	o, err = z.UserAudit.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "SMTP"
	o = append(o, 0xa4, 0x53, 0x4d, 0x54, 0x50)
	o, err = z.SMTP.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "TWILIO"
	// map header, size 3
	// string "ACCOUNT_SID"
	o = append(o, 0xa6, 0x54, 0x57, 0x49, 0x4c, 0x49, 0x4f, 0x83, 0xab, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x49, 0x44)
	o = msgp.AppendString(o, z.Twilio.AccountSID)
	// string "AUTH_TOKEN"
	o = append(o, 0xaa, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e)
	o = msgp.AppendString(o, z.Twilio.AuthToken)
	// string "FROM"
	o = append(o, 0xa4, 0x46, 0x52, 0x4f, 0x4d)
	o = msgp.AppendString(o, z.Twilio.From)
	// string "NEXMO"
	// map header, size 3
	// string "API_KEY"
	o = append(o, 0xa5, 0x4e, 0x45, 0x58, 0x4d, 0x4f, 0x83, 0xa7, 0x41, 0x50, 0x49, 0x5f, 0x4b, 0x45, 0x59)
	o = msgp.AppendString(o, z.Nexmo.APIKey)
	// string "API_SECRET"
	o = append(o, 0xaa, 0x41, 0x50, 0x49, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	o = msgp.AppendString(o, z.Nexmo.APISecret)
	// string "FROM"
	o = append(o, 0xa4, 0x46, 0x52, 0x4f, 0x4d)
	o = msgp.AppendString(o, z.Nexmo.From)
	// string "FORGOT_PASSWORD"
	o = append(o, 0xaf, 0x46, 0x4f, 0x52, 0x47, 0x4f, 0x54, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44)
	o, err = z.ForgotPassword.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "WELCOME_EMAIL"
	o = append(o, 0xad, 0x57, 0x45, 0x4c, 0x43, 0x4f, 0x4d, 0x45, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c)
	o, err = z.WelcomeEmail.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "SSO_SETTING"
	o = append(o, 0xab, 0x53, 0x53, 0x4f, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x49, 0x4e, 0x47)
	o, err = z.SSOSetting.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "SSO_PROVIDERS"
	o = append(o, 0xad, 0x53, 0x53, 0x4f, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x53)
	o = msgp.AppendArrayHeader(o, uint32(len(z.SSOProviders)))
	for zqyh := range z.SSOProviders {
		o = msgp.AppendString(o, z.SSOProviders[zqyh])
	}
	// string "SSO_CONFIGS"
	o = append(o, 0xab, 0x53, 0x53, 0x4f, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x53)
	o = msgp.AppendArrayHeader(o, uint32(len(z.SSOConfigs)))
	for zyzr := range z.SSOConfigs {
		o, err = z.SSOConfigs[zyzr].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "USER_VERIFY"
	o = append(o, 0xab, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59)
	o, err = z.UserVerify.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TenantConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsbz uint32
	zsbz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsbz > 0 {
		zsbz--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "DATABASE_URL":
			z.DBConnectionStr, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "API_KEY":
			z.APIKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "MASTER_KEY":
			z.MasterKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "APP_NAME":
			z.AppName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "CORS_HOST":
			z.CORSHost, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "AUTH":
			bts, err = z.Auth.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "TOKEN_STORE":
			var zrjx uint32
			zrjx, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zrjx > 0 {
				zrjx--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "SECRET":
					z.TokenStore.Secret, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "EXPIRY":
					z.TokenStore.Expiry, bts, err = msgp.ReadInt64Bytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "USER_PROFILE":
			var zawn uint32
			zawn, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zawn > 0 {
				zawn--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "IMPLEMENTATION":
					z.UserProfile.ImplName, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "IMPL_STORE_URL":
					z.UserProfile.ImplStoreURL, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "USER_AUDIT":
			bts, err = z.UserAudit.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "SMTP":
			bts, err = z.SMTP.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "TWILIO":
			var zwel uint32
			zwel, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zwel > 0 {
				zwel--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "ACCOUNT_SID":
					z.Twilio.AccountSID, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "AUTH_TOKEN":
					z.Twilio.AuthToken, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "FROM":
					z.Twilio.From, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "NEXMO":
			var zrbe uint32
			zrbe, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zrbe > 0 {
				zrbe--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "API_KEY":
					z.Nexmo.APIKey, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "API_SECRET":
					z.Nexmo.APISecret, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "FROM":
					z.Nexmo.From, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		case "FORGOT_PASSWORD":
			bts, err = z.ForgotPassword.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "WELCOME_EMAIL":
			bts, err = z.WelcomeEmail.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "SSO_SETTING":
			bts, err = z.SSOSetting.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "SSO_PROVIDERS":
			var zmfd uint32
			zmfd, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.SSOProviders) >= int(zmfd) {
				z.SSOProviders = (z.SSOProviders)[:zmfd]
			} else {
				z.SSOProviders = make([]string, zmfd)
			}
			for zqyh := range z.SSOProviders {
				z.SSOProviders[zqyh], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "SSO_CONFIGS":
			var zzdc uint32
			zzdc, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.SSOConfigs) >= int(zzdc) {
				z.SSOConfigs = (z.SSOConfigs)[:zzdc]
			} else {
				z.SSOConfigs = make([]SSOConfiguration, zzdc)
			}
			for zyzr := range z.SSOConfigs {
				bts, err = z.SSOConfigs[zyzr].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "USER_VERIFY":
			bts, err = z.UserVerify.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TenantConfiguration) Msgsize() (s int) {
	s = 3 + 13 + msgp.StringPrefixSize + len(z.DBConnectionStr) + 8 + msgp.StringPrefixSize + len(z.APIKey) + 11 + msgp.StringPrefixSize + len(z.MasterKey) + 9 + msgp.StringPrefixSize + len(z.AppName) + 10 + msgp.StringPrefixSize + len(z.CORSHost) + 5 + z.Auth.Msgsize() + 12 + 1 + 7 + msgp.StringPrefixSize + len(z.TokenStore.Secret) + 7 + msgp.Int64Size + 13 + 1 + 15 + msgp.StringPrefixSize + len(z.UserProfile.ImplName) + 15 + msgp.StringPrefixSize + len(z.UserProfile.ImplStoreURL) + 11 + z.UserAudit.Msgsize() + 5 + z.SMTP.Msgsize() + 7 + 1 + 12 + msgp.StringPrefixSize + len(z.Twilio.AccountSID) + 11 + msgp.StringPrefixSize + len(z.Twilio.AuthToken) + 5 + msgp.StringPrefixSize + len(z.Twilio.From) + 6 + 1 + 8 + msgp.StringPrefixSize + len(z.Nexmo.APIKey) + 11 + msgp.StringPrefixSize + len(z.Nexmo.APISecret) + 5 + msgp.StringPrefixSize + len(z.Nexmo.From) + 16 + z.ForgotPassword.Msgsize() + 14 + z.WelcomeEmail.Msgsize() + 12 + z.SSOSetting.Msgsize() + 14 + msgp.ArrayHeaderSize
	for zqyh := range z.SSOProviders {
		s += msgp.StringPrefixSize + len(z.SSOProviders[zqyh])
	}
	s += 12 + msgp.ArrayHeaderSize
	for zyzr := range z.SSOConfigs {
		s += z.SSOConfigs[zyzr].Msgsize()
	}
	s += 12 + z.UserVerify.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TokenStoreConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zelx uint32
	zelx, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zelx > 0 {
		zelx--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "SECRET":
			z.Secret, err = dc.ReadString()
			if err != nil {
				return
			}
		case "EXPIRY":
			z.Expiry, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z TokenStoreConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "SECRET"
	err = en.Append(0x82, 0xa6, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Secret)
	if err != nil {
		return
	}
	// write "EXPIRY"
	err = en.Append(0xa6, 0x45, 0x58, 0x50, 0x49, 0x52, 0x59)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Expiry)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TokenStoreConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "SECRET"
	o = append(o, 0x82, 0xa6, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	o = msgp.AppendString(o, z.Secret)
	// string "EXPIRY"
	o = append(o, 0xa6, 0x45, 0x58, 0x50, 0x49, 0x52, 0x59)
	o = msgp.AppendInt64(o, z.Expiry)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TokenStoreConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbal uint32
	zbal, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbal > 0 {
		zbal--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "SECRET":
			z.Secret, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "EXPIRY":
			z.Expiry, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TokenStoreConfiguration) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.Secret) + 7 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TwilioConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zjqz uint32
	zjqz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zjqz > 0 {
		zjqz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ACCOUNT_SID":
			z.AccountSID, err = dc.ReadString()
			if err != nil {
				return
			}
		case "AUTH_TOKEN":
			z.AuthToken, err = dc.ReadString()
			if err != nil {
				return
			}
		case "FROM":
			z.From, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z TwilioConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "ACCOUNT_SID"
	err = en.Append(0x83, 0xab, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x49, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteString(z.AccountSID)
	if err != nil {
		return
	}
	// write "AUTH_TOKEN"
	err = en.Append(0xaa, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.AuthToken)
	if err != nil {
		return
	}
	// write "FROM"
	err = en.Append(0xa4, 0x46, 0x52, 0x4f, 0x4d)
	if err != nil {
		return err
	}
	err = en.WriteString(z.From)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z TwilioConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "ACCOUNT_SID"
	o = append(o, 0x83, 0xab, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x49, 0x44)
	o = msgp.AppendString(o, z.AccountSID)
	// string "AUTH_TOKEN"
	o = append(o, 0xaa, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e)
	o = msgp.AppendString(o, z.AuthToken)
	// string "FROM"
	o = append(o, 0xa4, 0x46, 0x52, 0x4f, 0x4d)
	o = msgp.AppendString(o, z.From)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TwilioConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zkct uint32
	zkct, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zkct > 0 {
		zkct--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ACCOUNT_SID":
			z.AccountSID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "AUTH_TOKEN":
			z.AuthToken, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FROM":
			z.From, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z TwilioConfiguration) Msgsize() (s int) {
	s = 1 + 12 + msgp.StringPrefixSize + len(z.AccountSID) + 11 + msgp.StringPrefixSize + len(z.AuthToken) + 5 + msgp.StringPrefixSize + len(z.From)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserAuditConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zana uint32
	zana, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zana > 0 {
		zana--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ENABLED":
			z.Enabled, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "TRAIL_HANDLER_URL":
			z.TrailHandlerURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "PW_MIN_LENGTH":
			z.PwMinLength, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "PW_UPPERCASE_REQUIRED":
			z.PwUppercaseRequired, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "PW_LOWERCASE_REQUIRED":
			z.PwLowercaseRequired, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "PW_DIGIT_REQUIRED":
			z.PwDigitRequired, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "PW_SYMBOL_REQUIRED":
			z.PwSymbolRequired, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "PW_MIN_GUESSABLE_LEVEL":
			z.PwMinGuessableLevel, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "PW_EXCLUDED_KEYWORDS":
			var ztyy uint32
			ztyy, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.PwExcludedKeywords) >= int(ztyy) {
				z.PwExcludedKeywords = (z.PwExcludedKeywords)[:ztyy]
			} else {
				z.PwExcludedKeywords = make([]string, ztyy)
			}
			for ztmt := range z.PwExcludedKeywords {
				z.PwExcludedKeywords[ztmt], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "PW_EXCLUDED_FIELDS":
			var zinl uint32
			zinl, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.PwExcludedFields) >= int(zinl) {
				z.PwExcludedFields = (z.PwExcludedFields)[:zinl]
			} else {
				z.PwExcludedFields = make([]string, zinl)
			}
			for ztco := range z.PwExcludedFields {
				z.PwExcludedFields[ztco], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "PW_HISTORY_SIZE":
			z.PwHistorySize, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "PW_HISTORY_DAYS":
			z.PwHistoryDays, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "PW_EXPIRY_DAYS":
			z.PwExpiryDays, err = dc.ReadInt()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *UserAuditConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 13
	// write "ENABLED"
	err = en.Append(0x8d, 0xa7, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Enabled)
	if err != nil {
		return
	}
	// write "TRAIL_HANDLER_URL"
	err = en.Append(0xb1, 0x54, 0x52, 0x41, 0x49, 0x4c, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x52, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.TrailHandlerURL)
	if err != nil {
		return
	}
	// write "PW_MIN_LENGTH"
	err = en.Append(0xad, 0x50, 0x57, 0x5f, 0x4d, 0x49, 0x4e, 0x5f, 0x4c, 0x45, 0x4e, 0x47, 0x54, 0x48)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.PwMinLength)
	if err != nil {
		return
	}
	// write "PW_UPPERCASE_REQUIRED"
	err = en.Append(0xb5, 0x50, 0x57, 0x5f, 0x55, 0x50, 0x50, 0x45, 0x52, 0x43, 0x41, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.PwUppercaseRequired)
	if err != nil {
		return
	}
	// write "PW_LOWERCASE_REQUIRED"
	err = en.Append(0xb5, 0x50, 0x57, 0x5f, 0x4c, 0x4f, 0x57, 0x45, 0x52, 0x43, 0x41, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.PwLowercaseRequired)
	if err != nil {
		return
	}
	// write "PW_DIGIT_REQUIRED"
	err = en.Append(0xb1, 0x50, 0x57, 0x5f, 0x44, 0x49, 0x47, 0x49, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.PwDigitRequired)
	if err != nil {
		return
	}
	// write "PW_SYMBOL_REQUIRED"
	err = en.Append(0xb2, 0x50, 0x57, 0x5f, 0x53, 0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.PwSymbolRequired)
	if err != nil {
		return
	}
	// write "PW_MIN_GUESSABLE_LEVEL"
	err = en.Append(0xb6, 0x50, 0x57, 0x5f, 0x4d, 0x49, 0x4e, 0x5f, 0x47, 0x55, 0x45, 0x53, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.PwMinGuessableLevel)
	if err != nil {
		return
	}
	// write "PW_EXCLUDED_KEYWORDS"
	err = en.Append(0xb4, 0x50, 0x57, 0x5f, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.PwExcludedKeywords)))
	if err != nil {
		return
	}
	for ztmt := range z.PwExcludedKeywords {
		err = en.WriteString(z.PwExcludedKeywords[ztmt])
		if err != nil {
			return
		}
	}
	// write "PW_EXCLUDED_FIELDS"
	err = en.Append(0xb2, 0x50, 0x57, 0x5f, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.PwExcludedFields)))
	if err != nil {
		return
	}
	for ztco := range z.PwExcludedFields {
		err = en.WriteString(z.PwExcludedFields[ztco])
		if err != nil {
			return
		}
	}
	// write "PW_HISTORY_SIZE"
	err = en.Append(0xaf, 0x50, 0x57, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x49, 0x5a, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.PwHistorySize)
	if err != nil {
		return
	}
	// write "PW_HISTORY_DAYS"
	err = en.Append(0xaf, 0x50, 0x57, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x44, 0x41, 0x59, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.PwHistoryDays)
	if err != nil {
		return
	}
	// write "PW_EXPIRY_DAYS"
	err = en.Append(0xae, 0x50, 0x57, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x59, 0x5f, 0x44, 0x41, 0x59, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.PwExpiryDays)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *UserAuditConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 13
	// string "ENABLED"
	o = append(o, 0x8d, 0xa7, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44)
	o = msgp.AppendBool(o, z.Enabled)
	// string "TRAIL_HANDLER_URL"
	o = append(o, 0xb1, 0x54, 0x52, 0x41, 0x49, 0x4c, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x52, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.TrailHandlerURL)
	// string "PW_MIN_LENGTH"
	o = append(o, 0xad, 0x50, 0x57, 0x5f, 0x4d, 0x49, 0x4e, 0x5f, 0x4c, 0x45, 0x4e, 0x47, 0x54, 0x48)
	o = msgp.AppendInt(o, z.PwMinLength)
	// string "PW_UPPERCASE_REQUIRED"
	o = append(o, 0xb5, 0x50, 0x57, 0x5f, 0x55, 0x50, 0x50, 0x45, 0x52, 0x43, 0x41, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44)
	o = msgp.AppendBool(o, z.PwUppercaseRequired)
	// string "PW_LOWERCASE_REQUIRED"
	o = append(o, 0xb5, 0x50, 0x57, 0x5f, 0x4c, 0x4f, 0x57, 0x45, 0x52, 0x43, 0x41, 0x53, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44)
	o = msgp.AppendBool(o, z.PwLowercaseRequired)
	// string "PW_DIGIT_REQUIRED"
	o = append(o, 0xb1, 0x50, 0x57, 0x5f, 0x44, 0x49, 0x47, 0x49, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44)
	o = msgp.AppendBool(o, z.PwDigitRequired)
	// string "PW_SYMBOL_REQUIRED"
	o = append(o, 0xb2, 0x50, 0x57, 0x5f, 0x53, 0x59, 0x4d, 0x42, 0x4f, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44)
	o = msgp.AppendBool(o, z.PwSymbolRequired)
	// string "PW_MIN_GUESSABLE_LEVEL"
	o = append(o, 0xb6, 0x50, 0x57, 0x5f, 0x4d, 0x49, 0x4e, 0x5f, 0x47, 0x55, 0x45, 0x53, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c)
	o = msgp.AppendInt(o, z.PwMinGuessableLevel)
	// string "PW_EXCLUDED_KEYWORDS"
	o = append(o, 0xb4, 0x50, 0x57, 0x5f, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x53)
	o = msgp.AppendArrayHeader(o, uint32(len(z.PwExcludedKeywords)))
	for ztmt := range z.PwExcludedKeywords {
		o = msgp.AppendString(o, z.PwExcludedKeywords[ztmt])
	}
	// string "PW_EXCLUDED_FIELDS"
	o = append(o, 0xb2, 0x50, 0x57, 0x5f, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x53)
	o = msgp.AppendArrayHeader(o, uint32(len(z.PwExcludedFields)))
	for ztco := range z.PwExcludedFields {
		o = msgp.AppendString(o, z.PwExcludedFields[ztco])
	}
	// string "PW_HISTORY_SIZE"
	o = append(o, 0xaf, 0x50, 0x57, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x49, 0x5a, 0x45)
	o = msgp.AppendInt(o, z.PwHistorySize)
	// string "PW_HISTORY_DAYS"
	o = append(o, 0xaf, 0x50, 0x57, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x44, 0x41, 0x59, 0x53)
	o = msgp.AppendInt(o, z.PwHistoryDays)
	// string "PW_EXPIRY_DAYS"
	o = append(o, 0xae, 0x50, 0x57, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x59, 0x5f, 0x44, 0x41, 0x59, 0x53)
	o = msgp.AppendInt(o, z.PwExpiryDays)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserAuditConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zare uint32
	zare, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zare > 0 {
		zare--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ENABLED":
			z.Enabled, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "TRAIL_HANDLER_URL":
			z.TrailHandlerURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "PW_MIN_LENGTH":
			z.PwMinLength, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "PW_UPPERCASE_REQUIRED":
			z.PwUppercaseRequired, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "PW_LOWERCASE_REQUIRED":
			z.PwLowercaseRequired, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "PW_DIGIT_REQUIRED":
			z.PwDigitRequired, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "PW_SYMBOL_REQUIRED":
			z.PwSymbolRequired, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "PW_MIN_GUESSABLE_LEVEL":
			z.PwMinGuessableLevel, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "PW_EXCLUDED_KEYWORDS":
			var zljy uint32
			zljy, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.PwExcludedKeywords) >= int(zljy) {
				z.PwExcludedKeywords = (z.PwExcludedKeywords)[:zljy]
			} else {
				z.PwExcludedKeywords = make([]string, zljy)
			}
			for ztmt := range z.PwExcludedKeywords {
				z.PwExcludedKeywords[ztmt], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "PW_EXCLUDED_FIELDS":
			var zixj uint32
			zixj, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.PwExcludedFields) >= int(zixj) {
				z.PwExcludedFields = (z.PwExcludedFields)[:zixj]
			} else {
				z.PwExcludedFields = make([]string, zixj)
			}
			for ztco := range z.PwExcludedFields {
				z.PwExcludedFields[ztco], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "PW_HISTORY_SIZE":
			z.PwHistorySize, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "PW_HISTORY_DAYS":
			z.PwHistoryDays, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "PW_EXPIRY_DAYS":
			z.PwExpiryDays, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UserAuditConfiguration) Msgsize() (s int) {
	s = 1 + 8 + msgp.BoolSize + 18 + msgp.StringPrefixSize + len(z.TrailHandlerURL) + 14 + msgp.IntSize + 22 + msgp.BoolSize + 22 + msgp.BoolSize + 18 + msgp.BoolSize + 19 + msgp.BoolSize + 23 + msgp.IntSize + 21 + msgp.ArrayHeaderSize
	for ztmt := range z.PwExcludedKeywords {
		s += msgp.StringPrefixSize + len(z.PwExcludedKeywords[ztmt])
	}
	s += 19 + msgp.ArrayHeaderSize
	for ztco := range z.PwExcludedFields {
		s += msgp.StringPrefixSize + len(z.PwExcludedFields[ztco])
	}
	s += 16 + msgp.IntSize + 16 + msgp.IntSize + 15 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserProfileConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrsc uint32
	zrsc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrsc > 0 {
		zrsc--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "IMPLEMENTATION":
			z.ImplName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "IMPL_STORE_URL":
			z.ImplStoreURL, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z UserProfileConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "IMPLEMENTATION"
	err = en.Append(0x82, 0xae, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ImplName)
	if err != nil {
		return
	}
	// write "IMPL_STORE_URL"
	err = en.Append(0xae, 0x49, 0x4d, 0x50, 0x4c, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ImplStoreURL)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z UserProfileConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "IMPLEMENTATION"
	o = append(o, 0x82, 0xae, 0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e)
	o = msgp.AppendString(o, z.ImplName)
	// string "IMPL_STORE_URL"
	o = append(o, 0xae, 0x49, 0x4d, 0x50, 0x4c, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.ImplStoreURL)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserProfileConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zctn uint32
	zctn, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zctn > 0 {
		zctn--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "IMPLEMENTATION":
			z.ImplName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "IMPL_STORE_URL":
			z.ImplStoreURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z UserProfileConfiguration) Msgsize() (s int) {
	s = 1 + 15 + msgp.StringPrefixSize + len(z.ImplName) + 15 + msgp.StringPrefixSize + len(z.ImplStoreURL)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserVerifyConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zrus uint32
	zrus, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zrus > 0 {
		zrus--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "URL_PREFIX":
			z.URLPrefix, err = dc.ReadString()
			if err != nil {
				return
			}
		case "AUTO_UPDATE":
			z.AutoUpdate, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "AUTO_SEND_SIGNUP":
			z.AutoSendOnSignup, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "AUTO_SEND_UPDATE":
			z.AutoSendOnUpdate, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "REQUIRED":
			z.Required, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "CRITERIA":
			z.Criteria, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ERROR_REDIRECT":
			z.ErrorRedirect, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ERROR_HTML_URL":
			z.ErrorHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "KEYS":
			var zsvm uint32
			zsvm, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Keys) >= int(zsvm) {
				z.Keys = (z.Keys)[:zsvm]
			} else {
				z.Keys = make([]string, zsvm)
			}
			for zswy := range z.Keys {
				z.Keys[zswy], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "KEY_CONFIGS":
			var zaoz uint32
			zaoz, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.KeyConfigs) >= int(zaoz) {
				z.KeyConfigs = (z.KeyConfigs)[:zaoz]
			} else {
				z.KeyConfigs = make([]UserVerifyKeyConfiguration, zaoz)
			}
			for znsg := range z.KeyConfigs {
				err = z.KeyConfigs[znsg].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *UserVerifyConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 10
	// write "URL_PREFIX"
	err = en.Append(0x8a, 0xaa, 0x55, 0x52, 0x4c, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58)
	if err != nil {
		return err
	}
	err = en.WriteString(z.URLPrefix)
	if err != nil {
		return
	}
	// write "AUTO_UPDATE"
	err = en.Append(0xab, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.AutoUpdate)
	if err != nil {
		return
	}
	// write "AUTO_SEND_SIGNUP"
	err = en.Append(0xb0, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x55, 0x50)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.AutoSendOnSignup)
	if err != nil {
		return
	}
	// write "AUTO_SEND_UPDATE"
	err = en.Append(0xb0, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.AutoSendOnUpdate)
	if err != nil {
		return
	}
	// write "REQUIRED"
	err = en.Append(0xa8, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Required)
	if err != nil {
		return
	}
	// write "CRITERIA"
	err = en.Append(0xa8, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x41)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Criteria)
	if err != nil {
		return
	}
	// write "ERROR_REDIRECT"
	err = en.Append(0xae, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ErrorRedirect)
	if err != nil {
		return
	}
	// write "ERROR_HTML_URL"
	err = en.Append(0xae, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ErrorHTMLURL)
	if err != nil {
		return
	}
	// write "KEYS"
	err = en.Append(0xa4, 0x4b, 0x45, 0x59, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Keys)))
	if err != nil {
		return
	}
	for zswy := range z.Keys {
		err = en.WriteString(z.Keys[zswy])
		if err != nil {
			return
		}
	}
	// write "KEY_CONFIGS"
	err = en.Append(0xab, 0x4b, 0x45, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x53)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.KeyConfigs)))
	if err != nil {
		return
	}
	for znsg := range z.KeyConfigs {
		err = z.KeyConfigs[znsg].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *UserVerifyConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "URL_PREFIX"
	o = append(o, 0x8a, 0xaa, 0x55, 0x52, 0x4c, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58)
	o = msgp.AppendString(o, z.URLPrefix)
	// string "AUTO_UPDATE"
	o = append(o, 0xab, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45)
	o = msgp.AppendBool(o, z.AutoUpdate)
	// string "AUTO_SEND_SIGNUP"
	o = append(o, 0xb0, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x55, 0x50)
	o = msgp.AppendBool(o, z.AutoSendOnSignup)
	// string "AUTO_SEND_UPDATE"
	o = append(o, 0xb0, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45)
	o = msgp.AppendBool(o, z.AutoSendOnUpdate)
	// string "REQUIRED"
	o = append(o, 0xa8, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44)
	o = msgp.AppendBool(o, z.Required)
	// string "CRITERIA"
	o = append(o, 0xa8, 0x43, 0x52, 0x49, 0x54, 0x45, 0x52, 0x49, 0x41)
	o = msgp.AppendString(o, z.Criteria)
	// string "ERROR_REDIRECT"
	o = append(o, 0xae, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54)
	o = msgp.AppendString(o, z.ErrorRedirect)
	// string "ERROR_HTML_URL"
	o = append(o, 0xae, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.ErrorHTMLURL)
	// string "KEYS"
	o = append(o, 0xa4, 0x4b, 0x45, 0x59, 0x53)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Keys)))
	for zswy := range z.Keys {
		o = msgp.AppendString(o, z.Keys[zswy])
	}
	// string "KEY_CONFIGS"
	o = append(o, 0xab, 0x4b, 0x45, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x53)
	o = msgp.AppendArrayHeader(o, uint32(len(z.KeyConfigs)))
	for znsg := range z.KeyConfigs {
		o, err = z.KeyConfigs[znsg].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserVerifyConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zfzb uint32
	zfzb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zfzb > 0 {
		zfzb--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "URL_PREFIX":
			z.URLPrefix, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "AUTO_UPDATE":
			z.AutoUpdate, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "AUTO_SEND_SIGNUP":
			z.AutoSendOnSignup, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "AUTO_SEND_UPDATE":
			z.AutoSendOnUpdate, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "REQUIRED":
			z.Required, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "CRITERIA":
			z.Criteria, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ERROR_REDIRECT":
			z.ErrorRedirect, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ERROR_HTML_URL":
			z.ErrorHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "KEYS":
			var zsbo uint32
			zsbo, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Keys) >= int(zsbo) {
				z.Keys = (z.Keys)[:zsbo]
			} else {
				z.Keys = make([]string, zsbo)
			}
			for zswy := range z.Keys {
				z.Keys[zswy], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "KEY_CONFIGS":
			var zjif uint32
			zjif, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.KeyConfigs) >= int(zjif) {
				z.KeyConfigs = (z.KeyConfigs)[:zjif]
			} else {
				z.KeyConfigs = make([]UserVerifyKeyConfiguration, zjif)
			}
			for znsg := range z.KeyConfigs {
				bts, err = z.KeyConfigs[znsg].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UserVerifyConfiguration) Msgsize() (s int) {
	s = 1 + 11 + msgp.StringPrefixSize + len(z.URLPrefix) + 12 + msgp.BoolSize + 17 + msgp.BoolSize + 17 + msgp.BoolSize + 9 + msgp.BoolSize + 9 + msgp.StringPrefixSize + len(z.Criteria) + 15 + msgp.StringPrefixSize + len(z.ErrorRedirect) + 15 + msgp.StringPrefixSize + len(z.ErrorHTMLURL) + 5 + msgp.ArrayHeaderSize
	for zswy := range z.Keys {
		s += msgp.StringPrefixSize + len(z.Keys[zswy])
	}
	s += 12 + msgp.ArrayHeaderSize
	for znsg := range z.KeyConfigs {
		s += z.KeyConfigs[znsg].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserVerifyKeyConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zqgz uint32
	zqgz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zqgz > 0 {
		zqgz--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "KEY":
			z.Key, err = dc.ReadString()
			if err != nil {
				return
			}
		case "CODE_FORMAT":
			z.CodeFormat, err = dc.ReadString()
			if err != nil {
				return
			}
		case "EXPIRY":
			z.Expiry, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "SUCCESS_REDIRECT":
			z.SuccessRedirect, err = dc.ReadString()
			if err != nil {
				return
			}
		case "SUCCESS_HTML_URL":
			z.SuccessHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ERROR_REDIRECT":
			z.ErrorRedirect, err = dc.ReadString()
			if err != nil {
				return
			}
		case "ERROR_HTML_URL":
			z.ErrorHTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "PROVIDER":
			z.Provider, err = dc.ReadString()
			if err != nil {
				return
			}
		case "PROVIDER_CONFIG":
			err = z.ProviderConfig.DecodeMsg(dc)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *UserVerifyKeyConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "KEY"
	err = en.Append(0x89, 0xa3, 0x4b, 0x45, 0x59)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Key)
	if err != nil {
		return
	}
	// write "CODE_FORMAT"
	err = en.Append(0xab, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.CodeFormat)
	if err != nil {
		return
	}
	// write "EXPIRY"
	err = en.Append(0xa6, 0x45, 0x58, 0x50, 0x49, 0x52, 0x59)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Expiry)
	if err != nil {
		return
	}
	// write "SUCCESS_REDIRECT"
	err = en.Append(0xb0, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SuccessRedirect)
	if err != nil {
		return
	}
	// write "SUCCESS_HTML_URL"
	err = en.Append(0xb0, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SuccessHTMLURL)
	if err != nil {
		return
	}
	// write "ERROR_REDIRECT"
	err = en.Append(0xae, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ErrorRedirect)
	if err != nil {
		return
	}
	// write "ERROR_HTML_URL"
	err = en.Append(0xae, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ErrorHTMLURL)
	if err != nil {
		return
	}
	// write "PROVIDER"
	err = en.Append(0xa8, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Provider)
	if err != nil {
		return
	}
	// write "PROVIDER_CONFIG"
	err = en.Append(0xaf, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47)
	if err != nil {
		return err
	}
	err = z.ProviderConfig.EncodeMsg(en)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *UserVerifyKeyConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "KEY"
	o = append(o, 0x89, 0xa3, 0x4b, 0x45, 0x59)
	o = msgp.AppendString(o, z.Key)
	// string "CODE_FORMAT"
	o = append(o, 0xab, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54)
	o = msgp.AppendString(o, z.CodeFormat)
	// string "EXPIRY"
	o = append(o, 0xa6, 0x45, 0x58, 0x50, 0x49, 0x52, 0x59)
	o = msgp.AppendInt64(o, z.Expiry)
	// string "SUCCESS_REDIRECT"
	o = append(o, 0xb0, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54)
	o = msgp.AppendString(o, z.SuccessRedirect)
	// string "SUCCESS_HTML_URL"
	o = append(o, 0xb0, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.SuccessHTMLURL)
	// string "ERROR_REDIRECT"
	o = append(o, 0xae, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54)
	o = msgp.AppendString(o, z.ErrorRedirect)
	// string "ERROR_HTML_URL"
	o = append(o, 0xae, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.ErrorHTMLURL)
	// string "PROVIDER"
	o = append(o, 0xa8, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52)
	o = msgp.AppendString(o, z.Provider)
	// string "PROVIDER_CONFIG"
	o = append(o, 0xaf, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47)
	o, err = z.ProviderConfig.MarshalMsg(o)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserVerifyKeyConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zsnw uint32
	zsnw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zsnw > 0 {
		zsnw--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "KEY":
			z.Key, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "CODE_FORMAT":
			z.CodeFormat, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "EXPIRY":
			z.Expiry, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "SUCCESS_REDIRECT":
			z.SuccessRedirect, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SUCCESS_HTML_URL":
			z.SuccessHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ERROR_REDIRECT":
			z.ErrorRedirect, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ERROR_HTML_URL":
			z.ErrorHTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "PROVIDER":
			z.Provider, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "PROVIDER_CONFIG":
			bts, err = z.ProviderConfig.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UserVerifyKeyConfiguration) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Key) + 12 + msgp.StringPrefixSize + len(z.CodeFormat) + 7 + msgp.Int64Size + 17 + msgp.StringPrefixSize + len(z.SuccessRedirect) + 17 + msgp.StringPrefixSize + len(z.SuccessHTMLURL) + 15 + msgp.StringPrefixSize + len(z.ErrorRedirect) + 15 + msgp.StringPrefixSize + len(z.ErrorHTMLURL) + 9 + msgp.StringPrefixSize + len(z.Provider) + 16 + z.ProviderConfig.Msgsize()
	return
}

// DecodeMsg implements msgp.Decodable
func (z *UserVerifyKeyProviderConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var ztls uint32
	ztls, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for ztls > 0 {
		ztls--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "SUBJECT":
			z.Subject, err = dc.ReadString()
			if err != nil {
				return
			}
		case "SENDER":
			z.Sender, err = dc.ReadString()
			if err != nil {
				return
			}
		case "SENDER_NAME":
			z.SenderName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "REPLY_TO":
			z.ReplyTo, err = dc.ReadString()
			if err != nil {
				return
			}
		case "REPLY_TO_NAME":
			z.ReplyToName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "TEXT_URL":
			z.TextURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "HTML_URL":
			z.HTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *UserVerifyKeyProviderConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "SUBJECT"
	err = en.Append(0x87, 0xa7, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Subject)
	if err != nil {
		return
	}
	// write "SENDER"
	err = en.Append(0xa6, 0x53, 0x45, 0x4e, 0x44, 0x45, 0x52)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Sender)
	if err != nil {
		return
	}
	// write "SENDER_NAME"
	err = en.Append(0xab, 0x53, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SenderName)
	if err != nil {
		return
	}
	// write "REPLY_TO"
	err = en.Append(0xa8, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x54, 0x4f)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ReplyTo)
	if err != nil {
		return
	}
	// write "REPLY_TO_NAME"
	err = en.Append(0xad, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ReplyToName)
	if err != nil {
		return
	}
	// write "TEXT_URL"
	err = en.Append(0xa8, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.TextURL)
	if err != nil {
		return
	}
	// write "HTML_URL"
	err = en.Append(0xa8, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.HTMLURL)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *UserVerifyKeyProviderConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "SUBJECT"
	o = append(o, 0x87, 0xa7, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54)
	o = msgp.AppendString(o, z.Subject)
	// string "SENDER"
	o = append(o, 0xa6, 0x53, 0x45, 0x4e, 0x44, 0x45, 0x52)
	o = msgp.AppendString(o, z.Sender)
	// string "SENDER_NAME"
	o = append(o, 0xab, 0x53, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	o = msgp.AppendString(o, z.SenderName)
	// string "REPLY_TO"
	o = append(o, 0xa8, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x54, 0x4f)
	o = msgp.AppendString(o, z.ReplyTo)
	// string "REPLY_TO_NAME"
	o = append(o, 0xad, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	o = msgp.AppendString(o, z.ReplyToName)
	// string "TEXT_URL"
	o = append(o, 0xa8, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.TextURL)
	// string "HTML_URL"
	o = append(o, 0xa8, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.HTMLURL)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *UserVerifyKeyProviderConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zmvo uint32
	zmvo, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zmvo > 0 {
		zmvo--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "SUBJECT":
			z.Subject, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SENDER":
			z.Sender, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SENDER_NAME":
			z.SenderName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "REPLY_TO":
			z.ReplyTo, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "REPLY_TO_NAME":
			z.ReplyToName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "TEXT_URL":
			z.TextURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "HTML_URL":
			z.HTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *UserVerifyKeyProviderConfiguration) Msgsize() (s int) {
	s = 1 + 8 + msgp.StringPrefixSize + len(z.Subject) + 7 + msgp.StringPrefixSize + len(z.Sender) + 12 + msgp.StringPrefixSize + len(z.SenderName) + 9 + msgp.StringPrefixSize + len(z.ReplyTo) + 14 + msgp.StringPrefixSize + len(z.ReplyToName) + 9 + msgp.StringPrefixSize + len(z.TextURL) + 9 + msgp.StringPrefixSize + len(z.HTMLURL)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *WelcomeEmailConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zigk uint32
	zigk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zigk > 0 {
		zigk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ENABLED":
			z.Enabled, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "SENDER_NAME":
			z.SenderName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "SENDER":
			z.Sender, err = dc.ReadString()
			if err != nil {
				return
			}
		case "SUBJECT":
			z.Subject, err = dc.ReadString()
			if err != nil {
				return
			}
		case "REPLY_TO_NAME":
			z.ReplyToName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "REPLY_TO":
			z.ReplyTo, err = dc.ReadString()
			if err != nil {
				return
			}
		case "TEXT_URL":
			z.TextURL, err = dc.ReadString()
			if err != nil {
				return
			}
		case "HTML_URL":
			z.HTMLURL, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *WelcomeEmailConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 8
	// write "ENABLED"
	err = en.Append(0x88, 0xa7, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44)
	if err != nil {
		return err
	}
	err = en.WriteBool(z.Enabled)
	if err != nil {
		return
	}
	// write "SENDER_NAME"
	err = en.Append(0xab, 0x53, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteString(z.SenderName)
	if err != nil {
		return
	}
	// write "SENDER"
	err = en.Append(0xa6, 0x53, 0x45, 0x4e, 0x44, 0x45, 0x52)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Sender)
	if err != nil {
		return
	}
	// write "SUBJECT"
	err = en.Append(0xa7, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Subject)
	if err != nil {
		return
	}
	// write "REPLY_TO_NAME"
	err = en.Append(0xad, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ReplyToName)
	if err != nil {
		return
	}
	// write "REPLY_TO"
	err = en.Append(0xa8, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x54, 0x4f)
	if err != nil {
		return err
	}
	err = en.WriteString(z.ReplyTo)
	if err != nil {
		return
	}
	// write "TEXT_URL"
	err = en.Append(0xa8, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.TextURL)
	if err != nil {
		return
	}
	// write "HTML_URL"
	err = en.Append(0xa8, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return err
	}
	err = en.WriteString(z.HTMLURL)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *WelcomeEmailConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "ENABLED"
	o = append(o, 0x88, 0xa7, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44)
	o = msgp.AppendBool(o, z.Enabled)
	// string "SENDER_NAME"
	o = append(o, 0xab, 0x53, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	o = msgp.AppendString(o, z.SenderName)
	// string "SENDER"
	o = append(o, 0xa6, 0x53, 0x45, 0x4e, 0x44, 0x45, 0x52)
	o = msgp.AppendString(o, z.Sender)
	// string "SUBJECT"
	o = append(o, 0xa7, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54)
	o = msgp.AppendString(o, z.Subject)
	// string "REPLY_TO_NAME"
	o = append(o, 0xad, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	o = msgp.AppendString(o, z.ReplyToName)
	// string "REPLY_TO"
	o = append(o, 0xa8, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x5f, 0x54, 0x4f)
	o = msgp.AppendString(o, z.ReplyTo)
	// string "TEXT_URL"
	o = append(o, 0xa8, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.TextURL)
	// string "HTML_URL"
	o = append(o, 0xa8, 0x48, 0x54, 0x4d, 0x4c, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.HTMLURL)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *WelcomeEmailConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zopb uint32
	zopb, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zopb > 0 {
		zopb--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ENABLED":
			z.Enabled, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "SENDER_NAME":
			z.SenderName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SENDER":
			z.Sender, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SUBJECT":
			z.Subject, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "REPLY_TO_NAME":
			z.ReplyToName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "REPLY_TO":
			z.ReplyTo, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "TEXT_URL":
			z.TextURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "HTML_URL":
			z.HTMLURL, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *WelcomeEmailConfiguration) Msgsize() (s int) {
	s = 1 + 8 + msgp.BoolSize + 12 + msgp.StringPrefixSize + len(z.SenderName) + 7 + msgp.StringPrefixSize + len(z.Sender) + 8 + msgp.StringPrefixSize + len(z.Subject) + 14 + msgp.StringPrefixSize + len(z.ReplyToName) + 9 + msgp.StringPrefixSize + len(z.ReplyTo) + 9 + msgp.StringPrefixSize + len(z.TextURL) + 9 + msgp.StringPrefixSize + len(z.HTMLURL)
	return
}
