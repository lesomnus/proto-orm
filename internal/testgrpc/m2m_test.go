package testgrpc_test

import (
	"context"
	"testing"

	orm "github.com/lesomnus/proto-orm"
	"github.com/lesomnus/proto-orm/internal/example/library"
	"github.com/stretchr/testify/require"
)

func TestM2mPatch(t *testing.T) {
	var (
		author1 *library.Author
		author2 *library.Author
		author3 *library.Author

		book1 *library.Book
		book2 *library.Book
	)

	setup := func(ctx context.Context, x *require.Assertions, c *Client) {
		var err error
		author1, err = c.Author.Add(ctx, library.AuthorAddRequest_builder{
			Alias: "author1",
			Name:  "author1",
		}.Build())
		x.NoError(err)
		author2, err = c.Author.Add(ctx, library.AuthorAddRequest_builder{
			Alias: "author2",
			Name:  "author2",
		}.Build())
		x.NoError(err)
		author3, err = c.Author.Add(ctx, library.AuthorAddRequest_builder{
			Alias: "author3",
			Name:  "author3",
		}.Build())
		x.NoError(err)

		book1, err = c.Book.Add(ctx, library.BookAddRequest_builder{Title: "book1"}.Build())
		x.NoError(err)
		book2, err = c.Book.Add(ctx, library.BookAddRequest_builder{Title: "book2"}.Build())
		x.NoError(err)
	}

	t.Run("add by secondary key", T(func(ctx context.Context, x *require.Assertions, c *Client) {
		setup(ctx, x, c)

		_, err := c.Book.Patch(ctx, library.BookPatchRequest_builder{
			Key: book1.Pick(),
			Authors: []*library.BookPatchAuthors{
				library.BookPatchAuthors_builder{
					Op: orm.PatchOp_ADD,
					Items: []*library.AuthorGetRequest{
						library.AuthorByAlias(author1.GetAlias()),
						library.AuthorByAlias(author2.GetAlias()),
					},
				}.Build(),
			},
		}.Build())
		x.NoError(err)
		{
			v, err := c.Book.Get(ctx, book1.Pick())
			x.NoError(err)

			authors := v.GetAuthors()
			x.Len(authors, 2)
			x.ElementsMatch(
				[][]byte{
					author1.GetId(),
					author2.GetId(),
				},
				[][]byte{
					authors[0].GetId(),
					authors[1].GetId(),
				},
			)
		}
	}))
	t.Run("add and erase", T(func(ctx context.Context, x *require.Assertions, c *Client) {
		setup(ctx, x, c)

		_, err := c.Book.Patch(ctx, library.BookPatchRequest_builder{
			Key: book1.Pick(),
			Authors: []*library.BookPatchAuthors{
				library.BookPatchAuthors_builder{
					Op: orm.PatchOp_ADD,
					Items: []*library.AuthorGetRequest{
						author1.Pick(),
						author2.Pick(),
					},
				}.Build(),
			},
		}.Build())
		x.NoError(err)
		{
			v, err := c.Book.Get(ctx, book1.Pick())
			x.NoError(err)

			authors := v.GetAuthors()
			x.Len(authors, 2)
			x.ElementsMatch(
				[][]byte{
					author1.GetId(),
					author2.GetId(),
				},
				[][]byte{
					authors[0].GetId(),
					authors[1].GetId(),
				},
			)
		}

		_, err = c.Book.Patch(ctx, library.BookPatchRequest_builder{
			Key: book1.Pick(),
			Authors: []*library.BookPatchAuthors{
				library.BookPatchAuthors_builder{
					Op: orm.PatchOp_ERASE,
					Items: []*library.AuthorGetRequest{
						author1.Pick(),
					},
				}.Build(),
			},
		}.Build())
		x.NoError(err)
		{
			v, err := c.Book.Get(ctx, book1.Pick())
			x.NoError(err)

			authors := v.GetAuthors()
			x.Len(authors, 1)
			x.Equal(author2.GetId(), authors[0].GetId())
		}
	}))
	t.Run("clear", T(func(ctx context.Context, x *require.Assertions, c *Client) {
		setup(ctx, x, c)

		_, err := c.Book.Patch(ctx, library.BookPatchRequest_builder{
			Key: book1.Pick(),
			Authors: []*library.BookPatchAuthors{
				library.BookPatchAuthors_builder{
					Op: orm.PatchOp_ADD,
					Items: []*library.AuthorGetRequest{
						author1.Pick(),
						author2.Pick(),
					},
				}.Build(),
			},
		}.Build())
		x.NoError(err)
		{
			v, err := c.Book.Get(ctx, book1.Pick())
			x.NoError(err)

			authors := v.GetAuthors()
			x.Len(authors, 2)
		}

		_, err = c.Book.Patch(ctx, library.BookPatchRequest_builder{
			Key: book1.Pick(),
			Authors: []*library.BookPatchAuthors{
				library.BookPatchAuthors_builder{
					Op: orm.PatchOp_CLEAR,
				}.Build(),
			},
		}.Build())
		x.NoError(err)
		{
			v, err := c.Book.Get(ctx, book1.Pick())
			x.NoError(err)

			authors := v.GetAuthors()
			x.Empty(authors)
		}
	}))
	t.Run("shared", T(func(ctx context.Context, x *require.Assertions, c *Client) {
		setup(ctx, x, c)

		_, err := c.Book.Patch(ctx, library.BookPatchRequest_builder{
			Key: book1.Pick(),
			Authors: []*library.BookPatchAuthors{
				library.BookPatchAuthors_builder{
					Op: orm.PatchOp_ADD,
					Items: []*library.AuthorGetRequest{
						author1.Pick(),
						author2.Pick(),
					},
				}.Build(),
			},
		}.Build())
		x.NoError(err)
		{
			v, err := c.Book.Get(ctx, book1.Pick())
			x.NoError(err)

			authors := v.GetAuthors()
			x.Len(authors, 2)
			x.ElementsMatch(
				[][]byte{
					author1.GetId(),
					author2.GetId(),
				},
				[][]byte{
					authors[0].GetId(),
					authors[1].GetId(),
				},
			)
		}

		_, err = c.Book.Patch(ctx, library.BookPatchRequest_builder{
			Key: book2.Pick(),
			Authors: []*library.BookPatchAuthors{
				library.BookPatchAuthors_builder{
					Op: orm.PatchOp_ADD,
					Items: []*library.AuthorGetRequest{
						author2.Pick(),
						author3.Pick(),
					},
				}.Build(),
			},
		}.Build())
		x.NoError(err)
		{
			v, err := c.Book.Get(ctx, book2.Pick())
			x.NoError(err)

			authors := v.GetAuthors()
			x.Len(authors, 2)
			x.ElementsMatch(
				[][]byte{
					author2.GetId(),
					author3.GetId(),
				},
				[][]byte{
					authors[0].GetId(),
					authors[1].GetId(),
				},
			)
		}

	}))
}
