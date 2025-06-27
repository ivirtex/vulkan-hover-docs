# vkCreateImageView(3) Manual Page

## Name

vkCreateImageView - Create an image view from an existing image



## <a href="#_c_specification" class="anchor"></a>C Specification

To create an image view, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkCreateImageView(
    VkDevice                                    device,
    const VkImageViewCreateInfo*                pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkImageView*                                pView);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the image view.

- `pCreateInfo` is a pointer to a `VkImageViewCreateInfo` structure
  containing parameters to be used to create the image view.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pView` is a pointer to a [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle in
  which the resulting image view object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCreateImageView-device-09667"
  id="VUID-vkCreateImageView-device-09667"></a>
  VUID-vkCreateImageView-device-09667  
  `device` **must** support at least one queue family with one of the
  `VK_QUEUE_VIDEO_ENCODE_BIT_KHR`, `VK_QUEUE_VIDEO_DECODE_BIT_KHR`,
  `VK_QUEUE_COMPUTE_BIT`, or `VK_QUEUE_GRAPHICS_BIT` capabilities

- <a href="#VUID-vkCreateImageView-image-09179"
  id="VUID-vkCreateImageView-image-09179"></a>
  VUID-vkCreateImageView-image-09179  
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html)::`image` **must**
  have been created from `device`

Valid Usage (Implicit)

- <a href="#VUID-vkCreateImageView-device-parameter"
  id="VUID-vkCreateImageView-device-parameter"></a>
  VUID-vkCreateImageView-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateImageView-pCreateInfo-parameter"
  id="VUID-vkCreateImageView-pCreateInfo-parameter"></a>
  VUID-vkCreateImageView-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html) structure

- <a href="#VUID-vkCreateImageView-pAllocator-parameter"
  id="VUID-vkCreateImageView-pAllocator-parameter"></a>
  VUID-vkCreateImageView-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateImageView-pView-parameter"
  id="VUID-vkCreateImageView-pView-parameter"></a>
  VUID-vkCreateImageView-pView-parameter  
  `pView` **must** be a valid pointer to a
  [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html),
[VkImageViewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateImageView"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
