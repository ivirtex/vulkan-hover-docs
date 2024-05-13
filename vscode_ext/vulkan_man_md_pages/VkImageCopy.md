# VkImageCopy(3) Manual Page

## Name

VkImageCopy - Structure specifying an image copy operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkImageCopy` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkImageCopy {
    VkImageSubresourceLayers    srcSubresource;
    VkOffset3D                  srcOffset;
    VkImageSubresourceLayers    dstSubresource;
    VkOffset3D                  dstOffset;
    VkExtent3D                  extent;
} VkImageCopy;
```

## <a href="#_members" class="anchor"></a>Members

- `srcSubresource` and `dstSubresource` are
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) structures
  specifying the image subresources of the images used for the source
  and destination image data, respectively.

- `srcOffset` and `dstOffset` select the initial `x`, `y`, and `z`
  offsets in texels of the sub-regions of the source and destination
  image data.

- `extent` is the size in texels of the image to copy in `width`,
  `height` and `depth`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkImageCopy-apiVersion-07940"
  id="VUID-VkImageCopy-apiVersion-07940"></a>
  VUID-VkImageCopy-apiVersion-07940  
  If the
  [VK_KHR_sampler_ycbcr_conversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_sampler_ycbcr_conversion.html)
  extension is not enabled, and
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, the `aspectMask` member of `srcSubresource`
  and `dstSubresource` **must** match

- <a href="#VUID-VkImageCopy-apiVersion-07941"
  id="VUID-VkImageCopy-apiVersion-07941"></a>
  VUID-VkImageCopy-apiVersion-07941  
  If the [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html) extension is
  not enabled, and
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, the `layerCount` member of `srcSubresource`
  and `dstSubresource` **must** match

- <a href="#VUID-VkImageCopy-extent-06668"
  id="VUID-VkImageCopy-extent-06668"></a>
  VUID-VkImageCopy-extent-06668  
  `extent.width` **must** not be 0

- <a href="#VUID-VkImageCopy-extent-06669"
  id="VUID-VkImageCopy-extent-06669"></a>
  VUID-VkImageCopy-extent-06669  
  `extent.height` **must** not be 0

- <a href="#VUID-VkImageCopy-extent-06670"
  id="VUID-VkImageCopy-extent-06670"></a>
  VUID-VkImageCopy-extent-06670  
  `extent.depth` **must** not be 0

Valid Usage (Implicit)

- <a href="#VUID-VkImageCopy-srcSubresource-parameter"
  id="VUID-VkImageCopy-srcSubresource-parameter"></a>
  VUID-VkImageCopy-srcSubresource-parameter  
  `srcSubresource` **must** be a valid
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) structure

- <a href="#VUID-VkImageCopy-dstSubresource-parameter"
  id="VUID-VkImageCopy-dstSubresource-parameter"></a>
  VUID-VkImageCopy-dstSubresource-parameter  
  `dstSubresource` **must** be a valid
  [VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkExtent3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent3D.html),
[VkImageSubresourceLayers](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSubresourceLayers.html),
[VkOffset3D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset3D.html), [vkCmdCopyImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImage.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageCopy"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
