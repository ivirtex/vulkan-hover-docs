# VkPipelineRenderingCreateInfo(3) Manual Page

## Name

VkPipelineRenderingCreateInfo - Structure specifying attachment formats



## [](#_c_specification)C Specification

The `VkPipelineRenderingCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkPipelineRenderingCreateInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           viewMask;
    uint32_t           colorAttachmentCount;
    const VkFormat*    pColorAttachmentFormats;
    VkFormat           depthAttachmentFormat;
    VkFormat           stencilAttachmentFormat;
} VkPipelineRenderingCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_dynamic_rendering
typedef VkPipelineRenderingCreateInfo VkPipelineRenderingCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `viewMask` is a bitfield of view indices describing which views are active during rendering. It **must** match [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`viewMask` when rendering.
- `colorAttachmentCount` is the number of entries in `pColorAttachmentFormats`
- `pColorAttachmentFormats` is a pointer to an array of [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) values defining the format of color attachments used in this pipeline.
- `depthAttachmentFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value defining the format of the depth attachment used in this pipeline.
- `stencilAttachmentFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value defining the format of the stencil attachment used in this pipeline.

## [](#_description)Description

When a pipeline is created without a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html), if the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html) includes this structure, it specifies the view mask and format of attachments used for rendering. If this structure is not specified, and the pipeline does not include a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html), `viewMask` and `colorAttachmentCount` are `0`, and `depthAttachmentFormat` and `stencilAttachmentFormat` are `VK_FORMAT_UNDEFINED`. If a graphics pipeline is created with a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html), parameters of this structure are ignored.

If `depthAttachmentFormat`, `stencilAttachmentFormat`, or any element of `pColorAttachmentFormats` is `VK_FORMAT_UNDEFINED`, it indicates that the corresponding attachment is unused within the render pass. Valid formats indicate that an attachment **can** be used - but it is still valid to set the attachment to `NULL` when beginning rendering.

If the render pass is going to be used with an external format resolve attachment, a [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html) structure **must** also be included in the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), defining the external format of the resolve attachment that will be used.

Valid Usage

- [](#VUID-VkPipelineRenderingCreateInfo-colorAttachmentCount-09533)VUID-VkPipelineRenderingCreateInfo-colorAttachmentCount-09533  
  `colorAttachmentCount` **must** be less than or equal to [`maxColorAttachments`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxColorAttachments)

Valid Usage (Implicit)

- [](#VUID-VkPipelineRenderingCreateInfo-sType-sType)VUID-VkPipelineRenderingCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_RENDERING_CREATE_INFO`

## [](#_see_also)See Also

[VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineRenderingCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0