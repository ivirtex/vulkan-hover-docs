# VkSampleLocationEXT(3) Manual Page

## Name

VkSampleLocationEXT - Structure specifying the coordinates of a sample
location



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSampleLocationEXT` structure is defined as:

``` c
// Provided by VK_EXT_sample_locations
typedef struct VkSampleLocationEXT {
    float    x;
    float    y;
} VkSampleLocationEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `x` is the horizontal coordinate of the sample’s location.

- `y` is the vertical coordinate of the sample’s location.

## <a href="#_description" class="anchor"></a>Description

The domain space of the sample location coordinates has an upper-left
origin within the pixel in framebuffer space.

The values specified in a `VkSampleLocationEXT` structure are always
clamped to the implementation-dependent sample location coordinate range
\[`sampleLocationCoordinateRange`\[0\],`sampleLocationCoordinateRange`\[1\]\]
that **can** be queried using
[VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html).

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_sample_locations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sample_locations.html),
[VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSampleLocationEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
