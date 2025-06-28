# VkValidationFeatureEnableEXT(3) Manual Page

## Name

VkValidationFeatureEnableEXT - Specify validation features to enable



## [](#_c_specification)C Specification

Possible values of elements of the [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeaturesEXT.html)::`pEnabledValidationFeatures` array, specifying validation features to be enabled, are:

```c++
// Provided by VK_EXT_validation_features
typedef enum VkValidationFeatureEnableEXT {
    VK_VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_EXT = 0,
    VK_VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_RESERVE_BINDING_SLOT_EXT = 1,
    VK_VALIDATION_FEATURE_ENABLE_BEST_PRACTICES_EXT = 2,
    VK_VALIDATION_FEATURE_ENABLE_DEBUG_PRINTF_EXT = 3,
    VK_VALIDATION_FEATURE_ENABLE_SYNCHRONIZATION_VALIDATION_EXT = 4,
} VkValidationFeatureEnableEXT;
```

## [](#_description)Description

- `VK_VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_EXT` specifies that GPU-assisted validation is enabled. Activating this feature instruments shader programs to generate additional diagnostic data. This feature is disabled by default.
- `VK_VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_RESERVE_BINDING_SLOT_EXT` specifies that the validation layers reserve a descriptor set binding slot for their own use. The layer reports a value for [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`maxBoundDescriptorSets` that is one less than the value reported by the device. If the device supports the binding of only one descriptor set, the validation layer does not perform GPU-assisted validation. This feature is disabled by default.
- `VK_VALIDATION_FEATURE_ENABLE_BEST_PRACTICES_EXT` specifies that Vulkan best-practices validation is enabled. Activating this feature enables the output of warnings related to common misuse of the API, but which are not explicitly prohibited by the specification. This feature is disabled by default.
- `VK_VALIDATION_FEATURE_ENABLE_DEBUG_PRINTF_EXT` specifies that the layers will process `debugPrintfEXT` operations in shaders and send the resulting output to the debug callback. This feature is disabled by default.
- `VK_VALIDATION_FEATURE_ENABLE_SYNCHRONIZATION_VALIDATION_EXT` specifies that Vulkan synchronization validation is enabled. This feature reports resource access conflicts due to missing or incorrect synchronization operations between actions (Draw, Copy, Dispatch, Blit) reading or writing the same regions of memory. This feature is disabled by default.

## [](#_see_also)See Also

[VK\_EXT\_validation\_features](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_features.html), [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeaturesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkValidationFeatureEnableEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0