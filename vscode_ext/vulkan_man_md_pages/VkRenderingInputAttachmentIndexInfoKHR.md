# VkRenderingInputAttachmentIndexInfoKHR(3) Manual Page

## Name

VkRenderingInputAttachmentIndexInfoKHR - Structure specifying input
attachment indices



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderingInputAttachmentIndexInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_dynamic_rendering_local_read
typedef struct VkRenderingInputAttachmentIndexInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           colorAttachmentCount;
    const uint32_t*    pColorAttachmentInputIndices;
    const uint32_t*    pDepthInputAttachmentIndex;
    const uint32_t*    pStencilInputAttachmentIndex;
} VkRenderingInputAttachmentIndexInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `colorAttachmentCount` is the number of elements in
  `pColorAttachmentInputIndices`.

- `pColorAttachmentInputIndices` is a pointer to an array of
  `colorAttachmentCount` `uint32_t` values defining indices for color
  attachments to be used as input attachments.

- `pDepthInputAttachmentIndex` is either `NULL`, or a pointer to a
  `uint32_t` value defining the index for the depth attachment to be
  used as an input attachment.

- `pStencilInputAttachmentIndex` is either `NULL`, or a pointer to a
  `uint32_t` value defining the index for the stencil attachment to be
  used as an input attachment.

## <a href="#_description" class="anchor"></a>Description

This structure allows applications to remap attachments to different
input attachment indices.

Each element of `pColorAttachmentInputIndices` set to a value of
`VK_ATTACHMENT_UNUSED` indicates that the corresponding attachment will
not be used as an input attachment in this pipeline. Any other value in
each of those elements will map the corresponding attachment to a
`InputAttachmentIndex` value defined in shader code.

If `pColorAttachmentInputIndices` is `NULL`, it is equivalent to setting
each element to its index within the array.

If `pDepthInputAttachmentIndex` or `pStencilInputAttachmentIndex` are
set to `NULL`, they map to input attachments without a
`InputAttachmentIndex` decoration. If they point to a value of
`VK_ATTACHMENT_UNUSED`, it indicates that the corresponding attachment
will not be used as an input attachment in this pipeline. If they point
to any other value it maps the corresponding attachment to a
`InputAttachmentIndex` value defined in shader code.

This structure **can** be included in the `pNext` chain of a
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)
structure to set this state for a pipeline. If this structure is not
included in the `pNext` chain of
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html), it is
equivalent to specifying this structure with the following properties:

- `colorAttachmentCount` set to
  [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRenderingCreateInfo.html)::`colorAttachmentCount`.

- `pColorAttachmentInputIndices` set to `NULL`.

- `pDepthInputAttachmentIndex` set to `NULL`.

- `pStencilInputAttachmentIndex` set to `NULL`.

This structure **can** be included in the `pNext` chain of a
[VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
structure to specify inherited state from the primary command buffer. If
this structure is not included in the `pNext` chain of
[VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html),
it is equivalent to specifying this structure with the following
properties:

- `colorAttachmentCount` set to
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html)::`colorAttachmentCount`.

- `pColorAttachmentInputIndices` set to `NULL`.

- `pDepthInputAttachmentIndex` set to `NULL`.

- `pStencilInputAttachmentIndex` set to `NULL`.

Valid Usage

- <a
  href="#VUID-VkRenderingInputAttachmentIndexInfoKHR-dynamicRenderingLocalRead-09519"
  id="VUID-VkRenderingInputAttachmentIndexInfoKHR-dynamicRenderingLocalRead-09519"></a>
  VUID-VkRenderingInputAttachmentIndexInfoKHR-dynamicRenderingLocalRead-09519  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicRenderingLocalRead"
  target="_blank"
  rel="noopener"><code>dynamicRenderingLocalRead</code></a> feature is
  not enabled, and `pColorAttachmentInputIndices` is not `NULL`, each
  element **must** be set to `VK_ATTACHMENT_UNUSED`

- <a
  href="#VUID-VkRenderingInputAttachmentIndexInfoKHR-dynamicRenderingLocalRead-09520"
  id="VUID-VkRenderingInputAttachmentIndexInfoKHR-dynamicRenderingLocalRead-09520"></a>
  VUID-VkRenderingInputAttachmentIndexInfoKHR-dynamicRenderingLocalRead-09520  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicRenderingLocalRead"
  target="_blank"
  rel="noopener"><code>dynamicRenderingLocalRead</code></a> feature is
  not enabled, `pDepthInputAttachmentIndex` **must** be a valid pointer
  to a value of `VK_ATTACHMENT_UNUSED`

- <a
  href="#VUID-VkRenderingInputAttachmentIndexInfoKHR-dynamicRenderingLocalRead-09521"
  id="VUID-VkRenderingInputAttachmentIndexInfoKHR-dynamicRenderingLocalRead-09521"></a>
  VUID-VkRenderingInputAttachmentIndexInfoKHR-dynamicRenderingLocalRead-09521  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-dynamicRenderingLocalRead"
  target="_blank"
  rel="noopener"><code>dynamicRenderingLocalRead</code></a> feature is
  not enabled, `pStencilInputAttachmentIndex` **must** be a valid
  pointer to a value of `VK_ATTACHMENT_UNUSED`

- <a
  href="#VUID-VkRenderingInputAttachmentIndexInfoKHR-pColorAttachmentInputIndices-09522"
  id="VUID-VkRenderingInputAttachmentIndexInfoKHR-pColorAttachmentInputIndices-09522"></a>
  VUID-VkRenderingInputAttachmentIndexInfoKHR-pColorAttachmentInputIndices-09522  
  Elements of `pColorAttachmentInputIndices` that are not
  `VK_ATTACHMENT_UNUSED` **must** each be unique

- <a
  href="#VUID-VkRenderingInputAttachmentIndexInfoKHR-pColorAttachmentInputIndices-09523"
  id="VUID-VkRenderingInputAttachmentIndexInfoKHR-pColorAttachmentInputIndices-09523"></a>
  VUID-VkRenderingInputAttachmentIndexInfoKHR-pColorAttachmentInputIndices-09523  
  Elements of `pColorAttachmentInputIndices` that are not
  `VK_ATTACHMENT_UNUSED` **must** not take the same value as the content
  of `pDepthInputAttachmentIndex`

- <a
  href="#VUID-VkRenderingInputAttachmentIndexInfoKHR-pColorAttachmentInputIndices-09524"
  id="VUID-VkRenderingInputAttachmentIndexInfoKHR-pColorAttachmentInputIndices-09524"></a>
  VUID-VkRenderingInputAttachmentIndexInfoKHR-pColorAttachmentInputIndices-09524  
  Elements of `pColorAttachmentInputIndices` that are not
  `VK_ATTACHMENT_UNUSED` **must** not take the same value as the content
  of `pStencilInputAttachmentIndex`

- <a
  href="#VUID-VkRenderingInputAttachmentIndexInfoKHR-colorAttachmentCount-09525"
  id="VUID-VkRenderingInputAttachmentIndexInfoKHR-colorAttachmentCount-09525"></a>
  VUID-VkRenderingInputAttachmentIndexInfoKHR-colorAttachmentCount-09525  
  `colorAttachmentCount` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxColorAttachments"
  target="_blank" rel="noopener"><code>maxColorAttachments</code></a>

Valid Usage (Implicit)

- <a href="#VUID-VkRenderingInputAttachmentIndexInfoKHR-sType-sType"
  id="VUID-VkRenderingInputAttachmentIndexInfoKHR-sType-sType"></a>
  VUID-VkRenderingInputAttachmentIndexInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RENDERING_INPUT_ATTACHMENT_INDEX_INFO_KHR`

- <a
  href="#VUID-VkRenderingInputAttachmentIndexInfoKHR-pColorAttachmentInputIndices-parameter"
  id="VUID-VkRenderingInputAttachmentIndexInfoKHR-pColorAttachmentInputIndices-parameter"></a>
  VUID-VkRenderingInputAttachmentIndexInfoKHR-pColorAttachmentInputIndices-parameter  
  If `colorAttachmentCount` is not `0`, and
  `pColorAttachmentInputIndices` is not `NULL`,
  `pColorAttachmentInputIndices` **must** be a valid pointer to an array
  of `colorAttachmentCount` `uint32_t` values

- <a
  href="#VUID-VkRenderingInputAttachmentIndexInfoKHR-pDepthInputAttachmentIndex-parameter"
  id="VUID-VkRenderingInputAttachmentIndexInfoKHR-pDepthInputAttachmentIndex-parameter"></a>
  VUID-VkRenderingInputAttachmentIndexInfoKHR-pDepthInputAttachmentIndex-parameter  
  If `pDepthInputAttachmentIndex` is not `NULL`,
  `pDepthInputAttachmentIndex` **must** be a valid pointer to a valid
  `uint32_t` value

- <a
  href="#VUID-VkRenderingInputAttachmentIndexInfoKHR-pStencilInputAttachmentIndex-parameter"
  id="VUID-VkRenderingInputAttachmentIndexInfoKHR-pStencilInputAttachmentIndex-parameter"></a>
  VUID-VkRenderingInputAttachmentIndexInfoKHR-pStencilInputAttachmentIndex-parameter  
  If `pStencilInputAttachmentIndex` is not `NULL`,
  `pStencilInputAttachmentIndex` **must** be a valid pointer to a valid
  `uint32_t` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_dynamic_rendering_local_read](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering_local_read.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdSetRenderingInputAttachmentIndicesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetRenderingInputAttachmentIndicesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderingInputAttachmentIndexInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
