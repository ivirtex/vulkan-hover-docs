# VkRenderingAreaInfoKHR(3) Manual Page

## Name

VkRenderingAreaInfoKHR - Structure describing rendering area granularity
query info



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderingAreaInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance5
typedef struct VkRenderingAreaInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           viewMask;
    uint32_t           colorAttachmentCount;
    const VkFormat*    pColorAttachmentFormats;
    VkFormat           depthAttachmentFormat;
    VkFormat           stencilAttachmentFormat;
} VkRenderingAreaInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `viewMask` is the viewMask used for rendering.

- `colorAttachmentCount` is the number of entries in
  `pColorAttachmentFormats`

- `pColorAttachmentFormats` is a pointer to an array of
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) values defining the format of color
  attachments used in the render pass instance.

- `depthAttachmentFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value defining
  the format of the depth attachment used in the render pass instance.

- `stencilAttachmentFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value
  defining the format of the stencil attachment used in the render pass
  instance.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkRenderingAreaInfoKHR-sType-sType"
  id="VUID-VkRenderingAreaInfoKHR-sType-sType"></a>
  VUID-VkRenderingAreaInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDERING_AREA_INFO_KHR`

- <a href="#VUID-VkRenderingAreaInfoKHR-pNext-pNext"
  id="VUID-VkRenderingAreaInfoKHR-pNext-pNext"></a>
  VUID-VkRenderingAreaInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html),
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetRenderingAreaGranularityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetRenderingAreaGranularityKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderingAreaInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
