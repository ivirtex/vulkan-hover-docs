# VkSampleLocationEXT(3) Manual Page

## Name

VkSampleLocationEXT - Structure specifying the coordinates of a sample location



## [](#_c_specification)C Specification

The `VkSampleLocationEXT` structure is defined as:

```c++
// Provided by VK_EXT_sample_locations
typedef struct VkSampleLocationEXT {
    float    x;
    float    y;
} VkSampleLocationEXT;
```

## [](#_members)Members

- `x` is the horizontal coordinate of the sample’s location.
- `y` is the vertical coordinate of the sample’s location.

## [](#_description)Description

The domain space of the sample location coordinates has an upper-left origin within the pixel in framebuffer space.

The values specified in a `VkSampleLocationEXT` structure are always clamped to the implementation-dependent sample location coordinate range \[`sampleLocationCoordinateRange`\[0],`sampleLocationCoordinateRange`\[1]] that **can** be queried using [VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html).

## [](#_see_also)See Also

[VK\_EXT\_sample\_locations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sample_locations.html), [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationsInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSampleLocationEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0