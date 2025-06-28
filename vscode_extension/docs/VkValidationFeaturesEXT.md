# VkValidationFeaturesEXT(3) Manual Page

## Name

VkValidationFeaturesEXT - Specify validation features to enable or disable for a Vulkan instance



## [](#_c_specification)C Specification

When creating a Vulkan instance for which you wish to enable or disable specific validation features, add a [VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeaturesEXT.html) structure to the `pNext` chain of the [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html) structure, specifying the features to be enabled or disabled.

```c++
// Provided by VK_EXT_validation_features
typedef struct VkValidationFeaturesEXT {
    VkStructureType                         sType;
    const void*                             pNext;
    uint32_t                                enabledValidationFeatureCount;
    const VkValidationFeatureEnableEXT*     pEnabledValidationFeatures;
    uint32_t                                disabledValidationFeatureCount;
    const VkValidationFeatureDisableEXT*    pDisabledValidationFeatures;
} VkValidationFeaturesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `enabledValidationFeatureCount` is the number of features to enable.
- `pEnabledValidationFeatures` is a pointer to an array of [VkValidationFeatureEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeatureEnableEXT.html) values specifying the validation features to be enabled.
- `disabledValidationFeatureCount` is the number of features to disable.
- `pDisabledValidationFeatures` is a pointer to an array of [VkValidationFeatureDisableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeatureDisableEXT.html) values specifying the validation features to be disabled.

## [](#_description)Description

Valid Usage

- [](#VUID-VkValidationFeaturesEXT-pEnabledValidationFeatures-02967)VUID-VkValidationFeaturesEXT-pEnabledValidationFeatures-02967  
  If the `pEnabledValidationFeatures` array contains `VK_VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_RESERVE_BINDING_SLOT_EXT`, then it **must** also contain `VK_VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_EXT` or `VK_VALIDATION_FEATURE_ENABLE_DEBUG_PRINTF_EXT`

Valid Usage (Implicit)

- [](#VUID-VkValidationFeaturesEXT-sType-sType)VUID-VkValidationFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VALIDATION_FEATURES_EXT`
- [](#VUID-VkValidationFeaturesEXT-pEnabledValidationFeatures-parameter)VUID-VkValidationFeaturesEXT-pEnabledValidationFeatures-parameter  
  If `enabledValidationFeatureCount` is not `0`, `pEnabledValidationFeatures` **must** be a valid pointer to an array of `enabledValidationFeatureCount` valid [VkValidationFeatureEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeatureEnableEXT.html) values
- [](#VUID-VkValidationFeaturesEXT-pDisabledValidationFeatures-parameter)VUID-VkValidationFeaturesEXT-pDisabledValidationFeatures-parameter  
  If `disabledValidationFeatureCount` is not `0`, `pDisabledValidationFeatures` **must** be a valid pointer to an array of `disabledValidationFeatureCount` valid [VkValidationFeatureDisableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeatureDisableEXT.html) values

## [](#_see_also)See Also

[VK\_EXT\_validation\_features](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_features.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkValidationFeatureDisableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeatureDisableEXT.html), [VkValidationFeatureEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFeatureEnableEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkValidationFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0