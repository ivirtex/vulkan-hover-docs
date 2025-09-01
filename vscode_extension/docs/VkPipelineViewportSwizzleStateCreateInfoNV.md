# VkPipelineViewportSwizzleStateCreateInfoNV(3) Manual Page

## Name

VkPipelineViewportSwizzleStateCreateInfoNV - Structure specifying swizzle applied to primitive clip coordinates



## [](#_c_specification)C Specification

Each primitive sent to a given viewport has a swizzle and **optional** negation applied to its clip coordinates. The swizzle that is applied depends on the viewport index, and is controlled by the `VkPipelineViewportSwizzleStateCreateInfoNV` pipeline state:

```c++
// Provided by VK_NV_viewport_swizzle
typedef struct VkPipelineViewportSwizzleStateCreateInfoNV {
    VkStructureType                                sType;
    const void*                                    pNext;
    VkPipelineViewportSwizzleStateCreateFlagsNV    flags;
    uint32_t                                       viewportCount;
    const VkViewportSwizzleNV*                     pViewportSwizzles;
} VkPipelineViewportSwizzleStateCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `viewportCount` is the number of viewport swizzles used by the pipeline.
- `pViewportSwizzles` is a pointer to an array of [VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportSwizzleNV.html) structures, defining the viewport swizzles.

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineViewportSwizzleStateCreateInfoNV-viewportCount-01215)VUID-VkPipelineViewportSwizzleStateCreateInfoNV-viewportCount-01215  
  `viewportCount` **must** be greater than or equal to the `viewportCount` set in `VkPipelineViewportStateCreateInfo`

Valid Usage (Implicit)

- [](#VUID-VkPipelineViewportSwizzleStateCreateInfoNV-sType-sType)VUID-VkPipelineViewportSwizzleStateCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_SWIZZLE_STATE_CREATE_INFO_NV`
- [](#VUID-VkPipelineViewportSwizzleStateCreateInfoNV-flags-zerobitmask)VUID-VkPipelineViewportSwizzleStateCreateInfoNV-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkPipelineViewportSwizzleStateCreateInfoNV-pViewportSwizzles-parameter)VUID-VkPipelineViewportSwizzleStateCreateInfoNV-pViewportSwizzles-parameter  
  `pViewportSwizzles` **must** be a valid pointer to an array of `viewportCount` valid [VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportSwizzleNV.html) structures
- [](#VUID-VkPipelineViewportSwizzleStateCreateInfoNV-viewportCount-arraylength)VUID-VkPipelineViewportSwizzleStateCreateInfoNV-viewportCount-arraylength  
  `viewportCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_NV\_viewport\_swizzle](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_viewport_swizzle.html), [VkPipelineViewportSwizzleStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportSwizzleStateCreateFlagsNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewportSwizzleNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineViewportSwizzleStateCreateInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0