## PG Back Web - Contribution Guidelines

Thank you for your interest in contributing to the PG Back Web project! Please follow these guidelines to ensure smooth collaboration and consistent code quality.

### Main Branch Policy  
- The **main branch** reflects the latest **stable release** of the project.  
- **No direct commits** should ever be made to the main branch.  
- All development work should happen in feature branches and merged via **Pull Requests (PRs)** into the **develop** branch.  
- The **develop branch** contains the latest code under active development. Once a new release is ready, the main branch will be updated by merging from the development branch.

### Development Workflow  
1. **Fork the repository** and create a feature branch from the `develop` branch.  
   - Use descriptive names for your branches, e.g., `feature/add-new-endpoint` or `bugfix/fix-connection-issue`.  

2. **Make your changes** in your feature branch.

3. **Ensure all tests pass** and the code follows the project’s style guidelines.

4. **Open a Pull Request (PR)** against the `develop` branch.

5. Your PR will be reviewed by maintainers. Please be responsive to feedback.

6. Once approved, the changes will be merged into the `develop` branch. A merge into the `main` branch will only occur when releasing a new version.

### Project Dependencies  
This project requires the following dependencies installed on your system:  
- [**Taskfile**](https://taskfile.dev/) – Used to automate development tasks.  
- **Docker** – For containerized environments.  
- **Docker Compose** – To manage multi-container setups.

### How to Use Taskfile Commands  
- To see all available commands, run:
  ```bash
  task --list
  ```

- The **primary command** for development is:
  ```bash
  task on
  ```
  This command should be run from your **host environment** to start a local development server.

### Additional Notes  
- Always **write clear commit messages** that explain the purpose of your changes.
- **Keep your fork up to date** with the latest changes from the `develop` branch.
- Be respectful and follow the project’s code of conduct when interacting with other contributors.

We appreciate your contributions and are excited to have you on board!
