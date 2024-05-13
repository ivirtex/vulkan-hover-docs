# VkPhysicalDeviceYcbcrImageArraysFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceYcbcrImageArraysFeaturesEXT - Structure describing
extended Y′C<sub>B</sub>C<sub>R</sub> image creation features that can
be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceYcbcrImageArraysFeaturesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_ycbcr_image_arrays
typedef struct VkPhysicalDeviceYcbcrImageArraysFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           ycbcrImageArrays;
} VkPhysicalDeviceYcbcrImageArraysFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-ycbcrImageArrays"></span> `ycbcrImageArrays`
  indicates that the implementation supports creating images with a
  format that requires <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
  target="_blank" rel="noopener">Y′C<sub>B</sub>C<sub>R</sub>
  conversion</a> and has multiple array layers.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceYcbcrImageArraysFeaturesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceYcbcrImageArraysFeaturesEXT` **can** also be used in
the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceYcbcrImageArraysFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceYcbcrImageArraysFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceYcbcrImageArraysFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_YCBCR_IMAGE_ARRAYS_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_ycbcr_image_arrays](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_ycbcr_image_arrays.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceYcbcrImageArraysFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
