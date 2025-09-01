# VkRenderingInputAttachmentIndexInfo(3) Manual Page

## Name

VkRenderingInputAttachmentIndexInfo - Structure specifying input attachment indices



## [](#_c_specification)C Specification

The `VkRenderingInputAttachmentIndexInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkRenderingInputAttachmentIndexInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           colorAttachmentCount;
    const uint32_t*    pColorAttachmentInputIndices;
    const uint32_t*    pDepthInputAttachmentIndex;
    const uint32_t*    pStencilInputAttachmentIndex;
} VkRenderingInputAttachmentIndexInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_dynamic_rendering_local_read
typedef VkRenderingInputAttachmentIndexInfo VkRenderingInputAttachmentIndexInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `colorAttachmentCount` is the number of elements in `pColorAttachmentInputIndices`.
- `pColorAttachmentInputIndices` is a pointer to an array of `colorAttachmentCount` `uint32_t` values defining indices for color attachments to be used as input attachments.
- `pDepthInputAttachmentIndex` is either `NULL`, or a pointer to a `uint32_t` value defining the index for the depth attachment to be used as an input attachment.
- `pStencilInputAttachmentIndex` is either `NULL`, or a pointer to a `uint32_t` value defining the index for the stencil attachment to be used as an input attachment.

## [](#_description)Description

This structure allows applications to remap attachments to different input attachment indices.

Each element of `pColorAttachmentInputIndices` set to a value of `VK_ATTACHMENT_UNUSED` indicates that the corresponding attachment will not be used as an input attachment in this pipeline. Any other value in each of those elements will map the corresponding attachment to a `InputAttachmentIndex` value defined in shader code.

If `pColorAttachmentInputIndices` is `NULL`, it is equivalent to setting each element to its index within the array.

If `pDepthInputAttachmentIndex` or `pStencilInputAttachmentIndex` are set to `NULL`, they map to input attachments without a `InputAttachmentIndex` decoration. If they point to a value of `VK_ATTACHMENT_UNUSED`, it indicates that the corresponding attachment will not be used as an input attachment in this pipeline. If they point to any other value it maps the corresponding attachment to a `InputAttachmentIndex` value defined in shader code.

This structure **can** be included in the `pNext` chain of a [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html) structure to set this state for a pipeline. If this structure is not included in the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), it is equivalent to specifying this structure with the following properties:

- `colorAttachmentCount` set to [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRenderingCreateInfo.html)::`colorAttachmentCount`.
- `pColorAttachmentInputIndices` set to `NULL`.
- `pDepthInputAttachmentIndex` set to `NULL`.
- `pStencilInputAttachmentIndex` set to `NULL`.

This structure **can** be included in the `pNext` chain of a [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html) structure to specify inherited state from the primary command buffer. If this structure is not included in the `pNext` chain of [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html), it is equivalent to specifying this structure with the following properties:

- `colorAttachmentCount` set to [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceRenderingInfo.html)::`colorAttachmentCount`.
- `pColorAttachmentInputIndices` set to `NULL`.
- `pDepthInputAttachmentIndex` set to `NULL`.
- `pStencilInputAttachmentIndex` set to `NULL`.

Valid Usage

- [](#VUID-VkRenderingInputAttachmentIndexInfo-dynamicRenderingLocalRead-09519)VUID-VkRenderingInputAttachmentIndexInfo-dynamicRenderingLocalRead-09519  
  If the [`dynamicRenderingLocalRead`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicRenderingLocalRead) feature is not enabled, and `pColorAttachmentInputIndices` is not `NULL`, each element **must** be `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkRenderingInputAttachmentIndexInfo-dynamicRenderingLocalRead-09520)VUID-VkRenderingInputAttachmentIndexInfo-dynamicRenderingLocalRead-09520  
  If the [`dynamicRenderingLocalRead`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicRenderingLocalRead) feature is not enabled, `pDepthInputAttachmentIndex` **must** be a valid pointer to a value of `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkRenderingInputAttachmentIndexInfo-dynamicRenderingLocalRead-09521)VUID-VkRenderingInputAttachmentIndexInfo-dynamicRenderingLocalRead-09521  
  If the [`dynamicRenderingLocalRead`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicRenderingLocalRead) feature is not enabled, `pStencilInputAttachmentIndex` **must** be a valid pointer to a value of `VK_ATTACHMENT_UNUSED`
- [](#VUID-VkRenderingInputAttachmentIndexInfo-pColorAttachmentInputIndices-09522)VUID-VkRenderingInputAttachmentIndexInfo-pColorAttachmentInputIndices-09522  
  Elements of `pColorAttachmentInputIndices` that are not `VK_ATTACHMENT_UNUSED` **must** each be unique
- [](#VUID-VkRenderingInputAttachmentIndexInfo-pColorAttachmentInputIndices-09523)VUID-VkRenderingInputAttachmentIndexInfo-pColorAttachmentInputIndices-09523  
  Elements of `pColorAttachmentInputIndices` that are not `VK_ATTACHMENT_UNUSED` **must** not take the same value as the content of `pDepthInputAttachmentIndex`
- [](#VUID-VkRenderingInputAttachmentIndexInfo-pColorAttachmentInputIndices-09524)VUID-VkRenderingInputAttachmentIndexInfo-pColorAttachmentInputIndices-09524  
  Elements of `pColorAttachmentInputIndices` that are not `VK_ATTACHMENT_UNUSED` **must** not take the same value as the content of `pStencilInputAttachmentIndex`
- [](#VUID-VkRenderingInputAttachmentIndexInfo-colorAttachmentCount-09525)VUID-VkRenderingInputAttachmentIndexInfo-colorAttachmentCount-09525  
  `colorAttachmentCount` **must** be less than or equal to [`maxColorAttachments`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxColorAttachments)

Valid Usage (Implicit)

- [](#VUID-VkRenderingInputAttachmentIndexInfo-sType-sType)VUID-VkRenderingInputAttachmentIndexInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDERING_INPUT_ATTACHMENT_INDEX_INFO`
- [](#VUID-VkRenderingInputAttachmentIndexInfo-pColorAttachmentInputIndices-parameter)VUID-VkRenderingInputAttachmentIndexInfo-pColorAttachmentInputIndices-parameter  
  If `colorAttachmentCount` is not `0`, and `pColorAttachmentInputIndices` is not `NULL`, `pColorAttachmentInputIndices` **must** be a valid pointer to an array of `colorAttachmentCount` `uint32_t` values
- [](#VUID-VkRenderingInputAttachmentIndexInfo-pDepthInputAttachmentIndex-parameter)VUID-VkRenderingInputAttachmentIndexInfo-pDepthInputAttachmentIndex-parameter  
  If `pDepthInputAttachmentIndex` is not `NULL`, `pDepthInputAttachmentIndex` **must** be a valid pointer to a valid `uint32_t` value
- [](#VUID-VkRenderingInputAttachmentIndexInfo-pStencilInputAttachmentIndex-parameter)VUID-VkRenderingInputAttachmentIndexInfo-pStencilInputAttachmentIndex-parameter  
  If `pStencilInputAttachmentIndex` is not `NULL`, `pStencilInputAttachmentIndex` **must** be a valid pointer to a valid `uint32_t` value

## [](#_see_also)See Also

[VK\_KHR\_dynamic\_rendering\_local\_read](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering_local_read.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdSetRenderingInputAttachmentIndices](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRenderingInputAttachmentIndices.html), [vkCmdSetRenderingInputAttachmentIndices](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRenderingInputAttachmentIndices.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderingInputAttachmentIndexInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0