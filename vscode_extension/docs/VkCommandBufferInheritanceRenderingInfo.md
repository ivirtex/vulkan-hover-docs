# VkCommandBufferInheritanceRenderingInfo(3) Manual Page

## Name

VkCommandBufferInheritanceRenderingInfo - Structure specifying command buffer inheritance info for dynamic render pass instances



## [](#_c_specification)C Specification

The `VkCommandBufferInheritanceRenderingInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkCommandBufferInheritanceRenderingInfo {
    VkStructureType          sType;
    const void*              pNext;
    VkRenderingFlags         flags;
    uint32_t                 viewMask;
    uint32_t                 colorAttachmentCount;
    const VkFormat*          pColorAttachmentFormats;
    VkFormat                 depthAttachmentFormat;
    VkFormat                 stencilAttachmentFormat;
    VkSampleCountFlagBits    rasterizationSamples;
} VkCommandBufferInheritanceRenderingInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_dynamic_rendering
typedef VkCommandBufferInheritanceRenderingInfo VkCommandBufferInheritanceRenderingInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure
- `flags` is a bitmask of [VkRenderingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFlagBits.html) used by the render pass instance.
- `viewMask` is the view mask used for rendering.
- `colorAttachmentCount` is the number of color attachments specified in the render pass instance.
- `pColorAttachmentFormats` is a pointer to an array of [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) values defining the format of color attachments.
- `depthAttachmentFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value defining the format of the depth attachment.
- `stencilAttachmentFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value defining the format of the stencil attachment.
- `rasterizationSamples` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) specifying the number of samples used in rasterization.

## [](#_description)Description

If the `pNext` chain of [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html) includes a `VkCommandBufferInheritanceRenderingInfo` structure, then that structure controls parameters of dynamic render pass instances that the [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) **can** be executed within. If [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html)::`renderPass` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), or `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT` is not specified in [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferBeginInfo.html)::`flags`, parameters of this structure are ignored.

If `colorAttachmentCount` is `0` and the [`variableMultisampleRate`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-variableMultisampleRate) feature is enabled, `rasterizationSamples` is ignored.

If `depthAttachmentFormat`, `stencilAttachmentFormat`, or any element of `pColorAttachmentFormats` is `VK_FORMAT_UNDEFINED`, it indicates that the corresponding attachment is unused within the render pass and writes to those attachments are discarded.

Valid Usage

- [](#VUID-VkCommandBufferInheritanceRenderingInfo-colorAttachmentCount-06004)VUID-VkCommandBufferInheritanceRenderingInfo-colorAttachmentCount-06004  
  If `colorAttachmentCount` is not `0`, `rasterizationSamples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-variableMultisampleRate-06005)VUID-VkCommandBufferInheritanceRenderingInfo-variableMultisampleRate-06005  
  If the [`variableMultisampleRate`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-variableMultisampleRate) feature is not enabled, `rasterizationSamples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06540)VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06540  
  If `depthAttachmentFormat` is not `VK_FORMAT_UNDEFINED`, it **must** be a format that includes a depth component
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06007)VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06007  
  If `depthAttachmentFormat` is not `VK_FORMAT_UNDEFINED`, it **must** be a format with [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) that include `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-pColorAttachmentFormats-06492)VUID-VkCommandBufferInheritanceRenderingInfo-pColorAttachmentFormats-06492  
  If any element of `pColorAttachmentFormats` is not `VK_FORMAT_UNDEFINED`, it **must** be a format with [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) that include `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` , or `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV` if the [`linearColorAttachment`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-linearColorAttachment) feature is enabled
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-06541)VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-06541  
  If `stencilAttachmentFormat` is not `VK_FORMAT_UNDEFINED`, it **must** be a format that includes a stencil aspect
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-06199)VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-06199  
  If `stencilAttachmentFormat` is not `VK_FORMAT_UNDEFINED`, it **must** be a format with [potential format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#potential-format-features) that include `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06200)VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06200  
  If `depthAttachmentFormat` is not `VK_FORMAT_UNDEFINED` and `stencilAttachmentFormat` is not `VK_FORMAT_UNDEFINED`, `depthAttachmentFormat` **must** equal `stencilAttachmentFormat`
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-multiview-06008)VUID-VkCommandBufferInheritanceRenderingInfo-multiview-06008  
  If the [`multiview`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-multiview) feature is not enabled, `viewMask` **must** be `0`
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-viewMask-06009)VUID-VkCommandBufferInheritanceRenderingInfo-viewMask-06009  
  The index of the most significant bit in `viewMask` **must** be less than [`maxMultiviewViewCount`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxMultiviewViewCount)

Valid Usage (Implicit)

- [](#VUID-VkCommandBufferInheritanceRenderingInfo-sType-sType)VUID-VkCommandBufferInheritanceRenderingInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_RENDERING_INFO`
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-flags-parameter)VUID-VkCommandBufferInheritanceRenderingInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkRenderingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFlagBits.html) values
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-pColorAttachmentFormats-parameter)VUID-VkCommandBufferInheritanceRenderingInfo-pColorAttachmentFormats-parameter  
  If `colorAttachmentCount` is not `0`, `pColorAttachmentFormats` **must** be a valid pointer to an array of `colorAttachmentCount` valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) values
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-parameter)VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-parameter  
  `depthAttachmentFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-parameter)VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-parameter  
  `stencilAttachmentFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkCommandBufferInheritanceRenderingInfo-rasterizationSamples-parameter)VUID-VkCommandBufferInheritanceRenderingInfo-rasterizationSamples-parameter  
  If `rasterizationSamples` is not `0`, `rasterizationSamples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value

## [](#_see_also)See Also

[VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkRenderingFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingFlags.html), [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCommandBufferInheritanceRenderingInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0