!/bin/bash

if [ -z "$1" ]; then
  echo "Please provide a name as a command-line parameter."
  exit 1
fi

name="$1"

# Create the directory structure
mkdir -p ./app/controller
mkdir -p ./app/entity
mkdir -p ./app/repo
mkdir -p ./app/service

# Create the controller file
cat <<EOF > ./app/controller/${name}Controller.go
package ${name}Controller

// Add your controller code here
EOF

# Create the entity file
cat <<EOF > ./app/entity/${name}Entity.go
package ${name}Entity

import "github.com/jinzhu/gorm"

// ${name}Entity structure for our blog
type ${name}Entity struct {
	gorm.Model
	Variable1    string \`gorm:"column:Variable1" json:"Variable1"\`
	Variable2 string \`gorm:"column:variable2" json:"Variable2"\`
}

func (${name}Entity) TableName() string {
	return "${name}"
}

EOF

# Create the repository file
cat <<EOF > ./app/repo/${name}Repo.go
package ${name}Repo

// Add your repository code here
EOF

# Create the service file
cat <<EOF > ./app/service/${name}Service.go
package ${name}Service

// Add your service code here
EOF

echo "Files created successfully!"
