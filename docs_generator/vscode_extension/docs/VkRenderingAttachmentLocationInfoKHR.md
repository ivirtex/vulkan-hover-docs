# VkRenderingAttachmentLocationInfo(3) Manual Page

## Name

VkRenderingAttachmentLocationInfo - Structure specifying attachment locations



## [](#_c_specification)C Specification

The `VkRenderingAttachmentLocationInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkRenderingAttachmentLocationInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           colorAttachmentCount;
    const uint32_t*    pColorAttachmentLocations;
} VkRenderingAttachmentLocationInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_dynamic_rendering_local_read
typedef VkRenderingAttachmentLocationInfo VkRenderingAttachmentLocationInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `colorAttachmentCount` is the number of elements in `pColorAttachmentLocations`.
- `pColorAttachmentLocations` is a pointer to an array of `colorAttachmentCount` `uint32_t` values defining remapped locations for color attachments.

## [](#_description)Description

This structure allows applications to remap the locations of color attachments to different fragment shader output locations.

Each element of `pColorAttachmentLocations` set to `VK_ATTACHMENT_UNUSED` will be inaccessible to this pipeline as a color attachment; no location will map to it. Each element of `pColorAttachmentLocations` set to any other value will map the specified location value to the color attachment specified in the render pass at the corresponding index in the `pColorAttachmentLocations` array. Any writes to a fragment output location that is not mapped to an attachment **must** be discarded.

If `pColorAttachmentLocations` is `NULL`, it is equivalent to setting each element to its index within the array.

This structure **can** be included in the `pNext` chain of a [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html) structure to set this state for a pipeline. If this structure is not included in the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), it is equivalent to specifying this structure with the following properties:

- `colorAttachmentCount` set to [VkPipelineRenderingCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRenderingCreateInfo.html)::`colorAttachmentCount`.
- `pColorAttachmentLocations` set to `NULL`.

This structure **can** be included in the `pNext` chain of a [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html) structure to specify inherited state from the primary command buffer. If [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html)::renderPass is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), or `VK_COMMAND_BUFFER_USAGE_RENDER_PASS_CONTINUE_BIT` is not specified in [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferBeginInfo.html)::flags, members of this structure are ignored. If this structure is not included in the `pNext` chain of [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html), it is equivalent to specifying this structure with the following properties:

- `colorAttachmentCount` set to [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceRenderingInfo.html)::`colorAttachmentCount`.
- `pColorAttachmentLocations` set to `NULL`.

Valid Usage

- [](#VUID-VkRenderingAttachmentLocationInfo-dynamicRenderingLocalRead-09512)VUID-VkRenderingAttachmentLocationInfo-dynamicRenderingLocalRead-09512  
  If the [`dynamicRenderingLocalRead`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-dynamicRenderingLocalRead) feature is not enabled, and `pColorAttachmentLocations` is not `NULL`, each element **must** be the value of its index within the array
- [](#VUID-VkRenderingAttachmentLocationInfo-pColorAttachmentLocations-09513)VUID-VkRenderingAttachmentLocationInfo-pColorAttachmentLocations-09513  
  Elements of `pColorAttachmentLocations` that are not `VK_ATTACHMENT_UNUSED` **must** each be unique
- [](#VUID-VkRenderingAttachmentLocationInfo-colorAttachmentCount-09514)VUID-VkRenderingAttachmentLocationInfo-colorAttachmentCount-09514  
  `colorAttachmentCount` **must** be less than or equal to [`maxColorAttachments`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxColorAttachments)
- [](#VUID-VkRenderingAttachmentLocationInfo-pColorAttachmentLocations-09515)VUID-VkRenderingAttachmentLocationInfo-pColorAttachmentLocations-09515  
  Each element of `pColorAttachmentLocations` **must** be less than [`maxColorAttachments`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxColorAttachments)

Valid Usage (Implicit)

- [](#VUID-VkRenderingAttachmentLocationInfo-sType-sType)VUID-VkRenderingAttachmentLocationInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDERING_ATTACHMENT_LOCATION_INFO`

## [](#_see_also)See Also

[VK\_KHR\_dynamic\_rendering\_local\_read](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering_local_read.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdSetRenderingAttachmentLocations](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRenderingAttachmentLocations.html), [vkCmdSetRenderingAttachmentLocationsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetRenderingAttachmentLocationsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderingAttachmentLocationInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0