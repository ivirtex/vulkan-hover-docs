# VkPipelineColorBlendAdvancedStateCreateInfoEXT(3) Manual Page

## Name

VkPipelineColorBlendAdvancedStateCreateInfoEXT - Structure specifying
parameters that affect advanced blend operations



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html)
includes a `VkPipelineColorBlendAdvancedStateCreateInfoEXT` structure,
then that structure includes parameters that affect advanced blend
operations.

The `VkPipelineColorBlendAdvancedStateCreateInfoEXT` structure is
defined as:

``` c
// Provided by VK_EXT_blend_operation_advanced
typedef struct VkPipelineColorBlendAdvancedStateCreateInfoEXT {
    VkStructureType      sType;
    const void*          pNext;
    VkBool32             srcPremultiplied;
    VkBool32             dstPremultiplied;
    VkBlendOverlapEXT    blendOverlap;
} VkPipelineColorBlendAdvancedStateCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `srcPremultiplied` specifies whether the source color of the blend
  operation is treated as premultiplied.

- `dstPremultiplied` specifies whether the destination color of the
  blend operation is treated as premultiplied.

- `blendOverlap` is a [VkBlendOverlapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOverlapEXT.html) value
  specifying how the source and destination sample’s coverage is
  correlated.

## <a href="#_description" class="anchor"></a>Description

If this structure is not present, `srcPremultiplied` and
`dstPremultiplied` are both considered to be `VK_TRUE`, and
`blendOverlap` is considered to be `VK_BLEND_OVERLAP_UNCORRELATED_EXT`.

Valid Usage

- <a
  href="#VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-srcPremultiplied-01424"
  id="VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-srcPremultiplied-01424"></a>
  VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-srcPremultiplied-01424  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-advancedBlendNonPremultipliedSrcColor"
  target="_blank" rel="noopener">non-premultiplied source color</a>
  property is not supported, `srcPremultiplied` **must** be `VK_TRUE`

- <a
  href="#VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-dstPremultiplied-01425"
  id="VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-dstPremultiplied-01425"></a>
  VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-dstPremultiplied-01425  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-advancedBlendNonPremultipliedDstColor"
  target="_blank" rel="noopener">non-premultiplied destination color</a>
  property is not supported, `dstPremultiplied` **must** be `VK_TRUE`

- <a
  href="#VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-blendOverlap-01426"
  id="VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-blendOverlap-01426"></a>
  VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-blendOverlap-01426  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-advancedBlendCorrelatedOverlap"
  target="_blank" rel="noopener">correlated overlap</a> property is not
  supported, `blendOverlap` **must** be
  `VK_BLEND_OVERLAP_UNCORRELATED_EXT`

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-sType-sType"
  id="VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-sType-sType"></a>
  VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_COLOR_BLEND_ADVANCED_STATE_CREATE_INFO_EXT`

- <a
  href="#VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-blendOverlap-parameter"
  id="VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-blendOverlap-parameter"></a>
  VUID-VkPipelineColorBlendAdvancedStateCreateInfoEXT-blendOverlap-parameter  
  `blendOverlap` **must** be a valid
  [VkBlendOverlapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOverlapEXT.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_blend_operation_advanced](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_blend_operation_advanced.html),
[VkBlendOverlapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlendOverlapEXT.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineColorBlendAdvancedStateCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
