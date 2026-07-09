CREATE TABLE
    "user" (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        username VARCHAR(255) NOT NULL UNIQUE,
        email VARCHAR(255) NOT NULL UNIQUE,
        password_hash VARCHAR(255) NOT NULL,
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
    );

CREATE TABLE
    "profile_installer" (
        profile_id UUID REFERENCES "profile" (id) ON DELETE CASCADE,
        installer_id UUID REFERENCES "installer" (id) ON DELETE CASCADE,
        execution_order INT NOT NULL,
        PRIMARY KEY (profile_id, installer_id)
    );

CREATE TABLE
    "global_package" (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        name VARCHAR(255) NOT NULL,
        version VARCHAR(50) NOT NULL,
        description TEXT NOT NULL,
        category VARCHAR(50) NOT NULL,
        package_script JSONB NOT NULL,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    "profile_package" (
        profile_id UUID REFERENCES "profile" (id) ON DELETE CASCADE,
        global_package_id UUID REFERENCES "global_package" (id) ON DELETE CASCADE,
        PRIMARY KEY (profile_id, global_package_id)
    );

CREATE TABLE
    "git_repository" (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        profile_id UUID REFERENCES "profile" (id) ON DELETE CASCADE,
        name VARCHAR(255) NOT NULL,
        repo_url VARCHAR(255),
        branch VARCHAR(50),
        target_directory VARCHAR(255),
        visibility VARCHAR(50) NOT NULL,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE
    "custom_script" (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
        profile_id UUID REFERENCES "profile" (id) ON DELETE CASCADE,
        name VARCHAR(255),
        description TEXT,
        shell_type VARCHAR(50),
        run_order VARCHAR(50),
        content TEXT,
        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
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