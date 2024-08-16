-- name: DatabasesServiceSetTestData :exec
UPDATE databases
SET test_ok = @test_ok,
    test_error = @test_error,
    last_test_at = NOW()
WHERE id = @database_id;
