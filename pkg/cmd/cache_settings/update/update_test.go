package update

import (
	"fmt"
	"github.com/aziontech/azion-cli/pkg/logger"
	"go.uber.org/zap/zapcore"
	"testing"

	msg "github.com/aziontech/azion-cli/messages/cache_settings"
	"github.com/aziontech/azion-cli/pkg/httpmock"
	"github.com/aziontech/azion-cli/pkg/testutils"
	"github.com/aziontech/azion-cli/utils"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	logger.New(zapcore.DebugLevel)
	t.Run("create new cache_setting", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, stdout, _ := testutils.NewFactory(mock)
		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"-a", "1673635841",
			"-c", "112233",
			"--name", "fmaiswaybetter",
			"--adaptive-delivery-action", "ignore",
			"--browser-cache-settings", "override",
			"--browser-cache-settings-maximum-ttl", "60",
			"--cdn-cache-settings", "honor",
			"--cnd-cache-settings-maximum-ttl", "60",
			"--cache-by-cookies", "whitelist",
			"--cookie-names", "aa,123,987",
			"--cache-by-query-string", "whitelist",
			"--query-string-fields", "heyy,yoo",
			"--enable-caching-for-options", "true",
			"--enable-caching-for-post", "true",
			"--enable-caching-string-sort", "true",
			"--l2-caching-enabled", "true",
			"--slice-configuration-enabled", "false",
			"--slice-l2-caching-enabled", "false",
		})

		err := cmd.Execute()
		require.NoError(t, err)
		require.Equal(t, fmt.Sprintf(msg.CacheSettingsUpdateOutputSuccess, 112233), stdout.String())
	})

	t.Run("create with file", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, stdout, _ := testutils.NewFactory(mock)

		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"--application-id", "1673635841",
			"--in", "./fixtures/update.json",
		})

		err := cmd.Execute()
		require.NoError(t, err)
		require.Equal(t, fmt.Sprintf(msg.CacheSettingsUpdateOutputSuccess, 112233), stdout.String())
	})

	t.Run("no acceleration error --in flag", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result_no_acceleration.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, _, _ := testutils.NewFactory(mock)

		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"--application-id", "1673635841",
			"--in", "./fixtures/update.json",
		})

		err := cmd.Execute()
		require.ErrorIs(t, err, msg.ErrorApplicationAccelerationNotEnabled)
	})

	t.Run("missing one mandatory flag", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result_no_acceleration.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, _, _ := testutils.NewFactory(mock)

		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"--application-id", "1673635841",
		})

		err := cmd.Execute()
		require.ErrorIs(t, err, msg.ErrorMandatoryUpdateFlags)
	})

	t.Run("no acceleration error with args", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result_no_acceleration.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, _, _ := testutils.NewFactory(mock)

		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"-a", "1673635841",
			"-c", "112233",
			"--name", "fmaiswaybetter",
			"--adaptive-delivery-action", "ignore",
			"--browser-cache-settings", "override",
			"--browser-cache-settings-maximum-ttl", "60",
			"--cdn-cache-settings", "honor",
			"--cnd-cache-settings-maximum-ttl", "60",
			"--cache-by-cookies", "whitelist",
			"--cookie-names", "aa,123,987",
			"--cache-by-query-string", "whitelist",
			"--query-string-fields", "heyy,yoo",
			"--enable-caching-for-options", "true",
			"--enable-caching-for-post", "true",
			"--enable-caching-string-sort", "true",
			"--l2-caching-enabled", "true",
			"--slice-configuration-enabled", "false",
			"--slice-l2-caching-enabled", "false",
		})

		err := cmd.Execute()
		require.ErrorIs(t, err, msg.ErrorApplicationAccelerationNotEnabled)
	})

	t.Run("override but no ttl", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, _, _ := testutils.NewFactory(mock)

		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"-a", "1673635841",
			"-c", "112233",
			"--name", "fmaiswaybetter",
			"--adaptive-delivery-action", "ignore",
			"--browser-cache-settings", "override",
			"--cdn-cache-settings", "honor",
			"--cnd-cache-settings-maximum-ttl", "60",
			"--cache-by-cookies", "whitelist",
			"--cookie-names", "aa,123,987",
			"--cache-by-query-string", "whitelist",
			"--query-string-fields", "heyy,yoo",
			"--enable-caching-for-options", "true",
			"--enable-caching-for-post", "true",
			"--enable-caching-string-sort", "true",
			"--l2-caching-enabled", "true",
			"--slice-configuration-enabled", "false",
			"--slice-l2-caching-enabled", "false",
		})

		err := cmd.Execute()
		require.ErrorIs(t, err, msg.ErrorBrowserMaximumTtlNotSent)
	})

	t.Run("no acceleration error with args", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result_no_acceleration.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, _, _ := testutils.NewFactory(mock)

		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"-a", "1673635841",
			"-c", "112233",
			"--name", "fmaiswaybetter",
			"--adaptive-delivery-action", "ignore",
			"--browser-cache-settings", "override",
			"--browser-cache-settings-maximum-ttl", "60",
			"--cdn-cache-settings", "honor",
			"--cnd-cache-settings-maximum-ttl", "60",
			"--cache-by-cookies", "whitelist",
			"--cookie-names", "aa,123,987",
			"--cache-by-query-string", "whitelist",
			"--query-string-fields", "heyy,yoo",
			"--enable-caching-for-options", "true",
			"--enable-caching-for-post", "true",
			"--enable-caching-string-sort", "true",
			"--l2-caching-enabled", "true",
			"--slice-configuration-enabled", "false",
			"--slice-l2-caching-enabled", "false",
		})

		err := cmd.Execute()
		require.ErrorIs(t, err, msg.ErrorApplicationAccelerationNotEnabled)
	})

	t.Run("wrong l2 boolean var", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, _, _ := testutils.NewFactory(mock)

		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"-a", "1673635841",
			"-c", "112233",
			"--name", "fmaiswaybetter",
			"--adaptive-delivery-action", "ignore",
			"--browser-cache-settings", "override",
			"--cdn-cache-settings", "honor",
			"--cnd-cache-settings-maximum-ttl", "60",
			"--browser-cache-settings-maximum-ttl", "60",
			"--cache-by-cookies", "whitelist",
			"--cookie-names", "aa,123,987",
			"--cache-by-query-string", "whitelist",
			"--query-string-fields", "heyy,yoo",
			"--enable-caching-for-options", "true",
			"--enable-caching-for-post", "true",
			"--enable-caching-string-sort", "true",
			"--l2-caching-enabled", "true",
			"--slice-configuration-enabled", "false",
			"--slice-l2-caching-enabled", "troo",
		})

		err := cmd.Execute()
		require.ErrorIs(t, err, msg.ErrorSliceL2CachingFlag)
	})

	t.Run("wrong caching for options boolean var", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, _, _ := testutils.NewFactory(mock)

		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"-a", "1673635841",
			"-c", "112233",
			"--name", "fmaiswaybetter",
			"--adaptive-delivery-action", "ignore",
			"--browser-cache-settings", "override",
			"--cdn-cache-settings", "honor",
			"--cnd-cache-settings-maximum-ttl", "60",
			"--browser-cache-settings-maximum-ttl", "60",
			"--cache-by-cookies", "whitelist",
			"--cookie-names", "aa,123,987",
			"--cache-by-query-string", "whitelist",
			"--query-string-fields", "heyy,yoo",
			"--enable-caching-for-options", "untrue",
			"--enable-caching-for-post", "true",
			"--enable-caching-string-sort", "true",
			"--l2-caching-enabled", "true",
			"--slice-configuration-enabled", "false",
			"--slice-l2-caching-enabled", "false",
		})

		err := cmd.Execute()
		require.ErrorIs(t, err, msg.ErrorCachingForOptionsFlag)
	})

	t.Run("wrong caching for post boolean var", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, _, _ := testutils.NewFactory(mock)

		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"-a", "1673635841",
			"-c", "112233",
			"--name", "fmaiswaybetter",
			"--adaptive-delivery-action", "ignore",
			"--browser-cache-settings", "override",
			"--cdn-cache-settings", "honor",
			"--cnd-cache-settings-maximum-ttl", "60",
			"--browser-cache-settings-maximum-ttl", "60",
			"--cache-by-cookies", "whitelist",
			"--cookie-names", "aa,123,987",
			"--cache-by-query-string", "whitelist",
			"--query-string-fields", "heyy,yoo",
			"--enable-caching-for-options", "true",
			"--enable-caching-for-post", "incorrect",
			"--enable-caching-string-sort", "true",
			"--l2-caching-enabled", "true",
			"--slice-configuration-enabled", "false",
			"--slice-l2-caching-enabled", "false",
		})

		err := cmd.Execute()
		require.ErrorIs(t, err, msg.ErrorCachingForPostFlag)
	})

	t.Run("wrong caching string sort boolean var", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, _, _ := testutils.NewFactory(mock)

		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"-a", "1673635841",
			"-c", "112233",
			"--name", "fmaiswaybetter",
			"--adaptive-delivery-action", "ignore",
			"--browser-cache-settings", "override",
			"--cdn-cache-settings", "honor",
			"--cnd-cache-settings-maximum-ttl", "60",
			"--browser-cache-settings-maximum-ttl", "60",
			"--cache-by-cookies", "whitelist",
			"--cookie-names", "aa,123,987",
			"--cache-by-query-string", "whitelist",
			"--query-string-fields", "heyy,yoo",
			"--enable-caching-for-options", "false",
			"--enable-caching-for-post", "false",
			"--enable-caching-string-sort", "precise",
			"--l2-caching-enabled", "true",
			"--slice-configuration-enabled", "false",
			"--slice-l2-caching-enabled", "false",
		})

		err := cmd.Execute()
		require.ErrorIs(t, err, msg.ErrorCachingStringSortFlag)
	})

	t.Run("wrong slice configuration enabled boolean var", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, _, _ := testutils.NewFactory(mock)

		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"-a", "1673635841",
			"-c", "112233",
			"--name", "fmaiswaybetter",
			"--adaptive-delivery-action", "ignore",
			"--browser-cache-settings", "override",
			"--cdn-cache-settings", "honor",
			"--cnd-cache-settings-maximum-ttl", "60",
			"--browser-cache-settings-maximum-ttl", "60",
			"--cache-by-cookies", "whitelist",
			"--cookie-names", "aa,123,987",
			"--cache-by-query-string", "whitelist",
			"--query-string-fields", "heyy,yoo",
			"--enable-caching-for-options", "true",
			"--enable-caching-for-post", "true",
			"--enable-caching-string-sort", "true",
			"--l2-caching-enabled", "true",
			"--slice-configuration-enabled", "faithful",
			"--slice-l2-caching-enabled", "false",
		})

		err := cmd.Execute()
		require.ErrorIs(t, err, msg.ErrorSliceConfigurationFlag)
	})

	t.Run("wrong slice l2 caching enabled boolean var", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, _, _ := testutils.NewFactory(mock)

		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"-a", "1673635841",
			"-c", "112233",
			"--name", "fmaiswaybetter",
			"--adaptive-delivery-action", "ignore",
			"--browser-cache-settings", "override",
			"--cdn-cache-settings", "honor",
			"--cnd-cache-settings-maximum-ttl", "60",
			"--browser-cache-settings-maximum-ttl", "60",
			"--cache-by-cookies", "whitelist",
			"--cookie-names", "aa,123,987",
			"--cache-by-query-string", "whitelist",
			"--query-string-fields", "heyy,yoo",
			"--enable-caching-for-options", "true",
			"--enable-caching-for-post", "true",
			"--enable-caching-string-sort", "true",
			"--l2-caching-enabled", "true",
			"--slice-configuration-enabled", "true",
			"--slice-l2-caching-enabled", "erroneous",
		})

		err := cmd.Execute()
		require.ErrorIs(t, err, msg.ErrorSliceL2CachingFlag)
	})

	t.Run("error unmarshall", func(t *testing.T) {
		mock := &httpmock.Registry{}

		mock.Register(
			httpmock.REST("GET", "edge_applications/1673635841"),
			httpmock.JSONFromFile("./fixtures/app_result.json"),
		)

		mock.Register(
			httpmock.REST("PATCH", "edge_applications/1673635841/cache_settings/112233"),
			httpmock.JSONFromFile("./fixtures/result.json"),
		)

		f, _, _ := testutils.NewFactory(mock)

		cmd := NewCmd(f)
		cmd.SetArgs([]string{
			"--application-id", "1673635841",
			"--in", "./fixtures/error",
		})

		err := cmd.Execute()
		require.ErrorIs(t, err, utils.ErrorUnmarshalReader)
	})

}
