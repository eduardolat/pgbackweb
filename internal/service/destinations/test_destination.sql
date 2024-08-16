-- name: DestinationsServiceSetTestData :exec
UPDATE destinations
SET test_ok = @test_ok,
    test_error = @test_error,
    last_test_at = NOW()
WHERE id = @destination_id;
