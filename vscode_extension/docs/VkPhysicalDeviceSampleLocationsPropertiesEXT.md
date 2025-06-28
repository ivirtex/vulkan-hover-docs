# VkPhysicalDeviceSampleLocationsPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceSampleLocationsPropertiesEXT - Structure describing sample location limits that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceSampleLocationsPropertiesEXT` structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`sampleLocationSampleCounts` is a bitmask of [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) indicating the sample counts supporting custom sample locations.
- []()`maxSampleLocationGridSize` is the maximum size of the pixel grid in which sample locations **can** vary that is supported for all sample counts in `sampleLocationSampleCounts`.
- []()`sampleLocationCoordinateRange`\[2] is the range of supported sample location coordinates.
- []()`sampleLocationSubPixelBits` is the number of bits of subpixel precision for sample locations.
- []()`variableSampleLocations` specifies whether the sample locations used by all pipelines that will be bound to a command buffer during a subpass **must** match. If set to `VK_TRUE`, the implementation supports variable sample locations in a subpass. If set to `VK_FALSE`, then the sample locations **must** stay constant in each subpass.

## [](#_description)Description

If the `VkPhysicalDeviceSampleLocationsPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceSampleLocationsPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceSampleLocationsPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLE_LOCATIONS_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_sample\_locations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sample_locations.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceSampleLocationsPropertiesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0