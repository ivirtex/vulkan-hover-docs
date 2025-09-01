# VkRenderingAreaInfo(3) Manual Page

## Name

VkRenderingAreaInfo - Structure describing rendering area granularity query info



## [](#_c_specification)C Specification

The `VkRenderingAreaInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_4
typedef struct VkRenderingAreaInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           viewMask;
    uint32_t           colorAttachmentCount;
    const VkFormat*    pColorAttachmentFormats;
    VkFormat           depthAttachmentFormat;
    VkFormat           stencilAttachmentFormat;
} VkRenderingAreaInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_maintenance5
typedef VkRenderingAreaInfo VkRenderingAreaInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `viewMask` is the viewMask used for rendering.
- `colorAttachmentCount` is the number of entries in `pColorAttachmentFormats`
- `pColorAttachmentFormats` is a pointer to an array of [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) values defining the format of color attachments used in the render pass instance.
- `depthAttachmentFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value defining the format of the depth attachment used in the render pass instance.
- `stencilAttachmentFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value defining the format of the stencil attachment used in the render pass instance.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkRenderingAreaInfo-sType-sType)VUID-VkRenderingAreaInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDERING_AREA_INFO`
- [](#VUID-VkRenderingAreaInfo-pNext-pNext)VUID-VkRenderingAreaInfo-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetRenderingAreaGranularity](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRenderingAreaGranularity.html), [vkGetRenderingAreaGranularity](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRenderingAreaGranularity.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderingAreaInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0