# VkPipelineViewportSwizzleStateCreateInfoNV(3) Manual Page

## Name

VkPipelineViewportSwizzleStateCreateInfoNV - Structure specifying
swizzle applied to primitive clip coordinates



## <a href="#_c_specification" class="anchor"></a>C Specification

Each primitive sent to a given viewport has a swizzle and **optional**
negation applied to its clip coordinates. The swizzle that is applied
depends on the viewport index, and is controlled by the
`VkPipelineViewportSwizzleStateCreateInfoNV` pipeline state:

``` c
// Provided by VK_NV_viewport_swizzle
typedef struct VkPipelineViewportSwizzleStateCreateInfoNV {
    VkStructureType                                sType;
    const void*                                    pNext;
    VkPipelineViewportSwizzleStateCreateFlagsNV    flags;
    uint32_t                                       viewportCount;
    const VkViewportSwizzleNV*                     pViewportSwizzles;
} VkPipelineViewportSwizzleStateCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `viewportCount` is the number of viewport swizzles used by the
  pipeline.

- `pViewportSwizzles` is a pointer to an array of
  [VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportSwizzleNV.html) structures, defining
  the viewport swizzles.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkPipelineViewportSwizzleStateCreateInfoNV-viewportCount-01215"
  id="VUID-VkPipelineViewportSwizzleStateCreateInfoNV-viewportCount-01215"></a>
  VUID-VkPipelineViewportSwizzleStateCreateInfoNV-viewportCount-01215  
  `viewportCount` **must** be greater than or equal to the
  `viewportCount` set in `VkPipelineViewportStateCreateInfo`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineViewportSwizzleStateCreateInfoNV-sType-sType"
  id="VUID-VkPipelineViewportSwizzleStateCreateInfoNV-sType-sType"></a>
  VUID-VkPipelineViewportSwizzleStateCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_SWIZZLE_STATE_CREATE_INFO_NV`

- <a
  href="#VUID-VkPipelineViewportSwizzleStateCreateInfoNV-flags-zerobitmask"
  id="VUID-VkPipelineViewportSwizzleStateCreateInfoNV-flags-zerobitmask"></a>
  VUID-VkPipelineViewportSwizzleStateCreateInfoNV-flags-zerobitmask  
  `flags` **must** be `0`

- <a
  href="#VUID-VkPipelineViewportSwizzleStateCreateInfoNV-pViewportSwizzles-parameter"
  id="VUID-VkPipelineViewportSwizzleStateCreateInfoNV-pViewportSwizzles-parameter"></a>
  VUID-VkPipelineViewportSwizzleStateCreateInfoNV-pViewportSwizzles-parameter  
  `pViewportSwizzles` **must** be a valid pointer to an array of
  `viewportCount` valid [VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportSwizzleNV.html)
  structures

- <a
  href="#VUID-VkPipelineViewportSwizzleStateCreateInfoNV-viewportCount-arraylength"
  id="VUID-VkPipelineViewportSwizzleStateCreateInfoNV-viewportCount-arraylength"></a>
  VUID-VkPipelineViewportSwizzleStateCreateInfoNV-viewportCount-arraylength  
  `viewportCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_viewport_swizzle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_viewport_swizzle.html),
[VkPipelineViewportSwizzleStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportSwizzleStateCreateFlagsNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkViewportSwizzleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportSwizzleNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineViewportSwizzleStateCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
