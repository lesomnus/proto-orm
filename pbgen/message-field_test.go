package pbgen_test

import (
	"bytes"
	"testing"

	"github.com/lesomnus/proto-orm/pbgen"
	"github.com/stretchr/testify/require"
)

func TestMessageField(t *testing.T) {
	t.Run("normal field", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.MessageField{
			Type:   pbgen.TypeBytes,
			Name:   "id",
			Number: 1,
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal("bytes id = 1;", v)
	})

	t.Run("labels", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.MessageField{
			Label:  pbgen.LabelOptional,
			Type:   pbgen.TypeBytes,
			Name:   "id",
			Number: 1,
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal("optional bytes id = 1;", v)
	})

	t.Run("full ident", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.MessageField{
			Type:   "foo.bar.baz",
			Name:   "id",
			Number: 1,
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal("foo.bar.baz id = 1;", v)
	})

	t.Run("single option", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.MessageField{
			Type:    pbgen.TypeBytes,
			Name:    "id",
			Number:  1,
			Options: []pbgen.Option{pbgen.FeatureFieldPresenceExplicit},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`bytes id = 1 [features.field_presence = EXPLICIT];`, v)
	})

	t.Run("multiple options", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.MessageField{
			Type:   pbgen.TypeBytes,
			Name:   "id",
			Number: 1,
			Options: []pbgen.Option{
				pbgen.FeatureFieldPresenceLegacyRequired,
				pbgen.FeatureFieldPresenceExplicit,
				pbgen.FeatureFieldPresenceImplicit,
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`bytes id = 1 [
	features.field_presence = LEGACY_REQUIRED,
	features.field_presence = EXPLICIT,
	features.field_presence = IMPLICIT
];`, v)
	})

	t.Run("all", func(t *testing.T) {
		require := require.New(t)

		d := pbgen.MessageField{
			Label:  pbgen.LabelOptional,
			Type:   "foo.bar.baz",
			Name:   "id",
			Number: 1,
			Options: []pbgen.Option{
				pbgen.FeatureFieldPresenceLegacyRequired,
				pbgen.FeatureFieldPresenceExplicit,
				pbgen.FeatureFieldPresenceImplicit,
			},
		}
		o := bytes.Buffer{}
		err := pbgen.Execute(&o, &d)
		require.NoError(err)

		v := o.String()
		require.Equal(`optional foo.bar.baz id = 1 [
	features.field_presence = LEGACY_REQUIRED,
	features.field_presence = EXPLICIT,
	features.field_presence = IMPLICIT
];`, v)
	})
}
