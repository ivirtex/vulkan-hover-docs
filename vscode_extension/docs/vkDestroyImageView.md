# vkDestroyImageView(3) Manual Page

## Name

vkDestroyImageView - Destroy an image view object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy an image view, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyImageView(
    VkDevice                                    device,
    VkImageView                                 imageView,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the image view.

- `imageView` is the image view to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyImageView-imageView-01026"
  id="VUID-vkDestroyImageView-imageView-01026"></a>
  VUID-vkDestroyImageView-imageView-01026  
  All submitted commands that refer to `imageView` **must** have
  completed execution

- <a href="#VUID-vkDestroyImageView-imageView-01027"
  id="VUID-vkDestroyImageView-imageView-01027"></a>
  VUID-vkDestroyImageView-imageView-01027  
  If `VkAllocationCallbacks` were provided when `imageView` was created,
  a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyImageView-imageView-01028"
  id="VUID-vkDestroyImageView-imageView-01028"></a>
  VUID-vkDestroyImageView-imageView-01028  
  If no `VkAllocationCallbacks` were provided when `imageView` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyImageView-device-parameter"
  id="VUID-vkDestroyImageView-device-parameter"></a>
  VUID-vkDestroyImageView-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyImageView-imageView-parameter"
  id="VUID-vkDestroyImageView-imageView-parameter"></a>
  VUID-vkDestroyImageView-imageView-parameter  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle

- <a href="#VUID-vkDestroyImageView-pAllocator-parameter"
  id="VUID-vkDestroyImageView-pAllocator-parameter"></a>
  VUID-vkDestroyImageView-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyImageView-imageView-parent"
  id="VUID-vkDestroyImageView-imageView-parent"></a>
  VUID-vkDestroyImageView-imageView-parent  
  If `imageView` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `imageView` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyImageView"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
