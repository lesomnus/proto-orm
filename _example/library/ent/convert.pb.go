// Code generated by "protoc-gen-orm-ent". DO NOT EDIT

package ent

import (
	library "github.com/lesomnus/proto-orm/_example/library"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (e *Author) Proto() *library.Author {
	m := &library.Author{}
	m.Id = e.ID[:]
	m.Alias = e.Alias
	m.Name = e.Name
	for _, v := range e.Edges.Books {
		m.Books = append(m.Books, v.Proto())
	}
	m.DateCreated = timestamppb.New(e.DateCreated)

	return m
}

func (e *Book) Proto() *library.Book {
	m := &library.Book{}
	m.Id = e.ID[:]
	m.Title = e.Title
	for _, v := range e.Edges.Authors {
		m.Authors = append(m.Authors, v.Proto())
	}
	m.Index = e.Index
	m.Genres = e.Genres
	m.DateCreated = timestamppb.New(e.DateCreated)

	return m
}

func (e *Like) Proto() *library.Like {
	m := &library.Like{}
	m.Id = e.ID[:]
	if v := e.Edges.Book; v != nil {
		m.Book = v.Proto()
	}
	if v := e.Edges.Member; v != nil {
		m.Member = v.Proto()
	}
	m.DateCreated = timestamppb.New(e.DateCreated)

	return m
}

func (e *Loan) Proto() *library.Loan {
	m := &library.Loan{}
	m.Id = e.ID[:]
	m.BookId = e.BookID[:]
	if v := e.Edges.Book; v != nil {
		m.Book = v.Proto()
	}
	m.MemberId = e.MemberID[:]
	if v := e.Edges.Member; v != nil {
		m.Member = v.Proto()
	}
	m.DateDue = timestamppb.New(e.DateDue)
	if v := e.DateReturn; v != nil {
		m.DateReturn = timestamppb.New(*e.DateReturn)
	}
	m.DateCreated = timestamppb.New(e.DateCreated)

	return m
}

func (e *Member) Proto() *library.Member {
	m := &library.Member{}
	m.Id = e.ID[:]
	m.Name = e.Name
	m.Level = library.Member_Level(e.Level)
	m.Labels = e.Labels
	m.DateCreated = timestamppb.New(e.DateCreated)

	return m
}