# VkPhysicalDeviceSampleLocationsPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceSampleLocationsPropertiesEXT - Structure describing
sample location limits that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceSampleLocationsPropertiesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_sample_locations
typedef struct VkPhysicalDeviceSampleLocationsPropertiesEXT {
    VkStructureType       sType;
    void*                 pNext;
    VkSampleCountFlags    sampleLocationSampleCounts;
    VkExtent2D            maxSampleLocationGridSize;
    float                 sampleLocationCoordinateRange[2];
    uint32_t              sampleLocationSubPixelBits;
    VkBool32              variableSampleLocations;
} VkPhysicalDeviceSampleLocationsPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-sampleLocationSampleCounts"></span>
  `sampleLocationSampleCounts` is a bitmask of
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) indicating the
  sample counts supporting custom sample locations.

- <span id="limits-maxSampleLocationGridSize"></span>
  `maxSampleLocationGridSize` is the maximum size of the pixel grid in
  which sample locations **can** vary that is supported for all sample
  counts in `sampleLocationSampleCounts`.

- <span id="limits-sampleLocationCoordinateRange"></span>
  `sampleLocationCoordinateRange`\[2\] is the range of supported sample
  location coordinates.

- <span id="limits-sampleLocationSubPixelBits"></span>
  `sampleLocationSubPixelBits` is the number of bits of subpixel
  precision for sample locations.

- <span id="limits-variableSampleLocations"></span>
  `variableSampleLocations` specifies whether the sample locations used
  by all pipelines that will be bound to a command buffer during a
  subpass **must** match. If set to `VK_TRUE`, the implementation
  supports variable sample locations in a subpass. If set to `VK_FALSE`,
  then the sample locations **must** stay constant in each subpass.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceSampleLocationsPropertiesEXT` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceSampleLocationsPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceSampleLocationsPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceSampleLocationsPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLE_LOCATIONS_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_sample_locations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sample_locations.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html),
[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceSampleLocationsPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
