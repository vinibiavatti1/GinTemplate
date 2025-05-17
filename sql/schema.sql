CREATE TABLE `Users` (
    `ID`            INT AUTO_INCREMENT PRIMARY KEY,
    `Name`          VARCHAR(255) NOT NULL,
    `Email`         VARCHAR(255) NOT NULL,
    `PasswordHash`  VARCHAR(255) NOT NULL,
    `Role`          VARCHAR(255) NOT NULL,
    `CreatedAt`     DATETIME NOT NULL DEFAULT NOW(),
    `UpdatedAt`     DATETIME NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    `Active`        BOOLEAN NOT NULL DEFAULT TRUE,
    `LastLogin`     DATETIME NULL,
    `Hash`          VARCHAR(255) NULL,
    `EmailVerified` BOOLEAN NOT NULL DEFAULT FALSE
);
