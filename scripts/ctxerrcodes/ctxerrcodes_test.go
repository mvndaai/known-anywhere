package main

import (
	"testing"
)

func TestEverything(t *testing.T) {
	t.Skip("The test does't matter, just the diff. Uncomment and reun the script to see changes")

	///* Check invalid*/
	//ctx := context.Background()
	//ctxerr.New(ctx, "invalid-uuid0", "invalid uuid")
	//ctxerr.Newf(ctx, "invalid-uuid1", "invalid %s", "uuid")
	//ctxerr.NewHTTP(ctx, "invalid-uuid2", "invalid uuid", http.StatusBadRequest)
	//err := ctxerr.NewHTTPf(ctx, "invalid-uuid3", "invalid uuid", http.StatusBadRequest, "missing %s", "uuid")
	//ctxerr.Wrap(ctx, err, "invalid-uuid0", "invalid uuid")
	//ctxerr.Wrapf(ctx, err, "invalid-uuid1", "invalid %s", "uuid")
	//ctxerr.WrapHTTP(ctx, err, "invalid-uuid2", "invalid uuid", http.StatusBadRequest)
	//ctxerr.WrapHTTPf(ctx, err, "invalid-uuid3", "invalid uuid", http.StatusBadRequest, "missing %s", "uuid")

	///* Check valid duplicates */
	//ctxerr.New(ctx, "cf85d8c4-2f3d-44eb-8c10-c3ee95ba865a", "valid uuid0")
	//ctxerr.New(ctx, "cf85d8c4-2f3d-44eb-8c10-c3ee95ba865a", "valid uuid1")
	//ctxerr.Newf(ctx, "cf85d8c4-2f3d-44eb-8c10-c3ee95ba865a", "valid %s2", "uuid")
	//ctxerr.NewHTTP(ctx, "cf85d8c4-2f3d-44eb-8c10-c3ee95ba865a", "valid uuid3", http.StatusBadRequest)
	//err = ctxerr.NewHTTPf(ctx, "cf85d8c4-2f3d-44eb-8c10-c3ee95ba865a", "valid uuid4", http.StatusBadRequest, "missing %s", "uuid")
	//ctxerr.Wrap(ctx, err, "cf85d8c4-2f3d-44eb-8c10-c3ee95ba865a", "valid uuid5")
	//ctxerr.Wrapf(ctx, err, "cf85d8c4-2f3d-44eb-8c10-c3ee95ba865a", "valid %s6", "uuid")
	//ctxerr.WrapHTTP(ctx, err, "cf85d8c4-2f3d-44eb-8c10-c3ee95ba865a", "valid uuid7", http.StatusBadRequest)
	//ctxerr.WrapHTTPf(ctx, err, "cf85d8c4-2f3d-44eb-8c10-c3ee95ba865a", "valid uuid8", http.StatusBadRequest, "missing %s", "uuid")

	///* Check empty codes */
	//ctxerr.New(ctx, "", "empty uuid")
	//ctxerr.Newf(ctx, "", "empty %s", "uuid")
	//ctxerr.NewHTTP(ctx, "", "empty uuid", http.StatusBadRequest)
	//err = ctxerr.NewHTTPf(ctx, "", "empty uuid", http.StatusBadRequest, "missing %s", "uuid")
	//ctxerr.Wrap(ctx, err, "", "empty uuid")
	//ctxerr.Wrapf(ctx, err, "", "empty %s", "uuid")
	//ctxerr.WrapHTTP(ctx, err, "", "empty uuid", http.StatusBadRequest)
	//ctxerr.WrapHTTPf(ctx, err, "", "empty uuid", http.StatusBadRequest, "missing %s", "uuid")

	///* Check when not alone */
	//func() error {
	//	return ctxerr.New(context.Background(), "invalid-uuid", "invalid uuid")
	//}()

	//func() (string, error) {
	//	err := ctxerr.New(ctx, "invalid-uuid", "invalid uuid")
	//	return "", ctxerr.WrapHTTPf(ctx, err, "", "action", http.StatusBadRequest, "missing %s", "uuid")
	//}()
}
