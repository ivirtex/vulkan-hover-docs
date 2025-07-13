# VkReleaseSwapchainImagesInfoKHR(3) Manual Page

## Name

VkReleaseSwapchainImagesInfoKHR - Structure describing a list of swapchain image indices to be released



## [](#_c_specification)C Specification

The `VkReleaseSwapchainImagesInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_swapchain_maintenance1
typedef struct VkReleaseSwapchainImagesInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkSwapchainKHR     swapchain;
    uint32_t           imageIndexCount;
    const uint32_t*    pImageIndices;
} VkReleaseSwapchainImagesInfoKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_swapchain_maintenance1
typedef VkReleaseSwapchainImagesInfoKHR VkReleaseSwapchainImagesInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `swapchain` is a swapchain to which images are being released.
- `imageIndexCount` is the number of image indices to be released.
- `pImageIndices` is a pointer to an array of indices into the array of `swapchain`’s presentable images, with `imageIndexCount` entries.

## [](#_description)Description

Valid Usage

- [](#VUID-VkReleaseSwapchainImagesInfoKHR-pImageIndices-07785)VUID-VkReleaseSwapchainImagesInfoKHR-pImageIndices-07785  
  Each element of `pImageIndices` **must** be the index of a presentable image acquired from the swapchain specified by `swapchain`
- [](#VUID-VkReleaseSwapchainImagesInfoKHR-pImageIndices-07786)VUID-VkReleaseSwapchainImagesInfoKHR-pImageIndices-07786  
  All uses of presentable images identified by elements of `pImageIndices` **must** have completed execution

Valid Usage (Implicit)

- [](#VUID-VkReleaseSwapchainImagesInfoKHR-sType-sType)VUID-VkReleaseSwapchainImagesInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RELEASE_SWAPCHAIN_IMAGES_INFO_KHR`
- [](#VUID-VkReleaseSwapchainImagesInfoKHR-pNext-pNext)VUID-VkReleaseSwapchainImagesInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkReleaseSwapchainImagesInfoKHR-swapchain-parameter)VUID-VkReleaseSwapchainImagesInfoKHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-VkReleaseSwapchainImagesInfoKHR-pImageIndices-parameter)VUID-VkReleaseSwapchainImagesInfoKHR-pImageIndices-parameter  
  `pImageIndices` **must** be a valid pointer to an array of `imageIndexCount` `uint32_t` values
- [](#VUID-VkReleaseSwapchainImagesInfoKHR-imageIndexCount-arraylength)VUID-VkReleaseSwapchainImagesInfoKHR-imageIndexCount-arraylength  
  `imageIndexCount` **must** be greater than `0`

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_EXT\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_swapchain_maintenance1.html), [VK\_KHR\_swapchain\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain_maintenance1.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html), [vkReleaseSwapchainImagesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseSwapchainImagesEXT.html), [vkReleaseSwapchainImagesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseSwapchainImagesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkReleaseSwapchainImagesInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0