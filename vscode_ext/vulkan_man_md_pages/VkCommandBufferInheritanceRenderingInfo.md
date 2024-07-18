# VkCommandBufferInheritanceRenderingInfo(3) Manual Page

## Name

VkCommandBufferInheritanceRenderingInfo - Structure specifying command
buffer inheritance info for dynamic render pass instances



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCommandBufferInheritanceRenderingInfo` structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_dynamic_rendering
typedef VkCommandBufferInheritanceRenderingInfo VkCommandBufferInheritanceRenderingInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this structure

- `flags` is a bitmask of
  [VkRenderingFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFlagBits.html) used by the render
  pass instance.

- `viewMask` is the view mask used for rendering.

- `colorAttachmentCount` is the number of color attachments specified in
  the render pass instance.

- `pColorAttachmentFormats` is a pointer to an array of
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) values defining the format of color
  attachments.

- `depthAttachmentFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value defining
  the format of the depth attachment.

- `stencilAttachmentFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value
  defining the format of the stencil attachment.

- `rasterizationSamples` is a
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) specifying the
  number of samples used in rasterization.

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain of
[VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
includes a `VkCommandBufferInheritanceRenderingInfo` structure, then
that structure controls parameters of dynamic render pass instances that
the [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) **can** be executed within.
If
[VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)::`renderPass`
is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), or
`VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT` is not specified in
[VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html)::`flags`,
parameters of this structure are ignored.

If `colorAttachmentCount` is `0` and the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-variableMultisampleRate"
target="_blank" rel="noopener"><code>variableMultisampleRate</code></a>
feature is enabled, `rasterizationSamples` is ignored.

If `depthAttachmentFormat`, `stencilAttachmentFormat`, or any element of
`pColorAttachmentFormats` is `VK_FORMAT_UNDEFINED`, it indicates that
the corresponding attachment is unused within the render pass and writes
to those attachments are discarded.

Valid Usage

- <a
  href="#VUID-VkCommandBufferInheritanceRenderingInfo-colorAttachmentCount-06004"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-colorAttachmentCount-06004"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-colorAttachmentCount-06004  
  If `colorAttachmentCount` is not `0`, `rasterizationSamples` **must**
  be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value

- <a
  href="#VUID-VkCommandBufferInheritanceRenderingInfo-variableMultisampleRate-06005"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-variableMultisampleRate-06005"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-variableMultisampleRate-06005  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-variableMultisampleRate"
  target="_blank" rel="noopener"><code>variableMultisampleRate</code></a>
  feature is not enabled, `rasterizationSamples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value

- <a
  href="#VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06540"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06540"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06540  
  If `depthAttachmentFormat` is not `VK_FORMAT_UNDEFINED`, it **must**
  be a format that includes a depth component

- <a
  href="#VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06007"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06007"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06007  
  If `depthAttachmentFormat` is not `VK_FORMAT_UNDEFINED`, it **must**
  be a format with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> that
  include `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a
  href="#VUID-VkCommandBufferInheritanceRenderingInfo-pColorAttachmentFormats-06492"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-pColorAttachmentFormats-06492"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-pColorAttachmentFormats-06492  
  If any element of `pColorAttachmentFormats` is not
  `VK_FORMAT_UNDEFINED`, it **must** be a format with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> that
  include `VK_FORMAT_FEATURE_COLOR_ATTACHMENT_BIT` , or
  `VK_FORMAT_FEATURE_2_LINEAR_COLOR_ATTACHMENT_BIT_NV` if the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-linearColorAttachment"
  target="_blank" rel="noopener"><code>linearColorAttachment</code></a>
  feature is enabled

- <a
  href="#VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-06541"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-06541"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-06541  
  If `stencilAttachmentFormat` is not `VK_FORMAT_UNDEFINED`, it **must**
  be a format that includes a stencil aspect

- <a
  href="#VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-06199"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-06199"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-06199  
  If `stencilAttachmentFormat` is not `VK_FORMAT_UNDEFINED`, it **must**
  be a format with <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#potential-format-features"
  target="_blank" rel="noopener">potential format features</a> that
  include `VK_FORMAT_FEATURE_DEPTH_STENCIL_ATTACHMENT_BIT`

- <a
  href="#VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06200"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06200"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-06200  
  If `depthAttachmentFormat` is not `VK_FORMAT_UNDEFINED` and
  `stencilAttachmentFormat` is not `VK_FORMAT_UNDEFINED`,
  `depthAttachmentFormat` **must** equal `stencilAttachmentFormat`

- <a href="#VUID-VkCommandBufferInheritanceRenderingInfo-multiview-06008"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-multiview-06008"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-multiview-06008  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiview"
  target="_blank" rel="noopener"><code>multiview</code></a> feature is
  not enabled, `viewMask` **must** be `0`

- <a href="#VUID-VkCommandBufferInheritanceRenderingInfo-viewMask-06009"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-viewMask-06009"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-viewMask-06009  
  The index of the most significant bit in `viewMask` **must** be less
  than <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxMultiviewViewCount"
  target="_blank" rel="noopener"><code>maxMultiviewViewCount</code></a>

Valid Usage (Implicit)

- <a href="#VUID-VkCommandBufferInheritanceRenderingInfo-sType-sType"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-sType-sType"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_RENDERING_INFO`

- <a href="#VUID-VkCommandBufferInheritanceRenderingInfo-flags-parameter"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-flags-parameter"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkRenderingFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFlagBits.html) values

- <a
  href="#VUID-VkCommandBufferInheritanceRenderingInfo-pColorAttachmentFormats-parameter"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-pColorAttachmentFormats-parameter"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-pColorAttachmentFormats-parameter  
  If `colorAttachmentCount` is not `0`, `pColorAttachmentFormats`
  **must** be a valid pointer to an array of `colorAttachmentCount`
  valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) values

- <a
  href="#VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-parameter"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-parameter"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-depthAttachmentFormat-parameter  
  `depthAttachmentFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html)
  value

- <a
  href="#VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-parameter"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-parameter"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-stencilAttachmentFormat-parameter  
  `stencilAttachmentFormat` **must** be a valid
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a
  href="#VUID-VkCommandBufferInheritanceRenderingInfo-rasterizationSamples-parameter"
  id="VUID-VkCommandBufferInheritanceRenderingInfo-rasterizationSamples-parameter"></a>
  VUID-VkCommandBufferInheritanceRenderingInfo-rasterizationSamples-parameter  
  If `rasterizationSamples` is not `0`, `rasterizationSamples` **must**
  be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkRenderingFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFlags.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCommandBufferInheritanceRenderingInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
