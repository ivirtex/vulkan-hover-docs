# VkSurfaceFormat2KHR(3) Manual Page

## Name

VkSurfaceFormat2KHR - Structure describing a supported swapchain format
tuple



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSurfaceFormat2KHR` structure is defined as:

``` c
// Provided by VK_KHR_get_surface_capabilities2
typedef struct VkSurfaceFormat2KHR {
    VkStructureType       sType;
    void*                 pNext;
    VkSurfaceFormatKHR    surfaceFormat;
} VkSurfaceFormat2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `surfaceFormat` is a [VkSurfaceFormatKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFormatKHR.html)
  structure describing a format-color space pair that is compatible with
  the specified surface.

## <a href="#_description" class="anchor"></a>Description

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-imageCompressionControlSwapchain"
target="_blank"
rel="noopener"><code>imageCompressionControlSwapchain</code></a> feature
is supported and a
[VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionPropertiesEXT.html)
structure is included in the `pNext` chain of this structure, then it
will be filled with the compression properties that are supported for
the `surfaceFormat`.

Valid Usage

- <a href="#VUID-VkSurfaceFormat2KHR-pNext-06750"
  id="VUID-VkSurfaceFormat2KHR-pNext-06750"></a>
  VUID-VkSurfaceFormat2KHR-pNext-06750  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-imageCompressionControlSwapchain"
  target="_blank"
  rel="noopener"><code>imageCompressionControlSwapchain</code></a>
  feature is not enabled, the `pNext` chain **must** not include an
  [VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionPropertiesEXT.html)
  structure

Valid Usage (Implicit)

- <a href="#VUID-VkSurfaceFormat2KHR-sType-sType"
  id="VUID-VkSurfaceFormat2KHR-sType-sType"></a>
  VUID-VkSurfaceFormat2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_FORMAT_2_KHR`

- <a href="#VUID-VkSurfaceFormat2KHR-pNext-pNext"
  id="VUID-VkSurfaceFormat2KHR-pNext-pNext"></a>
  VUID-VkSurfaceFormat2KHR-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkImageCompressionPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionPropertiesEXT.html)

- <a href="#VUID-VkSurfaceFormat2KHR-sType-unique"
  id="VUID-VkSurfaceFormat2KHR-sType-unique"></a>
  VUID-VkSurfaceFormat2KHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_get_surface_capabilities2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_surface_capabilities2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSurfaceFormatKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceFormatKHR.html),
[vkGetPhysicalDeviceSurfaceFormats2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceFormats2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfaceFormat2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
