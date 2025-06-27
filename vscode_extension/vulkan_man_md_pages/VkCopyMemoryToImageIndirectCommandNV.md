# VkCopyMemoryToImageIndirectCommandNV(3) Manual Page

## Name

VkCopyMemoryToImageIndirectCommandNV - Structure specifying indirect
buffer image copy operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkCopyMemoryToImageIndirectCommandNV` is defined as:

``` c
// Provided by VK_NV_copy_memory_indirect
typedef struct VkCopyMemoryToImageIndirectCommandNV {
    VkDeviceAddress             srcAddress;
    uint32_t                    bufferRowLength;
    uint32_t                    bufferImageHeight;
    VkImageSubresourceLayers    imageSubresource;
    VkOffset3D                  imageOffset;
    VkExtent3D                  imageExtent;
} VkCopyMemoryToImageIndirectCommandNV;
```

## <a href="#_members" class="anchor"></a>Members

- `srcAddress` is the starting address of the source device memory to
  copy from.

- `bufferRowLength` and `bufferImageHeight` specify in texels a
  subregion of a larger two- or three-dimensional image in buffer
  memory, and control the addressing calculations. If either of these
  values is zero, that aspect of the buffer memory is considered to be
  tightly packed according to the `imageExtent`.

- `imageSubresource` is a
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) used to
  specify the specific image subresources of the image used for the
  destination image data, which **must** match the values specified in
  `pImageSubresources` parameter of
  [vkCmdCopyMemoryToImageIndirectNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToImageIndirectNV.html)
  during command recording.

- `imageOffset` selects the initial `x`, `y`, `z` offsets in texels of
  the sub-region of the destination image data.

- `imageExtent` is the size in texels of the destination image in
  `width`, `height` and `depth`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkCopyMemoryToImageIndirectCommandNV-srcAddress-07678"
  id="VUID-VkCopyMemoryToImageIndirectCommandNV-srcAddress-07678"></a>
  VUID-VkCopyMemoryToImageIndirectCommandNV-srcAddress-07678  
  The `srcAddress` **must** be 4 byte aligned

- <a
  href="#VUID-VkCopyMemoryToImageIndirectCommandNV-bufferRowLength-07679"
  id="VUID-VkCopyMemoryToImageIndirectCommandNV-bufferRowLength-07679"></a>
  VUID-VkCopyMemoryToImageIndirectCommandNV-bufferRowLength-07679  
  `bufferRowLength` **must** be `0`, or greater than or equal to the
  `width` member of `imageExtent`

- <a
  href="#VUID-VkCopyMemoryToImageIndirectCommandNV-bufferImageHeight-07680"
  id="VUID-VkCopyMemoryToImageIndirectCommandNV-bufferImageHeight-07680"></a>
  VUID-VkCopyMemoryToImageIndirectCommandNV-bufferImageHeight-07680  
  `bufferImageHeight` **must** be `0`, or greater than or equal to the
  `height` member of `imageExtent`

- <a href="#VUID-VkCopyMemoryToImageIndirectCommandNV-imageOffset-07681"
  id="VUID-VkCopyMemoryToImageIndirectCommandNV-imageOffset-07681"></a>
  VUID-VkCopyMemoryToImageIndirectCommandNV-imageOffset-07681  
  `imageOffset` **must** specify a valid offset in the destination image

- <a href="#VUID-VkCopyMemoryToImageIndirectCommandNV-imageExtent-07682"
  id="VUID-VkCopyMemoryToImageIndirectCommandNV-imageExtent-07682"></a>
  VUID-VkCopyMemoryToImageIndirectCommandNV-imageExtent-07682  
  `imageExtent` **must** specify a valid region in the destination image
  and can be `0`

Valid Usage (Implicit)

- <a
  href="#VUID-VkCopyMemoryToImageIndirectCommandNV-imageSubresource-parameter"
  id="VUID-VkCopyMemoryToImageIndirectCommandNV-imageSubresource-parameter"></a>
  VUID-VkCopyMemoryToImageIndirectCommandNV-imageSubresource-parameter  
  `imageSubresource` **must** be a valid
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_copy_memory_indirect](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_copy_memory_indirect.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html),
[VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html),
[VkOffset3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset3D.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyMemoryToImageIndirectCommandNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
