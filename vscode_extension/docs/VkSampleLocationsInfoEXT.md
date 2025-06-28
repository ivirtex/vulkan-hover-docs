# VkSampleLocationsInfoEXT(3) Manual Page

## Name

VkSampleLocationsInfoEXT - Structure specifying a set of sample locations



## [](#_c_specification)C Specification

The `VkSampleLocationsInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_sample_locations
typedef struct VkSampleLocationsInfoEXT {
    VkStructureType               sType;
    const void*                   pNext;
    VkSampleCountFlagBits         sampleLocationsPerPixel;
    VkExtent2D                    sampleLocationGridSize;
    uint32_t                      sampleLocationsCount;
    const VkSampleLocationEXT*    pSampleLocations;
} VkSampleLocationsInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `sampleLocationsPerPixel` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value specifying the number of sample locations per pixel.
- `sampleLocationGridSize` is the size of the sample location grid to select custom sample locations for.
- `sampleLocationsCount` is the number of sample locations in `pSampleLocations`.
- `pSampleLocations` is a pointer to an array of `sampleLocationsCount` [VkSampleLocationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationEXT.html) structures.

## [](#_description)Description

This structure **can** be used either to specify the sample locations to be used for rendering or to specify the set of sample locations an image subresource has been last rendered with for the purposes of layout transitions of depth/stencil images created with `VK_IMAGE_CREATE_SAMPLE_LOCATIONS_COMPATIBLE_DEPTH_BIT_EXT`.

The sample locations in `pSampleLocations` specify `sampleLocationsPerPixel` number of sample locations for each pixel in the grid of the size specified in `sampleLocationGridSize`. The sample location for sample i at the pixel grid location (x,y) is taken from `pSampleLocations`\[(x + y × `sampleLocationGridSize.width`) × `sampleLocationsPerPixel` + i].

If the render pass has a fragment density map, the implementation will choose the sample locations for the fragment and the contents of `pSampleLocations` **may** be ignored.

Valid Usage

- [](#VUID-VkSampleLocationsInfoEXT-sampleLocationsPerPixel-01526)VUID-VkSampleLocationsInfoEXT-sampleLocationsPerPixel-01526  
  `sampleLocationsPerPixel` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value that is set in [VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html)::`sampleLocationSampleCounts`
- [](#VUID-VkSampleLocationsInfoEXT-sampleLocationsCount-01527)VUID-VkSampleLocationsInfoEXT-sampleLocationsCount-01527  
  `sampleLocationsCount` **must** equal `sampleLocationsPerPixel` × `sampleLocationGridSize.width` × `sampleLocationGridSize.height`

Valid Usage (Implicit)

- [](#VUID-VkSampleLocationsInfoEXT-sType-sType)VUID-VkSampleLocationsInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLE_LOCATIONS_INFO_EXT`
- [](#VUID-VkSampleLocationsInfoEXT-pSampleLocations-parameter)VUID-VkSampleLocationsInfoEXT-pSampleLocations-parameter  
  If `sampleLocationsCount` is not `0`, `pSampleLocations` **must** be a valid pointer to an array of `sampleLocationsCount` [VkSampleLocationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationEXT.html) structures

## [](#_see_also)See Also

[VK\_EXT\_sample\_locations](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sample_locations.html), [VkAttachmentSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentSampleLocationsEXT.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html), [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html), [VkSampleLocationEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleLocationEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubpassSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassSampleLocationsEXT.html), [vkCmdSetSampleLocationsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetSampleLocationsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSampleLocationsInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0