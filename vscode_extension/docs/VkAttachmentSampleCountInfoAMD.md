# VkAttachmentSampleCountInfoAMD(3) Manual Page

## Name

VkAttachmentSampleCountInfoAMD - Structure specifying command buffer
inheritance info for dynamic render pass instances



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAttachmentSampleCountInfoAMD` or `VkAttachmentSampleCountInfoNV`
structure is defined as:

``` c
// Provided by VK_KHR_dynamic_rendering with VK_AMD_mixed_attachment_samples
typedef struct VkAttachmentSampleCountInfoAMD {
    VkStructureType                 sType;
    const void*                     pNext;
    uint32_t                        colorAttachmentCount;
    const VkSampleCountFlagBits*    pColorAttachmentSamples;
    VkSampleCountFlagBits           depthStencilAttachmentSamples;
} VkAttachmentSampleCountInfoAMD;
```

or the equivalent

``` c
// Provided by VK_KHR_dynamic_rendering with VK_NV_framebuffer_mixed_samples
typedef VkAttachmentSampleCountInfoAMD VkAttachmentSampleCountInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this structure

- `colorAttachmentCount` is the number of color attachments specified in
  a render pass instance.

- `pColorAttachmentSamples` is a pointer to an array of
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) values defining
  the sample count of color attachments.

- `depthStencilAttachmentSamples` is a
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value defining the
  sample count of a depth/stencil attachment.

## <a href="#_description" class="anchor"></a>Description

If
[VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)::`renderPass`
is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
`VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT` is specified in
[VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`flags`, and
the `pNext` chain of
[VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
includes `VkAttachmentSampleCountInfoAMD`, then this structure defines
the sample counts of each attachment within the render pass instance. If
`VkAttachmentSampleCountInfoAMD` is not included, the value of
[VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)::`rasterizationSamples`
is used as the sample count for each attachment. If
[VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)::`renderPass`
is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), or
`VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT` is not specified in
[VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`flags`,
parameters of this structure are ignored.

`VkAttachmentSampleCountInfoAMD` **can** also be included in the `pNext`
chain of
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html). When
a graphics pipeline is created without a
[VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html), if this structure is included in the
`pNext` chain of
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html), it
specifies the sample count of attachments used for rendering. If this
structure is not specified, and the pipeline does not include a
[VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html), the value of
[VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html)::`rasterizationSamples`
is used as the sample count for each attachment. If a graphics pipeline
is created with a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html), parameters of
this structure are ignored.

Valid Usage (Implicit)

- <a href="#VUID-VkAttachmentSampleCountInfoAMD-sType-sType"
  id="VUID-VkAttachmentSampleCountInfoAMD-sType-sType"></a>
  VUID-VkAttachmentSampleCountInfoAMD-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ATTACHMENT_SAMPLE_COUNT_INFO_AMD`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_mixed_attachment_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_mixed_attachment_samples.html),
[VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html),
[VK_NV_framebuffer_mixed_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_framebuffer_mixed_samples.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAttachmentSampleCountInfoAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
