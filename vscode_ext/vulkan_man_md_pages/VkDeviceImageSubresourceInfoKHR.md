# VkDeviceImageSubresourceInfoKHR(3) Manual Page

## Name

VkDeviceImageSubresourceInfoKHR - Image creation information for
querying subresource layout



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceImageSubresourceInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_maintenance5
typedef struct VkDeviceImageSubresourceInfoKHR {
    VkStructureType                  sType;
    const void*                      pNext;
    const VkImageCreateInfo*         pCreateInfo;
    const VkImageSubresource2KHR*    pSubresource;
} VkDeviceImageSubresourceInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pCreateInfo` is a pointer to a
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure containing
  parameters affecting creation of the image to query.

- `pSubresource` pSubresource is a pointer to a
  [VkImageSubresource2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource2KHR.html) structure
  selecting a specific image subresource for the query.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkDeviceImageSubresourceInfoKHR-aspectMask-00997"
  id="VUID-VkDeviceImageSubresourceInfoKHR-aspectMask-00997"></a>
  VUID-VkDeviceImageSubresourceInfoKHR-aspectMask-00997  
  The `aspectMask` member of `pSubresource` **must** only have a single
  bit set

- <a href="#VUID-VkDeviceImageSubresourceInfoKHR-mipLevel-01716"
  id="VUID-VkDeviceImageSubresourceInfoKHR-mipLevel-01716"></a>
  VUID-VkDeviceImageSubresourceInfoKHR-mipLevel-01716  
  The `mipLevel` member of `pSubresource` **must** be less than the
  `mipLevels` specified in `pCreateInfo`

- <a href="#VUID-VkDeviceImageSubresourceInfoKHR-arrayLayer-01717"
  id="VUID-VkDeviceImageSubresourceInfoKHR-arrayLayer-01717"></a>
  VUID-VkDeviceImageSubresourceInfoKHR-arrayLayer-01717  
  The `arrayLayer` member of `pSubresource` **must** be less than the
  `arrayLayers` specified in `pCreateInfo`

- <a href="#VUID-VkDeviceImageSubresourceInfoKHR-format-08886"
  id="VUID-VkDeviceImageSubresourceInfoKHR-format-08886"></a>
  VUID-VkDeviceImageSubresourceInfoKHR-format-08886  
  If `format` of the `image` is a color format that is not a
  [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), and `tiling` of
  the `pCreateInfo` is `VK_IMAGE_TILING_LINEAR` or
  `VK_IMAGE_TILING_OPTIMAL`, the `aspectMask` member of `pSubresource`
  **must** be `VK_IMAGE_ASPECT_COLOR_BIT`

- <a href="#VUID-VkDeviceImageSubresourceInfoKHR-format-04462"
  id="VUID-VkDeviceImageSubresourceInfoKHR-format-04462"></a>
  VUID-VkDeviceImageSubresourceInfoKHR-format-04462  
  If `format` of the `pCreateInfo` has a depth component, the
  `aspectMask` member of `pSubresource` **must** contain
  `VK_IMAGE_ASPECT_DEPTH_BIT`

- <a href="#VUID-VkDeviceImageSubresourceInfoKHR-format-04463"
  id="VUID-VkDeviceImageSubresourceInfoKHR-format-04463"></a>
  VUID-VkDeviceImageSubresourceInfoKHR-format-04463  
  If `format` of the `pCreateInfo` has a stencil component, the
  `aspectMask` member of `pSubresource` **must** contain
  `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-VkDeviceImageSubresourceInfoKHR-format-04464"
  id="VUID-VkDeviceImageSubresourceInfoKHR-format-04464"></a>
  VUID-VkDeviceImageSubresourceInfoKHR-format-04464  
  If `format` of the `pCreateInfo` does not contain a stencil or depth
  component, the `aspectMask` member of `pSubresource` **must** not
  contain `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT`

- <a href="#VUID-VkDeviceImageSubresourceInfoKHR-tiling-08717"
  id="VUID-VkDeviceImageSubresourceInfoKHR-tiling-08717"></a>
  VUID-VkDeviceImageSubresourceInfoKHR-tiling-08717  
  If the `tiling` of the `pCreateInfo` is `VK_IMAGE_TILING_LINEAR` and
  has a [multi-planar image
  format](#formats-requiring-sampler-ycbcr-conversion), then the
  `aspectMask` member of `pSubresource` **must** be a single valid
  [multi-planar aspect mask](#formats-planes-image-aspect) bit

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceImageSubresourceInfoKHR-sType-sType"
  id="VUID-VkDeviceImageSubresourceInfoKHR-sType-sType"></a>
  VUID-VkDeviceImageSubresourceInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_IMAGE_SUBRESOURCE_INFO_KHR`

- <a href="#VUID-VkDeviceImageSubresourceInfoKHR-pNext-pNext"
  id="VUID-VkDeviceImageSubresourceInfoKHR-pNext-pNext"></a>
  VUID-VkDeviceImageSubresourceInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDeviceImageSubresourceInfoKHR-pCreateInfo-parameter"
  id="VUID-VkDeviceImageSubresourceInfoKHR-pCreateInfo-parameter"></a>
  VUID-VkDeviceImageSubresourceInfoKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) structure

- <a href="#VUID-VkDeviceImageSubresourceInfoKHR-pSubresource-parameter"
  id="VUID-VkDeviceImageSubresourceInfoKHR-pSubresource-parameter"></a>
  VUID-VkDeviceImageSubresourceInfoKHR-pSubresource-parameter  
  `pSubresource` **must** be a valid pointer to a valid
  [VkImageSubresource2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource2KHR.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance5](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance5.html),
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html),
[VkImageSubresource2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresource2KHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDeviceImageSubresourceLayoutKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceImageSubresourceLayoutKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceImageSubresourceInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
