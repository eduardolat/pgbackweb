# Define command aliases
alias t='task'
alias td='task dev'
alias tb='task build'
alias tt='task test'
alias tl='task lint'
alias tls='task --list'
alias tf='task format'
alias ll='ls -alF'
alias la='ls -A'
alias l='ls -CF'
alias ..='cd ..'
alias c='clear'
echo "[OK] aliases set"

# Set the user file-creation mode mask to 000, which allows all
# users read, write, and execute permissions for newly created files.
umask 000
echo "[OK] umask set"

# Run the 'fixperms' task that fixes the permissions of the files and
# directories in the project.
task fixperms
echo "[OK] permissions fixed"

# Configure Git to ignore ownership and file mode changes.
git config --global --add safe.directory '*'
git config --global core.fileMode false
git config --unset core.fileMode
git config core.fileMode false
echo "[OK] git configured"

echo "
───────────────────────────────────────────────
── Website: https://eduardo.lat ───────────────
── Github: https://github.com/eduardolat ──────
───────────────────────────────────────────────
── Development environment is ready to use! ───
───────────────────────────────────────────────
"