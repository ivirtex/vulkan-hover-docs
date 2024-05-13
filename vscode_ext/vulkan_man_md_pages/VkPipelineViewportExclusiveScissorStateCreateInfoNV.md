# VkPipelineViewportExclusiveScissorStateCreateInfoNV(3) Manual Page

## Name

VkPipelineViewportExclusiveScissorStateCreateInfoNV - Structure
specifying parameters controlling exclusive scissor testing



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineViewportExclusiveScissorStateCreateInfoNV` structure is
defined as:

``` c
// Provided by VK_NV_scissor_exclusive
typedef struct VkPipelineViewportExclusiveScissorStateCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           exclusiveScissorCount;
    const VkRect2D*    pExclusiveScissors;
} VkPipelineViewportExclusiveScissorStateCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `exclusiveScissorCount` is the number of exclusive scissor rectangles.

- `pExclusiveScissors` is a pointer to an array of
  [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures defining exclusive scissor
  rectangles.

## <a href="#_description" class="anchor"></a>Description

If the `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV` dynamic state is enabled
for a pipeline, the `pExclusiveScissors` member is ignored.

When this structure is included in the `pNext` chain of
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html), it
defines parameters of the exclusive scissor test. If this structure is
not included in the `pNext` chain, it is equivalent to specifying this
structure with an `exclusiveScissorCount` of `0`.

Valid Usage

- <a
  href="#VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02027"
  id="VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02027"></a>
  VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02027  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `exclusiveScissorCount` **must** be `0` or `1`

- <a
  href="#VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02028"
  id="VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02028"></a>
  VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02028  
  `exclusiveScissorCount` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxViewports`

- <a
  href="#VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02029"
  id="VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02029"></a>
  VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02029  
  `exclusiveScissorCount` **must** be `0` or greater than or equal to
  the `viewportCount` member of
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-sType-sType"
  id="VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-sType-sType"></a>
  VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_EXCLUSIVE_SCISSOR_STATE_CREATE_INFO_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_scissor_exclusive](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_scissor_exclusive.html),
[VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineViewportExclusiveScissorStateCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
