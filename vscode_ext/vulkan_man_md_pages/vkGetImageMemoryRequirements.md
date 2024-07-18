# vkGetImageMemoryRequirements(3) Manual Page

## Name

vkGetImageMemoryRequirements - Returns the memory requirements for
specified Vulkan object



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the memory requirements for an image resource which is not
created with the `VK_IMAGE_CREATE_DISJOINT_BIT` flag set, call:

``` c
// Provided by VK_VERSION_1_0
void vkGetImageMemoryRequirements(
    VkDevice                                    device,
    VkImage                                     image,
    VkMemoryRequirements*                       pMemoryRequirements);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the image.

- `image` is the image to query.

- `pMemoryRequirements` is a pointer to a
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure in which
  the memory requirements of the image object are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetImageMemoryRequirements-image-01588"
  id="VUID-vkGetImageMemoryRequirements-image-01588"></a>
  VUID-vkGetImageMemoryRequirements-image-01588  
  `image` **must** not have been created with the
  `VK_IMAGE_CREATE_DISJOINT_BIT` flag set

- <a href="#VUID-vkGetImageMemoryRequirements-image-04004"
  id="VUID-vkGetImageMemoryRequirements-image-04004"></a>
  VUID-vkGetImageMemoryRequirements-image-04004  
  If `image` was created with the
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`
  external memory handle type, then `image` **must** be bound to memory

- <a href="#VUID-vkGetImageMemoryRequirements-image-08960"
  id="VUID-vkGetImageMemoryRequirements-image-08960"></a>
  VUID-vkGetImageMemoryRequirements-image-08960  
  If `image` was created with the
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_SCREEN_BUFFER_BIT_QNX` external memory
  handle type, then `image` **must** be bound to memory

Valid Usage (Implicit)

- <a href="#VUID-vkGetImageMemoryRequirements-device-parameter"
  id="VUID-vkGetImageMemoryRequirements-device-parameter"></a>
  VUID-vkGetImageMemoryRequirements-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetImageMemoryRequirements-image-parameter"
  id="VUID-vkGetImageMemoryRequirements-image-parameter"></a>
  VUID-vkGetImageMemoryRequirements-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a
  href="#VUID-vkGetImageMemoryRequirements-pMemoryRequirements-parameter"
  id="VUID-vkGetImageMemoryRequirements-pMemoryRequirements-parameter"></a>
  VUID-vkGetImageMemoryRequirements-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure

- <a href="#VUID-vkGetImageMemoryRequirements-image-parent"
  id="VUID-vkGetImageMemoryRequirements-image-parent"></a>
  VUID-vkGetImageMemoryRequirements-image-parent  
  `image` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetImageMemoryRequirements"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
