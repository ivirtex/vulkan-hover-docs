# vkGetPhysicalDeviceSparseImageFormatProperties2(3) Manual Page

## Name

vkGetPhysicalDeviceSparseImageFormatProperties2 - Retrieve properties of
an image format applied to sparse images



## <a href="#_c_specification" class="anchor"></a>C Specification

`vkGetPhysicalDeviceSparseImageFormatProperties2` returns an array of
[VkSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties2.html).
Each element describes properties for one set of image aspects that are
bound simultaneously for a `VkImage` created with the provided image
creation parameters. This is usually one element for each aspect in the
image, but for interleaved depth/stencil images there is only one
element describing the combined aspects.

``` c
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceSparseImageFormatProperties2(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceSparseImageFormatInfo2* pFormatInfo,
    uint32_t*                                   pPropertyCount,
    VkSparseImageFormatProperties2*             pProperties);
```

or the equivalent command

``` c
// Provided by VK_KHR_get_physical_device_properties2
void vkGetPhysicalDeviceSparseImageFormatProperties2KHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceSparseImageFormatInfo2* pFormatInfo,
    uint32_t*                                   pPropertyCount,
    VkSparseImageFormatProperties2*             pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the sparse
  image format properties.

- `pFormatInfo` is a pointer to a
  [VkPhysicalDeviceSparseImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSparseImageFormatInfo2.html)
  structure containing input parameters to the command.

- `pPropertyCount` is a pointer to an integer related to the number of
  sparse format properties available or queried, as described below.

- `pProperties` is either `NULL` or a pointer to an array of
  [VkSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties2.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

`vkGetPhysicalDeviceSparseImageFormatProperties2` behaves identically to
[vkGetPhysicalDeviceSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSparseImageFormatProperties.html),
with the ability to return extended information by adding extending
structures to the `pNext` chain of its `pProperties` parameter.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pFormatInfo-parameter"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pFormatInfo-parameter"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pFormatInfo-parameter  
  `pFormatInfo` **must** be a valid pointer to a valid
  [VkPhysicalDeviceSparseImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSparseImageFormatInfo2.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pPropertyCount-parameter"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pPropertyCount-parameter"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pProperties-parameter"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties2-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pPropertyCount`
  [VkSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties2.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceSparseImageFormatInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSparseImageFormatInfo2.html),
[VkSparseImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceSparseImageFormatProperties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
