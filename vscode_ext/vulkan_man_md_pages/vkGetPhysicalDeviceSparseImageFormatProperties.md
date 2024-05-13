# vkGetPhysicalDeviceSparseImageFormatProperties(3) Manual Page

## Name

vkGetPhysicalDeviceSparseImageFormatProperties - Retrieve properties of
an image format applied to sparse images



## <a href="#_c_specification" class="anchor"></a>C Specification

`vkGetPhysicalDeviceSparseImageFormatProperties` returns an array of
[VkSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties.html).
Each element describes properties for one set of image aspects that are
bound simultaneously for a `VkImage` created with the provided image
creation parameters. This is usually one element for each aspect in the
image, but for interleaved depth/stencil images there is only one
element describing the combined aspects.

``` c
// Provided by VK_VERSION_1_0
void vkGetPhysicalDeviceSparseImageFormatProperties(
    VkPhysicalDevice                            physicalDevice,
    VkFormat                                    format,
    VkImageType                                 type,
    VkSampleCountFlagBits                       samples,
    VkImageUsageFlags                           usage,
    VkImageTiling                               tiling,
    uint32_t*                                   pPropertyCount,
    VkSparseImageFormatProperties*              pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the sparse
  image format properties.

- `format` is the image format.

- `type` is the dimensionality of the image.

- `samples` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html)
  value specifying the number of samples per texel.

- `usage` is a bitmask describing the intended usage of the image.

- `tiling` is the tiling arrangement of the texel blocks in memory.

- `pPropertyCount` is a pointer to an integer related to the number of
  sparse format properties available or queried, as described below.

- `pProperties` is either `NULL` or a pointer to an array of
  [VkSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

If `pProperties` is `NULL`, then the number of sparse format properties
available is returned in `pPropertyCount`. Otherwise, `pPropertyCount`
**must** point to a variable set by the user to the number of elements
in the `pProperties` array, and on return the variable is overwritten
with the number of structures actually written to `pProperties`. If
`pPropertyCount` is less than the number of sparse format properties
available, at most `pPropertyCount` structures will be written.

If `VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT` is not supported for the given
arguments, `pPropertyCount` will be set to zero upon return, and no data
will be written to `pProperties`.

Multiple aspects are returned for depth/stencil images that are
implemented as separate planes by the implementation. The depth and
stencil data planes each have unique `VkSparseImageFormatProperties`
data.

Depth/stencil images with depth and stencil data interleaved into a
single plane will return a single `VkSparseImageFormatProperties`
structure with the `aspectMask` set to `VK_IMAGE_ASPECT_DEPTH_BIT` \|
`VK_IMAGE_ASPECT_STENCIL_BIT`.

Valid Usage

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties-samples-01094"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties-samples-01094"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties-samples-01094  
  `samples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value that is set
  in `VkImageFormatProperties`::`sampleCounts` returned by
  `vkGetPhysicalDeviceImageFormatProperties` with `format`, `type`,
  `tiling`, and `usage` equal to those in this command

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties-format-parameter"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties-format-parameter"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties-type-parameter"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties-type-parameter"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties-type-parameter  
  `type` **must** be a valid [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html) value

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties-samples-parameter"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties-samples-parameter"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties-samples-parameter  
  `samples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties-usage-parameter"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties-usage-parameter"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties-usage-parameter  
  `usage` **must** be a valid combination of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) values

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties-usage-requiredbitmask"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties-usage-requiredbitmask"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties-usage-requiredbitmask  
  `usage` **must** not be `0`

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties-tiling-parameter"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties-tiling-parameter"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties-tiling-parameter  
  `tiling` **must** be a valid [VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html) value

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties-pPropertyCount-parameter"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties-pPropertyCount-parameter"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties-pPropertyCount-parameter  
  `pPropertyCount` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceSparseImageFormatProperties-pProperties-parameter"
  id="VUID-vkGetPhysicalDeviceSparseImageFormatProperties-pProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceSparseImageFormatProperties-pProperties-parameter  
  If the value referenced by `pPropertyCount` is not `0`, and
  `pProperties` is not `NULL`, `pProperties` **must** be a valid pointer
  to an array of `pPropertyCount`
  [VkSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties.html)
  structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html), [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html),
[VkSparseImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceSparseImageFormatProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
