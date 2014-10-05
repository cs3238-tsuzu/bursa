default_attributes(
  "java" => {
    "install_flavor" => "oracle",
    "jdk_version" => "7",
    "oracle" => {
      "accept_oracle_download_terms" => true
    },
    "accept_license_agreement" => true
  },
  "elasticsearch" => {
    "version" => "1.4.0.Beta1",
    "cluster" => {
      "name" => "elasticpurse"
    },
    "plugins" => {
      "karmi/elasticsearch-paramedic" => {
        "url" => "https://github.com/karmi/elasticsearch-paramedic/archive/master.zip"
      },
      "marvel" => {
        "url" => "http://download.elasticsearch.org/elasticsearch/marvel/marvel-latest.zip"
      }
    }
  },
  "postgresql" => {
    "password" => {
      "postgres" => "sa_securemebaby"
    }
  }
)
