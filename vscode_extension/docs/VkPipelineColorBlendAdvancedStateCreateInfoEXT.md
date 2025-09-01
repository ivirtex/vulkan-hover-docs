# VkPipelineColorBlendAdvancedStateCreateInfoEXT(3) Manual Page

## Name

VkPipelineColorBlendAdvancedStateCreateInfoEXT - Structure specifying parameters that affect advanced blend operations



## [](#_c_specification)C Specification

If the `pNext` chain of [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendStateCreateInfo.html) includes a `VkPipelineColorBlendAdvancedStateCreateInfoEXT` structure, then that structure includes parameters that affect advanced blend operations.

The `VkPipelineColorBlendAdvancedStateCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_blend_operation_advanced
typedef struct VkPipelineColorBlendAdvancedStateCreateInfoEXT {
    VkStructureType      sType;
    const void*          pNext;
    VkBool32             srcPremultiplied;
    VkBool32             dstPremultiplied;
    VkBlendOverlapEXT    blendOverlap;
} VkPipelineColorBlendAdvancedStateCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `srcPremultiplied` specifies whether the source color of the blend operation is treated as premultiplied.
- `dstPremultiplied` specifies whether the destination color of the blend operation is treated as premultiplied.
- `blendOverlap` is a [VkBlendOverlapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOverlapEXT.html) value specifying how the source and destination sample’s coverage is correlated.

## [](#_description)Description

If this structure is not present, `srcPremultiplied` and `dstPremultiplied` are both considered to be `VK_TRUE`, and `blendOverlap` is considered to be `VK_BLEND_OVERLAP_UNCORRELATED_EXT`.

Valid Usage

- [](#VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-srcPremultiplied-01424)VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-srcPremultiplied-01424  
  If the [non-premultiplied source color](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-advancedBlendNonPremultipliedSrcColor) property is not supported, `srcPremultiplied` **must** be `VK_TRUE`
- [](#VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-dstPremultiplied-01425)VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-dstPremultiplied-01425  
  If the [non-premultiplied destination color](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-advancedBlendNonPremultipliedDstColor) property is not supported, `dstPremultiplied` **must** be `VK_TRUE`
- [](#VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-blendOverlap-01426)VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-blendOverlap-01426  
  If the [correlated overlap](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-advancedBlendCorrelatedOverlap) property is not supported, `blendOverlap` **must** be `VK_BLEND_OVERLAP_UNCORRELATED_EXT`

Valid Usage (Implicit)

- [](#VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-sType-sType)VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_ADVANCED_STATE_CREATE_INFO_EXT`
- [](#VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-blendOverlap-parameter)VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-blendOverlap-parameter  
  `blendOverlap` **must** be a valid [VkBlendOverlapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOverlapEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_blend\_operation\_advanced](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_blend_operation_advanced.html), [VkBlendOverlapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlendOverlapEXT.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineColorBlendAdvancedStateCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0