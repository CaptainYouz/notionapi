package main

const (
	// PermissionTypeUser describes permissions for a user
	PermissionTypeUser = "user_permission"
	// PermissionTypePublic describes permissions for public
	PermissionTypePublic = "public_permission"
)

const (
	// TypePage is a notion Page
	TypePage = "page"
	// TypeText is a text block
	TypeText = "text"
	// TypeBookmark is a bookmark block
	TypeBookmark = "bookmark"
	// TypeGist is a gist block
	TypeGist = "gist"
	// TypeBulletedList is a bulleted list block
	TypeBulletedList = "bulleted_list"
	// TypeNumberedList is a numbered list block
	TypeNumberedList = "numbered_list"
	// TypeToggle is a toggle block
	TypeToggle = "toggle"
	// TypeTodo is a todo block
	TypeTodo = "to_do"
	// TypeDivider is a divider block
	TypeDivider = "divider"
	// TypeImage is an image block
	TypeImage = "image"
	// TypeHeader is a header block
	TypeHeader = "header"
	// TypeSubHeader is a header block
	TypeSubHeader = "sub_header"
	// TypeQuote is a quote block
	TypeQuote = "quote"
	// TypeComment is a comment block
	TypeComment = "comment"
	// TypeCode is a code block
	TypeCode = "code"
	// TypeTable is a table block
	TypeTable = "table"
	// TypeCollectionView is a collection view block
	TypeCollectionView = "collection_view"
)

const (
	// TableSpace represents a Notion workspace
	TableSpace = "space"
	// TableBlock represents a Notion block
	TableBlock = "block"
)

const (
	// RoleReader represents a reader
	RoleReader = "reader"
	// RoleEditor represents an editor
	RoleEditor = "editor"
)