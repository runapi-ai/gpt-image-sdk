# frozen_string_literal: true

module RunApi
  module GptImage
    CONTRACT = {
      "edit-image" => {
        "models" => ["gpt-image-1.5"],
        "fields_by_model" => {
          "gpt-image-1.5" => {
            "aspect_ratio" => {
              "enum" => ["1:1", "2:3", "3:2"],
              "required" => true
            },
            "prompt" => {
              "required" => true
            },
            "quality" => {
              "enum" => ["medium", "high"],
              "required" => true
            },
            "source_image_urls" => {
              "required" => true
            }
          }
        }
      },
      "text-to-image" => {
        "models" => ["gpt-image-1.5"],
        "fields_by_model" => {
          "gpt-image-1.5" => {
            "aspect_ratio" => {
              "enum" => ["1:1", "2:3", "3:2"],
              "required" => true
            },
            "prompt" => {
              "required" => true
            },
            "quality" => {
              "enum" => ["medium", "high"],
              "required" => true
            }
          }
        }
      }
    }.freeze
  end
end
