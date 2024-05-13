# VkReleaseSwapchainImagesInfoEXT(3) Manual Page

## Name

VkReleaseSwapchainImagesInfoEXT - Structure describing a list of
swapchain image indices to be released



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkReleaseSwapchainImagesInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_swapchain_maintenance1
typedef struct VkReleaseSwapchainImagesInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkSwapchainKHR     swapchain;
    uint32_t           imageIndexCount;
    const uint32_t*    pImageIndices;
} VkReleaseSwapchainImagesInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `swapchain` is a swapchain to which images are being released.

- `imageIndexCount` is the number of image indices to be released.

- `pImageIndices` is a pointer to an array of indices into the array of
  `swapchain`’s presentable images, with `imageIndexCount` entries.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkReleaseSwapchainImagesInfoEXT-pImageIndices-07785"
  id="VUID-VkReleaseSwapchainImagesInfoEXT-pImageIndices-07785"></a>
  VUID-VkReleaseSwapchainImagesInfoEXT-pImageIndices-07785  
  Each element of `pImageIndices` **must** be the index of a presentable
  image acquired from the swapchain specified by `swapchain`

- <a href="#VUID-VkReleaseSwapchainImagesInfoEXT-pImageIndices-07786"
  id="VUID-VkReleaseSwapchainImagesInfoEXT-pImageIndices-07786"></a>
  VUID-VkReleaseSwapchainImagesInfoEXT-pImageIndices-07786  
  All uses of presentable images identified by elements of
  `pImageIndices` **must** have completed execution

Valid Usage (Implicit)

- <a href="#VUID-VkReleaseSwapchainImagesInfoEXT-sType-sType"
  id="VUID-VkReleaseSwapchainImagesInfoEXT-sType-sType"></a>
  VUID-VkReleaseSwapchainImagesInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RELEASE_SWAPCHAIN_IMAGES_INFO_EXT`

- <a href="#VUID-VkReleaseSwapchainImagesInfoEXT-pNext-pNext"
  id="VUID-VkReleaseSwapchainImagesInfoEXT-pNext-pNext"></a>
  VUID-VkReleaseSwapchainImagesInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkReleaseSwapchainImagesInfoEXT-swapchain-parameter"
  id="VUID-VkReleaseSwapchainImagesInfoEXT-swapchain-parameter"></a>
  VUID-VkReleaseSwapchainImagesInfoEXT-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-VkReleaseSwapchainImagesInfoEXT-pImageIndices-parameter"
  id="VUID-VkReleaseSwapchainImagesInfoEXT-pImageIndices-parameter"></a>
  VUID-VkReleaseSwapchainImagesInfoEXT-pImageIndices-parameter  
  `pImageIndices` **must** be a valid pointer to an array of
  `imageIndexCount` `uint32_t` values

- <a
  href="#VUID-VkReleaseSwapchainImagesInfoEXT-imageIndexCount-arraylength"
  id="VUID-VkReleaseSwapchainImagesInfoEXT-imageIndexCount-arraylength"></a>
  VUID-VkReleaseSwapchainImagesInfoEXT-imageIndexCount-arraylength  
  `imageIndexCount` **must** be greater than `0`

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_swapchain_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_swapchain_maintenance1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html),
[vkReleaseSwapchainImagesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkReleaseSwapchainImagesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkReleaseSwapchainImagesInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
