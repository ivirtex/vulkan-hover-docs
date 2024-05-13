# vkDestroyImage(3) Manual Page

## Name

vkDestroyImage - Destroy an image object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy an image, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyImage(
    VkDevice                                    device,
    VkImage                                     image,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the image.

- `image` is the image to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyImage-image-01000"
  id="VUID-vkDestroyImage-image-01000"></a>
  VUID-vkDestroyImage-image-01000  
  All submitted commands that refer to `image`, either directly or via a
  `VkImageView`, **must** have completed execution

- <a href="#VUID-vkDestroyImage-image-01001"
  id="VUID-vkDestroyImage-image-01001"></a>
  VUID-vkDestroyImage-image-01001  
  If `VkAllocationCallbacks` were provided when `image` was created, a
  compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyImage-image-01002"
  id="VUID-vkDestroyImage-image-01002"></a>
  VUID-vkDestroyImage-image-01002  
  If no `VkAllocationCallbacks` were provided when `image` was created,
  `pAllocator` **must** be `NULL`

- <a href="#VUID-vkDestroyImage-image-04882"
  id="VUID-vkDestroyImage-image-04882"></a>
  VUID-vkDestroyImage-image-04882  
  `image` **must** not have been acquired from
  [vkGetSwapchainImagesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetSwapchainImagesKHR.html)

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyImage-device-parameter"
  id="VUID-vkDestroyImage-device-parameter"></a>
  VUID-vkDestroyImage-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyImage-image-parameter"
  id="VUID-vkDestroyImage-image-parameter"></a>
  VUID-vkDestroyImage-image-parameter  
  If `image` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `image`
  **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-vkDestroyImage-pAllocator-parameter"
  id="VUID-vkDestroyImage-pAllocator-parameter"></a>
  VUID-vkDestroyImage-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyImage-image-parent"
  id="VUID-vkDestroyImage-image-parent"></a>
  VUID-vkDestroyImage-image-parent  
  If `image` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `image` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyImage"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
