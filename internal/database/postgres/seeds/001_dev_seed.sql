INSERT INTO users (id, email, password_hash, role, created_at, updated_at)
VALUES
    (
        '11111111-1111-1111-1111-111111111111',
        'alice@example.com',
        '$2a$10$r0PwbM8zhx6qhmQfXIK53ewxP0gJf31M6fJY6M4ZaChkCGbdJrY8.',
        'user',
        NOW() - INTERVAL '3 days',
        NOW() - INTERVAL '3 days'
    ),
    (
        '22222222-2222-2222-2222-222222222222',
        'bob@example.com',
        '$2a$10$r0PwbM8zhx6qhmQfXIK53ewxP0gJf31M6fJY6M4ZaChkCGbdJrY8.',
        'user',
        NOW() - INTERVAL '2 days',
        NOW() - INTERVAL '2 days'
    ),
    (
        '33333333-3333-3333-3333-333333333333',
        'admin@example.com',
        '$2a$10$r0PwbM8zhx6qhmQfXIK53ewxP0gJf31M6fJY6M4ZaChkCGbdJrY8.',
        'admin',
        NOW() - INTERVAL '1 day',
        NOW() - INTERVAL '1 day'
    )
ON CONFLICT (id) DO NOTHING;

INSERT INTO posts (id, user_id, content, created_at, updated_at)
VALUES
    (
        'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1',
        '11111111-1111-1111-1111-111111111111',
        'Building this backend to learn production-ready Go, Postgres, and API design.',
        NOW() - INTERVAL '36 hours',
        NOW() - INTERVAL '36 hours'
    ),
    (
        'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa2',
        '22222222-2222-2222-2222-222222222222',
        'Today I switched from Firestore to Postgres so I can understand relational modeling deeply.',
        NOW() - INTERVAL '18 hours',
        NOW() - INTERVAL '18 hours'
    ),
    (
        'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa3',
        '11111111-1111-1111-1111-111111111111',
        'Next up is authentication with access and refresh tokens backed by the database.',
        NOW() - INTERVAL '6 hours',
        NOW() - INTERVAL '6 hours'
    )
ON CONFLICT (id) DO NOTHING;

INSERT INTO comments (id, post_id, user_id, parent_id, content, created_at, updated_at)
VALUES
    (
        'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb1',
        'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1',
        '22222222-2222-2222-2222-222222222222',
        NULL,
        'Nice direction. Postgres will teach you a lot about schema and query design.',
        NOW() - INTERVAL '30 hours',
        NOW() - INTERVAL '30 hours'
    ),
    (
        'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb2',
        'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa1',
        '11111111-1111-1111-1111-111111111111',
        'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb1',
        'That is exactly the goal. I want stronger backend fundamentals.',
        NOW() - INTERVAL '28 hours',
        NOW() - INTERVAL '28 hours'
    ),
    (
        'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbb3',
        'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaa2',
        '33333333-3333-3333-3333-333333333333',
        NULL,
        'Good call. Add pagination and indexes early and the app will age much better.',
        NOW() - INTERVAL '12 hours',
        NOW() - INTERVAL '12 hours'
    )
ON CONFLICT (id) DO NOTHING;

INSERT INTO refresh_tokens (id, user_id, token_hash, expires_at, revoked_at, user_agent, ip_address, created_at)
VALUES
    (
        'cccccccc-cccc-cccc-cccc-ccccccccccc1',
        '11111111-1111-1111-1111-111111111111',
        'seeded-refresh-token-hash-alice',
        NOW() + INTERVAL '7 days',
        NULL,
        'seed-script',
        '127.0.0.1',
        NOW() - INTERVAL '1 hour'
    ),
    (
        'cccccccc-cccc-cccc-cccc-ccccccccccc2',
        '22222222-2222-2222-2222-222222222222',
        'seeded-refresh-token-hash-bob',
        NOW() + INTERVAL '7 days',
        NULL,
        'seed-script',
        '127.0.0.1',
        NOW() - INTERVAL '1 hour'
    )
ON CONFLICT (id) DO NOTHING;
