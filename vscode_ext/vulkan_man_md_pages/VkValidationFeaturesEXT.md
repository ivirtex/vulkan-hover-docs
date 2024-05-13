# VkValidationFeaturesEXT(3) Manual Page

## Name

VkValidationFeaturesEXT - Specify validation features to enable or
disable for a Vulkan instance



## <a href="#_c_specification" class="anchor"></a>C Specification

When creating a Vulkan instance for which you wish to enable or disable
specific validation features, add a
[VkValidationFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeaturesEXT.html) structure to the
`pNext` chain of the [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html)
structure, specifying the features to be enabled or disabled.

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `enabledValidationFeatureCount` is the number of features to enable.

- `pEnabledValidationFeatures` is a pointer to an array of
  [VkValidationFeatureEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeatureEnableEXT.html)
  values specifying the validation features to be enabled.

- `disabledValidationFeatureCount` is the number of features to disable.

- `pDisabledValidationFeatures` is a pointer to an array of
  [VkValidationFeatureDisableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeatureDisableEXT.html)
  values specifying the validation features to be disabled.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkValidationFeaturesEXT-pEnabledValidationFeatures-02967"
  id="VUID-VkValidationFeaturesEXT-pEnabledValidationFeatures-02967"></a>
  VUID-VkValidationFeaturesEXT-pEnabledValidationFeatures-02967  
  If the `pEnabledValidationFeatures` array contains
  `VK_VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_RESERVE_BINDING_SLOT_EXT`,
  then it **must** also contain
  `VK_VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_EXT`

- <a href="#VUID-VkValidationFeaturesEXT-pEnabledValidationFeatures-02968"
  id="VUID-VkValidationFeaturesEXT-pEnabledValidationFeatures-02968"></a>
  VUID-VkValidationFeaturesEXT-pEnabledValidationFeatures-02968  
  If the `pEnabledValidationFeatures` array contains
  `VK_VALIDATION_FEATURE_ENABLE_DEBUG_PRINTF_EXT`, then it **must** not
  contain `VK_VALIDATION_FEATURE_ENABLE_GPU_ASSISTED_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkValidationFeaturesEXT-sType-sType"
  id="VUID-VkValidationFeaturesEXT-sType-sType"></a>
  VUID-VkValidationFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VALIDATION_FEATURES_EXT`

- <a
  href="#VUID-VkValidationFeaturesEXT-pEnabledValidationFeatures-parameter"
  id="VUID-VkValidationFeaturesEXT-pEnabledValidationFeatures-parameter"></a>
  VUID-VkValidationFeaturesEXT-pEnabledValidationFeatures-parameter  
  If `enabledValidationFeatureCount` is not `0`,
  `pEnabledValidationFeatures` **must** be a valid pointer to an array
  of `enabledValidationFeatureCount` valid
  [VkValidationFeatureEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeatureEnableEXT.html)
  values

- <a
  href="#VUID-VkValidationFeaturesEXT-pDisabledValidationFeatures-parameter"
  id="VUID-VkValidationFeaturesEXT-pDisabledValidationFeatures-parameter"></a>
  VUID-VkValidationFeaturesEXT-pDisabledValidationFeatures-parameter  
  If `disabledValidationFeatureCount` is not `0`,
  `pDisabledValidationFeatures` **must** be a valid pointer to an array
  of `disabledValidationFeatureCount` valid
  [VkValidationFeatureDisableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeatureDisableEXT.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_validation_features](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_validation_features.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkValidationFeatureDisableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeatureDisableEXT.html),
[VkValidationFeatureEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFeatureEnableEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkValidationFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
