# VkValidationFeatureDisableEXT(3) Manual Page

## Name

VkValidationFeatureDisableEXT - Specify validation features to disable



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of elements of the
[VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeaturesEXT.html)::`pDisabledValidationFeatures`
array, specifying validation features to be disabled, are:

``` c
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

## <a href="#_description" class="anchor"></a>Description

- `VK_VALIDATION_FEATURE_DISABLE_ALL_EXT` specifies that all validation
  checks are disabled.

- `VK_VALIDATION_FEATURE_DISABLE_SHADERS_EXT` specifies that shader
  validation is disabled. This feature is enabled by default.

- `VK_VALIDATION_FEATURE_DISABLE_THREAD_SAFETY_EXT` specifies that
  thread safety validation is disabled. This feature is enabled by
  default.

- `VK_VALIDATION_FEATURE_DISABLE_API_PARAMETERS_EXT` specifies that
  stateless parameter validation is disabled. This feature is enabled by
  default.

- `VK_VALIDATION_FEATURE_DISABLE_OBJECT_LIFETIMES_EXT` specifies that
  object lifetime validation is disabled. This feature is enabled by
  default.

- `VK_VALIDATION_FEATURE_DISABLE_CORE_CHECKS_EXT` specifies that core
  validation checks are disabled. This feature is enabled by default. If
  this feature is disabled, the shader validation and GPU-assisted
  validation features are also disabled.

- `VK_VALIDATION_FEATURE_DISABLE_UNIQUE_HANDLES_EXT` specifies that
  protection against duplicate non-dispatchable object handles is
  disabled. This feature is enabled by default.

- `VK_VALIDATION_FEATURE_DISABLE_SHADER_VALIDATION_CACHE_EXT` specifies
  that there will be no caching of shader validation results and every
  shader will be validated on every application execution. Shader
  validation caching is enabled by default.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_validation_features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_validation_features.html),
[VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeaturesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkValidationFeatureDisableEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
