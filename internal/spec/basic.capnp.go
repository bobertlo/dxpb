// Code generated by capnpc-go. DO NOT EDIT.

package spec

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type GithubEvent struct{ capnp.Struct }
type GithubEvent_commit GithubEvent
type GithubEvent_pullRequest GithubEvent
type GithubEvent_issue GithubEvent
type GithubEvent_Which uint16

const (
	GithubEvent_Which_commit      GithubEvent_Which = 0
	GithubEvent_Which_pullRequest GithubEvent_Which = 1
	GithubEvent_Which_issue       GithubEvent_Which = 2
)

func (w GithubEvent_Which) String() string {
	const s = "commitpullRequestissue"
	switch w {
	case GithubEvent_Which_commit:
		return s[0:6]
	case GithubEvent_Which_pullRequest:
		return s[6:17]
	case GithubEvent_Which_issue:
		return s[17:22]

	}
	return "GithubEvent_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// GithubEvent_TypeID is the unique identifier for the type GithubEvent.
const GithubEvent_TypeID = 0x992b0cc20a124b84

func NewGithubEvent(s *capnp.Segment) (GithubEvent, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4})
	return GithubEvent{st}, err
}

func NewRootGithubEvent(s *capnp.Segment) (GithubEvent, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4})
	return GithubEvent{st}, err
}

func ReadRootGithubEvent(msg *capnp.Message) (GithubEvent, error) {
	root, err := msg.RootPtr()
	return GithubEvent{root.Struct()}, err
}

func (s GithubEvent) String() string {
	str, _ := text.Marshal(0x992b0cc20a124b84, s.Struct)
	return str
}

func (s GithubEvent) Which() GithubEvent_Which {
	return GithubEvent_Which(s.Struct.Uint16(0))
}
func (s GithubEvent) Who() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s GithubEvent) HasWho() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s GithubEvent) WhoBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s GithubEvent) SetWho(v string) error {
	return s.Struct.SetText(0, v)
}

func (s GithubEvent) Commit() GithubEvent_commit { return GithubEvent_commit(s) }

func (s GithubEvent) SetCommit() {
	s.Struct.SetUint16(0, 0)
}

func (s GithubEvent_commit) Hash() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s GithubEvent_commit) HasHash() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s GithubEvent_commit) HashBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s GithubEvent_commit) SetHash(v string) error {
	return s.Struct.SetText(1, v)
}

func (s GithubEvent_commit) Branch() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s GithubEvent_commit) HasBranch() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s GithubEvent_commit) BranchBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s GithubEvent_commit) SetBranch(v string) error {
	return s.Struct.SetText(2, v)
}

func (s GithubEvent_commit) Msg() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s GithubEvent_commit) HasMsg() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s GithubEvent_commit) MsgBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s GithubEvent_commit) SetMsg(v string) error {
	return s.Struct.SetText(3, v)
}

func (s GithubEvent) PullRequest() GithubEvent_pullRequest { return GithubEvent_pullRequest(s) }

func (s GithubEvent) SetPullRequest() {
	s.Struct.SetUint16(0, 1)
}

func (s GithubEvent_pullRequest) Number() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s GithubEvent_pullRequest) SetNumber(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

func (s GithubEvent_pullRequest) Msg() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s GithubEvent_pullRequest) HasMsg() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s GithubEvent_pullRequest) MsgBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s GithubEvent_pullRequest) SetMsg(v string) error {
	return s.Struct.SetText(1, v)
}

func (s GithubEvent_pullRequest) Action() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s GithubEvent_pullRequest) HasAction() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s GithubEvent_pullRequest) ActionBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s GithubEvent_pullRequest) SetAction(v string) error {
	return s.Struct.SetText(2, v)
}

func (s GithubEvent) Issue() GithubEvent_issue { return GithubEvent_issue(s) }

func (s GithubEvent) SetIssue() {
	s.Struct.SetUint16(0, 2)
}

func (s GithubEvent_issue) Number() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s GithubEvent_issue) SetNumber(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

func (s GithubEvent_issue) Msg() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s GithubEvent_issue) HasMsg() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s GithubEvent_issue) MsgBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s GithubEvent_issue) SetMsg(v string) error {
	return s.Struct.SetText(1, v)
}

func (s GithubEvent_issue) Action() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s GithubEvent_issue) HasAction() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s GithubEvent_issue) ActionBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s GithubEvent_issue) SetAction(v string) error {
	return s.Struct.SetText(2, v)
}

// GithubEvent_List is a list of GithubEvent.
type GithubEvent_List struct{ capnp.List }

// NewGithubEvent creates a new list of GithubEvent.
func NewGithubEvent_List(s *capnp.Segment, sz int32) (GithubEvent_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4}, sz)
	return GithubEvent_List{l}, err
}

func (s GithubEvent_List) At(i int) GithubEvent { return GithubEvent{s.List.Struct(i)} }

func (s GithubEvent_List) Set(i int, v GithubEvent) error { return s.List.SetStruct(i, v.Struct) }

func (s GithubEvent_List) String() string {
	str, _ := text.MarshalList(0x992b0cc20a124b84, s.List)
	return str
}

// GithubEvent_Promise is a wrapper for a GithubEvent promised by a client call.
type GithubEvent_Promise struct{ *capnp.Pipeline }

func (p GithubEvent_Promise) Struct() (GithubEvent, error) {
	s, err := p.Pipeline.Struct()
	return GithubEvent{s}, err
}

func (p GithubEvent_Promise) Commit() GithubEvent_commit_Promise {
	return GithubEvent_commit_Promise{p.Pipeline}
}

// GithubEvent_commit_Promise is a wrapper for a GithubEvent_commit promised by a client call.
type GithubEvent_commit_Promise struct{ *capnp.Pipeline }

func (p GithubEvent_commit_Promise) Struct() (GithubEvent_commit, error) {
	s, err := p.Pipeline.Struct()
	return GithubEvent_commit{s}, err
}

func (p GithubEvent_Promise) PullRequest() GithubEvent_pullRequest_Promise {
	return GithubEvent_pullRequest_Promise{p.Pipeline}
}

// GithubEvent_pullRequest_Promise is a wrapper for a GithubEvent_pullRequest promised by a client call.
type GithubEvent_pullRequest_Promise struct{ *capnp.Pipeline }

func (p GithubEvent_pullRequest_Promise) Struct() (GithubEvent_pullRequest, error) {
	s, err := p.Pipeline.Struct()
	return GithubEvent_pullRequest{s}, err
}

func (p GithubEvent_Promise) Issue() GithubEvent_issue_Promise {
	return GithubEvent_issue_Promise{p.Pipeline}
}

// GithubEvent_issue_Promise is a wrapper for a GithubEvent_issue promised by a client call.
type GithubEvent_issue_Promise struct{ *capnp.Pipeline }

func (p GithubEvent_issue_Promise) Struct() (GithubEvent_issue, error) {
	s, err := p.Pipeline.Struct()
	return GithubEvent_issue{s}, err
}

type IrcBot struct{ Client capnp.Client }

// IrcBot_TypeID is the unique identifier for the type IrcBot.
const IrcBot_TypeID = 0xbb6513902c830e39

func (c IrcBot) NoteGhEvent(ctx context.Context, params func(IrcBot_noteGhEvent_Params) error, opts ...capnp.CallOption) IrcBot_noteGhEvent_Results_Promise {
	if c.Client == nil {
		return IrcBot_noteGhEvent_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xbb6513902c830e39,
			MethodID:      0,
			InterfaceName: "basic.capnp:IrcBot",
			MethodName:    "noteGhEvent",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(IrcBot_noteGhEvent_Params{Struct: s}) }
	}
	return IrcBot_noteGhEvent_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type IrcBot_Server interface {
	NoteGhEvent(IrcBot_noteGhEvent) error
}

func IrcBot_ServerToClient(s IrcBot_Server) IrcBot {
	c, _ := s.(server.Closer)
	return IrcBot{Client: server.New(IrcBot_Methods(nil, s), c)}
}

func IrcBot_Methods(methods []server.Method, s IrcBot_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbb6513902c830e39,
			MethodID:      0,
			InterfaceName: "basic.capnp:IrcBot",
			MethodName:    "noteGhEvent",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := IrcBot_noteGhEvent{c, opts, IrcBot_noteGhEvent_Params{Struct: p}, IrcBot_noteGhEvent_Results{Struct: r}}
			return s.NoteGhEvent(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// IrcBot_noteGhEvent holds the arguments for a server call to IrcBot.noteGhEvent.
type IrcBot_noteGhEvent struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  IrcBot_noteGhEvent_Params
	Results IrcBot_noteGhEvent_Results
}

type IrcBot_noteGhEvent_Params struct{ capnp.Struct }

// IrcBot_noteGhEvent_Params_TypeID is the unique identifier for the type IrcBot_noteGhEvent_Params.
const IrcBot_noteGhEvent_Params_TypeID = 0xa5eec3de87bdd251

func NewIrcBot_noteGhEvent_Params(s *capnp.Segment) (IrcBot_noteGhEvent_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IrcBot_noteGhEvent_Params{st}, err
}

func NewRootIrcBot_noteGhEvent_Params(s *capnp.Segment) (IrcBot_noteGhEvent_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return IrcBot_noteGhEvent_Params{st}, err
}

func ReadRootIrcBot_noteGhEvent_Params(msg *capnp.Message) (IrcBot_noteGhEvent_Params, error) {
	root, err := msg.RootPtr()
	return IrcBot_noteGhEvent_Params{root.Struct()}, err
}

func (s IrcBot_noteGhEvent_Params) String() string {
	str, _ := text.Marshal(0xa5eec3de87bdd251, s.Struct)
	return str
}

func (s IrcBot_noteGhEvent_Params) Gh() (GithubEvent, error) {
	p, err := s.Struct.Ptr(0)
	return GithubEvent{Struct: p.Struct()}, err
}

func (s IrcBot_noteGhEvent_Params) HasGh() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s IrcBot_noteGhEvent_Params) SetGh(v GithubEvent) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewGh sets the gh field to a newly
// allocated GithubEvent struct, preferring placement in s's segment.
func (s IrcBot_noteGhEvent_Params) NewGh() (GithubEvent, error) {
	ss, err := NewGithubEvent(s.Struct.Segment())
	if err != nil {
		return GithubEvent{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// IrcBot_noteGhEvent_Params_List is a list of IrcBot_noteGhEvent_Params.
type IrcBot_noteGhEvent_Params_List struct{ capnp.List }

// NewIrcBot_noteGhEvent_Params creates a new list of IrcBot_noteGhEvent_Params.
func NewIrcBot_noteGhEvent_Params_List(s *capnp.Segment, sz int32) (IrcBot_noteGhEvent_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return IrcBot_noteGhEvent_Params_List{l}, err
}

func (s IrcBot_noteGhEvent_Params_List) At(i int) IrcBot_noteGhEvent_Params {
	return IrcBot_noteGhEvent_Params{s.List.Struct(i)}
}

func (s IrcBot_noteGhEvent_Params_List) Set(i int, v IrcBot_noteGhEvent_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s IrcBot_noteGhEvent_Params_List) String() string {
	str, _ := text.MarshalList(0xa5eec3de87bdd251, s.List)
	return str
}

// IrcBot_noteGhEvent_Params_Promise is a wrapper for a IrcBot_noteGhEvent_Params promised by a client call.
type IrcBot_noteGhEvent_Params_Promise struct{ *capnp.Pipeline }

func (p IrcBot_noteGhEvent_Params_Promise) Struct() (IrcBot_noteGhEvent_Params, error) {
	s, err := p.Pipeline.Struct()
	return IrcBot_noteGhEvent_Params{s}, err
}

func (p IrcBot_noteGhEvent_Params_Promise) Gh() GithubEvent_Promise {
	return GithubEvent_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type IrcBot_noteGhEvent_Results struct{ capnp.Struct }

// IrcBot_noteGhEvent_Results_TypeID is the unique identifier for the type IrcBot_noteGhEvent_Results.
const IrcBot_noteGhEvent_Results_TypeID = 0xf5bdbbc1d36794d7

func NewIrcBot_noteGhEvent_Results(s *capnp.Segment) (IrcBot_noteGhEvent_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return IrcBot_noteGhEvent_Results{st}, err
}

func NewRootIrcBot_noteGhEvent_Results(s *capnp.Segment) (IrcBot_noteGhEvent_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return IrcBot_noteGhEvent_Results{st}, err
}

func ReadRootIrcBot_noteGhEvent_Results(msg *capnp.Message) (IrcBot_noteGhEvent_Results, error) {
	root, err := msg.RootPtr()
	return IrcBot_noteGhEvent_Results{root.Struct()}, err
}

func (s IrcBot_noteGhEvent_Results) String() string {
	str, _ := text.Marshal(0xf5bdbbc1d36794d7, s.Struct)
	return str
}

// IrcBot_noteGhEvent_Results_List is a list of IrcBot_noteGhEvent_Results.
type IrcBot_noteGhEvent_Results_List struct{ capnp.List }

// NewIrcBot_noteGhEvent_Results creates a new list of IrcBot_noteGhEvent_Results.
func NewIrcBot_noteGhEvent_Results_List(s *capnp.Segment, sz int32) (IrcBot_noteGhEvent_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return IrcBot_noteGhEvent_Results_List{l}, err
}

func (s IrcBot_noteGhEvent_Results_List) At(i int) IrcBot_noteGhEvent_Results {
	return IrcBot_noteGhEvent_Results{s.List.Struct(i)}
}

func (s IrcBot_noteGhEvent_Results_List) Set(i int, v IrcBot_noteGhEvent_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s IrcBot_noteGhEvent_Results_List) String() string {
	str, _ := text.MarshalList(0xf5bdbbc1d36794d7, s.List)
	return str
}

// IrcBot_noteGhEvent_Results_Promise is a wrapper for a IrcBot_noteGhEvent_Results promised by a client call.
type IrcBot_noteGhEvent_Results_Promise struct{ *capnp.Pipeline }

func (p IrcBot_noteGhEvent_Results_Promise) Struct() (IrcBot_noteGhEvent_Results, error) {
	s, err := p.Pipeline.Struct()
	return IrcBot_noteGhEvent_Results{s}, err
}

const schema_902d2bfc3bbf0d43 = "x\xda\xa4\xd2?hSQ\x14\x06\xf0\xef\xbb\xf7\xa5\xaf" +
	"\xd5Ds}\x11\xa4K\xa1\x88\x83m\x83\xb6\x086\"" +
	")j\xa8E\x0b\xb9EAP\x87\x97\xf0h\x02\xf9S" +
	"\xf3^tR\xd0*\x828t\xd0AA\x07\xc1Q;" +
	"8\x95R\xc4:\x8bT\x117\x07\x177\x07EDQ" +
	"\x9f\\k^\"\xa4\x93[\xee\xcd\xe1\xdc\xdfw\xce\xdb" +
	"S\xe1\x84\xd8\x1b{)\x01=\x1c\xeb\x09\xaf\x1e\xdb\xb6" +
	"i5>t\x07:N\x11\x1eN<=\xf0chd" +
	"\x019\xcb\x96\x80\xf3\x82\xab\xce\x1b\xda\x80\xb3\xc6\x0f`" +
	"\xa8_\xad\\\x7f\xf7\xfc\xe3C(\x87@\xcc\xfc3v" +
	"R\x8c\x12t\xce\x8a,\x18\x8eo\x99\x1f^p\xbce" +
	"\xa8\xb8lw\x03\x9d\x8b\xe2\xaesM\x98N\x97\xc5\xa4" +
	"\xf3\xc8\xfc\x0a\x1f\xdfp\x0e\xbe\xbf_^\x83\xdeA\xd1" +
	"\x96l\xb7l\x02c\xb7\xc5\xa0\xe9{O,\x82\xe1\x93" +
	"\x9b\xa7No\xde\xf5\xe0s\xd7Rg\x9f\xfc\x04:\xe3" +
	"\xd2T\xbe\xbd5\xfb\xfa\xd9\xf2\xca\x97u\xa2e\x84k" +
	"2CX\xe1\xa5+\xf3\xfb\x97\xbe\xf5\x7f\xed\xdecI" +
	"~\x07\x9d\x15\xb9\x88\x91\xb0\xe0\xfa\xe5b\xba\xe8\x8a\xb9" +
	"\xda\\f\xb2\x1c\x94\x9a\x85\xdcy\xaf\x16 O\xea\xa4" +
	"\xb4\x00\x8b\x80r\x07\x01}FR\x97\x04\x13\x0cC\xb2" +
	"\xfd\x86\xf22\x10\x09\xf1\xcb\\F9\xd5t\x01\"!" +
	"\x7f\x9a\xcb(\x91\x1a\x1f\x85\xb0/\x94\xea\x8cC0\x0e" +
	"f\x8b\xf5j\xb5\x1c\x84s\xcdJe\xc6;\xd7\x84\xed" +
	"\xf9\xc1@\xd9\xf7\x9b^$\x93F6\xd5(\x1e\xaa\x07" +
	"\xe9Z=\xf0&K\x7f\x80;\xf3\x03n\xc3\xad\xfa\xda" +
	"\x8a\x90\x89~@\xf7J\xea\x94\xa0\x9c-1\xd9N\x0e" +
	"2\x09F=\xd9\xea)\xeb\x81\x09j\xc9\x18\x10\xad\x9c" +
	"\xad\xc1*U\x80P}v\xd8z\x17\xb6W\x0b&\x98" +
	"'\xff\xe5u\x0c.\xdd\xcab\xa2\xe8\xb8\xb4\x92L\xd1" +
	"\x02T.\x03\xe8\x09I}\\P\x91)\xc6\x005e" +
	"\xe6zDR\xe7\x05\x95\x10)\xf6\x00j\xdaT\x1e\x95" +
	"\xd4'\x04\xb3\xb5f\xb5\xe05\x18\x83`\x0c\xb4\xab\xfe" +
	"l4<\xb7\x18\x94\xeb\xb5\xd6q\xc3U\xa6\xd7\x07\xda" +
	"\xb2\xd8\xdd,\xbd\xdd,}\xffo\xd9hy3Y\xcf" +
	"oV\x02\x7fc\xf4\xdfO\xa3\xa56\xfb\xcd\xed\xeeT" +
	"\x8b\x14\x85Qg:\xd52Ei\xd4\x83m\xf5\xd6\x92" +
	"\xeb\x97\"g\xa1\xe1\xd6\x8a\xd1\xb13\xc2\xef\x00\x00\x00" +
	"\xff\xff\x16\xa7\x0e\x1c"

func init() {
	schemas.Register(schema_902d2bfc3bbf0d43,
		0x992b0cc20a124b84,
		0xa5eec3de87bdd251,
		0xbb6513902c830e39,
		0xd1699ee23d138aae,
		0xf3a2260b5b588cb3,
		0xf5bdbbc1d36794d7,
		0xf71af9b93883827e)
}