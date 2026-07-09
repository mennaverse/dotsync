CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE
    "user" (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        name VARCHAR(255),
        username VARCHAR(255) NOT NULL UNIQUE,
        email VARCHAR(255) NOT NULL UNIQUE,
        email_verified BOOLEAN DEFAULT FALSE,
        password_hash VARCHAR(255) NOT NULL,
        banned BOOLEAN DEFAULT FALSE NOT NULL,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    "profile" (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        user_id UUID REFERENCES "user" (id) ON DELETE CASCADE,
        name VARCHAR(255),
        description TEXT,
        visibility VARCHAR(50) NOT NULL,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    "installer" (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        user_id UUID REFERENCES "user" (id) ON DELETE CASCADE,
        name VARCHAR(255) NOT NULL,
        description TEXT,
        category VARCHAR(50) NOT NULL,
        visibility VARCHAR(50) NOT NULL,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    "installer_script" (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        installer_id UUID REFERENCES "installer" (id) ON DELETE CASCADE,
        os_family VARCHAR(50) NOT NULL,
        shell_type VARCHAR(50) NOT NULL,
        pre_install_script TEXT,
        main_install_script TEXT NOT NULL,
        post_install_script TEXT,
        uninstall_script TEXT,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT unique_installer_os_shell UNIQUE (installer_id, os_family, shell_type)
    );

CREATE TABLE
    "profile_installer" (
        profile_id UUID REFERENCES "profile" (id) ON DELETE CASCADE,
        installer_id UUID REFERENCES "installer" (id) ON DELETE CASCADE,
        execution_order INT DEFAULT 0 NOT NULL,
        PRIMARY KEY (profile_id, installer_id)
    );

CREATE TABLE
    "repository" (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        profile_id UUID REFERENCES "profile" (id) ON DELETE CASCADE,
        repo_url TEXT NOT NULL UNIQUE,
        branch VARCHAR(50) DEFAULT 'main' NOT NULL,
        visibility VARCHAR(50) NOT NULL,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    "profile_repository" (
        profile_id UUID REFERENCES "profile" (id) ON DELETE CASCADE,
        repository_id UUID REFERENCES "repository" (id) ON DELETE CASCADE,
        target_directory TEXT NOT NULL,
        branch VARCHAR(50) DEFAULT 'main' NOT NULL,
        PRIMARY KEY (profile_id, repository_id)
    );

CREATE TABLE
    "cli_token" (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        user_id UUID REFERENCES "user" (id) ON DELETE CASCADE,
        token VARCHAR(255) NOT NULL UNIQUE,
        device_name VARCHAR(255),
        last_used_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );