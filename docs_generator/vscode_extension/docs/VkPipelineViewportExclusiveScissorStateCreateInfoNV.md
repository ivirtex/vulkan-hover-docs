# VkPipelineViewportExclusiveScissorStateCreateInfoNV(3) Manual Page

## Name

VkPipelineViewportExclusiveScissorStateCreateInfoNV - Structure specifying parameters controlling exclusive scissor testing



## [](#_c_specification)C Specification

The `VkPipelineViewportExclusiveScissorStateCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_scissor_exclusive
typedef struct VkPipelineViewportExclusiveScissorStateCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           exclusiveScissorCount;
    const VkRect2D*    pExclusiveScissors;
} VkPipelineViewportExclusiveScissorStateCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `exclusiveScissorCount` is the number of exclusive scissor rectangles.
- `pExclusiveScissors` is a pointer to an array of [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html) structures defining exclusive scissor rectangles.

## [](#_description)Description

If the `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV` dynamic state is enabled for a pipeline, the `pExclusiveScissors` member is ignored.

When this structure is included in the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), it defines parameters of the exclusive scissor test. If this structure is not included in the `pNext` chain, it is equivalent to specifying this structure with an `exclusiveScissorCount` of `0`.

Valid Usage

- [](#VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02027)VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02027  
  If the [`multiViewport`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiViewport) feature is not enabled, `exclusiveScissorCount` **must** be `0` or `1`
- [](#VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02028)VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02028  
  `exclusiveScissorCount` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxViewports`
- [](#VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02029)VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-exclusiveScissorCount-02029  
  `exclusiveScissorCount` **must** be `0` or greater than or equal to the `viewportCount` member of [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html)

Valid Usage (Implicit)

- [](#VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-sType-sType)VUID-VkPipelineViewportExclusiveScissorStateCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_EXCLUSIVE_SCISSOR_STATE_CREATE_INFO_NV`

## [](#_see_also)See Also

[VK\_NV\_scissor\_exclusive](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_scissor_exclusive.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRect2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineViewportExclusiveScissorStateCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0