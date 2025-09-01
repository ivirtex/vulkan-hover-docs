# VkValidationFeatureDisableEXT(3) Manual Page

## Name

VkValidationFeatureDisableEXT - Specify validation features to disable



## [](#_c_specification)C Specification

Possible values of elements of the [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeaturesEXT.html)::`pDisabledValidationFeatures` array, specifying validation features to be disabled, are:

```c++
// Provided by VK_EXT_validation_features
typedef enum VkValidationFeatureDisableEXT {
    VK_VALIDATION_FEATURE_DISABLE_ALL_EXT = 0,
    VK_VALIDATION_FEATURE_DISABLE_SHADERS_EXT = 1,
    VK_VALIDATION_FEATURE_DISABLE_THREAD_SAFETY_EXT = 2,
    VK_VALIDATION_FEATURE_DISABLE_API_PARAMETERS_EXT = 3,
    VK_VALIDATION_FEATURE_DISABLE_OBJECT_LIFETIMES_EXT = 4,
    VK_VALIDATION_FEATURE_DISABLE_CORE_CHECKS_EXT = 5,
    VK_VALIDATION_FEATURE_DISABLE_UNIQUE_HANDLES_EXT = 6,
    VK_VALIDATION_FEATURE_DISABLE_SHADER_VALIDATION_CACHE_EXT = 7,
} VkValidationFeatureDisableEXT;
```

## [](#_description)Description

- `VK_VALIDATION_FEATURE_DISABLE_ALL_EXT` specifies that all validation checks are disabled.
- `VK_VALIDATION_FEATURE_DISABLE_SHADERS_EXT` specifies that shader validation, both runtime and standalone, is disabled. This validation occurs inside [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateInfoEXT.html) and [VkShaderModuleCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModuleCreateInfo.html). This feature is enabled by default.
- `VK_VALIDATION_FEATURE_DISABLE_THREAD_SAFETY_EXT` specifies that thread safety validation is disabled. This feature is enabled by default.
- `VK_VALIDATION_FEATURE_DISABLE_API_PARAMETERS_EXT` specifies that stateless parameter validation is disabled. This feature is enabled by default.
- `VK_VALIDATION_FEATURE_DISABLE_OBJECT_LIFETIMES_EXT` specifies that object lifetime validation is disabled. This feature is enabled by default.
- `VK_VALIDATION_FEATURE_DISABLE_CORE_CHECKS_EXT` specifies that core validation checks are disabled. This feature is enabled by default. If this feature is disabled, `VK_VALIDATION_FEATURE_DISABLE_SHADERS_EXT` is implied.
- `VK_VALIDATION_FEATURE_DISABLE_UNIQUE_HANDLES_EXT` specifies that protection against duplicate non-dispatchable object handles is disabled. This feature is enabled by default.
- `VK_VALIDATION_FEATURE_DISABLE_SHADER_VALIDATION_CACHE_EXT` specifies that there will be no caching of shader validation results and every shader will be validated on every application execution. Shader validation caching is enabled by default.

## [](#_see_also)See Also

[VK\_EXT\_validation\_features](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_features.html), [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeaturesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkValidationFeatureDisableEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0